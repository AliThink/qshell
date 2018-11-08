package qshell

import (
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/cdn"
	"os"
)

func GetCdnManager() *cdn.CdnManager {
	account, gErr := GetAccount()
	if gErr != nil {
		fmt.Fprintf(os.Stderr, "GetCdnManager error: %v\n", gErr)
		os.Exit(1)
	}
	mac := qbox.NewMac(account.AccessKey, account.SecretKey)
	return cdn.NewCdnManager(mac)
}
