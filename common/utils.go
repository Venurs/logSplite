package common

import (
	"flag"
	"fmt"
	"os"
)

func ArgsParse() (string, string, string) {
	var logPath, splitPath, date string
	flag.StringVar(&logPath, "l", "", "log path")
	flag.StringVar(&splitPath, "s", "", "split path")
	flag.StringVar(&date, "d", "ALL", "which day log, default all days")
	flag.Parse()
	fmt.Println(logPath, splitPath, date)
	if flag.NFlag() < 3 {
		flag.Usage()
		os.Exit(0)
	}
	return logPath, splitPath, date
}
