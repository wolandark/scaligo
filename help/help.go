package help

import "github.com/gdamore/tcell/v2"

var showOverlay = false

func ShowHelp(s tcell.Screen, termWidth int, termHeight int) {
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
}

func Toggle() {
	showOverlay = !showOverlay
}
