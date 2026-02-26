package functions

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func SimpleLS(w io.Writer, args []string, useColor bool) {
	if len(args) == 0 {
		args = []string{"."}
	}

	var files []string
	var dirs []string

	// Partition targets
	for _, target := range args {
		info, err := os.Lstat(target)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gols: cannot access '%s': %v\n", target, err)
			continue
		}

		if info.IsDir() {
			dirs = append(dirs, target)
		} else {
			files = append(files, target)
		}
	}

	// Sort top-level targets
	sort.Strings(files)
	sort.Strings(dirs)

	color := NewColor(useColor)

	// Print file targets first
	for _, file := range files {
		info, err := os.Lstat(file)
		if err != nil {
			continue
		}

		base := filepath.Base(file)
		color.ColorPrint(w, base, info)
	}

	// Print directory targets
	for i, dir := range dirs {

		// Print header if:
		// - more than one directory OR
		// - there were file targets
		if len(dirs) > 1 || len(files) > 0 {
			if i > 0 || len(files) > 0 {
				w.Write([]byte("\n"))
			}
			w.Write([]byte(dir + ":\n"))
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gols: cannot access '%s': %v\n", dir, err)
			continue
		}

		entries = dirFilter(entries)

		// Sort entries lexicographically
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Name() < entries[j].Name()
		})

		for _, entry := range entries {
			fullPath := filepath.Join(dir, entry.Name())

			info, err := os.Lstat(fullPath)
			if err != nil {
				continue
			}

			color.ColorPrint(w, entry.Name(), info)
		}
	}
}
