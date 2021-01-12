package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jn-lp/se-exam/pkg/tree"
)

func main() {
	input, err := os.Open("inputTree.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to open file,", err)
		return
	}
	defer input.Close()

	t, err := tree.New(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot decode tree from that reader,", err)
		return
	}

	if _, err := http.Post("http://localhost:8000/tree", "application/json", &t); err != nil {
		fmt.Fprintln(os.Stderr, "Cannot send POST request,", err)
		return
	} 
}
