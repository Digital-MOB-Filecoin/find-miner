package cmd

import (
	"github.com/Digital-MOB-Filecoin/find-miner/fmtool"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find miners",
	Long:  "Find miners",
	Run: func(cmd *cobra.Command, args []string) {

		size := viper.GetInt64("size")
		region := viper.GetString("region")
		verifiedSPL := viper.GetInt64("verified-storage-price-limit")
		skip := viper.GetInt64("skip-miners")

		x := fmtool.NewWorkerLib(
			size,
			region,
			verifiedSPL,
			skip,
			fmtool.Config{
				TargetURL: viper.GetString("rsv-api"),
			})

		err := x.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(findCmd)
	findCmd.Flags().Int64("size", 0, "Deal size")
	findCmd.Flags().String("region", "", "Miner's region : ap|cn|na|eu")
	findCmd.Flags().Int64("verified-storage-price-limit", -1, "Maximum acceptable verified storage price (in FIL)")
	findCmd.Flags().Int64("skip-miners", 0, "The first N miners that would normally be returned are skipped")
}
