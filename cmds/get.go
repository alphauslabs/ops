package cmds

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alphauslabs/blue-internal-go/operations/v1"
	protosinternal "github.com/alphauslabs/blue-internal-go/protos"
	"github.com/alphauslabs/bluectl/pkg/logger"
	"github.com/alphauslabs/ops/params"
	"github.com/alphauslabs/ops/pkg/connection"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/structpb"
)

func GetCmd() *cobra.Command {
	var (
		asOf        string
		includeDone bool
	)

	cmd := &cobra.Command{
		Use:   "get <orgId>",
		Short: "Describe a long operation item.",
		Long:  `Describe a long operation item.`,
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
			req := operations.GetOperationRequest{Name: args[0]}
			resp, err := client.GetOperation(ctx, &req)
			if err != nil {
				logger.Errorf("GetOperation failed: %v", err)
				return
			}

			switch {
			case params.OutFmt == "json":
				b, _ := json.Marshal(resp)
				logger.Info(string(b))
			default:
				logger.Infof("name: %v", resp.Name)
				var spb structpb.Struct
				resp.Metadata.UnmarshalTo(&spb)
				m := spb.AsMap()
				logger.Info("metadata:")
				for k, v := range m {
					logger.Infof("- %v: %v", k, fmt.Sprintf("%v", v))
				}

				logger.Infof("done: %v", resp.Done)
				if resp.Done {
					switch resp.Result.(type) {
					case *protosinternal.Operation_Response:
						logger.Infof("result: ok")
					case *protosinternal.Operation_Error:
						err := resp.Result.(*protosinternal.Operation_Error)
						logger.Infof("result: %q", err.Error.Message)
					}
				} else {
					logger.Infof("result:")
				}
			}
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringVar(&asOf, "as-of", asOf, "get items starting this date, fmt: yyyymmdd")
	cmd.Flags().BoolVar(&includeDone, "include-done", includeDone, "include completed operations")
	return cmd
}
