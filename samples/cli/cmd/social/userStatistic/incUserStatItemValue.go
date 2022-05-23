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

// IncUserStatItemValueCmd represents the IncUserStatItemValue command
var IncUserStatItemValueCmd = &cobra.Command{
	Use:   "incUserStatItemValue",
	Short: "Inc user stat item value",
	Long:  `Inc user stat item value`,
	RunE: func(cmd *cobra.Command, args []string) error {
		userStatisticService := &social.UserStatisticService{
			Client:          factory.NewSocialClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		statCode, _ := cmd.Flags().GetString("statCode")
		userId, _ := cmd.Flags().GetString("userId")
		bodyString := cmd.Flag("body").Value.String()
		var body *socialclientmodels.StatItemInc
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		input := &user_statistic.IncUserStatItemValueParams{
			Body:      body,
			Namespace: namespace,
			StatCode:  statCode,
			UserID:    userId,
		}
		ok, err := userStatisticService.IncUserStatItemValueShort(input)
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
	IncUserStatItemValueCmd.Flags().String("body", "", "Body")
	IncUserStatItemValueCmd.Flags().String("namespace", "", "Namespace")
	_ = IncUserStatItemValueCmd.MarkFlagRequired("namespace")
	IncUserStatItemValueCmd.Flags().String("statCode", "", "Stat code")
	_ = IncUserStatItemValueCmd.MarkFlagRequired("statCode")
	IncUserStatItemValueCmd.Flags().String("userId", "", "User id")
	_ = IncUserStatItemValueCmd.MarkFlagRequired("userId")
}
