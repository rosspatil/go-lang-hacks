package main

import (
	"github.com/Jeffail/gabs"
	"github.com/tidwall/sjson"
)

func main() {

}

func SetInGabs() string {

	c, err := gabs.New().Set(map[string]interface{}{"name": "hello"}, "test.child")
	if err != nil {
		panic(err)
	}
	return c.String()
}

func SetInGjson() string {
	s, err := sjson.Set("{}", "test.child", map[string]interface{}{"name": "hello"})
	if err != nil {
		panic(err)
	}
	return s
}
