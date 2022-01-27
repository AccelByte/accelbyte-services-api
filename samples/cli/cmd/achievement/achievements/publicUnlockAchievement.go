// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package achievements

import (
	"github.com/AccelByte/accelbyte-go-sdk/achievement-sdk/pkg/achievementclient/achievements"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/achievement"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// publicUnlockAchievementCmd represents the publicUnlockAchievement command
var publicUnlockAchievementCmd = &cobra.Command{
	Use:   "publicUnlockAchievement",
	Short: "Public unlock achievement",
	Long:  `Public unlock achievement`,
	RunE: func(cmd *cobra.Command, args []string) error {
		achievementsService := &achievement.AchievementsService{
			Client:          factory.NewAchievementClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		achievementCode, _ := cmd.Flags().GetString("achievementCode")
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		input := &achievements.PublicUnlockAchievementParams{
			AchievementCode: achievementCode,
			Namespace:       namespace,
			UserID:          userId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := achievementsService.PublicUnlockAchievement(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(publicUnlockAchievementCmd)
	publicUnlockAchievementCmd.Flags().StringP("achievementCode", "ae", " ", "Achievement code")
	_ = publicUnlockAchievementCmd.MarkFlagRequired("achievementCode")
	publicUnlockAchievementCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = publicUnlockAchievementCmd.MarkFlagRequired("namespace")
	publicUnlockAchievementCmd.Flags().StringP("userId", "ud", " ", "User id")
	_ = publicUnlockAchievementCmd.MarkFlagRequired("userId")
}
