package main

import (
	"3/calc"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1:]
	n, err := strconv.ParseInt(arg[0], 10, 32)
	if err != nil {
		fmt.Println("Error parse int")
	}

	result := int64(0)

	if len(arg) > 1 {
		if arg[1] == "-log" {
			result = calc.Calcilate(n, true)
		}
	} else {
		result = calc.Calcilate(n, false)
	}
	log.Info(result)

}
