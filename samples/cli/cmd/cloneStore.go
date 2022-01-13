// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.
package cmd

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/store"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloneStoreCmd represents the cloneStore command
var cloneStoreCmd = &cobra.Command{
	Use:   "cloneStore",
	Short: "Clone one store to another store. Source store and target store should in same namespace, same region and same language.",
	Long:  `Clone one store to another store. Source store and target store should in same namespace, same region and same language.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		storeService := &platform.StoreService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace := cmd.Flag("namespace").Value.String()
		storeId := cmd.Flag("sourceStoreId").Value.String()
		targetStoreId := cmd.Flag("targetStoreId").Value.String()
		input := &store.CloneStoreParams{
			Namespace:     namespace,
			StoreID:       storeId,
			TargetStoreID: &targetStoreId,
		}
		//nolint:staticcheck // SA1019 To be deprecated later
		//lint:ignore SA1019 Ignore the deprecation warnings
		storeInfo, err := storeService.CloneStore(input)
		if err != nil {
			return err
		}
		response, err := json.Marshal(storeInfo)
		if err != nil {
			return err
		}
		logrus.Infof("Response %s", response)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(cloneStoreCmd)
	cloneStoreCmd.Flags().StringP("namespace", "n", "", "Store namespace")
	_ = cloneStoreCmd.MarkFlagRequired("namespace")
	cloneStoreCmd.Flags().StringP("sourceStoreId", "s", "", "Source store UD")
	_ = cloneStoreCmd.MarkFlagRequired("sourceStoreId")
	cloneStoreCmd.Flags().StringP("targetStoreId", "t", "", "Target store ID")
	_ = cloneStoreCmd.MarkFlagRequired("targetStoreId")
}
