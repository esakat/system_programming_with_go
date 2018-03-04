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