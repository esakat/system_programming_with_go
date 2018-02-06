# 5章のメモ(システムコール)

この章ではシステムコールについて掘り下げていく

## システムコールとは

特権モードでOSの機能を呼ぶこと

### 特権モードとは

CPUの動作モードと呼ばれる仕組みでは、モードによって利用できる機能を設定できる  
OSが動作するのが特権モード。 一般的なアプリケーションはユーザモード  
特権モードではほとんど全てのCPU機能が使える　 
(ユーザモードでは機能が制限され、メモリ割り当てやファイル入出力などができない)

### 一般アプリケーションでもメモリ割り当てとかをしたい

多くのOSではシステムコールを介して、アプリケーションでも特権モードのみが使える機能を提供している  
(厳密に言うと、システムコールでは自由にいつでも特権モードの機能を使えるわけではない(一回で1つのプロシージャしか呼び出せない))

### システムコールがないとどうなるか

例えばあるアプリケーションが計算したデータがあるとして  
システムコールがなければ

* 計算結果を画面に表示できない(外部表示は別プロセスが管理。プロセス間通信は特権モードでのみ許可されている)
* 計算結果をファイルに保存できない(ファイル入出力も特権モードのみ許可)
* 共有メモリへ書き出しできない(共有メモリがあれば別プログラムから計算結果を参照可能)(共有メモリ作成は特権も(ry..)
* 計算結果を外部Webサービスにも送れない(外部Webサービスへの通信はカーネルが提供する機能でないと、不可)

## Go言語のシステムコール

前章で紹介した、`os.File`が満たすインタフェース(`io.Reader`など)も最終的にsyscallで定義された関数を呼び出す。

例えばファイルを開くsyscallは下のように定義されている

```Go
func syscall.Open(path String, mode int, perm uint32)
```

`os.Create()`メソッドを下っていってみる

```Go
func Create(name string) (*File, error) {
	return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}
```

`os.Create()`は`os.OpenFile()`を簡易的に呼べるようにしているらしい  
`os.OpenFile()`はfile_unix.goで定義されている  

大切なのは次の行

```Go
r, e = syscall.Open(name, flag|syscall.O_CLOEXEC, syscallMode(perm))
```

syscall.Openはmacの場合、zsycall_darwin_amd64.go内の`Open()`関数が呼び出される

```Go
func Open(path string, mode int, perm uint32) (fd int, err error) {
  var _p0 *byte
  // Go言語での文字列(path引数)をC言語形式の文字列(先頭要素へのポインタ表現)に変換
  // システムコールへは数値しか渡せないため
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
  }
  // Syscallの呼び出し
  // 何番の処理をやるかを最初に指定している。 SYS_OPENは定数
  // 定数はmacの場合、asm_darwin_amd64.sというGo言語の低レベルアセンブリ言語で書かれたコードで定義されている。これ以降はOS側での処理になる
	r0, _, e1 := Syscall(SYS_OPEN, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(perm))
	use(unsafe.Pointer(_p0))
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
```

Go言語で普通にSyscallを呼び出すと、処理時間のかかるスレッドにマーキングをして処理時間をあげる工夫がされている  
スレッド関係の処理を行わないシステムコール呼び出しに`RawSyscall()`がある

