# package 概念

- 概念就像 java 一個 "class" 的 scope, 也就是說散在各檔案的東西只要是同個 package 就是同個 scope, 也可看作其他語言的 namespace 概念.
- package 與所在的 folder name 無相關性(但習慣上會保持一致), 然後在同個 folder 裡的檔案會強制視為同個 scope, 所以它們的 package 全都要一樣
    - 命名習慣: 小寫字母/沒有底線/沒有駝峰, 也就是說應該要全是英文小寫.
        - 這樣的習慣是為了"命名明確", 就可以避免過多含意的 package 名稱
        - 也避免使用像這種 util/controller/manager 抽象過高卻有可能包含太多領域的爛大街名稱
        - 但 [/cmd](../) 裡的 `main()` 其 folder name 則比較沒有此限定, 原因是它只是個入口, 並非主要程式邏輯內容
        - struct 盡量不要包含 package 語意
        ```
        package user // package user 裡不要定義 type User 這種重複命名
        type User struct {} // user.User 很饒口
        type Entity struct {} // user.Entity 好懂多了
        ```

由於上述的概念, 所以當多維度交錯時則應該以 **領域** 為主劃分 package, \
例如 user/order (領域) 跟 service/repository (角色) 交錯時, \
應以 user/order (領域) 為 package 劃分

```
internal/
├── order/
│   ├── service.go        // package order
│   ├── repository.go
│   └── model.go
├── user/
│   ├── service.go        // package user
│   ├── repository.go
│   └── model.go
```

- 這樣就以 **領域**(user/order) 劃分所有面向(SRP), 業務邏輯就可以收斂.
- **領域** **角色** 可以簡單用 **業務需求** 或 **系統需求** 來區分:
  - user/order: 是業務邏輯劃分出來的 **領域** 概念
  - service/repository: 是程式系統操作上劃分出來的 **角色** 概念.

