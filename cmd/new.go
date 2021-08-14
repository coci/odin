package cmd

import (
	"github.com/coci/odin/odin"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create new post",
	Long:  `This will create new post template inside the content directory.`,
	Run: func(cmd *cobra.Command, args []string) {

		if args[0] == "" {
			log.Println("you have to provide title for your post.")
			os.Exit(0)
		}

		odin.New(args[0])

	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
