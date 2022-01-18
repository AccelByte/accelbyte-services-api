// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package achievement

import (
	"github.com/AccelByte/accelbyte-go-sdk/achievement-sdk/pkg/achievementclient"
	"github.com/AccelByte/accelbyte-go-sdk/achievement-sdk/pkg/achievementclient/achievements"
	"github.com/AccelByte/accelbyte-go-sdk/achievement-sdk/pkg/achievementclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type AchievementsService struct {
	Client          *achievementclient.JusticeAchievementService
	TokenRepository repository.TokenRepository
}

func (a *AchievementsService) AdminListAchievements(input *achievements.AdminListAchievementsParams) (*achievementclientmodels.ModelsPaginatedAchievementResponse, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, internalServerError, err := a.Client.Achievements.AdminListAchievements(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) AdminCreateNewAchievement(input *achievements.AdminCreateNewAchievementParams) (*achievementclientmodels.ModelsAchievementResponse, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, unauthorized, internalServerError, err := a.Client.Achievements.AdminCreateNewAchievement(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) ExportAchievements(input *achievements.ExportAchievementsParams) ([]*achievementclientmodels.ModelsAchievement, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, internalServerError, err := a.Client.Achievements.ExportAchievements(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) ImportAchievements(input *achievements.ImportAchievementsParams) (*achievementclientmodels.ServiceImportConfigResponse, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, internalServerError, err := a.Client.Achievements.ImportAchievements(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) AdminGetAchievement(input *achievements.AdminGetAchievementParams) (*achievementclientmodels.ModelsAchievementResponse, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, internalServerError, err := a.Client.Achievements.AdminGetAchievement(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) AdminUpdateAchievement(input *achievements.AdminUpdateAchievementParams) (*achievementclientmodels.ModelsAchievementResponse, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, internalServerError, err := a.Client.Achievements.AdminUpdateAchievement(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) AdminDeleteAchievement(input *achievements.AdminDeleteAchievementParams) error {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, notFound, internalServerError, err := a.Client.Achievements.AdminDeleteAchievement(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) AdminUpdateAchievementListOrder(input *achievements.AdminUpdateAchievementListOrderParams) error {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, notFound, internalServerError, err := a.Client.Achievements.AdminUpdateAchievementListOrder(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) AdminListUserAchievements(input *achievements.AdminListUserAchievementsParams) (*achievementclientmodels.ModelsPaginatedUserAchievementResponse, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, internalServerError, err := a.Client.Achievements.AdminListUserAchievements(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) AdminUnlockAchievement(input *achievements.AdminUnlockAchievementParams) error {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, internalServerError, err := a.Client.Achievements.AdminUnlockAchievement(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) PublicListAchievements(input *achievements.PublicListAchievementsParams) (*achievementclientmodels.ModelsPublicAchievementsResponse, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, internalServerError, err := a.Client.Achievements.PublicListAchievements(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) PublicGetAchievement(input *achievements.PublicGetAchievementParams) (*achievementclientmodels.ModelsPublicAchievementResponse, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, internalServerError, err := a.Client.Achievements.PublicGetAchievement(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) PublicListUserAchievements(input *achievements.PublicListUserAchievementsParams) (*achievementclientmodels.ModelsPaginatedUserAchievementResponse, error) {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, internalServerError, err := a.Client.Achievements.PublicListUserAchievements(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) PublicUnlockAchievement(input *achievements.PublicUnlockAchievementParams) error {
	accessToken, err := a.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, internalServerError, err := a.Client.Achievements.PublicUnlockAchievement(input, client.BearerToken(*accessToken.AccessToken))
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

func (a *AchievementsService) AdminListAchievementsShort(input *achievements.AdminListAchievementsParams, authInfo runtime.ClientAuthInfoWriter) (*achievementclientmodels.ModelsPaginatedAchievementResponse, error) {
	ok, err := a.Client.Achievements.AdminListAchievementsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AchievementsService) AdminCreateNewAchievementShort(input *achievements.AdminCreateNewAchievementParams, authInfo runtime.ClientAuthInfoWriter) (*achievementclientmodels.ModelsAchievementResponse, error) {
	created, err := a.Client.Achievements.AdminCreateNewAchievementShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (a *AchievementsService) ExportAchievementsShort(input *achievements.ExportAchievementsParams, authInfo runtime.ClientAuthInfoWriter) ([]*achievementclientmodels.ModelsAchievement, error) {
	ok, err := a.Client.Achievements.ExportAchievementsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AchievementsService) ImportAchievementsShort(input *achievements.ImportAchievementsParams, authInfo runtime.ClientAuthInfoWriter) (*achievementclientmodels.ServiceImportConfigResponse, error) {
	ok, err := a.Client.Achievements.ImportAchievementsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AchievementsService) AdminGetAchievementShort(input *achievements.AdminGetAchievementParams, authInfo runtime.ClientAuthInfoWriter) (*achievementclientmodels.ModelsAchievementResponse, error) {
	ok, err := a.Client.Achievements.AdminGetAchievementShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AchievementsService) AdminUpdateAchievementShort(input *achievements.AdminUpdateAchievementParams, authInfo runtime.ClientAuthInfoWriter) (*achievementclientmodels.ModelsAchievementResponse, error) {
	ok, err := a.Client.Achievements.AdminUpdateAchievementShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AchievementsService) AdminDeleteAchievementShort(input *achievements.AdminDeleteAchievementParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := a.Client.Achievements.AdminDeleteAchievementShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (a *AchievementsService) AdminUpdateAchievementListOrderShort(input *achievements.AdminUpdateAchievementListOrderParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := a.Client.Achievements.AdminUpdateAchievementListOrderShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (a *AchievementsService) AdminListUserAchievementsShort(input *achievements.AdminListUserAchievementsParams, authInfo runtime.ClientAuthInfoWriter) (*achievementclientmodels.ModelsPaginatedUserAchievementResponse, error) {
	ok, err := a.Client.Achievements.AdminListUserAchievementsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AchievementsService) AdminUnlockAchievementShort(input *achievements.AdminUnlockAchievementParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := a.Client.Achievements.AdminUnlockAchievementShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (a *AchievementsService) PublicListAchievementsShort(input *achievements.PublicListAchievementsParams, authInfo runtime.ClientAuthInfoWriter) (*achievementclientmodels.ModelsPublicAchievementsResponse, error) {
	ok, err := a.Client.Achievements.PublicListAchievementsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AchievementsService) PublicGetAchievementShort(input *achievements.PublicGetAchievementParams, authInfo runtime.ClientAuthInfoWriter) (*achievementclientmodels.ModelsPublicAchievementResponse, error) {
	ok, err := a.Client.Achievements.PublicGetAchievementShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AchievementsService) PublicListUserAchievementsShort(input *achievements.PublicListUserAchievementsParams, authInfo runtime.ClientAuthInfoWriter) (*achievementclientmodels.ModelsPaginatedUserAchievementResponse, error) {
	ok, err := a.Client.Achievements.PublicListUserAchievementsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (a *AchievementsService) PublicUnlockAchievementShort(input *achievements.PublicUnlockAchievementParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := a.Client.Achievements.PublicUnlockAchievementShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}
