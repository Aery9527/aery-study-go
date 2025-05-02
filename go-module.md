
- **go.mod** 定義用了哪些 module/版本是多少(類似 maven pom.xml), **go.sum** 是記錄這些 module 的內容 checksum 確保下載來的沒被改動
- Go 採用 MVS (Minimal Version Selection) 的版本解決策略, 例如 A 相依 C:v1.1, B 相依 C:v1.2, 那麼整體會使用 C:v1.2, 因為不支援多版本共存
- 為避免上述的 A 使用到 C:v1.2 而炸掉, 所以 Go 社群推崇 semver (Semantic Versioning, 語意化版本), 也就是說小版號不應該有 breaking change,
  而是不向下相容時跳大版號, 因此不同版號而炸掉是開發者的問題!