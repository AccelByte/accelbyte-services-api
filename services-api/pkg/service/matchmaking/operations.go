// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package matchmaking

import (
	"github.com/AccelByte/accelbyte-go-sdk/matchmaking-sdk/pkg/matchmakingclient"
	"github.com/AccelByte/accelbyte-go-sdk/matchmaking-sdk/pkg/matchmakingclient/operations"
	"github.com/AccelByte/accelbyte-go-sdk/matchmaking-sdk/pkg/matchmakingclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type OperationsService struct {
	Client          *matchmakingclient.JusticeMatchmakingService
	TokenRepository repository.TokenRepository
}

func (o *OperationsService) GetHealthcheckInfo(input *operations.GetHealthcheckInfoParams) error {
	accessToken, err := o.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, err = o.Client.Operations.GetHealthcheckInfo(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return err
	}
	return nil
}

func (o *OperationsService) HandlerV3Healthz(input *operations.HandlerV3HealthzParams) error {
	accessToken, err := o.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, err = o.Client.Operations.HandlerV3Healthz(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return err
	}
	return nil
}

func (o *OperationsService) PublicGetMessages(input *operations.PublicGetMessagesParams) ([]*matchmakingclientmodels.LogAppMessageDeclaration, error) {
	accessToken, err := o.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, internalServerError, err := o.Client.Operations.PublicGetMessages(input, client.BearerToken(*accessToken.AccessToken))
	if internalServerError != nil {
		return nil, internalServerError
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (o *OperationsService) VersionCheckHandler(input *operations.VersionCheckHandlerParams) error {
	accessToken, err := o.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, err = o.Client.Operations.VersionCheckHandler(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return err
	}
	return nil
}

func (o *OperationsService) GetHealthcheckInfoShort(input *operations.GetHealthcheckInfoParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := o.Client.Operations.GetHealthcheckInfoShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (o *OperationsService) HandlerV3HealthzShort(input *operations.HandlerV3HealthzParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := o.Client.Operations.HandlerV3HealthzShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (o *OperationsService) PublicGetMessagesShort(input *operations.PublicGetMessagesParams, authInfo runtime.ClientAuthInfoWriter) ([]*matchmakingclientmodels.LogAppMessageDeclaration, error) {
	ok, err := o.Client.Operations.PublicGetMessagesShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (o *OperationsService) VersionCheckHandlerShort(input *operations.VersionCheckHandlerParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := o.Client.Operations.VersionCheckHandlerShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}
