// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package iam

import (
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient"
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/users"
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type UsersService struct {
	Client          *iamclient.JusticeIamService
	TokenRepository repository.TokenRepository
}

func (u *UsersService) CreateUser(input *users.CreateUserParams) (*iamclientmodels.ModelUserCreateResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, unauthorized, forbidden, conflict, err := u.Client.Users.CreateUser(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) GetAdminUsersByRoleID(input *users.GetAdminUsersByRoleIDParams) (*iamclientmodels.ModelGetAdminUsersResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.GetAdminUsersByRoleID(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) GetUserByLoginID(input *users.GetUserByLoginIDParams) (*iamclientmodels.ModelPublicUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, internalServerError, err := u.Client.Users.GetUserByLoginID(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
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

func (u *UsersService) GetUserByPlatformUserID(input *users.GetUserByPlatformUserIDParams) (*iamclientmodels.ModelPublicUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.GetUserByPlatformUserID(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) ForgotPassword(input *users.ForgotPasswordParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.ForgotPassword(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) GetUsersByLoginIds(input *users.GetUsersByLoginIdsParams) (*iamclientmodels.ModelPublicUsersResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, err := u.Client.Users.GetUsersByLoginIds(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) ResetPassword(input *users.ResetPasswordParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, forbidden, notFound, internalServerError, err := u.Client.Users.ResetPassword(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
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

func (u *UsersService) SearchUser(input *users.SearchUserParams) (*iamclientmodels.ModelSearchUsersResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, err := u.Client.Users.SearchUser(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) GetUserByUserID(input *users.GetUserByUserIDParams) (*iamclientmodels.ModelUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, internalServerError, err := u.Client.Users.GetUserByUserID(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) UpdateUser(input *users.UpdateUserParams) (*iamclientmodels.ModelUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, conflict, internalServerError, err := u.Client.Users.UpdateUser(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) DeleteUser(input *users.DeleteUserParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, forbidden, notFound, err := u.Client.Users.DeleteUser(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) BanUser(input *users.BanUserParams) (*iamclientmodels.ModelUserBanResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.BanUser(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) GetUserBanHistory(input *users.GetUserBanHistoryParams) ([]*iamclientmodels.ModelUserBanResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, err := u.Client.Users.GetUserBanHistory(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) DisableUserBan(input *users.DisableUserBanParams) (*iamclientmodels.ModelUserBanResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.DisableUserBan(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) EnableUserBan(input *users.EnableUserBanParams) (*iamclientmodels.ModelUserBanResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.EnableUserBan(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) ListCrossNamespaceAccountLink(input *users.ListCrossNamespaceAccountLinkParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.ListCrossNamespaceAccountLink(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) DisableUser(input *users.DisableUserParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.DisableUser(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) EnableUser(input *users.EnableUserParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.EnableUser(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) GetUserInformation(input *users.GetUserInformationParams) (*iamclientmodels.ModelUserInformation, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, err := u.Client.Users.GetUserInformation(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) DeleteUserInformation(input *users.DeleteUserInformationParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, forbidden, notFound, err := u.Client.Users.DeleteUserInformation(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) GetUserLoginHistories(input *users.GetUserLoginHistoriesParams) (*iamclientmodels.ModelLoginHistoriesResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, err := u.Client.Users.GetUserLoginHistories(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) UpdatePassword(input *users.UpdatePasswordParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.UpdatePassword(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) SaveUserPermission(input *users.SaveUserPermissionParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.SaveUserPermission(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AddUserPermission(input *users.AddUserPermissionParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.AddUserPermission(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) DeleteUserPermission(input *users.DeleteUserPermissionParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.DeleteUserPermission(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) GetUserPlatformAccounts(input *users.GetUserPlatformAccountsParams) ([]*iamclientmodels.AccountcommonUserLinkedPlatform, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, err := u.Client.Users.GetUserPlatformAccounts(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) GetUserMapping(input *users.GetUserMappingParams) (*iamclientmodels.ModelGetUserMapping, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.GetUserMapping(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) GetUserJusticePlatformAccount(input *users.GetUserJusticePlatformAccountParams) (*iamclientmodels.ModelGetUserJusticePlatformAccountResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, err := u.Client.Users.GetUserJusticePlatformAccount(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PlatformLink(input *users.PlatformLinkParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, conflict, internalServerError, err := u.Client.Users.PlatformLink(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PlatformUnlink(input *users.PlatformUnlinkParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.PlatformUnlink(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) GetPublisherUser(input *users.GetPublisherUserParams) (*iamclientmodels.ModelGetPublisherUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.GetPublisherUser(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) SaveUserRoles(input *users.SaveUserRolesParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, conflict, err := u.Client.Users.SaveUserRoles(input, client.BearerToken(*accessToken.AccessToken))
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
	if conflict != nil {
		return conflict
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AddUserRole(input *users.AddUserRoleParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, forbidden, notFound, conflict, internalServerError, err := u.Client.Users.AddUserRole(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return unauthorized
	}
	if forbidden != nil {
		return forbidden
	}
	if notFound != nil {
		return notFound
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

func (u *UsersService) DeleteUserRole(input *users.DeleteUserRoleParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.DeleteUserRole(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) UpgradeHeadlessAccount(input *users.UpgradeHeadlessAccountParams) (*iamclientmodels.ModelUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, conflict, err := u.Client.Users.UpgradeHeadlessAccount(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) UpgradeHeadlessAccountWithVerificationCode(input *users.UpgradeHeadlessAccountWithVerificationCodeParams) (*iamclientmodels.ModelUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, conflict, err := u.Client.Users.UpgradeHeadlessAccountWithVerificationCode(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) UserVerification(input *users.UserVerificationParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.UserVerification(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) SendVerificationCode(input *users.SendVerificationCodeParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, conflict, tooManyRequests, internalServerError, err := u.Client.Users.SendVerificationCode(input, client.BearerToken(*accessToken.AccessToken))
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
	if conflict != nil {
		return conflict
	}
	if tooManyRequests != nil {
		return tooManyRequests
	}
	if internalServerError != nil {
		return internalServerError
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminGetAgeRestrictionStatusV2(input *users.AdminGetAgeRestrictionStatusV2Params) (*iamclientmodels.ModelAgeRestrictionResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, err := u.Client.Users.AdminGetAgeRestrictionStatusV2(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminUpdateAgeRestrictionConfigV2(input *users.AdminUpdateAgeRestrictionConfigV2Params) (*iamclientmodels.ModelAgeRestrictionResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.AdminUpdateAgeRestrictionConfigV2(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) GetListCountryAgeRestriction(input *users.GetListCountryAgeRestrictionParams) ([]*iamclientmodels.AccountcommonCountryAgeRestriction, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, err := u.Client.Users.GetListCountryAgeRestriction(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) UpdateCountryAgeRestriction(input *users.UpdateCountryAgeRestrictionParams) (*iamclientmodels.ModelCountry, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.UpdateCountryAgeRestriction(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminSearchUsersV2(input *users.AdminSearchUsersV2Params) (*iamclientmodels.ModelSearchUsersByPlatformIDResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, err := u.Client.Users.AdminSearchUsersV2(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminGetUserByUserIDV2(input *users.AdminGetUserByUserIDV2Params) (*iamclientmodels.ModelUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, internalServerError, err := u.Client.Users.AdminGetUserByUserIDV2(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminUpdateUserV2(input *users.AdminUpdateUserV2Params) (*iamclientmodels.ModelUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, conflict, internalServerError, err := u.Client.Users.AdminUpdateUserV2(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminBanUserV2(input *users.AdminBanUserV2Params) (*iamclientmodels.ModelUserBanResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminBanUserV2(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminGetUserBanV2(input *users.AdminGetUserBanV2Params) ([]*iamclientmodels.ModelUserBanResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, err := u.Client.Users.AdminGetUserBanV2(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminDisableUserV2(input *users.AdminDisableUserV2Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminDisableUserV2(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminEnableUserV2(input *users.AdminEnableUserV2Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminEnableUserV2(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminResetPasswordV2(input *users.AdminResetPasswordV2Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminResetPasswordV2(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminDeletePlatformLinkV2(input *users.AdminDeletePlatformLinkV2Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminDeletePlatformLinkV2(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminPutUserRolesV2(input *users.AdminPutUserRolesV2Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.AdminPutUserRolesV2(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminCreateUserRolesV2(input *users.AdminCreateUserRolesV2Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, conflict, err := u.Client.Users.AdminCreateUserRolesV2(input, client.BearerToken(*accessToken.AccessToken))
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
	if conflict != nil {
		return conflict
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicCreateUserV2(input *users.PublicCreateUserV2Params) (*iamclientmodels.ModelUserCreateResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, unauthorized, forbidden, conflict, err := u.Client.Users.PublicCreateUserV2(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) PublicForgotPasswordV2(input *users.PublicForgotPasswordV2Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, notFound, tooManyRequests, internalServerError, err := u.Client.Users.PublicForgotPasswordV2(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if notFound != nil {
		return notFound
	}
	if tooManyRequests != nil {
		return tooManyRequests
	}
	if internalServerError != nil {
		return internalServerError
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicResetPasswordV2(input *users.PublicResetPasswordV2Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, forbidden, notFound, internalServerError, err := u.Client.Users.PublicResetPasswordV2(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
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

func (u *UsersService) PublicGetUserByUserIDV2(input *users.PublicGetUserByUserIDV2Params) (*iamclientmodels.ModelUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, internalServerError, err := u.Client.Users.PublicGetUserByUserIDV2(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicUpdateUserV2(input *users.PublicUpdateUserV2Params) ([]*iamclientmodels.ModelUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, conflict, internalServerError, err := u.Client.Users.PublicUpdateUserV2(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicGetUserBan(input *users.PublicGetUserBanParams) ([]*iamclientmodels.ModelUserBanResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, err := u.Client.Users.PublicGetUserBan(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicUpdatePasswordV2(input *users.PublicUpdatePasswordV2Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.PublicUpdatePasswordV2(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) GetListJusticePlatformAccounts(input *users.GetListJusticePlatformAccountsParams) ([]*iamclientmodels.ModelGetUserMapping, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := u.Client.Users.GetListJusticePlatformAccounts(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicPlatformLinkV2(input *users.PublicPlatformLinkV2Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, conflict, internalServerError, err := u.Client.Users.PublicPlatformLinkV2(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicDeletePlatformLinkV2(input *users.PublicDeletePlatformLinkV2Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.PublicDeletePlatformLinkV2(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) ListAdminsV3(input *users.ListAdminsV3Params) (*iamclientmodels.ModelGetUsersResponseWithPaginationV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, internalServerError, err := u.Client.Users.ListAdminsV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminGetAgeRestrictionStatusV3(input *users.AdminGetAgeRestrictionStatusV3Params) (*iamclientmodels.ModelAgeRestrictionResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminGetAgeRestrictionStatusV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminUpdateAgeRestrictionConfigV3(input *users.AdminUpdateAgeRestrictionConfigV3Params) (*iamclientmodels.ModelAgeRestrictionResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, internalServerError, err := u.Client.Users.AdminUpdateAgeRestrictionConfigV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminGetListCountryAgeRestrictionV3(input *users.AdminGetListCountryAgeRestrictionV3Params) ([]*iamclientmodels.ModelCountryV3Response, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, internalServerError, err := u.Client.Users.AdminGetListCountryAgeRestrictionV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminUpdateCountryAgeRestrictionV3(input *users.AdminUpdateCountryAgeRestrictionV3Params) (*iamclientmodels.ModelCountryV3Response, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminUpdateCountryAgeRestrictionV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminGetUserByPlatformUserIDV3(input *users.AdminGetUserByPlatformUserIDV3Params) (*iamclientmodels.ModelUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminGetUserByPlatformUserIDV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) GetAdminUsersByRoleIDV3(input *users.GetAdminUsersByRoleIDV3Params) (*iamclientmodels.ModelGetUsersResponseWithPaginationV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.GetAdminUsersByRoleIDV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminGetUserByEmailAddressV3(input *users.AdminGetUserByEmailAddressV3Params) (*iamclientmodels.ModelUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminGetUserByEmailAddressV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminListUserIDByUserIDsV3(input *users.AdminListUserIDByUserIDsV3Params) (*iamclientmodels.ModelListUserInformationResult, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, internalServerError, err := u.Client.Users.AdminListUserIDByUserIDsV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminInviteUserV3(input *users.AdminInviteUserV3Params) (*iamclientmodels.ModelInviteUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, notFound, conflict, unprocessableEntity, internalServerError, err := u.Client.Users.AdminInviteUserV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if internalServerError != nil {
		return nil, internalServerError
	}
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) AdminListUsersV3(input *users.AdminListUsersV3Params) (*iamclientmodels.AccountcommonListUsersWithPlatformAccountsResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, internalServerError, err := u.Client.Users.AdminListUsersV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminSearchUserV3(input *users.AdminSearchUserV3Params) (*iamclientmodels.ModelSearchUsersResponseWithPaginationV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, internalServerError, err := u.Client.Users.AdminSearchUserV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminGetBulkUserByEmailAddressV3(input *users.AdminGetBulkUserByEmailAddressV3Params) (*iamclientmodels.ModelListUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminGetBulkUserByEmailAddressV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminGetUserByUserIDV3(input *users.AdminGetUserByUserIDV3Params) (*iamclientmodels.ModelUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminGetUserByUserIDV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminUpdateUserV3(input *users.AdminUpdateUserV3Params) (*iamclientmodels.ModelUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, conflict, internalServerError, err := u.Client.Users.AdminUpdateUserV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminGetUserBanV3(input *users.AdminGetUserBanV3Params) (*iamclientmodels.ModelGetUserBanV3Response, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.AdminGetUserBanV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminBanUserV3(input *users.AdminBanUserV3Params) (*iamclientmodels.ModelUserBanResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.AdminBanUserV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) AdminUpdateUserBanV3(input *users.AdminUpdateUserBanV3Params) (*iamclientmodels.ModelUserBanResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminUpdateUserBanV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminSendVerificationCodeV3(input *users.AdminSendVerificationCodeV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, conflict, tooManyRequests, err := u.Client.Users.AdminSendVerificationCodeV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if conflict != nil {
		return conflict
	}
	if tooManyRequests != nil {
		return tooManyRequests
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminVerifyAccountV3(input *users.AdminVerifyAccountV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminVerifyAccountV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) GetUserVerificationCode(input *users.GetUserVerificationCodeParams) (*iamclientmodels.ModelVerificationCodeResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.GetUserVerificationCode(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminGetUserDeletionStatusV3(input *users.AdminGetUserDeletionStatusV3Params) (*iamclientmodels.ModelUserDeletionStatusResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminGetUserDeletionStatusV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminUpdateUserDeletionStatusV3(input *users.AdminUpdateUserDeletionStatusV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminUpdateUserDeletionStatusV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminUpgradeHeadlessAccountV3(input *users.AdminUpgradeHeadlessAccountV3Params) (*iamclientmodels.ModelUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, conflict, internalServerError, err := u.Client.Users.AdminUpgradeHeadlessAccountV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminDeleteUserInformationV3(input *users.AdminDeleteUserInformationV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, unauthorized, forbidden, notFound, err := u.Client.Users.AdminDeleteUserInformationV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminGetUserLoginHistoriesV3(input *users.AdminGetUserLoginHistoriesV3Params) (*iamclientmodels.ModelLoginHistoriesResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, err := u.Client.Users.AdminGetUserLoginHistoriesV3(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminUpdateUserPermissionV3(input *users.AdminUpdateUserPermissionV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.AdminUpdateUserPermissionV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminAddUserPermissionsV3(input *users.AdminAddUserPermissionsV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.AdminAddUserPermissionsV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminDeleteUserPermissionBulkV3(input *users.AdminDeleteUserPermissionBulkV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.AdminDeleteUserPermissionBulkV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminDeleteUserPermissionV3(input *users.AdminDeleteUserPermissionV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.AdminDeleteUserPermissionV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminGetUserPlatformAccountsV3(input *users.AdminGetUserPlatformAccountsV3Params) (*iamclientmodels.AccountcommonUserLinkedPlatformsResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.AdminGetUserPlatformAccountsV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminGetListJusticePlatformAccounts(input *users.AdminGetListJusticePlatformAccountsParams) ([]*iamclientmodels.ModelGetUserMapping, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminGetListJusticePlatformAccounts(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminCreateJusticeUser(input *users.AdminCreateJusticeUserParams) (*iamclientmodels.ModelCreateJusticeUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminCreateJusticeUser(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminLinkPlatformAccount(input *users.AdminLinkPlatformAccountParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, internalServerError, err := u.Client.Users.AdminLinkPlatformAccount(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if unauthorized != nil {
		return unauthorized
	}
	if forbidden != nil {
		return forbidden
	}
	if internalServerError != nil {
		return internalServerError
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminPlatformUnlinkV3(input *users.AdminPlatformUnlinkV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.AdminPlatformUnlinkV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminPlatformLinkV3(input *users.AdminPlatformLinkV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, conflict, internalServerError, err := u.Client.Users.AdminPlatformLinkV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) AdminDeleteUserRolesV3(input *users.AdminDeleteUserRolesV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.AdminDeleteUserRolesV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminSaveUserRoleV3(input *users.AdminSaveUserRoleV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, forbidden, notFound, unprocessableEntity, internalServerError, err := u.Client.Users.AdminSaveUserRoleV3(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if forbidden != nil {
		return forbidden
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

func (u *UsersService) AdminAddUserRoleV3(input *users.AdminAddUserRoleV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, conflict, err := u.Client.Users.AdminAddUserRoleV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if conflict != nil {
		return conflict
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminDeleteUserRoleV3(input *users.AdminDeleteUserRoleV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.AdminDeleteUserRoleV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminUpdateUserStatusV3(input *users.AdminUpdateUserStatusV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.AdminUpdateUserStatusV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminVerifyUserWithoutVerificationCodeV3(input *users.AdminVerifyUserWithoutVerificationCodeV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, conflict, err := u.Client.Users.AdminVerifyUserWithoutVerificationCodeV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if conflict != nil {
		return conflict
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminGetMyUserV3(input *users.AdminGetMyUserV3Params) (*iamclientmodels.ModelUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, internalServerError, err := u.Client.Users.AdminGetMyUserV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicListUserIDByPlatformUserIDsV3(input *users.PublicListUserIDByPlatformUserIDsV3Params) (*iamclientmodels.AccountcommonUserPlatforms, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, internalServerError, err := u.Client.Users.PublicListUserIDByPlatformUserIDsV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicGetUserByPlatformUserIDV3(input *users.PublicGetUserByPlatformUserIDV3Params) (*iamclientmodels.ModelUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.PublicGetUserByPlatformUserIDV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicGetAsyncStatus(input *users.PublicGetAsyncStatusParams) (*iamclientmodels.ModelLinkRequest, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, internalServerError, err := u.Client.Users.PublicGetAsyncStatus(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicSearchUserV3(input *users.PublicSearchUserV3Params) (*iamclientmodels.ModelPublicUserInformationResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, internalServerError, err := u.Client.Users.PublicSearchUserV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicCreateUserV3(input *users.PublicCreateUserV3Params) (*iamclientmodels.ModelUserCreateResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, notFound, conflict, internalServerError, err := u.Client.Users.PublicCreateUserV3(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
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

func (u *UsersService) CheckUserAvailability(input *users.CheckUserAvailabilityParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, unprocessableEntity, err := u.Client.Users.CheckUserAvailability(input, client.BearerToken(*accessToken.AccessToken))
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
	if unprocessableEntity != nil {
		return unprocessableEntity
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicBulkGetUsers(input *users.PublicBulkGetUsersParams) (*iamclientmodels.ModelListBulkUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, internalServerError, err := u.Client.Users.PublicBulkGetUsers(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if internalServerError != nil {
		return nil, internalServerError
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicForgotPasswordV3(input *users.PublicForgotPasswordV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, notFound, tooManyRequests, err := u.Client.Users.PublicForgotPasswordV3(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if notFound != nil {
		return notFound
	}
	if tooManyRequests != nil {
		return tooManyRequests
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) GetAdminInvitationV3(input *users.GetAdminInvitationV3Params) (*iamclientmodels.ModelUserInvitationV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, internalServerError, err := u.Client.Users.GetAdminInvitationV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) CreateUserFromInvitationV3(input *users.CreateUserFromInvitationV3Params) (*iamclientmodels.ModelUserCreateResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, notFound, internalServerError, err := u.Client.Users.CreateUserFromInvitationV3(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
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

func (u *UsersService) UpdateUserV3(input *users.UpdateUserV3Params) ([]*iamclientmodels.ModelUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, conflict, internalServerError, err := u.Client.Users.UpdateUserV3(input, client.BearerToken(*accessToken.AccessToken))
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
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicUpdateUserV3(input *users.PublicUpdateUserV3Params) ([]*iamclientmodels.ModelUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, conflict, internalServerError, err := u.Client.Users.PublicUpdateUserV3(input, client.BearerToken(*accessToken.AccessToken))
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
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicSendVerificationCodeV3(input *users.PublicSendVerificationCodeV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, notFound, conflict, tooManyRequests, err := u.Client.Users.PublicSendVerificationCodeV3(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if unauthorized != nil {
		return unauthorized
	}
	if notFound != nil {
		return notFound
	}
	if conflict != nil {
		return conflict
	}
	if tooManyRequests != nil {
		return tooManyRequests
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicUserVerificationV3(input *users.PublicUserVerificationV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, conflict, err := u.Client.Users.PublicUserVerificationV3(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if unauthorized != nil {
		return unauthorized
	}
	if forbidden != nil {
		return forbidden
	}
	if conflict != nil {
		return conflict
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicUpgradeHeadlessAccountV3(input *users.PublicUpgradeHeadlessAccountV3Params) (*iamclientmodels.ModelUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, conflict, internalServerError, err := u.Client.Users.PublicUpgradeHeadlessAccountV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicVerifyHeadlessAccountV3(input *users.PublicVerifyHeadlessAccountV3Params) (*iamclientmodels.ModelUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, conflict, internalServerError, err := u.Client.Users.PublicVerifyHeadlessAccountV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicUpdatePasswordV3(input *users.PublicUpdatePasswordV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, internalServerError, err := u.Client.Users.PublicUpdatePasswordV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicCreateJusticeUser(input *users.PublicCreateJusticeUserParams) (*iamclientmodels.ModelCreateJusticeUserResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.PublicCreateJusticeUser(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicPlatformLinkV3(input *users.PublicPlatformLinkV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, notFound, conflict, internalServerError, err := u.Client.Users.PublicPlatformLinkV3(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if unauthorized != nil {
		return unauthorized
	}
	if notFound != nil {
		return notFound
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

func (u *UsersService) PublicPlatformUnlinkV3(input *users.PublicPlatformUnlinkV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, notFound, internalServerError, err := u.Client.Users.PublicPlatformUnlinkV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicWebLinkPlatform(input *users.PublicWebLinkPlatformParams) (*iamclientmodels.ModelWebLinkingResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, notFound, err := u.Client.Users.PublicWebLinkPlatform(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if unauthorized != nil {
		return nil, unauthorized
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicWebLinkPlatformEstablish(input *users.PublicWebLinkPlatformEstablishParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, err = u.Client.Users.PublicWebLinkPlatformEstablish(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) ResetPasswordV3(input *users.ResetPasswordV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, forbidden, notFound, err := u.Client.Users.ResetPasswordV3(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
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

func (u *UsersService) PublicGetUserByUserIDV3(input *users.PublicGetUserByUserIDV3Params) (*iamclientmodels.ModelPublicUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, internalServerError, err := u.Client.Users.PublicGetUserByUserIDV3(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
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

func (u *UsersService) PublicGetUserBanHistoryV3(input *users.PublicGetUserBanHistoryV3Params) (*iamclientmodels.ModelGetUserBanV3Response, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.PublicGetUserBanHistoryV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicGetUserLoginHistoriesV3(input *users.PublicGetUserLoginHistoriesV3Params) (*iamclientmodels.ModelLoginHistoriesResponse, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, notFound, err := u.Client.Users.PublicGetUserLoginHistoriesV3(input, client.BearerToken(*accessToken.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicGetUserPlatformAccountsV3(input *users.PublicGetUserPlatformAccountsV3Params) (*iamclientmodels.AccountcommonUserLinkedPlatformsResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, unauthorized, forbidden, notFound, err := u.Client.Users.PublicGetUserPlatformAccountsV3(input, client.BearerToken(*accessToken.AccessToken))
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
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicLinkPlatformAccount(input *users.PublicLinkPlatformAccountParams) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, internalServerError, err := u.Client.Users.PublicLinkPlatformAccount(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if unauthorized != nil {
		return unauthorized
	}
	if forbidden != nil {
		return forbidden
	}
	if internalServerError != nil {
		return internalServerError
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicValidateUserByUserIDAndPasswordV3(input *users.PublicValidateUserByUserIDAndPasswordV3Params) error {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err := u.Client.Users.PublicValidateUserByUserIDAndPasswordV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) PublicGetMyUserV3(input *users.PublicGetMyUserV3Params) (*iamclientmodels.ModelUserResponseV3, error) {
	accessToken, err := u.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, internalServerError, err := u.Client.Users.PublicGetMyUserV3(input, client.BearerToken(*accessToken.AccessToken))
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

func (u *UsersService) CreateUserShort(input *users.CreateUserParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserCreateResponse, error) {
	created, err := u.Client.Users.CreateUserShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) GetAdminUsersByRoleIDShort(input *users.GetAdminUsersByRoleIDParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelGetAdminUsersResponse, error) {
	ok, err := u.Client.Users.GetAdminUsersByRoleIDShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) GetUserByLoginIDShort(input *users.GetUserByLoginIDParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelPublicUserResponse, error) {
	ok, err := u.Client.Users.GetUserByLoginIDShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) GetUserByPlatformUserIDShort(input *users.GetUserByPlatformUserIDParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelPublicUserResponse, error) {
	ok, err := u.Client.Users.GetUserByPlatformUserIDShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) ForgotPasswordShort(input *users.ForgotPasswordParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.ForgotPasswordShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) GetUsersByLoginIdsShort(input *users.GetUsersByLoginIdsParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelPublicUsersResponse, error) {
	ok, err := u.Client.Users.GetUsersByLoginIdsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) ResetPasswordShort(input *users.ResetPasswordParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.ResetPasswordShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) SearchUserShort(input *users.SearchUserParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelSearchUsersResponse, error) {
	ok, err := u.Client.Users.SearchUserShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) GetUserByUserIDShort(input *users.GetUserByUserIDParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponse, error) {
	ok, err := u.Client.Users.GetUserByUserIDShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) UpdateUserShort(input *users.UpdateUserParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponse, error) {
	ok, err := u.Client.Users.UpdateUserShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) DeleteUserShort(input *users.DeleteUserParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.DeleteUserShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) BanUserShort(input *users.BanUserParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserBanResponse, error) {
	created, err := u.Client.Users.BanUserShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) GetUserBanHistoryShort(input *users.GetUserBanHistoryParams, authInfo runtime.ClientAuthInfoWriter) ([]*iamclientmodels.ModelUserBanResponse, error) {
	ok, err := u.Client.Users.GetUserBanHistoryShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) DisableUserBanShort(input *users.DisableUserBanParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserBanResponse, error) {
	ok, err := u.Client.Users.DisableUserBanShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) EnableUserBanShort(input *users.EnableUserBanParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserBanResponse, error) {
	ok, err := u.Client.Users.EnableUserBanShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) ListCrossNamespaceAccountLinkShort(input *users.ListCrossNamespaceAccountLinkParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.ListCrossNamespaceAccountLinkShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) DisableUserShort(input *users.DisableUserParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.DisableUserShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) EnableUserShort(input *users.EnableUserParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.EnableUserShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) GetUserInformationShort(input *users.GetUserInformationParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserInformation, error) {
	ok, err := u.Client.Users.GetUserInformationShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) DeleteUserInformationShort(input *users.DeleteUserInformationParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.DeleteUserInformationShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) GetUserLoginHistoriesShort(input *users.GetUserLoginHistoriesParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelLoginHistoriesResponse, error) {
	ok, err := u.Client.Users.GetUserLoginHistoriesShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) UpdatePasswordShort(input *users.UpdatePasswordParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.UpdatePasswordShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) SaveUserPermissionShort(input *users.SaveUserPermissionParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.SaveUserPermissionShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AddUserPermissionShort(input *users.AddUserPermissionParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AddUserPermissionShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) DeleteUserPermissionShort(input *users.DeleteUserPermissionParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.DeleteUserPermissionShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) GetUserPlatformAccountsShort(input *users.GetUserPlatformAccountsParams, authInfo runtime.ClientAuthInfoWriter) ([]*iamclientmodels.AccountcommonUserLinkedPlatform, error) {
	ok, err := u.Client.Users.GetUserPlatformAccountsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) GetUserMappingShort(input *users.GetUserMappingParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelGetUserMapping, error) {
	ok, err := u.Client.Users.GetUserMappingShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) GetUserJusticePlatformAccountShort(input *users.GetUserJusticePlatformAccountParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelGetUserJusticePlatformAccountResponse, error) {
	ok, err := u.Client.Users.GetUserJusticePlatformAccountShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PlatformLinkShort(input *users.PlatformLinkParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PlatformLinkShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PlatformUnlinkShort(input *users.PlatformUnlinkParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PlatformUnlinkShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) GetPublisherUserShort(input *users.GetPublisherUserParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelGetPublisherUserResponse, error) {
	ok, err := u.Client.Users.GetPublisherUserShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) SaveUserRolesShort(input *users.SaveUserRolesParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.SaveUserRolesShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AddUserRoleShort(input *users.AddUserRoleParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AddUserRoleShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) DeleteUserRoleShort(input *users.DeleteUserRoleParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.DeleteUserRoleShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) UpgradeHeadlessAccountShort(input *users.UpgradeHeadlessAccountParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponse, error) {
	ok, err := u.Client.Users.UpgradeHeadlessAccountShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) UpgradeHeadlessAccountWithVerificationCodeShort(input *users.UpgradeHeadlessAccountWithVerificationCodeParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponse, error) {
	ok, err := u.Client.Users.UpgradeHeadlessAccountWithVerificationCodeShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) UserVerificationShort(input *users.UserVerificationParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.UserVerificationShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) SendVerificationCodeShort(input *users.SendVerificationCodeParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.SendVerificationCodeShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminGetAgeRestrictionStatusV2Short(input *users.AdminGetAgeRestrictionStatusV2Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelAgeRestrictionResponse, error) {
	ok, err := u.Client.Users.AdminGetAgeRestrictionStatusV2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminUpdateAgeRestrictionConfigV2Short(input *users.AdminUpdateAgeRestrictionConfigV2Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelAgeRestrictionResponse, error) {
	ok, err := u.Client.Users.AdminUpdateAgeRestrictionConfigV2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) GetListCountryAgeRestrictionShort(input *users.GetListCountryAgeRestrictionParams, authInfo runtime.ClientAuthInfoWriter) ([]*iamclientmodels.AccountcommonCountryAgeRestriction, error) {
	ok, err := u.Client.Users.GetListCountryAgeRestrictionShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) UpdateCountryAgeRestrictionShort(input *users.UpdateCountryAgeRestrictionParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelCountry, error) {
	ok, err := u.Client.Users.UpdateCountryAgeRestrictionShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminSearchUsersV2Short(input *users.AdminSearchUsersV2Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelSearchUsersByPlatformIDResponse, error) {
	ok, err := u.Client.Users.AdminSearchUsersV2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminGetUserByUserIDV2Short(input *users.AdminGetUserByUserIDV2Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponse, error) {
	ok, err := u.Client.Users.AdminGetUserByUserIDV2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminUpdateUserV2Short(input *users.AdminUpdateUserV2Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponse, error) {
	ok, err := u.Client.Users.AdminUpdateUserV2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminBanUserV2Short(input *users.AdminBanUserV2Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserBanResponse, error) {
	created, err := u.Client.Users.AdminBanUserV2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) AdminGetUserBanV2Short(input *users.AdminGetUserBanV2Params, authInfo runtime.ClientAuthInfoWriter) ([]*iamclientmodels.ModelUserBanResponse, error) {
	ok, err := u.Client.Users.AdminGetUserBanV2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminDisableUserV2Short(input *users.AdminDisableUserV2Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminDisableUserV2Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminEnableUserV2Short(input *users.AdminEnableUserV2Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminEnableUserV2Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminResetPasswordV2Short(input *users.AdminResetPasswordV2Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminResetPasswordV2Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminDeletePlatformLinkV2Short(input *users.AdminDeletePlatformLinkV2Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminDeletePlatformLinkV2Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminPutUserRolesV2Short(input *users.AdminPutUserRolesV2Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminPutUserRolesV2Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminCreateUserRolesV2Short(input *users.AdminCreateUserRolesV2Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminCreateUserRolesV2Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicCreateUserV2Short(input *users.PublicCreateUserV2Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserCreateResponse, error) {
	created, err := u.Client.Users.PublicCreateUserV2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) PublicForgotPasswordV2Short(input *users.PublicForgotPasswordV2Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicForgotPasswordV2Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicResetPasswordV2Short(input *users.PublicResetPasswordV2Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicResetPasswordV2Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicGetUserByUserIDV2Short(input *users.PublicGetUserByUserIDV2Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponse, error) {
	ok, err := u.Client.Users.PublicGetUserByUserIDV2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicUpdateUserV2Short(input *users.PublicUpdateUserV2Params, authInfo runtime.ClientAuthInfoWriter) ([]*iamclientmodels.ModelUserResponse, error) {
	ok, err := u.Client.Users.PublicUpdateUserV2Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicGetUserBanShort(input *users.PublicGetUserBanParams, authInfo runtime.ClientAuthInfoWriter) ([]*iamclientmodels.ModelUserBanResponse, error) {
	ok, err := u.Client.Users.PublicGetUserBanShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicUpdatePasswordV2Short(input *users.PublicUpdatePasswordV2Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicUpdatePasswordV2Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) GetListJusticePlatformAccountsShort(input *users.GetListJusticePlatformAccountsParams, authInfo runtime.ClientAuthInfoWriter) ([]*iamclientmodels.ModelGetUserMapping, error) {
	ok, err := u.Client.Users.GetListJusticePlatformAccountsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicPlatformLinkV2Short(input *users.PublicPlatformLinkV2Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicPlatformLinkV2Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicDeletePlatformLinkV2Short(input *users.PublicDeletePlatformLinkV2Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicDeletePlatformLinkV2Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) ListAdminsV3Short(input *users.ListAdminsV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelGetUsersResponseWithPaginationV3, error) {
	ok, err := u.Client.Users.ListAdminsV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminGetAgeRestrictionStatusV3Short(input *users.AdminGetAgeRestrictionStatusV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelAgeRestrictionResponseV3, error) {
	ok, err := u.Client.Users.AdminGetAgeRestrictionStatusV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminUpdateAgeRestrictionConfigV3Short(input *users.AdminUpdateAgeRestrictionConfigV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelAgeRestrictionResponseV3, error) {
	ok, err := u.Client.Users.AdminUpdateAgeRestrictionConfigV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminGetListCountryAgeRestrictionV3Short(input *users.AdminGetListCountryAgeRestrictionV3Params, authInfo runtime.ClientAuthInfoWriter) ([]*iamclientmodels.ModelCountryV3Response, error) {
	ok, err := u.Client.Users.AdminGetListCountryAgeRestrictionV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminUpdateCountryAgeRestrictionV3Short(input *users.AdminUpdateCountryAgeRestrictionV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelCountryV3Response, error) {
	ok, err := u.Client.Users.AdminUpdateCountryAgeRestrictionV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminGetUserByPlatformUserIDV3Short(input *users.AdminGetUserByPlatformUserIDV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponseV3, error) {
	ok, err := u.Client.Users.AdminGetUserByPlatformUserIDV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) GetAdminUsersByRoleIDV3Short(input *users.GetAdminUsersByRoleIDV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelGetUsersResponseWithPaginationV3, error) {
	ok, err := u.Client.Users.GetAdminUsersByRoleIDV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminGetUserByEmailAddressV3Short(input *users.AdminGetUserByEmailAddressV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponseV3, error) {
	ok, err := u.Client.Users.AdminGetUserByEmailAddressV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminListUserIDByUserIDsV3Short(input *users.AdminListUserIDByUserIDsV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelListUserInformationResult, error) {
	ok, err := u.Client.Users.AdminListUserIDByUserIDsV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminInviteUserV3Short(input *users.AdminInviteUserV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelInviteUserResponseV3, error) {
	created, err := u.Client.Users.AdminInviteUserV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) AdminListUsersV3Short(input *users.AdminListUsersV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.AccountcommonListUsersWithPlatformAccountsResponse, error) {
	ok, err := u.Client.Users.AdminListUsersV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminSearchUserV3Short(input *users.AdminSearchUserV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelSearchUsersResponseWithPaginationV3, error) {
	ok, err := u.Client.Users.AdminSearchUserV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminGetBulkUserByEmailAddressV3Short(input *users.AdminGetBulkUserByEmailAddressV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelListUserResponseV3, error) {
	ok, err := u.Client.Users.AdminGetBulkUserByEmailAddressV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminGetUserByUserIDV3Short(input *users.AdminGetUserByUserIDV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponseV3, error) {
	ok, err := u.Client.Users.AdminGetUserByUserIDV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminUpdateUserV3Short(input *users.AdminUpdateUserV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponseV3, error) {
	ok, err := u.Client.Users.AdminUpdateUserV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminGetUserBanV3Short(input *users.AdminGetUserBanV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelGetUserBanV3Response, error) {
	ok, err := u.Client.Users.AdminGetUserBanV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminBanUserV3Short(input *users.AdminBanUserV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserBanResponseV3, error) {
	created, err := u.Client.Users.AdminBanUserV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) AdminUpdateUserBanV3Short(input *users.AdminUpdateUserBanV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserBanResponseV3, error) {
	ok, err := u.Client.Users.AdminUpdateUserBanV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminSendVerificationCodeV3Short(input *users.AdminSendVerificationCodeV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminSendVerificationCodeV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminVerifyAccountV3Short(input *users.AdminVerifyAccountV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminVerifyAccountV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) GetUserVerificationCodeShort(input *users.GetUserVerificationCodeParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelVerificationCodeResponse, error) {
	ok, err := u.Client.Users.GetUserVerificationCodeShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminGetUserDeletionStatusV3Short(input *users.AdminGetUserDeletionStatusV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserDeletionStatusResponse, error) {
	ok, err := u.Client.Users.AdminGetUserDeletionStatusV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminUpdateUserDeletionStatusV3Short(input *users.AdminUpdateUserDeletionStatusV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminUpdateUserDeletionStatusV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminUpgradeHeadlessAccountV3Short(input *users.AdminUpgradeHeadlessAccountV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponseV3, error) {
	ok, err := u.Client.Users.AdminUpgradeHeadlessAccountV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminDeleteUserInformationV3Short(input *users.AdminDeleteUserInformationV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminDeleteUserInformationV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminGetUserLoginHistoriesV3Short(input *users.AdminGetUserLoginHistoriesV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelLoginHistoriesResponse, error) {
	ok, err := u.Client.Users.AdminGetUserLoginHistoriesV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminUpdateUserPermissionV3Short(input *users.AdminUpdateUserPermissionV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminUpdateUserPermissionV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminAddUserPermissionsV3Short(input *users.AdminAddUserPermissionsV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminAddUserPermissionsV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminDeleteUserPermissionBulkV3Short(input *users.AdminDeleteUserPermissionBulkV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminDeleteUserPermissionBulkV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminDeleteUserPermissionV3Short(input *users.AdminDeleteUserPermissionV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminDeleteUserPermissionV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminGetUserPlatformAccountsV3Short(input *users.AdminGetUserPlatformAccountsV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.AccountcommonUserLinkedPlatformsResponseV3, error) {
	ok, err := u.Client.Users.AdminGetUserPlatformAccountsV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminGetListJusticePlatformAccountsShort(input *users.AdminGetListJusticePlatformAccountsParams, authInfo runtime.ClientAuthInfoWriter) ([]*iamclientmodels.ModelGetUserMapping, error) {
	ok, err := u.Client.Users.AdminGetListJusticePlatformAccountsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) AdminCreateJusticeUserShort(input *users.AdminCreateJusticeUserParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelCreateJusticeUserResponse, error) {
	created, err := u.Client.Users.AdminCreateJusticeUserShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) AdminLinkPlatformAccountShort(input *users.AdminLinkPlatformAccountParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminLinkPlatformAccountShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminPlatformUnlinkV3Short(input *users.AdminPlatformUnlinkV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminPlatformUnlinkV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminPlatformLinkV3Short(input *users.AdminPlatformLinkV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminPlatformLinkV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminDeleteUserRolesV3Short(input *users.AdminDeleteUserRolesV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminDeleteUserRolesV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminSaveUserRoleV3Short(input *users.AdminSaveUserRoleV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminSaveUserRoleV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminAddUserRoleV3Short(input *users.AdminAddUserRoleV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminAddUserRoleV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminDeleteUserRoleV3Short(input *users.AdminDeleteUserRoleV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminDeleteUserRoleV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminUpdateUserStatusV3Short(input *users.AdminUpdateUserStatusV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminUpdateUserStatusV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminVerifyUserWithoutVerificationCodeV3Short(input *users.AdminVerifyUserWithoutVerificationCodeV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.AdminVerifyUserWithoutVerificationCodeV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) AdminGetMyUserV3Short(input *users.AdminGetMyUserV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponseV3, error) {
	ok, err := u.Client.Users.AdminGetMyUserV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicListUserIDByPlatformUserIDsV3Short(input *users.PublicListUserIDByPlatformUserIDsV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.AccountcommonUserPlatforms, error) {
	ok, err := u.Client.Users.PublicListUserIDByPlatformUserIDsV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicGetUserByPlatformUserIDV3Short(input *users.PublicGetUserByPlatformUserIDV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponseV3, error) {
	ok, err := u.Client.Users.PublicGetUserByPlatformUserIDV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicGetAsyncStatusShort(input *users.PublicGetAsyncStatusParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelLinkRequest, error) {
	ok, err := u.Client.Users.PublicGetAsyncStatusShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicSearchUserV3Short(input *users.PublicSearchUserV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelPublicUserInformationResponseV3, error) {
	ok, err := u.Client.Users.PublicSearchUserV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicCreateUserV3Short(input *users.PublicCreateUserV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserCreateResponseV3, error) {
	created, err := u.Client.Users.PublicCreateUserV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) CheckUserAvailabilityShort(input *users.CheckUserAvailabilityParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.CheckUserAvailabilityShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicBulkGetUsersShort(input *users.PublicBulkGetUsersParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelListBulkUserResponse, error) {
	ok, err := u.Client.Users.PublicBulkGetUsersShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicForgotPasswordV3Short(input *users.PublicForgotPasswordV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicForgotPasswordV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) GetAdminInvitationV3Short(input *users.GetAdminInvitationV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserInvitationV3, error) {
	ok, err := u.Client.Users.GetAdminInvitationV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) CreateUserFromInvitationV3Short(input *users.CreateUserFromInvitationV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserCreateResponseV3, error) {
	created, err := u.Client.Users.CreateUserFromInvitationV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) UpdateUserV3Short(input *users.UpdateUserV3Params, authInfo runtime.ClientAuthInfoWriter) ([]*iamclientmodels.ModelUserResponseV3, error) {
	ok, err := u.Client.Users.UpdateUserV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicUpdateUserV3Short(input *users.PublicUpdateUserV3Params, authInfo runtime.ClientAuthInfoWriter) ([]*iamclientmodels.ModelUserResponseV3, error) {
	ok, err := u.Client.Users.PublicUpdateUserV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicSendVerificationCodeV3Short(input *users.PublicSendVerificationCodeV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicSendVerificationCodeV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicUserVerificationV3Short(input *users.PublicUserVerificationV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicUserVerificationV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicUpgradeHeadlessAccountV3Short(input *users.PublicUpgradeHeadlessAccountV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponseV3, error) {
	ok, err := u.Client.Users.PublicUpgradeHeadlessAccountV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicVerifyHeadlessAccountV3Short(input *users.PublicVerifyHeadlessAccountV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponseV3, error) {
	ok, err := u.Client.Users.PublicVerifyHeadlessAccountV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicUpdatePasswordV3Short(input *users.PublicUpdatePasswordV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicUpdatePasswordV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicCreateJusticeUserShort(input *users.PublicCreateJusticeUserParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelCreateJusticeUserResponse, error) {
	created, err := u.Client.Users.PublicCreateJusticeUserShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (u *UsersService) PublicPlatformLinkV3Short(input *users.PublicPlatformLinkV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicPlatformLinkV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicPlatformUnlinkV3Short(input *users.PublicPlatformUnlinkV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicPlatformUnlinkV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicWebLinkPlatformShort(input *users.PublicWebLinkPlatformParams, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelWebLinkingResponse, error) {
	ok, err := u.Client.Users.PublicWebLinkPlatformShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicWebLinkPlatformEstablishShort(input *users.PublicWebLinkPlatformEstablishParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicWebLinkPlatformEstablishShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) ResetPasswordV3Short(input *users.ResetPasswordV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.ResetPasswordV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicGetUserByUserIDV3Short(input *users.PublicGetUserByUserIDV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelPublicUserResponseV3, error) {
	ok, err := u.Client.Users.PublicGetUserByUserIDV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicGetUserBanHistoryV3Short(input *users.PublicGetUserBanHistoryV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelGetUserBanV3Response, error) {
	ok, err := u.Client.Users.PublicGetUserBanHistoryV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicGetUserLoginHistoriesV3Short(input *users.PublicGetUserLoginHistoriesV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelLoginHistoriesResponse, error) {
	ok, err := u.Client.Users.PublicGetUserLoginHistoriesV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicGetUserPlatformAccountsV3Short(input *users.PublicGetUserPlatformAccountsV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.AccountcommonUserLinkedPlatformsResponseV3, error) {
	ok, err := u.Client.Users.PublicGetUserPlatformAccountsV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (u *UsersService) PublicLinkPlatformAccountShort(input *users.PublicLinkPlatformAccountParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicLinkPlatformAccountShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicValidateUserByUserIDAndPasswordV3Short(input *users.PublicValidateUserByUserIDAndPasswordV3Params, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := u.Client.Users.PublicValidateUserByUserIDAndPasswordV3Short(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersService) PublicGetMyUserV3Short(input *users.PublicGetMyUserV3Params, authInfo runtime.ClientAuthInfoWriter) (*iamclientmodels.ModelUserResponseV3, error) {
	ok, err := u.Client.Users.PublicGetMyUserV3Short(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}
