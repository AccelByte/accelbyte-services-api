// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package wallet

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/wallet"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// PayWithUserWalletCmd represents the PayWithUserWallet command
var PayWithUserWalletCmd = &cobra.Command{
	Use:   "payWithUserWallet",
	Short: "Pay with user wallet",
	Long:  `Pay with user wallet`,
	RunE: func(cmd *cobra.Command, args []string) error {
		walletService := &platform.WalletService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		currencyCode, _ := cmd.Flags().GetString("currencyCode")
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		bodyString := cmd.Flag("body").Value.String()
		var body *platformclientmodels.PaymentRequest
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		input := &wallet.PayWithUserWalletParams{
			Body:         body,
			CurrencyCode: currencyCode,
			Namespace:    namespace,
			UserID:       userId,
		}
		ok, err := walletService.PayWithUserWalletShort(input)
		if err != nil {
			logrus.Error(err)
			return err
		} else {
			logrus.Infof("Response CLI success: %+v", ok)
		}
		return nil
	},
}

func init() {
	PayWithUserWalletCmd.Flags().StringP("body", "", "", "Body")
	PayWithUserWalletCmd.Flags().StringP("currencyCode", "", "", "Currency code")
	_ = PayWithUserWalletCmd.MarkFlagRequired("currencyCode")
	PayWithUserWalletCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = PayWithUserWalletCmd.MarkFlagRequired("namespace")
	PayWithUserWalletCmd.Flags().StringP("userId", "", "", "User id")
	_ = PayWithUserWalletCmd.MarkFlagRequired("userId")
}
