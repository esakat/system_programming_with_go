# 13章のメモ(Go言語と並列処理)

## 並列と並行のちがい

### 並行

CPU数、コア数の限界を超えて複数の仕事を同時に行う

シングルコアでも複数の作業を同時におこなるのは並行処理  
1つの作業でプログラム全体がブロックされてしまうのを防ぐため

### 並列

複数あるCPUのコアを効率よく扱う

## goでの並列処理

goroutineとチャネルで実現できる

チャネルはデータの入出力を直列化できる。

複数のgoroutineでデータを処理しても、チャネルを経由するだけで、複雑なロックが不要になる