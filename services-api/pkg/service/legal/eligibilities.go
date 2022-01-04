// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package legal

import (
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclient/eligibilities"
	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type EligibilitiesService struct {
	Client          *legalclient.JusticeLegalService
	TokenRepository repository.TokenRepository
}

func (e *EligibilitiesService) RetrieveEligibilitiesPublic(input *eligibilities.RetrieveEligibilitiesPublicParams) ([]*legalclientmodels.RetrieveUserEligibilitiesResponse, error) {
	accessToken, err := e.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := e.Client.Eligibilities.RetrieveEligibilitiesPublic(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (e *EligibilitiesService) RetrieveEligibilitiesPublicIndirect(input *eligibilities.RetrieveEligibilitiesPublicIndirectParams) (*legalclientmodels.RetrieveUserEligibilitiesIndirectResponse, error) {
	accessToken, err := e.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := e.Client.Eligibilities.RetrieveEligibilitiesPublicIndirect(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (e *EligibilitiesService) RetrieveEligibilitiesPublicShort(input *eligibilities.RetrieveEligibilitiesPublicParams, authInfo runtime.ClientAuthInfoWriter) ([]*legalclientmodels.RetrieveUserEligibilitiesResponse, error) {
	ok, err := e.Client.Eligibilities.RetrieveEligibilitiesPublicShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (e *EligibilitiesService) RetrieveEligibilitiesPublicIndirectShort(input *eligibilities.RetrieveEligibilitiesPublicIndirectParams, authInfo runtime.ClientAuthInfoWriter) (*legalclientmodels.RetrieveUserEligibilitiesIndirectResponse, error) {
	ok, err := e.Client.Eligibilities.RetrieveEligibilitiesPublicIndirectShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}
