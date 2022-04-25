// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package social

import (
	"io"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils/auth"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclient"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclient/slot"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type SlotService struct {
	Client          *socialclient.JusticeSocialService
	TokenRepository repository.TokenRepository
}

// Deprecated: Use GetUserNamespaceSlotsShort instead
func (s *SlotService) GetUserNamespaceSlots(input *slot.GetUserNamespaceSlotsParams) ([]*socialclientmodels.SlotInfo, error) {
	token, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Slot.GetUserNamespaceSlots(input, client.BearerToken(*token.AccessToken))
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use GetSlotDataShort instead
func (s *SlotService) GetSlotData(input *slot.GetSlotDataParams, writer io.Writer) (io.Writer, error) {
	token, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := s.Client.Slot.GetSlotData(input, client.BearerToken(*token.AccessToken), writer)
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use PublicGetUserNamespaceSlotsShort instead
func (s *SlotService) PublicGetUserNamespaceSlots(input *slot.PublicGetUserNamespaceSlotsParams) ([]*socialclientmodels.SlotInfo, error) {
	token, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Slot.PublicGetUserNamespaceSlots(input, client.BearerToken(*token.AccessToken))
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use PublicCreateUserNamespaceSlotShort instead
func (s *SlotService) PublicCreateUserNamespaceSlot(input *slot.PublicCreateUserNamespaceSlotParams) error {
	token, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, conflict, err := s.Client.Slot.PublicCreateUserNamespaceSlot(input, client.BearerToken(*token.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if conflict != nil {
		return conflict
	}
	if err != nil {
		return err
	}

	return nil
}

// Deprecated: Use PublicGetSlotDataShort instead
func (s *SlotService) PublicGetSlotData(input *slot.PublicGetSlotDataParams, writer io.Writer) (io.Writer, error) {
	token, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := s.Client.Slot.PublicGetSlotData(input, client.BearerToken(*token.AccessToken), writer)
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use PublicUpdateUserNamespaceSlotShort instead
func (s *SlotService) PublicUpdateUserNamespaceSlot(input *slot.PublicUpdateUserNamespaceSlotParams) (*socialclientmodels.SlotInfo, error) {
	token, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := s.Client.Slot.PublicUpdateUserNamespaceSlot(input, client.BearerToken(*token.AccessToken))
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

// Deprecated: Use PublicDeleteUserNamespaceSlotShort instead
func (s *SlotService) PublicDeleteUserNamespaceSlot(input *slot.PublicDeleteUserNamespaceSlotParams) error {
	token, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, notFound, err := s.Client.Slot.PublicDeleteUserNamespaceSlot(input, client.BearerToken(*token.AccessToken))
	if notFound != nil {
		return notFound
	}
	if err != nil {
		return err
	}

	return nil
}

// Deprecated: Use PublicUpdateUserNamespaceSlotMetadataShort instead
func (s *SlotService) PublicUpdateUserNamespaceSlotMetadata(input *slot.PublicUpdateUserNamespaceSlotMetadataParams) (*socialclientmodels.SlotInfo, error) {
	token, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := s.Client.Slot.PublicUpdateUserNamespaceSlotMetadata(input, client.BearerToken(*token.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:USER:{userId}:SLOTDATA [READ]'], 'authorization': []}]
func (s *SlotService) GetUserNamespaceSlotsShort(input *slot.GetUserNamespaceSlotsParams, authInfoWriter runtime.ClientAuthInfoWriter) ([]*socialclientmodels.SlotInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(s.TokenRepository, nil, security, "")
	}
	ok, err := s.Client.Slot.GetUserNamespaceSlotsShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:USER:{userId}:SLOTDATA [READ]'], 'authorization': []}]
func (s *SlotService) GetSlotDataShort(input *slot.GetSlotDataParams, authInfoWriter runtime.ClientAuthInfoWriter, writer io.Writer) (io.Writer, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(s.TokenRepository, nil, security, "")
	}
	ok, err := s.Client.Slot.GetSlotDataShort(input, authInfoWriter, writer)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['NAMESPACE:{namespace}:USER:{userId}:SLOTDATA [READ]'], 'authorization': []}]
func (s *SlotService) PublicGetUserNamespaceSlotsShort(input *slot.PublicGetUserNamespaceSlotsParams, authInfoWriter runtime.ClientAuthInfoWriter) ([]*socialclientmodels.SlotInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(s.TokenRepository, nil, security, "")
	}
	ok, err := s.Client.Slot.PublicGetUserNamespaceSlotsShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['NAMESPACE:{namespace}:USER:{userId}:SLOTDATA [CREATE]'], 'authorization': []}]
func (s *SlotService) PublicCreateUserNamespaceSlotShort(input *slot.PublicCreateUserNamespaceSlotParams, authInfoWriter runtime.ClientAuthInfoWriter) error {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(s.TokenRepository, nil, security, "")
	}
	_, err := s.Client.Slot.PublicCreateUserNamespaceSlotShort(input, authInfoWriter)
	if err != nil {
		return err
	}

	return nil
}

// [{'authorization': []}, {'HasPermission': ['NAMESPACE:{namespace}:USER:{userId}:SLOTDATA [READ]'], 'authorization': []}]
func (s *SlotService) PublicGetSlotDataShort(input *slot.PublicGetSlotDataParams, authInfoWriter runtime.ClientAuthInfoWriter, writer io.Writer) (io.Writer, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(s.TokenRepository, nil, security, "")
	}
	ok, err := s.Client.Slot.PublicGetSlotDataShort(input, authInfoWriter, writer)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['NAMESPACE:{namespace}:USER:{userId}:SLOTDATA [UPDATE]'], 'authorization': []}]
func (s *SlotService) PublicUpdateUserNamespaceSlotShort(input *slot.PublicUpdateUserNamespaceSlotParams, authInfoWriter runtime.ClientAuthInfoWriter) (*socialclientmodels.SlotInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(s.TokenRepository, nil, security, "")
	}
	ok, err := s.Client.Slot.PublicUpdateUserNamespaceSlotShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['NAMESPACE:{namespace}:USER:{userId}:SLOTDATA [DELETE]'], 'authorization': []}]
func (s *SlotService) PublicDeleteUserNamespaceSlotShort(input *slot.PublicDeleteUserNamespaceSlotParams, authInfoWriter runtime.ClientAuthInfoWriter) error {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(s.TokenRepository, nil, security, "")
	}
	_, err := s.Client.Slot.PublicDeleteUserNamespaceSlotShort(input, authInfoWriter)
	if err != nil {
		return err
	}

	return nil
}

// [{'authorization': []}, {'HasPermission': ['NAMESPACE:{namespace}:USER:{userId}:SLOTDATA [UPDATE]'], 'authorization': []}]
func (s *SlotService) PublicUpdateUserNamespaceSlotMetadataShort(input *slot.PublicUpdateUserNamespaceSlotMetadataParams, authInfoWriter runtime.ClientAuthInfoWriter) (*socialclientmodels.SlotInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(s.TokenRepository, nil, security, "")
	}
	ok, err := s.Client.Slot.PublicUpdateUserNamespaceSlotMetadataShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}
