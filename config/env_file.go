package config

import (
	"encoding/json"
	"os"
	"sync"
)

var onceConfig sync.Once

// type global
var (
	payload map[string]string
)

func GetConfig(projectPath string) map[string]string {
	if len(payload) > 0 {
		return payload
	}

	onceConfig.Do(func() {
		// 生成配置
		envFile, err2 := os.Getwd()
		if err2 != nil {
			return
		}
		if os.Getenv("CODENATION_ENV") == "prod" {
			envFile = envFile + projectPath + "/env-prod.json"
		} else if os.Getenv("CODENATION_ENV") == "test" {
			envFile = envFile + projectPath + "/env-test.json"
		} else {
			envFile = envFile + projectPath + "/env-dev.json"
		}
		content, err := os.ReadFile(envFile)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(content, &payload)
		if err != nil {
			panic(err)
		}
	})

	return payload
}
