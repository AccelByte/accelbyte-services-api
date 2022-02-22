// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.
package factory

import (
	"strings"

	"github.com/AccelByte/accelbyte-go-sdk/eventlog-sdk/pkg/eventlogclient"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/sirupsen/logrus"
)

var eventlogClientInstance *eventlogclient.JusticeEventlogService

func NewEventlogClient(configRepository repository.ConfigRepository) *eventlogclient.JusticeEventlogService {
	if eventlogClientInstance == nil {
		baseUrl := configRepository.GetJusticeBaseUrl()
		if len(baseUrl) > 0 {
			logrus.Infof("Base URL : %v", baseUrl)
			baseUrlSplit := strings.Split(baseUrl, "://")
			httpClientConfig := &eventlogclient.TransportConfig{
				Host:     baseUrlSplit[1],
				BasePath: "",
				Schemes:  []string{baseUrlSplit[0]},
			}
			eventlogClientInstance = eventlogclient.NewHTTPClientWithConfig(nil, httpClientConfig)
		} else {
			eventlogClientInstance = eventlogclient.NewHTTPClient(nil)
		}

	}

	return eventlogClientInstance
}
