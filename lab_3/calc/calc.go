package calc

import (
	log "github.com/sirupsen/logrus"
)

func Calcilate(n int64, flag bool) int64 {
	if flag {
		log.Print("Start calculate...")
		log.Printf("Calculate <%d>...\n", n)
	}
	res := int64(1)

	for i := int64(1); i-1 < n; i++ {
		res *= i
	}

	if flag {
		log.Print("Calculation complete!\n")
	}

	return res
}
