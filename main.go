package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/nugielim/gpubsubbeat/beater"
)

func main() {
	err := beat.Run("gpubsubbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
