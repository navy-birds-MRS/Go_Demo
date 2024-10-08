package flag

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var greetCmd = &cobra.Command{
	Use:   "Net",
	Short: "Basic network commands",
	Run: func(cmd *cobra.Command, args []string) {
		ipAddr := cmd.Flag("ping").Value.String()

		pingcmd := exec.Command("ping", "-n", "1", ipAddr)

		// Execute the command and capture the output
		var out bytes.Buffer
		pingcmd.Stdout = &out
		err := pingcmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print the output
		fmt.Println(out.String())

		// Check if the ping was successful
		if bytes.Contains(out.Bytes(), []byte("Reply from")) {
			fmt.Println("Ping successful!")
		} else {
			fmt.Println("Ping failed!")
		}
	},
}

func init() {
	greetCmd.Flags().StringP("ping", "p", "127.0.0.1", "Ping IP or domain")
	rootCmd.AddCommand(greetCmd)
}

var rootCmd = &cobra.Command{Use: "Genie"}

func flag() {
	rootCmd.Execute()
}
