package pkg

import (
	"fmt"
	"time"
)

var (
	InfoLogMessage  = "Info message"
	WarnLogMessage  = "Warn message"
	ErrLogMessage   = "Err message"
	DebugLogMessage = "Debug message"
	LogDir          = "logs"

	LogTime     = time.Now().UTC()
	LogFileName = fmt.Sprintf("%s.txt", LogTime.Format("01-02-2006"))
)
