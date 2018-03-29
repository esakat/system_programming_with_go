package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/edsrzf/mmap-go"
)

func main() {
	// テストデータを書き込み
	var testData = []byte("0123456789ABCDEF")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	err := ioutil.WriteFile(testPath, testData, 0644)
	if err != nil {
		panic(err)
	}

	// メモリにマッピング
	// mは[]byteのエイリアスなので添字アクセス可能
	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// 指定したファイルの内容をメモリ上に展開
	// 1つ目の引数に指定されたファイルの全内容がメモリにマップされる
	// 2つ目で許可する操作を指定
	// 3つ目はフラグで、mmap.ANONを渡すと、メモリにマッピングせず、領域だけ確保する
	m, err := mmap.Map(f, mmap.RDWR, 0)
	if err != nil {
		panic(err)
	}
	// メモリ上に展開された内容を削除して閉じる
	defer m.Unmap()

	// メモリ上のデータを修正して書き込む
	m[9] = 'X' // 数字の9をXに書き換えている

	// 書きかけの内容をファイルへ保存する
	m.Flush()

	// 読み込んでみる
	fileData, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("original: %s\n", testData)
	fmt.Printf("mmap:	%s\n", m)
	fmt.Printf("file: %s\n", fileData)
}
