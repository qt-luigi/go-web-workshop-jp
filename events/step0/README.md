# ステップ0: 基本アーキテクチャー

このステップの目的はwebアプリケーションへの2つのエントリーポイントを定義することです。

コーディングを始める前に実行することで、現在の動作を確認することができます。

    $ goapp serve

これを実装したら、`localhost:8080` にアクセスしてカンファレンスの一覧を表示して、
"新しいイベント" をクリックしても何も起こらないはずです。

それをデプロイすることもできます。

    $ goapp deploy --version=step0 --application=your-project-id

そうしたら https://step0.your-project-id.appspot.com にアクセスしてください。

作業が完了したら、[指示](../../section05/README.md#congratulations) に戻りましょう。
