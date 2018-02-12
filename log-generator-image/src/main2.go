package main

import (
	"bufio"
	"os"
	"log/syslog"
	"log"
	"fmt"
)

func main() {
	sysLog, err := syslog.Dial("tcp", "127.0.0.1:603", syslog.LOG_INFO|syslog.LOG_SYSLOG, "log-generator|host|writer")
	if err != nil {
		log.Fatal(err)
	}
	sysLog.Info("")
}