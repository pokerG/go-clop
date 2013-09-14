package main

import (
	"fmt"
	"io/ioutil"
	"jsonOperate"
	"bytes"
	"net/http"
	//"net/url"
	//"net"
	"mime/multipart"
	"os"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}

}

var js *jsonOperate.Json
var urlc string = "http://clop.hit.edu.cn/toStulogin.action"

func initjs() {
	file, err := os.Open("port.json")
	handleError(err)
	defer file.Close()

	lines, _ := ioutil.ReadFile(file.Name())
	js, err = jsonOperate.NewJson(lines)
	handleError(err)
}

func main() {
	initjs()
	fmt.Println(js)
	connect(urlc)
}

func connect(urlc string) {
	bodyBuf := &bytes.Buffer{}
    bodyWriter := multipart.NewWriter(bodyBuf)

    for 
}
