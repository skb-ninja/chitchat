## mux =>マルチプレクサを生成する
   - HandleFunc は第一引数にパス、第二引数にハンドラ関数
   - この場合ルートにアクセスしたらindexにリダイレクトする仕様
   - 全てのハンドラ関数は第一引数にResponseWriter、第二引数にRequestを取るため改めて渡す必要はない。

##　files => 静的なファイルを返送するFileServerを使用
 - lpm