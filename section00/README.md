# 0: Hello world

webサーバーを書き始める前に、簡単なGoプログラムを分析してみましょう。

[embedmd]:# (hello/main.go /package main/ $)
```go
package main

import    "fmt"

func         main() {
	fmt.Println("hello, world!")        }
```

あなたはこのプログラムのすべての行を理解できるはずです。
もしそうでないなら、[Goツアー][1] をチェックする時間でしょう。

## コードを実行する

この時点で、上記のコードを `main.go` という名前のファイルに貼り付けることができます。
このファイルはあなたの `$GOPATH` のどこにあっても構いません。
私は `$GOPATH/src/hello/main.go` というパスを持つフォルダーの中にあると仮定します。

あなたはコードをいくつかの方法で実行できるようになりました。
`main.go` を含むディレクトリーから次を実行します:

- `go run main.go` は `main.go` だけをコンパイルして実行します。

- `go install` はカレントディレクトリーのコードをコンパイルし、`$GOPATH/bin` に `hello` という名前のバイナリーを生成します。

## コードをフォーマットする: gofmt

もしあなたがいくつかのGoコードを見慣れているなら、このコードは標準的な方法でフォーマットされていないことに気づいたかもしれません。

`gofmt main.go` を実行してみると、出力がどのように表示されるかわかります。
`gofmt -d main.go` を実行すると、ファイルがフォーマットされたバージョンとどのように違うのかを見ることできます。
最後に `gofmt -w main.go` を実行すると、単にフォーマットされたバージョンでファイルを書き換えることができます。

## import文を管理する: goimports

コンパイラーがあなたの使用するパッケージを知るためにプログラムの先頭にimport文が必要です。
しかし、これは自分で書く必要があることを意味していません！

`main.go` から `import "fmt"` 行を削除して実行すると、エラーを確認できるはずです:

```bash
$ go run main.go
# command-line-arguments
./main.go:5: undefined: fmt in fmt.Println
```

あなたは不足しているimport文を手動で追加したり `goimports` を使うことでこのエラーを修正できます。

あなたのマシンに `goimports` がインストールされていないなら、次を実行することで簡単にインストールできます:

```bash
$ go get golang.org/x/tools/cmd/goimports
```

これは `$GOAPTH/bin` に `goimports` のバイナリーをインストールします。
基本的にはGitHubからコードを取得して `go install` を実行するのと同じです。

あなたは `gofmt` の代わりに `goimports` を実行できるようになりました。
`goimports` はコードをフォーマットします*し*、不足しているパッケージは追加して、未使用のパッケージは削除して、あなたのインポートしたパッケージリストを修正します。

`gofmt` と同様に次を実行できます:

- `goimports main.go` は、修正されたファイルがどのように表示されるかわかります。
- `goimports -d main.go` は、現在のバージョンと修正されたバージョンの違いが表示されます。
- `goimports -w main.go` は、ファイルを修正されたバージョンで書き換えます。

## あなたの人生を楽にする

`gofmt` と `goimports` はあなたの人生を楽にするために使用するかもしれないツールの2つに過ぎません。
これらのツールがファイルを保存するたびに呼び出されるようにするには、お気に入りのエディターにプラグインを追加することを検討する必要があります。
私は個人的には [vim][2] と [vim-go][3] が好きですが、Goプラグインを追加した [VS Code][4] も使用しています。
しかしGoをサポートするエディターはもっとたくさんあります。 あなたは [ここ][5] でそれらを見つけることができます。

# おめでとうございます！

あなたはセクション...0で終わりました。 さて、どこかで始める必要があります！
しかしですね、少なくとも今あなたは [セクション1][6] に取り組む準備ができています。

[1]: https://tour.golang.org
[2]: http://www.vim.org/
[3]: https://github.com/fatih/vim-go
[4]: https://code.visualstudio.com
[5]: https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins
[6]: ../section01/README.md
