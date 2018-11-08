package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tonycai653/qshell/iqshell"
	"os"
)

func init() {
	RootCmd.AddCommand(cmdAccount)
}

var cmdAccount = &cobra.Command{
	Use:   "account [<AccessKey> <SecretKey> <Name>]",
	Short: "Get/Set AccessKey and SecretKey",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 && len(args) != 3 {
			return fmt.Errorf("command account receives zero or three args, received %d\n", len(args))
		}
		return nil
	},
	Run: Account,
}

func Account(cmd *cobra.Command, params []string) {
	if len(params) == 0 {
		account, gErr := iqshell.GetAccount()
		if gErr != nil {
			fmt.Println(gErr)
			os.Exit(iqshell.STATUS_ERROR)
		}
		fmt.Println(account.String())
	} else if len(params) == 3 {
		accessKey := params[0]
		secretKey := params[1]
		name := params[2]

		pt, oldPath := iqshell.AccPath(), iqshell.OldAccPath()
		sErr := iqshell.SetAccount2(accessKey, secretKey, name, pt, oldPath)
		if sErr != nil {
			fmt.Println(sErr)
			os.Exit(iqshell.STATUS_ERROR)
		}
	}
}
