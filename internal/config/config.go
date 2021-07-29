package config

import (
	"flag"
)

type CommandLine struct {
	Addr        string
	DBPath      string
	LogFormat   string
	LogLevel    string
	MetricsAddr string
}

func ParseCommandLine(progamName string, args ...string) (*CommandLine, error) {
	cl := &CommandLine{}

	fs := flag.NewFlagSet(progamName, flag.ContinueOnError)

	fs.StringVar(&cl.Addr, "addr", ":8080", "the address on which the app should listen")
	fs.StringVar(&cl.DBPath, "db-path", "db.sqlite3", "path to the sqlite database")
	fs.StringVar(&cl.LogFormat, "log-format", "", "the format to use when logging messages")
	fs.StringVar(&cl.LogLevel, "log-level", "info", "the log level to use - increase for more verbosity")
	fs.StringVar(&cl.MetricsAddr, "metrics-addr", ":9090", "the address on which the app should listen for metrics")

	return cl, fs.Parse(args)
}
