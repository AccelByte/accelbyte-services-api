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

// AdminUpdateInventoryCmd represents the AdminUpdateInventory command
var AdminUpdateInventoryCmd = &cobra.Command{
	Use:   "adminUpdateInventory",
	Short: "Admin update inventory",
	Long:  `Admin update inventory`,
	RunE: func(cmd *cobra.Command, args []string) error {
		adminInventoriesService := &inventory.AdminInventoriesService{
			Client:          factory.NewInventoryClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		bodyString := cmd.Flag("body").Value.String()
		var body *inventoryclientmodels.ApimodelsUpdateInventoryReq
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		inventoryId, _ := cmd.Flags().GetString("inventoryId")
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &admin_inventories.AdminUpdateInventoryParams{
			Body:        body,
			InventoryID: inventoryId,
			Namespace:   namespace,
		}
		ok, errOK := adminInventoriesService.AdminUpdateInventoryShort(input)
		if errOK != nil {
			logrus.Error(errOK)

			return errOK
		}

		logrus.Infof("Response CLI success: %+v", ok)

		return nil
	},
}

func init() {
	AdminUpdateInventoryCmd.Flags().String("body", "", "Body")
	_ = AdminUpdateInventoryCmd.MarkFlagRequired("body")
	AdminUpdateInventoryCmd.Flags().String("inventoryId", "", "Inventory id")
	_ = AdminUpdateInventoryCmd.MarkFlagRequired("inventoryId")
	AdminUpdateInventoryCmd.Flags().String("namespace", "", "Namespace")
	_ = AdminUpdateInventoryCmd.MarkFlagRequired("namespace")
}
