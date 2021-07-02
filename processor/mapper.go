package processor

import (
	"chasqi/rules"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

func GetNavigationTree(filePath string) *rules.NavigationTree {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	tree := rules.NavigationTree{}
	err = yaml.Unmarshal(file, &tree)

	for _, n := range tree.Variables {
		if strings.Contains(n.Value, "random::") {
			n.Value = ReplaceRandomsInString(n.Value)
		}
	}

	if err != nil {
		panic(err)
	}

	return &tree
}
