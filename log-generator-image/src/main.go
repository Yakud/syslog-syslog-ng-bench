package main

import (
	"os"
	"log/syslog"
	"log"
	"flag"
	"bufio"
	"time"
	"fmt"
	"sync"
)

func main() {
	syslogMode := flag.Bool("syslog", false, "as bool")
	stdoutMode := flag.Bool("stdout", true, "as bool")
	workers := flag.Int("workers", 1, "as bool")
	flag.Parse()

	if *syslogMode && *stdoutMode || !*syslogMode && !*stdoutMode {
		log.Fatal("wrong arguments")
	}

	if *syslogMode {
		counterMutex := &sync.Mutex{}
		count := 0
		prevCount := 0

		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("ticker stats")
			tick := time.Tick(time.Second)
			for ; true; <-tick {
				fmt.Print(count - prevCount, " msg/sec\n")
				prevCount = count
			}
		}()

		wg.Add(*workers)
		for i := 0; i < *workers; i ++ {
			go func() {
				defer wg.Done()
				writer, err := syslog.Dial("tcp", "127.0.0.1:603", syslog.LOG_INFO|syslog.LOG_SYSLOG, "log-generator|host|writer")
				if err != nil {
					log.Fatal(err)
				}
				time.Sleep(time.Millisecond)

				for {
					err := writer.Info("hello\n")
					if err != nil {
						log.Fatal(err)
					}
					counterMutex.Lock()
					count += 1
					counterMutex.Unlock()
				}
			}()
		}

		wg.Wait()
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