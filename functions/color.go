package functions

import (
	"io"
	"os"
)

const (
	Blue  = "\033[34m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

type color struct {
	useColor bool
}

func NewColor(use bool) color {
	return color{useColor: use}
}

func (c color) ColorPrint(w io.Writer, name string, info os.FileInfo) {
	if !c.useColor {
		w.Write([]byte(name + "\n"))
		return
	}

	mode := info.Mode()

	if info.IsDir() {
		w.Write([]byte(Blue + name + Reset + "\n"))
		return
	}

	if mode.IsRegular() && (mode&0111) != 0 {
		w.Write([]byte(Green + name + Reset + "\n"))
		return
	}

	w.Write([]byte(name + "\n"))
}
