// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package achievements

import (
	"github.com/AccelByte/accelbyte-go-sdk/achievement-sdk/pkg/achievementclient/achievements"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/achievement"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AdminUnlockAchievementCmd represents the AdminUnlockAchievement command
var AdminUnlockAchievementCmd = &cobra.Command{
	Use:   "adminUnlockAchievement",
	Short: "Admin unlock achievement",
	Long:  `Admin unlock achievement`,
	RunE: func(cmd *cobra.Command, args []string) error {
		achievementsService := &achievement.AchievementsService{
			Client:          factory.NewAchievementClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		achievementCode, _ := cmd.Flags().GetString("achievementCode")
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		input := &achievements.AdminUnlockAchievementParams{
			AchievementCode: achievementCode,
			Namespace:       namespace,
			UserID:          userId,
		}
		errInput := achievementsService.AdminUnlockAchievementShort(input)
		if errInput != nil {
			logrus.Error(errInput)

			return errInput
		}

		return nil
	},
}

func init() {
	AdminUnlockAchievementCmd.Flags().StringP("achievementCode", "", "", "Achievement code")
	_ = AdminUnlockAchievementCmd.MarkFlagRequired("achievementCode")
	AdminUnlockAchievementCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = AdminUnlockAchievementCmd.MarkFlagRequired("namespace")
	AdminUnlockAchievementCmd.Flags().StringP("userId", "", "", "User id")
	_ = AdminUnlockAchievementCmd.MarkFlagRequired("userId")
}
