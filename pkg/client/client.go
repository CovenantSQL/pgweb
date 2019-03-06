package client

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/CovenantSQL/pgweb/pkg/command"
	"github.com/CovenantSQL/pgweb/pkg/connection"
	"github.com/CovenantSQL/pgweb/pkg/history"
	"github.com/CovenantSQL/pgweb/pkg/shared"
	"github.com/CovenantSQL/pgweb/pkg/statements"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	_ "github.com/CovenantSQL/CovenantSQL/client"
)

type Client struct {
	db               *sqlx.DB
	tunnel           *Tunnel
	lastQueryTime    time.Time
	External         bool
	History          []history.Record `json:"history"`
	ConnectionString string           `json:"connection_string"`
}

// Struct to hold table rows browsing options
type RowsOptions struct {
	Where      string // Custom filter
	Offset     int    // Number of rows to skip
	Limit      int    // Number of rows to fetch
	SortColumn string // Column to sort by
	SortOrder  string // Sort direction (ASC, DESC)
}

func New() (*Client, error) {
	str, err := connection.BuildStringFromOptions(command.Opts)

	if command.Opts.Debug && str != "" {
		fmt.Println("Creating a new client for:", str)
	}

	if err != nil {
		return nil, err
	}

	db, err := sqlx.Open("covenantsql", str)
	if err != nil {
		return nil, err
	}

	client := Client{
		db:               db,
		ConnectionString: str,
		History:          history.New(),
	}

	return &client, nil
}

func NewFromUrl(url string, sshInfo *shared.SSHInfo) (*Client, error) {
	var tunnel *Tunnel

	if sshInfo != nil {
		return nil, errors.New("ssh mode is not supported")
	}

	if command.Opts.Debug {
		fmt.Println("Creating a new client for:", url)
	}

	db, err := sqlx.Open("covenantsql", url)
	if err != nil {
		return nil, err
	}

	client := Client{
		db:               db,
		tunnel:           tunnel,
		ConnectionString: url,
		History:          history.New(),
	}

	return &client, nil
}

func (client *Client) Test() error {
	return client.db.Ping()
}

func (client *Client) Info() (*Result, error) {
	return client.query(statements.Info)
}

func (client *Client) Databases() ([]string, error) {
	return client.fetchRows(statements.Databases)
}

func (client *Client) Schemas() ([]string, error) {
	return client.fetchRows(statements.Schemas)
}

func (client *Client) Objects() (*Result, error) {
	return client.query(statements.Objects)
}

func (client *Client) Table(table string) (*Result, error) {
	schemaResult, err := client.query(statements.TableSchema + table)
	if err != nil {
		return nil, err
	}

	schemaResult.Columns = []string{
		"column_name",
		"data_type",
		"is_nullable",
		"column_default",
		"pk",
	}

	for i := range schemaResult.Rows {
		// covenantsql returns: cid,name,type,notnull,dflt_value,pk
		schemaResult.Rows[i] = Row{
			schemaResult.Rows[i][1],                           // name
			schemaResult.Rows[i][2],                           // type
			fmt.Sprintf("%v", schemaResult.Rows[i][3]) == "0", // is_nullable
			schemaResult.Rows[i][4],                           // dflt_value
			schemaResult.Rows[i][5],                           // pk
		}
	}

	return schemaResult, nil
}

func (client *Client) MaterializedView(name string) (*Result, error) {
	// not supported
	return nil, nil
}

func (client *Client) TableRows(table string, opts RowsOptions) (*Result, error) {
	sql := fmt.Sprintf(`SELECT * FROM %s`, table)

	if opts.Where != "" {
		sql += fmt.Sprintf(" WHERE %s", opts.Where)
	}

	if opts.SortColumn != "" {
		if opts.SortOrder == "" {
			opts.SortOrder = "ASC"
		}

		sql += fmt.Sprintf(` ORDER BY %s %s`, opts.SortColumn, opts.SortOrder)
	}

	if opts.Limit > 0 {
		sql += fmt.Sprintf(" LIMIT %d", opts.Limit)
	}

	if opts.Offset > 0 {
		sql += fmt.Sprintf(" OFFSET %d", opts.Offset)
	}

	return client.query(sql)
}

func (client *Client) TableRowsCount(table string, opts RowsOptions) (*Result, error) {
	sql := fmt.Sprintf(`SELECT COUNT(1) FROM %s`, table)

	if opts.Where != "" {
		sql += fmt.Sprintf(" WHERE %s", opts.Where)
	}

	return client.query(sql)
}

func (client *Client) TableInfo(table string) (*Result, error) {
	return client.query(statements.TableInfo + table)
}

func (client *Client) TableIndexes(table string) (*Result, error) {
	res, err := client.query(statements.TableIndexes + table)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (client *Client) TableConstraints(table string) (*Result, error) {
	// not supported
	return nil, nil
}

// Returns all active queriers on the server
func (client *Client) Activity() (*Result, error) {
	// not supported
	return nil, nil
}

func (client *Client) Query(query string) (*Result, error) {
	res, err := client.query(query)

	// Save history records only if query did not fail
	if err == nil && !client.hasHistoryRecord(query) {
		client.History = append(client.History, history.NewRecord(query))
	}

	return res, err
}

func (client *Client) ServerVersion() string {
	return "CovenantSQL/1.0"
}

func (client *Client) query(query string, args ...interface{}) (*Result, error) {
	// Update the last usage time
	defer func() {
		client.lastQueryTime = time.Now().UTC()
	}()

	action := strings.ToLower(strings.Split(query, " ")[0])
	if action == "update" || action == "delete" ||
		action == "create" || action == "insert" ||
		action == "replace" {
		res, err := client.db.Exec(query, args...)
		if err != nil {
			return nil, err
		}

		affected, err := res.RowsAffected()
		if err != nil {
			return nil, err
		}

		result := Result{
			Columns: []string{"Rows Affected"},
			Rows: []Row{
				{affected},
			},
		}

		return &result, nil
	}

	rows, err := client.db.Queryx(query, args...)
	if err != nil {
		if command.Opts.Debug {
			log.Println("Failed query:", query, "\nArgs:", args)
		}
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Make sure to never return null colums
	if cols == nil {
		cols = []string{}
	}

	result := Result{
		Columns: cols,
		Rows:    []Row{},
	}

	for rows.Next() {
		obj, err := rows.SliceScan()

		for i, item := range obj {
			if item == nil {
				obj[i] = nil
			} else {
				t := reflect.TypeOf(item).Kind().String()

				if t == "slice" {
					obj[i] = string(item.([]byte))
				}
			}
		}

		if err == nil {
			result.Rows = append(result.Rows, obj)
		}
	}

	result.PrepareBigints()

	return &result, nil
}

// Close database connection
func (client *Client) Close() error {
	if client.tunnel != nil {
		client.tunnel.Close()
	}

	if client.db != nil {
		return client.db.Close()
	}

	return nil
}

func (client *Client) IsIdle() bool {
	mins := int(time.Since(client.lastQueryTime).Minutes())

	if command.Opts.ConnectionIdleTimeout > 0 {
		return mins >= command.Opts.ConnectionIdleTimeout
	}

	return false
}

// Fetch all rows as strings for a single column
func (client *Client) fetchRows(q string) ([]string, error) {
	res, err := client.query(q)

	if err != nil {
		return nil, err
	}

	// Init empty slice so json.Marshal will encode it to "[]" instead of "null"
	results := make([]string, 0)

	for _, row := range res.Rows {
		results = append(results, row[0].(string))
	}

	return results, nil
}

func (client *Client) hasHistoryRecord(query string) bool {
	result := false

	for _, record := range client.History {
		if record.Query == query {
			result = true
			break
		}
	}

	return result
}
