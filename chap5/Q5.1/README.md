結果

```
$ sudo dtruss go run main.go
Password:
dtrace: system integrity protection is on, some features will not be available

SYSCALL(args) 		 = return
dtrace: 1367 dynamic variable drops with non-empty dirty list
open("/dev/dtracehelper\0", 0x2, 0xFFFFFFFFEFBFEC20)		 = 3 0
ioctl(0x3, 0x80086804, 0x7FFEEFBFEB80)		 = 0 0
close(0x3)		 = 0 0
access("/AppleInternal/XBS/.isChrooted\0", 0x0, 0x0)		 = -1 Err#2
thread_selfid(0x0, 0x0, 0x0)		 = 40250 0
bsdthread_register(0x7FFF7290EC50, 0x7FFF7290EC40, 0x2000)		 = 1073742047 0
mprotect(0x173A000, 0x1000, 0x0)		 = 0 0
...
```