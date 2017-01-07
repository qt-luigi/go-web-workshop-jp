# ステップ4: 一時的な結果をMemcacheに保存する

前のステップで完成させたアプリケーションでは、誰かがイベントをリスト化するたびに、
各イベントの天気APIに新しいリクエストを送信することによって、
時間とネットワークリソースを無駄にしていました。

今はApp EngineのMemcacheを使用して修正します。
これは完全に管理されたMemcacheインスタンスであり、使い始めるために何もする必要はありません！
それはクールではないですか？

この演習を終えると、アプリケーションが遥かに高速で、リソースが少なくて済むことがわかります。

おめでとうございます、あなたは終わりました！
あなたは [ステップ5](../step5) のものと比較してあなたのコードをチェックすることができます。
最後の指示について [指示](../../section09/README.md#congratulations) に戻ってください。