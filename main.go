package main

import (
	"flag"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"gs/keys"
	"gs/scales"
	"os"
	"slices"
	"strings"
)

var showOverlay = false

// ======================================================================================
// Constants
// ======================================================================================
const (
	fretSpacing   = 8
	stringSpacing = 2

	noteRune     = '⬤'
	sharpRune    = '⬤'
	inlayRune    = '◆'
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
	fretStyle       = tcell.StyleDefault.Foreground(tcell.ColorGold)
	matchStyle      = tcell.StyleDefault.Foreground(tcell.ColorDarkCyan)
	matchSharpStyle = tcell.StyleDefault.Foreground(tcell.ColorRed)
	inlayStyle      = tcell.StyleDefault.Foreground(tcell.ColorWhiteSmoke)
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
	// flag section
	//======================================================================================
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

	scale, exists, title := scales.GetScale(*s_flag)

	if !exists {
		fmt.Println("Unknown scale:", *s_flag)
		os.Exit(1)
		return
	}

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

		//======================================================================================
		// help overlay
		//======================================================================================
		if showOverlay {
			helpLines := []string{

				"To Change Scales, Use Their Abreviations With The Formula Below:",
				"",
				"Any Major Key Will Be Displayed By Its Capital Letter Name, Such As <C>, <A>, <E> etc...",
				"",
				"Any Major Key Can Also Be Summoned With Its Lower Case Letter Followed By A Capital <M>, Such As <cM>, <eM> etc...",
				"",
				"Any Minor Key Can Also Be Summoned With Its Lower Case Letter Followed By A Lower Case <m>, Such As <cm>, <em> etc...",
				"",
				"Any Minor/Major Pentatonic Scale Can Be Summoned By A <p> In The End, Such As: <amp> For A Minor Pentatonic, And <aMp> For A Major Pentatonic.",
				"",
				"Example:",
				"C / cM - C Major",
				"cm - C Minor",
				"cMp - C Major Pentatonic",
				"",
				"Use The Same Rules For Other Scales",
				"",
				"q  - Quit",
				"Help - Press h To Close",
			}
			boxStyle := tcell.StyleDefault.Background(tcell.ColorDarkSlateGray).Foreground(tcell.ColorWhite)
			textStyle := boxStyle.Foreground(tcell.ColorGold)

			// box size from the help text lenght
			inner := 0
			for _, l := range helpLines {
				if len(l) > inner {
					inner = len(l)
				}
			}
			const pad = 2
			bw, bh := inner+pad*2+2, len(helpLines)+pad*2+2
			bx, by := (termWidth-bw)/2, (termHeight-bh)/2

			// fill
			for y := by; y < by+bh; y++ {
				for x := bx; x < bx+bw; x++ {
					s.SetContent(x, y, ' ', nil, boxStyle)
				}
			}
			// border
			for x := bx + 1; x < bx+bw-1; x++ {
				s.SetContent(x, by, '─', nil, boxStyle)
				s.SetContent(x, by+bh-1, '─', nil, boxStyle)
			}
			for y := by + 1; y < by+bh-1; y++ {
				s.SetContent(bx, y, '│', nil, boxStyle)
				s.SetContent(bx+bw-1, y, '│', nil, boxStyle)
			}
			s.SetContent(bx, by, '┌', nil, boxStyle)
			s.SetContent(bx+bw-1, by, '┐', nil, boxStyle)
			s.SetContent(bx, by+bh-1, '└', nil, boxStyle)
			s.SetContent(bx+bw-1, by+bh-1, '┘', nil, boxStyle)

			// text — same background as the fill, so no stripes
			for i, line := range helpLines {
				for j, r := range line {
					s.SetContent(bx+pad+1+j, by+pad+1+i, r, nil, textStyle)
				}
			}
		}

		s.Show()

		// read keys
		newScale, newTitle, toggleOverlay, running := keys.ReadKeys(s)
		if !running {
			break
		}

		if toggleOverlay {
			showOverlay = !showOverlay
		}

		if newScale != nil {
			scale = newScale
			title = newTitle
		}
	}
}
