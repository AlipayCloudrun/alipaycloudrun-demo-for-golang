package util

import "os"

func GetEnvDefault(key, defVal string) string {
	env, exist := os.LookupEnv(key)
	if !exist {
		return defVal
	}

	return env
}
