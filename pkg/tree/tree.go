package tree

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type BinaryTree struct {
	Value  int         `json:",omitempty"`
	Left   *BinaryTree `json:",omitempty"`
	Middle *BinaryTree `json:",omitempty"`
	Right  *BinaryTree `json:",omitempty"`
}

func New(r io.Reader) (t BinaryTree, e error) {
	e = json.NewDecoder(r).Decode(&t)
	return
}

func (t *BinaryTree) Read(p []byte) (n int, err error) {
	data, err := json.Marshal(t)
	return bytes.NewReader(data).Read(p)
}

func (t *BinaryTree) Encode() (p []byte) {
	data, _ := json.Marshal(t)
	return data
}

func (t *BinaryTree) InverseBranches() {
	if t.Left == nil && t.Right == nil {
		return
	}

	t.Left, t.Right = t.Right, t.Left
	t.Left.InverseBranches()
	t.Right.InverseBranches()
}

func HTTPHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handleTreeSave(r, rw)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleTreeSave(r *http.Request, rw http.ResponseWriter) {
	t, err := New(r.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot decode tree from that reader,", err)
		return
	}

	t.InverseBranches()

	output, err := os.Create("outputTree.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to open file:", err)
		return
	}
	defer output.Close()

	output.Write(t.Encode())
}
