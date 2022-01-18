// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package dsmc

import (
	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclient"
	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclient/public"
	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type PublicService struct {
	Client          *dsmcclient.JusticeDsmcService
	TokenRepository repository.TokenRepository
}

func (p *PublicService) GetDefaultProvider(input *public.GetDefaultProviderParams) (*dsmcclientmodels.ModelsDefaultProvider, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := p.Client.Public.GetDefaultProvider(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicService) ListProviders(input *public.ListProvidersParams) ([]string, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := p.Client.Public.ListProviders(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicService) ListProvidersByRegion(input *public.ListProvidersByRegionParams) (*dsmcclientmodels.ModelsDefaultProvider, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := p.Client.Public.ListProvidersByRegion(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicService) GetDefaultProviderShort(input *public.GetDefaultProviderParams, authInfo runtime.ClientAuthInfoWriter) (*dsmcclientmodels.ModelsDefaultProvider, error) {
	ok, err := p.Client.Public.GetDefaultProviderShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicService) ListProvidersShort(input *public.ListProvidersParams, authInfo runtime.ClientAuthInfoWriter) ([]string, error) {
	ok, err := p.Client.Public.ListProvidersShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicService) ListProvidersByRegionShort(input *public.ListProvidersByRegionParams, authInfo runtime.ClientAuthInfoWriter) (*dsmcclientmodels.ModelsDefaultProvider, error) {
	ok, err := p.Client.Public.ListProvidersByRegionShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}
