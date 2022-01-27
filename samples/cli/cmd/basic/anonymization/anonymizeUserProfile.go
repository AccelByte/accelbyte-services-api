// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package anonymization

import (
	"github.com/AccelByte/accelbyte-go-sdk/basic-sdk/pkg/basicclient/anonymization"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/basic"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// anonymizeUserProfileCmd represents the anonymizeUserProfile command
var anonymizeUserProfileCmd = &cobra.Command{
	Use:   "anonymizeUserProfile",
	Short: "Anonymize user profile",
	Long:  `Anonymize user profile`,
	RunE: func(cmd *cobra.Command, args []string) error {
		anonymizationService := &basic.AnonymizationService{
			Client:          factory.NewBasicClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		input := &anonymization.AnonymizeUserProfileParams{
			Namespace: namespace,
			UserID:    userId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := anonymizationService.AnonymizeUserProfile(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(anonymizeUserProfileCmd)
	anonymizeUserProfileCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = anonymizeUserProfileCmd.MarkFlagRequired("namespace")
	anonymizeUserProfileCmd.Flags().StringP("userId", "ud", " ", "User id")
	_ = anonymizeUserProfileCmd.MarkFlagRequired("userId")
}
