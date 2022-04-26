// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package imageConfig

import (
	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclient/image_config"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/dsmc"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// DeleteImageCmd represents the DeleteImage command
var DeleteImageCmd = &cobra.Command{
	Use:   "deleteImage",
	Short: "Delete image",
	Long:  `Delete image`,
	RunE: func(cmd *cobra.Command, args []string) error {
		imageConfigService := &dsmc.ImageConfigService{
			Client:          factory.NewDsmcClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		imageURI, _ := cmd.Flags().GetString("imageURI")
		version, _ := cmd.Flags().GetString("version")
		input := &image_config.DeleteImageParams{
			Namespace: namespace,
			ImageURI:  imageURI,
			Version:   version,
		}
		errInput := imageConfigService.DeleteImageShort(input)
		if errInput != nil {
			logrus.Error(errInput)

			return errInput
		}

		return nil
	},
}

func init() {
	DeleteImageCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = DeleteImageCmd.MarkFlagRequired("namespace")
	DeleteImageCmd.Flags().StringP("imageURI", "", "", "Image URI")
	_ = DeleteImageCmd.MarkFlagRequired("imageURI")
	DeleteImageCmd.Flags().StringP("version", "", "", "Version")
	_ = DeleteImageCmd.MarkFlagRequired("version")
}
