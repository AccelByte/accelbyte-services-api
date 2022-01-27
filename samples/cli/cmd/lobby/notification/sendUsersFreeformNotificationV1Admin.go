// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package notification

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclient/notification"
	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/lobby"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sendUsersFreeformNotificationV1AdminCmd represents the sendUsersFreeformNotificationV1Admin command
var sendUsersFreeformNotificationV1AdminCmd = &cobra.Command{
	Use:   "sendUsersFreeformNotificationV1Admin",
	Short: "Send users freeform notification V1 admin",
	Long:  `Send users freeform notification V1 admin`,
	RunE: func(cmd *cobra.Command, args []string) error {
		notificationService := &lobby.NotificationService{
			Client:          factory.NewLobbyClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		bodyString := cmd.Flag("body").Value.String()
		var body *lobbyclientmodels.ModelFreeFormNotificationRequestV1
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &notification.SendUsersFreeformNotificationV1AdminParams{
			Body:      body,
			Namespace: namespace,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := notificationService.SendUsersFreeformNotificationV1Admin(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(sendUsersFreeformNotificationV1AdminCmd)
	sendUsersFreeformNotificationV1AdminCmd.Flags().StringP("body", "by", " ", "Body")
	_ = sendUsersFreeformNotificationV1AdminCmd.MarkFlagRequired("body")
	sendUsersFreeformNotificationV1AdminCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = sendUsersFreeformNotificationV1AdminCmd.MarkFlagRequired("namespace")
}
