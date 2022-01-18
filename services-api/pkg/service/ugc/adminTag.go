// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package ugc

import (
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclient"
	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclient/admin_tag"
	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclientmodels"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type AdminTagService struct {
	Client          *ugcclient.JusticeUgcService
	TokenRepository repository.TokenRepository
}

func (a *AdminTagService) AdminGetTag(input *admin_tag.AdminGetTagParams) (*ugcclientmodels.ModelsPaginatedGetTagResponse, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, notFound, internalServerError, err := a.Client.AdminTag.AdminGetTag(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AdminTagService) AdminCreateTag(input *admin_tag.AdminCreateTagParams) (*ugcclientmodels.ModelsCreateTagResponse, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, unauthorized, conflict, internalServerError, err := a.Client.AdminTag.AdminCreateTag(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if unauthorized != nil {
		return nil, unauthorized
	}
	if conflict != nil {
		return nil, conflict
	}
	if internalServerError != nil {
		return nil, internalServerError
	}
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (a *AdminTagService) AdminUpdateTag(input *admin_tag.AdminUpdateTagParams) (*ugcclientmodels.ModelsCreateTagResponse, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, conflict, internalServerError, err := a.Client.AdminTag.AdminUpdateTag(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if unauthorized != nil {
		return nil, unauthorized
	}
	if notFound != nil {
		return nil, notFound
	}
	if conflict != nil {
		return nil, conflict
	}
	if internalServerError != nil {
		return nil, internalServerError
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AdminTagService) AdminDeleteTag(input *admin_tag.AdminDeleteTagParams) error {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, notFound, internalServerError, err := a.Client.AdminTag.AdminDeleteTag(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return unauthorized
	}
	if notFound != nil {
		return notFound
	}
	if internalServerError != nil {
		return internalServerError
	}
	if err != nil {
		return err
	}
	return nil
}

func (a *AdminTagService) AdminGetTagShort(input *admin_tag.AdminGetTagParams, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsPaginatedGetTagResponse, error) {
	ok, err := a.Client.AdminTag.AdminGetTagShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AdminTagService) AdminCreateTagShort(input *admin_tag.AdminCreateTagParams, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsCreateTagResponse, error) {
	created, err := a.Client.AdminTag.AdminCreateTagShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (a *AdminTagService) AdminUpdateTagShort(input *admin_tag.AdminUpdateTagParams, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsCreateTagResponse, error) {
	ok, err := a.Client.AdminTag.AdminUpdateTagShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AdminTagService) AdminDeleteTagShort(input *admin_tag.AdminDeleteTagParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := a.Client.AdminTag.AdminDeleteTagShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}
