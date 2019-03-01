package log

import (
	"log"
	"os"
)

var (
	Error = log.New(os.Stderr, "ERROR ", log.LstdFlags)
	Info  = log.New(os.Stdout, "INFO ", log.LstdFlags)
)
