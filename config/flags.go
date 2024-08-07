package config

import "github.com/urfave/cli/v2"

var (
	LogFlags = []cli.Flag{
		&cli.StringFlag{
			Name:    "agsf-log-dir-path",
			EnvVars: []string{"AGSF_LOG_DIR_PATH"},
			Usage:   "Directory path for log files",
		},
	}
	DatabaseFlags = []cli.Flag{
		&cli.StringFlag{
			Name:    "agsf-db-base-path",
			EnvVars: []string{"AGSF_DB_BASE_PATH"},
			Usage:   "Base directory path for sqlite file",
		},
	}
)

func ConcatFlags(flagGroups ...[]cli.Flag) []cli.Flag {
	var result []cli.Flag
	for _, flags := range flagGroups {
		result = append(result, flags...)
	}
	return result
}
