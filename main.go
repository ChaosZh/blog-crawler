package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chaoszh/blog-crawler/pkg/note"
	"github.com/kjk/notionapi"
	"github.com/kjk/notionapi/tomarkdown"
)


var client = &notionapi.Client{}
var datasetId = "deba983740544616a2b8b399087ad180"

var xmapper = map[string]string {
	"Title": "title",
	"Tags": "OpfN",
	"Status": "Nm\\r",
	"isPrivate": "__lm",
}

type NoteCollection []note.Note

func getNotes() ([]note.Note, error) {

	notes := make([]note.Note, 0)

	archivedPage, err := client.DownloadPage(datasetId)
	if err != nil {
		log.Fatalf("Failed at %s\n", err)
		return nil, err
	}
	catalogue := archivedPage.TableViews[0]

	for _, row := range catalogue.Rows {
		meta := note.NoteMeta{}		
		meta.Id = row.Page.GetNotionID().NoDashID
		meta.LastEditedTime = time.Unix(0, row.Page.LastEditedTime * int64(time.Millisecond)).Format("2006-01-02 15:04:05")
		for k, v := range xmapper {
			property := row.Page.Properties[v]
			if property != nil {
				property = property.([]interface{})[0]	//extract data from notion result
				property = property.([]interface{})[0]
				switch k{
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
			Meta: meta,
			Content: nil,
		})
	}

	return notes, nil
}

type CacheStrategy int

const (
	All = iota
	Update
)

func cacheNotes(notes []note.Note, strategy CacheStrategy) {
	// Todo: add update stratrgy, compare local cached with new note.metas, update local cached
	if strategy == All {
		for _, n := range notes {
			page, err := client.DownloadPage(n.Meta.Id)
			if err != nil {
				fmt.Println("Something wrong when downloading page...")
			}
			n.Content = tomarkdown.ToMarkdown(page)
			n.Dump("./cache")
		}
	}
}

func main() {
	notes, _ := getNotes()
	//fmt.Printf("%+v \n", notes) // debug

	cacheNotes(notes, All)
	//markdowns := getMarkdownsFromArchivedPage(arcPage)
	//cacheMarkdowns(markdowns)
}