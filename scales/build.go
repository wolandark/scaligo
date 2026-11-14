package scales

import "slices"

var chromatic = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

// key spellings must match what ScaleAliases and ScaleMap references
var roots = map[string]string{
	"A": "A", "A#": "A_Sharp",
	"B": "B",
	"C": "C", "C#": "C_Sharp",
	"D": "D", "D#": "D_Sharp",
	"E": "E",
	"F": "F", "F#": "F_Sharp",
	"G": "G", "G#": "G_Sharp",
}

var families = map[string][]int{
	"Major":                  {0, 2, 4, 5, 7, 9, 11},
	"Minor":                  {0, 2, 3, 5, 7, 8, 10},
	"Major_Pentatonic":       {0, 2, 4, 7, 9},
	"Minor_Pentatonic":       {0, 3, 5, 7, 10},
	"Major_Pentatonic_Blues": {0, 2, 3, 4, 7, 9},
	"Minor_Pentatonic_Blues": {0, 3, 5, 6, 7, 10},
}

func BuildScale(root string, steps []int) []string {
	base := slices.Index(chromatic, root)
	notes := make([]string, len(steps))
	for i, s := range steps {
		notes[i] = chromatic[(base+s)%12]
	}
	return notes
}

var Scales = BuildAll()

func BuildAll() map[string][]string {
	m := make(map[string][]string, 12*len(families))
	for root, spelling := range roots {
		for name, steps := range families {
			m[spelling+"_"+name] = BuildScale(root, steps)
		}
	}
	return m
}
