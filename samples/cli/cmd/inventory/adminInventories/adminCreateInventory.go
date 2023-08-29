// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package adminInventories

import (
	"encoding/json"

	"github.com/AccelByte/accelbyte-go-sdk/inventory-sdk/pkg/inventoryclient/admin_inventories"
	"github.com/AccelByte/accelbyte-go-sdk/inventory-sdk/pkg/inventoryclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/inventory"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AdminCreateInventoryCmd represents the AdminCreateInventory command
var AdminCreateInventoryCmd = &cobra.Command{
	Use:   "adminCreateInventory",
	Short: "Admin create inventory",
	Long:  `Admin create inventory`,
	RunE: func(cmd *cobra.Command, args []string) error {
		adminInventoriesService := &inventory.AdminInventoriesService{
			Client:          factory.NewInventoryClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		bodyString := cmd.Flag("body").Value.String()
		var body *inventoryclientmodels.ApimodelsCreateInventoryReq
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &admin_inventories.AdminCreateInventoryParams{
			Body:      body,
			Namespace: namespace,
		}
		created, errCreated := adminInventoriesService.AdminCreateInventoryShort(input)
		if errCreated != nil {
			logrus.Error(errCreated)

			return errCreated
		}

		logrus.Infof("Response CLI success: %+v", created)

		return nil
	},
}

func init() {
	AdminCreateInventoryCmd.Flags().String("body", "", "Body")
	_ = AdminCreateInventoryCmd.MarkFlagRequired("body")
	AdminCreateInventoryCmd.Flags().String("namespace", "", "Namespace")
	_ = AdminCreateInventoryCmd.MarkFlagRequired("namespace")
}
