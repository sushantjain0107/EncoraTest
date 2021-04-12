package main

import (
	"encoding/json"
	"log"
)

func parse(v string) (*node, error) {
	root := &node{}

	// Parse here

	return root, nil
}

type node struct {
	Name     string  `json:"name"`
	Children []*node `json:"children,omitempty"`
}

var examples = []string{
	"[a,b,c]",
	"[a[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]",
}

func main() {
	for i, example := range examples {
		result, err := parse(example)
		if err != nil {
			panic(err)
		}
		j, err := json.MarshalIndent(result, " ", " ")
		if err != nil {
			panic(err)
		}
		log.Printf("Example %d: %s - %s", i, example, string(j))
	}
}
