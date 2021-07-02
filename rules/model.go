package rules

import "fmt"

type NavigationTree struct {
	Host           string      `yaml:"host"`
	AmountOfAgents int         `yaml:"agents"`
	Nodes          []Node      `yaml:"nodes"`
	Variables      []*Variable `yaml:"variables"`
}

type Variable struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
	Type  string `yaml:"dt"`
}
type Node struct {
	Method          string                 `yaml:"method"`
	ExposesResponse []string               `yaml:"exposes"`
	Name            string                 `yaml:"name"`
	RequestBody     map[string]interface{} `yaml:"body"`
	RequestHeaders  map[string]interface{} `yaml:"header"`
	Path            string                 `yaml:"path"`
	ReturnsArray    bool                   `yaml:"returns_array"`
}

func (n *Node) GetRequestBodyAsStringMap() map[string]string {
	strValueMap := make(map[string]string)

	for k, v := range n.RequestBody {
		strValueMap[k] = fmt.Sprint(v)
	}

	return strValueMap
}
