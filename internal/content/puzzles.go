package content

type Chapter struct {
	Puzzles []*Puzzle
	Title   string
}

type content struct {
	Story             string
	Solution          string
	CompletionMessage string
}

type Puzzle struct {
	Cipher                Cipher
	ID                    string
	Title                 string
	IsCompleted           bool
	IsLocked              bool
	Content               []content
	CompletedContentIndex int
	UnlocksIDs            []string
}

var Chapters = generateChapters()

func generateChapters() []Chapter {
	return []Chapter{
		{Title: "Chapter 1 Substitution Ciphers", Puzzles: []*Puzzle{
			{ID: "caesar", Title: "Caesar Cipher", IsCompleted: false, IsLocked: false, Cipher: CaesarCipher{Shift: 3}, Content: pickIntercepts(CaesarCipher{Shift: 3}, caesarIntercepts), CompletedContentIndex: -1, UnlocksIDs: []string{"atbash"}},
			{ID: "atbash", Title: "Atbash Cipher", IsCompleted: false, IsLocked: true, Cipher: AtbashCipher{}, Content: pickIntercepts(AtbashCipher{}, atbashIntercepts), CompletedContentIndex: -1, UnlocksIDs: []string{"keyword"}},
			// scripts: keywordIntercepts; expected: KeywordCipher, key JULIUS
			{ID: "keyword", Title: "Keyword Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1, UnlocksIDs: []string{"affine"}},
			// scripts: affineIntercepts; expected: AffineCipher, keys 5, 8
			{ID: "affine", Title: "Affine Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1, UnlocksIDs: []string{"scytale"}},
		}},
		{Title: "Chapter 2 Transposition Ciphers", Puzzles: []*Puzzle{
			// scripts: scytaleIntercepts; expected: ScytaleCipher, staff width
			{ID: "scytale", Title: "Scytale Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1, UnlocksIDs: []string{"columnar"}},
			// scripts: columnarIntercepts; expected: ColumnarCipher, keyword / column count
			{ID: "columnar", Title: "Columnar Transposition", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1, UnlocksIDs: []string{"rail_fence"}},
			// scripts: railFenceIntercepts; expected: RailFenceCipher, rail count
			{ID: "rail_fence", Title: "Rail Fence Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1, UnlocksIDs: []string{"route"}},
			// scripts: routeIntercepts; expected: RouteCipher, spiral / deck plan
			{ID: "route", Title: "Route Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1, UnlocksIDs: []string{"vigenere"}},
		}},
		{Title: "Chapter 3 Polyalphabetic Ciphers", Puzzles: []*Puzzle{
			// scripts: vigenereIntercepts; expected: VigenereCipher, repeating keyword AMORE
			{ID: "vigenere", Title: "Vigenère Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1, UnlocksIDs: []string{"gronsfeld"}},
			// scripts: gronsfeldIntercepts; expected: GronsfeldCipher, digits 31415
			{ID: "gronsfeld", Title: "Gronsfeld Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1, UnlocksIDs: []string{"autokey"}},
			// scripts: autokeyIntercepts; expected: AutokeyCipher, primer AMORE
			{ID: "autokey", Title: "Autokey Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1, UnlocksIDs: []string{"beaufort"}},
			// scripts: beaufortIntercepts; expected: BeaufortCipher, reciprocal of the same family
			{ID: "beaufort", Title: "Beaufort Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1, UnlocksIDs: []string{"playfair"}},
		}},
		{Title: "Chapter 4 Modern Ciphers", Puzzles: []*Puzzle{
			// scripts: playfairIntercepts; expected: PlayfairCipher, 5x5 square, key BOINK
			{ID: "playfair", Title: "Playfair Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1, UnlocksIDs: []string{"bifid"}},
			// scripts: bifidIntercepts; expected: BifidCipher, Polybius + split, period 5
			{ID: "bifid", Title: "Bifid Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1, UnlocksIDs: []string{"boinkers"}},
			// scripts: boinkersIntercepts; expected: repeating XOR / fake OTP, key BOINK
			{ID: "boinkers", Title: "Boinkers Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1, UnlocksIDs: []string{"vernam"}},
			// scripts: vernamIntercepts; expected: VernamCipher, true OTP
			{ID: "vernam", Title: "Vernam Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1},
		}},
	}
}
