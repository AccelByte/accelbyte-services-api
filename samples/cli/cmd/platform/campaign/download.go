// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package campaign

import (
	"bytes"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/campaign"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// DownloadCmd represents the Download command
var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download",
	Long:  `Download`,
	RunE: func(cmd *cobra.Command, args []string) error {
		campaignService := &platform.CampaignService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		campaignId, _ := cmd.Flags().GetString("campaignId")
		namespace, _ := cmd.Flags().GetString("namespace")
		batchNo, _ := cmd.Flags().GetInt32("batchNo")
		input := &campaign.DownloadParams{
			CampaignID: campaignId,
			Namespace:  namespace,
			BatchNo:    &batchNo,
		}
		writer := bytes.NewBuffer(nil)
		ok, err := campaignService.DownloadShort(input, writer)
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
	DownloadCmd.Flags().String("campaignId", "", "Campaign id")
	_ = DownloadCmd.MarkFlagRequired("campaignId")
	DownloadCmd.Flags().String("namespace", "", "Namespace")
	_ = DownloadCmd.MarkFlagRequired("namespace")
	DownloadCmd.Flags().Int32("batchNo", 0, "Batch no")
}
