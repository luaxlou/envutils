package envutils

import (
	"github.com/joho/godotenv"
	"github.com/luaxlou/fileutils"
	"log"
	"os"
)

var currAppEnv = "dev"

const (
	Prod = "prod"
	Test = "test"
	Dev  = "dev"
)

func init() {

	if env := os.Getenv("APP_ENV"); env != "" {

		SetEnv(env)
	}

}

func Getenv(key string) string {

	return os.Getenv(key)
}

func MustGetenv(key string) string {

	v := os.Getenv(key)

	if v == "" {
		panic("env", key, "must required")
	}

	return v
}

// @Deprecated use SetAppEnv
func SetEnv(env string) {

	currAppEnv = env

}

// @Deprecated use GetAppEnv
func GetEnv() string {
	return currAppEnv
}

func GetAppEnv() string {
	return currAppEnv
}

func IsProd() bool {
	return currAppEnv == Prod

}

func IsTest() bool {
	return currAppEnv == Test

}

func IsDev() bool {

	return currAppEnv == Dev

}

// 当处于开发环境时，加载当前目录Env文件
func LoadEnvOnlyDev() {

	if IsDev() {
		LoadEnv()
	}
}

func LoadEnvByPathname(pathname string) {
	err := godotenv.Load(pathname)
	if err != nil {
		log.Fatal(err)
	}

}

// 加载当前目录Env文件
func LoadEnv() {
	//多目录检查主要是为了测试代码对于环境变量的加载
	checkPaths := []string{
		"./.env",
		"~/.env",
		"../.env",
		"../../.env",
		"../../../.env",
		"../../../../.env",
		"../../../../../.env",
	}

	for _, p := range checkPaths {
		if fileutils.Exists(p) {
			LoadEnvByPathname(p)
			return
		}
	}

}
