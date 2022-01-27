// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package wallet

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/wallet"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getWalletCmd represents the getWallet command
var getWalletCmd = &cobra.Command{
	Use:   "getWallet",
	Short: "Get wallet",
	Long:  `Get wallet`,
	RunE: func(cmd *cobra.Command, args []string) error {
		walletService := &platform.WalletService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		walletId, _ := cmd.Flags().GetString("walletId")
		input := &wallet.GetWalletParams{
			Namespace: namespace,
			WalletID:  walletId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := walletService.GetWallet(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(getWalletCmd)
	getWalletCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = getWalletCmd.MarkFlagRequired("namespace")
	getWalletCmd.Flags().StringP("walletId", "wd", " ", "Wallet id")
	_ = getWalletCmd.MarkFlagRequired("walletId")
}
