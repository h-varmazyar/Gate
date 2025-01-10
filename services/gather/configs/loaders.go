package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func loadString(envName string) string {
	validate(envName)
	return viper.GetString(envName)
}

func loadInt(envName string) int {
	validate(envName)
	return viper.GetInt(envName)
}

func loadUint(envName string) uint {
	validate(envName)
	return viper.GetUint(envName)
}

func loadInt64(envName string) int64 {
	validate(envName)
	return viper.GetInt64(envName)
}

func loadBool(envName string) bool {
	validate(envName)
	return viper.GetBool(envName)
}

//func loadUint(envName string) uint {
//	validate(envName)
//	return viper.GetUint(envName)
//}

func loadFloat(envName string) float64 {
	validate(envName)
	return viper.GetFloat64(envName)
}

func loadStringSlice(envName string) []string {
	validate(envName)
	return viper.GetStringSlice(envName)
}

// nolint:unused
func loadDuration(envName string) time.Duration {
	validate(envName)
	return viper.GetDuration(envName)
}

func validate(envName string) error {
	exists := viper.IsSet(envName)
	if !exists {
		return fmt.Errorf("environment variable [%s] does not exist", envName)
	}
	return nil
}
