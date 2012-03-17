package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type CompData struct {
	args string
	out  http.ResponseWriter
	back chan int
}

var buf_r *bufio.Reader
var buf_w *bufio.Writer
var targets = make(map[string]int)
var channel = make(chan CompData)

func Compiler() {
	for {
		data := <-channel
		log.Print("write")

		target, ok := targets[data.args]
		if ok {
			fmt.Fprintf(buf_w, "compile %v\n", target)
			buf_w.Flush()
		} else {
			fmt.Fprintf(buf_w, "mxmlc %v\n", data.args)
			buf_w.Flush()
			str, _ := buf_r.ReadString('\n')
			var i int
			fmt.Sscanf(str, " fcsh: Assigned %v as the compiler target id\n", &i)
			log.Print("Target ID:", i)
			targets[data.args] = i
		}

		waitForPrompt(data.out)
		data.back <- 0
	}
}

func Compile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprint(w, "Please send requests via POST")
		return
	}

	log.Print("Handling client.\n")

	reader := bufio.NewReader(r.Body)
	args, _ := reader.ReadString('\n') //  
	back := make(chan int)
	channel <- CompData{args: args, out: w, back: back}
	<-back
	fmt.Fprint(w, "\n")
}

func waitForPrompt(out io.Writer) {
	log.Print("Waiting for prompt")
	for {
		str, _ := buf_r.ReadString(')')
		fmt.Fprint(out, str)
		length := len(str)
		if length > 6 && str[length-6:] == "(fcsh)" {
			fmt.Fprint(out, "\n")
			log.Print("Finished waiting\n")
			break
		}
	}
}

func main() {
	log.Print("Starting server.\n")
	go Compiler()

	log.Print("Starting fcsh.\n")
	cmd := exec.Command("fcsh")
	in, _ := cmd.StdinPipe()
	out, _ := cmd.StdoutPipe()
	buf_r = bufio.NewReader(out)
	buf_w = bufio.NewWriter(in)

	if err := cmd.Start(); err != nil {
		log.Fatal("Could not start fcsh")
	}

	waitForPrompt(os.Stdout)
	waitForPrompt(os.Stdout)

	log.Print("Listening.\n")
	http.HandleFunc("/compile", Compile)
	http.ListenAndServe(":7950", nil)
}
