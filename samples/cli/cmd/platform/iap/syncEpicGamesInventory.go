// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package iap

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/i_a_p"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// syncEpicGamesInventoryCmd represents the syncEpicGamesInventory command
var syncEpicGamesInventoryCmd = &cobra.Command{
	Use:   "syncEpicGamesInventory",
	Short: "Sync epic games inventory",
	Long:  `Sync epic games inventory`,
	RunE: func(cmd *cobra.Command, args []string) error {
		iapService := &platform.IAPService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		bodyString := cmd.Flag("body").Value.String()
		var body *platformclientmodels.EpicGamesReconcileRequest
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		input := &i_a_p.SyncEpicGamesInventoryParams{
			Body:      body,
			Namespace: namespace,
			UserID:    userId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := iapService.SyncEpicGamesInventory(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(syncEpicGamesInventoryCmd)
	syncEpicGamesInventoryCmd.Flags().StringP("body", "by", " ", "Body")
	syncEpicGamesInventoryCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = syncEpicGamesInventoryCmd.MarkFlagRequired("namespace")
	syncEpicGamesInventoryCmd.Flags().StringP("userId", "ud", " ", "User id")
	_ = syncEpicGamesInventoryCmd.MarkFlagRequired("userId")
}
