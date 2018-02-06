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

## POSIXとC言語の標準規格

ここまではOS前までのシステムコール呼び出し方法、ここからはOSより先でのシステムコールの処理について

### POSIXとは

Portable Operating System Interfaceのこと  
OS間で共通のシステムコールを定義して、アプリケーションの移植性を高めるためのIEEE企画

システムコールを呼び出すインタフェース  
(具体的にはC言語の関数名と引数、返り値が定義されている)
(必ずしもシステムコールとインタフェースが1対1で結びついているわけではない)

例えば、ファイル入出力はPOSIXにおいては5つのシステムコールで定義されている  
`open()`,`read()`,`write()`,`close()`,`lseek()`

### OSでのシステムコール

GoではSYSCALLで終わっていた、ここから先はLinuxのソースコードを確認していく

[リポジトリ](https://github.com/torvalds/linux/)

fs/read_write.c#L566
```c
SYSCALL_DEFINE3(write, unsigned int, fd, const char __user *, buf,
		size_t, count)
{
	...
}
```

Linuxカーネルでは、writeシステムコールは上のように定義されている.  
SYSCALL_DEFINE`x`マクロ(xには0~6の数値が入る)

これを展開すると

```c
asmlinkage long sys_write(...)
```

asmlinkageは引数をCPUのレジスタ経由で渡すようにするためのフラグ

通常の関数呼び出しでは、呼び出し側と呼ばれる側がスタックメモリで隣接したメモリブロックに、
それぞれをスコープに含まれるローカル変数を格納する。

システムコールにおいてはメモリの共有ができない(ユーザーモード領域とカーネル領域ではそれぞれ別々にスタックメモリが用意されている)

これを解決するためにレジスタを使用している

`sys_write`は`sys_call_table`配列に格納されている

arch/x86/entry/common.c#L269
```c
__visible void do_syscall_64(struct pt_regs *regs)
{
	struct thread_info *ti = current_thread_info();
	unsigned long nr = regs->orig_ax;
	...
	if (likely((nr & __SYSCALL_MASK) < NR_syscalls)) {
		nr = array_index_nospec(nr & __SYSCALL_MASK, NR_syscalls);
		// axレジスタからシステムコール番号を読みだして、
		// 6つのレジスタを関数の引数として渡している
		regs->ax = sys_call_table[nr](
			regs->di, regs->si, regs->dx,
			regs->r10, regs->r8, regs->r9);
	}
	...
}
```

この後、`call do_syscall_64`が呼び出される
これが扱うのはCPUそのもの

`syscall.Write()`→`SYSCALL`→`レジスタ`→`sys_call_table`という流れになる

### システムコールのモニタリング

macの場合は`dtruss`コマンドでできる

### エラー処理

POSIXのシステムコールであるwrite()ではエラーは次のように定義されている

> 成功した場合、書き込まれたバイト数が返される(ゼロは何も書き込まれなかったことを示す).
> エラーの場合、-1を返し、errnoにエラーを示す値をセットする。

errnoには数値のみ、なぜかというと、システムコールでやっているから  
レジスタ経由、つまり数値しか扱えないのである

これは低レイヤの話で、最近の言語だと例外が使われることが多い
Nodeだと、コールバック関数、Goだとレスポンスにエラーを返すとかね

## まとめ

システムコールはOSしか使えないはずの機能を一般的なアプリケーションでも使える機能だった  
システムコールがないと、アプリケーションはほとんど何もできない(他のプロセスやリソースを使うものが使えない)ということも理解した