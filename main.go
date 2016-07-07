package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func loopLines(limit int64, r io.Reader, w io.Writer) error {
	sleep := time.Second / time.Duration(limit)
	s := bufio.NewScanner(r)
	ticker := time.NewTicker(sleep)
	for s.Scan() {
		line := s.Text()
		fmt.Fprintln(w, line)
		<-ticker.C
	}
	ticker.Stop()
	return s.Err()
}

func main() {
	limit := flag.Int64("l", 0, "lines per second (0 means no limit)")
	flag.Parse()

	err := loopLines(*limit, os.Stdin, os.Stdout)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
