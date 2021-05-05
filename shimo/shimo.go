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
	fmt.Println("排骨的石墨 MarkDown 图片下载工具")

	if len(os.Args) < 3 { 
		fmt.Println("lack of parameters")
		return
	}

	filePath := os.Args[1]
	token := os.Args[2]
	folder := filepath.Dir(filePath)

	content, _ := ioutil.ReadFile(filePath)
	md := string(content);
	md = strings.Replace(md, ")![", ")\n![", -1)
	re := regexp.MustCompile(`!\[(.*)\]\((https://uploader.shimo.im/.*)\)`)
	matches := re.FindAllStringSubmatch(md, -1)

	for _, match := range matches {
		url := match[2]
		url = url[0:strings.Index(url, "!thumbnail")]
		fileName := url[strings.LastIndex(url, "/") + 1:]
		url = url + "?accessToken=" + token
		fmt.Println("pulling ", fileName)
	  Download(url, filepath.Join(folder, fileName))
	}
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