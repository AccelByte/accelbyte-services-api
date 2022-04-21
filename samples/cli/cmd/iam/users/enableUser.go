// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package users

import (
	"net/http"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/users"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// EnableUserCmd represents the EnableUser command
var EnableUserCmd = &cobra.Command{
	Use:   "enableUser",
	Short: "Enable user",
	Long:  `Enable user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		usersService := &iam.UsersService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		httpClient := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		input := &users.EnableUserParams{
			Namespace:  namespace,
			UserID:     userId,
			HTTPClient: httpClient,
		}
		errInput := usersService.EnableUserShort(input)
		if errInput != nil {
			logrus.Error(errInput)

			return errInput
		}

		return nil
	},
}

func init() {
	EnableUserCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = EnableUserCmd.MarkFlagRequired("namespace")
	EnableUserCmd.Flags().StringP("userId", "", "", "User id")
	_ = EnableUserCmd.MarkFlagRequired("userId")
}
