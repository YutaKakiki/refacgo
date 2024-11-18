package utils

import (
	"errors"
	"regexp"
)

// マークダウンテキストにおけるコードブロックの内容を抜き出す正規表現
var codeBlockRegex = regexp.MustCompile("```[a-zA-Z ]*\n*([\\s\\S]*?)```")

func DevideCodeAndText(mix string) (code, text string, err error) {
	m := codeBlockRegex.FindStringSubmatch(mix)
	// コードブロックが含まれていなかった場合はエラーを返す
	if len(m) == 0 {
		return "", "", errors.New("failed to match codeblock in md text")
	}
	// コードブロックにマッチした文字列
	code = m[1]
	// mixからコードブロックを削除した文字列
	text = codeBlockRegex.ReplaceAllString(mix, "")

	return code, text, nil
}
