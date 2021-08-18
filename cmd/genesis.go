package cmd

/*
	special cmd to generate a genesis file payload
*/

import (
	"github.com/spf13/cobra"
	"log"
)

const EMPTYBLOCKROOT = "Llx000000002069732061207368697474792073656e74656e636520746f2074657374Lx41444452"

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
