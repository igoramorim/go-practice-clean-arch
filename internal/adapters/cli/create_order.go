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
	useCase := application.NewCreateOrderService(repo, nil)
	createOrderCmd := newCreateOrderCmd(useCase)
	orderCmd.AddCommand(createOrderCmd)

	// TODO: Handle errors

	createOrderCmd.Flags().Int64("id", 0, "ID of the order. Must be greater than 0.")
	_ = createOrderCmd.MarkFlagRequired("id")

	createOrderCmd.Flags().Float64P("price", "p", 0, "Price of the order. Must be greater than 0.")
	_ = createOrderCmd.MarkFlagRequired("price")

	createOrderCmd.Flags().Float64P("tax", "t", 0, "Tax that will be applied to the order price. Must be greater than 0.")
	_ = createOrderCmd.MarkFlagRequired("tax")
}

func newCreateOrderCmd(useCase dorder.CreateOrderUseCase) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new order.",
		Long:  `Create a new order.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cmd.Flags().GetInt64("id")
			if err != nil {
				return err
			}

			price, err := cmd.Flags().GetFloat64("price")
			if err != nil {
				return err
			}

			tax, err := cmd.Flags().GetFloat64("tax")
			if err != nil {
				return err
			}

			fmt.Println("id", id)
			fmt.Println("price", price)
			fmt.Println("tax", tax)

			in := dorder.CreateOrderUseCaseInput{
				ID:    id,
				Price: price,
				Tax:   tax,
			}
			out, err := useCase.Execute(context.Background(), in)
			fmt.Println(out)

			return err
		},
	}
}
