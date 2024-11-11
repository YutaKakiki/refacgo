package loadfile

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func LoadFile(filepath string) ([]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// バッファを持った空のバイト配列
	var buf bytes.Buffer
	// 行単位でファイルを読み込む
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
