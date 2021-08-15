package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "lama",
		Short: "Llama is a next generation blockchain for fun and no drama!",
		Long:  "Llama is a next generation blockchain for fun and no drama!",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	// Used for flags.
	cfgFile            string
	mainAccountAddress string
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config/.lama.json)")
	rootCmd.PersistentFlags().StringVarP(&mainAccountAddress, "account", "w", "", "LLx Address to use this node (your wallet address)")
	viper.BindPFlag("mainAddress", rootCmd.PersistentFlags().Lookup("mainAccountAddress"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home := "./config"

		// Search config in home directory with name ".lama" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".lama")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
