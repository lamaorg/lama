package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	P2PRPC = &cobra.Command{
		Use:   "rpc",
		Short: "start, stop, send and request the rpc service",
		Long:  "start, stop, send and request the rpc service",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("running libp2p rpc services for LamaChain")
		},
	}
)
