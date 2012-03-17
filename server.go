package main

import (
	"log"
	"fmt"
	"net/http"
	"os/exec"
	"bufio"
)

var buf_r *bufio.Reader
var buf_w *bufio.Writer
var targets = make(map[string]int)
var channel = make(chan string)

func Compiler() {
	for {
		args := <-channel
		log.Print("Channel recv")
		log.Print("write")

		target, ok := targets[args]
		if ok {
			fmt.Fprintf(buf_w, "compile %v\n", target)
			buf_w.Flush()
		} else {
			fmt.Fprintf(buf_w, "mxmlc %v\n", args)
			buf_w.Flush()
			str, _ := buf_r.ReadString('\n')
			var i int
			fmt.Sscanf(str, " fcsh: Assigned %v as the compiler target id\n", &i)
			log.Print("Target ID:", i)
			targets[args] = i
		}
		

		waitForPrompt()
	}
}

func Compile(w http.ResponseWriter, r *http.Request) {
	log.Print("Handling client.\n")
	fmt.Fprintf(w, "Hello. client!")
	args := "H:/work/scraps/code/notzelda/main.mxml"	
	channel <- args
	log.Print("Channel send")
}

func waitForPrompt() {
	log.Print("Waiting for prompt")
	for {
		str, _ := buf_r.ReadString(')')
		fmt.Print(str)
		length := len(str)
		if(length > 6 && str[length-6:] == "(fcsh)") {
			fmt.Print("\n")
			log.Print("Finished waiting\n")
			break;	
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
	
	waitForPrompt()
	waitForPrompt()
	
	log.Print("Listening.\n")
	http.HandleFunc("/compile", Compile)
	http.ListenAndServe(":7950", nil)	
}
