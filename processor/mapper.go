package processor

import (
	"chasqi/rules"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func GetNavigationTree(filePath string) *rules.NavigationTree {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	tree := rules.NavigationTree{}
	err = yaml.Unmarshal(file, &tree)

	if err != nil {
		panic(err)
	}

	return &tree
}
