package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Trainer 是一個範例結構，用於儲存到 MongoDB
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// --- 1. 設定並連線到 MongoDB ---
	// 設定客戶端選項，指定 MongoDB 連線 URI
	// 這裡我們使用本地開發環境的預設位址
	clientOptions := options.Client().ApplyURI("mongodb://root:123456@localhost:27017")

	// 建立一個帶有超時的 context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 使用設定好的選項和 context 來連線到 MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 在函式結束時，確保斷開連線
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// 檢查連線是否成功
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功連線到 MongoDB!")

	// --- 2. 寫入資料 ---
	// 選擇要操作的資料庫和集合 (collection)
	// 如果它們不存在，MongoDB 會在第一次寫入時自動建立
	collection := client.Database("game").Collection("kerker")

	// 建立一個要寫入的資料實體
	ash := Trainer{Name: "Aery", Age: 18, City: "BigBig"}

	// 使用 InsertOne 將單一文件寫入集合中
	insertResult, err := collection.InsertOne(ctx, ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功寫入一筆資料，ID:", insertResult.InsertedID)

	// --- 3. 讀取資料 ---
	// 我們將使用剛剛寫入的 ID 來查詢該筆資料
	var result Trainer
	// 建立一個 filter 來尋找 _id 符合我們插入的 ID 的文件
	filter := bson.M{"_id": insertResult.InsertedID}

	// 使用 FindOne 查詢單一文件
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		// 如果沒有找到文件，會回傳 mongo.ErrNoDocuments
		if err == mongo.ErrNoDocuments {
			fmt.Println("找不到對應的資料")
			return
		}
		log.Fatal(err)
	}

	fmt.Printf("成功讀取到資料: %+v\n", result)
}
