package main

import (
	"fmt"

	"github.com/JohnAllanNielsen/map2xml"
)

func main() {
	ori := map[string]interface{}{
		"a": 1,
		"b": "abekat",
		"tivoli": map[string]interface{}{
			"int":  42,
			"crap": false,
		},
	}
	awesome := map2xml.New(ori, "goodfellas").AsCData().WithIndent("", "  ").WithStartAttributes(map[string]string{"status": "awesome"})
	str, err := awesome.MarshalToString()
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
