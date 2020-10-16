package printer

import (
	"github.com/vultr/govultr"
)

func Os(vultrOS []govultr.OS, meta *govultr.Meta) {
	col := columns{"OSID", "NAME", "ARCH", "FAMILY"}
	display(col)
	for _, os := range vultrOS {
		display(columns{os.ID, os.Name, os.Arch, os.Family})
	}

	Meta(meta)
	flush()
}
