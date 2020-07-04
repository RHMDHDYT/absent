package curl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var METHOD_GET = "GET"
var METHOD_POST = "POST"

type Curl struct {
	BaseUrl         string
	Method          string
	Body            map[string]string
	Header          map[string]string
	HeaderMultiPart map[string]string
	File            map[string]string
}

func (c *Curl) Send() ([]byte, error) {
	var client = &http.Client{}
	var request *http.Request
	switch c.Method {
	case METHOD_GET:
		r, _ := http.NewRequest(c.Method, c.BaseUrl, nil)
		request = r
	case METHOD_POST:
		var b io.Reader
		if len(c.getMultiPartHeader()) <= 0 {
			b = c.bodyMultiPart()
		} else {
			b = c.body()
		}
		r, _ := http.NewRequest(c.Method, c.BaseUrl, b)
		request = r
	default:
		return nil, errors.New("Method Not Available")
	}
	request = c.header(request)
	response, _ := client.Do(request)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return body, err
}

func (c *Curl) body() io.Reader {
	bytes, _ := json.Marshal(c.Body)
	return strings.NewReader(string(bytes))
}

func (c *Curl) header(r *http.Request) *http.Request {
	if len(c.getMultiPartHeader()) <= 0 {
		r.Header.Add("Content-Type", "application/json")
	} else {
		for k, v := range c.getMultiPartHeader() {
			r.Header.Add(k, v)
		}
	}
	for k, v := range c.Header {
		r.Header.Add(k, v)
	}
	return r
}

func (c *Curl) setMultiPartHeader(writer *multipart.Writer) {
	c.HeaderMultiPart = map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}
}

func (c *Curl) getMultiPartHeader() map[string]string {
	return c.HeaderMultiPart
}

func (c *Curl) bodyMultiPart() io.Reader {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for k, v := range c.Body {
		_ = writer.WriteField(k, v)
	}
	writer = c.setFile(writer)
	c.setMultiPartHeader(writer)
	return payload
}

func (c *Curl) setFile(writer *multipart.Writer) *multipart.Writer {
	for kFile, vFile := range c.File {
		file, errFile4 := os.Open(vFile)
		defer file.Close()
		part4, errFile4 := writer.CreateFormFile(kFile, filepath.Base(vFile))
		_, errFile4 = io.Copy(part4, file)
		if errFile4 != nil {
			fmt.Println(errFile4)
		}
		err := writer.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
	return writer
}
