# 8: urlfetchでリモートリソースを取得する

時には、アプリケーションが外の世界とやりとりしたり、POSTリクエストを通じてデータを送信したり、
GETを使用して情報を取得したりすることが必要になります。

App Engineのフレームワークはスケーラビリティとパフォーマンスを保証するためにできることを制限していて、
これは `net/http` パッケージを直接使用して `http.Get("https://google.com")` を実行できないことを意味していますが、
App Engineランタイムによって提供される [`appengine/urlfetch`](https://cloud.google.com/appengine/docs/go/urlfetch/reference) パッケージを使用することで簡単に行えます。

`urlfetch` パッケージの最も重要な関数は `Client` です:

```go
func Client(context appengine.Context) *http.Client
```

`appengine.Context` を指定すると、HTTPクライアントが取得され、そこから開始することができます。
Googleのホームページを取得したい場合は、次のようにします:

```go
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// 新しいHTTPクライアントを生成する
	c := urlfetch.Client(ctx)

	// Googleホームページのリクエストに使用する
	res, err := c.Get("https://google.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	// この関数の最後でbodyを閉じる必要がある
	defer res.Body.Close()

	// 出力先に全てのwebページを出力できる
	_, err = io.Copy(w, res.Body)
	if err != nil {
		log.Errorf(ctx, "could not copy the response: %v", err)
	}
}
```

この [ディレクトリー](fetch) で見つかるこのコードを実行すると、
すべてのローカルリンクが壊されているため、Googleホームページの版が少し壊れているはずです。

# 演習: イベントの場所の天気を取得する

前に学習したJSONのエンコードとデコードの知識を使って、
[ステップ3](../events/step3/README.md) に取り組む時です。

# おめてとうございます！

あなたは `urlfetch` パッケージのおかげで、
App EngineアプリケーションをHTTPリクエストを使用して他のwebとやりとりできるようにしました！

しかしリクエストを処理するたびに外部リソースを取得することは本当に良い考えですか？
後で使用するためにこのデータの一部をキャッシュしてはいけませんか？

memcacheを使用してApp Engineから情報をキャッシュして取得する方法を学習するため [次の章](../section09/README.md) に進みます。
