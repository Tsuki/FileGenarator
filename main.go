package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"bufio"
)

var input = flag.String("file", "input.txt", "source file")
var output = flag.String("output", "out.txt", "output file")
var sleep = flag.Int("time", 1, "time")

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(*input)
	check(err)
	fmt.Print(string(dat))
	write(dat)
}

func write(value []byte) {
	f, err := os.OpenFile(*output, os.O_WRONLY|os.O_CREATE, 0660)
	check(err)
	defer f.Close()
	for {
		w := bufio.NewWriter(f)
		n, err := w.Write(value)
		check(err)
		err = w.Flush()
		check(err)
		fmt.Printf("wrote %d bytes\n", n)
		time.Sleep(time.Duration(*sleep) * time.Second)
	}

}
