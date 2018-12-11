package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"time"
)

var cmd = &cobra.Command{
	Use: "cobra test",
	Short: "cobra short key",
	Long: "cobra long key",
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	time.Sleep(60 * time.Second)
}
