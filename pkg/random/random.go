package random

import (
	"fmt"
	"math/rand"
)

const (
	testMode = 1
	code     = "1111"
)

func GetRandomCode(mode int) string {
	if mode == testMode {
		return code
	}
	return fmt.Sprintf("%d",
		(rand.Intn(9)+1)*100000)
}
