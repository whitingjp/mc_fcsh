package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	log.Printf("Starting client.\n")
	client := new(http.Client)
	file := os.Args[1]
	if !path.IsAbs(file) {
		wd, _ := os.Getwd()
		file = wd + "/" + file
	}
	args := strings.Join(os.Args[2:], " ")
	log.Printf("Sending cmd: '"+file+" "+args+"'\n")
	request, _ := http.NewRequest("POST", "http://localhost:7950/compile", strings.NewReader(file+" "+args))
	resp, _ := client.Do(request)
	reader := bufio.NewReader(resp.Body)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			fmt.Print(str)
		}
	}
}
