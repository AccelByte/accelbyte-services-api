// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package userAction

import (
	"github.com/spf13/cobra"
)

// ReportUserCmd represents the ReportUser command
var ReportUserCmd = &cobra.Command{
	Use:   "reportUser",
	Short: "Report user",
	Long:  `Report user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//userActionService := &basic.UserActionService{
		//	Client:          factory.NewBasicClient(&repository.ConfigRepositoryImpl{}),
		//	TokenRepository: &repository.TokenRepositoryImpl{},
		//}
		//namespace, _ := cmd.Flags().GetString("namespace")
		//bodyString := cmd.Flag("body").Value.String()
		//var body *basicclientmodels.UserReportRequest
		//errBody := json.Unmarshal([]byte(bodyString), &body)
		//if errBody != nil {
		//	return errBody
		//}
		//input := &user_action.ReportUserParams{
		//	Body:      body,
		//	Namespace: namespace,
		//}
		////lint:ignore SA1019 Ignore the deprecation warnings
		//ok, err := userActionService.ReportUser(input)
		//if err != nil {
		//	logrus.Error(err)
		//	return err
		//} else {
		//	response, errIndent := json.MarshalIndent(ok, "", "    ")
		//	if errIndent != nil {
		//		return errIndent
		//	} else {
		//		logrus.Infof("Response %s", string(response))
		//	}
		//}
		return nil
	},
}

func init() {
	//ReportUserCmd.Flags().StringP("body", "", "", "Body")
	//ReportUserCmd.Flags().StringP("namespace", "", "", "Namespace")
	//_ = ReportUserCmd.MarkFlagRequired("namespace")
}
