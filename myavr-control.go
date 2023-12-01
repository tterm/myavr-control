package main

import (
	"strings"
	"flag"
	"github.com/tarm/serial"
	"log"
	"fmt"
)

func printMap(modeMap map[string]string) string {
	var sb strings.Builder
	for key, value := range modeMap {
		sb.WriteString(key)
		sb.WriteString("(")
		sb.WriteString(value)
		sb.WriteString(")\n")
	}
	return sb.String()
}

func main() {
	const commandPrefix = "\xE6\xB5\xBA\xB9\xB2\xB3\xA9";
	modeTable := map[string]string {
		"p": "Program",
		"d": "Data",
		"r": "Reset Board",
		"R": "Reset Programmer",
		"+": "Power On Board",
		"-": "Power Off Board",
		"i": "Status",
	}
	var device string
	var baud int
	var mode string

	flag.StringVar(&device, "device", "", "Specify device name")
	flag.StringVar(&mode, "mode", "p", "Specify mode: \n" + fmt.Sprint(printMap(modeTable)))
	flag.IntVar(&baud, "baud", 19200, "Specify baudrate")
	flag.Parse()

	if len(device) == 0 {
		log.Printf("Device must be set\n")
		flag.PrintDefaults()
		return
	}

	_, found := modeTable[mode]
	if !found {
		log.Printf("Wrong mode set\n")
		flag.PrintDefaults()
		return
	}

	command := commandPrefix + mode
	fmt.Println(command)

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
