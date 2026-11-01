package dummydata

import (
	"fmt"
	"math/rand"
)

type Person struct {
	FamilyName     string
	GivenName      string
	FamilyNameKana string
	GivenNameKana  string
	Gender         string
	Tel            string
	Prefecture     string
	Address1       string
	Address2       string
	Area           string
	Age            int
}

func GeneratePerson(opt Options) (p Person, e error) {
	p.FamilyName, p.GivenName = generateNames(opt)
	p.FamilyNameKana, p.GivenNameKana = generateKana(opt)
	p.Gender = GetRandomGender(opt.IncludeUnanswered)

	pref, err := getPrefectureByRate(rand.Float64())
	if err != nil {
		return p, err
	}
	if len(pref.Areas) == 0 {
		return p, fmt.Errorf("no postcodes for %s", pref.Name)
	}
	pc := pref.Areas[rand.Intn(len(pref.Areas))]
	p.Area = pc.Code
	p.Tel = NextTel(pc.Tel)
	p.Prefecture = pref.Name
	p.Address1 = generateAddress1(opt, pref, pc)
	p.Address2 = generateAddress2(opt)

	p.Age, err = getRandomAgeInRange(opt.MinAge, opt.MaxAge)
	return p, err
}

func generateNames(opt Options) (family, given string) {
	if opt.UseAlpha {
		// 氏名合計で文字数上限以下とし、氏と名の長さはそれぞれランダムにする
		familyLen := randRange(1, maxNameLength-1)
		return randomChars(familyLen, jisx0208Alphabet), randomChars(randRange(1, maxNameLength-familyLen), jisx0208Alphabet)
	}
	familyLen := randomNameLength()
	givenLen := min(randomNameLength(), maxNameLength-familyLen)
	return randomWeightedChars(familyLen, mjSources(opt)...), randomChars(givenLen, jinmeikanji, hiragana, katakana)
}

// nameLengthRates は氏・名の文字数（1〜4文字）ごとの出現割合（%）
var nameLengthRates = []int{5, 80, 14, 1}

func randomNameLength() int {
	n := rand.Intn(100)
	for i, rate := range nameLengthRates {
		if n < rate {
			return i + 1
		}
		n -= rate
	}
	return len(nameLengthRates)
}

func generateKana(opt Options) (family, given string) {
	if opt.UseAlpha {
		// 氏名合計で文字数上限以下とし、氏と名の長さはそれぞれランダムにする
		familyLen := randRange(1, maxKanaLength-1)
		return randomChars(familyLen, katakana), randomChars(randRange(1, maxKanaLength-familyLen), katakana)
	}
	familyLen := randRange(1, 8)
	givenLen := randRange(1, min(8, maxKanaLength-familyLen))
	return randomChars(familyLen, katakana), randomChars(givenLen, katakana)
}
