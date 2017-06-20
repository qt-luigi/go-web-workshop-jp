# About this repository（このリポジトリーについて）

The original is [campoy/go-web-workshop](https://github.com/campoy/go-web-workshop).
The author is [Francesc Campoy](https://campoy.cat/).
This repository is a translated each README.md files in [campoy/go-web-workshop](https://github.com/campoy/go-web-workshop) into Japanese.
I got permission from [@francesc](https://twitter.com/francesc) to publish this repos.
At first I was translating it by myself, then using Google Translate by GNMT from the middle and lastly I checked and fixed it by myself again.

オリジナルは [campoy/go-web-workshop](https://github.com/campoy/go-web-workshop) です。
著者は [Francesc Campoy](https://campoy.cat/) 氏です。
このリポジトリーは [campoy/go-web-workshop](https://github.com/campoy/go-web-workshop) の各README.mdファイルを日本語に翻訳したものです。
私は [@francesc](https://twitter.com/francesc) 氏からこのリポジトリーを公開する許可を得ました。
最初は自分で翻訳していたのですが、途中からGNMTによるGoogle翻訳を使用し、最後にもう一度自分でチェックして修正しました。

[![Build Status](https://travis-ci.org/campoy/go-web-workshop.svg)](https://travis-ci.org/campoy/go-web-workshop) [![Go Report Card](https://goreportcard.com/badge/github.com/campoy/go-web-workshop)](https://goreportcard.com/report/github.com/campoy/go-web-workshop)

# GoでWebアプリケーションを構築する

ようこそ、ゴーファー！ あなたはゴーファーではありませんか？
さて、このワークショップはゴーファー、および [Goプログラミング言語][1] を使用する人たちに向けたものです。
しかし前に一度もGoを書いたことがないからと恐れることはありません！
はじめに [Goツアー][2] で言語の基礎を学ぶことをお勧めします。

このワークショップは一流のインストラクターで何回か開催されましたが、
このリポジトリーの目的は個人でコンテンツを追うことをできるだけ簡単にすることです。
もしあなたがどこかの箇所で詰まったなら気軽にファイルのissusに質問してください。

## ワークスペースを設定する

このワークショップを進めるには、次の準備が必要です:

1. [Goプログラミング言語][1] をインストールしている。
1. [How To Write Go Code][9] チュートリアルに従って `GOPATH` を設定している。
1. Goの基礎にある程度精通している。（ [Goツアー][2] はGoをはじめるのにとても良い場所です）
1. Googleアカウントを取得していて、[Google Cloud SDK][3] をインストールしている。

## コンテンツ

Goでも他の言語でも、webアプリケーションを構築する方法については語ることが多くあります。
しかし私たちには1日しかないので、あまり多くをカバーしようとはしていません。
その代わり、私たちは基礎をカバーするので、あなたは後で他のソリューションやフレームワークを探求することができます。

ワークショップは11のセクションに分かれています:

- [0: Hello world](section00/README.md)
- [1: Webクライアント](section01/README.md)
- [2: Webサーバー](section02/README.md)
- [3: 入力の検証とステータスコード](section03/README.md)
- [4: App Engineへデプロイする](section04/README.md)
- [5: Hello, HTML](section05/README.md)
- [6: JSONのエンコードとデコード](section06/README.md)
- [7: Cloud Datastoreで永続化ストレージ](section07/README.md)
- [8: urlfetchでリモートリソースを取得する](section08/README.md)
- [9: Memcacheとは何かとApp Engineから使用する方法](section09/README.md)
- [10: おめでとうございます！](section10/README.md)

## リソース

これらの場所でGoのさらなる情報を見つけることができます:

- [golang.org](https://golang.org)
- [godoc.org](https://godoc.org) 、様々なパッケージのドキュメントを見つけることができます。
- [The Go Programming Language Blog](https://blog.golang.org)

私の好きなGoの一面はコミュニティであり、あなたはすでにその一員です。 ようこそ！

Goコミュニティの新人として、あなたは疑問を抱いたりどこかの箇所で行き詰まるかもしれません。
これは全くもって正常なことであり、私たちはあなたを助けるためにここにいます。
ゴーファーたちがよく出入りする場所のいくつかは次のとおりです:

- [The Go Forum](https://forum.golangbridge.org/)
- [Freenode](https://freenode.net/) の #go-nuts IRCチャンネル
- [Slack](https://gophers.slack.com/messages/general/) のゴーファーたちのコミュニティ (アカウントで [ここ](https://invite.slack.golangbridge.org/) にサインアップ).
- Twitterの [@golang](https://twitter.com/golang) と [#golang](https://twitter.com/search?q=%23golang)
- Google+の [Go+ community](https://plus.google.com/communities/114112804251407510571)
- [Go user meetups](https://go-meetups.appspot.com/)
- golang-nuts [メーリングリスト](https://groups.google.com/forum/?fromgroups#!forum/golang-nuts)
- Go community [Wiki](https://github.com/golang/go/wiki)

### 免責事項

このワークショップはGoogle公認のプロダクト（試験的またはその他）ではなく、Googleが所有しているのはコードだけです。

[1]: https://golang.org
[2]: https://tour.golang.org
[3]: https://cloud.google.com/sdk/downloads
[9]: https://golang.org/doc/code.html#Organization
