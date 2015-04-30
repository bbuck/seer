package log

type LogLevel int

func (l LogLevel) String() string {
	switch l {
	case LogError:
		return "error"
	case LogInfo:
		return "info"
	case LogVerbose:
		return "verbose"
	case LogDebug:
		return "debug"
	default:
		return "INVALID"
	}
}

const (
	LogError LogLevel = iota
	LogInfo
	LogVerbose
	LogDebug
)

var MaxLogLevel LogLevel = LogDebug
