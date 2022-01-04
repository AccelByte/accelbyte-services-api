// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package cloudsave

import (
	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclient"
	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclient/admin_game_record"
	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type AdminGameRecordService struct {
	Client          *cloudsaveclient.JusticeCloudsaveService
	TokenRepository repository.TokenRepository
}

func (a *AdminGameRecordService) ListGameRecordsHandlerV1(input *admin_game_record.ListGameRecordsHandlerV1Params) (*cloudsaveclientmodels.ModelsListGameRecordKeys, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, internalServerError, err := a.Client.AdminGameRecord.ListGameRecordsHandlerV1(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if internalServerError != nil {
		return nil, internalServerError
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AdminGameRecordService) AdminGetGameRecordHandlerV1(input *admin_game_record.AdminGetGameRecordHandlerV1Params) (*cloudsaveclientmodels.ModelsGameRecord, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, notFound, internalServerError, err := a.Client.AdminGameRecord.AdminGetGameRecordHandlerV1(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if notFound != nil {
		return nil, notFound
	}
	if internalServerError != nil {
		return nil, internalServerError
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AdminGameRecordService) AdminPutGameRecordHandlerV1(input *admin_game_record.AdminPutGameRecordHandlerV1Params) error {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, internalServerError, err := a.Client.AdminGameRecord.AdminPutGameRecordHandlerV1(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return unauthorized
	}
	if internalServerError != nil {
		return internalServerError
	}
	if err != nil {
		return err
	}
	return nil
}

func (a *AdminGameRecordService) AdminPostGameRecordHandlerV1(input *admin_game_record.AdminPostGameRecordHandlerV1Params) error {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, internalServerError, err := a.Client.AdminGameRecord.AdminPostGameRecordHandlerV1(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return unauthorized
	}
	if internalServerError != nil {
		return internalServerError
	}
	if err != nil {
		return err
	}
	return nil
}

func (a *AdminGameRecordService) AdminDeleteGameRecordHandlerV1(input *admin_game_record.AdminDeleteGameRecordHandlerV1Params) error {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, internalServerError, err := a.Client.AdminGameRecord.AdminDeleteGameRecordHandlerV1(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return unauthorized
	}
	if internalServerError != nil {
		return internalServerError
	}
	if err != nil {
		return err
	}
	return nil
}

func (a *AdminGameRecordService) ListGameRecordsHandlerV1Short(input *admin_game_record.ListGameRecordsHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*cloudsaveclientmodels.ModelsListGameRecordKeys, error) {
	ok, err := a.Client.AdminGameRecord.ListGameRecordsHandlerV1Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AdminGameRecordService) AdminGetGameRecordHandlerV1Short(input *admin_game_record.AdminGetGameRecordHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*cloudsaveclientmodels.ModelsGameRecord, error) {
	ok, err := a.Client.AdminGameRecord.AdminGetGameRecordHandlerV1Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AdminGameRecordService) AdminPutGameRecordHandlerV1Short(input *admin_game_record.AdminPutGameRecordHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := a.Client.AdminGameRecord.AdminPutGameRecordHandlerV1Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (a *AdminGameRecordService) AdminPostGameRecordHandlerV1Short(input *admin_game_record.AdminPostGameRecordHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := a.Client.AdminGameRecord.AdminPostGameRecordHandlerV1Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (a *AdminGameRecordService) AdminDeleteGameRecordHandlerV1Short(input *admin_game_record.AdminDeleteGameRecordHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := a.Client.AdminGameRecord.AdminDeleteGameRecordHandlerV1Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}
