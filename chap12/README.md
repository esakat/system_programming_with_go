# 12章のメモ(シグナルによるプロセス間の通信)

この章ではプロセスに対して送られるシグナルを学ぶ

## シグナルの用途

### プロセス間通信

カーネルが仲介して、あるプロセスから、別のプロセスに対してシグナルを送ることができる。
自分自身に対しても送れる

### ソフトウェア割り込み

システムで発生したイベントはシグナルとして、プロセスに送られる。
シグナルを受け取ったプロセスは現在行なっているタスクを中断して、あらかじめ登録しておいた登録ルーチンを実行する

システムコールを受け取るカーネルは常に起動しているが、シグナルを受け付けるプロセスは停止していることもある
シグナルはその辺りも考慮して作られている。

## シグナルの種類

全て`SIG`から始まる
以下のコマンドで確認可能

```bash
$ man signal
```

### ハンドルできないシグナル

SIGKILLとSIGSTOPの2つはハンドルできないので注意

```bash
// プロセスID指定でKILL
$ kill -KILL 11111
// プロセス名を指定でSTOP
$ pkill -STOP ./sample
↓
$ fg ./sample //これでまた再開できる
```

### サーバーでハンドルするシグナル

`SIGTERN`: killコマンドなどがデフォで送るシグナル、プロセスの停止
`SIGHUP`: 設定ファイルの再読み込みなどに使われる

### コンソールでハンドルするシグナル

`SIGINT`: `Ctrl+C`で送られるシグナル
`SIGQUIT`: `Ctrl+\`で送られるシグナル
`SIGTSTP`: `Ctrl+Z`で送られるシグナル
`SIGCONT`: バックグラウンド動作から戻させる指令
`SIGWINCH`: ウィンドウサイズ変更
`SIGHUP`: 模擬ターミナルから切断されるときに呼ばれるシグナル

## Go言語でのシグナル

```go
var {
  Interrupt Signal = syscall.SIGINT
  kill      Signal = syscall.SIGKILL
}
```
これらが定義されている

## シグナルの応用例

サーバー系だと簡単にシャットダウンできない(アクセス中のユーザーへ正常終了返すまで。。) 
複数台サーバだとさらに難しくなる(グレイスフルリスタート)

この仕組み実現の為にServer:Starterというのがある  
サーバーの再起動が必要になったら、新しいサーバーを起動して新しいリクエストをそちらに流しつつ、古いサーバーのリクエストが終了したら、正しく終了する。

goでも実装されている

```bash
$ go get github.com/lestrrat/go-server-starter/cmd/start_server
```
以下はserverというプログラムグをServer:Starterで起動させている
```bash
$ start_server --port 8080 --pid-file -- ./server
```