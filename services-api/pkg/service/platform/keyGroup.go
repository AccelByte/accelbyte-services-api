// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package platform

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/key_group"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils/auth"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type KeyGroupService struct {
	Client          *platformclient.JusticePlatformService
	TokenRepository repository.TokenRepository
}

// Deprecated: Use QueryKeyGroupsShort instead
func (k *KeyGroupService) QueryKeyGroups(input *key_group.QueryKeyGroupsParams) (*platformclientmodels.KeyGroupPagingSlicedResult, error) {
	token, err := k.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := k.Client.KeyGroup.QueryKeyGroups(input, client.BearerToken(*token.AccessToken))
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use CreateKeyGroupShort instead
func (k *KeyGroupService) CreateKeyGroup(input *key_group.CreateKeyGroupParams) (*platformclientmodels.KeyGroupInfo, error) {
	token, err := k.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, conflict, unprocessableEntity, err := k.Client.KeyGroup.CreateKeyGroup(input, client.BearerToken(*token.AccessToken))
	if conflict != nil {
		return nil, conflict
	}
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}

	return created.GetPayload(), nil
}

// Deprecated: Use GetKeyGroupShort instead
func (k *KeyGroupService) GetKeyGroup(input *key_group.GetKeyGroupParams) (*platformclientmodels.KeyGroupInfo, error) {
	token, err := k.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := k.Client.KeyGroup.GetKeyGroup(input, client.BearerToken(*token.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use UpdateKeyGroupShort instead
func (k *KeyGroupService) UpdateKeyGroup(input *key_group.UpdateKeyGroupParams) (*platformclientmodels.KeyGroupInfo, error) {
	token, err := k.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, conflict, unprocessableEntity, err := k.Client.KeyGroup.UpdateKeyGroup(input, client.BearerToken(*token.AccessToken))
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

// Deprecated: Use GetKeyGroupDynamicShort instead
func (k *KeyGroupService) GetKeyGroupDynamic(input *key_group.GetKeyGroupDynamicParams) (*platformclientmodels.KeyGroupDynamicInfo, error) {
	token, err := k.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := k.Client.KeyGroup.GetKeyGroupDynamic(input, client.BearerToken(*token.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use ListKeysShort instead
func (k *KeyGroupService) ListKeys(input *key_group.ListKeysParams) (*platformclientmodels.KeyPagingSliceResult, error) {
	token, err := k.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := k.Client.KeyGroup.ListKeys(input, client.BearerToken(*token.AccessToken))
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use UploadKeysShort instead
func (k *KeyGroupService) UploadKeys(input *key_group.UploadKeysParams) (*platformclientmodels.BulkOperationResult, error) {
	token, err := k.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := k.Client.KeyGroup.UploadKeys(input, client.BearerToken(*token.AccessToken))
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

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:KEYGROUP [READ]'], 'authorization': []}]
func (k *KeyGroupService) QueryKeyGroupsShort(input *key_group.QueryKeyGroupsParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.KeyGroupPagingSlicedResult, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(k.TokenRepository, nil, security, "")
	}
	ok, err := k.Client.KeyGroup.QueryKeyGroupsShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:KEYGROUP [CREATE]'], 'authorization': []}]
func (k *KeyGroupService) CreateKeyGroupShort(input *key_group.CreateKeyGroupParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.KeyGroupInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(k.TokenRepository, nil, security, "")
	}
	created, err := k.Client.KeyGroup.CreateKeyGroupShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return created.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:KEYGROUP [READ]'], 'authorization': []}]
func (k *KeyGroupService) GetKeyGroupShort(input *key_group.GetKeyGroupParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.KeyGroupInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(k.TokenRepository, nil, security, "")
	}
	ok, err := k.Client.KeyGroup.GetKeyGroupShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:KEYGROUP [UPDATE]'], 'authorization': []}]
func (k *KeyGroupService) UpdateKeyGroupShort(input *key_group.UpdateKeyGroupParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.KeyGroupInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(k.TokenRepository, nil, security, "")
	}
	ok, err := k.Client.KeyGroup.UpdateKeyGroupShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:KEYGROUP [READ]'], 'authorization': []}]
func (k *KeyGroupService) GetKeyGroupDynamicShort(input *key_group.GetKeyGroupDynamicParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.KeyGroupDynamicInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(k.TokenRepository, nil, security, "")
	}
	ok, err := k.Client.KeyGroup.GetKeyGroupDynamicShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:KEYGROUP [READ]'], 'authorization': []}]
func (k *KeyGroupService) ListKeysShort(input *key_group.ListKeysParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.KeyPagingSliceResult, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(k.TokenRepository, nil, security, "")
	}
	ok, err := k.Client.KeyGroup.ListKeysShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:KEYGROUP [UPDATE]'], 'authorization': []}]
func (k *KeyGroupService) UploadKeysShort(input *key_group.UploadKeysParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.BulkOperationResult, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(k.TokenRepository, nil, security, "")
	}
	ok, err := k.Client.KeyGroup.UploadKeysShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}
