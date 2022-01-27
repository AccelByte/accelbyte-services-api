// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package deploymentConfig

import (
	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclient/deployment_config"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/dsmc"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// deleteDeploymentCmd represents the deleteDeployment command
var deleteDeploymentCmd = &cobra.Command{
	Use:   "deleteDeployment",
	Short: "Delete deployment",
	Long:  `Delete deployment`,
	RunE: func(cmd *cobra.Command, args []string) error {
		deploymentConfigService := &dsmc.DeploymentConfigService{
			Client:          factory.NewDsmcClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		deployment, _ := cmd.Flags().GetString("deployment")
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &deployment_config.DeleteDeploymentParams{
			Deployment: deployment,
			Namespace:  namespace,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := deploymentConfigService.DeleteDeployment(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(deleteDeploymentCmd)
	deleteDeploymentCmd.Flags().StringP("deployment", "dt", " ", "Deployment")
	_ = deleteDeploymentCmd.MarkFlagRequired("deployment")
	deleteDeploymentCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = deleteDeploymentCmd.MarkFlagRequired("namespace")
}
