package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"scaligo/draw"
	"scaligo/help"
	"scaligo/keys"
	"scaligo/opts"
)

var (
	scale, title, cfg = opts.ParseOpts()
	frets             = len(cfg.Tuning[0])
	str               = len(cfg.Tuning)
)

const (
	fretSpacing   = 8
	stringSpacing = 2
)

func main() {
	//==========================================================================================
	// tcell screen
	//==========================================================================================
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

	//==========================================================================================
	// main loop || game loop
	//==========================================================================================
	for {
		s.Clear()
		//======================================================================================
		// Dimensions and positioning
		//======================================================================================
		termWidth, termHeight := s.Size()

		boardWidth := (frets - 1) * fretSpacing
		boardHeight := (str-1)*stringSpacing + 1

		x := (termWidth - boardWidth) / 2
		y := (termHeight - boardHeight) / 2
		//======================================================================================
		// reject small terminal
		//======================================================================================
		draw.RejectSmallTerm(termWidth, termHeight, s)
		//======================================================================================
		// draw loops
		//======================================================================================
		draw.DrawInlays(s, x, y, boardHeight, cfg)
		draw.DrawFrets(s, x, y, boardHeight, cfg)
		draw.DrawFretNumbers(s, x, y, boardHeight, cfg)
		draw.DrawStrings(s, x, y, boardWidth, stringSpacing, cfg)
		draw.DrawTabs(s, x, y, boardWidth, stringSpacing, cfg)
		draw.DrawJunctions(s, x, y, fretSpacing, stringSpacing, cfg)
		draw.DrawScaleInAlphabet(s, x, y, scale, title, cfg)
		draw.DrawNotesOnTheFretboard(s, x, y, scale, cfg)
		help.ShowHelp(s, termWidth, termHeight)
		s.Show()
		//======================================================================================
		// read keys
		//======================================================================================
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
