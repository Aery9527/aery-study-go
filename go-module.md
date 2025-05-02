# Go Module

- 這基本上就是 java 的 maven 或 gradle 的相依管理,  主要是以 repository 為單位
- 1.11 引入的官方套件管理方式, 1.16起預設使用, 完全取代了早期的 `GOPATH`
- [go.mod](./go.mod) 記錄專案的 module 名稱, 依賴關係, 版本等資訊
- [go.sum](./go.sum) 是 module 的 checksum 檔案, 記錄每個版本的雜湊值, 防止供應鏈攻擊/確保一致性(不要手賤改它, 讓 go 維護)
- go module 的設計初衷是 **與版本控制系統整合(主要是 Git)**, 所以它乾脆直接把 module path 設計成 **來源識別子(source identifier)**
  - 用 URL 的形式(通常是 Git repo URL), 可以自然保證全世界 module path 不會撞名
  - go module system 會用 go get 自動去抓對應的 git 倉庫
  - 搭配 Semantic Import Versioning 流暢, 例如 **v2+** 的 major version 可以寫這樣 `import "github.com/foo/bar/v2"`
  - 支援多種 VCS (不只 GitHub): git, mercurial, svn, ...
  - 可以設置 proxy, 例如 `GOPROXY=https://proxy.golang.org` 來加速下載, 這個 proxy 會快取 module, 並且提供 checksum 檔案
  - 在開發時想用本地 module, 就使用 `replace` 指令, 例如 `replace example.com/foo => ../foo`
- go 採用 MVS (Minimal Version Selection) 的版本解決策略, 簡單來說就是最高版本. 例如 A 相依 C(v1.1), B 相依 C(v1.2), 那麼整體會使用 C(v1.2)
- 為避免上述的 A 使用到 C(v1.2) 而炸掉, 所以 Go 社群推崇 semver (Semantic Versioning, 語意化版本), 也就是說小版號不應該有 breaking change. \
  所以當修改內容不向下相容時, 就需要跳大版號, 因此不同版號而炸掉是開發者的問題!