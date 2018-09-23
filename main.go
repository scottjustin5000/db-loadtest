package main

import (
	"flag"
)

func main() {
	rps := flag.Int("rps", 0, "request per second")
	f := flag.String("f", "", "queries file")
	c := flag.Int("c", 0, "level of concurrency (only used when EOF = true)")
	cs := flag.String("cs", "", "connection string")
	d := flag.Int("d", 0, "time to run in seconds")
	eof := flag.Bool("eof", false, "end test when file exhausted")
	r := flag.Bool("r", false, "process queries randomly")
	s := flag.Bool("s", false, "process queries sequentially")

	flag.Parse()
	cnf := Config{RequestsPerSecond: *rps, ConnectionString: *cs, Random: *r, Sequential: *s, File: *f, Duration: *d, EOF: *eof, Concurrency: *c}

	Launch(cnf)
}
