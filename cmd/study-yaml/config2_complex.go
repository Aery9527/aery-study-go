package main

// ComplexConfig 複雜結構體配置 (對應 config2.yaml)
type ComplexConfig struct {
	User struct {
		ID      int `yaml:"id"`
		Profile struct {
			Name    string `yaml:"name"`
			Email   string `yaml:"email"`
			Age     int    `yaml:"age"`
			Address struct {
				Country string `yaml:"country"`
				City    string `yaml:"city"`
				Street  string `yaml:"street"`
				Zipcode string `yaml:"zipcode"`
			} `yaml:"address"`
		} `yaml:"profile"`
		Preferences struct {
			Language      string `yaml:"language"`
			Timezone      string `yaml:"timezone"`
			Notifications struct {
				Email bool `yaml:"email"`
				SMS   bool `yaml:"sms"`
				Push  bool `yaml:"push"`
			} `yaml:"notifications"`
		} `yaml:"preferences"`
	} `yaml:"user"`
	Environment map[string]map[string]string `yaml:"environment"`
	Metadata    struct {
		CreatedAt string `yaml:"created_at"`
		Tags      []struct {
			Type  string `yaml:"type"`
			Value string `yaml:"value"`
		} `yaml:"tags"`
		CustomFields map[string]interface{} `yaml:"custom_fields"`
	} `yaml:"metadata"`
}

// 載入複雜配置檔案
func loadComplexConfig(filename string) ComplexConfig {
	var config ComplexConfig
	data, err := readConfigFile(filename)
	if err != nil {
		logFatal("讀取檔案 %s 失敗: %v", filename, err)
	}

	err = unmarshalYAML(data, &config)
	if err != nil {
		logFatal("解析 YAML 檔案 %s 失敗: %v", filename, err)
	}

	return config
}
