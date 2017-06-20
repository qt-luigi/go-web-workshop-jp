# 3: 入力の検証とステータスコード

これまでは、サーバーが受け取るすべてのHTTPリクエストは正しいと仮定しました。
そしてWebサーバーを書くときに決してしてはいけないことがあるとすれば、それはあなたの入力を信頼していることです！

この章では、HTTPリクエストの様々な部分で送信された情報を抽出する方法を確認します。
その情報が得られれば、それらをどのように検証できるのか、様々なエラーをどのように伝えることができるのかがわかります。

それでは始めましょう！

## URLからパラメーターを読み込む

パスに応じてリクエストを別のハンドラーにルーティングする方法を見てきました。
今度は `?` の後のデータとして知られているリクエストのクエリー部分の情報を抽出する方法を見てみましょう。

`http.Request` 型は次のドキュメントにある `FormValue` メソッドを持っています:

    func (r *Request) FormValue(key string) string

    FormValueは、クエリーの名前付きコンポーネントの最初の値を返します。 POSTとPUTのボディのパラメーターはURLクエリー文字列の値よりも優先されます。 FormValueは必要に応じてParseMultipartFormとParseFormを呼び出し、これらの関数から返された任意のエラーを無視します。 キーが存在しない場合、FormValueは空の文字列を返します。 同じキーの複数の値にアクセスするには、ParseFormを呼び出してRequest.Formを直接検査します。

それは簡単です！ そのため `/hello?msg=world` というURLでパラメーター `msg` の値を取得したい場合、次のプログラムを書くことができます。

[embedmd]:# (examples/handlers/main.go /func paramHandler/ /^}/)
```go
func paramHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		name = "friend"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
```

### Hello, parameterの演習

`/hello?name=world` へのリクエストに `Hello、world!` というテキストのHTTPレスポンスで答えるwebサーバーを書いてください。
`name` が存在しない場合、`Hello、friend!` を表示する必要があります。

あなたは自分のブラウザーでテストすることができますが、`curl` でもいくつか試してみましょう。
これらを実行する前に、あなたが確認したいこととその理由を考えてください。

```bash
$ curl "localhost:8080/hello?name=world"

$ curl "localhost:8080/hello?name=world&name=francesc"

$ curl -X POST -d "name=francesc" "localhost:8080/hello"

$ curl -X POST -d "name=francesc" "localhost:8080/hello?name=potato"
```

あなたのプログラムが `name` に与えられたすべての値をどのように出力させるのかを考えてください。

## リクエストボディから読み込む

二つ前の章で `http.Response` の `Body` から読み込んだのと同様に、`http.Request` のbodyを読み込むことができます。

`http.Request` の `Body` の型が `io.ReadCloser` であっても、こちらのbodyはhttpハンドラーの実行終了時に自動的に閉じられるので、心配いりません。

`io.Reader` から読み込むことができる方法はたくさんありますが、とりあえず `ioutil.ReadAll` を使うことができ、何か問題が起これば `[]byte` と `error` を返します。

[embedmd]:# (examples/handlers/main.go /func bodyHandler/ /^}/)
```go
func bodyHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "could not read body: %v", err)
		return
	}
	name := string(b)
	if name == "" {
		name = "friend"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
```

### Hello, bodyの演習

前の演習を修正し、`name` 引数をクエリーやフォームの値から読み込む代わりに、`Body` の内容を使用してください。
同じように、`Body` が空であればレスポンスは `friend` と挨拶すべきです。

`ioutil.ReadAll` の呼び出しがエラーを返す場合、そのエラーメッセージを出力に書き出してください。

`curl` を使用してあなたの演習をテストできます:

```bash
$ curl -X POST -d "francesc" "localhost:8080/hello"
```

さらなる演習として、nameの周りの余分な空白スペースを削除してください。

## エラーを伝える

前の演習では `ioutil.ReadAll` の呼び出しが失敗した場合にエラーを表示することにしました。
それはあなたが想像しているように、実際にはかなり恐ろしい考えです 😁。

私たちはどうしたらよいですか？ HTTPプロトコルではレスポンスの性質を記述するのに役立つ一連のステータスコードが定義されています。
実際には `Get` を使用して取得したレスポンスが `OK` であるかどうかを確認してから、使用しています。

`ResponseWriter` でステータスコードを設定する方法は2つあります。

### ResponseWriter.WriteHeaderでのステータスコード

`ResponseWriter` で `WriteHeader` メソッドを使用すると、レスポンスのステータスコードを設定できます。
引数は `int` なので、任意の数値を渡すことができますが、`http` パッケージですでに定義されている定数を使う方が良いでしょう。
それらはすべて `Status` で始まり、[ここ](https://golang.org/pkg/net/http/#pkg-constants) で見つけることができます。

デフォルトでは、レスポンスのステータスコードは `200` で知られている `StatusOK` です。

#### より良いエラーの演習

前のプログラムを修正し、エラーの場合のレスポンスはステータスコードを `500` にしてください。
数値の `500` を使用する代わりに、対応する定数を見つけてください。

次にbodyが空の場合はレスポンスのステータスを `400` にしてください。

### http.Errorでのステータスコード

前の演習では、レスポンスのステータスコードを設定するときに、エラーの説明も書き込むことがよくあることに気付きました。
そのために `http.Error` 関数が存在します。

    func Error(w ResponseWriter, error string, code int)

    Errorは、指定されたエラーメッセージとHTTPコードでリクエストに応えます。 エラーメッセージはプレーンテキストである必要があります。

`Error` への呼び出しは `WriteHeader` への呼び出しに置き換えられ、その後に `ResponseWriter`への書き込みがいくつか呼び出されます。

#### http.Errorでのステータスコードの演習

前の演習での `WriteHeader` と `Fprintf` の呼び出しを `Error` の呼び出しに置き換えてください。

### レスポンスヘッダー

あなたはレスポンスに複数の行を送ると、あなたのブラウザーはそれを1つの行として表示することに気づいたかもしれません。何故でしょうか？

答えは `net/http` パッケージが出力はHTMLだと推測して、ブラウザーが行を連結するということです。
短い出力に対しては、推測が困難です。
`curl` コマンドに `-v` を追加することで、あなたのレスポンスのコンテンツタイプを見ることができます。

```bash
$ curl -v localhost:8080
< HTTP/1.1 200 OK
< Date: Mon, 25 Apr 2016 16:14:46 GMT
< Content-Length: 19
< Content-Type: text/html; charset=utf-8
```

それでは、`net/http` パッケージがコンテンツタイプを推測することを止めるにはどうすればよいでしょうか？ 私たちはそれを指定します！
これを行うには、ヘッダーの `"Content-Type"` を `"text/plain"` 値に設定する必要があります。
`ResponseWriter` の `Header` 関数でヘッダーをレスポンスに設定できます。

`Header` は、他のメソッドの中でも、`Set` メソッドを持つ [`http.Header`](https://golang.org/pkg/net/http/#Header) を返します。
このように `w` という名前の `ResponseWriter` でコンテンツタイプを設定することができます。

[embedmd]:# (examples/texthandler/main.go /w.Header.*/)
```go
w.Header().Set("Content-Type", "text/plain")
```

### 反復を避ける

あなたは何百もの異なるHTTPハンドラーを持っていて、それぞれにコンテンツタイプヘッダーを設定したいと考えているとします。
それは苦しいように聞こえませんか？

多くのハンドラーが共有する振る舞いを定義するのに役立つクールなテクニックについて教えてください。
いくつかの人々はそれらをデコレーターと呼び、そのほとんどはPythonも書いています 😛。

始めに `http.HandlerFunc` を含む `textHandler` という名前の新しい型を定義します。

[embedmd]:# (examples/texthandler/main.go /type textHandler/ /^}/)
```go
type textHandler struct {
	h http.HandlerFunc
}
```

ここで `textHandler` に `ServeHTTP` メソッドを定義し、`http.Handler` インターフェースを満たします。

[embedmd]:# (examples/texthandler/main.go /.*ServeHTTP/ /^}/)
```go
func (t textHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// コンテンツタイプを設定する
	w.Header().Set("Content-Type", "text/plain")
	// デコレートされたハンドラーでServeHTTPを呼び出す。
	t.h(w, r)
}
```

最後に `http.HandleFunc` での呼び出しを `http.Handle` に置き換えます。

[embedmd]:# (examples/texthandler/main.go /func main/ /^}/)
```go
func main() {
	http.Handle("/hello", textHandler{helloHandler})
	http.ListenAndServe(":8080", nil)
}
```

#### ヘッダーを一回だけ設定する演習

前の演習からのプログラムを修正し、レスポンスのコンテンツタイプを設定する行が1つしかないようにしてください。

#### 優れたエラー処理の演習（任意）

前に示した `textHandler` を `http.HandlerFunc` の代わりに `int` と `error` を返す関数を受け取るように修正してください。
`textHandler` の `ServeHTTP` メソッドは、その整数と `error` をチェックし、それに応じてステータスコードと内容を設定する必要があります。

_超難問_: ステータスコードに関する情報も含む新しいエラー型を定義してください。
これを成し遂げる最初の人は賞金を得るかもしれない ... ちょっと言い過ぎです。

# おめでとうございます！

これで、httpハンドラーへの入力を検証し、それに応じてステータスコードとコンテンツタイプを設定できるようになりました。
あなたははっきりと素晴らしいです！ 🎉

あなたが素晴らしい人になれることを、[セクション4](../section04/README.md) に進むことによって、保証します。
