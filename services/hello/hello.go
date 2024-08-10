package hello

import (
	"math/rand"
	"strconv"
)

func GenerateTimesLongMap(times int) map[string]int {
	result := make(map[string]int)
	for i := 0; i < times ;i++{
		result[strconv.Itoa(i)] = rand.Int()
	}
	return result
}