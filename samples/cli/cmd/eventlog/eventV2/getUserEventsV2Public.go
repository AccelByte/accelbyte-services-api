// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package eventV2

import (
	"github.com/AccelByte/accelbyte-go-sdk/eventlog-sdk/pkg/eventlogclient/event_v2"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/eventlog"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getUserEventsV2PublicCmd represents the getUserEventsV2Public command
var getUserEventsV2PublicCmd = &cobra.Command{
	Use:   "getUserEventsV2Public",
	Short: "Get user events V2 public",
	Long:  `Get user events V2 public`,
	RunE: func(cmd *cobra.Command, args []string) error {
		eventV2Service := &eventlog.EventV2Service{
			Client:          factory.NewEventlogClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		endDate, _ := cmd.Flags().GetString("endDate")
		eventName, _ := cmd.Flags().GetString("eventName")
		offset, _ := cmd.Flags().GetFloat64("offset")
		pageSize, _ := cmd.Flags().GetFloat64("pageSize")
		startDate, _ := cmd.Flags().GetString("startDate")
		input := &event_v2.GetUserEventsV2PublicParams{
			Namespace: namespace,
			UserID:    userId,
			EndDate:   &endDate,
			EventName: &eventName,
			Offset:    &offset,
			PageSize:  &pageSize,
			StartDate: &startDate,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := eventV2Service.GetUserEventsV2Public(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(getUserEventsV2PublicCmd)
	getUserEventsV2PublicCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = getUserEventsV2PublicCmd.MarkFlagRequired("namespace")
	getUserEventsV2PublicCmd.Flags().StringP("userId", "ud", " ", "User id")
	_ = getUserEventsV2PublicCmd.MarkFlagRequired("userId")
	getUserEventsV2PublicCmd.Flags().StringP("endDate", "ee", " ", "End date")
	getUserEventsV2PublicCmd.Flags().StringP("eventName", "ee", " ", "Event name")
	getUserEventsV2PublicCmd.Flags().Float64P("offset", "ot", 0, "Offset")
	getUserEventsV2PublicCmd.Flags().Float64P("pageSize", "pe", 1, "Page size")
	getUserEventsV2PublicCmd.Flags().StringP("startDate", "se", " ", "Start date")
}
