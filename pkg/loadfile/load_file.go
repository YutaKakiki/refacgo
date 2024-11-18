package loadfile

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

// コマンドラインの入力から受け取ったファイルパスからファイルを読み込むための関数
func LoadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
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
