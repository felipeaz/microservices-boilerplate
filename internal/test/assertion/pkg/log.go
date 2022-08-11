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

	LogTime = time.Date(1997, 12, 31, 10, 00, 00, 00, time.UTC)
	LogFile = fmt.Sprintf("%s.txt", LogTime.Format("01-02-2006"))

	LogDir = "logs"
)
