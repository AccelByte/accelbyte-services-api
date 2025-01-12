// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package dlc

import (
	"encoding/json"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/dlc"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// UpdatePlatformDLCConfigCmd represents the UpdatePlatformDLCConfig command
var UpdatePlatformDLCConfigCmd = &cobra.Command{
	Use:   "updatePlatformDLCConfig",
	Short: "Update platform DLC config",
	Long:  `Update platform DLC config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dlcService := &platform.DLCService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		bodyString := cmd.Flag("body").Value.String()
		var body *platformclientmodels.PlatformDLCConfigUpdate
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &dlc.UpdatePlatformDLCConfigParams{
			Body:      body,
			Namespace: namespace,
		}
		ok, errOK := dlcService.UpdatePlatformDLCConfigShort(input)
		if errOK != nil {
			logrus.Error(errOK)

			return errOK
		}

		logrus.Infof("Response CLI success: %+v", ok)

		return nil
	},
}

func init() {
	UpdatePlatformDLCConfigCmd.Flags().String("body", "", "Body")
	_ = UpdatePlatformDLCConfigCmd.MarkFlagRequired("body")
	UpdatePlatformDLCConfigCmd.Flags().String("namespace", "", "Namespace")
	_ = UpdatePlatformDLCConfigCmd.MarkFlagRequired("namespace")
}
