package functions

import (
	"os"
	"strings"
)

func dirFilter(entries []os.DirEntry) []os.DirEntry {
	var filtered []os.DirEntry

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		filtered = append(filtered, entry)
	}

	return filtered
}
