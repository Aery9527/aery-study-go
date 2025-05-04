# 目錄結構

此目錄結構參考 [project-layout](https://github.com/golang-standards/project-layout), \
除少數 go 本身定義的 folder (**/internal**, **/vendor**) 有特定功能外, \
其餘都是社群約定俗成的結構.

| 目錄                                      | 說明                   | 備註                          |
|-----------------------------------------|----------------------|-----------------------------|
| [/api](./api/README.md)                 | API 定義和協議相關的檔案       |                             |
| [/assets](./assets/README.md)           | 其餘要放入 repository 的檔案 |                             |
| [/build](./build/README.md)             | 持續整合和部署的相關檔案         |                             |
| [/cmd](./cmd/README.md)                 | 主要應用程式的入口點           |                             |
| [/configs](./configs/README.md)         | 程式設定檔                |                             |
| [/deployments](./deployments/README.md) | 部署相關的設定              |                             |
| [/docs](./docs/README.md)               | 專案文件                 |                             |
| [/examples](./examples/README.md)       | 使用範例                 |                             |
| [/init](./init/README.md)               | 系統初始化相關檔案            |                             |
| [/internal](./internal/README.md)       | 私有函式庫程式碼             | **compiler** 會強制封裝這裡面的內容不對外 |
| [/pkg](./pkg/README.md)                 | 公開函式庫程式碼             |                             |
| [/scripts](./scripts/README.md)         | 建置/安裝/分析等腳本          |                             |
| [/test](./test/README.md)               | 整合測試與測試資料            |                             |
| [/tool](./tool/README.md)               | 支援工具                 |                             |
| [/web](./web/README.md)                 | web 靜態資源             |                             |

### `/vendor`

- 這是個在根目錄底下特殊的目錄, 若此目錄存在, 則會優先使用裡面的 module, 否則才會根據 [go.mod](./go.mod) 去網路上抓相依

### `*/testdata`

- 這是個特殊的資料夾名稱, 在任何目錄底下有這個資料夾, 會被 go 工具鏈忽略(例如 `go test`)
- 裡面放測試用的資料, 像這樣讀取內容 `data, err := os.ReadFile("testdata/kerker.json")`
- 這個設計有助於測試資料的集中, 也不會干擾 `go build`
