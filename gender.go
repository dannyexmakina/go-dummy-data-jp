package dummydata

import (
	"math/rand"
)

var genders = []string{
	"男性",
	"女性",
}

var gendersWithUnanswered = []string{
	"男性",
	"女性",
	"未回答",
}

func GetRandomGender(includeUnanswered bool) (gender string) {
	if includeUnanswered {
		return gendersWithUnanswered[rand.Intn(len(gendersWithUnanswered))]
	}
	return genders[rand.Intn(len(genders))]
}
