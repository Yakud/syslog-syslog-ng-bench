package main

import (
	"os/exec"
	"time"
	"log"
	"strings"
	"strconv"
	"fmt"
)

// syslog-ng-ctl stats parser
func main() {
	_, err := exec.Command("syslog-ng-ctl","stats", "--reset").Output()
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]int)

	secondsInterval := 1

	tick := time.Tick(time.Second * time.Duration(secondsInterval))
	for ; true; <-tick {
		out, err := exec.Command("syslog-ng-ctl","stats").Output()
		if err != nil {
			log.Fatal(err)
		}

		lines := strings.Split(string(out), "\n")

		for _, line := range lines[1:] {
			parts := strings.Split(line, ";")
			metric := strings.Join(parts[0:2], ";")
			currentValue, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				currentValue = 0
			}

			prevValue, ok := data[metric]
			if !ok {
				data[metric] = currentValue
				prevValue = currentValue
			}

			fmt.Print(metric,"\t",(currentValue - prevValue) / secondsInterval,"/sec\n")
		}
		fmt.Println("")
	}
}