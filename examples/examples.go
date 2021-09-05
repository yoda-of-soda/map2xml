package main

import (
	"fmt"

	"github.com/yoda-of-soda/map2xml"
)

func main() {
	inputMap := map[string]interface{}{
		"first_name": "No",
		"last_name":  "Name",
		"age":        42,
		"got_a_job":  true,
		"secret":     nil,
		"address": map[string]interface{}{
			"street":   "124 Oxford Street",
			"city":     "Somewhere",
			"zip_code": 32784,
			"state":    "Deep state",
		},
	}

	config := map2xml.New(inputMap)
	config.WithIndent("", "  ")
	config.AsCData()
	config.WithRoot("person", map[string]string{"mood": "happy"})
	config.Print()

	xmlBytes, err := config.Marshal()
	fmt.Println(string(xmlBytes))

	str, err := map2xml.New(inputMap).AsCData().WithIndent("", "  ").WithRoot("person", map[string]string{"mood": "happy"}).Print().MarshalToString()
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
