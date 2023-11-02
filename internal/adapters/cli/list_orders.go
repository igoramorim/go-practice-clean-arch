package cli

import (
	"context"
	"fmt"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/repository/mysqlorder"
	"github.com/igoramorim/go-practice-clean-arch/internal/application"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"

	"github.com/spf13/cobra"
)

func init() {
	// TODO: Fix DI

	repo := mysqlorder.New(nil)
	useCase := application.NewFindAllOrdersByPageService(repo)
	listOrdersCmd := newListOrdersCmd(useCase)
	orderCmd.AddCommand(listOrdersCmd)

	listOrdersCmd.Flags().IntP("page", "p", 1, "Page to request.")
	listOrdersCmd.Flags().IntP("limit", "l", 10, "How many orders by page.")
	listOrdersCmd.Flags().StringP("sort", "s", "asc", "How the orders will be sorted by id.")
}

func newListOrdersCmd(useCase dorder.FindAllOrdersByPageUseCase) *cobra.Command {
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

			fmt.Println("page", page)
			fmt.Println("limit", limit)
			fmt.Println("sort", sort)

			in := dorder.FindAllOrdersByPageUseCaseInput{
				Page:  page,
				Limit: limit,
				Sort:  sort,
			}
			out, err := useCase.Execute(context.Background(), in)
			fmt.Println(out)

			return err
		},
	}
}
