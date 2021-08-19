package v1

import (
	"archive/zip"
	"encoding/xml"
	"io"
	"io/ioutil"
)

type Xmap struct {
	XMLName xml.Name `xml:"xmap-content"`
	Sheet   Sheet    `xml:"sheet"`
}

type Sheet struct {
	XMLName xml.Name `xml:"sheet"`
	Topic   Topic    `xml:"topic"`
}

type Topic struct {
	XMLName xml.Name `xml:"topic"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"title"`

	Note  string `xml:"notes>plain"`
	Label string `xml:"labels>label"`
	Html  string `xml:"notes>html"`

	Children []Topic `xml:"children>topics>topic"`
}

type Comments struct {
	XMLName xml.Name  `xml:"comments"`
	Child   []Comment `xml:"comment"`
}

type Comment struct {
	Content string `xml:"content"`
	ID      string `xml:"object-id,attr"`
}

func Parse(file string) *Xmap {
	zipFile, err := zip.OpenReader(file)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	contentFile := new(zip.File)
	commentsFile := new(zip.File)
	for _, f := range zipFile.File {
		switch f.Name {
		case "content.xml":
			contentFile = f
		case "comments.xml":
			commentsFile = f

		}
	}
	x := Xmap{}
	c := Comments{}
	parseXml(contentFile, &x)
	parseXml(commentsFile, &c)

	return &x
}

func parseXml(file *zip.File, data interface{}) {
	rc, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer rc.Close()

	buildXMlTree(rc, data)
}

func buildXMlTree(rc io.ReadCloser, data interface{}) {
	defer rc.Close()

	content, err := ioutil.ReadAll(rc)
	if err != nil {
		panic(err)
	}

	err = xml.Unmarshal(content, data)
	if err != nil {
		panic(err)
	}
}
