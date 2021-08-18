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
	rootCmd.AddCommand(WalletCMD)
	rootCmd.AddCommand(GenesisCMD)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./config/.lama.json", "config file (default is config/.lama.json)")
	rootCmd.PersistentFlags().StringVarP(&mainAccountAddress, "mainAccountAddress", "w", "", "LLx Address to use this node (your wallet address)")
	viper.BindPFlag("usewallet", rootCmd.PersistentFlags().Lookup("mainAccountAddress"))
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootCmd.PersistentFlags().Bool("testnet", false, "use Llama test network")
	rootCmd.PersistentFlags().Bool("local", false, "setup a local node for Llama")
	WalletCMD.PersistentFlags().Bool("create", false, "create a new wallet")

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

	viper.SetEnvPrefix("lama")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		//viper.Debug()
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

var WalletCMD = &cobra.Command{
	Use:   "wallet",
	Short: "LLama Wallet functions",
	Long:  "LLama Wallet functions",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.DebugFlags()
		fmt.Println("LLama Wallet Operator... ")
	},
}
