package listcmd

import (
	"fmt"

	"github.com/cli/cli/v2/internal/tableprinter"
	"github.com/cli/cli/v2/pkg/iostreams"
)

type ConsoleListRenderer struct {
	IO *iostreams.IOStreams
}

func (r ConsoleListRenderer) Render(sponsors []Sponsor) error {
	if len(sponsors) == 0 && r.IO.IsStdoutTTY() {
		fmt.Fprintln(r.IO.Out, "No sponsors found")
		return nil
	}

	tp := tableprinter.New(r.IO, tableprinter.WithHeader("SPONSOR"))
	for _, sponsor := range sponsors {
		tp.AddField(string(sponsor))
		tp.EndRow()
	}
	return tp.Render()
}
