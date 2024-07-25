package main

import (
	"log/slog"
	"os"

	"github.com/naivary/specraft/runtime"
)

func main() {
	if err := run(); err != nil {
		slog.Error("something went wrong", "err_msg", err.Error())
		os.Exit(1)
	}
}

func run() error {
	rt, err := runtime.JSONSchema("examples/api", "examples/auth_req.go")
	if err != nil {
		return err
	}
    return rt.Generate()
}
