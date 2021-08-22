package cmd

import (
	"github.com/Digital-MOB-Filecoin/find-miner/fmtool"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var log = logrus.WithField("module", "main")

var (
	config            string
	version           bool
	verbose, vverbose bool

	RootCmd = &cobra.Command{
		Use:   "find-miner",
		Short: "find-miner CLI",
		Long:  "The use case for find-miner is to enable a user of the Lotus CLI or API to select a miner for a given type of storage deal.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := viper.BindPFlags(cmd.Flags())
			if err != nil {
				log.Fatal(err)
			}

		},

		Run: func(cmd *cobra.Command, args []string) {
			size := viper.GetInt64("size")
			region := viper.GetString("region")
			verifiedSPL := viper.GetInt64("verified-storage-price-limit")
			skip := viper.GetInt64("skip-miners")
			verified := viper.GetString("verified")
			fastRetrieval := viper.GetString("fastRetrieval")

			x := fmtool.NewWorkerLib(
				size,
				region,
				verifiedSPL,
				skip,
				verified,
				fastRetrieval,
				fmtool.Config{
					RsvAPI: viper.GetString("rsv-api"),
				})

			err := x.Run()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	cobra.OnInitialize(func() {
		viper.Set("version", RootCmd.Version)
	})
	RootCmd.Flags().Int64("size", 0, "Deal size")
	RootCmd.Flags().String("region", "", "Miner's region : ap|cn|na|eu")
	RootCmd.Flags().Int64("verified-storage-price-limit", -1, "Maximum acceptable verified storage price (in FIL)")
	RootCmd.Flags().Int64("skip-miners", 0, "The first N miners that would normally be returned are skipped")
	RootCmd.Flags().String("verified", "null", "Verified")
	RootCmd.Flags().String("fastRetrieval", "null", "Fast Retrieval")

	RootCmd.PersistentFlags().String("rsv-api", "https://api.repsys.d.interplanetary.one/rpc", "RSV api")
}
