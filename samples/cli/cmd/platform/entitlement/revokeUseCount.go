// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package entitlement

import (
	"encoding/json"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/entitlement"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// RevokeUseCountCmd represents the RevokeUseCount command
var RevokeUseCountCmd = &cobra.Command{
	Use:   "revokeUseCount",
	Short: "Revoke use count",
	Long:  `Revoke use count`,
	RunE: func(cmd *cobra.Command, args []string) error {
		entitlementService := &platform.EntitlementService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		bodyString := cmd.Flag("body").Value.String()
		var body *platformclientmodels.RevokeUseCountRequest
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		entitlementId, _ := cmd.Flags().GetString("entitlementId")
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		input := &entitlement.RevokeUseCountParams{
			Body:          body,
			EntitlementID: entitlementId,
			Namespace:     namespace,
			UserID:        userId,
		}
		ok, errOK := entitlementService.RevokeUseCountShort(input)
		if errOK != nil {
			logrus.Error(errOK)

			return errOK
		}

		logrus.Infof("Response CLI success: %+v", ok)

		return nil
	},
}

func init() {
	RevokeUseCountCmd.Flags().String("body", "", "Body")
	_ = RevokeUseCountCmd.MarkFlagRequired("body")
	RevokeUseCountCmd.Flags().String("entitlementId", "", "Entitlement id")
	_ = RevokeUseCountCmd.MarkFlagRequired("entitlementId")
	RevokeUseCountCmd.Flags().String("namespace", "", "Namespace")
	_ = RevokeUseCountCmd.MarkFlagRequired("namespace")
	RevokeUseCountCmd.Flags().String("userId", "", "User id")
	_ = RevokeUseCountCmd.MarkFlagRequired("userId")
}
