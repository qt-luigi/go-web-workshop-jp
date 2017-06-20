# ステップ0: 基本アーキテクチャー

このステップの目的はwebアプリケーションへの2つのエントリーポイントを定義することです。

コーディングを始める前に実行することで、現在の動作を確認することができます。

    $ dev_appserver.py .

これを実装したら、`localhost:8080` にアクセスしてカンファレンスの一覧を表示して、
"新しいイベント" をクリックしても何も起こらないはずです。

デプロイするには、まず [gcloud](https://cloud.google.com/sdk/downloads) をインストールして設定する必要があります:

    $ gcloud init

_注意_: Google Compute Engineのゾーンを設定する必要はありません。

その後、次を実行します

    $ gcloud app deploy --version=step0 app.yaml

そうしたら https://step0.your-project-id.appspot.com にアクセスするか次を実行します

    $ gcloud app browse --version=step0

This will display a basic Events page but also an alert will be alerted. That's fine, for now.
これにより基本となるイベントページが表示されますがアラートにも警告が表示されます。 それは今のところ大丈夫です。

作業が完了したら、[指示](../../section05/README.md#congratulations) に戻りましょう。
