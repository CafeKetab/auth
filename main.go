package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/CafeKetab/auth-go/cmd"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main() {
	const short = "short description"
	const long = `long description`

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

	root := &cobra.Command{Short: short, Long: long}

	root.AddCommand(
		cmd.Server{}.Command(trap),
	)

	if err := root.Execute(); err != nil {
		log.Fatal("failed to execute root command", zap.Error(err))
	}
}