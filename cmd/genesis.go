package cmd

/*
	special cmd to generate a genesis file payload
*/

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	GenesisCMD = &cobra.Command{
		Use:   "genesis",
		Short: "Create and manage genesis block",
		Long:  "Create and manage genesis block (need generated operator keys)",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("running genesis")
		},
	}
)
