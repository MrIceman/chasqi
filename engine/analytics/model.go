package analytics

type LogEntry struct {
	StatusCode       int
	Path             string
	ResponseTimeInMs int
	ResponseBody     string
	Method           string
}
