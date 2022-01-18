// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package dsmc

import (
	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclient"
	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclient/image_config"
	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type ImageConfigService struct {
	Client          *dsmcclient.JusticeDsmcService
	TokenRepository repository.TokenRepository
}

func (i *ImageConfigService) UpdateImage(input *image_config.UpdateImageParams) error {
	accessToken, err := i.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, internalServerError, err := i.Client.ImageConfig.UpdateImage(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
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

func (i *ImageConfigService) CreateImage(input *image_config.CreateImageParams) error {
	accessToken, err := i.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, conflict, internalServerError, err := i.Client.ImageConfig.CreateImage(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if unauthorized != nil {
		return unauthorized
	}
	if conflict != nil {
		return conflict
	}
	if internalServerError != nil {
		return internalServerError
	}
	if err != nil {
		return err
	}
	return nil
}

func (i *ImageConfigService) ImportImages(input *image_config.ImportImagesParams) (*dsmcclientmodels.ModelsImportResponse, error) {
	accessToken, err := i.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, internalServerError, err := i.Client.ImageConfig.ImportImages(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if internalServerError != nil {
		return nil, internalServerError
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (i *ImageConfigService) ListImages(input *image_config.ListImagesParams) (*dsmcclientmodels.ModelsListImageResponse, error) {
	accessToken, err := i.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, internalServerError, err := i.Client.ImageConfig.ListImages(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
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

func (i *ImageConfigService) DeleteImage(input *image_config.DeleteImageParams) error {
	accessToken, err := i.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, notFound, unprocessableEntity, internalServerError, err := i.Client.ImageConfig.DeleteImage(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if unauthorized != nil {
		return unauthorized
	}
	if notFound != nil {
		return notFound
	}
	if unprocessableEntity != nil {
		return unprocessableEntity
	}
	if internalServerError != nil {
		return internalServerError
	}
	if err != nil {
		return err
	}
	return nil
}

func (i *ImageConfigService) ExportImages(input *image_config.ExportImagesParams) ([]*dsmcclientmodels.ModelsImageRecord, error) {
	accessToken, err := i.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, internalServerError, err := i.Client.ImageConfig.ExportImages(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
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

func (i *ImageConfigService) GetImageLimit(input *image_config.GetImageLimitParams) (*dsmcclientmodels.ModelsGetImageLimitResponse, error) {
	accessToken, err := i.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, internalServerError, err := i.Client.ImageConfig.GetImageLimit(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
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

func (i *ImageConfigService) GetImageDetail(input *image_config.GetImageDetailParams) (*dsmcclientmodels.ModelsGetImageDetailResponse, error) {
	accessToken, err := i.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, notFound, internalServerError, err := i.Client.ImageConfig.GetImageDetail(input, client.BearerToken(*accessToken.AccessToken))
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

func (i *ImageConfigService) ImageDetailClient(input *image_config.ImageDetailClientParams) (*dsmcclientmodels.ModelsGetImageDetailResponse, error) {
	accessToken, err := i.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, notFound, internalServerError, err := i.Client.ImageConfig.ImageDetailClient(input, client.BearerToken(*accessToken.AccessToken))
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

func (i *ImageConfigService) UpdateImageShort(input *image_config.UpdateImageParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := i.Client.ImageConfig.UpdateImageShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (i *ImageConfigService) CreateImageShort(input *image_config.CreateImageParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := i.Client.ImageConfig.CreateImageShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (i *ImageConfigService) ImportImagesShort(input *image_config.ImportImagesParams, authInfo runtime.ClientAuthInfoWriter) (*dsmcclientmodels.ModelsImportResponse, error) {
	ok, err := i.Client.ImageConfig.ImportImagesShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (i *ImageConfigService) ListImagesShort(input *image_config.ListImagesParams, authInfo runtime.ClientAuthInfoWriter) (*dsmcclientmodels.ModelsListImageResponse, error) {
	ok, err := i.Client.ImageConfig.ListImagesShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (i *ImageConfigService) DeleteImageShort(input *image_config.DeleteImageParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := i.Client.ImageConfig.DeleteImageShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (i *ImageConfigService) ExportImagesShort(input *image_config.ExportImagesParams, authInfo runtime.ClientAuthInfoWriter) ([]*dsmcclientmodels.ModelsImageRecord, error) {
	ok, err := i.Client.ImageConfig.ExportImagesShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (i *ImageConfigService) GetImageLimitShort(input *image_config.GetImageLimitParams, authInfo runtime.ClientAuthInfoWriter) (*dsmcclientmodels.ModelsGetImageLimitResponse, error) {
	ok, err := i.Client.ImageConfig.GetImageLimitShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (i *ImageConfigService) GetImageDetailShort(input *image_config.GetImageDetailParams, authInfo runtime.ClientAuthInfoWriter) (*dsmcclientmodels.ModelsGetImageDetailResponse, error) {
	ok, err := i.Client.ImageConfig.GetImageDetailShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (i *ImageConfigService) ImageDetailClientShort(input *image_config.ImageDetailClientParams, authInfo runtime.ClientAuthInfoWriter) (*dsmcclientmodels.ModelsGetImageDetailResponse, error) {
	ok, err := i.Client.ImageConfig.ImageDetailClientShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}
