package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "timeShelly",
	Short: "timeShelly splits up a video file into parts",
	Long:  "timeShelly can download a video from youtube and then cut it into parts all from a inputfile",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
