package dummydata

import (
	"fmt"
	"math/rand"

	pref "dummy-data-jp/prefectures"
)

type Prefecture = pref.Prefecture
type Area = pref.Area

var prefectures = []Prefecture{
	pref.Hokkaido,
	pref.Aomori,
	pref.Iwate,
	pref.Miyagi,
	pref.Akita,
	pref.Yamagata,
	pref.Fukushima,
	pref.Ibaraki,
	pref.Tochigi,
	pref.Gunma,
	pref.Saitama,
	pref.Chiba,
	pref.Tokyo,
	pref.Kanagawa,
	pref.Niigata,
	pref.Toyama,
	pref.Ishikawa,
	pref.Fukui,
	pref.Yamanashi,
	pref.Nagano,
	pref.Gifu,
	pref.Shizuoka,
	pref.Aichi,
	pref.Mie,
	pref.Shiga,
	pref.Kyoto,
	pref.Osaka,
	pref.Hyogo,
	pref.Nara,
	pref.Wakayama,
	pref.Tottori,
	pref.Shimane,
	pref.Okayama,
	pref.Hiroshima,
	pref.Yamaguchi,
	pref.Tokushima,
	pref.Kagawa,
	pref.Ehime,
	pref.Kochi,
	pref.Fukuoka,
	pref.Saga,
	pref.Nagasaki,
	pref.Kumamoto,
	pref.Oita,
	pref.Miyazaki,
	pref.Kagoshima,
	pref.Okinawa,
}

func GetPrefectureByRate(rate float64) (name string, e error) {
	if rate < 0 {
		return name, fmt.Errorf("rate cannot be negative")
	}
	if rate > 1 {
		return name, fmt.Errorf("rate cannot be greater than 1")
	}
	for i := range prefectures {
		if prefectures[i].Rate >= rate {
			return prefectures[i].Name, nil
		}
	}
	return name, fmt.Errorf("rate not found")
}
func GetRandomPrefecture() (name string, e error) {
	rate := rand.Float64()
	return GetPrefectureByRate(rate)
}

func generateBanchi() string {
	if rand.Intn(2) == 0 {
		return "０００"
	}
	return "０" + dash + "００" + dash + "００"
}

func generateAddress1(opt Options, pref Prefecture, pc Area) string {
	base := pref.Name + pc.Address
	banchi := generateBanchi()
	room := maxAddress1Length - len([]rune(base)) - len([]rune(banchi))
	koaza := ""
	if rand.Intn(2) == 0 && room >= 2 {
		koaza = randomWeightedChars(randRange(2, min(4, room)), append(mjSources(opt), uniform(hiragana), uniform(katakana))...)
	}
	return base + koaza + banchi
}

func generateAddress2(opt Options) string {
	if rand.Float64() < 0.2 {
		return ""
	}
	return randomWeightedChars(randRange(1, min(30, maxAddress2Length)), append(mjSources(opt), uniform(hiragana), uniform(katakana), uniform(jisx0208Kigou))...)
}

func getPrefectureByRate(rate float64) (Prefecture, error) {
	name, err := GetPrefectureByRate(rate)
	if err != nil {
		return Prefecture{}, err
	}
	for i := range prefectures {
		if prefectures[i].Name == name {
			return prefectures[i], nil
		}
	}
	return Prefecture{}, fmt.Errorf("prefecture not found: %s", name)
}
