package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"path/filepath"
)

func main(){
	fmt.Println("------------------------------")
	fmt.Println("排骨的石墨 MarkDown 图片下载工具")
	fmt.Println("------------------------------")

	if len(os.Args) < 3 { 
		fmt.Println("lack of parameters")
		return
	}

	mdFilePath := os.Args[1]
	token := os.Args[2]

	folder := strings.TrimRight(mdFilePath, ".md")
	os.MkdirAll(folder, os.ModePerm)
	content, _ := ioutil.ReadFile(mdFilePath)
	md := string(content);
	md = strings.Replace(md, ")![", ")\n![", -1)
	re := regexp.MustCompile(`!\[(.*)\]\((https://uploader.shimo.im/.*)\)`)
	matches := re.FindAllStringSubmatch(md, -1)

	for _, match := range matches {
		originUrl := match[2]
		url := originUrl[0:strings.Index(originUrl, "!thumbnail")]
		fileName := url[strings.LastIndex(url, "/") + 1:]
		url = url + "?accessToken=" + token
		fmt.Println("pulling ", fileName)
	  Download(url, filepath.Join(folder, fileName))
		relativePath := "./" + filepath.Base(folder) + "/" + fileName
		md = strings.Replace(md, originUrl, relativePath, -1)
	}

	file, _ := os.OpenFile(mdFilePath, os.O_RDWR, 0755)
  defer file.Close()
  file.WriteString(md)
  file.Sync()
}

func Download(url string, filePath string) error {
	file, err := os.Create(filePath)

	if err == nil {
		defer file.Close()

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("Referer", "https://shimo.im")
		res, err := client.Do(req)
	
		if err == nil {			
			defer res.Body.Close()
			_, err = io.Copy(file, res.Body)
		}
	}
	
	return err
}