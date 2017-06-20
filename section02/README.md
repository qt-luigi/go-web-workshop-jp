# 2: Webサーバー

このセクションでは、簡単なHTTPサーバーをGoで書く方法を学習します。

これを行うには [`net/http`](https://golang.org/pkg/net/http/) パッケージを使用するので、
リンクをクリックしてドキュメントを参照してください。

### HTTPハンドラーを定義する

`net/http` パッケージは [`HandlerFunc`](https://golang.org/pkg/net/http/#HandlerFunc) 型を定義します:

```go
type HandlerFunc func(ResponseWriter, *Request)
```

この関数型の最初の引数は [`ResponseWriter`](https://golang.org/pkg/net/http/#ResponseWriter) で、
HTTPレスポンスでヘッダーを設定する手段を提供します。
また `io.Writer` インターフェースを満足させる `Write` メソッドも提供しています。

単に `"Hello、web"` を出力に書き出す非常に簡単なHTTPハンドラーを見てみましょう:

```go
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, web")
}
```

ご覧のとおり、最初の引数が `io.Writer` である `fmt.Fprintln` 関数を使用しています。

### HTTPハンドラーを登録する

ハンドラーが定義されたら、それについて `http` パッケージに通知し、いつ実行するかを指定する必要があります。
これを行うには `http.HandleFunc` 関数を使います:

```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

最初の引数はハンドラーをいつ実行するかを決定するために使用されるパターンであり、二つ目の引数はハンドラー自身です。

パターン名は固定されており、`"/favicon.ico"` のようなルート付きパスや、
`"/images/"` のようなルート付きサブツリー（末尾のスラッシュに注意してください）になります。
長いパターンは短いのよりも優先されるので、
`"/images/"` と `"/images/thumbnails/"` の両方にハンドラーが登録されている場合、
後者のハンドラーは `"/images/thumbnails/"` で始まるパスに対して呼び出され、
前者は `"/images/"` サブツリー内の任意の他のパスに対するリクエストを受け取ります。

上記で定義した `helloHandler` を登録する方法を見てみましょう:

```go
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, web")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
}
```
`main` 関数の一部としてハンドラーを登録していることに注意してください。

上記のコードを実行してみてください:

	$ go run main.go

何が起こりますか？ そう、私たちはパズルの最後のピースを欠いています: webサーバーを起動します！

#### ハンドラーインターフェース

`http.HandleFunc` を使用して `http.HandlerFunc` 型の値を渡すことは、かなり制限することができます。
`http.Handler` インターフェースを満足する任意の値を受け入れる `http.Handle` という別の関数もあります。

[`http.Handler`](https://golang.org/pkg/net/http/#Handler) インターフェースは `http` パッケージで次のように定義されています:

```go
type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
}
```

`http.HandlerFunc` 型は [`HandlerFunc.ServeHTTP`](https://golang.org/pkg/net/http/#HandlerFunc.ServeHTTP) のおかげで `http.Handler` を満たすものだと推測します。

`HandlerFunc` の `ServeHTTP` メソッドのコードは美しいものです。

```go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

このインターフェースは、Webフレームワークとツールキットが機能を追加する拡張ポイントであることがわかります。

### リスニングとサービング

ハンドラーを定義して登録したら、HTTPサーバーを起動してリクエストを待ち受けて、対応するハンドラーを実行する必要があります。

これを行うには [`http.ListenAndServe`](https://golang.org/pkg/net/http/#ListenAndServe) 関数を使用します:

```go
func ListenAndServe(addr string, handler Handler) error
```

最初の引数はサーバーで待ち受けたいアドレスで、
`"127.0.0.1:8080"` や `"localhost:80"` のようなものを使用できます。

二つ目の引数は `http.Handler` で、リクエストを処理する様々な方法を定義するための型です。
`HandleFunc` でデフォルトのメソッドを使用しているので、
ここで任意の値を指定する必要はありません: `nil` になります。

そして最後に、しかし間違いなく関数は `error` を返します。
Goでは、エラーは例外を投げるのではなく値を返すことによって処理されます。

`error` 型は（ `int` や `bool` と同様に）あらかじめ定義された型で、
一つのメソッドしか持たないインターフェースです:

```go
type error interface {
	Error() string
}
```

通常、エラーはメソッドと関数によって返される最後の値であり、
エラーが発生しなかった場合は返された値は `nil` に等しくなります。

したがって、サーバーが正常に起動したことを確認し、エラーをログに記録したい場合は、
`ListenAndServe` を呼び出すようにコードを修正します。

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, web")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

_注釈_: あなたは [main.go](main.go) で完全なコードを見つけることができます。

このコードを実行すると `127.0.0.1:8080` で待ち受けているwebサーバーが起動します。

実行してみてください:

	$ go run main.go

そして http://127.0.0.1:8080/hello にアクセスしてください。

### Bye, webの演習

上記のプログラムを修正し、
httpレスポンスに `"Bye、web"` を出力する `byeHandler` という名前の二つ目のハンドラーを追加してください。

### Hello, Handlerの演習

前の例からのプログラムを修正し、`http.HandleFunc` の呼び出しを `http.Handle` の呼び出しで置き換えられるようにしてください。
`helloHandler` という新しい型を定義し、その型が `http.Handler` インターフェースを満たすようにする必要があります。

### 優秀なマルチプレクサー

すぐに次のようなリクエストをハンドラーにルーティングするより複雑な要件が発生します:

- メソッドに応じたルーティング: `POST` と `GET` は異なるハンドラーをルーティングします。
- パスからの変数抽出: `/product/{productID}/part/{partID}`

これらのケースは手作業または既存の `net/http` パッケージに正しくプラグインされる
[Gorillaツールキット](http://www.gorillatoolkit.org/) とその `mux` パッケージのようなツールキットを使って処理することができます。


```go
package main

import (
	...

	"github.com/gorilla/mux"
)

func listProducts(w http.ResponseWriter, r *http.Request) {
	// すべてのproductをリスト化する
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	// productを追加する
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["productID"]
	// 特定のproductを取得する
}

func main() {
	r := mux.NewRouter()
	// /product/ にあるGETリクエストのみと一致する
	r.HandleFunc("/product/", listProducts).Methods("GET")

	// /product/ にあるPOSTリクエストのみと一致する
	r.HandleFunc("/product/", addProduct).Methods("POST")

	// productIDに関係なくGETに一致する
	r.HandleFunc("/product/{productID}", getProduct)

	// Gorillaルーターですべてのリクエストを処理する。
	http.Handle("/", r)
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal(err)
	}
}
```

Gorillaには、セッション管理、クッキーなどのパッケージも用意されています。
[ドキュメント](http://www.gorillatoolkit.org/) を参照してください。

#### Hello, {you}の演習

前の例から `mux` パッケージを使用して新しいwebサーバーを書いてください。
このサーバーは `"Hello, name"` というテキストを含むHTTPページで `/hello/name` に送られるすべてのHTTPリクエストを処理します。
この例の `name` はもちろん変更可能で、リクエストが `/hello/Francesc` だった場合、レスポンスは `"Hello、Francesc"` となるべきです。

_注釈_：マシンに `mux` パッケージをインストールするには、`go get` を使用できます:

```bash
$ go get github.com/gorilla/mux
```

# おめでとうございます！

Goで初めてのHTTPサーバーを書き上げました！ 素晴らしくないですか？
まぁ、まだそれほど多くはありませんが、最高のものが来るでしょう。

次の章では、HTTPエンドポイントの入力を検証する方法と、
HTTPレスポンスで様々な問題を通知する方法について学習します。

[次のセクション](../section03/README.md) に進みましょう。
