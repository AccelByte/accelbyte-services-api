// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package policies

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/policies"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/legal"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// UpdatePolicyCmd represents the UpdatePolicy command
var UpdatePolicyCmd = &cobra.Command{
	Use:   "updatePolicy",
	Short: "Update policy",
	Long:  `Update policy`,
	RunE: func(cmd *cobra.Command, args []string) error {
		policiesService := &legal.PoliciesService{
			Client:          factory.NewLegalClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		policyId, _ := cmd.Flags().GetString("policyId")
		bodyString := cmd.Flag("body").Value.String()
		var body *legalclientmodels.UpdatePolicyRequest
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		input := &policies.UpdatePolicyParams{
			Body:     body,
			PolicyID: policyId,
		}
		errInput := policiesService.UpdatePolicyShort(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	UpdatePolicyCmd.Flags().StringP("body", "", "", "Body")
	UpdatePolicyCmd.Flags().StringP("policyId", "", "", "Policy id")
	_ = UpdatePolicyCmd.MarkFlagRequired("policyId")
}
