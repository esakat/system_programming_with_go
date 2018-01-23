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

| var | *io.Reader* | *io.Writer* | *io.Seeker* | *io.Closer* | *io.ReaderAt* |
|:---:|:-----------:|:-----------:|:-----------:|:-----------:|:-------------:|
| os.Stdin | ○ |   |   | ○ |   |
| os.File  | ○ | ○ | ○ | ○ |   |
| net.Conn | ○ | ○ |   | ○ |   |
| bytes.Buffer | ○ | ○ |   |   |   |
| bytes.Reader  | ○ |   | ○ |   | ○ |
| strings.Reader | ○ |   | ○ |   | ○ |

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


### ファイル入力(os.File)

ファイルからの入力は`os.File`を利用する<br />
ファイルの新規作成は`os.Create()`を使用する。`os.Open()`を使うと既存のファイルを開ける<br >
この2つの処理は内部的には同じ、`os.OpenFile()`という関数のフラグ違いエイリアスで、同じシステムコールが呼ばれている

```Go
func Open(name string) (*File, error) {
  return OpenFile(name, O_RDONLY, 0)
}

func Create(name string) (*File, error) {
  return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}
```

`io.Copy()`を使って標準出力にファイルの内容を書きだし

```Go
package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
```

`defer`は`確実に行う後処理`を実行してくれる機能<br>
`defer XXXXXXX`は現在のスコープが終了したら、`XXXXXXX`処理を実行する

### ネットワーク通信の読み込み(net.Conn)

ネットワーク経由のやり取りは、送信データを送信者側から見ると、書き込みで<br>
受信者から見ると読み込みになる

前章で書いた処理も受信者からすると、読み込みになる

```Go
package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}
```

`net.Dial()`で返される`conn`が`net.Conn`型でこれを`io.Copy`を使って、標準出力に渡している<br>
これはシンプルだが、毎回HTTPの取得結果をRFCにしたがってパースするのは大変<br>
Go言語では、HTTPレスポンスをパースする、`http.ReadResponse()`がある<br>
この関数に`bufio,Reader`でラップした、`net.Conn`を渡すと、http.Response構造体のデータが返される

```Go
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP1.0\r\nHost: ascii.jp\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// ヘッダーを表示
	fmt.Println(res.Header)
	// ボディを表示
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

/*
実行結果
map[Content-Type:[text/html] X-Frame-Options:[SAMEORIGIN] X-Ua-Compatible:[IE=edge;IE=11;IE=10;IE=9] Content-Length:[135] Date:[Sun, 21 Jan 2018 23:35:54 GMT] Server:[Apache] Expires:[0]]
<html>
<head>
<meta http-equiv='refresh' content='1; url=http://ascii.jp/&arubalp=e95690d3-ceae-4a5a-a976-5065e0a815'>
</head>
</html>
*/
```

### メモリに蓄えた内容を読み込むバッファ

バッファは書き込みだけでなく読み込みにも使えるよってこと(`bytes.Buffer`)<br>
他に`bytes.Reader`と`strings.Reader`があるけど、ほとんどのケースで使い分けないので、`bytes.Buffer`さえ覚えておけばおk<br>
(これらが使われるのバイナリ解析用くらい)

バッファの初期化はいくつかやり方がある、初期データあるか,初期データの型は？

```Go
// 空バッファ
var buffer1 bytes.Buffer
// バイト列で初期化
buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})
// 文字列で初期化
buffer3 := bytes.NewBufferString("初期文字列")
```

一番上の空バッファの初期化だけは、ポインタではなく、実体なので`io.Writer`などに渡す際は`&buffer1`のようにポインタ値を渡す必要がある

## バイナリ解析関連

バイナリデータを読み込む時に必要な機能を紹介から

### 必要な部位を切り出す(io.LimiReader / io.SectionReader)

ファイルを先頭から必要なとこまでだけ読み込みたいという場合の方法
(ヘッダー領域を読み込むような)

```Go
// たくさんデータがあっても先頭16バイトしか読み込まない
lReader := io.LimitReader(reader, 16)
```

長さだけでなく、スタート位置も固定したい場合がある  
PNGファイルなど、データがいくつかのチャンク(データの塊)に別れている場合  
チャンク毎にReaderを用意して読み込めれば、コードの独立性が高まる  

その時に便利なのが、`io.SectionReader`  
これは`io.Reader`が使えず、代わりに`io.ReaderAt`を使う

```Go
package main

import (
  "os"
  "io"
  "strings"
)

func main() {
  // 文字列からSectionの部分だけを切り出したReaderを作成し、それをos.Stdoutで書きだしている
  reader := strings.NewReader("Example of io.SectionReader\n")
  sectionReader := io.NewSectionReader(reader, 14, 7)
  io.Copy(os.Stdout, sectionReader)
}
```

### エンディアン変換

バイナリ解析ではエンディアン解析が必須、現在のCPUの主流はリトルエンディアン(小さい桁からメモリ格納)  
しかしネットワーク上で転送されるデータの多くはビッグエンディアンのため、変換作業が必要になることが多々ある

任意のエンディアン数値を環境のエンディアンに変換するには`encoding/binary`パッケージを使う

```Go
package main

import (
  "bytes"
  "encoding/binary"
  "fmt"
)

// data変数に格納されたビッグエンディアンの10000という値が、他のエンディアン環境で10000で表示される
func main() {
  // 32ビットのビッグエンディアンデータ(10000)
  data := []byte{0x0, 0x0, 0x27, 0x10}
  var i int32
  // エンディアン変換
  binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
  fmt.Printf("data: %d\n", i)
}
```

### PNGファイルの分析

PNGファイルはバイナリフォーマットで先頭８バイトはシグニチャ(固定バイト列)  
それ以降はチャンクで構成されている

各チャンクと、その長さを列挙して見る  
読み込み用データは[これ](https://en.wikipedia.org/wiki/File:Lenna.png)を使用  
![これ](https://upload.wikimedia.org/wikipedia/en/2/24/Lenna.png)

コードは`png_parse`ディレクトリを参照

Go言語では配列に要素を追加するには  
```Go
array = append(array, elemet)
```

とする、多くのオブジェクト言語と違い、配列やスライス自体にメソッドを持たない

### png画像に秘密のテキストを入れる

PNG画像にはテキストを追加する`tEXt`というチャンクがある(それに圧縮をかけた`zTXt`もある)  
これらは画像に埋め込まれるだけで、表示はされない

`png_secret_msg`で先ほどの画像に隠し文字を入れて見る

## テキスト解析

テキスト解析ではバイナリのようにデータサイズが固定ではないため、探索しながら読み込み

## データ型を指定して解析

`io.Reader`から読み込んだデータを整数や浮動小数点に変換するには`fmt.Fscan()`を使う  
これを使うと、任意のデータ区切りをフォーマット文字列として指定できる

```Go
var source = "123 1.234 1.0e4 test"

func main() {
	reader := strings.NewReader(source)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
}
```

Go言語は型情報をデータが持っているので、`%v`と書いておけば、変数の型を読み取って変換してくれる。

## その他の決まった形式のフォーマット文字列の解析

`encoding`パッケージを使えば、決まった形式の文字列も扱える(例えばcsvとか)


