package curl

import (
	"encoding/json"
	"fmt"
	"io"
	//"os"
	"net/http"
	"net/url"
	"strings"
)

type myJsonStruct struct {
	Path string `json:"path"`
}

func RequestV1(requestUrl string) int8 {
	fullUrl, err := url.ParseRequestURI(requestUrl)
	if err != nil {
		//os.Exit(1)
		// return errorcode and exit
	}

	response, err := http.Get(requestUrl)
	if err != nil {
		//os.Exit(1)
		// return errorcode and exit
	}
	defer response.Body.Close()

	buf := new(strings.Builder)
	n, err := io.Copy(buf, response.Body)
	if err != nil || n < 0 {
		//os.Exit(1)
		// return errorcode and exit
	}

	lineCount := strings.Count(buf.String(), "\n")

	if lineCount == 0 {
		fmt.Printf("{\"%s\": \"%s\"}", fullUrl.RequestURI(), buf.String())
	} else {
		var tmpArray []json.RawMessage
		tmp := strings.Split(buf.String(), "\n")
		for _, s := range tmp {
			if len(s) != 0 {
				in, err := json.Marshal(&myJsonStruct{Path: s})
				if err != nil {
					//os.Exit(1)
					// return errorcode and exit
				}
				newPath := json.RawMessage(in)
				tmpArray = append(tmpArray, newPath)
			}
		}

		fmt.Printf("{\"%s\": [", fullUrl.RequestURI())
		for i, s := range tmpArray {
			fmt.Printf("  %s", string(s))
			if i < len(tmpArray)-1 {
				fmt.Println(",")
			} else {
				fmt.Println()
			}
		}
		fmt.Println("]}")
	}
}
