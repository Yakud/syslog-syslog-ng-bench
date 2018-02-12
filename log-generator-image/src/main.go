package main

import (
	"os"
	"log/syslog"
	"log"
	"flag"
	"bufio"
	"time"
	"fmt"
)

func main() {
	syslogMode := flag.Bool("syslog", false, "as bool")
	stdoutMode := flag.Bool("stdout", true, "as bool")
	flag.Parse()

	if *syslogMode && *stdoutMode || !*syslogMode && !*stdoutMode {
		log.Fatal("wrong arguments")
	}

	if *syslogMode {
		count := 0
		prevCount := 0

		go func() {
			for ;true;<-time.After(time.Second) {
				fmt.Println(count - prevCount)
				prevCount = count
			}
		}()

		go func() {
			writer, err := syslog.Dial("tcp", "127.0.0.1:603", syslog.LOG_INFO|syslog.LOG_SYSLOG, "log-generator|host|writer")
			if err != nil {
				log.Fatal(err)
			}

			for {
				err := writer.Info("hello\n")
				if err != nil {
					log.Fatal(err)
				}
				count += 1
			}
		}()
	}

	if *stdoutMode {
		writer := bufio.NewWriter(os.Stdout)
		for {
			writer.WriteString("hello\n")
			err := writer.Flush()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}