package dummydata

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

type Options struct {
	UseAlpha          bool
	IncludeUnanswered bool
	UseExtendedMJ     bool
	MinAge            int
	MaxAge            int
}

const (
	maxNameLength     = 205
	maxKanaLength     = 40
	maxAddress1Length = 100
	maxAddress2Length = 100

	// joyoRate は MJ 文字を選ぶ際に常用漢字から選ぶ確率
	joyoRate = 0.9
)

var (
	hiragana = filterRunes(jisx0208Kana, func(r rune) bool { return r >= 0x3041 && r <= 0x309F })
	katakana = filterRunes(jisx0208Kana, func(r rune) bool { return r >= 0x30A0 && r <= 0x30FF })
	digits   = filterRunes(jisx0208Alphabet, func(r rune) bool { return r >= 0xFF10 && r <= 0xFF19 })
	dash     = string(rune(0x2015))
	// smallKana は語頭に置かない拗音・促音などの小書き文字
	smallKana = []rune("ぁぃぅぇぉっゃゅょゎァィゥェォッャュョヮヵヶ")
	mjsAll    = append(append([][]rune{}, mjs...), mjsExtended...)
)

// toRunes は文字列の一覧を文字集合に変換する
func toRunes(src []string) [][]rune {
	dst := make([][]rune, len(src))
	for i, s := range src {
		dst[i] = []rune(s)
	}
	return dst
}

func filterRunes(src [][]rune, f func(rune) bool) [][]rune {
	var dst [][]rune
	for _, c := range src {
		if f(c[0]) {
			dst = append(dst, c)
		}
	}
	return dst
}

// mjsFor は UseExtendedMJ 指定時は表示できない文字も含めた MJ 文字を返す
func mjsFor(opt Options) [][]rune {
	if opt.UseExtendedMJ {
		return mjsAll
	}
	return mjs
}

// charSource は重み付きの文字集合
type charSource struct {
	runes  [][]rune
	weight float64
}

// uniform は文字数に比例した重みの文字集合を返す
func uniform(runes [][]rune) charSource {
	return charSource{runes, float64(len(runes))}
}

// mjSources は常用漢字を joyoRate の割合で、残りを mjsFor(opt) から選ぶ文字集合を返す
// 合計の重みは元の文字数のままなので、他の文字集合との比率は変わらない
func mjSources(opt Options) []charSource {
	total := float64(len(joyokanji) + len(mjsFor(opt)))
	return []charSource{
		{joyokanji, total * joyoRate},
		{mjsFor(opt), total * (1 - joyoRate)},
	}
}

// randomChars は sources の全文字を合わせた集合から一様に n 文字を生成する
func randomChars(n int, sources ...[][]rune) string {
	ws := make([]charSource, len(sources))
	for i, src := range sources {
		ws[i] = uniform(src)
	}
	return randomWeightedChars(n, ws...)
}

// randomWeightedChars は文字集合を重みに応じて選び、その中から一様に1文字ずつ n 文字を生成する
func randomWeightedChars(n int, sources ...charSource) string {
	total := 0.0
	for _, src := range sources {
		total += src.weight
	}
	var sb strings.Builder
	for i := 0; i < n; {
		r := rand.Float64() * total
		src := sources[len(sources)-1].runes
		for _, s := range sources {
			if r < s.weight {
				src = s.runes
				break
			}
			r -= s.weight
		}
		c := src[rand.Intn(len(src))]
		if i == 0 && slices.Contains(smallKana, c[0]) {
			continue
		}
		sb.WriteString(string(c))
		i++
	}
	return sb.String()
}

func randRange(min, max int) int {
	if max < min {
		return min
	}
	return min + rand.Intn(max-min+1)
}

// getRandomAgeInRange は年齢分布の重みを保ったまま [minAge, maxAge] から年齢を選ぶ
func getRandomAgeInRange(minAge, maxAge int) (int, error) {
	if minAge == 0 && maxAge >= ages[len(ages)-1].Age {
		return GetRandomAge()
	}
	var lo, hi float64
	prev := 0.0
	found := false
	for i := range ages {
		if ages[i].Age >= minAge && ages[i].Age <= maxAge {
			if !found {
				lo = prev
				found = true
			}
			hi = ages[i].Rate
		}
		prev = ages[i].Rate
	}
	if !found || hi <= lo {
		return 0, fmt.Errorf("no age in range %d-%d", minAge, maxAge)
	}
	return GetAgeByRate(lo + rand.Float64()*(hi-lo))
}
