package cmd

import "github.com/spf13/cobra"

func addFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("rsv-api", "", "RSV api.")
}
