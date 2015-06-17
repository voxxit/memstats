package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/beego/memcache"
)

var (
	hostname = flag.String("hostname", "localhost:11211", "IP/hostname + port of memcached instance to poll")
	seconds  = flag.Int("seconds", 1, "Number of seconds to wait between polls")
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	conn, err := memcache.Connect(*hostname)
	check(err)

	defer conn.Close()

	for {
		stats, err := conn.Stats("")
		check(err)

		fmt.Println(time.Now())
		fmt.Printf("%s\r\n", string(stats))

		time.Sleep(time.Duration(*seconds) * time.Second)
	}
}
