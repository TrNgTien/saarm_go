package utilities

func GetValueOrDefault(val interface{}, defaultVal interface{}) interface{} {

	if val == nil || val == "" {
		return defaultVal
	}

	return val
}

func GetValueEnv(val, defaultVal string) string {
	if val == "" {
		return defaultVal
	}

	return val
}

func ArrayIncludeString(arr []string, target string) bool {
  for _, item := range arr {
    if item == target {
      return true
    }
  }
  return false
}
