package main

import (
	"fmt"
	"os/exec"
	"time"
	"log"
)

func main() {
	_, err := exec.Command("syslog-ng-ctl","stats", "--reset").Output()
	if err != nil {
		log.Fatal(err)
	}

	tick := time.Tick(time.Second)
	for ; true; <-tick {
		out, err := exec.Command("syslog-ng-ctl","stats").Output()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(out))
	}
}