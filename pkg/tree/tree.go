package tree

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jn-lp/se-exam/pkg/tools"
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

func (t *BinaryTree) Encode() ([]byte, error) {
	return json.Marshal(t)
}

func (t *BinaryTree) Read(p []byte) (n int, err error) {
	data, err := t.Encode()
	return bytes.NewReader(data).Read(p)
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
		log.Println("Cannot decode tree from that reader,", err)
		tools.WriteJsonBadRequest(rw, "Cannot decode tree from that reader")
		return
	}

	t.InverseBranches()

	output, err := os.Create("outputTree.json")
	if err != nil {
		log.Println("Unable to open file:", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	defer output.Close()

	bytes, err := t.Encode()
	if err != nil {
		log.Println("Unable to encode tree struct:", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	output.Write(bytes)
}
