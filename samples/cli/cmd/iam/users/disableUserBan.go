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

// DisableUserBanCmd represents the DisableUserBan command
var DisableUserBanCmd = &cobra.Command{
	Use:   "disableUserBan",
	Short: "Disable user ban",
	Long:  `Disable user ban`,
	RunE: func(cmd *cobra.Command, args []string) error {
		usersService := &iam.UsersService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		banId, _ := cmd.Flags().GetString("banId")
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		httpClient := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		input := &users.DisableUserBanParams{
			BanID:      banId,
			Namespace:  namespace,
			UserID:     userId,
			HTTPClient: httpClient,
		}
		ok, err := usersService.DisableUserBanShort(input)
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
	DisableUserBanCmd.Flags().StringP("banId", "", "", "Ban id")
	_ = DisableUserBanCmd.MarkFlagRequired("banId")
	DisableUserBanCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = DisableUserBanCmd.MarkFlagRequired("namespace")
	DisableUserBanCmd.Flags().StringP("userId", "", "", "User id")
	_ = DisableUserBanCmd.MarkFlagRequired("userId")
}
