// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package gdpr

import (
	"github.com/AccelByte/accelbyte-go-sdk/gdpr-sdk/pkg/gdprclient"
	"github.com/AccelByte/accelbyte-go-sdk/gdpr-sdk/pkg/gdprclient/data_deletion"
	"github.com/AccelByte/accelbyte-go-sdk/gdpr-sdk/pkg/gdprclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type DataDeletionService struct {
	Client          *gdprclient.JusticeGdprService
	TokenRepository repository.TokenRepository
}

func (d *DataDeletionService) AdminGetListDeletionDataRequest(input *data_deletion.AdminGetListDeletionDataRequestParams) (*gdprclientmodels.ModelsListDeletionDataResponse, error) {
	accessToken, err := d.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err := d.Client.DataDeletion.AdminGetListDeletionDataRequest(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
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

func (d *DataDeletionService) AdminGetUserAccountDeletionRequest(input *data_deletion.AdminGetUserAccountDeletionRequestParams) (*gdprclientmodels.ModelsDeletionData, error) {
	accessToken, err := d.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, internalServerError, err := d.Client.DataDeletion.AdminGetUserAccountDeletionRequest(input, client.BearerToken(*accessToken.AccessToken))
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

func (d *DataDeletionService) AdminSubmitUserAccountDeletionRequest(input *data_deletion.AdminSubmitUserAccountDeletionRequestParams) (*gdprclientmodels.ModelsRequestDeleteResponse, error) {
	accessToken, err := d.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, unauthorized, forbidden, notFound, conflict, internalServerError, err := d.Client.DataDeletion.AdminSubmitUserAccountDeletionRequest(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
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
	return created.GetPayload(), nil
}

func (d *DataDeletionService) AdminCancelUserAccountDeletionRequest(input *data_deletion.AdminCancelUserAccountDeletionRequestParams) error {
	accessToken, err := d.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := d.Client.DataDeletion.AdminCancelUserAccountDeletionRequest(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if unauthorized != nil {
		return unauthorized
	}
	if forbidden != nil {
		return forbidden
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

func (d *DataDeletionService) PublicSubmitUserAccountDeletionRequest(input *data_deletion.PublicSubmitUserAccountDeletionRequestParams) (*gdprclientmodels.ModelsRequestDeleteResponse, error) {
	accessToken, err := d.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, unauthorized, forbidden, notFound, internalServerError, err := d.Client.DataDeletion.PublicSubmitUserAccountDeletionRequest(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
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
	return created.GetPayload(), nil
}

func (d *DataDeletionService) PublicCancelUserAccountDeletionRequest(input *data_deletion.PublicCancelUserAccountDeletionRequestParams) error {
	accessToken, err := d.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, forbidden, notFound, internalServerError, err := d.Client.DataDeletion.PublicCancelUserAccountDeletionRequest(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return unauthorized
	}
	if forbidden != nil {
		return forbidden
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

func (d *DataDeletionService) PublicGetUserAccountDeletionStatus(input *data_deletion.PublicGetUserAccountDeletionStatusParams) (*gdprclientmodels.ModelsDeletionStatus, error) {
	accessToken, err := d.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, internalServerError, err := d.Client.DataDeletion.PublicGetUserAccountDeletionStatus(input, client.BearerToken(*accessToken.AccessToken))
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

func (d *DataDeletionService) AdminGetListDeletionDataRequestShort(input *data_deletion.AdminGetListDeletionDataRequestParams, authInfo runtime.ClientAuthInfoWriter) (*gdprclientmodels.ModelsListDeletionDataResponse, error) {
	ok, err := d.Client.DataDeletion.AdminGetListDeletionDataRequestShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (d *DataDeletionService) AdminGetUserAccountDeletionRequestShort(input *data_deletion.AdminGetUserAccountDeletionRequestParams, authInfo runtime.ClientAuthInfoWriter) (*gdprclientmodels.ModelsDeletionData, error) {
	ok, err := d.Client.DataDeletion.AdminGetUserAccountDeletionRequestShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (d *DataDeletionService) AdminSubmitUserAccountDeletionRequestShort(input *data_deletion.AdminSubmitUserAccountDeletionRequestParams, authInfo runtime.ClientAuthInfoWriter) (*gdprclientmodels.ModelsRequestDeleteResponse, error) {
	created, err := d.Client.DataDeletion.AdminSubmitUserAccountDeletionRequestShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (d *DataDeletionService) AdminCancelUserAccountDeletionRequestShort(input *data_deletion.AdminCancelUserAccountDeletionRequestParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := d.Client.DataDeletion.AdminCancelUserAccountDeletionRequestShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (d *DataDeletionService) PublicSubmitUserAccountDeletionRequestShort(input *data_deletion.PublicSubmitUserAccountDeletionRequestParams, authInfo runtime.ClientAuthInfoWriter) (*gdprclientmodels.ModelsRequestDeleteResponse, error) {
	created, err := d.Client.DataDeletion.PublicSubmitUserAccountDeletionRequestShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (d *DataDeletionService) PublicCancelUserAccountDeletionRequestShort(input *data_deletion.PublicCancelUserAccountDeletionRequestParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := d.Client.DataDeletion.PublicCancelUserAccountDeletionRequestShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (d *DataDeletionService) PublicGetUserAccountDeletionStatusShort(input *data_deletion.PublicGetUserAccountDeletionStatusParams, authInfo runtime.ClientAuthInfoWriter) (*gdprclientmodels.ModelsDeletionStatus, error) {
	ok, err := d.Client.DataDeletion.PublicGetUserAccountDeletionStatusShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}
