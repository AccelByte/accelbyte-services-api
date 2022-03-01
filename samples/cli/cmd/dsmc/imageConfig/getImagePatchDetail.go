// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package imageConfig

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclient/image_config"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/dsmc"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// GetImagePatchDetailCmd represents the GetImagePatchDetail command
var GetImagePatchDetailCmd = &cobra.Command{
	Use:   "getImagePatchDetail",
	Short: "Get image patch detail",
	Long:  `Get image patch detail`,
	RunE: func(cmd *cobra.Command, args []string) error {
		imageConfigService := &dsmc.ImageConfigService{
			Client:          factory.NewDsmcClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		version, _ := cmd.Flags().GetString("version")
		versionPatch, _ := cmd.Flags().GetString("versionPatch")
		input := &image_config.GetImagePatchDetailParams{
			Namespace:    namespace,
			Version:      version,
			VersionPatch: versionPatch,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := imageConfigService.GetImagePatchDetail(input)
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
	GetImagePatchDetailCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = GetImagePatchDetailCmd.MarkFlagRequired("namespace")
	GetImagePatchDetailCmd.Flags().StringP("version", "", "", "Version")
	_ = GetImagePatchDetailCmd.MarkFlagRequired("version")
	GetImagePatchDetailCmd.Flags().StringP("versionPatch", "", "", "Version patch")
	_ = GetImagePatchDetailCmd.MarkFlagRequired("versionPatch")
}
