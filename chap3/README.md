# 3章のメモ

前章では`io.Writer`について、この章では`io.Reader`について説明

## io.Reader

`io.Writer`と同様に`io.Reader`インタフェースで外部からのデータ読み込みが抽象化されている

以下が`io.Reader`の`Read()`メソッド

```Go
type Reader interface {
  func Read(p []byte) (n int, err error)
}
```

引数の`p`は読み込み内容を一時的にいれておくバッファ、あらかじめメモリを用意しておいて、それを使う<br />
(Goでメモリを確保するには`make()`を使う)

こんな感じ
```Go
// 1024バイトのバッフォを作成
buffer := make([]byte, 1024)
// sizeは実際に読み込んだバイト数、errはエラー
size, err := r.Read(buffer)
```

これだと毎回バッファを作成して、読み込みのたびに引数を指定する必要があるが、不便
Go言語では補助の機能が豊富に提供されている

## io.Readerの補助関数

PythonやRubyではfileオブジェクトが補助関数を持ったりするけど<br />
Goの場合は例外を除き、外部のヘルパー関数を利用する

### ioutil.ReadAll()

終端記号が来るまで、すべてのデータを読み込んで返す<br />
メモリに収まりきらない場合を除いて、基本的にこれでいける

```Go
buffer, err := ioutil.ReadAll(reader)
```

### io.ReadFull()

決まったバイト数だけ読み込みたい場合に利用<br />
これを使うと指定したサイズ分まで読み込めない場合はエラーが返ってくる

```Go
// 4バイト読み込めないとエラーになる
buffer := make([]byte, 4)
size, err := io.ReadFull(reader, buffer)
```


