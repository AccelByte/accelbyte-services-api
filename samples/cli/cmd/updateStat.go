// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/social"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclient/stat_configuration"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// updateStatCmd represents the updateStat command
var updateStatCmd = &cobra.Command{
	Use:   "updateStat",
	Short: "Update Stat",
	Long:  `Update Stat`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("updateStat called")
		socialService := &social.StatConfigurationService{
			Client:          factory.NewSocialClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace := cmd.Flag("namespace").Value.String()
		statCode := cmd.Flag("statCode").Value.String()
		bodyString := cmd.Flag("body").Value.String()
		var body *socialclientmodels.StatUpdate
		errContent := json.Unmarshal([]byte(bodyString), &body)
		if errContent != nil {
			return errContent
		}
		input := &stat_configuration.UpdateStatParams{
			Body:      body,
			Namespace: namespace,
			StatCode:  statCode,
		}
		//nolint:staticcheck // SA1019 To be deprecated later
		//lint:ignore SA1019 Ignore the deprecation warnings
		stat, err := socialService.UpdateStat(input)
		if err != nil {
			return err
		}
		response, err := json.MarshalIndent(stat, "", "    ")
		if err != nil {
			return err
		}
		logrus.Infof("Response %v", string(response))
		return nil
	},
}

func init() {
	RootCmd.AddCommand(updateStatCmd)
	updateStatCmd.Flags().StringP("namespace", "n", "", "User namespace")
	_ = updateStatCmd.MarkFlagRequired("namespace")
	updateStatCmd.Flags().StringP("statCode", "s", "", "Stat Code")
	updateStatCmd.Flags().StringP("body", "b", "", "Slot Body. Example `{\"Key1\":\"Value1\",\"Key2\":\"Value2\"}'")
	_ = updateStatCmd.MarkFlagRequired("body")
}
