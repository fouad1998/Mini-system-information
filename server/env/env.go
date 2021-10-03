package env

import "github.com/spf13/viper"

type Settings struct {
	Username         string
	Password         string
	StorageDirectory string
	StorageFile      string
	Seperator        string
}

var Setting *Settings

func init() {
	viper.SetDefault("USERNAME", "admin")
	viper.SetDefault("PASSWORD", "admin")
	viper.SetDefault("STORAGE_DIRECTORY", "store")
	viper.SetDefault("STORAGE_FILE", "storage")
	viper.SetDefault("SEPERATOR", ";;||;;")

	Setting = &Settings{
		Username:         viper.GetString("USERNAME"),
		Password:         viper.GetString("PASSWORD"),
		StorageDirectory: viper.GetString("STORAGE_DIRECTORY"),
		StorageFile:      viper.GetString("STORAGE_FILE"),
		Seperator:        viper.GetString("SEPERATOR"),
	}
}
