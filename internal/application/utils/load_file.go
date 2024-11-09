package utils

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func LoadFile(filepath string) ([]byte, error) {
	// geminiの送信MAXの容量がわからんのでまた後々。
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var buf bytes.Buffer
	// ファイルを読み込む
	scannar := bufio.NewScanner(f)
	// 一行ごとにループ
	for scannar.Scan() {
		// bufに書き込み
		buf.Write(scannar.Bytes())
		// 文末に改行
		buf.WriteByte('\n') //1バイト書き込む
	}
	return buf.Bytes(), nil
}
