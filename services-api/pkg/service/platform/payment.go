// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package platform

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/payment"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type PaymentService struct {
	Client          *platformclient.JusticePlatformService
	TokenRepository repository.TokenRepository
}

func (p *PaymentService) QueryPaymentNotifications(input *payment.QueryPaymentNotificationsParams)(*platformclientmodels.PaymentNotificationPagingSlicedResult, error) {
    accessToken, err := p.TokenRepository.GetToken()
    if err != nil {
		return nil, err
	}
ok, err := p.Client.Payment.QueryPaymentNotifications(input, client.BearerToken(*accessToken.AccessToken))
    if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PaymentService) QueryPaymentOrders(input *payment.QueryPaymentOrdersParams)(*platformclientmodels.PaymentOrderPagingSlicedResult, error) {
    accessToken, err := p.TokenRepository.GetToken()
    if err != nil {
		return nil, err
	}
ok, err := p.Client.Payment.QueryPaymentOrders(input, client.BearerToken(*accessToken.AccessToken))
    if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PaymentService) ListExtOrderNoByExtTxID(input *payment.ListExtOrderNoByExtTxIDParams)([]string, error) {
    accessToken, err := p.TokenRepository.GetToken()
    if err != nil {
		return nil, err
	}
ok, err := p.Client.Payment.ListExtOrderNoByExtTxID(input, client.BearerToken(*accessToken.AccessToken))
    if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PaymentService) GetPaymentOrder(input *payment.GetPaymentOrderParams)(*platformclientmodels.PaymentOrderInfo, error) {
    accessToken, err := p.TokenRepository.GetToken()
    if err != nil {
		return nil, err
	}
ok, notFound, err := p.Client.Payment.GetPaymentOrder(input, client.BearerToken(*accessToken.AccessToken))
    if notFound != nil {
		return nil, notFound
    }
    if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PaymentService) ChargePaymentOrder(input *payment.ChargePaymentOrderParams)(*platformclientmodels.PaymentOrderInfo, error) {
    accessToken, err := p.TokenRepository.GetToken()
    if err != nil {
		return nil, err
	}
ok, badRequest, notFound, conflict, err := p.Client.Payment.ChargePaymentOrder(input, client.BearerToken(*accessToken.AccessToken))
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

func (p *PaymentService) SimulatePaymentOrderNotification(input *payment.SimulatePaymentOrderNotificationParams)(*platformclientmodels.NotificationProcessResult, error) {
    accessToken, err := p.TokenRepository.GetToken()
    if err != nil {
		return nil, err
	}
ok, badRequest, notFound, err := p.Client.Payment.SimulatePaymentOrderNotification(input, client.BearerToken(*accessToken.AccessToken))
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

func (p *PaymentService) GetPaymentOrderChargeStatus(input *payment.GetPaymentOrderChargeStatusParams)(*platformclientmodels.PaymentOrderChargeStatus, error) {
    accessToken, err := p.TokenRepository.GetToken()
    if err != nil {
		return nil, err
	}
ok, notFound, err := p.Client.Payment.GetPaymentOrderChargeStatus(input, client.BearerToken(*accessToken.AccessToken))
    if notFound != nil {
		return nil, notFound
    }
    if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PaymentService) CreateUserPaymentOrder(input *payment.CreateUserPaymentOrderParams)(*platformclientmodels.PaymentOrderInfo, error) {
    accessToken, err := p.TokenRepository.GetToken()
    if err != nil {
		return nil, err
	}
created, badRequest, forbidden, notFound, conflict, unprocessableEntity, err := p.Client.Payment.CreateUserPaymentOrder(input, client.BearerToken(*accessToken.AccessToken))
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

func (p *PaymentService) RefundUserPaymentOrder(input *payment.RefundUserPaymentOrderParams)(*platformclientmodels.PaymentOrderInfo, error) {
    accessToken, err := p.TokenRepository.GetToken()
    if err != nil {
		return nil, err
	}
ok, notFound, conflict, unprocessableEntity, err := p.Client.Payment.RefundUserPaymentOrder(input, client.BearerToken(*accessToken.AccessToken))
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

func (p *PaymentService) QueryPaymentNotificationsShort(input *payment.QueryPaymentNotificationsParams, authInfo runtime.ClientAuthInfoWriter)(*platformclientmodels.PaymentNotificationPagingSlicedResult, error) {
    ok, err := p.Client.Payment.QueryPaymentNotificationsShort(input, authInfo)
    if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PaymentService) QueryPaymentOrdersShort(input *payment.QueryPaymentOrdersParams, authInfo runtime.ClientAuthInfoWriter)(*platformclientmodels.PaymentOrderPagingSlicedResult, error) {
    ok, err := p.Client.Payment.QueryPaymentOrdersShort(input, authInfo)
    if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PaymentService) ListExtOrderNoByExtTxIDShort(input *payment.ListExtOrderNoByExtTxIDParams, authInfo runtime.ClientAuthInfoWriter)([]string, error) {
    ok, err := p.Client.Payment.ListExtOrderNoByExtTxIDShort(input, authInfo)
    if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PaymentService) GetPaymentOrderShort(input *payment.GetPaymentOrderParams, authInfo runtime.ClientAuthInfoWriter)(*platformclientmodels.PaymentOrderInfo, error) {
    ok, err := p.Client.Payment.GetPaymentOrderShort(input, authInfo)
    if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PaymentService) ChargePaymentOrderShort(input *payment.ChargePaymentOrderParams, authInfo runtime.ClientAuthInfoWriter)(*platformclientmodels.PaymentOrderInfo, error) {
    ok, err := p.Client.Payment.ChargePaymentOrderShort(input, authInfo)
    if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PaymentService) SimulatePaymentOrderNotificationShort(input *payment.SimulatePaymentOrderNotificationParams, authInfo runtime.ClientAuthInfoWriter)(*platformclientmodels.NotificationProcessResult, error) {
    ok, err := p.Client.Payment.SimulatePaymentOrderNotificationShort(input, authInfo)
    if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PaymentService) GetPaymentOrderChargeStatusShort(input *payment.GetPaymentOrderChargeStatusParams, authInfo runtime.ClientAuthInfoWriter)(*platformclientmodels.PaymentOrderChargeStatus, error) {
    ok, err := p.Client.Payment.GetPaymentOrderChargeStatusShort(input, authInfo)
    if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (p *PaymentService) CreateUserPaymentOrderShort(input *payment.CreateUserPaymentOrderParams, authInfo runtime.ClientAuthInfoWriter)(*platformclientmodels.PaymentOrderInfo, error) {
    created, err := p.Client.Payment.CreateUserPaymentOrderShort(input, authInfo)
    if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (p *PaymentService) RefundUserPaymentOrderShort(input *payment.RefundUserPaymentOrderParams, authInfo runtime.ClientAuthInfoWriter)(*platformclientmodels.PaymentOrderInfo, error) {
    ok, err := p.Client.Payment.RefundUserPaymentOrderShort(input, authInfo)
    if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

