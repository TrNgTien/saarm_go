package utilities

func GetValueOrDefault(val interface{}, defaultVal interface{}) interface{} {

  if val == nil || val == "" {
    return defaultVal
  }

  return val
}

