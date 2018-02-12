package main

import (
	"os"
	"log/syslog"
	"log"
	"flag"
	"bufio"
)

func main() {
	syslogMode := flag.Bool("syslog", false, "as bool")
	stdoutMode := flag.Bool("stdout", true, "as bool")
	flag.Parse()

	if *syslogMode && *stdoutMode || !*syslogMode && !*stdoutMode {
		log.Fatal("wrong arguments")
	}

	if *syslogMode {
		writer, err := syslog.Dial("tcp", "127.0.0.1:603", syslog.LOG_INFO|syslog.LOG_SYSLOG, "log-generator|host|writer")
		if err != nil {
			log.Fatal(err)
		}

		for {
			err := writer.Info("hello\n")
			if err != nil {
				log.Fatal(err)
			}
		}
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