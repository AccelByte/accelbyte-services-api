// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.
package factory

import (
	"strings"

	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclient"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/sirupsen/logrus"
)

var cloudsaveClientInstance *cloudsaveclient.JusticeCloudsaveService

func NewCloudsaveClient(configRepository repository.ConfigRepository) *cloudsaveclient.JusticeCloudsaveService {
	if cloudsaveClientInstance == nil {
		baseUrl := configRepository.GetJusticeBaseUrl()
		if len(baseUrl) > 0 {
			logrus.Infof("Base URL : %v", baseUrl)
			baseUrlSplit := strings.Split(baseUrl, "://")
			httpClientConfig := &cloudsaveclient.TransportConfig{
				Host:     baseUrlSplit[1],
				BasePath: "",
				Schemes:  []string{baseUrlSplit[0]},
			}
			cloudsaveClientInstance = cloudsaveclient.NewHTTPClientWithConfig(nil, httpClientConfig)
		} else {
			cloudsaveClientInstance = cloudsaveclient.NewHTTPClient(nil)
		}

	}

	return cloudsaveClientInstance
}
