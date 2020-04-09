package test

import (
	stdlog "log"
	"os"
)

type Mytest struct {
	stderr *stdlog.Logger
	stdout *stdlog.Logger
}

func CreateLogger() *stdLogger {
	return &Mytest{
		stdout: stdlog.New(os.Stdout, "", 0),
		stderr: stdlog.New(os.Stderr, "", 0),
	}
}
