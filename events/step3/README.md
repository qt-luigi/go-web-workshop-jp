# ステップ3: openweathermap.orgで天気を追加する

これまでは、webフォームから送信されたイベントを格納し、
webページに戻して表示することができるアプリケーションを用意していました。

イベントが行われている場所の天気を取得することでクールな機能を追加します。
これを行うには、外部APIから情報を取得する必要があり、
オープンなopenweathermap.orgを使用します。

APIキーを取得して app.yaml ファイルの WEATHER_API_KEY の値を置き換えるには、
https://openweathermap.org にサインアップする必要があります。

次にAPIへの呼び出しを実装し、リクエストを用意し、レスポンスをデコードし、
イベントに結果として生じる天気を追加します。

これはリクエストごとでイベントごとに1つのAPI呼び出しが発生するので、非常に無駄です。
しかし私たちはそれを後で修正するので心配しないでください。
今は [指示](../../section08/README.md#congratulations) に戻ってください。
