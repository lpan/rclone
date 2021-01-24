package merge

import (
	"context"
	"errors"

	"github.com/rclone/rclone/cmd"
	"github.com/rclone/rclone/fs/sync"
	"github.com/spf13/cobra"
)

func init() {
	cmd.Root.AddCommand(commandDefinition)
	_ = commandDefinition.Flags()
}

var commandDefinition = &cobra.Command{
	Use:   "merge source:path dest:path",
	Short: `Merge source and dest so that they are identical, modifying both source and dest.`,
	Long: `
TODO documentation
`,
	Run: func(command *cobra.Command, args []string) {
		cmd.CheckArgs(2, 2, command, args)
		fsrc, srcFileName, fdst := cmd.NewFsSrcFileDst(args)
		cmd.Run(true, true, command, func() error {
			if srcFileName == "" {
				return sync.Merge(context.Background(), fdst, fsrc)
			}

			return errors.New("Bruh")
		})
	},
}
