// 所有 godoc 註解都要在目標的前一行, 否則不會被解析, 且需要以目標 name 為開頭
// XXX 雖然部分支援 markdown, 但最常用的清單在 JetBrains 的 goland IDE 似乎不解析? 還有一些奇怪的狀況?
// XXX 而且奇怪的是我參考 flag 套件的 godoc, goland IDE 解析可以連結程式碼, 不確定是不是一定要在 github 上才有這功能?

/*
Package godoc 展示 godoc 寫法, XXX 這邊不知道為啥一定要 Package 大寫開頭, 才不會收到 IDE 的警告

# 對於 markdown 支援似乎只有標題?

- 像這個清單不支援?

- 但支援內部程式碼連結 [Sample]
*/
package godoc

// Sample public func 必須以 func name 開頭否則 IDE 會警告
func Sample() {
}

// 123456 private func 由於不會對外公開, 因此不遵循 func name 開頭似乎也不會被 IDE 警告
func sample() {
}
