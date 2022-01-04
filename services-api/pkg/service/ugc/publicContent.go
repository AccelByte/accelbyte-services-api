// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package ugc

import (
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclient"
	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclient/public_content"
	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclientmodels"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type PublicContentService struct {
	Client          *ugcclient.JusticeUgcService
	TokenRepository repository.TokenRepository
}

func (p *PublicContentService) SearchChannelSpecificContent(input *public_content.SearchChannelSpecificContentParams) (*ugcclientmodels.ModelsPaginatedContentDownloadResponse, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, notFound, internalServerError, err := p.Client.PublicContent.SearchChannelSpecificContent(input, client.BearerToken(*accessToken.AccessToken))
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

func (p *PublicContentService) PublicSearchContent(input *public_content.PublicSearchContentParams) (*ugcclientmodels.ModelsPaginatedContentDownloadResponse, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, notFound, internalServerError, err := p.Client.PublicContent.PublicSearchContent(input, client.BearerToken(*accessToken.AccessToken))
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

func (p *PublicContentService) DownloadContentByShareCode(input *public_content.DownloadContentByShareCodeParams) (*ugcclientmodels.ModelsContentDownloadResponse, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, notFound, internalServerError, err := p.Client.PublicContent.DownloadContentByShareCode(input, client.BearerToken(*accessToken.AccessToken))
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

func (p *PublicContentService) PublicDownloadContentByContentID(input *public_content.PublicDownloadContentByContentIDParams) (*ugcclientmodels.ModelsContentDownloadResponse, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, notFound, internalServerError, err := p.Client.PublicContent.PublicDownloadContentByContentID(input, client.BearerToken(*accessToken.AccessToken))
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

func (p *PublicContentService) PublicDownloadContentPreview(input *public_content.PublicDownloadContentPreviewParams) (*ugcclientmodels.ModelsGetContentPreviewResponse, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, notFound, internalServerError, err := p.Client.PublicContent.PublicDownloadContentPreview(input, client.BearerToken(*accessToken.AccessToken))
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

func (p *PublicContentService) CreateContentDirect(input *public_content.CreateContentDirectParams) (*ugcclientmodels.ModelsCreateContentResponse, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, unauthorized, internalServerError, err := p.Client.PublicContent.CreateContentDirect(input, client.BearerToken(*accessToken.AccessToken))
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
	return created.GetPayload(), nil
}

func (p *PublicContentService) CreateContentS3(input *public_content.CreateContentS3Params) (*ugcclientmodels.ModelsCreateContentResponse, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, unauthorized, internalServerError, err := p.Client.PublicContent.CreateContentS3(input, client.BearerToken(*accessToken.AccessToken))
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
	return created.GetPayload(), nil
}

func (p *PublicContentService) UpdateContentS3(input *public_content.UpdateContentS3Params) (*ugcclientmodels.ModelsCreateContentResponse, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, internalServerError, err := p.Client.PublicContent.UpdateContentS3(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
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

func (p *PublicContentService) UpdateContentDirect(input *public_content.UpdateContentDirectParams) (*ugcclientmodels.ModelsCreateContentResponse, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, internalServerError, err := p.Client.PublicContent.UpdateContentDirect(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
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

func (p *PublicContentService) DeleteContent(input *public_content.DeleteContentParams) error {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, notFound, internalServerError, err := p.Client.PublicContent.DeleteContent(input, client.BearerToken(*accessToken.AccessToken))
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

func (p *PublicContentService) PublicGetUserContent(input *public_content.PublicGetUserContentParams) (*ugcclientmodels.ModelsPaginatedContentDownloadResponse, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, notFound, internalServerError, err := p.Client.PublicContent.PublicGetUserContent(input, client.BearerToken(*accessToken.AccessToken))
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

func (p *PublicContentService) UpdateScreenshots(input *public_content.UpdateScreenshotsParams) (*ugcclientmodels.ModelsUpdateScreenshotResponse, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, internalServerError, err := p.Client.PublicContent.UpdateScreenshots(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
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

func (p *PublicContentService) UploadContentScreenshot(input *public_content.UploadContentScreenshotParams) (*ugcclientmodels.ModelsCreateScreenshotResponse, error) {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, unauthorized, internalServerError, err := p.Client.PublicContent.UploadContentScreenshot(input, client.BearerToken(*accessToken.AccessToken))
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
	return created.GetPayload(), nil
}

func (p *PublicContentService) DeleteContentScreenshot(input *public_content.DeleteContentScreenshotParams) error {
	accessToken, err := p.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, notFound, internalServerError, err := p.Client.PublicContent.DeleteContentScreenshot(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
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

func (p *PublicContentService) SearchChannelSpecificContentShort(input *public_content.SearchChannelSpecificContentParams, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsPaginatedContentDownloadResponse, error) {
	ok, err := p.Client.PublicContent.SearchChannelSpecificContentShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicContentService) PublicSearchContentShort(input *public_content.PublicSearchContentParams, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsPaginatedContentDownloadResponse, error) {
	ok, err := p.Client.PublicContent.PublicSearchContentShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicContentService) DownloadContentByShareCodeShort(input *public_content.DownloadContentByShareCodeParams, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsContentDownloadResponse, error) {
	ok, err := p.Client.PublicContent.DownloadContentByShareCodeShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicContentService) PublicDownloadContentByContentIDShort(input *public_content.PublicDownloadContentByContentIDParams, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsContentDownloadResponse, error) {
	ok, err := p.Client.PublicContent.PublicDownloadContentByContentIDShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicContentService) PublicDownloadContentPreviewShort(input *public_content.PublicDownloadContentPreviewParams, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsGetContentPreviewResponse, error) {
	ok, err := p.Client.PublicContent.PublicDownloadContentPreviewShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicContentService) CreateContentDirectShort(input *public_content.CreateContentDirectParams, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsCreateContentResponse, error) {
	created, err := p.Client.PublicContent.CreateContentDirectShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (p *PublicContentService) CreateContentS3Short(input *public_content.CreateContentS3Params, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsCreateContentResponse, error) {
	created, err := p.Client.PublicContent.CreateContentS3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (p *PublicContentService) UpdateContentS3Short(input *public_content.UpdateContentS3Params, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsCreateContentResponse, error) {
	ok, err := p.Client.PublicContent.UpdateContentS3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicContentService) UpdateContentDirectShort(input *public_content.UpdateContentDirectParams, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsCreateContentResponse, error) {
	ok, err := p.Client.PublicContent.UpdateContentDirectShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicContentService) DeleteContentShort(input *public_content.DeleteContentParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := p.Client.PublicContent.DeleteContentShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (p *PublicContentService) PublicGetUserContentShort(input *public_content.PublicGetUserContentParams, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsPaginatedContentDownloadResponse, error) {
	ok, err := p.Client.PublicContent.PublicGetUserContentShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicContentService) UpdateScreenshotsShort(input *public_content.UpdateScreenshotsParams, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsUpdateScreenshotResponse, error) {
	ok, err := p.Client.PublicContent.UpdateScreenshotsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PublicContentService) UploadContentScreenshotShort(input *public_content.UploadContentScreenshotParams, authInfo runtime.ClientAuthInfoWriter) (*ugcclientmodels.ModelsCreateScreenshotResponse, error) {
	created, err := p.Client.PublicContent.UploadContentScreenshotShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (p *PublicContentService) DeleteContentScreenshotShort(input *public_content.DeleteContentScreenshotParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := p.Client.PublicContent.DeleteContentScreenshotShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}
