package cli

import (
	"context"
	"fmt"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"

	"github.com/spf13/cobra"
)

func init() {
	ctx := context.Background()
	client := newOrderClient()

	listOrdersCmd := newListOrdersCmd(ctx, client)
	orderCmd.AddCommand(listOrdersCmd)

	listOrdersCmd.Flags().IntP("page", "p", 1, "Page to request.")
	listOrdersCmd.Flags().IntP("limit", "l", 10, "How many orders by page.")
	listOrdersCmd.Flags().StringP("sort", "s", "asc", "How the orders will be sorted by id.")
}

func newListOrdersCmd(ctx context.Context, client *orderClient) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all created orders.",
		Long:  `List all created orders by page.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			page, err := cmd.Flags().GetInt("page")
			if err != nil {
				return err
			}

			limit, err := cmd.Flags().GetInt("limit")
			if err != nil {
				return err
			}

			sort, err := cmd.Flags().GetString("sort")
			if err != nil {
				return err
			}

			in := dorder.FindAllOrdersByPageUseCaseInput{
				Page:  page,
				Limit: limit,
				Sort:  sort,
			}
			response, err := client.listOrders(ctx, in)
			if err != nil {
				return err
			}

			fmt.Println(response)
			return nil
		},
	}
}
