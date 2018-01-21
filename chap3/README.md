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

### io.Copy()

`io.Reader`から`io.Writer`にそのままデータを渡したいときはコピー系の補助関数を利用する<br />
`io.Copy()`ではすべてを読み込んで書き込む(ファイルを開いてHTTP転送や、ハッシュ計算も可能)<br />
指定したバイト数のみ読み込む場合は`io.CopyN()`が使える

```Go
// すべてコピー
writeSize, err := io.Copy(writer, reader)
// 指定したサイズのみをコピー
writeSize, err := io.CopyN(writer, reader, size)
```

### io.CopuBuffer()

あらかじめコピーする量が決まっていて、バッファを使いまわしたい場合に利用する<br />
`io.Copy()`はデフォだと、32KBのバッファを確保している

```Go
// 8kbのバッファを利用
buffer := make([]byte, 8 * 1024)
io.CopyBuffer(writer, reader, buffer)
```

## io.Writer,io.Reader以外の入出力インタフェース

読み書き以外にもクローズ処理など様々な処理が必要<br />
以下のようなものがあるよ

* io.Closerインタフェース
  * `func Close() error`メソッドを持つ
  * 使用し終えたファイルを閉じる
* io.Seekerインタフェース
  * `func Seek(offset int64, whence int) (int64, error)`メソッドを持つ
  * 読み書き位置を移動する
* io.ReaderAtインタフェース
  * `func ReadAt(p []byte, off int64) (n int, err error)`メソッドを持つ
  * 対象となるオブジェクトがランダムアクセスを行える場合に、好きな位置を自由にアクセスする際に使用する

## 入出力の複合インタフェース

`io.Closer`や`io.Seeker`だけを満たした、構造体を扱うことはほとんどない<br />
`io.Writer`などを組合わせた複合インタフェースがほとんど

例えば,`io.ReadWriter`は`io.Reader`と`io.Writer`を満たしている必要がある

### 複合インタフェースのキャスト

`io.ReadWriter`が要求されているが`io.Reader`を満たしていない場合がある<br />
例えば、ソケット読み込み関数を作成中に、引数は`io.ReadCloser`だが<br />
ユニットテスト用に、`io.Reader`インタフェースを満たすものが使いたいケースがある<br />
その場合は`ioutil.NopCloser()`を使うと、ダミーの`Close()`を持って`io.ReadCloser`のふりをするラッパーオブジェクトが得られる

```Go
import (
  "io"
  "io/ioutil"
  "strings"
)

var reader io.Reader = strings.NewReader("テストデータ")
var readCloser io.ReadCloser = ioutil.NopCloser(reader)
```

`bufio.NewReadWriter()`を使うと`io.Reader`と`io.Writer`を繋げて、`io.ReadWriter`型のオブジェクトを作れる

```Go
import "bufio"

var readWriter io.ReadWriter = bufio.NewReadWriter(reader, writer)
```

## よく使われるio.Readerを満たす構造体

```memo
<memo>
Go言語では多くの構造体がio.Writerとio.Reader両方を満たしているので
それぞれ用意する必要はほとんどないよ
```

| var | *io.Reader* | *io.Writer* | *io.Seeker* | *io.Closer* |
|:---:|:-----------:|:-----------:|:-----------:|:-----------:|
| os.Stdin | ○ |   |   | ○ |
| os.File  | ○ | ○ | ○ | ○ |
### 標準入力(os.Stdin)

標準入力に対応するオブジェクトが`os.Stdin`<br/>
このプログラムを実行すると入力待ちになり、Enterを押すたびに、結果が返ってくる

```Go
package main

import (
  "os"
  "fmt"
  "io"
)

func main() {
  for {
    buffer := make([]byte, 5)
    size, err := os.Stdin.Read(buffer)
    if err == io.EOF {
      fmt.Println("EOF")
      break
    }
    fmt.Printf("size=%d input='%s'\n", size, string(buffer))
  }
}

/* 
結果
$ go run main.go
fsafasf
size=5 input='fsafa'
size=3 input='sf
'
改行が変になる
*/
```

通常で実行すると、入力待ちでブロックしてしまう<br />
Go言語の`Read()`にはタイムアウトの仕組みがないため、このブロックは避けられない

他の言語だと、ノンブロッキング用APIとブロックするAPIが用意されていることが多いが<br />
Goでは並列処理がうまく扱えるため、それを使ってノンブロッキングな処理を実現する<br />
具体的にはgoroutineとチャネルで実現する


