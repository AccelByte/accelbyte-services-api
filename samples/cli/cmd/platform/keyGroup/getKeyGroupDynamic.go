// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package keyGroup

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/key_group"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getKeyGroupDynamicCmd represents the getKeyGroupDynamic command
var getKeyGroupDynamicCmd = &cobra.Command{
	Use:   "getKeyGroupDynamic",
	Short: "Get key group dynamic",
	Long:  `Get key group dynamic`,
	RunE: func(cmd *cobra.Command, args []string) error {
		keyGroupService := &platform.KeyGroupService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		keyGroupId, _ := cmd.Flags().GetString("keyGroupId")
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &key_group.GetKeyGroupDynamicParams{
			KeyGroupID: keyGroupId,
			Namespace:  namespace,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := keyGroupService.GetKeyGroupDynamic(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(getKeyGroupDynamicCmd)
	getKeyGroupDynamicCmd.Flags().StringP("keyGroupId", "kd", " ", "Key group id")
	_ = getKeyGroupDynamicCmd.MarkFlagRequired("keyGroupId")
	getKeyGroupDynamicCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = getKeyGroupDynamicCmd.MarkFlagRequired("namespace")
}
