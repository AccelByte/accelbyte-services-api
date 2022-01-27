// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package fileUpload

import (
	"github.com/AccelByte/accelbyte-go-sdk/basic-sdk/pkg/basicclient/file_upload"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/basic"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// publicGeneratedUploadUrlCmd represents the publicGeneratedUploadUrl command
var publicGeneratedUploadUrlCmd = &cobra.Command{
	Use:   "publicGeneratedUploadUrl",
	Short: "Public generated upload url",
	Long:  `Public generated upload url`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fileUploadService := &basic.FileUploadService{
			Client:          factory.NewBasicClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		folder, _ := cmd.Flags().GetString("folder")
		namespace, _ := cmd.Flags().GetString("namespace")
		fileType, _ := cmd.Flags().GetString("fileType")
		input := &file_upload.PublicGeneratedUploadURLParams{
			Folder:    folder,
			Namespace: namespace,
			FileType:  fileType,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := fileUploadService.PublicGeneratedUploadURL(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(publicGeneratedUploadUrlCmd)
	publicGeneratedUploadUrlCmd.Flags().StringP("folder", "fr", " ", "Folder")
	_ = publicGeneratedUploadUrlCmd.MarkFlagRequired("folder")
	publicGeneratedUploadUrlCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = publicGeneratedUploadUrlCmd.MarkFlagRequired("namespace")
	publicGeneratedUploadUrlCmd.Flags().StringP("fileType", "fe", " ", "File type")
	_ = publicGeneratedUploadUrlCmd.MarkFlagRequired("fileType")
}
