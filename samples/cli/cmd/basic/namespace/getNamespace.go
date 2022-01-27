// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package namespace

import (
	namespace_ "github.com/AccelByte/accelbyte-go-sdk/basic-sdk/pkg/basicclient/namespace"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/basic"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getNamespaceCmd represents the getNamespace command
var getNamespaceCmd = &cobra.Command{
	Use:   "getNamespace",
	Short: "Get namespace",
	Long:  `Get namespace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		namespaceService := &basic.NamespaceService{
			Client:          factory.NewBasicClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		activeOnly, _ := cmd.Flags().GetBool("activeOnly")
		input := &namespace_.GetNamespaceParams{
			Namespace:  namespace,
			ActiveOnly: &activeOnly,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := namespaceService.GetNamespace(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(getNamespaceCmd)
	getNamespaceCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = getNamespaceCmd.MarkFlagRequired("namespace")
	getNamespaceCmd.Flags().BoolP("activeOnly", "ay", false, "Active only")
}
