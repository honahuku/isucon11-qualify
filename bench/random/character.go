package random

import "math/rand"

func Character() string {
	return CharacterData[rand.Intn(len(CharacterData))]
}

// 昇順ソート済み
var CharacterData = []string{
	"いじっぱり",
	"うっかりや",
	"おくびょう",
	"おだやか",
	"おっとり",
	"おとなしい",
	"がんばりや",
	"きまぐれ",
	"さみしがり",
	"しんちょう",
	"すなお",
	"ずぶとい",
	"せっかち",
	"てれや",
	"なまいき",
	"のうてんき",
	"のんき",
	"ひかえめ",
	"まじめ",
	"むじゃき",
	"やんちゃ",
	"ゆうかん",
	"ようき",
	"れいせい",
	"わんぱく",
}
