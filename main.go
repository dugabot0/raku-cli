package main

import (
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/cmd/books"
	"github.com/lumiere/raku-cli/cmd/ichiba"
	"github.com/lumiere/raku-cli/cmd/misc"
	"github.com/lumiere/raku-cli/cmd/travel"
)

func main() {
	cmd.AddCommand(ichiba.NewCmd())
	cmd.AddCommand(books.NewCmd())
	cmd.AddCommand(travel.NewCmd())
	cmd.AddCommand(misc.NewCmd())
	cmd.Execute()
}
