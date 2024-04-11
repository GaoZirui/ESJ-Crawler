package main

import (
	"github.com/gocolly/colly/v2"
	"log"
	"os"
	"strconv"
)

type Chapter struct {
	Name string
	Url  string
	Id   int
}

const yamlPath = `./config/config.yaml`

func main() {
	loadConfig(yamlPath)
	makeDir(config.StorageRootPathForTXT + "/" + config.LiNovelName)
	makeDir(config.StorageRootPathForEPUB)

	chapterList := make([]Chapter, 0)
	id := 0
	var file *os.File

	collector := colly.NewCollector()

	collector.OnError(func(r *colly.Response, err error) {
		log.Fatal("Request URL: ", r.Request.URL, "\nfailed with response: ", r, "\nError: ", err)
	})

	// login with email and password
	err := collector.Post(config.LoginUrl,
		map[string]string{"email": config.Email, "pwd": config.PassWord, "remember_me": "on"})
	if err != nil {
		log.Fatal(err)
	}

	// get all chapters
	collector.OnHTML(config.ChapterListSelector, func(element *colly.HTMLElement) {
		element.ForEach("a", func(i int, el *colly.HTMLElement) {
			if id > config.UpperChapter {
				return
			}
			if id >= config.LowerChapter {
				chapterList = append(chapterList, Chapter{
					Name: el.Attr("data-title"),
					Url:  el.Attr("href"),
					Id:   id,
				})
			}
			id++
		})
	})

	err = collector.Visit(config.ChapterListUrl)
	if err != nil {
		log.Fatal(err)
	}

	collector.OnHTML(config.ChapterTextSelector, func(element *colly.HTMLElement) {
		_, err = file.WriteString(element.Text + "\n")
		if err != nil {
			log.Fatal(err)
		}
	})

	// get each chapter content
	for _, chapter := range chapterList {
		file, err = os.Create(config.StorageRootPathForTXT + "/" + config.LiNovelName + "/" + strconv.Itoa(chapter.Id) + ".txt")
		if err != nil {
			log.Fatal(err)
		}

		_, err = file.WriteString(chapter.Name + "\n\n\n")
		if err != nil {
			log.Fatal(err)
		}

		err = collector.Visit(chapter.Url)
		if err != nil {
			log.Fatal(err)
		}

		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	if config.MakeEPUB == "on" {
		makeEPUB(chapterList)
	}
}
