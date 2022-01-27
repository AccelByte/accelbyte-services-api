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

// addUserRoleCmd represents the addUserRole command
var addUserRoleCmd = &cobra.Command{
	Use:   "addUserRole",
	Short: "Add user role",
	Long:  `Add user role`,
	RunE: func(cmd *cobra.Command, args []string) error {
		usersService := &iam.UsersService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		roleId, _ := cmd.Flags().GetString("roleId")
		userId, _ := cmd.Flags().GetString("userId")
		input := &users.AddUserRoleParams{
			Namespace: namespace,
			RoleID:    roleId,
			UserID:    userId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := usersService.AddUserRole(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(addUserRoleCmd)
	addUserRoleCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = addUserRoleCmd.MarkFlagRequired("namespace")
	addUserRoleCmd.Flags().StringP("roleId", "rd", " ", "Role id")
	_ = addUserRoleCmd.MarkFlagRequired("roleId")
	addUserRoleCmd.Flags().StringP("userId", "ud", " ", "User id")
	_ = addUserRoleCmd.MarkFlagRequired("userId")
}
