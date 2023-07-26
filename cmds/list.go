package cmds

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/alphauslabs/blue-internal-go/operations/v1"
	"github.com/alphauslabs/bluectl/pkg/logger"
	"github.com/alphauslabs/ops/params"
	"github.com/alphauslabs/ops/pkg/connection"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	var (
		asOf        string
		includeDone bool
	)

	cmd := &cobra.Command{
		Use:   "list <orgId>",
		Short: "List long operations",
		Long:  `List long operations.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				logger.Error("<orgId> is required.")
				return
			}

			ctx := context.Background()
			con, err := connection.New(ctx)
			if err != nil {
				logger.Errorf("connection.New failed: %v", err)
				return
			}

			defer con.Close()
			client := operations.NewOperationsClient(con)
			req := operations.ListOperationsRequest{
				Parent:      args[0],
				AsOf:        asOf,
				IncludeDone: includeDone,
			}

			stream, err := client.ListOperations(ctx, &req)
			if err != nil {
				logger.Errorf("ListOperations failed: %v", err)
				return
			}

			render := true
			table := tablewriter.NewWriter(os.Stdout)
			table.SetBorder(false)
			table.SetAutoWrapText(false)
			table.SetHeaderLine(false)
			table.SetColumnSeparator("")
			table.SetTablePadding("  ")
			table.SetNoWhiteSpace(true)
			table.Append([]string{"PARENT", "NAME", "STATUS"})

		loop:
			for {
				v, err := stream.Recv()
				switch {
				case err == io.EOF:
					break loop
				case err != nil && err != io.EOF:
					logger.Error(err)
					break loop
				}

				switch {
				case params.OutFmt == "json":
					render = false
					b, _ := json.Marshal(v)
					logger.Info(string(b))
				default:
					sts := "-"
					if v.Done {
						sts = "done"
					}

					table.Append([]string{args[0], v.Name, sts})
				}
			}

			if render {
				table.Render()
			}
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringVar(&asOf, "as-of", asOf, "get items starting this date, fmt: yyyymmdd")
	cmd.Flags().BoolVar(&includeDone, "include-done", includeDone, "include completed operations")
	return cmd
}
