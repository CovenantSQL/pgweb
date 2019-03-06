package connection

import (
	neturl "net/url"

	"github.com/CovenantSQL/pgweb/pkg/command"

	"github.com/CovenantSQL/CovenantSQL/client"
)

// FormatURL reformats the existing connection string
func FormatURL(opts command.Options) (string, error) {
	url := opts.Url

	cfg, err := client.ParseDSN(url)
	if err != nil {
		return "", err
	}

	return cfg.FormatDSN(), nil
}

// IsBlank returns true if command options do not contain connection details
func IsBlank(opts command.Options) bool {
	return opts.DbName == "" && opts.Url == ""
}

// BuildStringFromOptions returns a new connection string built from options
func BuildStringFromOptions(opts command.Options) (string, error) {
	// If connection string is provided we just use that
	if opts.Url != "" {
		return FormatURL(opts)
	}

	url := neturl.URL{
		Scheme: "covenantsql",
		Host:   opts.DbName,
	}

	return url.String(), nil
}
