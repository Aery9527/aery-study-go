package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var configPrefix = "cmd/study-yaml/"

// 共用的輔助函數
func readConfigFile(filename string) ([]byte, error) {
	return os.ReadFile(configPrefix + filename)
}

func unmarshalYAML(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func marshalYAML(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func logFatal(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func main() {
	fmt.Println("=== yaml.v3 完整使用範例 ===\n")

	// 範例 1: 讀取基本類型配置
	fmt.Println("1. 讀取基本類型配置 (config1.yaml)")
	basicConfig := loadBasicConfig("config1.yaml")
	fmt.Printf("應用程式名稱: %s\n", basicConfig.Basic.Name)
	fmt.Printf("版本: %s\n", basicConfig.Basic.Version)
	fmt.Printf("除錯模式: %t\n", basicConfig.Basic.Debug)
	fmt.Printf("連接埠: %d\n", basicConfig.Basic.Port)
	fmt.Printf("逾時: %.1f 秒\n", basicConfig.Basic.Timeout)

	fmt.Println("伺服器列表:")
	for i, server := range basicConfig.Servers {
		fmt.Printf("  伺服器 %d: %s:%d\n", i+1, server.Host, server.Port)
	}

	fmt.Printf("標籤: %v\n", basicConfig.Tags)
	fmt.Printf("連接埠列表: %v\n", basicConfig.Ports)
	fmt.Println()

	// 範例 2: 讀取複雜結構體配置
	fmt.Println("2. 讀取複雜結構體配置 (config2.yaml)")
	complexConfig := loadComplexConfig("config2.yaml")
	fmt.Printf("使用者 ID: %d\n", complexConfig.User.ID)
	fmt.Printf("使用者姓名: %s\n", complexConfig.User.Profile.Name)
	fmt.Printf("電子郵件: %s\n", complexConfig.User.Profile.Email)
	fmt.Printf("年齡: %d\n", complexConfig.User.Profile.Age)
	fmt.Printf("地址: %s %s %s %s\n",
		complexConfig.User.Profile.Address.Country,
		complexConfig.User.Profile.Address.City,
		complexConfig.User.Profile.Address.Street,
		complexConfig.User.Profile.Address.Zipcode)

	fmt.Println("環境設定:")
	for env, config := range complexConfig.Environment {
		fmt.Printf("  %s 環境:\n", env)
		for key, value := range config {
			fmt.Printf("    %s: %s\n", key, value)
		}
	}

	fmt.Println("中繼資料標籤:")
	for _, tag := range complexConfig.Metadata.Tags {
		fmt.Printf("  %s: %s\n", tag.Type, tag.Value)
	}

	fmt.Println("自訂欄位:")
	for key, value := range complexConfig.Metadata.CustomFields {
		fmt.Printf("  %s: %v (類型: %T)\n", key, value, value)
	}
	fmt.Println()

	// 範例 3: 讀取高級配置
	fmt.Println("3. 讀取高級配置 (config3.yaml)")
	advancedConfig := loadAdvancedConfig("config3.yaml")
	fmt.Printf("主資料庫: %s:%d\n",
		advancedConfig.Database.Connections.Primary.Host,
		advancedConfig.Database.Connections.Primary.Port)
	fmt.Printf("副本資料庫: %s:%d\n",
		advancedConfig.Database.Connections.Replica.Host,
		advancedConfig.Database.Connections.Replica.Port)

	fmt.Printf("連線逾時: %v\n", advancedConfig.Timeouts.Connect)
	fmt.Printf("讀取逾時: %v\n", advancedConfig.Timeouts.Read)
	fmt.Printf("寫入逾時: %v\n", advancedConfig.Timeouts.Write)
	fmt.Printf("閒置逾時: %v\n", advancedConfig.Timeouts.Idle)

	fmt.Printf("新 UI 功能: %t\n", advancedConfig.OptionalConfig.FeatureFlags.NewUI)
	if advancedConfig.OptionalConfig.FeatureFlags.BetaFeatures != nil {
		fmt.Printf("Beta 功能: %t\n", *advancedConfig.OptionalConfig.FeatureFlags.BetaFeatures)
	} else {
		fmt.Println("Beta 功能: nil (未設定)")
	}

	fmt.Println("服務配置:")
	for _, service := range advancedConfig.Services {
		fmt.Printf("  服務類型: %s, 名稱: %s\n", service.Type, service.Name)
		if service.Type == "http" {
			if httpConfig := parseHTTPConfig(service.Config); httpConfig != nil {
				fmt.Printf("    HTTP 設定: %s:%d (SSL: %t)\n",
					httpConfig.Host, httpConfig.Port, httpConfig.SSL)
			}
		} else if service.Type == "grpc" {
			if grpcConfig := parseGRPCConfig(service.Config); grpcConfig != nil {
				fmt.Printf("    gRPC 設定: %s:%d\n", grpcConfig.Host, grpcConfig.Port)
				fmt.Printf("    TLS 證書: %s\n", grpcConfig.TLSCert)
			}
		}
	}
	fmt.Println()

	// 範例 4: 動態創建和寫入 YAML
	fmt.Println("4. 動態創建和寫入 YAML 範例")
	demonstrateYAMLCreation()
	fmt.Println()

	// 範例 5: 錯誤處理和驗證
	fmt.Println("5. 錯誤處理和驗證範例")
	demonstrateErrorHandling()
}

// 示範動態創建和寫入 YAML
func demonstrateYAMLCreation() {
	// 創建一個示例配置
	config := map[string]interface{}{
		"application": map[string]interface{}{
			"name":    "動態創建的應用程式",
			"version": "2.0.0",
			"author":  "AI 助手",
		},
		"database": map[string]interface{}{
			"host":     "localhost",
			"port":     5432,
			"username": "admin",
			"ssl":      true,
		},
		"features": []string{"authentication", "logging", "monitoring"},
		"limits": map[string]interface{}{
			"max_users":       1000,
			"request_timeout": "30s",
			"rate_limit":      100,
		},
	}

	// 將配置序列化為 YAML
	yamlData, err := marshalYAML(config)
	if err != nil {
		log.Printf("序列化 YAML 失敗: %v", err)
		return
	}

	fmt.Println("動態創建的 YAML 內容:")
	fmt.Println(string(yamlData))

	// 寫入到檔案
	err = os.WriteFile("dynamic_config.yaml", yamlData, 0644)
	if err != nil {
		log.Printf("寫入檔案失敗: %v", err)
		return
	}
	fmt.Println("已成功寫入 dynamic_config.yaml 檔案")
}

// 示範錯誤處理和驗證
func demonstrateErrorHandling() {
	// 嘗試讀取不存在的檔案
	_, err := readConfigFile("nonexistent.yaml")
	if err != nil {
		fmt.Printf("預期的檔案不存在錯誤: %v\n", err)
	}

	// 嘗試解析無效的 YAML
	invalidYAML := `
invalid yaml content:
  - this is: [invalid
    structure
`
	var testConfig map[string]interface{}
	err = unmarshalYAML([]byte(invalidYAML), &testConfig)
	if err != nil {
		fmt.Printf("預期的 YAML 解析錯誤: %v\n", err)
	}

	// 示範結構體驗證
	validationExample := `
user:
  name: ""  # 空字串
  age: -5   # 負數
  email: "invalid-email"  # 無效格式
`

	type User struct {
		Name  string `yaml:"name"`
		Age   int    `yaml:"age"`
		Email string `yaml:"email"`
	}

	type ValidationConfig struct {
		User User `yaml:"user"`
	}

	var validationConfig ValidationConfig
	err = unmarshalYAML([]byte(validationExample), &validationConfig)
	if err != nil {
		fmt.Printf("YAML 解析錯誤: %v\n", err)
	} else {
		// 自訂驗證邏輯
		if validationConfig.User.Name == "" {
			fmt.Println("驗證錯誤: 使用者姓名不能為空")
		}
		if validationConfig.User.Age < 0 {
			fmt.Printf("驗證錯誤: 年齡不能為負數，得到: %d\n", validationConfig.User.Age)
		}
		// 這裡可以加入更多驗證邏輯，例如 email 格式驗證等
	}
}
