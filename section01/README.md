# 1: Webクライアント

このセクションでは、GoでWebサーバーとクライアントを書くために知っておく必要があるすべてを学習します。

`net/http` パッケージはHTTPリクエストを送信するのに役立つ一連の関数と型を提供します。
最も重要な型は次のとおりです:

- [Client](https://golang.org/pkg/net/http/#Client)
- [Request](https://golang.org/pkg/net/http/#Request)
- [Response](https://golang.org/pkg/net/http/#Response)

これらの型がどのように機能するか短い時間ですが見ていきます。
しかしその前に、ここでも私たちの生活を楽にするヘルパー関数があることを認識することが重要です。

## Getメソッド

これらのヘルパー関数のひとつが [`Get`](https://golang.org/pkg/net/http/#Get) です。

```go
package main

import (
        "fmt"
        "log"
        "net/http"
)

func main() {
        res, err := http.Get("https://golang.org")
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(res.Status)
}
```
[ソースコード](examples/get.go)

このプログラムはGoのホームページにGETメソッドのHTTPリクエストを送信して、レスポンスのステータスコードを出力します。
何かが間違っている場合は、エラーをログに記録してプログラムの実行を停止します。

URLの値を変更して他のコードでは何が取得できるかを確認してください。
いくつかのアイデアで試すことができます:

- https://golang.org/foo
- https://thisurldoesntexist.com
- https:/thisisnotaurl

### ステータスコードの演習

`Get` 関数を使用すると `Response` と `error` を返します。
`Response` のドキュメントを読んで `get.go` を修正し、レスポンスのステータスコードが404のときだけメッセージを出力してください。

### bodyの演習

`Response` 型は `io.ReadCloser` 型の `Body` フィールドも保有しています。
このフィールドの型はそれで何ができるのかを知る大きなヒントです: それを読み込んで閉じます。

`get.go` プログラムを修正し、`Response` のbodyを標準出力に出力してください。
メモリリークするのを防ぐために `Body` はプログラムの最後で閉じなければならないことを思い出してください！

## 他のメソッド

私たちが使用した `Get` メソッドは `http.DefaultClient` の `Get` メソッドを呼び出すヘルパー関数です。
`Client` 型は私たちが知っている全てのHTTPメソッドに直接関係するいくつかの他のメソッドを提供します:

- [Client.Get](https://golang.org/pkg/net/http/#Client.Get)
- [Client.Post](https://golang.org/pkg/net/http/#Client.Post)
- [Client.PostForm](https://golang.org/pkg/net/http/#Client.PostForm)
- [Client.Head](https://golang.org/pkg/net/http/#Client.Head)

同等のヘルパー関数:

- [http.Get](https://golang.org/pkg/net/http/#Get)
- [http.Post](https://golang.org/pkg/net/http/#Post)
- [http.PostForm](https://golang.org/pkg/net/http/#PostForm)
- [http.Head](https://golang.org/pkg/net/http/#Head)

何かいくつかの他のメソッドを使いたいですか？
[`Do`](https://golang.org/pkg/net/http/#Client.Do) メソッドを参照してください。
このメソッドは引数として `http.Request` のポインターを受け取り、`http.Response` と `error` を返します。

`Request` メソッドはあらゆる種類のHTTPリクエストを送信するために必要なすべての表現を提供します。

例えば、当初の `get.go` プログラムに相当するリクエストを作成することができます:

`Client` 型は与えられた `Request` を送信する `Do` メソッドを公開し、`Response` と `error` を返します。

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://golang.org", nil)
	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request failed: %v", err)
	}
	fmt.Println(res.Status)
}
```
[ソースコード](examples/do-get.go)

これもまた `Request` に `Body` を提供するためで、`Response` で持っていたものと似ています。
大きな疑問は `io.Reader` を作る方法ですか？

それは、あなたが読み込みたいものの型に依存します。
もし `string` を提供したいのであれば、それはとても簡単です。
[`strings.NewReader`](https://golang.org/pkg/strings/#NewReader) で `string` から `io.Reader` を作成することができます。

### PUTリクエストの演習

上記のコードを修正してPUTリクエストを https://http-methods.appspot.com/YourName/Message に送信してください。
`YourName` をあなたの名前、または誰も使用していないユニークなものに置き換えてください。
これによりリクエストの本文に送信した文字列が保存され、後で取得することができます:

```
    $ curl https://http-methods.appspot.com/YourName/Message
```

## パラメーター: クエリーとフォーム

https://http-methods.appspot.com の背後にあるwebアプリケーションは名前空間の配下にあるすべてのキーの出力をサポートしています。
https://http-methods.appspot.com/Hungary/ にアクセスすると、その名前空間のすべてのキーを見ることができます。
値も表示するために、URLに `?v=true` を追加することができます: https://http-methods.appspot.com/Hungary/?v=true

### クエリーの演習

プログラムを書き、名前空間のすべてのキーと値を取得して表示してください。
`?v=true` を `NewRequest` に渡したURLの最後に追加するのではなく `Do` に渡す `Request` に直接追加する方法を見つけてください。

# おめでとうございます！

よくできました！ これでHTTPのクライアント、リクエスト、レスポンスについて（今日）知るべきことはすべて知っています。
さて、これをサーバー側でどのように処理するのかを理解するために、[次のセクション](../section02/README.md) に進みましょう。
