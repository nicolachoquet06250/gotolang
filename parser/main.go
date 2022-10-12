package parser

import (
	"gotolang/constants"
	"gotolang/types"
	"gotolang/utils"
	"strings"
)

func sendElementToArrayAndReset[T comparable](arr *[]T, element T) T {
	var v T
	if element != v {
		*arr = append(*arr, element)
		return v
	}
	return element
}

func sendSymbolToArray[T comparable](arr *[]T, s T) {
	*arr = append(*arr, s)
}

func createLine[T comparable](arr *[][]T, element *[]T) {
	*arr = append(*arr, *element)
}

func Parse(data string) (code *[][]string) {
	l := strings.Split(data, constants.LinesBreak)
	var splitCode [][]string

	for _, line := range l {
		splitLine := strings.Split(line, " ")
		var t []string
		for _, row := range splitLine {
			t = append(t, row)
		}
		splitCode = append(splitCode, t)
	}

	// debug
	utils.DebugAction(func() {
		for _, row := range splitCode {
			for _, col := range row {
				println(col)
			}
		}
	}, "", false)

	code = new([][]string)

	for _, row := range splitCode {
		tmp := new([]string)

		var lastWord string

		for i, col := range row {
			lastWord = sendElementToArrayAndReset(tmp, lastWord)

			if len(col) > 1 {
				for _, s := range strings.Split(col, "") {
					if types.Symbol(s).IsValid() {
						lastWord = sendElementToArrayAndReset(tmp, lastWord)
						sendSymbolToArray(tmp, s)
						continue
					}

					lastWord += s
				}
			} else {
				lastWord = sendElementToArrayAndReset(tmp, lastWord)
				sendSymbolToArray(tmp, col)
			}

			if i == len(row)-1 && lastWord != "" {
				lastWord = sendElementToArrayAndReset(tmp, lastWord)
			}
		}

		createLine(code, tmp)
	}

	// debug
	utils.DebugAction(func() {
		for _, row := range splitCode {
			for _, col := range row {
				println(col)
			}
		}
	}, "", false)

	return
}
