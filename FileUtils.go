package main

import (
	"github.com/bmaupin/go-epub"
	"log"
	"os"
	"strconv"
	"strings"
)

func isFolderExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func makeDir(path string) {
	if !isFolderExists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func makeEPUB(chapterList []Chapter) {
	book := epub.NewEpub(config.LiNovelName)

	for _, chapter := range chapterList {
		text, err := os.ReadFile(config.StorageRootPathForTXT + "/" + config.LiNovelName + "/" + strconv.Itoa(chapter.Id) + ".txt")
		fmtText := strings.Replace(string(text), "\n", "<p>", -1)
		if err != nil {
			log.Fatal(err)
		}
		_, err = book.AddSection(fmtText, chapter.Name, "", "")
		if err != nil {
			log.Fatal(err)
		}
	}

	err := book.Write(config.StorageRootPathForEPUB + "/" + config.LiNovelName + ".epub")
	if err != nil {
		log.Fatal(err)
	}
}
