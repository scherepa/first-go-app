// pkg/errors example
package practice

import (
	"fmt"
	"log"
	"os"

	// here we use not build in package for errors
	"github.com/pkg/errors"
)

// Config holds configuration
type Config struct {
	// configuration fields go here (redacted)
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't open configuration file")
	}
	defer file.Close()

	cfg := &Config{}
	// Parse file here (redacted)
	return cfg, nil

}

func setupLogging() {
	// if no such file create
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func ErrorTraceToLog() {
	setupLogging()
	// we have no such file as we want error to be thrown
	cfg, err := readConfig("/path/to/config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		// this is for trace into log file
		log.Printf("error: %+v", err)
		os.Exit(1)
	}

	// Normal operation (redacted)
	fmt.Println(cfg)
}
