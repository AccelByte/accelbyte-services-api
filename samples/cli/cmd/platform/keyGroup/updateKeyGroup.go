// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package keyGroup

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/key_group"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// UpdateKeyGroupCmd represents the UpdateKeyGroup command
var UpdateKeyGroupCmd = &cobra.Command{
	Use:   "updateKeyGroup",
	Short: "Update key group",
	Long:  `Update key group`,
	RunE: func(cmd *cobra.Command, args []string) error {
		keyGroupService := &platform.KeyGroupService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		keyGroupId, _ := cmd.Flags().GetString("keyGroupId")
		namespace, _ := cmd.Flags().GetString("namespace")
		bodyString := cmd.Flag("body").Value.String()
		var body *platformclientmodels.KeyGroupUpdate
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		input := &key_group.UpdateKeyGroupParams{
			Body:       body,
			KeyGroupID: keyGroupId,
			Namespace:  namespace,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := keyGroupService.UpdateKeyGroup(input)
		if err != nil {
			logrus.Error(err)
			return err
		} else {
			response, errIndent := json.MarshalIndent(ok, "", "    ")
			if errIndent != nil {
				return errIndent
			} else {
				logrus.Infof("Response %s", string(response))
			}
		}
		return nil
	},
}

func init() {
	UpdateKeyGroupCmd.Flags().StringP("body", "", "", "Body")
	UpdateKeyGroupCmd.Flags().StringP("keyGroupId", "", "", "Key group id")
	_ = UpdateKeyGroupCmd.MarkFlagRequired("keyGroupId")
	UpdateKeyGroupCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = UpdateKeyGroupCmd.MarkFlagRequired("namespace")
}
