package loadfile

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// アプリケーション内で利用する内部リソースを読み込む関数
func LoadInternal(relativePath string) ([]byte, error) {
	// 呼び出しもとのパスを得る
	_, caller_path, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("failed to get caller information")
	}
	basePath := filepath.Dir(caller_path)
	absPath := filepath.Join(basePath, relativePath)
	f, err := os.Open(absPath)
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

// コマンドラインの入力から受け取ったファイルパスからファイルを読み込むための関数
func LoadExternal(filename string) ([]byte, error) {
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
