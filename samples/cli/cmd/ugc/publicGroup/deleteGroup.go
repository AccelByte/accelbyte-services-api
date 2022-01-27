// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package publicGroup

import (
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/ugc"
	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclient/public_group"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// deleteGroupCmd represents the deleteGroup command
var deleteGroupCmd = &cobra.Command{
	Use:   "deleteGroup",
	Short: "Delete group",
	Long:  `Delete group`,
	RunE: func(cmd *cobra.Command, args []string) error {
		publicGroupService := &ugc.PublicGroupService{
			Client:          factory.NewUgcClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		groupId, _ := cmd.Flags().GetString("groupId")
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		input := &public_group.DeleteGroupParams{
			GroupID:   groupId,
			Namespace: namespace,
			UserID:    userId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := publicGroupService.DeleteGroup(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(deleteGroupCmd)
	deleteGroupCmd.Flags().StringP("groupId", "gd", " ", "Group id")
	_ = deleteGroupCmd.MarkFlagRequired("groupId")
	deleteGroupCmd.Flags().StringP("namespace", "ne", " ", "Namespace")
	_ = deleteGroupCmd.MarkFlagRequired("namespace")
	deleteGroupCmd.Flags().StringP("userId", "ud", " ", "User id")
	_ = deleteGroupCmd.MarkFlagRequired("userId")
}
