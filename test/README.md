# 整合測試與測試資料

- 主要是整合測試與測試資料
- 這邊也可以寫 `main.go` 直接執行或 `xxx_test.go` 跑 `go test` 執行
- 與 `xxx_test.go` 的差異對照表

  | | `xxx_test.go`                         | `/test`                     |
  |---|---------------------------------------|-----------------------------|
  | 定義方式 | `*_test.go` 檔案 + Go testing framework | 自訂測試工具、腳本、測資                |
  | 類型 | 單元測試(Unit Test)                       | 整合測試（Integration）、E2E、壓力測試  |
  | 執行方式 | `go test` 自動執行                       | 自己寫 script 或 CI/CD 執行  |
  | 存取資料 | 用 fixture / mock                      | 用 /test/data 中的實體檔案  |
  | 使用場景 | 測某一個 function 是否正確                      | 模擬真實環境，測一整條流程  |

