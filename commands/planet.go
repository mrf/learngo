// Copyright © 2014 Mark Ferree.

package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "learngo",
	Short: `Learngo is an awesome planet style RSS aggregator`,
	Long:  `longo long long`,
  Run: rootRun,
}

func rootRun(cmd *cobra.Command, args []string) {
  fmt.Println(viper.Get("feeds"))
  fmt.Println(viper.GetString("appname"))
}

func Execute() {
  addCommands()
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

var CfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is $HOME/learngo/config.yaml)")
}

func initConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/learngo/")
	viper.AddConfigPath("/home/mrf/.learngo/")
	viper.ReadInConfig()
}

func addCommands() {
  RootCmd.AddCommand(fetchCmd)
}
