package main

// BasicConfig 基本類型配置 (對應 config1.yaml)
type BasicConfig struct {
	Basic struct {
		Name    string  `yaml:"name"`
		Version string  `yaml:"version"`
		Debug   bool    `yaml:"debug"`
		Port    int     `yaml:"port"`
		Timeout float64 `yaml:"timeout"`
	} `yaml:"basic"`
	Servers []struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"servers"`
	Tags  []string `yaml:"tags"`
	Ports []int    `yaml:"ports"`
}

// 載入基本配置檔案
func loadBasicConfig(filename string) BasicConfig {
	var config BasicConfig
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
