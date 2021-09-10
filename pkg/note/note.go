package note

import (
	"encoding/json"
	"os"
)

type Note struct {
	Meta    NoteMeta
	Content []byte
}

type NoteMeta struct {
	Id             string
	Title          string
	Tags           []string
	LastEditedTime string
	Status         string
	IsPrivate      bool
}

func (n Note) DumpMetadata(folder string) {
	meta, _ := json.Marshal(n.Meta)
	_ = os.WriteFile(folder+"/meta/"+n.Meta.Id+".json", meta, 0644)
}

func (n Note) DumpContent(folder string) {
	_ = os.WriteFile(folder+"/content/"+n.Meta.Id+".md", n.Content, 0644)
}

func (n Note) Dump(folder string) {
	n.DumpMetadata(folder)
	n.DumpContent(folder)
}
