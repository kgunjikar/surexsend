package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UserAdd = &cobra.Command{
	Use:   "useradd",
	Short: "Add User",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add User")
	},
}

var UserDelete = &cobra.Command{
	Use:   "userdel",
	Short: "Delete User",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete User")
	},
}

var UserUpdate = &cobra.Command{
	Use:   "userupdate",
	Short: "Update User",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Update User")
	},
}

var UserList = &cobra.Command{
	Use:   "userlist",
	Short: "List User",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List User")
	},
}
