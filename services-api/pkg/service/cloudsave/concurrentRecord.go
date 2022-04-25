// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package cloudsave

import (
	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclient"
	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclient/concurrent_record"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils/auth"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type ConcurrentRecordService struct {
	Client          *cloudsaveclient.JusticeCloudsaveService
	TokenRepository repository.TokenRepository
}

// Deprecated: Use PutGameRecordConcurrentHandlerV1Short instead
func (c *ConcurrentRecordService) PutGameRecordConcurrentHandlerV1(input *concurrent_record.PutGameRecordConcurrentHandlerV1Params) error {
	token, err := c.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, preconditionFailed, internalServerError, err := c.Client.ConcurrentRecord.PutGameRecordConcurrentHandlerV1(input, client.BearerToken(*token.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if unauthorized != nil {
		return unauthorized
	}
	if preconditionFailed != nil {
		return preconditionFailed
	}
	if internalServerError != nil {
		return internalServerError
	}
	if err != nil {
		return err
	}

	return nil
}

// Deprecated: Use PutPlayerPublicRecordConcurrentHandlerV1Short instead
func (c *ConcurrentRecordService) PutPlayerPublicRecordConcurrentHandlerV1(input *concurrent_record.PutPlayerPublicRecordConcurrentHandlerV1Params) error {
	token, err := c.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, preconditionFailed, internalServerError, err := c.Client.ConcurrentRecord.PutPlayerPublicRecordConcurrentHandlerV1(input, client.BearerToken(*token.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if unauthorized != nil {
		return unauthorized
	}
	if preconditionFailed != nil {
		return preconditionFailed
	}
	if internalServerError != nil {
		return internalServerError
	}
	if err != nil {
		return err
	}

	return nil
}

// [{'HasPermission': ['NAMESPACE:{namespace}:CLOUDSAVE:RECORD [UPDATE]'], 'HasScope': ['social'], 'authorization': []}]
func (c *ConcurrentRecordService) PutGameRecordConcurrentHandlerV1Short(input *concurrent_record.PutGameRecordConcurrentHandlerV1Params, authInfoWriter runtime.ClientAuthInfoWriter) error {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(c.TokenRepository, nil, security, "")
	}
	_, err := c.Client.ConcurrentRecord.PutGameRecordConcurrentHandlerV1Short(input, authInfoWriter)
	if err != nil {
		return err
	}

	return nil
}

// [{'HasPermission': ['NAMESPACE:{namespace}:USER:{userId}:PUBLIC:CLOUDSAVE:RECORD [UPDATE]'], 'HasScope': ['social'], 'authorization': []}]
func (c *ConcurrentRecordService) PutPlayerPublicRecordConcurrentHandlerV1Short(input *concurrent_record.PutPlayerPublicRecordConcurrentHandlerV1Params, authInfoWriter runtime.ClientAuthInfoWriter) error {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(c.TokenRepository, nil, security, "")
	}
	_, err := c.Client.ConcurrentRecord.PutPlayerPublicRecordConcurrentHandlerV1Short(input, authInfoWriter)
	if err != nil {
		return err
	}

	return nil
}
