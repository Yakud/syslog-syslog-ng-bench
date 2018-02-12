package main

import (
	"os/exec"
	"time"
	"log"
	"strings"
	"strconv"
	"fmt"
	"flag"
)

// syslog-ng-ctl stats parser
func main() {
	interval := flag.Int("interval", 1, "as bool")
	flag.Parse()

	_, err := exec.Command("syslog-ng-ctl","stats", "--reset").Output()
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]int)


	tick := time.Tick(time.Second * time.Duration(*interval))
	for ; true; <-tick {
		out, err := exec.Command("syslog-ng-ctl","stats").Output()
		if err != nil {
			log.Fatal(err)
		}

		lines := strings.Split(string(out), "\n")

		for _, line := range lines[1:len(lines)-1] {
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

			fmt.Print(
				metric,
				"\t",
				fmt.Sprintf("%.3f",float64(currentValue - prevValue) / float64(*interval)),"/sec\n")
		}
		fmt.Println("")
	}
}