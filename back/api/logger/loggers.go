package logger

import "log"

var flags = log.LstdFlags | log.Lshortfile
var (
	Debug = log.New(log.Writer(), "[DEBUG] ", flags)
	Info  = log.New(log.Writer(), "[INFO] ", flags)
	Warn  = log.New(log.Writer(), "[WARN] ", flags)
	Error = log.New(log.Writer(), "[ERROR]  ", flags)
)
