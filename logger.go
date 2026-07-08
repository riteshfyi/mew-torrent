package main

import (
	"log/slog"
	"os"
)

// Log is initialized immediately when the package is loaded
var Log = slog.New(slog.NewJSONHandler(os.Stdout, nil));