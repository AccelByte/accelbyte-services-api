// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package users

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/users"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// SaveUserRolesCmd represents the SaveUserRoles command
var SaveUserRolesCmd = &cobra.Command{
	Use:   "saveUserRoles",
	Short: "Save user roles",
	Long:  `Save user roles`,
	RunE: func(cmd *cobra.Command, args []string) error {
		usersService := &iam.UsersService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		bodyString := cmd.Flag("body").Value.String()
		var body []string
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		input := &users.SaveUserRolesParams{
			Body:      body,
			Namespace: namespace,
			UserID:    userId,
		}
		errInput := usersService.SaveUserRolesShort(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	SaveUserRolesCmd.Flags().StringP("body", "", "", "Body")
	_ = SaveUserRolesCmd.MarkFlagRequired("body")
	SaveUserRolesCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = SaveUserRolesCmd.MarkFlagRequired("namespace")
	SaveUserRolesCmd.Flags().StringP("userId", "", "", "User id")
	_ = SaveUserRolesCmd.MarkFlagRequired("userId")
}
