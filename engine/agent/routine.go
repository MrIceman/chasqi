package agent

const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Route struct {
	level                 int
	previous              *Route
	next                  *Route
	path                  string
	response              map[string]string
	exposesResponseValues []string
	isCalled              bool
	method                string
	headers               map[string]interface{}
	name                  string
	body                  map[string]interface{}
	returnsArray          bool
}

func (r *Route) GetValuesFromPrevResponse(key string) string {
	if r.previous == nil {
		println("NO previous value")
	}

	for _, value := range r.previous.exposesResponseValues {
		if value == key {
			return r.previous.response[value]
		}
	}

	return r.previous.GetValuesFromPrevResponse(key)
}
