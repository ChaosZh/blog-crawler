package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kjk/notionapi"
	"github.com/kjk/notionapi/tomarkdown"
)

type MarkdownFile struct {
	Title			string
	Content			[]byte
	Tags			interface{}
	LastEditedTime	string
}

func (markdown MarkdownFile) toFile(folder string) {
	_ = os.WriteFile(folder + "/" + markdown.Title + ".md", markdown.Content, 0644)
}

var client = &notionapi.Client{}
var datasetId = "deba983740544616a2b8b399087ad180"

func getArchivedPage() (*notionapi.Page, error) {
	page, err := client.DownloadPage(datasetId)
	if err != nil {
		log.Fatalf("Failed at %s\n", err)
		return nil, err
	} else {
		return page, nil
	}
}

func getMarkdownsFromArchivedPage(archivedPage *notionapi.Page) ([]MarkdownFile) {
	//table_view := archivedPage.Root().Content[0]
	//print(archivedPage.TableViews)
	catalogue := archivedPage.TableViews[0]
	markdowns := []MarkdownFile{}
	for idx, row := range catalogue.Rows {
		print("dealing with " + fmt.Sprint(idx))
		nid := row.Page.GetNotionID().NoDashID
		page, _ := client.DownloadPage(nid)


		title := row.Page.Properties["title"]
		content := tomarkdown.ToMarkdown(page)
		tags := row.Page.Properties["OpfN"]
		time := time.Unix(0, row.Page.LastEditedTime * int64(time.Millisecond)).Format("2006-01-02 15:04:05")

		res := MarkdownFile {fmt.Sprint(title), content, tags, time}

		markdowns = append(markdowns, res)
		//break
	}
	return markdowns
}

func cacheMarkdowns(markdowns []MarkdownFile){
	for _, mdx := range markdowns {
		mdx.toFile("./cache")
	}
}

func print(obj interface{}) {
	fmt.Printf("%+v \n", obj)
	//fmt.Println(json.MarshalIndent(obj, "", "\t"))
}

func main() {
	arcPage, _ := getArchivedPage()
	markdowns := getMarkdownsFromArchivedPage(arcPage)
	cacheMarkdowns(markdowns)
	//print(markdowns[0])
	//fmt.Println(markdowns[0].Tags.([]interface{})[0].([]interface{})[0])
	//fmt.Println(markdowns[0].Title.([]interface{})[0].([]interface{}))
	//s:=reflect.ValueOf(&markdowns[0]).Elem()
	//for i := 0; i < s.NumField(); i++ {
	//	f := s.Field(i)
	//	fmt.Printf("%d: %s %s = %v\n", i,
	//		s.Type().Field(i).Name, f.Type(), f.Interface())
	//}
	//print()
}