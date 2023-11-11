package main

import (
	"fmt"
	"github.com/kgunjikar/surexsend/sctl/users"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:  "say",
	Long: "Root Command",
}

var CompletionCmd = &cobra.Command{
	Use:                   "completion [bash|sh|fish|powershell]",
	Short:                 "Generate Completion script",
	Long:                  "To load completions ",
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "powershell":
			cmd.Root().GenBashCompletion(os.Stdout)
		}
	},
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(users.UserAdd, users.UserDelete, users.UserUpdate, users.UserList, CompletionCmd)
}
