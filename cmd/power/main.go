package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("test")

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				cmd := exec.Command("/usr/bin/ls", "-1")
				out, _ := cmd.Output()
				fmt.Printf("%s", out)
				/// ReadCSV command
				// parse_csv_and_publish(path)
			}
		}
	}()
	time.Sleep(1 * time.Hour)
}
