package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tonycai653/iqshell/qshell"
	"os"
)

var (
	prefopCmd = &cobra.Command{
		Use:   "prefop <PersistentId>",
		Short: "Query the pfop status",
		Args:  cobra.ExactArgs(1),
		Run:   Prefop,
	}

	fopCmd = &cobra.Command{
		Use:   "pfop <Bucket> <Key> <fopCommand>",
		Short: "issue a request to process file in bucket",
		Args:  cobra.ExactArgs(3),
		Run:   Fop,
	}
)

var pipeline string

func init() {
	fopCmd.Flags().StringVarP(&pipeline, "pipeline", "p", "", "task pipeline")
	RootCmd.AddCommand(prefopCmd, fopCmd)
}

func Prefop(cmd *cobra.Command, params []string) {
	persistentId := params[0]

	fopRet, err := qshell.Prefop(persistentId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Prefop error: %v\n", err)
		os.Exit(qshell.STATUS_ERROR)
	} else {
		fmt.Println(fopRet.String())
	}
}

func Fop(cmd *cobra.Command, params []string) {
	bucket, key, fops := params[0], params[1], params[2]

	persistengId, err := qshell.Pfop(bucket, key, fops, pipeline)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Prefop error: %v\n", err)
		os.Exit(qshell.STATUS_ERROR)
	}
	fmt.Println(persistengId)
}
