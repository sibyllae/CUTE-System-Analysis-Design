# Go API 專案

### 介紹

這是一個使用 Go 語言和 Gin 框架開發的 API 專案。該專案提供以下端點來管理用戶、遊戲、遊戲類型、登入記錄和訂單。

### 安裝與運行

先決條件

- Go 1.16 及以上版本
- MySQL 資料庫

### clone

```bash
git clone https://github.com/yourusername/yourrepository.git
cd yourrepository
```

配置資料庫
請確保你的 MySQL 資料庫已啟動，並在 internal/database/db.go 檔案中配置正確的資料庫連接字串。
運行專案

```bash
go run cmd/app/main.go
```

專案將啟動一個 HTTP 伺服器，監聽 http://127.0.0.1:8080。
