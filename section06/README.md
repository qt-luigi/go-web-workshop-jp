# 6: JSONのエンコードとデコード

Goの標準ライブラリーはJSONのエンコードとデコードをパッケージ [`encoding/json`](https://golang.org/pkg/encoding/json/) で提供します。

## JSONをエンコードおよびデコードする

最初にJSONを一般的な方法でエンコードおよびデコードする方法を学び、
その後、HTTPサーバーの内部でそれを行う方法を見ていきましょう。

### JSONとGo構造体

GoでJSONオブジェクトをエンコードおよびデコードする最も簡単な方法は、
デコードしたいJSONオブジェクトの構造に一致するGoの型を作成することです。

次のようなJSONオブジェクトがあるとします:

```json
{
	"name": "gopher",
	"age_years": 5
}
```

同じフィールドを含む型を作成します:

```go
type Person struct {
	Name      string
	AgeYears  int
}
```

すべての識別子（型とフィールドの両方）は大文字で始まることに注意してください。
これは大文字で始まる識別子だけがパッケージの外部にエクスポートされるためです。
そのため、もし `Name` フィールドが `name` の場合、
`encoding/json` パッケージはそれがそこにあることを知ることさえできません。

幸いにも、私たちはフィールドタグを使用することができ、
GoのフィールドごとにJSONフォームで使用される名前を変更できます。

たとえば、前の例に次のフィールドタグを追加します:

```go
type Person struct {
	Name     string `json:"name"`
	AgeYears int    `json:"age_years"`
}
```

_注釈_: バッククォート ```(`)``` はGoで文字列を書く方法とはまったく異なります。
それらは二重引用符 `(")` を使用して複数の行にまたがることを許しています。

`Person` 型の新しい変数を宣言するために、2つのオプションがあります:

- `var` キーワードを使用してそのフィールドに初期値を与えない、

```go
	var p Person
	fmt.Println(p)
	// output: Person{ 0}

	// 構造体を出力するためのより良い方法
	fmt.Printf("%#v\n", p)
	// output: main.Person{Name:"", AgeYears:0}
```

- または、`:=` 演算子を使用してフィールドを初期化する。

```go
	p := Person{Name: "gopher", AgeYears: 5}
	fmt.Printf("%#v\n", p)
	// output: main.Person{Name:"gopher", AgeYears:5}
```

構造体の詳細については、Goツアーの [このセクション](https://tour.golang.org/moretypes/5) を参照してください。

### Go構造体をJSONにエンコードする

Go構造体をエンコードするために、私たちは便利な `Encode` メソッドを提供する、
[`json.Encoder`](https://golang.org/pkg/encoding/json/#Encoder) を使用します。

```go
package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	p := Person{"gopher", 5}

	// 標準出力に書き込むエンコーダを生成する。
	enc := json.NewEncoder(os.Stdout)
	// pをエンコードするためにエンコーダーを使用する。 失敗する可能性あり。
	err := enc.Encode(p)
	// 失敗した場合は、エラーを記録して実行を停止する。
	if err != nil {
		log.Fatal(err)
	}
}
```

このコードスニペットは、私たちが値をエンコードするたびにエラーを処理する方法を示していますが、
この例では、エンコーダーの出力がネットワーク接続を介して送信される可能性があると考えると、
エラーが発生しないように見えます。

`go run` ツールでコードを試したり、[ここ](https://play.golang.org/p/rsO0Vk-9Xl) のGo Playgroundを使用して試したりすることができます。

### JSONオブジェクトをGo構造体にデコードする

`json.Encoder` と同じように、`json.Decoder` とその使い方はとても似ています。

```go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// 空のPerson値を作成する。
	var p Person

	// 標準入力から読み込むデコーダーを作成する。
	dec := json.NewDecoder(os.Stdin)
	// 値をpにデコードするためにデコーダーを使用する。
	err := dec.Decode(&p)
	// 失敗した場合は、エラーを記録して実行を停止する。
	if err != nil {
		log.Fatal(err)
	}
	// それ以外の場合は、デコードしたものを出力する。
	fmt.Printf("decoded: %#v\n", p)
}
```

`dec.Decode` の引数は `p` ではなく `&p` であることに注意してください。
これは変数 `p` へのポインターであるため、`encoding/json` パッケージは `p` の値を変更することができます。
そうでないなら、`p` のコピーを渡すことになり、変更は副作用なしに行われます。

ポインターの詳細については [Goツアー](https://tour.golang.org/moretypes/1) をご覧ください。

## encoding/json + net/http = webサービス！

`http.HandlerFunc` 型をもう一度見てみましょう:

```go
type HandlerFunc func(ResponseWriter, *Request)
```

### JSONをhttp.ResponseWriterにエンコードする

前述のように `http.ResponseWriter` はメソッド `Write` を実装しているので、
`json.NewEncoder` が要求する `io.Writer `インターフェースを満たします。

そのため、HTTPレスポンスで `Person` を簡単にJSONエンコードできます:

```go
func handler(w http.ResponseWriter, r *http.Request) {
	p := Person{"gopher", 5}

	// Content-Typeヘッダを設定する。
	w.Header().Set("Content-Type", "application/json")

	// pを出力にエンコードする。
	enc := json.NewEncoder(w)
	err := enc.Encode(p)
	if  err != nil {
		// エンコーディングが失敗した場合は、コード500のエラーページを生成する。
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
```

### http.RequestからJSONをデコードする

`http.Request` 型は構造体であり、`io.ReadCloser` 型の `Body` という名前のフィールドと、
`Read` と `Close` のメソッドのインターフェースを持っています。

メソッド `Read` のシグネチャーは `io.Reader` のものと一致するので、
`io.ReadCloser` は `io.Reader` であると言うことができ、したがって、
`http.Request` の `Body` を `json.Decoder` の入力として使うことができます。

```go
func handler(w http.ResponseWriter, r *http.Request) {
	var p Person

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Name is %v and age is %v", p.Name, p.AgeYears)
}
```

このハンドラーをテストしたい場合、curlを使うことができます:

	$ curl -d '{"name": "gopher", "age_years": 5}' http://localhost:8080/
	Name is gopher and age is 5

## 演習

JSONのエンコードとデコードをイベントアプリケーションの[ステップ1](../events/step1/README.md) で追加してください。
その後、ここに戻って来てください！

# おめでとうございます！

バックエンドとフロントエンドがHTTPリクエストを超えてJSONメッセージを通じて対話するwebアプリケーションの構築に成功しました:
これはかなりRESTfulなものです！

しかし、私たちがデコードしている情報のいくつかを保存したいのであればどうでしょうか？

[次の章](../section07/README.md) に進んでGoogle Cloud Datastoreについて学習します。

