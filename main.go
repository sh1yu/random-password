package main

import (
	"flag"
	"fmt"
	"math/rand"
	"bytes"
	"time"
)

const (
	base   = "base"
	extend = "extend"
	all    = "all"
)

var base_d = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
var base_e = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
var base_a = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_:@"

func main() {
	length := flag.Int("len", 16, "length of password")
	inputbase := flag.String("base", extend, "base character for password, valid value is \"base\", \"extend\" and \"all\".")
	flag.Parse()

	if *length <= 0 {
		fmt.Println("invalid parameter -len. use -h for help.")
		return
	}

	var selectedbase string
	switch *inputbase {
	case base:
		selectedbase = base_d
	case extend:
		selectedbase = base_e
	case all:
		selectedbase = base_a
	default:
		fmt.Println("invalid parameter -base. use -h for help.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	buf := bytes.Buffer{}
	for i := 0; i < *length; i++ {
		number := rand.Intn(len(selectedbase))
		buf.WriteByte(selectedbase[number])
	}
	fmt.Println("Randon password: " + buf.String())
}
