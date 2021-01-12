package main

import (
	"github.com/jn-lp/se-exam/pkg/tree"
)

func ComposeApiServer(port int) *APIServer {
	return &APIServer{
		Port:        port,
		TreeHandler: tree.HTTPHandler(),
	}
}
