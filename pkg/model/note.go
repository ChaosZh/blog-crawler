package model

import (
	"encoding/json"
	"os"

	"github.com/chaoszh/blog-crawler/util"
)

type Note struct {
	Meta
	Content []byte
}

type Meta struct {
	Id             string
	Alias          string
	Title          string
	Tags           []string
	LastEditedTime string
	Status         string
	IsPrivate      bool
}

func (n *Note) Standardize() {
	n.Alias = util.GenAliasFor(n.Title)
}

func (n *Note) Dump(folder string) {
	n.Standardize()
	n.DumpMetadata(folder)
	n.DumpContent(folder)
}

func (n *Note) DumpMetadata(folder string) {
	meta, _ := json.Marshal(n.Meta)
	_ = os.WriteFile(folder+"/meta/"+n.Alias+".json", meta, 0644) // make sure that the folder path has been created
}

func (n *Note) DumpContent(folder string) {
	_ = os.WriteFile(folder+"/content/"+n.Alias+".md", n.Content, 0644)
}
