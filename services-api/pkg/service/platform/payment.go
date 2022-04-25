// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package platform

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/payment"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils/auth"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type PaymentService struct {
	Client          *platformclient.JusticePlatformService
	TokenRepository repository.TokenRepository
}

// Deprecated: Use QueryPaymentNotificationsShort instead
func (p *PaymentService) QueryPaymentNotifications(input *payment.QueryPaymentNotificationsParams) (*platformclientmodels.PaymentNotificationPagingSlicedResult, error) {
	token, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := p.Client.Payment.QueryPaymentNotifications(input, client.BearerToken(*token.AccessToken))
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use QueryPaymentOrdersShort instead
func (p *PaymentService) QueryPaymentOrders(input *payment.QueryPaymentOrdersParams) (*platformclientmodels.PaymentOrderPagingSlicedResult, error) {
	token, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := p.Client.Payment.QueryPaymentOrders(input, client.BearerToken(*token.AccessToken))
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use ListExtOrderNoByExtTxIDShort instead
func (p *PaymentService) ListExtOrderNoByExtTxID(input *payment.ListExtOrderNoByExtTxIDParams) ([]string, error) {
	token, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := p.Client.Payment.ListExtOrderNoByExtTxID(input, client.BearerToken(*token.AccessToken))
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use GetPaymentOrderShort instead
func (p *PaymentService) GetPaymentOrder(input *payment.GetPaymentOrderParams) (*platformclientmodels.PaymentOrderInfo, error) {
	token, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := p.Client.Payment.GetPaymentOrder(input, client.BearerToken(*token.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use ChargePaymentOrderShort instead
func (p *PaymentService) ChargePaymentOrder(input *payment.ChargePaymentOrderParams) (*platformclientmodels.PaymentOrderInfo, error) {
	token, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, conflict, err := p.Client.Payment.ChargePaymentOrder(input, client.BearerToken(*token.AccessToken))
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

// Deprecated: Use SimulatePaymentOrderNotificationShort instead
func (p *PaymentService) SimulatePaymentOrderNotification(input *payment.SimulatePaymentOrderNotificationParams) (*platformclientmodels.NotificationProcessResult, error) {
	token, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := p.Client.Payment.SimulatePaymentOrderNotification(input, client.BearerToken(*token.AccessToken))
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

// Deprecated: Use GetPaymentOrderChargeStatusShort instead
func (p *PaymentService) GetPaymentOrderChargeStatus(input *payment.GetPaymentOrderChargeStatusParams) (*platformclientmodels.PaymentOrderChargeStatus, error) {
	token, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := p.Client.Payment.GetPaymentOrderChargeStatus(input, client.BearerToken(*token.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use CreateUserPaymentOrderShort instead
func (p *PaymentService) CreateUserPaymentOrder(input *payment.CreateUserPaymentOrderParams) (*platformclientmodels.PaymentOrderInfo, error) {
	token, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, forbidden, notFound, conflict, unprocessableEntity, err := p.Client.Payment.CreateUserPaymentOrder(input, client.BearerToken(*token.AccessToken))
	if badRequest != nil {
		return nil, badRequest
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
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}

	return created.GetPayload(), nil
}

// Deprecated: Use RefundUserPaymentOrderShort instead
func (p *PaymentService) RefundUserPaymentOrder(input *payment.RefundUserPaymentOrderParams) (*platformclientmodels.PaymentOrderInfo, error) {
	token, err := p.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, conflict, unprocessableEntity, err := p.Client.Payment.RefundUserPaymentOrder(input, client.BearerToken(*token.AccessToken))
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

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:PAYMENT:NOTIFICATION [READ]'], 'authorization': []}]
func (p *PaymentService) QueryPaymentNotificationsShort(input *payment.QueryPaymentNotificationsParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.PaymentNotificationPagingSlicedResult, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(p.TokenRepository, nil, security, "")
	}
	ok, err := p.Client.Payment.QueryPaymentNotificationsShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:PAYMENT [READ]'], 'authorization': []}]
func (p *PaymentService) QueryPaymentOrdersShort(input *payment.QueryPaymentOrdersParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.PaymentOrderPagingSlicedResult, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(p.TokenRepository, nil, security, "")
	}
	ok, err := p.Client.Payment.QueryPaymentOrdersShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:PAYMENT [READ]'], 'authorization': []}]
func (p *PaymentService) ListExtOrderNoByExtTxIDShort(input *payment.ListExtOrderNoByExtTxIDParams, authInfoWriter runtime.ClientAuthInfoWriter) ([]string, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(p.TokenRepository, nil, security, "")
	}
	ok, err := p.Client.Payment.ListExtOrderNoByExtTxIDShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:PAYMENT [READ]'], 'authorization': []}]
func (p *PaymentService) GetPaymentOrderShort(input *payment.GetPaymentOrderParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.PaymentOrderInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(p.TokenRepository, nil, security, "")
	}
	ok, err := p.Client.Payment.GetPaymentOrderShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:PAYMENT [UPDATE]'], 'authorization': []}]
func (p *PaymentService) ChargePaymentOrderShort(input *payment.ChargePaymentOrderParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.PaymentOrderInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(p.TokenRepository, nil, security, "")
	}
	ok, err := p.Client.Payment.ChargePaymentOrderShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:PAYMENT [UPDATE]'], 'authorization': []}]
func (p *PaymentService) SimulatePaymentOrderNotificationShort(input *payment.SimulatePaymentOrderNotificationParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.NotificationProcessResult, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(p.TokenRepository, nil, security, "")
	}
	ok, err := p.Client.Payment.SimulatePaymentOrderNotificationShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:PAYMENT [READ]'], 'authorization': []}]
func (p *PaymentService) GetPaymentOrderChargeStatusShort(input *payment.GetPaymentOrderChargeStatusParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.PaymentOrderChargeStatus, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(p.TokenRepository, nil, security, "")
	}
	ok, err := p.Client.Payment.GetPaymentOrderChargeStatusShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:USER:{userId}:PAYMENT [CREATE]'], 'authorization': []}]
func (p *PaymentService) CreateUserPaymentOrderShort(input *payment.CreateUserPaymentOrderParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.PaymentOrderInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(p.TokenRepository, nil, security, "")
	}
	created, err := p.Client.Payment.CreateUserPaymentOrderShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return created.GetPayload(), nil
}

// [{'authorization': []}, {'HasPermission': ['ADMIN:NAMESPACE:{namespace}:USER:{userId}:PAYMENT [UPDATE]'], 'authorization': []}]
func (p *PaymentService) RefundUserPaymentOrderShort(input *payment.RefundUserPaymentOrderParams, authInfoWriter runtime.ClientAuthInfoWriter) (*platformclientmodels.PaymentOrderInfo, error) {
	if authInfoWriter == nil {
		security := [][]string{
			{"bearer"},
		}
		authInfoWriter = auth.AuthInfoWriter(p.TokenRepository, nil, security, "")
	}
	ok, err := p.Client.Payment.RefundUserPaymentOrderShort(input, authInfoWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}
