package cmd

import (
	"log"
	wxchatserver "wx-server/internal/wxchat-server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "wx chat server",
	Long:  `Run server`,
	Run: func(cmd *cobra.Command, args []string) {
		var c wxchatserver.Config
		err := viper.Unmarshal(&c)
		if err != nil {
			log.Panic(err)
		}
		s, err := wxchatserver.NewServer(&c)
		if err != nil {
			log.Panic(err)
			return
		}
		s.Run()
	},
}
