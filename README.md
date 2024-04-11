# ESJ-Crawler
A small go program used to crawl light novels from ESJ and package them into txt and epub formats.

***This program will not collect your personal information in any way!!!***
***Please do not use this software for any profit-making purposes, and obtain the translator's permission before crawling; the author is not responsible for any consequences! ! !***

1. <u>Complete config file.</u>
2. <u>Run Collector.exe</u>
3. <u>You are welcome to improve the code in any way you like!</u>

* config file

```yam
LiNovelName:
StorageRootPathForTXT: ./txt
StorageRootPathForEPUB: ./epub
LoginUrl: https://www.esjzone.me/inc/mem_login.php
ChapterListUrl:
LowerChapter: 0
UpperChapter: 100
Email:
PassWord:
ChapterListSelector: div[id='chapterList']
ChapterTextSelector: div[class='forum-content mt-3'] > p
MakeEPUB: on
```

* `LiNovelName`: name to be used when store in disk
* `StorageRootPathForTXT`: path to store txt files
* `StorageRootPathForEPUB`: path to store epub files
* `LoginUrl`: no need to change, use to login
* `ChapterListUrl`: enter the chapter list page url
* `LowerChapter & UpperChapter`: the chapter you pulling is behind `[L, U]`
* `Email`: your email use to login
* `PassWord`: your login password
* `ChapterListSelector`: `goquerySelector` use to match chapter list
* `ChapterTextSelector`: `goquerySelector` use to match chapter text
* `MakeEPUB`: whether need to make epub book
