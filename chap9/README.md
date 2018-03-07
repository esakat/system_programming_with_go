# 9章のメモ(ファイルシステムの基礎)

## ファイルシステムとは

コンピュータにデータを保存すると、ストレージのどこかにビットの羅列が記録されるはず  
それにアクセスするには、決まったルールなどで管理しておかないと難しいだろう  
そのための機能が`ファイルシステム`

## ファイルシステムの基礎

ストレージに対するファイルシステムでは、ストレージの全領域を512バイトから4kバイトの固定長データ配列として扱う  
ここにファイルの実データだけでなく、管理情報も格納する。この管理情報はLinuxではinodeと呼ばれている

inodeにはユニークな識別子がついており、inodeにアクセスできればファイルの実データにもアクセスができる  
以下のコマンドでinodeを確認できる

```bash
$ ls -i
```

### 実際のファイルシステム

実際は複雑に入り組んだファイルシステムで構成されている  
(例えばLinuxだと/proc以下は仮想的なファイルシステムになっている)

LinuxではVFSというAPIで統一的に扱えるようにしている。

## ファイル/ディレクトリを扱うGo言語の関数

### ディレクトリの作成

```go
// フォルダを1階層だけ作成
os.Mkdir("setting", 0755)

// 深いフォルダを1回で作成
os.Mkdir("setting/myapp/networksettings", 0755)
```

### ファイルの削除/移動/リネーム

`os.Remove()`=ファイルや子要素を持たないディレクトリの削除(rmとrmdirの兼用)  
対象を全て再帰的に削除するときは、`os.RemoveAll()`を使う

```go
// ファイルや空のディレクトリの削除
os.Remove("server.log")

// ディレクトリを中身ごと削除
os.Remove("workdir")
```

#### 特定の長さでファイルを切り落とす os.Truncate()

```go
// 先頭100バイトできる
os.Truncate("server.log", 100)

// Truncateメソッドを利用する場合
file, _ := os.Open("server.log")
file.Truncate(100)
```

#### ファイルの移動 os.Rename()

mvコマンドと違って、移動先ディレクトリ内での名前も指定する必要がある

```go
// リネーム
os.Rename("old_name.txt", "new_name.txt")

// 移動
os.Rename("oldfir/file.txt", "newdir/file.txt")

// 移動さきはディレクトリ✖️
os.Rename("olddir/file.txt", "newdir/") // エラー
```
マウントデバイスが異なる場合も、renameはエラーがでる

```go
err := os.Rename("sample.rst", "/tmp/sample.rst")
if err != nil {
  panic(err)
  // ここが実行され、コンソールにエラーが表示される
  // rename sample.rst /tmp/sample.rst: cross-device link
}
```

デバイスやドライバが異なる場合はファイルを開いてコピーする必要がある

```go
oldFile, err := os.Open("old_name.txt")
if err != nil {
  panic(err)
}
newFile, err := os.Create("/other_device/new_name.txt")
if err != nil {
  panic(err)
}
defer newFile.Close()
_, err := io.Copy(newFile, oldFile)
if err != nil {
  panic(err)
}
oldFile.Close()
os.Remove("old_name.txt")
```

## ファイルの属性取得

ファイルの属性は`os.Stat()`, `os.LStat()`で取得できます。  
これらは対象がシンボリックリンクの場合挙動が異なる

`os.Stat()`はシンボリックリンクの場合、リンク先情報を取得できる  
`os.LStat()`はリンクの情報を取得できる

`os.File()`で取得している場合は、`Stat()`で取得できる

### FileModeタイプについて

FileModeの実体は32ビットの非負の整数ですが、メソッドがいくつか使える

## ファイルの存在チェック

`os.Stat()`はファイルの存在チェックによく使われる

```go
info, err := os.Stat(ファイルパス)
if err == os.ErrNotExist {
  // ファイルが存在しない
} else if err != nil {
  // それ以外のエラー
} else {
  // 正常ケース
}
```

存在チェックそのものはシステムコールで提供されていない  
多言語でも、`os.Stat()`などでファイル存在をチェックしている

## OS固有のファイル属性を取得する

OS固有の情報を取得するには以下のコマンド

```go
// Windows
internalStat := info.Sys().(syscall.Win32FileAttributeData)

// Windows以外
internalStat := info.Sys().(*syscall.Stat_t)
```

## ファイルの同一性チェック

`os.FileInfo`が参照するファイルが同一か判定できる関数がある   
同一判定は内容が同じという判定ではなく、全く同じ実体を見ているかという判定

```go
if os.SameFile(fileInfo1, fileInfo2) {
  fmt.Println("同じファイル")
}
```

## ファイル属性の設定

```go
// ファイルのモードを変更
os.Chmod("setting.txt", 0644)

// ファイルのオーナーを変更
os.Chown("setting.txt", os.Getuid(), os.Getgid())

// ファイルの最終アクセス日時と変更日時を変更
os.Chtimes("setting.txt", time.Now(), time.Now())
```

## リンク

ハードリンク、シンボリックリンクの作成もGoで可能
```go
// ハードリンク
os.Link("oldfile", "newfile")

// シンボリックリンク
os.Symlink("oldfile", "newfile")

// シンボリックリンクのリンク先を入手
link, err := os.ReadLink("newfile-symlink")
```

WindowsでもVista以降は気軽にシンボリックリンク作れるよ

## ディレクトリ情報

ディレクトリ一覧取得はosパッケージにはない
os.Openで開いて、os.Fileのメソッドでディレクトリ内のファイル一覧を取得する

# OS内部のファイル高速化

通常は意識することはないけど、データベース管理システムを実装するようなケースだと意識する必要あり  
ディスクへの書き込みは遅い処理のため、なるべく最後までやらないようにしたい。  
LinuxではVFSの内部に設けられたバッファを使うことで、ディスクへの操作をなるべく回避している

ディスクにデータを書き込む時、一旦バッファへデータを格納します。(その時点でアプリに結果を返す)  
データ参照時にバッファにデータが残っていれば、バッファからデータを取るようにしている。
(バッファと実際のストレージの同期はアプリケーションが知らないところで非同期に行われている)

Go言語で、ストレージへの書き込みを確実に行いたいときは`file.Sync()`を使う

# ファイルパスとマルチプラットフォーム

アプリケーションのファイルアクセスは「どのファイル」に対して、「何をするか」で説明可能  
この「どのファイル」というのを指定するのがファイルパス

ファイルパス表記方法はLinuxとかのPOSIX表記と、Windowsでの表記がある

## Go言語でのパスの扱い

* path/filepath : osのファイルシステムに使う
* path: URLに使う

ファイルやディレクトリに対するパス表記を操作するには、path/filepathパッケージを使う  
path/filepathであれば、WindowsもPOSIX表記の違いも吸収して扱える

pathの方は常に`/`で区切るパス表記に対して使うパッケージで  
URLに対して使われる

