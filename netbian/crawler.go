package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
)

// Requirements: multitask crawl all the photos on the first page of each post

// Step: make a request, get the list page,
// Go to the details page and get the url, at the end of jpg
// Request the url, to get the resp, save file

// Tool function, which returns the content corresponding to url
func HandleUrl(url string) (Content string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	Content = string(bytes)
	fmt.Println(Content)
	return Content
}

// Tool function, saving pictures
func SaveImage(imageUrl string) {
	filePath := "/Users/batman/Downloads/image/" + path.Base(imageUrl)
	f, _ := os.Create(filePath)
	resp, _ := http.Get(imageUrl)
	defer f.Close()
	defer resp.Body.Close()
	reader := bufio.NewReaderSize(resp.Body, 32*1024)
	writer := bufio.NewWriter(f)
	_, _ = io.Copy(writer, reader)
	fmt.Println("图片保存完毕")
}

// Process each detail page and extract the url corresponding to the photo
func HandleDetail(detailUrl string) {
	detailContent := HandleUrl(detailUrl)
	reg := regexp.MustCompile("<img class=\"BDE_Image\" src=\".*?\" size=")
	imageTempSlice := reg.FindAllString(detailContent, -1)
	reg2 := regexp.MustCompile("http.*jpg")
	for _, i := range imageTempSlice {
		imgUrl := reg2.FindString(i)
		SaveImage(imgUrl)
	}
}

// Process each list page and extract the details page url
func HandleListContent(listContent string) {
	//<img src="/uploads/allimg/191022/000653-1571674013a6f4.jpg" alt="荷花4k壁纸3840x2160">
	//<a href="/tupian/21953.html" target="_blank">
	reg := regexp.MustCompile("<a href=\".*?\" target=")
	resultSlice := reg.FindAllString(listContent, -1)
	reg2 := regexp.MustCompile("/tupian.*html")

	seedUrl := "http://pic.netbian.com"
	for _, i := range resultSlice {
		detailUrl := seedUrl + reg2.FindString(i)
		Content := HandleUrl(detailUrl)
		imgReg := regexp.MustCompile("<img src=\".*?\" data-pic=")
		imgSlice := imgReg.FindAllString(Content, -1)
		imgReg2 := regexp.MustCompile("/uploads.*jpg")
		for _, j := range imgSlice {
			imgUrl := seedUrl + imgReg2.FindString(j)
			SaveImage(imgUrl)
		}

	}
}

// Build the url for each list page
func runListUrl() {
	startUrl := "http://pic.netbian.com/4kfengjing/index"
	for i := 1; i < 25; i++ {
		var Content string
		if i == 1 {
			Content = HandleUrl(startUrl + ".html")
		} else {
			Content = HandleUrl(startUrl + "_" + strconv.Itoa(i) + ".html")
		}
		HandleListContent(Content)
	}
}

func main() {
	runListUrl()
}
