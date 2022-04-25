// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package platform

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/reward"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils/auth"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type RewardService struct {
	Client          *platformclient.JusticePlatformService
	TokenRepository repository.TokenRepository
}

// Deprecated: Use CreateRewardShort instead
func (r *RewardService) CreateReward(input *reward.CreateRewardParams) (*platformclientmodels.RewardInfo, error) {
	token, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, conflict, unprocessableEntity, err := r.Client.Reward.CreateReward(input, client.BearerToken(*token.AccessToken))
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

// Deprecated: Use QueryRewardsShort instead
func (r *RewardService) QueryRewards(input *reward.QueryRewardsParams) (*platformclientmodels.RewardPagingSlicedResult, error) {
	token, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := r.Client.Reward.QueryRewards(input, client.BearerToken(*token.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use ExportRewardsShort instead
func (r *RewardService) ExportRewards(input *reward.ExportRewardsParams) error {
	token, err := r.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, err = r.Client.Reward.ExportRewards(input, client.BearerToken(*token.AccessToken))
	if err != nil {
		return err
	}

	return nil
}

// Deprecated: Use ImportRewardsShort instead
func (r *RewardService) ImportRewards(input *reward.ImportRewardsParams) error {
	token, err := r.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, err := r.Client.Reward.ImportRewards(input, client.BearerToken(*token.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if err != nil {
		return err
	}

	return nil
}

// Deprecated: Use GetRewardShort instead
func (r *RewardService) GetReward(input *reward.GetRewardParams) (*platformclientmodels.RewardInfo, error) {
	token, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := r.Client.Reward.GetReward(input, client.BearerToken(*token.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use UpdateRewardShort instead
func (r *RewardService) UpdateReward(input *reward.UpdateRewardParams) (*platformclientmodels.RewardInfo, error) {
	token, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, conflict, err := r.Client.Reward.UpdateReward(input, client.BearerToken(*token.AccessToken))
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

// Deprecated: Use DeleteRewardShort instead
func (r *RewardService) DeleteReward(input *reward.DeleteRewardParams) (*platformclientmodels.RewardInfo, error) {
	token, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := r.Client.Reward.DeleteReward(input, client.BearerToken(*token.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use CheckEventConditionShort instead
func (r *RewardService) CheckEventCondition(input *reward.CheckEventConditionParams) (*platformclientmodels.ConditionMatchResult, error) {
	token, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := r.Client.Reward.CheckEventCondition(input, client.BearerToken(*token.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use GetRewardByCodeShort instead
func (r *RewardService) GetRewardByCode(input *reward.GetRewardByCodeParams) (*platformclientmodels.RewardInfo, error) {
	token, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := r.Client.Reward.GetRewardByCode(input, client.BearerToken(*token.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use QueryRewards1Short instead
func (r *RewardService) QueryRewards1(input *reward.QueryRewards1Params) (*platformclientmodels.RewardPagingSlicedResult, error) {
	token, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unprocessableEntity, err := r.Client.Reward.QueryRewards1(input, client.BearerToken(*token.AccessToken))
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use GetReward1Short instead
func (r *RewardService) GetReward1(input *reward.GetReward1Params) (*platformclientmodels.RewardInfo, error) {
	token, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := r.Client.Reward.GetReward1(input, client.BearerToken(*token.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:REWARD [CREATE]'], 'authorization': []}]
func (r *RewardService) CreateRewardShort(input *reward.CreateRewardParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.RewardInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(r.TokenRepository, nil, security, "")
	}
	ok, err := r.Client.Reward.CreateRewardShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:REWARD [READ]'], 'authorization': []}]
func (r *RewardService) QueryRewardsShort(input *reward.QueryRewardsParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.RewardPagingSlicedResult, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(r.TokenRepository, nil, security, "")
	}
	ok, err := r.Client.Reward.QueryRewardsShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:REWARD [READ]'], 'authorization': []}]
func (r *RewardService) ExportRewardsShort(input *reward.ExportRewardsParams, authInfoWriter runtime.ClientAuthInfoWriter) error {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(r.TokenRepository, nil, security, "")
	}
	_, err := r.Client.Reward.ExportRewardsShort(input, authInfoWriter)
	if err != nil {
		return err
	}

	return nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:REWARD [CREATE]'], 'authorization': []}]
func (r *RewardService) ImportRewardsShort(input *reward.ImportRewardsParams, authInfoWriter runtime.ClientAuthInfoWriter) error {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(r.TokenRepository, nil, security, "")
	}
	_, err := r.Client.Reward.ImportRewardsShort(input, authInfoWriter)
	if err != nil {
		return err
	}

	return nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:REWARD [READ]'], 'authorization': []}]
func (r *RewardService) GetRewardShort(input *reward.GetRewardParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.RewardInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(r.TokenRepository, nil, security, "")
	}
	ok, err := r.Client.Reward.GetRewardShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:REWARD [UPDATE]'], 'authorization': []}]
func (r *RewardService) UpdateRewardShort(input *reward.UpdateRewardParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.RewardInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(r.TokenRepository, nil, security, "")
	}
	ok, err := r.Client.Reward.UpdateRewardShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:REWARD [DELETE]'], 'authorization': []}]
func (r *RewardService) DeleteRewardShort(input *reward.DeleteRewardParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.RewardInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(r.TokenRepository, nil, security, "")
	}
	ok, err := r.Client.Reward.DeleteRewardShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:REWARD [READ]'], 'authorization': []}]
func (r *RewardService) CheckEventConditionShort(input *reward.CheckEventConditionParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.ConditionMatchResult, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(r.TokenRepository, nil, security, "")
	}
	ok, err := r.Client.Reward.CheckEventConditionShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['NAMESPACE:{namespace}:REWARD [READ]'], 'authorization': []}]
func (r *RewardService) GetRewardByCodeShort(input *reward.GetRewardByCodeParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.RewardInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(r.TokenRepository, nil, security, "")
	}
	ok, err := r.Client.Reward.GetRewardByCodeShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['NAMESPACE:{namespace}:REWARD [READ]'], 'authorization': []}]
func (r *RewardService) QueryRewards1Short(input *reward.QueryRewards1Params, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.RewardPagingSlicedResult, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(r.TokenRepository, nil, security, "")
	}
	ok, err := r.Client.Reward.QueryRewards1Short(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['NAMESPACE:{namespace}:REWARD [READ]'], 'authorization': []}]
func (r *RewardService) GetReward1Short(input *reward.GetReward1Params, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.RewardInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(r.TokenRepository, nil, security, "")
	}
	ok, err := r.Client.Reward.GetReward1Short(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}
