// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package session

import (
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/sessionbrowser"
	"github.com/AccelByte/accelbyte-go-sdk/sessionbrowser-sdk/pkg/sessionbrowserclient/session"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getTotalActiveSessionCmd represents the getTotalActiveSession command
var getTotalActiveSessionCmd = &cobra.Command{
	Use:   "getTotalActiveSession",
	Short: "Get total active session",
	Long:  `Get total active session`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sessionService := &sessionbrowser.SessionService{
			Client:          factory.NewSessionbrowserClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		sessionType, _ := cmd.Flags().GetString("sessionType")
		input := &session.GetTotalActiveSessionParams{
			Namespace:   namespace,
			SessionType: &sessionType,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := sessionService.GetTotalActiveSession(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(getTotalActiveSessionCmd)
	getTotalActiveSessionCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = getTotalActiveSessionCmd.MarkFlagRequired("namespace")
	getTotalActiveSessionCmd.Flags().StringP("session_type", "se", " ", "Session type")
}
