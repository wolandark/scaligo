package main

import (
	"flag"
	"fmt"
	"maps"

	"scaligo/help"
	"scaligo/keys"
	"scaligo/scales"

	"os"
	"slices"
	"strings"

	"github.com/gdamore/tcell/v2"
)

// ======================================================================================
// Constants
// ======================================================================================
const (
	fretSpacing   = 8
	stringSpacing = 2

	fretRune     = '│'
	stringRune   = '─'
	junctionRune = '┼'
	tRune        = '┬'
	revTRune     = '┴'
)

// ======================================================================================
// Styles
// ======================================================================================
var (
	style           = tcell.StyleDefault.Foreground(tcell.ColorWhite)
	scaleName       = tcell.StyleDefault.Foreground(tcell.ColorGreen)
	stringStyle     = tcell.StyleDefault.Foreground(tcell.ColorSilver)
	fretStyle       = tcell.StyleDefault.Foreground(tcell.ColorLightYellow)
	matchStyle      = tcell.StyleDefault.Foreground(tcell.ColorDarkCyan)
	matchSharpStyle = tcell.StyleDefault.Foreground(tcell.ColorRed)
	inlayStyle      = tcell.StyleDefault.Foreground(tcell.ColorWhiteSmoke)
	noteRune        = '⬤'
	sharpRune       = '⬤'
	inlayRune       = '◆'
)

// ======================================================================================
// Fretboard data
// ======================================================================================
var fretboard_e = [][]string{
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
	{"B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
}

func main() {
	//======================================================================================
	// flag section
	//======================================================================================
	// s_flag := flag.String("s", "", "scale")

	s_flag := flag.String("s", "A_Minor_Pentatonic", "scale") //default scale if no -s
	l_flag := flag.Bool("l", false, "list known scales")      // list known scales
	a_flag := flag.Bool("a", false, "list abbreviations")     // list scale aliases
	h_flag := flag.Bool("h", false, "commandline help")       // obviously prints help
	e_flag := flag.Bool("e", false, "use emojis")
	c_flag := flag.Bool("c", false, "customize colors")

	c_style := flag.String("style", "", "Fret color (hex)")
	c_scaleName := flag.String("scalename", "", "Fret color (hex)")
	c_fretStyle := flag.String("fretstyle", "", "Fret color (hex)")
	c_stringStyle := flag.String("stringstyle", "", "String color (hex)")
	c_matchStyle := flag.String("matchstyle", "", "Match color (hex)")
	c_matchSharpStyle := flag.String("matchsharpstyle", "", "Match color (hex)")
	c_inlayStyle := flag.String("inlaystyle", "", "Inlay color (hex)")

	flag.Parse()

	if *c_flag {
		if *c_style != "" {
			style = tcell.StyleDefault.Foreground(tcell.GetColor(*c_style))
		}
		if *c_scaleName != "" {
			scaleName = tcell.StyleDefault.Foreground(tcell.GetColor(*c_scaleName))
		}
		if *c_matchSharpStyle != "" {
			matchSharpStyle = tcell.StyleDefault.Foreground(tcell.GetColor(*c_matchSharpStyle))
		}
		if *c_fretStyle != "" {
			fretStyle = tcell.StyleDefault.Foreground(tcell.GetColor(*c_fretStyle))
		}
		if *c_stringStyle != "" {
			stringStyle = tcell.StyleDefault.Foreground(tcell.GetColor(*c_stringStyle))
		}
		if *c_matchStyle != "" {
			matchStyle = tcell.StyleDefault.Foreground(tcell.GetColor(*c_matchStyle))
		}
		if *c_inlayStyle != "" {
			inlayStyle = tcell.StyleDefault.Foreground(tcell.GetColor(*c_inlayStyle))
		}
	}

	if *e_flag {
		noteRune = '🟢'
		sharpRune = '🔴'
		inlayRune = '💠'

	}

	if *h_flag {
		fmt.Println("Use ./scaligo to enter interactive visualizer")
		fmt.Println()
		fmt.Println("Flags:")
		fmt.Println("\t-l : list all known scales")
		fmt.Println("\t-a : list all known scales and their aliases (shortcuts in interactive mode)")
		fmt.Println("\t-s : use -s to start scaligo with a known scale. Example: <scaligo -s am>")
		fmt.Println("\t-h : this help menu")
		fmt.Println()
		fmt.Println("Quick Start:")
		fmt.Println("\tIn interactive mode, press any sequence from the -a flag to swtich scales")
		fmt.Println("\tPress <h> to read the help and follow along")
		return
	}

	if *a_flag {
		for i, k := range slices.Sorted(maps.Keys(scales.ScaleMap)) {
			v := scales.ScaleMap[k]
			fmt.Printf("\n%d - %s -%s", i, k, v)
			if (i+1)%4 == 0 {
				fmt.Println()
			}
		}
		fmt.Println()
		return
	}

	if *l_flag {
		for i, name := range slices.Sorted(maps.Keys(scales.Scales)) {
			fmt.Printf("%-35s", name)
			if (i+1)%4 == 0 {
				fmt.Println()
			}
		}
		fmt.Println()
		return
	}

	scale, exists, title := scales.GetScale(*s_flag)

	if !exists {
		fmt.Println("Unknown scale:", *s_flag)
		os.Exit(1)
	}

	//======================================================================================
	// tcell screen
	//======================================================================================
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := s.Init(); err != nil {
		fmt.Println(err)
		return
	}
	defer s.Fini()

	//======================================================================================
	// main loop || game loop
	//======================================================================================
	for {
		s.Clear()

		//======================================================================================
		// Dimensions and positioning
		//======================================================================================
		termWidth, termHeight := s.Size()

		frets := len(fretboard_e[0])
		str := len(fretboard_e)

		boardWidth := (frets - 1) * fretSpacing
		boardHeight := (str-1)*stringSpacing + 1

		x := (termWidth - boardWidth) / 2
		y := (termHeight - boardHeight) / 2

		//======================================================================================
		// reject small terminal
		//======================================================================================
		if termWidth < 105 && termHeight < 21 {
			fmt.Println("term too small, go fuck yourself!")

			for i, r := range "Hello, World!" {
				s.SetContent(x, y, r, nil, style)
				i++
			}
			s.SetContent(x, y, '!', nil, tcell.StyleDefault)

			s.Clear()
			defer s.Fini()
			return
		}

		//======================================================================================
		// draw loops
		//======================================================================================
		// draw inlays
		//======================================================================================
		inlays := []int{1, 3, 5, 7, 9, 12}
		for fretIndex := 0; fretIndex < frets; fretIndex++ {
			if slices.Contains(inlays, fretIndex) {
				for j := 4; j < 6; j++ {
					for i := 1; i < boardHeight-1; i++ {
						// s.SetContent(x+(fretIndex-1)*fretSpacing+j, y+i, '🔷', nil, style)
						s.SetContent(x+(fretIndex-1)*fretSpacing+fretSpacing/2, y+boardHeight/2, inlayRune, nil, inlayStyle)
					}
				}
			}
		}

		//======================================================================================
		// draw frets only
		//======================================================================================
		for i := 0; i < frets; i++ {
			for j := 0; j < boardHeight; j++ {
				s.SetContent(x+i*fretSpacing, y+j, fretRune, nil, fretStyle)
			}
		}

		//======================================================================================
		// draw strings only
		//======================================================================================
		for i := 0; i < str; i++ {
			for j := 0; j < boardWidth; j++ {
				s.SetContent(x+j, y+i*stringSpacing, stringRune, nil, stringStyle)
			}
		}

		//======================================================================================
		// junctions
		//======================================================================================
		for i := 0; i < frets; i++ {
			for j := 0; j < str; j++ {
				nx := x + i*fretSpacing
				ny := y + j*stringSpacing

				r := junctionRune

				if j == 0 {
					r = tRune
				} else if j == str-1 {
					r = revTRune
				}

				s.SetContent(nx, ny, r, nil, fretStyle)
			}
		}

		//======================================================================================
		// draw scale notes in abcd
		//======================================================================================
		offset := 0
		// title = strings.ReplaceAll(title, "maj", "Major")
		// title = strings.ReplaceAll(title, "min", "Minor")
		// title = strings.ReplaceAll(title, "sharp", "Sharp")
		// title = strings.ToUpper(title[:1]) + title[1:]
		title = strings.ReplaceAll(title, "_", " ")
		for i := 0; i < len(title); i++ {
			s.SetContent(x+i, y-5, rune(title[i]), nil, scaleName)
		}

		for i := 0; i < len(scale); i++ {
			for j := 0; j < len(scale[i]); j++ {
				a := rune(scale[i][j])
				s.SetContent(x+offset, y-3, a, nil, style)
				offset++
			}
			offset++
		}

		//======================================================================================
		// draw notes on the fretboard
		//======================================================================================
		for stringIndex := 0; stringIndex < len(fretboard_e); stringIndex++ {
			notes := fretboard_e[stringIndex]

			for fretIndex := 0; fretIndex < len(notes); fretIndex++ {
				note := notes[fretIndex]

				if !slices.Contains(scale, note) {
					continue
				}

				nx := x + fretIndex*fretSpacing
				ny := y + stringIndex*stringSpacing

				if fretIndex > 0 {
					nx -= fretSpacing / 2
				}

				runes := []rune(note)
				r := noteRune
				var runeColor = matchStyle

				if slices.Contains(runes, '#') {
					r = sharpRune
					runeColor = matchSharpStyle
				}
				if len(runes) > 1 {
					r = sharpRune
					runes = runes[:1]
				}

				for i := 0; i < len(runes); i++ {
					s.SetContent(nx+i, ny, runes[i], nil, runeColor)
					s.SetContent(nx+i, ny, r, nil, runeColor)
				}
			}
		}

		help.ShowHelp(s, termWidth, termHeight)

		s.Show()

		// read keys
		newScale, newTitle, toggleOverlay, toggleShortcuts, running := keys.ReadKeys(s)
		if !running {
			break
		}

		if toggleOverlay {
			help.Toggle()
			toggleShortcuts = false
		}

		if toggleShortcuts {
			help.ToggleShortcuts()
			toggleOverlay = false
		}

		if newScale != nil {
			scale = newScale
			title = newTitle
		}
	}
}
