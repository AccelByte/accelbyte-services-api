// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package platform

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/fulfillment_script"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils/auth"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type FulfillmentScriptService struct {
	Client          *platformclient.JusticePlatformService
	TokenRepository repository.TokenRepository
}

// Deprecated: Use ListFulfillmentScriptsShort instead
func (f *FulfillmentScriptService) ListFulfillmentScripts(input *fulfillment_script.ListFulfillmentScriptsParams) ([]*platformclientmodels.FulfillmentScriptInfo, error) {
	token, err := f.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := f.Client.FulfillmentScript.ListFulfillmentScripts(input, client.BearerToken(*token.AccessToken))
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use TestFulfillmentScriptEvalShort instead
func (f *FulfillmentScriptService) TestFulfillmentScriptEval(input *fulfillment_script.TestFulfillmentScriptEvalParams) (*platformclientmodels.FulfillmentScriptEvalTestResult, error) {
	token, err := f.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := f.Client.FulfillmentScript.TestFulfillmentScriptEval(input, client.BearerToken(*token.AccessToken))
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use GetFulfillmentScriptShort instead
func (f *FulfillmentScriptService) GetFulfillmentScript(input *fulfillment_script.GetFulfillmentScriptParams) (*platformclientmodels.FulfillmentScriptInfo, error) {
	token, err := f.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := f.Client.FulfillmentScript.GetFulfillmentScript(input, client.BearerToken(*token.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use CreateFulfillmentScriptShort instead
func (f *FulfillmentScriptService) CreateFulfillmentScript(input *fulfillment_script.CreateFulfillmentScriptParams) (*platformclientmodels.FulfillmentScriptInfo, error) {
	token, err := f.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, conflict, err := f.Client.FulfillmentScript.CreateFulfillmentScript(input, client.BearerToken(*token.AccessToken))
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}

	return created.GetPayload(), nil
}

// Deprecated: Use DeleteFulfillmentScriptShort instead
func (f *FulfillmentScriptService) DeleteFulfillmentScript(input *fulfillment_script.DeleteFulfillmentScriptParams) error {
	token, err := f.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, err = f.Client.FulfillmentScript.DeleteFulfillmentScript(input, client.BearerToken(*token.AccessToken))
	if err != nil {
		return err
	}

	return nil
}

// Deprecated: Use UpdateFulfillmentScriptShort instead
func (f *FulfillmentScriptService) UpdateFulfillmentScript(input *fulfillment_script.UpdateFulfillmentScriptParams) (*platformclientmodels.FulfillmentScriptInfo, error) {
	token, err := f.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, err := f.Client.FulfillmentScript.UpdateFulfillmentScript(input, client.BearerToken(*token.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:FULFILLMENT [READ]'], 'authorization': []}]
func (f *FulfillmentScriptService) ListFulfillmentScriptsShort(input *fulfillment_script.ListFulfillmentScriptsParams, authInfoWriter runtime.ClientAuthInfoWriter) ([]*platformclientmodels.FulfillmentScriptInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(f.TokenRepository, nil, security, "")
	}
	ok, err := f.Client.FulfillmentScript.ListFulfillmentScriptsShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:FULFILLMENT [READ]'], 'authorization': []}]
func (f *FulfillmentScriptService) TestFulfillmentScriptEvalShort(input *fulfillment_script.TestFulfillmentScriptEvalParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.FulfillmentScriptEvalTestResult, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(f.TokenRepository, nil, security, "")
	}
	ok, err := f.Client.FulfillmentScript.TestFulfillmentScriptEvalShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:FULFILLMENT [READ]'], 'authorization': []}]
func (f *FulfillmentScriptService) GetFulfillmentScriptShort(input *fulfillment_script.GetFulfillmentScriptParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.FulfillmentScriptInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(f.TokenRepository, nil, security, "")
	}
	ok, err := f.Client.FulfillmentScript.GetFulfillmentScriptShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:FULFILLMENT [CREATE]'], 'authorization': []}]
func (f *FulfillmentScriptService) CreateFulfillmentScriptShort(input *fulfillment_script.CreateFulfillmentScriptParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.FulfillmentScriptInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(f.TokenRepository, nil, security, "")
	}
	created, err := f.Client.FulfillmentScript.CreateFulfillmentScriptShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return created.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:FULFILLMENT [DELETE]'], 'authorization': []}]
func (f *FulfillmentScriptService) DeleteFulfillmentScriptShort(input *fulfillment_script.DeleteFulfillmentScriptParams, authInfoWriter runtime.ClientAuthInfoWriter) error {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(f.TokenRepository, nil, security, "")
	}
	_, err := f.Client.FulfillmentScript.DeleteFulfillmentScriptShort(input, authInfoWriter)
	if err != nil {
		return err
	}

	return nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:FULFILLMENT [UPDATE]'], 'authorization': []}]
func (f *FulfillmentScriptService) UpdateFulfillmentScriptShort(input *fulfillment_script.UpdateFulfillmentScriptParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.FulfillmentScriptInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(f.TokenRepository, nil, security, "")
	}
	ok, err := f.Client.FulfillmentScript.UpdateFulfillmentScriptShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}
