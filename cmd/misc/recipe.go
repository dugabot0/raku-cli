package misc

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newRecipeCmd() *cobra.Command {
	var (
		categoryType string
		categoryID   string
	)

	c := &cobra.Command{
		Use:   "recipe",
		Short: "Get recipe category list",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.RecipeCategoryList(client.RecipeParams{
				CategoryType: categoryType,
				CategoryID:   categoryID,
			})
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&categoryType, "category-type", "", "Category type: large/medium/small")
	f.StringVar(&categoryID, "category-id", "", "Category ID")

	return c
}
