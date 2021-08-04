// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package cmd

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getClientBynamespaceById represents the getClientBynamespaceById command
var getClientBynamespaceById = &cobra.Command{
	Use:   "getClientByNamespaceById",
	Short: "Admin Get client by namespace by Id",
	Long:  `Admin Get client by namespace by Id`,
	RunE: func(cmd *cobra.Command, args []string) error {
		clientService := &service.ClientService{
			IamClient:       factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		clientID := cmd.Flag("clientID").Value.String()
		namespace := cmd.Flag("namespace").Value.String()
		ok, err := clientService.AdminGetClientsbyNamespacebyIDV3(clientID, namespace)
		if err != nil {
			logrus.Error(err)
		} else {
			response, err := json.Marshal(ok)
			if err != nil {
				logrus.Error(err)
			} else {
				logrus.Infof("Response %s", response)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getClientBynamespaceById)
	getClientBynamespaceById.Flags().StringP("clientID", "c", "", "Client ID")
	getClientBynamespaceById.MarkFlagRequired("clientID")
	getClientBynamespaceById.Flags().StringP("namespace", "n", "", "Client Namespace")
	getClientBynamespaceById.MarkFlagRequired("namespace")
}
