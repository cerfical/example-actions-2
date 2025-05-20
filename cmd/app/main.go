package main

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting the application")
	greet(os.Stdout)
}

func greet(w io.Writer) {
	if _, err := fmt.Fprintf(w, "Hello, GitHub Actions!"); err != nil {
		panic(err)
	}
}
