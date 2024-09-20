利用パッケージ
JWT : golang-jwt/jwt (https://github.com/golang-jwt/jwt)

中部大学では学科間の購入を深めるために、以下のようなプロジェクトが進行している。その一環として自由に書き込みができる掲示板アプリを作成する。

プロジェクト : 
中部大学で進む学生の学力向上に向けたプロジェクト(PDF)
https://drive.google.com/file/d/14lBC3PIAs16_IP9kmMKwEzsawloZCvrv/view?usp=sharing

# なんのために

- 問題 : 中部大学では、他学部・学科の交流が少なく、多くの学部があるマンモス学校としての強みを出し切れていないと考える。
- 解決に向けて : この問題を解決するために、中部大学の学生専用の掲示板を作成する。ここでは、学部や学科に問わす、興味のある分野や共同で研究したい内容をシェア、仲間を募ることができる。

# 問題解決に向けた独自機能

1. 掲示板はUserが自由に立てられる。その中で、学部・学科を指定可能。該当の学部生はその掲示板が自動的に見れる状態になる。
2. 偏った学部生が集まってしまったときの為に、Helpボタンを用意する。Helpボタンを押した後に学科を指定することでその学科の生徒に協調表示され、アドバイスをもらいやすくできる。
3. アドバイスやサポートに対価を支払うことができる。

## usecase

1. User関連
    1. アカウント新規作成
    2. User情報確認
    3. User情報更新
    4. ログイン
    5. ログアウト
2. UserPost関連
    1. Postsの投稿
    2. 投稿者ごとのPostsの確認
3. Post関連
    1. 他のPostsの検索
    2. 投稿時のタグ付け機能
    3. サポート要請機能
    4. Postの訂正機能 (削除や修正ではない)
4. 掲示板関連
    1. 掲示板作成機能
    2. タグ付け
    3. フィルタリング機能
    4. 通報機能
    5. (通知機能)

# 導入
1. リポジトリをクローンする
2. cdコマンドでGolangDockerのリポジトリに移動する
3. 以下のコマンドでdocker composeを起動する

   ```docker compose up -d --build```

5. docker ps で起動している事を確認する

