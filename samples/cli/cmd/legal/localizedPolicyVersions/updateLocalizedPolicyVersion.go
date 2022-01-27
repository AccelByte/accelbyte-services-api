// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package localizedPolicyVersions

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/localized_policy_versions"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/legal"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// updateLocalizedPolicyVersionCmd represents the updateLocalizedPolicyVersion command
var updateLocalizedPolicyVersionCmd = &cobra.Command{
	Use:   "updateLocalizedPolicyVersion",
	Short: "Update localized policy version",
	Long:  `Update localized policy version`,
	RunE: func(cmd *cobra.Command, args []string) error {
		localizedPolicyVersionsService := &legal.LocalizedPolicyVersionsService{
			Client:          factory.NewLegalClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		localizedPolicyVersionId, _ := cmd.Flags().GetString("localizedPolicyVersionId")
		bodyString := cmd.Flag("body").Value.String()
		var body *legalclientmodels.UpdateLocalizedPolicyVersionRequest
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		input := &localized_policy_versions.UpdateLocalizedPolicyVersionParams{
			Body:                     body,
			LocalizedPolicyVersionID: localizedPolicyVersionId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := localizedPolicyVersionsService.UpdateLocalizedPolicyVersion(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(updateLocalizedPolicyVersionCmd)
	updateLocalizedPolicyVersionCmd.Flags().StringP("body", "by", " ", "Body")
	updateLocalizedPolicyVersionCmd.Flags().StringP("localizedPolicyVersionId", "ld", " ", "Localized policy version id")
	_ = updateLocalizedPolicyVersionCmd.MarkFlagRequired("localizedPolicyVersionId")
}
