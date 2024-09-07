package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/chrisbradleydev/go-merge/pkg/utils"
	"gopkg.in/yaml.v3"
)

func main() {
	c1 := []byte(`name: start
list:
- name: one
  over: abc
  prev: prev1
- name: two
  over: abc
  prev: prev2
- name: three
  over: abc
  prev: prev3`)
	c2 := []byte(`name: finish
list:
- name: one
  over: xyz
- name: two
  over: xyz
- name: three
  over: xyz`)

	map1 := map[string]interface{}{}
	if err := yaml.Unmarshal(c1, &map1); err != nil {
		log.Fatal(err)
	}
	map2 := map[string]interface{}{}
	if err := yaml.Unmarshal(c2, &map2); err != nil {
		log.Fatal(err)
	}

	newMap := utils.MergeMaps(map1, map2)
	merged, err := yaml.Marshal(newMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(merged))
}
