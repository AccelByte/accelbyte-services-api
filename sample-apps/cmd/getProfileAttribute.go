// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getProfileAttribute represents the getProfileAttribute command
var getProfileAttribute = &cobra.Command{
	Use:   "getProfileAttribute",
	Short: "Public Get profile attribute",
	Long:  `Public Get profile attribute`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("getProfileAttribute called")
		gameProfileService := &service.GameProfileService{
			SocialServiceClient: factory.NewSocialClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository:     &repository.TokenRepositoryImpl{},
		}
		namespace := cmd.Flag("namespace").Value.String()
		profileId := cmd.Flag("profileId").Value.String()
		userId := cmd.Flag("userId").Value.String()
		attributeName := cmd.Flag("attributeName").Value.String()
		ok, err := gameProfileService.PublicGetProfileAttribute(namespace, userId, attributeName, profileId)
		if err != nil {
			return err
		} else {
			response, err := json.Marshal(ok)
			if err != nil {
				return err
			} else {
				logrus.Infof("Response %s", response)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getProfileAttribute)
	getProfileAttribute.Flags().StringP("namespace", "n", "", "User namespace")
	getProfileAttribute.MarkFlagRequired("namespace")
	getProfileAttribute.Flags().StringP("userId", "u", "", "User ID")
	getProfileAttribute.MarkFlagRequired("userId")
	getProfileAttribute.Flags().StringP("profileId", "p", "", "Profile ID")
	getProfileAttribute.MarkFlagRequired("profileId")
	getProfileAttribute.Flags().StringP("attributeName", "a", "", "Attribute Name")
	getProfileAttribute.MarkFlagRequired("attributeName")
}
