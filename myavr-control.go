package main

import (
	"flag"
	"github.com/tarm/serial"
	"log"
)

func main() {
	var device string
	var baud int
	var mode string

	/* There are more modes, but we will do it later */
	data := "\xE6\xB5\xBA\xB9\xB2\xB3\xA9d"
	program := "\xE6\xB5\xBA\xB9\xB2\xB3\xA9p"

	flag.StringVar(&device, "device", "", "Specify device name")
	flag.StringVar(&mode, "mode", "p", "Specify mode d->data mode | p->programmer mode")
	flag.IntVar(&baud, "baud", 19200, "Specify baudrate")

	flag.Parse()

	if len(device) == 0 {
		flag.PrintDefaults()
		return
	}

	command := data

	if mode == "p" {
		command = program
	}

	c := &serial.Config{Name: device, Baud: baud}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte(command))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])
}
