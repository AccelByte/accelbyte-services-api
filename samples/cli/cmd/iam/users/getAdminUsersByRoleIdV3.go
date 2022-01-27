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

// getAdminUsersByRoleIdV3Cmd represents the getAdminUsersByRoleIdV3 command
var getAdminUsersByRoleIdV3Cmd = &cobra.Command{
	Use:   "getAdminUsersByRoleIdV3",
	Short: "Get admin users by role id V3",
	Long:  `Get admin users by role id V3`,
	RunE: func(cmd *cobra.Command, args []string) error {
		usersService := &iam.UsersService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		roleId, _ := cmd.Flags().GetString("roleId")
		after, _ := cmd.Flags().GetInt64("after")
		before, _ := cmd.Flags().GetInt64("before")
		limit, _ := cmd.Flags().GetInt64("limit")
		input := &users.GetAdminUsersByRoleIDV3Params{
			Namespace: namespace,
			RoleID:    roleId,
			After:     &after,
			Before:    &before,
			Limit:     &limit,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := usersService.GetAdminUsersByRoleIDV3(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(getAdminUsersByRoleIdV3Cmd)
	getAdminUsersByRoleIdV3Cmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = getAdminUsersByRoleIdV3Cmd.MarkFlagRequired("namespace")
	getAdminUsersByRoleIdV3Cmd.Flags().StringP("roleId", "rd", " ", "Role id")
	_ = getAdminUsersByRoleIdV3Cmd.MarkFlagRequired("roleId")
	getAdminUsersByRoleIdV3Cmd.Flags().Int64P("after", "ar", 0, "After")
	getAdminUsersByRoleIdV3Cmd.Flags().Int64P("before", "be", 0, "Before")
	getAdminUsersByRoleIdV3Cmd.Flags().Int64P("limit", "lt", 20, "Limit")
}
