package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	note "github.com/chaoszh/blog-crawler/pkg/model"
	"github.com/kjk/notionapi"
	"github.com/kjk/notionapi/tomarkdown"
)

var client = &notionapi.Client{}

func logger(message string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), message)
}

var xmapper = map[string]string{
	"Title":     "title",
	"Tags":      "OpfN",
	"Status":    "Nm\\r",
	"isPrivate": "__lm",
}

func getNotes(pageid string) ([]note.Note, error) {
	logger("get all notes from notion")

	notes := make([]note.Note, 0)

	archivedPage, err := client.DownloadPage(pageid)
	if err != nil {
		log.Fatalf("Failed at %s\n", err)
		return nil, err
	}

	catalogue := archivedPage.TableViews[0]
	for _, row := range catalogue.Rows {
		meta := note.Meta{}
		meta.Id = row.Page.GetNotionID().NoDashID
		meta.LastEditedTime = time.Unix(0, row.Page.LastEditedTime*int64(time.Millisecond)).Format("2006-01-02 15:04:05")

		//extract data from notion result
		for k, v := range xmapper {
			property := row.Page.Properties[v]
			if property != nil {
				property = property.([]interface{})[0]
				property = property.([]interface{})[0]
				switch k {
				case "Title":
					meta.Title = property.(string)
				case "Status":
					meta.Status = property.(string)
				case "Tags":
					meta.Tags = strings.Split(property.(string), ",")
				case "isPrivate":
					meta.IsPrivate = property.(string) == "Yes"
				}
			}
		}

		notes = append(notes, note.Note{
			Meta:    meta,
			Content: nil,
		})
	}

	return notes, nil
}

func downloadNotes(notes []note.Note) {
	logger("download all notes")
	for idx, n := range notes {
		page, err := client.DownloadPage(n.Meta.Id)
		if err != nil {
			fmt.Println("download failed")
		}
		n.Content = tomarkdown.ToMarkdown(page)
		n.Dump("./cache")

		msg := fmt.Sprintf("downloading...%d/%d", idx+1, len(notes))
		logger(msg)
	}
}

func main() {
	datasetId := "deba983740544616a2b8b399087ad180"
	notes, _ := getNotes(datasetId)
	downloadNotes(notes)
}
