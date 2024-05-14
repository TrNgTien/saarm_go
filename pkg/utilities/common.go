package utilities

import "os"

func GetValueOrDefault(val interface{}, defaultVal interface{}) interface{} {

	if val == nil || val == "" {
		return defaultVal
	}

	return val
}

func GetValueEnv(keyVal, defaultVal string) string {
	envData := os.Getenv(keyVal)

	if envData == "" {
		return defaultVal
	}

	return os.Getenv(keyVal)
}

func ArrayIncludeString(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}
