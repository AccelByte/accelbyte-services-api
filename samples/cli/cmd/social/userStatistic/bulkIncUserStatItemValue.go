// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package userStatistic

import (
	"encoding/json"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/social"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclient/user_statistic"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// BulkIncUserStatItemValueCmd represents the BulkIncUserStatItemValue command
var BulkIncUserStatItemValueCmd = &cobra.Command{
	Use:   "bulkIncUserStatItemValue",
	Short: "Bulk inc user stat item value",
	Long:  `Bulk inc user stat item value`,
	RunE: func(cmd *cobra.Command, args []string) error {
		userStatisticService := &social.UserStatisticService{
			Client:          factory.NewSocialClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		bodyString := cmd.Flag("body").Value.String()
		var body []*socialclientmodels.BulkUserStatItemInc
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &user_statistic.BulkIncUserStatItemValueParams{
			Body:      body,
			Namespace: namespace,
		}
		ok, errOK := userStatisticService.BulkIncUserStatItemValueShort(input)
		if errOK != nil {
			logrus.Error(errOK)

			return errOK
		}

		logrus.Infof("Response CLI success: %+v", ok)

		return nil
	},
}

func init() {
	BulkIncUserStatItemValueCmd.Flags().String("body", "", "Body")
	_ = BulkIncUserStatItemValueCmd.MarkFlagRequired("body")
	BulkIncUserStatItemValueCmd.Flags().String("namespace", "", "Namespace")
	_ = BulkIncUserStatItemValueCmd.MarkFlagRequired("namespace")
}
