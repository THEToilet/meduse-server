ping表示

## 仕様

- サーバが受け取るもの

|送信元| 通信要求番号 | 16進数 | 2進数 | 10進数 |
|:---: | :---: | :---: | :---: | :---: |
| client&host_client | ユーザ登録リクエスト | 0x01 |
| client&host_client | ユーザ一覧取得リクエスト（定期的） | 0x02 |
| client | パケット中継 | 0x03 | 2byte
| client&host_client | アプリ離脱 | 0x04 |
| host_client | コントローラユーザを観客席に送る | 0x05
| host_client | ユーザをBANする | 0x06
| host_client | ユーザをコントローラに登録する | 0x07
| client&host_client | pong | 0x08 |

タイプ + ユーザ番号 + ボタン 0x04 + 0x01 + 0x0070

| ボタン | 16進数 | 2進数 | 10進数 |
| :---: | :---: | :---: | :---: |
| 上 |  0x01 | 0B00000001 | 1 |
| 下 | 0x02 | 0B00000010 | 2 |
| 右 | 0x04 | 0B00000100 | 4 |
| 左 | 0x08 | 0B00001000 | 8 |
|Aボタン | 0x10 | 0B00010000 | 16 |
|Bボタン | 0x20 | 0B00100000 | 32 |
|1ボタン | 0x40 | 0B01000000  | 64 |
|2ボタン | 0x80 | 0B10000000 | 128 |
|+ボタン | 0x0100 | 0B0000000100000000 | 256 |
|-ボタン | 0x0200 | 0B0000001000000000 | 512 |
|homeボタン | 0x0400 | 0B0000010000000000 | 1024 |
|シェイク | 0x0800 | 0B0000100000000000 | 2048 |

[ A,B,1 ] : 0x0070 0B0000000001110000

|+コントローラー番号 +コントローラータイプ | host接続断

| ユーザ番号 | 16進数 | 2進数 | 10進数 |
| :---: | :---: | :---: | :---: |
|  1 | 0x01 | 0B00000001 |
| 2 | 0x02 | 0B00000010 |
| 3 | 0x03  | 0B00000011 |
| 4 | 0x04 | 0B00000100 |

## 機能

クライアント接続断

- クライアントとサーバ間ではping-pongを行う

ユーザマッチング機能 パケット中継

- 部屋は4部屋で
- 部屋はホスト1人と参加者3人で構成される
- 参加者は1人以上3人以下

- 8人ぐらいのスペース
- 途中でプレイヤーの番号変えられる
- プレイヤーの名前追加
- コントローラータイプ
- 最低2人
- ゲーム開始後の途中参加OK

- COM1~4存在
- あと4人ぐらい予備メンバがいる
- その中から補充できる
- 参加者は8人
- プレイ中の1~4人以外は待機、ただ疎通確認はしておく
- クライアントホスト側で7人の操作を行う
- 定期的にプレイヤー情報も更新させる5sに一回ぐらい
- set com4 to player7
- ユーザをコントローラに登録する
- 上書き
- ban user
- ユーザを完全にBAN
- back to queue
- COMから外して予備地に送る

### Client -> Server

- サーバに登録する（名前とともに）
- コントローラ情報を送信する（自分が操作側の場合）
- プレイヤー情報の更新(5s)
- 部屋を抜ける

### Server -> Client

- 疎通確認

### Server -> HostClient

- コントローラ情報を送る
- 疎通確認

### HostClient -> Server

- サーバに登録
- コントローラユーザを観客席に送る
- ユーザをBANする
- ユーザをコントローラに登録する
- プレイヤー情報の更新(5s)
- 部屋を抜ける

## 通信タイプ

通信タイプ 1byte UUID 16byte

