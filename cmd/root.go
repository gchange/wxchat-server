package cmd

import (
	"fmt"
	"wx-server/internal/logging"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "wxchat-server",
		Short: "A wechat server.",
		Long:  `A wechat server.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	var lc logging.ZapConfig
	sv := viper.Sub("logger")
	if sv != nil {
		err := sv.Unmarshal(&lc)
		if err != nil {
			fmt.Printf("init logger config error: %s\n", err.Error())
			return
		}
	}
	logging.Init(&lc)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./wxchat-server.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".wxchat-server")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
