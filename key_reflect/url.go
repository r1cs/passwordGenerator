package key_reflect

import (
	"log"
	"math/big"
)

type ASCII uint8

func (a ASCII) String() string {
	return string(a)
}

var UrlKeyMap UrlKeys

func init() {
	//97->122 a->z
	letterL := make([]ASCII, 26)
	for i := range letterL {
		letterL[i] = ASCII(uint8(i) + 97)
	}
	//65->90 A->Z
	letterC := make([]ASCII, 26)
	for i := range letterC {
		letterC[i] = ASCII(uint8(i) + 65)
	}

	//48 -> 57 0->9
	num := make([]ASCII, 10)
	for i := range num {
		num[i] = ASCII(uint8(i) + 48)
	}
	//"," add a specific ascii
	comma := ASCII(44)
	UrlKeyMap = append(letterL, letterC...)
	UrlKeyMap = append(UrlKeyMap, num...)
	UrlKeyMap = append(UrlKeyMap, comma)
	log.Printf("urk key map register: %s", UrlKeyMap)
}
type UrlKeys []ASCII

func(k UrlKeys)Size()*big.Int{
	return new(big.Int).SetInt64(int64(len(k)))
}

func(k UrlKeys)IndexAt(index int64)string{
	return k[index].String()
}