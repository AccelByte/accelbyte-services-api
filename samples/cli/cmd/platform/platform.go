// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package platform

import (
	"github.com/AccelByte/sample-apps/cmd/platform/achievementPlatform"
	"github.com/AccelByte/sample-apps/cmd/platform/anonymization"
	"github.com/AccelByte/sample-apps/cmd/platform/campaign"
	"github.com/AccelByte/sample-apps/cmd/platform/catalogChanges"
	"github.com/AccelByte/sample-apps/cmd/platform/category"
	"github.com/AccelByte/sample-apps/cmd/platform/currency"
	"github.com/AccelByte/sample-apps/cmd/platform/dlc"
	"github.com/AccelByte/sample-apps/cmd/platform/entitlement"
	"github.com/AccelByte/sample-apps/cmd/platform/fulfillment"
	"github.com/AccelByte/sample-apps/cmd/platform/fulfillmentScript"
	"github.com/AccelByte/sample-apps/cmd/platform/iap"
	"github.com/AccelByte/sample-apps/cmd/platform/invoice"
	"github.com/AccelByte/sample-apps/cmd/platform/item"
	"github.com/AccelByte/sample-apps/cmd/platform/keyGroup"
	"github.com/AccelByte/sample-apps/cmd/platform/order"
	"github.com/AccelByte/sample-apps/cmd/platform/orderDedicated"
	"github.com/AccelByte/sample-apps/cmd/platform/payment"
	"github.com/AccelByte/sample-apps/cmd/platform/paymentAccount"
	"github.com/AccelByte/sample-apps/cmd/platform/paymentCallbackConfig"
	"github.com/AccelByte/sample-apps/cmd/platform/paymentConfig"
	"github.com/AccelByte/sample-apps/cmd/platform/paymentDedicated"
	"github.com/AccelByte/sample-apps/cmd/platform/paymentStation"
	"github.com/AccelByte/sample-apps/cmd/platform/revocation"
	"github.com/AccelByte/sample-apps/cmd/platform/reward"
	"github.com/AccelByte/sample-apps/cmd/platform/section"
	"github.com/AccelByte/sample-apps/cmd/platform/servicePluginConfig"
	"github.com/AccelByte/sample-apps/cmd/platform/sessionPlatform"
	"github.com/AccelByte/sample-apps/cmd/platform/store"
	"github.com/AccelByte/sample-apps/cmd/platform/subscription"
	"github.com/AccelByte/sample-apps/cmd/platform/ticket"
	"github.com/AccelByte/sample-apps/cmd/platform/view"
	"github.com/AccelByte/sample-apps/cmd/platform/wallet"
	"github.com/spf13/cobra"
)

// PlatformCmd represents the service's child command
var PlatformCmd = &cobra.Command{
	Use:   "Platform",
	Short: "Platform to get all the commands",
}

func init() {
	PlatformCmd.AddCommand(fulfillmentScript.ListFulfillmentScriptsCmd)
	PlatformCmd.AddCommand(fulfillmentScript.TestFulfillmentScriptEvalCmd)
	PlatformCmd.AddCommand(fulfillmentScript.GetFulfillmentScriptCmd)
	PlatformCmd.AddCommand(fulfillmentScript.CreateFulfillmentScriptCmd)
	PlatformCmd.AddCommand(fulfillmentScript.DeleteFulfillmentScriptCmd)
	PlatformCmd.AddCommand(fulfillmentScript.UpdateFulfillmentScriptCmd)
	PlatformCmd.AddCommand(item.ListItemTypeConfigsCmd)
	PlatformCmd.AddCommand(item.CreateItemTypeConfigCmd)
	PlatformCmd.AddCommand(item.SearchItemTypeConfigCmd)
	PlatformCmd.AddCommand(item.GetItemTypeConfigCmd)
	PlatformCmd.AddCommand(item.UpdateItemTypeConfigCmd)
	PlatformCmd.AddCommand(item.DeleteItemTypeConfigCmd)
	PlatformCmd.AddCommand(campaign.QueryCampaignsCmd)
	PlatformCmd.AddCommand(campaign.CreateCampaignCmd)
	PlatformCmd.AddCommand(campaign.GetCampaignCmd)
	PlatformCmd.AddCommand(campaign.UpdateCampaignCmd)
	PlatformCmd.AddCommand(campaign.GetCampaignDynamicCmd)
	PlatformCmd.AddCommand(servicePluginConfig.GetLootBoxPluginConfigCmd)
	PlatformCmd.AddCommand(servicePluginConfig.UpdateLootBoxPluginConfigCmd)
	PlatformCmd.AddCommand(servicePluginConfig.DeleteLootBoxPluginConfigCmd)
	PlatformCmd.AddCommand(servicePluginConfig.UplodLootBoxPluginConfigCertCmd)
	PlatformCmd.AddCommand(servicePluginConfig.GetLootBoxGrpcInfoCmd)
	PlatformCmd.AddCommand(servicePluginConfig.GetSectionPluginConfigCmd)
	PlatformCmd.AddCommand(servicePluginConfig.UpdateSectionPluginConfigCmd)
	PlatformCmd.AddCommand(servicePluginConfig.DeleteSectionPluginConfigCmd)
	PlatformCmd.AddCommand(servicePluginConfig.UploadSectionPluginConfigCertCmd)
	PlatformCmd.AddCommand(category.GetRootCategoriesCmd)
	PlatformCmd.AddCommand(category.CreateCategoryCmd)
	PlatformCmd.AddCommand(category.ListCategoriesBasicCmd)
	PlatformCmd.AddCommand(category.GetCategoryCmd)
	PlatformCmd.AddCommand(category.UpdateCategoryCmd)
	PlatformCmd.AddCommand(category.DeleteCategoryCmd)
	PlatformCmd.AddCommand(category.GetChildCategoriesCmd)
	PlatformCmd.AddCommand(category.GetDescendantCategoriesCmd)
	PlatformCmd.AddCommand(campaign.QueryCodesCmd)
	PlatformCmd.AddCommand(campaign.CreateCodesCmd)
	PlatformCmd.AddCommand(campaign.DownloadCmd)
	PlatformCmd.AddCommand(campaign.BulkDisableCodesCmd)
	PlatformCmd.AddCommand(campaign.BulkEnableCodesCmd)
	PlatformCmd.AddCommand(campaign.QueryRedeemHistoryCmd)
	PlatformCmd.AddCommand(campaign.GetCodeCmd)
	PlatformCmd.AddCommand(campaign.DisableCodeCmd)
	PlatformCmd.AddCommand(campaign.EnableCodeCmd)
	PlatformCmd.AddCommand(servicePluginConfig.GetServicePluginConfigCmd)
	PlatformCmd.AddCommand(servicePluginConfig.UpdateServicePluginConfigCmd)
	PlatformCmd.AddCommand(servicePluginConfig.DeleteServicePluginConfigCmd)
	PlatformCmd.AddCommand(currency.ListCurrenciesCmd)
	PlatformCmd.AddCommand(currency.CreateCurrencyCmd)
	PlatformCmd.AddCommand(currency.UpdateCurrencyCmd)
	PlatformCmd.AddCommand(currency.DeleteCurrencyCmd)
	PlatformCmd.AddCommand(currency.GetCurrencyConfigCmd)
	PlatformCmd.AddCommand(currency.GetCurrencySummaryCmd)
	PlatformCmd.AddCommand(dlc.GetDLCItemConfigCmd)
	PlatformCmd.AddCommand(dlc.UpdateDLCItemConfigCmd)
	PlatformCmd.AddCommand(dlc.DeleteDLCItemConfigCmd)
	PlatformCmd.AddCommand(dlc.GetPlatformDLCConfigCmd)
	PlatformCmd.AddCommand(dlc.UpdatePlatformDLCConfigCmd)
	PlatformCmd.AddCommand(dlc.DeletePlatformDLCConfigCmd)
	PlatformCmd.AddCommand(entitlement.QueryEntitlementsCmd)
	PlatformCmd.AddCommand(entitlement.QueryEntitlements1Cmd)
	PlatformCmd.AddCommand(entitlement.GrantEntitlementsCmd)
	PlatformCmd.AddCommand(entitlement.RevokeEntitlementsCmd)
	PlatformCmd.AddCommand(entitlement.GetEntitlementCmd)
	PlatformCmd.AddCommand(fulfillment.QueryFulfillmentHistoriesCmd)
	PlatformCmd.AddCommand(iap.GetAppleIAPConfigCmd)
	PlatformCmd.AddCommand(iap.UpdateAppleIAPConfigCmd)
	PlatformCmd.AddCommand(iap.DeleteAppleIAPConfigCmd)
	PlatformCmd.AddCommand(iap.GetEpicGamesIAPConfigCmd)
	PlatformCmd.AddCommand(iap.UpdateEpicGamesIAPConfigCmd)
	PlatformCmd.AddCommand(iap.DeleteEpicGamesIAPConfigCmd)
	PlatformCmd.AddCommand(iap.GetGoogleIAPConfigCmd)
	PlatformCmd.AddCommand(iap.UpdateGoogleIAPConfigCmd)
	PlatformCmd.AddCommand(iap.DeleteGoogleIAPConfigCmd)
	PlatformCmd.AddCommand(iap.UpdateGoogleP12FileCmd)
	PlatformCmd.AddCommand(iap.GetIAPItemConfigCmd)
	PlatformCmd.AddCommand(iap.UpdateIAPItemConfigCmd)
	PlatformCmd.AddCommand(iap.DeleteIAPItemConfigCmd)
	PlatformCmd.AddCommand(iap.GetOculusIAPConfigCmd)
	PlatformCmd.AddCommand(iap.UpdateOculusIAPConfigCmd)
	PlatformCmd.AddCommand(iap.DeleteOculusIAPConfigCmd)
	PlatformCmd.AddCommand(iap.GetPlayStationIAPConfigCmd)
	PlatformCmd.AddCommand(iap.UpdatePlaystationIAPConfigCmd)
	PlatformCmd.AddCommand(iap.DeletePlaystationIAPConfigCmd)
	PlatformCmd.AddCommand(iap.GetSteamIAPConfigCmd)
	PlatformCmd.AddCommand(iap.UpdateSteamIAPConfigCmd)
	PlatformCmd.AddCommand(iap.DeleteSteamIAPConfigCmd)
	PlatformCmd.AddCommand(iap.GetTwitchIAPConfigCmd)
	PlatformCmd.AddCommand(iap.UpdateTwitchIAPConfigCmd)
	PlatformCmd.AddCommand(iap.DeleteTwitchIAPConfigCmd)
	PlatformCmd.AddCommand(iap.GetXblIAPConfigCmd)
	PlatformCmd.AddCommand(iap.UpdateXblIAPConfigCmd)
	PlatformCmd.AddCommand(iap.DeleteXblAPConfigCmd)
	PlatformCmd.AddCommand(iap.UpdateXblBPCertFileCmd)
	PlatformCmd.AddCommand(invoice.DownloadInvoiceDetailsCmd)
	PlatformCmd.AddCommand(invoice.GenerateInvoiceSummaryCmd)
	PlatformCmd.AddCommand(item.SyncInGameItemCmd)
	PlatformCmd.AddCommand(item.CreateItemCmd)
	PlatformCmd.AddCommand(item.GetItemByAppIdCmd)
	PlatformCmd.AddCommand(item.QueryItemsCmd)
	PlatformCmd.AddCommand(item.ListBasicItemsByFeaturesCmd)
	PlatformCmd.AddCommand(item.GetItemBySkuCmd)
	PlatformCmd.AddCommand(item.GetLocaleItemBySkuCmd)
	PlatformCmd.AddCommand(item.GetItemIdBySkuCmd)
	PlatformCmd.AddCommand(item.GetBulkItemIdBySkusCmd)
	PlatformCmd.AddCommand(item.BulkGetLocaleItemsCmd)
	PlatformCmd.AddCommand(item.GetAvailablePredicateTypesCmd)
	PlatformCmd.AddCommand(item.ValidateItemPurchaseConditionCmd)
	PlatformCmd.AddCommand(item.BulkUpdateRegionDataCmd)
	PlatformCmd.AddCommand(item.SearchItemsCmd)
	PlatformCmd.AddCommand(item.QueryUncategorizedItemsCmd)
	PlatformCmd.AddCommand(item.GetItemCmd)
	PlatformCmd.AddCommand(item.UpdateItemCmd)
	PlatformCmd.AddCommand(item.DeleteItemCmd)
	PlatformCmd.AddCommand(item.AcquireItemCmd)
	PlatformCmd.AddCommand(item.GetAppCmd)
	PlatformCmd.AddCommand(item.UpdateAppCmd)
	PlatformCmd.AddCommand(item.DisableItemCmd)
	PlatformCmd.AddCommand(item.GetItemDynamicDataCmd)
	PlatformCmd.AddCommand(item.EnableItemCmd)
	PlatformCmd.AddCommand(item.FeatureItemCmd)
	PlatformCmd.AddCommand(item.DefeatureItemCmd)
	PlatformCmd.AddCommand(item.GetLocaleItemCmd)
	PlatformCmd.AddCommand(item.UpdateItemPurchaseConditionCmd)
	PlatformCmd.AddCommand(item.ReturnItemCmd)
	PlatformCmd.AddCommand(keyGroup.QueryKeyGroupsCmd)
	PlatformCmd.AddCommand(keyGroup.CreateKeyGroupCmd)
	PlatformCmd.AddCommand(keyGroup.GetKeyGroupByBoothNameCmd)
	PlatformCmd.AddCommand(keyGroup.GetKeyGroupCmd)
	PlatformCmd.AddCommand(keyGroup.UpdateKeyGroupCmd)
	PlatformCmd.AddCommand(keyGroup.GetKeyGroupDynamicCmd)
	PlatformCmd.AddCommand(keyGroup.ListKeysCmd)
	PlatformCmd.AddCommand(keyGroup.UploadKeysCmd)
	PlatformCmd.AddCommand(order.QueryOrdersCmd)
	PlatformCmd.AddCommand(order.GetOrderStatisticsCmd)
	PlatformCmd.AddCommand(order.GetOrderCmd)
	PlatformCmd.AddCommand(order.RefundOrderCmd)
	PlatformCmd.AddCommand(paymentCallbackConfig.GetPaymentCallbackConfigCmd)
	PlatformCmd.AddCommand(paymentCallbackConfig.UpdatePaymentCallbackConfigCmd)
	PlatformCmd.AddCommand(payment.QueryPaymentNotificationsCmd)
	PlatformCmd.AddCommand(payment.QueryPaymentOrdersCmd)
	PlatformCmd.AddCommand(paymentDedicated.CreatePaymentOrderByDedicatedCmd)
	PlatformCmd.AddCommand(payment.ListExtOrderNoByExtTxIdCmd)
	PlatformCmd.AddCommand(payment.GetPaymentOrderCmd)
	PlatformCmd.AddCommand(payment.ChargePaymentOrderCmd)
	PlatformCmd.AddCommand(paymentDedicated.RefundPaymentOrderByDedicatedCmd)
	PlatformCmd.AddCommand(payment.SimulatePaymentOrderNotificationCmd)
	PlatformCmd.AddCommand(payment.GetPaymentOrderChargeStatusCmd)
	PlatformCmd.AddCommand(wallet.GetPlatformWalletConfigCmd)
	PlatformCmd.AddCommand(wallet.UpdatePlatformWalletConfigCmd)
	PlatformCmd.AddCommand(wallet.ResetPlatformWalletConfigCmd)
	PlatformCmd.AddCommand(revocation.GetRevocationConfigCmd)
	PlatformCmd.AddCommand(revocation.UpdateRevocationConfigCmd)
	PlatformCmd.AddCommand(revocation.DeleteRevocationConfigCmd)
	PlatformCmd.AddCommand(revocation.QueryRevocationHistoriesCmd)
	PlatformCmd.AddCommand(servicePluginConfig.GetLootBoxPluginConfig1Cmd)
	PlatformCmd.AddCommand(servicePluginConfig.UpdateRevocationPluginConfigCmd)
	PlatformCmd.AddCommand(servicePluginConfig.DeleteLootBoxPluginConfig1Cmd)
	PlatformCmd.AddCommand(servicePluginConfig.UploadRevocationPluginConfigCertCmd)
	PlatformCmd.AddCommand(reward.CreateRewardCmd)
	PlatformCmd.AddCommand(reward.QueryRewardsCmd)
	PlatformCmd.AddCommand(reward.ExportRewardsCmd)
	PlatformCmd.AddCommand(reward.ImportRewardsCmd)
	PlatformCmd.AddCommand(reward.GetRewardCmd)
	PlatformCmd.AddCommand(reward.UpdateRewardCmd)
	PlatformCmd.AddCommand(reward.DeleteRewardCmd)
	PlatformCmd.AddCommand(reward.CheckEventConditionCmd)
	PlatformCmd.AddCommand(reward.DeleteRewardConditionRecordCmd)
	PlatformCmd.AddCommand(section.QuerySectionsCmd)
	PlatformCmd.AddCommand(section.CreateSectionCmd)
	PlatformCmd.AddCommand(section.PurgeExpiredSectionCmd)
	PlatformCmd.AddCommand(section.GetSectionCmd)
	PlatformCmd.AddCommand(section.UpdateSectionCmd)
	PlatformCmd.AddCommand(section.DeleteSectionCmd)
	PlatformCmd.AddCommand(store.ListStoresCmd)
	PlatformCmd.AddCommand(store.CreateStoreCmd)
	PlatformCmd.AddCommand(store.ImportStoreCmd)
	PlatformCmd.AddCommand(store.GetPublishedStoreCmd)
	PlatformCmd.AddCommand(store.DeletePublishedStoreCmd)
	PlatformCmd.AddCommand(store.GetPublishedStoreBackupCmd)
	PlatformCmd.AddCommand(store.RollbackPublishedStoreCmd)
	PlatformCmd.AddCommand(store.GetStoreCmd)
	PlatformCmd.AddCommand(store.UpdateStoreCmd)
	PlatformCmd.AddCommand(store.DeleteStoreCmd)
	PlatformCmd.AddCommand(catalogChanges.QueryChangesCmd)
	PlatformCmd.AddCommand(catalogChanges.PublishAllCmd)
	PlatformCmd.AddCommand(catalogChanges.PublishSelectedCmd)
	PlatformCmd.AddCommand(catalogChanges.SelectAllRecordsCmd)
	PlatformCmd.AddCommand(catalogChanges.GetStatisticCmd)
	PlatformCmd.AddCommand(catalogChanges.UnselectAllRecordsCmd)
	PlatformCmd.AddCommand(catalogChanges.SelectRecordCmd)
	PlatformCmd.AddCommand(catalogChanges.UnselectRecordCmd)
	PlatformCmd.AddCommand(store.CloneStoreCmd)
	PlatformCmd.AddCommand(store.ExportStoreCmd)
	PlatformCmd.AddCommand(subscription.QuerySubscriptionsCmd)
	PlatformCmd.AddCommand(subscription.RecurringChargeSubscriptionCmd)
	PlatformCmd.AddCommand(ticket.GetTicketDynamicCmd)
	PlatformCmd.AddCommand(ticket.DecreaseTicketSaleCmd)
	PlatformCmd.AddCommand(ticket.GetTicketBoothIDCmd)
	PlatformCmd.AddCommand(ticket.IncreaseTicketSaleCmd)
	PlatformCmd.AddCommand(achievementPlatform.UnlockSteamUserAchievementCmd)
	PlatformCmd.AddCommand(achievementPlatform.GetXblUserAchievementsCmd)
	PlatformCmd.AddCommand(achievementPlatform.UpdateXblUserAchievementCmd)
	PlatformCmd.AddCommand(anonymization.AnonymizeCampaignCmd)
	PlatformCmd.AddCommand(anonymization.AnonymizeEntitlementCmd)
	PlatformCmd.AddCommand(anonymization.AnonymizeFulfillmentCmd)
	PlatformCmd.AddCommand(anonymization.AnonymizeIntegrationCmd)
	PlatformCmd.AddCommand(anonymization.AnonymizeOrderCmd)
	PlatformCmd.AddCommand(anonymization.AnonymizePaymentCmd)
	PlatformCmd.AddCommand(anonymization.AnonymizeRevocationCmd)
	PlatformCmd.AddCommand(anonymization.AnonymizeSubscriptionCmd)
	PlatformCmd.AddCommand(anonymization.AnonymizeWalletCmd)
	PlatformCmd.AddCommand(dlc.GetUserDLCByPlatformCmd)
	PlatformCmd.AddCommand(dlc.GetUserDLCCmd)
	PlatformCmd.AddCommand(entitlement.QueryUserEntitlementsCmd)
	PlatformCmd.AddCommand(entitlement.GrantUserEntitlementCmd)
	PlatformCmd.AddCommand(entitlement.GetUserAppEntitlementByAppIdCmd)
	PlatformCmd.AddCommand(entitlement.QueryUserEntitlementsByAppTypeCmd)
	PlatformCmd.AddCommand(entitlement.GetUserEntitlementByItemIdCmd)
	PlatformCmd.AddCommand(entitlement.GetUserActiveEntitlementsByItemIdsCmd)
	PlatformCmd.AddCommand(entitlement.GetUserEntitlementBySkuCmd)
	PlatformCmd.AddCommand(entitlement.ExistsAnyUserActiveEntitlementCmd)
	PlatformCmd.AddCommand(entitlement.ExistsAnyUserActiveEntitlementByItemIdsCmd)
	PlatformCmd.AddCommand(entitlement.GetUserAppEntitlementOwnershipByAppIdCmd)
	PlatformCmd.AddCommand(entitlement.GetUserEntitlementOwnershipByItemIdCmd)
	PlatformCmd.AddCommand(entitlement.GetUserEntitlementOwnershipByItemIdsCmd)
	PlatformCmd.AddCommand(entitlement.GetUserEntitlementOwnershipBySkuCmd)
	PlatformCmd.AddCommand(entitlement.RevokeAllEntitlementsCmd)
	PlatformCmd.AddCommand(entitlement.RevokeUserEntitlementsCmd)
	PlatformCmd.AddCommand(entitlement.GetUserEntitlementCmd)
	PlatformCmd.AddCommand(entitlement.UpdateUserEntitlementCmd)
	PlatformCmd.AddCommand(entitlement.ConsumeUserEntitlementCmd)
	PlatformCmd.AddCommand(entitlement.DisableUserEntitlementCmd)
	PlatformCmd.AddCommand(entitlement.EnableUserEntitlementCmd)
	PlatformCmd.AddCommand(entitlement.GetUserEntitlementHistoriesCmd)
	PlatformCmd.AddCommand(entitlement.RevokeUserEntitlementCmd)
	PlatformCmd.AddCommand(entitlement.RevokeUseCountCmd)
	PlatformCmd.AddCommand(entitlement.SellUserEntitlementCmd)
	PlatformCmd.AddCommand(fulfillment.FulfillItemCmd)
	PlatformCmd.AddCommand(fulfillment.RedeemCodeCmd)
	PlatformCmd.AddCommand(fulfillment.FulfillRewardsCmd)
	PlatformCmd.AddCommand(iap.QueryUserIAPOrdersCmd)
	PlatformCmd.AddCommand(iap.QueryAllUserIAPOrdersCmd)
	PlatformCmd.AddCommand(iap.QueryUserIAPConsumeHistoryCmd)
	PlatformCmd.AddCommand(iap.MockFulfillIAPItemCmd)
	PlatformCmd.AddCommand(order.QueryUserOrdersCmd)
	PlatformCmd.AddCommand(order.AdminCreateUserOrderCmd)
	PlatformCmd.AddCommand(order.CountOfPurchasedItemCmd)
	PlatformCmd.AddCommand(order.GetUserOrderCmd)
	PlatformCmd.AddCommand(order.UpdateUserOrderStatusCmd)
	PlatformCmd.AddCommand(order.FulfillUserOrderCmd)
	PlatformCmd.AddCommand(order.GetUserOrderGrantCmd)
	PlatformCmd.AddCommand(order.GetUserOrderHistoriesCmd)
	PlatformCmd.AddCommand(order.ProcessUserOrderNotificationCmd)
	PlatformCmd.AddCommand(order.DownloadUserOrderReceiptCmd)
	PlatformCmd.AddCommand(payment.CreateUserPaymentOrderCmd)
	PlatformCmd.AddCommand(payment.RefundUserPaymentOrderCmd)
	PlatformCmd.AddCommand(campaign.ApplyUserRedemptionCmd)
	PlatformCmd.AddCommand(revocation.DoRevocationCmd)
	PlatformCmd.AddCommand(sessionPlatform.RegisterXblSessionsCmd)
	PlatformCmd.AddCommand(subscription.QueryUserSubscriptionsCmd)
	PlatformCmd.AddCommand(subscription.GetUserSubscriptionActivitiesCmd)
	PlatformCmd.AddCommand(subscription.PlatformSubscribeSubscriptionCmd)
	PlatformCmd.AddCommand(subscription.CheckUserSubscriptionSubscribableByItemIdCmd)
	PlatformCmd.AddCommand(subscription.GetUserSubscriptionCmd)
	PlatformCmd.AddCommand(subscription.DeleteUserSubscriptionCmd)
	PlatformCmd.AddCommand(subscription.CancelSubscriptionCmd)
	PlatformCmd.AddCommand(subscription.GrantDaysToSubscriptionCmd)
	PlatformCmd.AddCommand(subscription.GetUserSubscriptionBillingHistoriesCmd)
	PlatformCmd.AddCommand(subscription.ProcessUserSubscriptionNotificationCmd)
	PlatformCmd.AddCommand(ticket.AcquireUserTicketCmd)
	PlatformCmd.AddCommand(wallet.QueryUserCurrencyWalletsCmd)
	PlatformCmd.AddCommand(wallet.DebitUserWalletByCurrencyCodeCmd)
	PlatformCmd.AddCommand(wallet.ListUserCurrencyTransactionsCmd)
	PlatformCmd.AddCommand(wallet.CheckWalletCmd)
	PlatformCmd.AddCommand(wallet.CreditUserWalletCmd)
	PlatformCmd.AddCommand(wallet.PayWithUserWalletCmd)
	PlatformCmd.AddCommand(wallet.GetUserWalletCmd)
	PlatformCmd.AddCommand(wallet.DebitUserWalletCmd)
	PlatformCmd.AddCommand(wallet.DisableUserWalletCmd)
	PlatformCmd.AddCommand(wallet.EnableUserWalletCmd)
	PlatformCmd.AddCommand(wallet.ListUserWalletTransactionsCmd)
	PlatformCmd.AddCommand(view.ListViewsCmd)
	PlatformCmd.AddCommand(view.CreateViewCmd)
	PlatformCmd.AddCommand(view.GetViewCmd)
	PlatformCmd.AddCommand(view.UpdateViewCmd)
	PlatformCmd.AddCommand(view.DeleteViewCmd)
	PlatformCmd.AddCommand(wallet.QueryWalletsCmd)
	PlatformCmd.AddCommand(wallet.BulkCreditCmd)
	PlatformCmd.AddCommand(wallet.BulkDebitCmd)
	PlatformCmd.AddCommand(wallet.GetWalletCmd)
	PlatformCmd.AddCommand(orderDedicated.SyncOrdersCmd)
	PlatformCmd.AddCommand(paymentConfig.TestAdyenConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.TestAliPayConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.TestCheckoutConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.DebugMatchedPaymentMerchantConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.TestPayPalConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.TestStripeConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.TestWxPayConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.TestXsollaConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.GetPaymentMerchantConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.UpdateAdyenConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.TestAdyenConfigByIdCmd)
	PlatformCmd.AddCommand(paymentConfig.UpdateAliPayConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.TestAliPayConfigByIdCmd)
	PlatformCmd.AddCommand(paymentConfig.UpdateCheckoutConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.TestCheckoutConfigByIdCmd)
	PlatformCmd.AddCommand(paymentConfig.UpdatePayPalConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.TestPayPalConfigByIdCmd)
	PlatformCmd.AddCommand(paymentConfig.UpdateStripeConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.TestStripeConfigByIdCmd)
	PlatformCmd.AddCommand(paymentConfig.UpdateWxPayConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.UpdateWxPayConfigCertCmd)
	PlatformCmd.AddCommand(paymentConfig.TestWxPayConfigByIdCmd)
	PlatformCmd.AddCommand(paymentConfig.UpdateXsollaConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.TestXsollaConfigByIdCmd)
	PlatformCmd.AddCommand(paymentConfig.UpdateXsollaUIConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.QueryPaymentProviderConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.CreatePaymentProviderConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.GetAggregatePaymentProvidersCmd)
	PlatformCmd.AddCommand(paymentConfig.DebugMatchedPaymentProviderConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.GetSpecialPaymentProvidersCmd)
	PlatformCmd.AddCommand(paymentConfig.UpdatePaymentProviderConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.DeletePaymentProviderConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.GetPaymentTaxConfigCmd)
	PlatformCmd.AddCommand(paymentConfig.UpdatePaymentTaxConfigCmd)
	PlatformCmd.AddCommand(paymentDedicated.SyncPaymentOrdersCmd)
	PlatformCmd.AddCommand(category.PublicGetRootCategoriesCmd)
	PlatformCmd.AddCommand(category.DownloadCategoriesCmd)
	PlatformCmd.AddCommand(category.PublicGetCategoryCmd)
	PlatformCmd.AddCommand(category.PublicGetChildCategoriesCmd)
	PlatformCmd.AddCommand(category.PublicGetDescendantCategoriesCmd)
	PlatformCmd.AddCommand(currency.PublicListCurrenciesCmd)
	PlatformCmd.AddCommand(iap.GetIAPItemMappingCmd)
	PlatformCmd.AddCommand(item.PublicGetItemByAppIdCmd)
	PlatformCmd.AddCommand(item.PublicQueryItemsCmd)
	PlatformCmd.AddCommand(item.PublicGetItemBySkuCmd)
	PlatformCmd.AddCommand(item.PublicBulkGetItemsCmd)
	PlatformCmd.AddCommand(item.PublicValidateItemPurchaseConditionCmd)
	PlatformCmd.AddCommand(item.PublicSearchItemsCmd)
	PlatformCmd.AddCommand(item.PublicGetAppCmd)
	PlatformCmd.AddCommand(item.PublicGetItemDynamicDataCmd)
	PlatformCmd.AddCommand(item.PublicGetItemCmd)
	PlatformCmd.AddCommand(paymentStation.GetPaymentCustomizationCmd)
	PlatformCmd.AddCommand(paymentStation.PublicGetPaymentUrlCmd)
	PlatformCmd.AddCommand(paymentStation.PublicGetPaymentMethodsCmd)
	PlatformCmd.AddCommand(paymentStation.PublicGetUnpaidPaymentOrderCmd)
	PlatformCmd.AddCommand(paymentStation.PayCmd)
	PlatformCmd.AddCommand(paymentStation.PublicCheckPaymentOrderPaidStatusCmd)
	PlatformCmd.AddCommand(paymentStation.GetPaymentPublicConfigCmd)
	PlatformCmd.AddCommand(paymentStation.PublicGetQRCodeCmd)
	PlatformCmd.AddCommand(paymentStation.PublicNormalizePaymentReturnUrlCmd)
	PlatformCmd.AddCommand(paymentStation.GetPaymentTaxValueCmd)
	PlatformCmd.AddCommand(reward.GetRewardByCodeCmd)
	PlatformCmd.AddCommand(reward.QueryRewards1Cmd)
	PlatformCmd.AddCommand(reward.GetReward1Cmd)
	PlatformCmd.AddCommand(store.PublicListStoresCmd)
	PlatformCmd.AddCommand(entitlement.PublicExistsAnyMyActiveEntitlementCmd)
	PlatformCmd.AddCommand(entitlement.PublicGetMyAppEntitlementOwnershipByAppIdCmd)
	PlatformCmd.AddCommand(entitlement.PublicGetMyEntitlementOwnershipByItemIdCmd)
	PlatformCmd.AddCommand(entitlement.PublicGetMyEntitlementOwnershipBySkuCmd)
	PlatformCmd.AddCommand(entitlement.PublicGetEntitlementOwnershipTokenCmd)
	PlatformCmd.AddCommand(iap.SyncTwitchDropsEntitlementCmd)
	PlatformCmd.AddCommand(wallet.PublicGetMyWalletCmd)
	PlatformCmd.AddCommand(dlc.SyncEpicGameDLCCmd)
	PlatformCmd.AddCommand(dlc.SyncOculusDLCCmd)
	PlatformCmd.AddCommand(dlc.PublicSyncPsnDlcInventoryCmd)
	PlatformCmd.AddCommand(dlc.PublicSyncPsnDlcInventoryWithMultipleServiceLabelsCmd)
	PlatformCmd.AddCommand(dlc.SyncSteamDLCCmd)
	PlatformCmd.AddCommand(dlc.SyncXboxDLCCmd)
	PlatformCmd.AddCommand(entitlement.PublicQueryUserEntitlementsCmd)
	PlatformCmd.AddCommand(entitlement.PublicGetUserAppEntitlementByAppIdCmd)
	PlatformCmd.AddCommand(entitlement.PublicQueryUserEntitlementsByAppTypeCmd)
	PlatformCmd.AddCommand(entitlement.PublicGetUserEntitlementByItemIdCmd)
	PlatformCmd.AddCommand(entitlement.PublicGetUserEntitlementBySkuCmd)
	PlatformCmd.AddCommand(entitlement.PublicExistsAnyUserActiveEntitlementCmd)
	PlatformCmd.AddCommand(entitlement.PublicGetUserAppEntitlementOwnershipByAppIdCmd)
	PlatformCmd.AddCommand(entitlement.PublicGetUserEntitlementOwnershipByItemIdCmd)
	PlatformCmd.AddCommand(entitlement.PublicGetUserEntitlementOwnershipByItemIdsCmd)
	PlatformCmd.AddCommand(entitlement.PublicGetUserEntitlementOwnershipBySkuCmd)
	PlatformCmd.AddCommand(entitlement.PublicGetUserEntitlementCmd)
	PlatformCmd.AddCommand(entitlement.PublicConsumeUserEntitlementCmd)
	PlatformCmd.AddCommand(entitlement.PublicSellUserEntitlementCmd)
	PlatformCmd.AddCommand(fulfillment.PublicRedeemCodeCmd)
	PlatformCmd.AddCommand(iap.PublicFulfillAppleIAPItemCmd)
	PlatformCmd.AddCommand(iap.SyncEpicGamesInventoryCmd)
	PlatformCmd.AddCommand(iap.PublicFulfillGoogleIAPItemCmd)
	PlatformCmd.AddCommand(iap.SyncOculusConsumableEntitlementsCmd)
	PlatformCmd.AddCommand(iap.PublicReconcilePlayStationStoreCmd)
	PlatformCmd.AddCommand(iap.PublicReconcilePlayStationStoreWithMultipleServiceLabelsCmd)
	PlatformCmd.AddCommand(iap.SyncSteamInventoryCmd)
	PlatformCmd.AddCommand(iap.SyncTwitchDropsEntitlement1Cmd)
	PlatformCmd.AddCommand(iap.SyncXboxInventoryCmd)
	PlatformCmd.AddCommand(order.PublicQueryUserOrdersCmd)
	PlatformCmd.AddCommand(order.PublicCreateUserOrderCmd)
	PlatformCmd.AddCommand(order.PublicGetUserOrderCmd)
	PlatformCmd.AddCommand(order.PublicCancelUserOrderCmd)
	PlatformCmd.AddCommand(order.PublicGetUserOrderHistoriesCmd)
	PlatformCmd.AddCommand(order.PublicDownloadUserOrderReceiptCmd)
	PlatformCmd.AddCommand(paymentAccount.PublicGetPaymentAccountsCmd)
	PlatformCmd.AddCommand(paymentAccount.PublicDeletePaymentAccountCmd)
	PlatformCmd.AddCommand(section.PublicListActiveSectionsCmd)
	PlatformCmd.AddCommand(subscription.PublicQueryUserSubscriptionsCmd)
	PlatformCmd.AddCommand(subscription.PublicSubscribeSubscriptionCmd)
	PlatformCmd.AddCommand(subscription.PublicCheckUserSubscriptionSubscribableByItemIdCmd)
	PlatformCmd.AddCommand(subscription.PublicGetUserSubscriptionCmd)
	PlatformCmd.AddCommand(subscription.PublicChangeSubscriptionBillingAccountCmd)
	PlatformCmd.AddCommand(subscription.PublicCancelSubscriptionCmd)
	PlatformCmd.AddCommand(subscription.PublicGetUserSubscriptionBillingHistoriesCmd)
	PlatformCmd.AddCommand(view.PublicListViewsCmd)
	PlatformCmd.AddCommand(wallet.PublicGetWalletCmd)
	PlatformCmd.AddCommand(wallet.PublicListUserWalletTransactionsCmd)
	PlatformCmd.AddCommand(item.QueryItems1Cmd)
	PlatformCmd.AddCommand(store.ImportStore1Cmd)
	PlatformCmd.AddCommand(store.ExportStore1Cmd)
}
