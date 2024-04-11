package main

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	LiNovelName            string `yaml:"LiNovelName"`
	StorageRootPathForTXT  string `yaml:"StorageRootPathForTXT"`
	StorageRootPathForEPUB string `yaml:"StorageRootPathForEPUB"`
	LoginUrl               string `yaml:"LoginUrl"`
	ChapterListUrl         string `yaml:"ChapterListUrl"`
	LowerChapter           int    `yaml:"LowerChapter"`
	UpperChapter           int    `yaml:"UpperChapter"`
	Email                  string `yaml:"Email"`
	PassWord               string `yaml:"PassWord"`
	ChapterListSelector    string `yaml:"ChapterListSelector"`
	ChapterTextSelector    string `yaml:"ChapterTextSelector"`
	MakeEPUB               string `yaml:"MakeEPUB"`
}

var config *Config

func loadConfig(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("[Config File Path Error]")
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("[Read Config File Error]")
	}

	if config.LowerChapter < 0 || config.UpperChapter < 0 || config.LowerChapter > config.UpperChapter {
		log.Fatal("[Chapter Bounds Error")
	}
}
