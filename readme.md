# Hakka Story

一個基於 Go 開發的客語故事系統後端，包括 Admin 和 API 兩個部分。

[API Doc](http://174.138.28.65:8080/swagger/index.html)

## 功能

主要功能：
- [X] 故事的 CRUD，包括標題、敘述、每一頁的中文和客語翻譯內容
- [ ] 故事的分類
- [ ] 故事封面圖片的上傳
- [ ] 故事內容的語音檔上傳
- [ ] AI 產生客語翻譯
- [ ] AI 產生故事封面
- [ ] AI 產生故事語音
- [ ] 故事分類

附加功能：
- [ ] 基於故事內容的客語練習題

## Auth

目前 Auth 使用 Auth0 進行管理，在 Auth0 設定 Access Token 帶入：
1. email
2. user_roles

## 資料庫

使用 MySQL 做為資料庫，並使用 goose 做為 migration 工具。

UP
```shell
goose -dir migrations mysql "${user}:${pwd}@${host}:${port}/${db} up
```

DOWN
```shell
goose -dir migrations mysql "${user}:${pwd}@${host}:${port}/${db} down
```

## Deploy

使用 CI Deploy 到 DigitalOcean，實在沒有時間所以不用 Kubernetes。
