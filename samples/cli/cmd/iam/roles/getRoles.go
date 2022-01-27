// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package roles

import (
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/roles"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getRolesCmd represents the getRoles command
var getRolesCmd = &cobra.Command{
	Use:   "getRoles",
	Short: "Get roles",
	Long:  `Get roles`,
	RunE: func(cmd *cobra.Command, args []string) error {
		rolesService := &iam.RolesService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		isWildcard, _ := cmd.Flags().GetString("isWildcard")
		input := &roles.GetRolesParams{
			IsWildcard: &isWildcard,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := rolesService.GetRoles(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(getRolesCmd)
	getRolesCmd.Flags().StringP("isWildcard", "id", " ", "Is wildcard")
}
