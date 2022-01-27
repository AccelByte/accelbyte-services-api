// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package slotConfig

import (
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/social"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclient/slot_config"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// deleteNamespaceSlotConfigCmd represents the deleteNamespaceSlotConfig command
var deleteNamespaceSlotConfigCmd = &cobra.Command{
	Use:   "deleteNamespaceSlotConfig",
	Short: "Delete namespace slot config",
	Long:  `Delete namespace slot config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		slotConfigService := &social.SlotConfigService{
			Client:          factory.NewSocialClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &slot_config.DeleteNamespaceSlotConfigParams{
			Namespace: namespace,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := slotConfigService.DeleteNamespaceSlotConfig(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(deleteNamespaceSlotConfigCmd)
	deleteNamespaceSlotConfigCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = deleteNamespaceSlotConfigCmd.MarkFlagRequired("namespace")
}
