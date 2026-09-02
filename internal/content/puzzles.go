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
			{ID: "atbash", Title: "Atbash Cipher", IsCompleted: false, IsLocked: true, Cipher: AtbashCipher{}, Content: pickIntercepts(AtbashCipher{}, atbashIntercepts), CompletedContentIndex: -1},
			{ID: "vernam", Title: "Boinkers Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1},
		}},
		{Title: "Chapter 2 Transposition Cipher", Puzzles: []*Puzzle{
			{ID: "transposition", Title: "Transposition Cipher", IsCompleted: false, IsLocked: false, Content: []content{}, CompletedContentIndex: -1},
			{ID: "rail_fence", Title: "Rail Fence Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1},
		}},
	}
}
