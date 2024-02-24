package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var tableTests = []struct {
	in  []string
	out []string
}{
	{[]string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
		"Thanks.", "I love music of Kartik.", "I love music of Kartik."}, []string{"I love music.", " ", "I love music of Kartik.",
		"Thanks.", "I love music of Kartik."}},
	{[]string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
		"Thanks.", "I love music of Kartik.", "I love music of Kartik."}, []string{"3 I love music.", "1  ", "2 I love music of Kartik.",
		"1 Thanks.", "2 I love music of Kartik."}},
	{[]string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
		"Thanks.", "I love music of Kartik.", "I love music of Kartik."}, []string{"I love music.", "I love music of Kartik.", "I love music of Kartik."}},
	{[]string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
		"Thanks.", "I love music of Kartik.", "I love music of Kartik."}, []string{" ", "Thanks."}},
	{[]string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", " ", "I love MuSIC of Kartik.", "I love music of kartik.",
		"Thanks.", "I love music of kartik.", "I love MuSIC of Kartik."}, []string{"I LOVE MUSIC.", " ", "I love MuSIC of Kartik.", "Thanks.", "I love music of kartik."}},
	{[]string{"We love music.", "I love music.", "They love music.", " ", "I love music of Kartik.", "We love music of Kartik.",
		"Thanks."}, []string{"We love music.", " ", "I love music of Kartik.", "Thanks."}},
	{[]string{"I love music.", "A love music.", "C love music.", " ", "I love music of Kartik.", "We love music of Kartik.",
		"Thanks."}, []string{"I love music.", " ", "I love music of Kartik.", "We love music of Kartik.", "Thanks."}},
	{[]string{"Небо", "Небо", "Облако", "Облако", "Пляж", "Анапа",
		"Анапа"}, []string{"Небо", "Облако", "Пляж", "Анапа"}},
	{[]string{"Небо", "Небо", "Облако", "Облако", "Пляж", "Анапа",
		"Анапа"}, []string{"2 Небо", "2 Облако", "1 Пляж", "2 Анапа"}},
	{[]string{"Небо", "Небо", "Облако", "Облако", "Пляж", "Анапа",
		"Анапа"}, []string{"Небо", "Облако", "Анапа"}},
	{[]string{"Небо", "Небо", "Облако", "Облако", "Пляж", "Анапа",
		"Анапа"}, []string{"Пляж"}},
	{[]string{"Небо голуБое", "Небо голубое", "Облако Голубое", "Облако ГолУбое", "Пляж чистый", "Анапа солнечнаЯ",
		"Анапа солНечная", "Сторона СолНечная"}, []string{"Небо голуБое", "Пляж чистый", "Анапа солнечнаЯ"}},
	{[]string{"Небо голуБое", "Небо голубое", "Облако Голубое", "Облако ГолУбое", "Пляж чистый", "Анапа солнечнаЯ",
		"Анапа солНечная", "Сторона СолНечная"}, []string{"4 Небо голуБое", "1 Пляж чистый", "3 Анапа солнечнаЯ"}},
	{[]string{"Небо голуБое", "Небо голубое", "Облако Голубое", "Облако ГолУбое", "Пляж чистый", "Анапа солнечнаЯ",
		"Анапа солНечная", "Сторона СолНечная"}, []string{"Небо голуБое", "Анапа солнечнаЯ"}},
	{[]string{"Небо голуБое", "Небо голубое", "Облако Голубое", "Облако ГолУбое", "Пляж чистый", "Анапа солнечнаЯ",
		"Анапа солНечная", "Сторона СолНечная"}, []string{"Пляж чистый"}},
	{[]string{"Небо и голуБое", "Небо я голубое", "Облако Б Голубое", "Облако Д Голубое", "Пляж Я чистый", "Анапа т солнечнаЯ",
		"Анапа г солНечная", "Сторона ф СолНечная"}, []string{"Небо и голуБое", "Пляж Я чистый", "Анапа т солнечнаЯ"}},
	{[]string{"Небо и голуБое", "Небо я голубое", "Облако Б Голубое", "Облако А Голубое", "Пляж Я чистый", "Анапа т солнечнаЯ",
		"Анапа г солНечная", "Сторона ф СолНечная"}, []string{"4 Небо и голуБое", "1 Пляж Я чистый", "3 Анапа т солнечнаЯ"}},
	{[]string{"Небо и голуБое", "Небо я голубое", "Облако Б Голубое", "Облако Д Голубое", "Пляж Я чистый", "Анапа т солнечнаЯ",
		"Анапа г солНечная", "Сторона ф СолНечная"}, []string{"Небо и голуБое", "Анапа т солнечнаЯ"}},
	{[]string{"Небо и голуБое", "Небо я голубое", "Облако Б Голубое", "Облако Д Голубое", "Пляж Я чистый", "Анапа т солнечнаЯ",
		"Анапа г солНечная", "Сторона ф СолНечная"}, []string{"Пляж Я чистый"}},
}

func TestTableParserSuccess(t *testing.T) {
	require.Equal(t, tableTests[0].out, Uniq(tableTests[0].in, Options{}), "The two slices should be the same.")
	require.Equal(t, tableTests[1].out, Uniq(tableTests[1].in, Options{C: true}), "The two slices should be the same.")
	require.Equal(t, tableTests[2].out, Uniq(tableTests[2].in, Options{D: true}), "The two slices should be the same.")
	require.Equal(t, tableTests[3].out, Uniq(tableTests[3].in, Options{U: true}), "The two slices should be the same.")
	require.Equal(t, tableTests[4].out, Uniq(tableTests[4].in, Options{I: true}), "The two slices should be the same.")
	require.Equal(t, tableTests[5].out, Uniq(tableTests[5].in, Options{F: 1}), "The two slices should be the same.")
	require.Equal(t, tableTests[6].out, Uniq(tableTests[6].in, Options{S: 1}), "The two slices should be the same.")
	require.Equal(t, tableTests[7].out, Uniq(tableTests[7].in, Options{}), "The two slices should be the same.")
	require.Equal(t, tableTests[8].out, Uniq(tableTests[8].in, Options{C: true}), "The two slices should be the same.")
	require.Equal(t, tableTests[9].out, Uniq(tableTests[9].in, Options{D: true}), "The two slices should be the same.")
	require.Equal(t, tableTests[10].out, Uniq(tableTests[10].in, Options{U: true}), "The two slices should be the same.")
	require.Equal(t, tableTests[11].out, Uniq(tableTests[11].in, Options{I: true, F: 1}), "The two slices should be the same.")
	require.Equal(t, tableTests[12].out, Uniq(tableTests[12].in, Options{I: true, F: 1, C: true}), "The two slices should be the same.")
	require.Equal(t, tableTests[13].out, Uniq(tableTests[13].in, Options{I: true, F: 1, D: true}), "The two slices should be the same.")
	require.Equal(t, tableTests[14].out, Uniq(tableTests[14].in, Options{I: true, F: 1, U: true}), "The two slices should be the same.")
	require.Equal(t, tableTests[15].out, Uniq(tableTests[15].in, Options{I: true, F: 1, S: 3}), "The two slices should be the same.")
	require.Equal(t, tableTests[16].out, Uniq(tableTests[16].in, Options{I: true, F: 1, S: 3, C: true}), "The two slices should be the same.")
	require.Equal(t, tableTests[17].out, Uniq(tableTests[17].in, Options{I: true, F: 1, S: 3, D: true}), "The two slices should be the same.")
	require.Equal(t, tableTests[18].out, Uniq(tableTests[18].in, Options{I: true, F: 1, S: 3, U: true}), "The two slices should be the same.")
}

func TestTableParserFail1(t *testing.T) {
	require.Equal(t, tableTests[0].out, Uniq(tableTests[4].in, Options{}), "The two slices should be the same.")
}

func TestTableParserFail2(t *testing.T) {
	require.Equal(t, tableTests[1].out, Uniq(tableTests[16].in, Options{C: true}), "The two slices should be the same.")
}

func TestTableParserFail3(t *testing.T) {
	require.Equal(t, tableTests[2].out, Uniq(tableTests[15].in, Options{D: true}), "The two slices should be the same.")
}

func TestTableParserFail4(t *testing.T) {
	require.Equal(t, tableTests[3].out, Uniq(tableTests[7].in, Options{U: true}), "The two slices should be the same.")
}

func TestTableParserFail5(t *testing.T) {
	require.Equal(t, tableTests[4].out, Uniq(tableTests[14].in, Options{I: true}), "The two slices should be the same.")
}

func TestTableParserFail6(t *testing.T) {
	require.Equal(t, tableTests[5].out, Uniq(tableTests[9].in, Options{F: 1}), "The two slices should be the same.")
}

func TestTableParserFail7(t *testing.T) {
	require.Equal(t, tableTests[6].out, Uniq(tableTests[8].in, Options{S: 1}), "The two slices should be the same.")
}

func TestTableParserFail8(t *testing.T) {
	require.Equal(t, tableTests[7].out, Uniq(tableTests[1].in, Options{}), "The two slices should be the same.")

}

func TestTableParserFail9(t *testing.T) {
	require.Equal(t, tableTests[8].out, Uniq(tableTests[2].in, Options{C: true}), "The two slices should be the same.")

}
func TestTableParserFail10(t *testing.T) {
	require.Equal(t, tableTests[9].out, Uniq(tableTests[18].in, Options{D: true}), "The two slices should be the same.")

}
func TestTableParserFail11(t *testing.T) {
	require.Equal(t, tableTests[10].out, Uniq(tableTests[5].in, Options{U: true}), "The two slices should be the same.")

}
func TestTableParserFail12(t *testing.T) {
	require.Equal(t, tableTests[11].out, Uniq(tableTests[3].in, Options{I: true, F: 1}), "The two slices should be the same.")

}

func TestTableParserFail13(t *testing.T) {
	require.Equal(t, tableTests[12].out, Uniq(tableTests[8].in, Options{I: true, F: 1, C: true}), "The two slices should be the same.")

}
func TestTableParserFail14(t *testing.T) {
	require.Equal(t, tableTests[13].out, Uniq(tableTests[6].in, Options{I: true, F: 1, D: true}), "The two slices should be the same.")

}
func TestTableParserFail15(t *testing.T) {
	require.Equal(t, tableTests[14].out, Uniq(tableTests[4].in, Options{I: true, F: 1, U: true}), "The two slices should be the same.")

}
func TestTableParserFail16(t *testing.T) {
	require.Equal(t, tableTests[15].out, Uniq(tableTests[0].in, Options{I: true, F: 1, S: 3}), "The two slices should be the same.")

}
func TestTableParserFail17(t *testing.T) {
	require.Equal(t, tableTests[16].out, Uniq(tableTests[10].in, Options{I: true, F: 1, S: 3, C: true}), "The two slices should be the same.")

}

func TestTableParserFail118(t *testing.T) {
	require.Equal(t, tableTests[17].out, Uniq(tableTests[14].in, Options{I: true, F: 1, S: 3, D: true}), "The two slices should be the same.")

}
func TestTableParserFail19(t *testing.T) {
	require.Equal(t, tableTests[18].out, Uniq(tableTests[14].in, Options{I: true, F: 1, S: 3, U: true}), "The two slices should be the same.")
}
