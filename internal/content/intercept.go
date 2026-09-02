package content

import (
	"fmt"
	"math/rand"
	"strings"
)

type choice struct {
	message             string
	solution            string
	completionQuote     string
	completionNarration string
}

type interceptStage struct {
	tag      string
	preamble string
	intro    []string
	prompt   string
	payloads []choice
}

func pickIntercepts(encoder Cipher, stages []interceptStage) []content {
	out := make([]content, 0, len(stages))
	for _, stage := range stages {
		c := stage.payloads[rand.Intn(len(stage.payloads))]
		out = append(out, content{
			Story:             interceptStory(encoder, stage, c),
			Solution:          c.solution,
			CompletionMessage: formatCompletion(c.completionQuote, c.completionNarration),
		})
	}
	return out
}

func interceptStory(encoder Cipher, stage interceptStage, c choice) string {
	var b strings.Builder

	b.WriteString(TagStyle.Render(stage.tag))
	b.WriteString("\n\n")
	if stage.preamble != "" {
		b.WriteString(stage.preamble)
		if !strings.HasSuffix(stage.preamble, "\n") {
			b.WriteString("\n")
		}
	}
	if len(stage.intro) > 0 {
		b.WriteString(formatQuotedDialogue(stage.intro...))
	}

	b.WriteString("\n\n")
	b.WriteString(TagStyle.Render("[!] INTERCEPTED PAYLOAD [!]"))
	b.WriteString("\n")

	payload := fmt.Sprintf(c.message, PayloadTagStyle.Render(encoder.Encrypt(c.solution)))
	b.WriteString(PayloadBoxStyle.Render(payload))

	if stage.prompt != "" {
		b.WriteString("\n\n")
		b.WriteString(stage.prompt)
	}

	return b.String()
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

func quoteAccent(word, rest string) string {
	return QuoteAccent.Render(word) + QuoteTextStyle.Render(rest)
}
