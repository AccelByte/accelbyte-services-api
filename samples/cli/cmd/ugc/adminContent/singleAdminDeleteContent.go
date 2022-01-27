// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package adminContent

import (
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/ugc"
	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclient/admin_content"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// singleAdminDeleteContentCmd represents the singleAdminDeleteContent command
var singleAdminDeleteContentCmd = &cobra.Command{
	Use:   "singleAdminDeleteContent",
	Short: "Single admin delete content",
	Long:  `Single admin delete content`,
	RunE: func(cmd *cobra.Command, args []string) error {
		adminContentService := &ugc.AdminContentService{
			Client:          factory.NewUgcClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		channelId, _ := cmd.Flags().GetString("channelId")
		contentId, _ := cmd.Flags().GetString("contentId")
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &admin_content.SingleAdminDeleteContentParams{
			ChannelID: channelId,
			ContentID: contentId,
			Namespace: namespace,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := adminContentService.SingleAdminDeleteContent(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(singleAdminDeleteContentCmd)
	singleAdminDeleteContentCmd.Flags().StringP("channelId", "cd", " ", "Channel id")
	_ = singleAdminDeleteContentCmd.MarkFlagRequired("channelId")
	singleAdminDeleteContentCmd.Flags().StringP("contentId", "cd", " ", "Content id")
	_ = singleAdminDeleteContentCmd.MarkFlagRequired("contentId")
	singleAdminDeleteContentCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = singleAdminDeleteContentCmd.MarkFlagRequired("namespace")
}
