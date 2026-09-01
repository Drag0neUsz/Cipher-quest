package content

import (
	"fmt"
	"math/rand"
	"strings"
)

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
			{ID: "caesar", Title: "Caesar Cipher", IsCompleted: false, IsLocked: false, Cipher: CaesarCipher{Shift: 3}, Content: pickCaesar(), CompletedContentIndex: -1, UnlocksIDs: []string{"vernam", "rail_fence"}},
			{ID: "atbash", Title: "Atbash Cipher", IsCompleted: false, IsLocked: false, Content: []content{}, CompletedContentIndex: -1},
			{ID: "vernam", Title: "Boinkers Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1},
		}},
		{Title: "Chapter 2 Transposition Cipher", Puzzles: []*Puzzle{
			{ID: "transposition", Title: "Transposition Cipher", IsCompleted: false, IsLocked: false, Content: []content{}, CompletedContentIndex: -1},
			{ID: "rail_fence", Title: "Rail Fence Cipher", IsCompleted: false, IsLocked: true, Content: []content{}, CompletedContentIndex: -1},
		}},
	}
}

type choice struct {
	message             string
	solution            string
	completionQuote     string
	completionNarration string
}

func formatQuotedDialogue(lines ...string) string {
	styled := make([]string, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			styled = append(styled, "")
			continue
		}
		styled = append(styled, QuoteTextStyle.Render(line))
	}
	return QuoteStyle.Render(strings.Join(styled, "\n"))
}

func formatCompletion(quote string, narration string) string {
	var b strings.Builder
	if quote != "" {
		b.WriteString(formatQuotedDialogue(strings.Split(quote, "\n")...))
	}
	if narration != "" {
		if quote != "" {
			b.WriteString("\n\n")
		}
		b.WriteString(narration)
	}
	return b.String()
}

func pickCaesar() []content {
	encoder := CaesarCipher{Shift: 3}
	stageOne := []choice{
		{
			message:             "Meet me in the %s at midnight.",
			solution:            "CATACOMBS",
			completionQuote:     "The catacombs at midnight. Creepy, but it checks out.",
			completionNarration: "Hold on — another packet just hit the wire.",
		},
		{
			message:             "The password is %s.",
			solution:            "ILOVEROMANS",
			completionQuote:     "ILOVEROMANS. Not exactly subtle, Caesar.",
			completionNarration: "Sergeant Butts is waving another intercept at you.",
		},
		{
			message:             "The signal fire burns on the %s.",
			solution:            "PALATINE",
			completionQuote:     "The Palatine hill. Of course he would light it there.",
			completionNarration: "A second transmission is coming through...",
		},
	}

	stageTwo := []choice{
		{
			message:             "We cross the %s at dawn. Tell no one.",
			solution:            "RUBICON",
			completionQuote:     "The Rubicon. Once you cross it, there is no going back.\nNice work, smartpants.",
			completionNarration: "Caesar's luck with the number 3 just ran out.",
		},
		{
			message:             "The watchword at the gate is %s.",
			solution:            "AQUILA",
			completionQuote:     "Aquila — the eagle. That's the watchword.",
			completionNarration: "The gate guards will be listening for a bird. We can get there first.",
		},
		{
			message:             "Trust no one in the senate but %s.",
			solution:            "BRUTUS",
			completionQuote:     "Brutus? Yeah, I'd keep an eye on that one.",
			completionNarration: "History might have a few notes about this later.",
		},
	}

	first := stageOne[rand.Intn(len(stageOne))]
	second := stageTwo[rand.Intn(len(stageTwo))]

	return []content{
		{
			Story:             caesarInterceptStory(encoder, first, 0),
			Solution:          first.solution,
			CompletionMessage: formatCompletion(first.completionQuote, first.completionNarration),
		},
		{
			Story:             caesarInterceptStory(encoder, second, 1),
			Solution:          second.solution,
			CompletionMessage: formatCompletion(second.completionQuote, second.completionNarration),
		},
	}
}

func caesarInterceptStory(encoder CaesarCipher, c choice, stage int) string {
	var b strings.Builder

	switch stage {
	case 0:
		b.WriteString(TagStyle.Render("INTERCEPTED TRANSMISSION"))
		b.WriteString("\n\n")
		b.WriteString(formatQuotedDialogue(
			"Hey smartpants!",
			"",
			"Sergeant Butts needs you to decipher this message from",
			"the Roman Caesar to his army. All we know is that he LOVES the",
			"number "+QuoteAccent.Render("3")+QuoteTextStyle.Render(", but we're not sure what to do with this..."),
		))
	case 1:
		b.WriteString(TagStyle.Render("FOLLOW-UP INTERCEPT"))
		b.WriteString("\n\n")
		b.WriteString(formatQuotedDialogue(
			"You cracked the first one!",
			"",
			"Sergeant Butts snagged a second message from the same courier.",
			"Same hand, same obsession with the number "+QuoteAccent.Render("3")+QuoteTextStyle.Render("."),
			"This one sounds like actual marching orders...",
		))
	default:
		b.WriteString(TagStyle.Render(fmt.Sprintf("INTERCEPT STAGE %d", stage+1)))
		b.WriteString("\n\n")
		b.WriteString(formatQuotedDialogue(
			"Another one just came through.",
			"",
			"Same courier, same lucky number "+QuoteAccent.Render("3")+QuoteTextStyle.Render("."),
		))
	}

	b.WriteString("\n\n")
	b.WriteString(TagStyle.Render("[!] INTERCEPTED PAYLOAD [!]"))
	b.WriteString("\n")

	payload := fmt.Sprintf(c.message, PayloadTagStyle.Render(encoder.Encrypt(c.solution)))
	b.WriteString(PayloadBoxStyle.Render(payload))
	b.WriteString("\n\n")

	switch stage {
	case 0:
		b.WriteString("Do you think you can help us out?")
	case 1:
		b.WriteString("Decode the missing word and we can warn the city.")
	default:
		b.WriteString("The pattern hasn't changed. Shift it back.")
	}

	return b.String()
}
