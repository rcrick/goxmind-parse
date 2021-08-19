package main

import (
	"encoding/json"
	"example/xmind/v1"
	"fmt"
)

func main() {
	f := "xmindold/xmind_old.xmind"
	tree := v1.Parse(f)
	empJSON, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Printf("MarshalIndent funnction output \n%s\n", string(empJSON))
}
