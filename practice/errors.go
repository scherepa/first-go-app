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

func killServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		// here we do not wrap at all origionally
		// lets add wrap
		//return err
		return errors.Wrap(err, "bad process ID")
	}
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return errors.Wrap(err, "bad process ID")
	}

	// Simulate kill
	fmt.Printf("killing server with pid=%d\n", pid)

	if err := os.Remove(pidFile); err != nil {
		// here i did not get
		// We can go on if we fail here
		log.Printf("warning: can't remove pid file - %s", err)
	}

	return nil
}

func ErrorTraceToLog() {
	setupLogging()
	// we have no such file as we want error to be thrown
	cfg, err := readConfig("/path/to/config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		// this is for trace into log file
		log.Printf("error: %+v", err)
		//os.Exit(1)
		return
	}

	// Normal operation (redacted)
	fmt.Println(cfg)
}

func SimulateKill() {
	//setupLogging()
	// file is in app directory as no folder before in the path
	if err := killServer("server.pid"); err != nil {
		// we've got error when adding server.pid with word "seven inside as integer is expected"
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		//trace
		//log.Printf("error: %+v", err)
		//os.Exit(1)
		// lets panic but recover
		defer func() {
			if e := recover(); e != nil {
				err = fmt.Errorf("%v", e)
			}
		}()
		// %+v for trace
		log.Panicf("Error: %+v", err)
	}
}
