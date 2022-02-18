// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package event

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/eventlog-sdk/pkg/eventlogclient/event"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/eventlog"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// GetEventByNamespaceHandlerCmd represents the GetEventByNamespaceHandler command
var GetEventByNamespaceHandlerCmd = &cobra.Command{
	Use:   "getEventByNamespaceHandler",
	Short: "Get event by namespace handler",
	Long:  `Get event by namespace handler`,
	RunE: func(cmd *cobra.Command, args []string) error {
		eventService := &eventlog.EventService{
			Client:          factory.NewEventlogClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		endDate, _ := cmd.Flags().GetString("endDate")
		pageSize, _ := cmd.Flags().GetFloat64("pageSize")
		startDate, _ := cmd.Flags().GetString("startDate")
		offset, _ := cmd.Flags().GetFloat64("offset")
		input := &event.GetEventByNamespaceHandlerParams{
			Namespace: namespace,
			Offset:    &offset,
			EndDate:   endDate,
			PageSize:  pageSize,
			StartDate: startDate,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := eventService.GetEventByNamespaceHandler(input)
		if err != nil {
			logrus.Error(err)
			return err
		} else {
			response, errIndent := json.MarshalIndent(ok, "", "    ")
			if errIndent != nil {
				return errIndent
			} else {
				logrus.Infof("Response %s", string(response))
			}
		}
		return nil
	},
}

func init() {
	GetEventByNamespaceHandlerCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = GetEventByNamespaceHandlerCmd.MarkFlagRequired("namespace")
	GetEventByNamespaceHandlerCmd.Flags().Float64P("offset", "", 0, "Offset")
	GetEventByNamespaceHandlerCmd.Flags().StringP("endDate", "", "", "End date")
	_ = GetEventByNamespaceHandlerCmd.MarkFlagRequired("endDate")
	GetEventByNamespaceHandlerCmd.Flags().Float64P("pageSize", "", 1, "Page size")
	_ = GetEventByNamespaceHandlerCmd.MarkFlagRequired("pageSize")
	GetEventByNamespaceHandlerCmd.Flags().StringP("startDate", "", "", "Start date")
	_ = GetEventByNamespaceHandlerCmd.MarkFlagRequired("startDate")
}
