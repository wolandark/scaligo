package main

import (
	"flag"
	"fmt"
	"gs/scales"
	"os"
	"slices"
	"strings"

	"github.com/gdamore/tcell/v2"
	"golang.org/x/term"
)

func GetTerminalSize() (int, int) {
	fd := int(os.Stdin.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		fmt.Printf("Failed to obtain terminal size: %v\n", err)
	}
	return width, height
}

func GetScale(input string) ([]string, bool) {
	key := strings.ToLower(strings.TrimSpace(input))
	if alias, ok := scales.ScaleAliases[key]; ok {
		key = alias
	} else {
		key = strings.ReplaceAll(key, " ", "_")
	}
	scale, ok := scales.Scales[key]
	return scale, ok
}

func CalcCenter() (int, int, int, int) {
	w, h := GetTerminalSize()
	cw := w / 2
	ch := h / 2
	return w, h, cw, ch
}

func main() {
	s_flag := flag.String("s", "", "scale")
	flag.Parse()

	provided := false

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "s" {
			provided = true
		}
	})

	if !provided {
		fmt.Println("NO SCALE PROVIDED!")
		return
	}

	scale, exists := GetScale(*s_flag)

	if !exists {
		fmt.Println("Unknown scale:", *s_flag)
		return
	}

	s, _ := tcell.NewScreen()
	s.Init()
	defer s.Fini()

	style := tcell.StyleDefault.Foreground(tcell.ColorLavender)

	matchStyle := tcell.StyleDefault.Foreground(tcell.ColorGreen)

	lineCount := 13
	lineSpacing := 8
	termWidth, termHeight, centerWidth, centerHeight := CalcCenter()

	//-----------------------------------
	// Minimum  129 × 32 (columns × rows)
	//-----------------------------------
	if termWidth < 129 || termHeight < 32 {
		fmt.Println("term too small, go fuck yourself!")

		for i, r := range "Hello, World!" {
			s.SetContent(centerWidth, centerHeight, r, nil, style)
			i++
		}
		os.Exit(1)
	}
	//-----------------------------------
	//-----------------------------------

	x := centerWidth
	stringY, stringHeight := centerHeight, 6
	stringWidth := (lineCount - 1) * lineSpacing
	fretY, fretHeight := centerHeight, 15

	fretX := centerWidth / 2
	stringX := fretX

	fretboard_e := [][]string{
		{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
		{"B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
		{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
		{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
		{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
		{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
	}

	// draw frets only
	for i := 0; i < lineCount; i++ {
		for j := 0; j < fretHeight-4; j++ {
			s.SetContent(fretX+i*lineSpacing, fretY+j, '│', nil, style)
		}
	}

	// draw strings only
	for i := 0; i < stringHeight; i++ {
		for j := 0; j < stringWidth; j++ {
			s.SetContent(stringX+j, stringY+i*2, '-', nil, style)
		}
	}

	for i := 0; i < len(scale); i++ {
		for j := 0; j < len(scale[i]); j++ {
			a := rune(scale[i][j])
			s.SetContent(x, 15, a, nil, style)
			x++
		}
		x++
	}

	for stringIndex := 0; stringIndex < len(fretboard_e); stringIndex++ {
		notes := fretboard_e[stringIndex]

		for fretIndex := 0; fretIndex < len(notes); fretIndex++ {
			note := notes[fretIndex]

			if !slices.Contains(scale, note) {
				continue
			}

			nx := fretX + fretIndex*lineSpacing
			ny := stringY + stringIndex*2

			if fretIndex > 0 {
				nx -= lineSpacing / 2
			}

			runes := []rune(note)
			r := '●'
			if slices.Contains(runes, '#') {
				r = 'Δ'
			}
			if len(runes) > 1 {
				r = 'Δ'
				runes = runes[:1]
			}

			for i := 0; i < len(runes); i++ {
				s.SetContent(nx+i, ny, runes[i], nil, matchStyle)
				s.SetContent(nx+i, ny, r, nil, matchStyle)
			}
		}
	}

	s.Show()
	s.PollEvent()
	s.PollEvent()
}
