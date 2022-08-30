package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/*
Поддерживает протоколы HTTP, FTP и HTTPS, а также поддерживает работу через HTTP прокси-сервер
‘-m’
‘--mirror’
Turn on options suitable for mirroring. This option turns on recursion and time-stamping,
sets infinite recursion depth and keeps FTP directory listings.
It is currently equivalent to ‘-r -N -l inf --no-remove-listing’.
[1] https://habr.com/ru/company/ruvds/blog/346640/
[2] https://www.gnu.org/software/wget/manual/wget.html

> go get golang.org/x/net/html
go: downloading golang.org/x/net v0.0.0-20220826154423-83b083e8dc8b
go get: added golang.org/x/net v0.0.0-20220826154423-83b083e8dc8b
*/

func main() {
	mirrorFlag := flag.Bool("", false, "download site (mirror)")

	flag.Parse()

	url := flag.Arg(0)
	fmt.Printf("Downloading :: from \n", url)

	if url == "" {
		fmt.Fprintf(os.Stderr, "Error :: uncorrect URL!\n")
		return
	}

	// 2 branches of deviation . mirror flag case and base case
	if *mirrorFlag {
		pathLocal, err := ParseURL(url, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error:: incorrect URL!\n")
		}
		pathLocal = path.Base(url)
		fmt.Println("Checking :: if " + pathLocal + " exists ...")
	} else {
		pathLocal := path.Base(url)
		fmt.Println("Checking :: if " + pathLocal + " exists ...")
		if _, err := os.Stat(pathLocal); os.IsNotExist(err) {
			err := DownloadFile(url, pathLocal)
			if err != nil {
				panic(err)
			}
			fmt.Println("Success :: " + pathLocal + " saved!")
		} else {
			fmt.Println("Attention :: " + pathLocal + " already exist!")
		}
	}

}

//CheckURL on site belonging
func CheckURL(url string) bool {
	if url != "" {
		runes := []rune(url)
		return runes[0] == '/'
	}
	return false
}

//ParseURL parse and validate url
func ParseURL(url string, domain bool) (string, error) {
	if url == "" {
		return "", errors.New("incorrect url")
	}
	var localURL = ""
	splitFn := func(c rune) bool {
		return c == '/'
	}
	parts := strings.FieldsFunc(url, splitFn)
	// first case - get domain, for example: http://domain.com
	if domain {
		localURL = parts[0] + "//" + parts[1]
		return localURL, nil
	}
	// second case - get local path, for example: /domain1/domain2
	if len(parts) > 2 {
		for i := 2; i < len(parts); i++ {
			localURL = localURL + "/" + parts[i]
		}
		return localURL, nil
	}
	return "/", nil
}

//DownloadFile method - get html page downloaded
func DownloadFile(url string, path string) error {
	client := http.Client{} // Use simple http client
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"User-Agent": {"a"}, //fix response 410 from Apache
	}

	// getting html page
	response, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer response.Body.Close()

	// create file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// write data to file
	_, err = io.Copy(file, response.Body)

	return err
}

//DownloadSite - downloads the entire site
func DownloadSite(url string, path string) error {
	fmt.Println("Downloading :: ", url, " to ", path)

	// get path, for example - /domain1/domain2
	urlWithoutDomain, err := ParseURL(url, false)
	if err != nil {
		return err
	}

	// get domain, for example - https://domain.com
	domain, err := ParseURL(url, true)
	if err != nil {
		return err
	}

	// map contains all paths (in the end)
	mapPath := make(map[string]bool)
	mapPath[urlWithoutDomain] = false

	// download
	_, err = DownloadPages(domain, path, urlWithoutDomain, mapPath)
	if err != nil {
		return err
	}

	return nil
}

//DownloadPages - downloads all pages (recursion)
func DownloadPages(url, path, urlWithoutDomain string, mapPath map[string]bool) (map[string]bool, error) {

	//page is loading (no repeat)
	mapPath[urlWithoutDomain] = true

	client := http.Client{}
	req, _ := http.NewRequest("GET", url+urlWithoutDomain, nil)
	req.Header = http.Header{
		"User-Agent": {"a"}, //fix response 410 from Apache
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	// save page to file
	if urlWithoutDomain == "/" {
		err = CreateFile(path+urlWithoutDomain+"index.html", bytes.NewBuffer(body))
	} else {
		err = CreateFile(path+urlWithoutDomain+".html", bytes.NewBuffer(body))
	}
	if err != nil {
		return nil, err
	}

	// parse page and get links from page
	mapPath = GetLinks(bytes.NewBuffer(body), mapPath)

	// call recursion
	for k := range mapPath {
		if !mapPath[k] {
			mapPath, err = DownloadPages(url, path, k, mapPath)
			if err != nil {
				return nil, err
			}
		}
	}

	return mapPath, nil
}

//GetLinks - parse body of page (get tag "a -> href")
func GetLinks(body io.Reader, mapPath map[string]bool) map[string]bool {
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return mapPath
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						if CheckURL(attr.Val) && !mapPath[attr.Val] {
							mapPath[attr.Val] = false
						}
					}
				}
			}
		}
	}
}

//CreateFile - creates a file and all folders along the way
func CreateFile(path string, body io.Reader) error {
	parts := strings.Split(path, "/")
	var tempPath string
	var lenPath = len(parts)
	for i, val := range parts {
		// create paths
		if i == 0 {
			tempPath = val
		} else {
			tempPath = tempPath + "/" + val
		}

		if i != lenPath-1 {
			if _, err := os.Stat(tempPath); os.IsNotExist(err) {
				err := os.Mkdir(tempPath, 0750)
				if err != nil {
					return err
				}
			}
		} else {
			// create file
			file, err := os.Create(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(file, body)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

/*
go run .\task.go "https://stackoverflow.com/questions/22197496/how-to-perform-action-on-clicking-a-qmenu-object-only"
Downloading :: from
%!(EXTRA string=https://stackoverflow.com/questions/22197496/how-to-perform-action-on-clicking-a-qmenu-object-only)Checking :: if how-to-perform-action-on-clicking-a-qmenu-object-only exists ...
Success :: how-to-perform-action-on-clicking-a-qmenu-object-only saved!
*/
