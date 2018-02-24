# 8章のメモ(UNIXドメインソケット)

## UNIXドメインソケットとは

コンピュータ内部でしか使えない代わりに高速な通信が可能なソケット  
TCP型とUDP型がある


一般的なTCP,UDP通信では外部に繋がるインタフェースに接続するが  
Unixドメインソケットの場合は専用の高速な通信が可能なインタフェースへ接続する

Unixドメインソケットを開くにはファイルシステムのパスを指定する  
サーバプロセスを立ち上げると、指定したパスにファイルが作成される。
クライアントはそのファイルへ接続するようになる

クライアント側はIPとポートで接続先を決めるのではなく、ファイルパスで接続先を決めるようになる

ここで作られるファイルはソケットファイルと呼ばれる特殊なファイルで、通常のファイルのように実体がない。

## Unixドメインソケットの使い方

クライアント
```go
conn, err := net.Dial("unix", "socketfile")
if err != nil {
  panic(err)
} 
// connを使った読み書き
```

サーバ
```go
listener, err := net.Listern("unix", "socketfile")
if err != nil {
  panic(err)
}
defer listener.Close() // Go言語だと、クローズしたらsocketfileは削除される
conn, err := listener.Accept()
if err != nil {
  // エラー処理
}
// connを使った読み書き
```

## ベンチマーク取ってみる

対象を全て同一フォルダに入れて
```
$ go test -bench .
```
