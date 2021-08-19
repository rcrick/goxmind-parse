package v2

import (
	"archive/zip"
	"encoding/json"
	"io/ioutil"
)

type Sheet struct {
	ID       string `json:"id"`
	Topic    Topic  `json:"rootTopic"`
	Title    string `json:"title"`
	Class    string `json:"class"`
	Position string `json:"topicPositioning"`
}

type Topic struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Class    string   `json:"class"`
	Notes    Notes    `json:"notes"`
	Labels   []string `json:"labels"`
	Children Children `json:"children"`
}

type Children struct {
	Attached []Attached `json:"attached"`
}

type Attached struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Notes    Notes    `json:"notes"`
	Labels   []string `json:"labels"`
	Children Children `json:"children"`
}

type Notes struct {
	Plain Plain `json:"plain"`
}

type Plain struct {
	Content string `json:"content"`
}

func Parse(file string) []Sheet {
	zipFile, err := zip.OpenReader(file)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	contentFile := new(zip.File)
	for _, f := range zipFile.File {
		switch f.Name {
		case "content.json":
			contentFile = f
		}
	}
	x := []Sheet{}
	rc, _ := contentFile.Open()
	content, err := ioutil.ReadAll(rc)
	json.Unmarshal(content, &x)
	//fmt.Printf("%+v\n", x)
	//
	//empJSON, _ := json.MarshalIndent(x, "", "  ")
	//fmt.Printf("MarshalIndent funnction output \n%s\n", string(empJSON))

	return x
}
