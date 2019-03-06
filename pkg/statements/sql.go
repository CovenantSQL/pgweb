package statements

const (
	Databases = `SELECT "main"`

	// ---------------------------------------------------------------------------

	Schemas = `SHOW TABLES`

	// ---------------------------------------------------------------------------

	Info = `SELECT "ok" AS info`

	// ---------------------------------------------------------------------------

	TableIndexes = `SHOW INDEX FROM TABLE `

	// ---------------------------------------------------------------------------

	TableInfo = `SELECT COUNT(1) AS rows_count FROM `

	// ---------------------------------------------------------------------------

	TableSchema = `DESC `

	// ---------------------------------------------------------------------------

	Objects = `
SELECT
  "main" as "schema",
  name as "name",
  type as "type"
FROM
  sqlite_master`
)
