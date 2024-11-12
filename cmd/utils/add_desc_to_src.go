package utils

import "fmt"

func AddDescToSrc(src []byte, desc string) []byte {
	// ソースコードと感覚を空けておくために改行を追加
	b := []byte(fmt.Sprintf("%v :\n\n\n", desc))
	src = append(b, src...)
	return src
}
