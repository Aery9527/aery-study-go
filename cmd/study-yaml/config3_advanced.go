package main

import "time"

// AdvancedConfig 特殊標籤配置 (對應 config3.yaml)
type AdvancedConfig struct {
	Database struct {
		Connections struct {
			Primary struct {
				Host     string `yaml:"host"`
				Port     int    `yaml:"port"`
				Username string `yaml:"username"`
				Password string `yaml:"password"`
				SSLMode  string `yaml:"ssl_mode"`
			} `yaml:"primary"`
			Replica struct {
				Host     string `yaml:"host"`
				Port     int    `yaml:"port"`
				Username string `yaml:"username"`
				Password string `yaml:"password"`
				SSLMode  string `yaml:"ssl_mode"`
			} `yaml:"replica"`
		} `yaml:"connections"`
	} `yaml:"database"`
	Timeouts struct {
		Connect time.Duration `yaml:"connect"`
		Read    time.Duration `yaml:"read"`
		Write   time.Duration `yaml:"write"`
		Idle    time.Duration `yaml:"idle"`
	} `yaml:"timeouts"`
	OptionalConfig struct {
		FeatureFlags struct {
			NewUI        bool  `yaml:"new_ui"`
			BetaFeatures *bool `yaml:"beta_features"` // 指標類型，可為 nil
			Experimental bool  `yaml:"experimental"`
		} `yaml:"feature_flags"`
		Limits struct {
			MaxConnections int            `yaml:"max_connections"`
			RateLimit      *int           `yaml:"rate_limit"` // 指標類型，可為 nil
			Timeout        *time.Duration `yaml:"timeout"`    // 指標類型，可為 nil
		} `yaml:"limits"`
	} `yaml:"optional_config"`
	Services []ServiceConfig `yaml:"services"`
}

// ServiceConfig 服務配置 - 支援多種服務類型
type ServiceConfig struct {
	Type   string                 `yaml:"type"`
	Name   string                 `yaml:"name"`
	Config map[string]interface{} `yaml:"config"`
}

// HTTPServiceConfig HTTP 服務專用配置
type HTTPServiceConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	SSL  bool   `yaml:"ssl"`
}

// GRPCServiceConfig gRPC 服務專用配置
type GRPCServiceConfig struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	TLSCert string `yaml:"tls_cert"`
	TLSKey  string `yaml:"tls_key"`
}

// 載入高級配置檔案
func loadAdvancedConfig(filename string) AdvancedConfig {
	var config AdvancedConfig
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

// 解析 HTTP 服務配置
func parseHTTPConfig(config map[string]interface{}) *HTTPServiceConfig {
	yamlData, err := marshalYAML(config)
	if err != nil {
		return nil
	}

	var httpConfig HTTPServiceConfig
	err = unmarshalYAML(yamlData, &httpConfig)
	if err != nil {
		return nil
	}

	return &httpConfig
}

// 解析 gRPC 服務配置
func parseGRPCConfig(config map[string]interface{}) *GRPCServiceConfig {
	yamlData, err := marshalYAML(config)
	if err != nil {
		return nil
	}

	var grpcConfig GRPCServiceConfig
	err = unmarshalYAML(yamlData, &grpcConfig)
	if err != nil {
		return nil
	}

	return &grpcConfig
}
