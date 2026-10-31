package dummydata

import (
	"math/rand"
	"strconv"
)

var telHeaders = []string{
	"090",
	"080",
	"070",
	"050",
}

const telLength = 11
const telRand = 0.5

var telCounter = 99999999999

func NextTel(head string) string {
	r := rand.Float64()
	if r < telRand || head == "" {
		head = telHeaders[rand.Intn(len(telHeaders))]
	}
	headLen := len(head)
	restLen := telLength - headLen
	restStr := strconv.Itoa(telCounter)

	telStr := restStr[headLen : headLen+restLen]

	telCounter--

	return head + telStr
}
