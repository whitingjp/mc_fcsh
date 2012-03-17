package main

import(
	"log"
	"net/http"
	"fmt"
	"strings"
	"bufio"
	"io"
)

func main() {
	log.Printf("Starting client.\n")
	client := new (http.Client)
	args := "H:/work/scraps/code/notzelda/main.mxml"
	request, _ := http.NewRequest("POST", "http://localhost:7950/compile", strings.NewReader(args))
	resp, _ := client.Do(request)
	reader := bufio.NewReader(resp.Body)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else
		{
			fmt.Print(str)
		}
	}
}
