// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package users

import (
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/users"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// deleteUserInformationCmd represents the deleteUserInformation command
var deleteUserInformationCmd = &cobra.Command{
	Use:   "deleteUserInformation",
	Short: "Delete user information",
	Long:  `Delete user information`,
	RunE: func(cmd *cobra.Command, args []string) error {
		usersService := &iam.UsersService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		input := &users.DeleteUserInformationParams{
			Namespace: namespace,
			UserID:    userId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := usersService.DeleteUserInformation(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(deleteUserInformationCmd)
	deleteUserInformationCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = deleteUserInformationCmd.MarkFlagRequired("namespace")
	deleteUserInformationCmd.Flags().StringP("userId", "ud", " ", "User id")
	_ = deleteUserInformationCmd.MarkFlagRequired("userId")
}
