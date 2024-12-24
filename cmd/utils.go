package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func resolveUrl(urlStr string) *url.URL {
	var reqUrl string
	if !strings.HasPrefix(urlStr, "http://") {
		reqUrl += "http://" + urlStr
	} else if strings.HasPrefix(urlStr, "https://") {
		reqUrl = strings.Replace(reqUrl, "https://", "http://", 1)
	} else {
		reqUrl = urlStr
	}
	fullUrl, err := url.ParseRequestURI(reqUrl)
	if err != nil {
		return &url.URL{}
	}
	return fullUrl
}

func processOutput(reqUrl string, input string) string {
	var result string
	lineCount := strings.Count(input, "\n")

	if lineCount == 0 {
		result += fmt.Sprintf("{\"%s\": \"%s\"}", reqUrl, input)
	} else {
		var tmpArray []json.RawMessage
		tmp := strings.Split(input, "\n")
		for _, s := range tmp {
			if len(s) != 0 {
				in, err := json.Marshal(&myJsonStruct{Path: s})
				if err != nil {
					return ""
				}
				newPath := json.RawMessage(in)
				tmpArray = append(tmpArray, newPath)
			}
		}

		result += fmt.Sprintf("{\"%s\": [", reqUrl)
		for i, s := range tmpArray {
			result += fmt.Sprintf("  %s", string(s))
			if i < len(tmpArray)-1 {
				result += ",\n"
			} else {
				result += "\n"
			}
		}
		result += "]}\n"
	}
	return result
}

func storeToken(token string) {
	f, _ := os.Create("./.token")
	defer f.Close()
	f.WriteString(token)
}

func retrieveToken() string {
	if _, err := os.Stat("./.token"); err == nil {
		token, _ := os.ReadFile("./.token")
		return string(token)
	}
	return ""
}

func processResponseString(response http.Response) (string, error) {
	buf := new(strings.Builder)
	n, err := io.Copy(buf, response.Body)
	if err != nil || n <= 0 || response.StatusCode != 200 {
		fmt.Print(buf)
		return buf.String(), fmt.Errorf("failed")
	}
	return buf.String(), nil
}

func TestQueryHelper(urlStr string) (*http.Request, string, error) {
	reqUrl := resolveUrl(urlStr)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return req, "", fmt.Errorf("Failed")
	}
	response, err := client.Do(req)
	if err != nil {
		return req, "", fmt.Errorf("Failed")
	}
	defer response.Body.Close()

	buf, err := processResponseString(*response)
	if err != nil {
		fmt.Print(buf)
		return req, "", fmt.Errorf("Failed")
	}
	return req, buf, nil
}

func TestQueryHelperV2(urlStr string) (*http.Request, string, error) {
	reqUrl := resolveUrl(urlStr)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return req, "", fmt.Errorf("Failed")
	}
	str := retrieveToken()
	if len(str) != 0 {
		req.Header.Set("X-aws-ec2-metadata-token", str)
	}

	response, err := client.Do(req)
	if err != nil {
		return req, "", fmt.Errorf("Failed")
	}
	defer response.Body.Close()

	buf, err := processResponseString(*response)
	if err != nil {
		fmt.Print(buf)
		return req, "", fmt.Errorf("Failed")
	}
	return req, buf, nil
}

func TestGetTokenHelper() {
	reqUrl := resolveUrl("localhost:1338/latest/api/token")
	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodPut, reqUrl.String(), nil)
	req.Header.Set("X-aws-ec2-metadata-token-ttl-seconds", "21600")
	res, err := client.Do(req)
	if err != nil {
		os.Exit(1)
	}
	defer res.Body.Close()

	t, _ := io.ReadAll(res.Body)
	storeToken(string(t))
}
