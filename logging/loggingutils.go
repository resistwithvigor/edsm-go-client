package logging

import (
	"io"
	"log"
)

func Init(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	const (
		FLAGS = log.Ldate | log.Ltime | log.Lshortfile
	)

	Trace = log.New(traceHandle, "TRACE: ", FLAGS)
	Info = log.New(infoHandle, "INFO: ", FLAGS)
	Warning = log.New(warningHandle, "WARNING: ", FLAGS)
	Error = log.New(errorHandle, "ERROR: ", FLAGS)
}
