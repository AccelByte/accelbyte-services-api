// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package cmd

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/users"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getPublicMyUserCmd represents the getMyUser command
var getPublicMyUserCmd = &cobra.Command{
	Use:   "getPublicMyUser",
	Short: "Get public my user",
	Long:  `Get public my user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		userService := &iam.UsersService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		input := &users.PublicGetMyUserV3Params{}
		//nolint:staticcheck // SA1019 To be deprecated later
		//lint:ignore SA1019 Ignore the deprecation warnings
		user, err := userService.PublicGetMyUserV3(input)
		if err != nil {
			return err
		}
		response, err := json.MarshalIndent(user, "", "    ")
		if err != nil {
			return err
		}
		logrus.Infof("Response %v", string(response))
		return nil
	},
}

func init() {
	RootCmd.AddCommand(getPublicMyUserCmd)

}
