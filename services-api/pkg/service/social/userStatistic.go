// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package social

import (
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclient"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclient/user_statistic"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type UserStatisticService struct {
	Client          *socialclient.JusticeSocialService
	TokenRepository repository.TokenRepository
}

func (u *UserStatisticService) BulkFetchStatItems(input *user_statistic.BulkFetchStatItemsParams) ([]*socialclientmodels.UserStatItemInfo, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkFetchStatItems(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkIncUserStatItem(input *user_statistic.BulkIncUserStatItemParams) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkIncUserStatItem(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkIncUserStatItemValue(input *user_statistic.BulkIncUserStatItemValueParams) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkIncUserStatItemValue(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkResetUserStatItem(input *user_statistic.BulkResetUserStatItemParams) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkResetUserStatItem(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) GetUserStatItems(input *user_statistic.GetUserStatItemsParams) (*socialclientmodels.UserStatItemPagingSlicedResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := u.Client.UserStatistic.GetUserStatItems(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkCreateUserStatItems(input *user_statistic.BulkCreateUserStatItemsParams) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkCreateUserStatItems(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkIncUserStatItem1(input *user_statistic.BulkIncUserStatItem1Params) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkIncUserStatItem1(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkIncUserStatItemValue1(input *user_statistic.BulkIncUserStatItemValue1Params) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkIncUserStatItemValue1(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkResetUserStatItem1(input *user_statistic.BulkResetUserStatItem1Params) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkResetUserStatItem1(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) CreateUserStatItem(input *user_statistic.CreateUserStatItemParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, notFound, conflict, err := u.Client.UserStatistic.CreateUserStatItem(input, client.BearerToken(*accessToken.AccessToken))
	if notFound != nil {
		return notFound
	}
	if conflict != nil {
		return conflict
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UserStatisticService) DeleteUserStatItems(input *user_statistic.DeleteUserStatItemsParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, forbidden, notFound, err := u.Client.UserStatistic.DeleteUserStatItems(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return unauthorized
	}
	if forbidden != nil {
		return forbidden
	}
	if notFound != nil {
		return notFound
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UserStatisticService) IncUserStatItemValue(input *user_statistic.IncUserStatItemValueParams) (*socialclientmodels.StatItemIncResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, conflict, err := u.Client.UserStatistic.IncUserStatItemValue(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) ResetUserStatItemValue(input *user_statistic.ResetUserStatItemValueParams) (*socialclientmodels.StatItemIncResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := u.Client.UserStatistic.ResetUserStatItemValue(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkFetchStatItems1(input *user_statistic.BulkFetchStatItems1Params) ([]*socialclientmodels.UserStatItemInfo, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkFetchStatItems1(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicBulkIncUserStatItem(input *user_statistic.PublicBulkIncUserStatItemParams) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.PublicBulkIncUserStatItem(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicBulkIncUserStatItemValue(input *user_statistic.PublicBulkIncUserStatItemValueParams) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.PublicBulkIncUserStatItemValue(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkResetUserStatItem2(input *user_statistic.BulkResetUserStatItem2Params) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkResetUserStatItem2(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicQueryUserStatItems(input *user_statistic.PublicQueryUserStatItemsParams) (*socialclientmodels.UserStatItemPagingSlicedResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := u.Client.UserStatistic.PublicQueryUserStatItems(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicBulkCreateUserStatItems(input *user_statistic.PublicBulkCreateUserStatItemsParams) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.PublicBulkCreateUserStatItems(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicBulkIncUserStatItem1(input *user_statistic.PublicBulkIncUserStatItem1Params) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.PublicBulkIncUserStatItem1(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkIncUserStatItemValue2(input *user_statistic.BulkIncUserStatItemValue2Params) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkIncUserStatItemValue2(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkResetUserStatItem3(input *user_statistic.BulkResetUserStatItem3Params) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkResetUserStatItem3(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicCreateUserStatItem(input *user_statistic.PublicCreateUserStatItemParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, notFound, conflict, err := u.Client.UserStatistic.PublicCreateUserStatItem(input, client.BearerToken(*accessToken.AccessToken))
	if notFound != nil {
		return notFound
	}
	if conflict != nil {
		return conflict
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UserStatisticService) DeleteUserStatItems1(input *user_statistic.DeleteUserStatItems1Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, forbidden, notFound, err := u.Client.UserStatistic.DeleteUserStatItems1(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return unauthorized
	}
	if forbidden != nil {
		return forbidden
	}
	if notFound != nil {
		return notFound
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UserStatisticService) PublicIncUserStatItem(input *user_statistic.PublicIncUserStatItemParams) (*socialclientmodels.StatItemIncResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, conflict, err := u.Client.UserStatistic.PublicIncUserStatItem(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicIncUserStatItemValue(input *user_statistic.PublicIncUserStatItemValueParams) (*socialclientmodels.StatItemIncResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, conflict, err := u.Client.UserStatistic.PublicIncUserStatItemValue(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) ResetUserStatItemValue1(input *user_statistic.ResetUserStatItemValue1Params) (*socialclientmodels.StatItemIncResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := u.Client.UserStatistic.ResetUserStatItemValue1(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkUpdateUserStatItemV2(input *user_statistic.BulkUpdateUserStatItemV2Params) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkUpdateUserStatItemV2(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkUpdateUserStatItem(input *user_statistic.BulkUpdateUserStatItemParams) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkUpdateUserStatItem(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) DeleteUserStatItems2(input *user_statistic.DeleteUserStatItems2Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, forbidden, notFound, err := u.Client.UserStatistic.DeleteUserStatItems2(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return unauthorized
	}
	if forbidden != nil {
		return forbidden
	}
	if notFound != nil {
		return notFound
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UserStatisticService) UpdateUserStatItemValue(input *user_statistic.UpdateUserStatItemValueParams) (*socialclientmodels.StatItemIncResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, conflict, unprocessableEntity, err := u.Client.UserStatistic.UpdateUserStatItemValue(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if conflict != nil {
		return nil, conflict
	}
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkUpdateUserStatItem1(input *user_statistic.BulkUpdateUserStatItem1Params) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkUpdateUserStatItem1(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkUpdateUserStatItem2(input *user_statistic.BulkUpdateUserStatItem2Params) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := u.Client.UserStatistic.BulkUpdateUserStatItem2(input, client.BearerToken(*accessToken.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) UpdateUserStatItemValue1(input *user_statistic.UpdateUserStatItemValue1Params) (*socialclientmodels.StatItemIncResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, conflict, unprocessableEntity, err := u.Client.UserStatistic.UpdateUserStatItemValue1(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if conflict != nil {
		return nil, conflict
	}
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkFetchStatItemsShort(input *user_statistic.BulkFetchStatItemsParams, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.UserStatItemInfo, error) {
	ok, err := u.Client.UserStatistic.BulkFetchStatItemsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkIncUserStatItemShort(input *user_statistic.BulkIncUserStatItemParams, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkIncUserStatItemShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkIncUserStatItemValueShort(input *user_statistic.BulkIncUserStatItemValueParams, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkIncUserStatItemValueShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkResetUserStatItemShort(input *user_statistic.BulkResetUserStatItemParams, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkResetUserStatItemShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) GetUserStatItemsShort(input *user_statistic.GetUserStatItemsParams, authInfo runtime.ClientAuthInfoWriter) (*socialclientmodels.UserStatItemPagingSlicedResult, error) {
	ok, err := u.Client.UserStatistic.GetUserStatItemsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkCreateUserStatItemsShort(input *user_statistic.BulkCreateUserStatItemsParams, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkCreateUserStatItemsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkIncUserStatItem1Short(input *user_statistic.BulkIncUserStatItem1Params, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkIncUserStatItem1Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkIncUserStatItemValue1Short(input *user_statistic.BulkIncUserStatItemValue1Params, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkIncUserStatItemValue1Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkResetUserStatItem1Short(input *user_statistic.BulkResetUserStatItem1Params, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkResetUserStatItem1Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) CreateUserStatItemShort(input *user_statistic.CreateUserStatItemParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.UserStatistic.CreateUserStatItemShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserStatisticService) DeleteUserStatItemsShort(input *user_statistic.DeleteUserStatItemsParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.UserStatistic.DeleteUserStatItemsShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserStatisticService) IncUserStatItemValueShort(input *user_statistic.IncUserStatItemValueParams, authInfo runtime.ClientAuthInfoWriter) (*socialclientmodels.StatItemIncResult, error) {
	ok, err := u.Client.UserStatistic.IncUserStatItemValueShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) ResetUserStatItemValueShort(input *user_statistic.ResetUserStatItemValueParams, authInfo runtime.ClientAuthInfoWriter) (*socialclientmodels.StatItemIncResult, error) {
	ok, err := u.Client.UserStatistic.ResetUserStatItemValueShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkFetchStatItems1Short(input *user_statistic.BulkFetchStatItems1Params, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.UserStatItemInfo, error) {
	ok, err := u.Client.UserStatistic.BulkFetchStatItems1Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicBulkIncUserStatItemShort(input *user_statistic.PublicBulkIncUserStatItemParams, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.PublicBulkIncUserStatItemShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicBulkIncUserStatItemValueShort(input *user_statistic.PublicBulkIncUserStatItemValueParams, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.PublicBulkIncUserStatItemValueShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkResetUserStatItem2Short(input *user_statistic.BulkResetUserStatItem2Params, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkResetUserStatItem2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicQueryUserStatItemsShort(input *user_statistic.PublicQueryUserStatItemsParams, authInfo runtime.ClientAuthInfoWriter) (*socialclientmodels.UserStatItemPagingSlicedResult, error) {
	ok, err := u.Client.UserStatistic.PublicQueryUserStatItemsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicBulkCreateUserStatItemsShort(input *user_statistic.PublicBulkCreateUserStatItemsParams, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.PublicBulkCreateUserStatItemsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicBulkIncUserStatItem1Short(input *user_statistic.PublicBulkIncUserStatItem1Params, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.PublicBulkIncUserStatItem1Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkIncUserStatItemValue2Short(input *user_statistic.BulkIncUserStatItemValue2Params, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkIncUserStatItemValue2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkResetUserStatItem3Short(input *user_statistic.BulkResetUserStatItem3Params, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkResetUserStatItem3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicCreateUserStatItemShort(input *user_statistic.PublicCreateUserStatItemParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.UserStatistic.PublicCreateUserStatItemShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserStatisticService) DeleteUserStatItems1Short(input *user_statistic.DeleteUserStatItems1Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.UserStatistic.DeleteUserStatItems1Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserStatisticService) PublicIncUserStatItemShort(input *user_statistic.PublicIncUserStatItemParams, authInfo runtime.ClientAuthInfoWriter) (*socialclientmodels.StatItemIncResult, error) {
	ok, err := u.Client.UserStatistic.PublicIncUserStatItemShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) PublicIncUserStatItemValueShort(input *user_statistic.PublicIncUserStatItemValueParams, authInfo runtime.ClientAuthInfoWriter) (*socialclientmodels.StatItemIncResult, error) {
	ok, err := u.Client.UserStatistic.PublicIncUserStatItemValueShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) ResetUserStatItemValue1Short(input *user_statistic.ResetUserStatItemValue1Params, authInfo runtime.ClientAuthInfoWriter) (*socialclientmodels.StatItemIncResult, error) {
	ok, err := u.Client.UserStatistic.ResetUserStatItemValue1Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkUpdateUserStatItemV2Short(input *user_statistic.BulkUpdateUserStatItemV2Params, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkUpdateUserStatItemV2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkUpdateUserStatItemShort(input *user_statistic.BulkUpdateUserStatItemParams, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkUpdateUserStatItemShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) DeleteUserStatItems2Short(input *user_statistic.DeleteUserStatItems2Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.UserStatistic.DeleteUserStatItems2Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserStatisticService) UpdateUserStatItemValueShort(input *user_statistic.UpdateUserStatItemValueParams, authInfo runtime.ClientAuthInfoWriter) (*socialclientmodels.StatItemIncResult, error) {
	ok, err := u.Client.UserStatistic.UpdateUserStatItemValueShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkUpdateUserStatItem1Short(input *user_statistic.BulkUpdateUserStatItem1Params, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkUpdateUserStatItem1Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) BulkUpdateUserStatItem2Short(input *user_statistic.BulkUpdateUserStatItem2Params, authInfo runtime.ClientAuthInfoWriter) ([]*socialclientmodels.BulkStatItemOperationResult, error) {
	ok, err := u.Client.UserStatistic.BulkUpdateUserStatItem2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UserStatisticService) UpdateUserStatItemValue1Short(input *user_statistic.UpdateUserStatItemValue1Params, authInfo runtime.ClientAuthInfoWriter) (*socialclientmodels.StatItemIncResult, error) {
	ok, err := u.Client.UserStatistic.UpdateUserStatItemValue1Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}
