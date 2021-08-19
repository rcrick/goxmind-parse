package main

import (
	"encoding/json"
	"example/xmind/v2"
	"fmt"
)

func main() {
	f := "xmindnew/xmindnew.xmind"
	tree := v2.Parse(f)
	empJSON, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Printf("MarshalIndent funnction output \n%s\n", string(empJSON))
}
