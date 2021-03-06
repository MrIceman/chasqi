package processor

import (
	"fmt"
	"testing"
)

func TestGetNavigationTree(t *testing.T) {
	filePath := "./fixture/test_nav.yaml"

	tree := GetNavigationTree(filePath)
	if tree.Host != "http://www.loremio.com" {
		println(tree.Host)
		t.Error("Failed to parse correct host")
	}

	if len(tree.Nodes) != 4 {
		t.Error("Invalid amount of Nodes")
	}
	// TODO write more fancy test
	t0 := tree.Nodes[0]
	validateValue(t0.Name, "create_account", t)
	validateValue(t0.Method, "POST", t)

	for k, v := range t0.RequestBody {
		println(k, ": ", fmt.Sprint(v))
	}

	t1 := tree.Nodes[1]
	for k, v := range t1.RequestBody {
		println(k, ": ", fmt.Sprint(v))
	}

	for _, v := range tree.Nodes {
		if v.Name == "get_subscription" {
			// Check that returns array is false
			if v.ReturnsArray == false {
				t.Error("Returns array for get_subscription is false")
			}
		} else {
			if v.ReturnsArray != false {
				t.Error("Returns array for " + v.Name + " is not false")
			}
		}
	}
}

func TestShouldSetRandomValuesForEmailAndPassword(t *testing.T) {
	filePath := "./fixture/test_nav.yaml"

	tree := GetNavigationTree(filePath)

	for _, v := range tree.Variables {
		if v.Name == "email" {
			println(v.Value)
			// 6 random chars + "@" + 8 random chars + ".de" == 18 chars
			if len(v.Value) != 18 {
				t.Error("Email doesn't have 18 chars")
			}
		}
		if v.Name == "password" {
			println(v.Value)
			if len(v.Value) != 8 {
				t.Error("Password doesn't have 8 chars")
			}
		}
	}
}

func validateValue(v0 string, v1 string, t *testing.T) {
	if v0 != v1 {
		t.Error(v0, " is not ", v1)
	}
}
