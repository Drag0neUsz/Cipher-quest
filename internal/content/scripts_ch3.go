package content

// Chapter 3 — Polyalphabetic Ciphers: Amore Infinito
// Carnival in Venice. Maestro Recitativo sells state secrets inside his flop
// opera (one good review: his mother). Arc: buyer → drop → masquerade → double-cross.

// scripts: vigenereIntercepts; expected: VigenereCipher, repeating keyword AMORE
var vigenereIntercepts = []interceptStage{
	{
		tag:      "INTERCEPTED TRANSMISSION",
		preamble: "*You are in a box at the opera. Butts is not a box-at-the-opera person.*\n",
		intro: []string{
			"I don't do culture, smartpants. I do intercepts.",
			"",
			"Maestro Recitativo hid a buyer list in *Amore Infinito* — five hours,",
			"one plot, a love-duet word that " + quoteAccent("never ends", "."),
			"If that chorus repeats one more time I will decrypt it out of spite.",
		},
		prompt: "Ride the repeating key and name the buyer.",
		payloads: []choice{
			{
				message:             "The buyer wears a %s mask.",
				solution:            "RAVEN",
				completionQuote:     "A raven mask. In Venice. He might as well have worn a nametag.",
				completionNarration: "Act two just started. Unfortunately, so did the next packet.",
			},
			{
				message:             "Sell the roster to the %s.",
				solution:            "DOGE",
				completionQuote:     "The Doge. That's not a buyer, that's a performance review.",
				completionNarration: "There's a second verse about the merchandise. Of course there is.",
			},
			{
				message:             "Our friend in the box is %s.",
				solution:            "OTHELLO",
				completionQuote:     "Othello. Wrong play, Maestro. I don't care. Neither does Venice.",
				completionNarration: "Butts just asked if this is the one with the ghosts. Decode the goods.",
			},
		},
	},
	{
		tag: "FOLLOW-UP INTERCEPT",
		intro: []string{
			"Same key, same duet, slightly more treason per measure.",
			"This one says what he's selling. If it's a fleet, someone in a mask already paid.",
		},
		prompt: "Name the goods before the encore starts bargaining.",
		payloads: []choice{
			{
				message:             "Tonight's aria hides the %s.",
				solution:            "FLEET",
				completionQuote:     "The fleet. He put a navy in a love song. Critics will be furious.",
				completionNarration: "Harbor master can count ships that still exist.",
			},
			{
				message:             "The tenor's high C marks the %s.",
				solution:            "TREATY",
				completionQuote:     "The treaty. Nothing says diplomacy like a man in tights hitting C.",
				completionNarration: "Butts wants the ink watched. I want earplugs.",
			},
			{
				message:             "Act two reveals the %s.",
				solution:            "ARMORY",
				completionQuote:     "The armory. He rhymed it with 'amory.' I checked. It does not.",
				completionNarration: "If the weapons are in act two, we can be in the basement in act one.",
			},
		},
	},
}

// scripts: gronsfeldIntercepts; expected: GronsfeldCipher, digits 31415
var gronsfeldIntercepts = []interceptStage{
	{
		tag: "INTERCEPTED TRANSMISSION",
		intro: []string{
			"The Maestro has bills, smartpants. This is not a metaphor.",
			"",
			"Next note is numbers — " + quoteAccent("box-office receipts", ", or his unpaid tailor."),
			"Same repeating trick, just with digits, because art is expensive and he is not paid in applause.",
		},
		prompt: "Run the numbers and find the drop.",
		payloads: []choice{
			{
				message:             "Leave the purse in the %s.",
				solution:            "GONDOLA",
				completionQuote:     "A gondola. Very discreet, if you ignore the singing.",
				completionNarration: "Price tag is in the next bar. I am afraid.",
			},
			{
				message:             "Coins under the %s.",
				solution:            "RIALTO",
				completionQuote:     "The Rialto. He picked the busiest bridge in the city. Genius.",
				completionNarration: "Whatever he wants in return is on the invoice.",
			},
			{
				message:             "Meet at the %s.",
				solution:            "BRIDGE",
				completionQuote:     "The bridge. He did not specify which. Venice has a few.",
				completionNarration: "Decode the fee. I need to know how embarrassed to be.",
			},
		},
	},
	{
		tag:      "FOLLOW-UP INTERCEPT",
		preamble: "*The tailor's invoice is pinned to the score. There are tears. Not musical ones.*\n",
		intro: []string{
			"Same digits. Same debt.",
			"This is what Recitativo actually wants. Try to keep a straight face.",
		},
		prompt: "Name the price. Butts is already judging him.",
		payloads: []choice{
			{
				message:             "He wants payment in %s.",
				solution:            "DUCATS",
				completionQuote:     "Ducats. Traditional. Boring. Almost respectable.",
				completionNarration: "We can mark the coins. Or the Maestro. Either works.",
			},
			{
				message:             "No paper. Bring %s.",
				solution:            "RUBIES",
				completionQuote:     "Rubies. He has taste and no pockets big enough.",
				completionNarration: "If he wants stones, we can bring him geology.",
			},
			{
				message:             "The price is one %s.",
				solution:            "BASSOON",
				completionQuote:     "A bassoon. He is selling a fleet for a woodwind. I need a drink.",
				completionNarration: "Butts says if that's the market, the republic is going to be fine.",
			},
		},
	},
}

// scripts: autokeyIntercepts; expected: AutokeyCipher, primer AMORE
var autokeyIntercepts = []interceptStage{
	{
		tag: "INTERCEPTED TRANSMISSION",
		intro: []string{
			"He read a pamphlet, smartpants. That's never good.",
			"",
			"Recitativo got " + quoteAccent("paranoid", ". The key is now the message itself,"),
			"after a little primer from the duet. He thinks this makes him modern.",
			"It makes him louder.",
		},
		prompt: "Chase the autokey and name the next meeting.",
		payloads: []choice{
			{
				message:             "Masquerade at the %s.",
				solution:            "ARSENALE",
				completionQuote:     "The Arsenale. A costume party at the shipyard. Subtle as a fanfare.",
				completionNarration: "There's a cleanup note. He's spooked. Good.",
			},
			{
				message:             "After the encore, the %s.",
				solution:            "LAGOON",
				completionQuote:     "The lagoon. Romantic, wet, and full of witnesses in boats.",
				completionNarration: "Second page says what to burn. I have a guess.",
			},
			{
				message:             "Midnight in the %s.",
				solution:            "CAMPANILE",
				completionQuote:     "The campanile. If you're going to betray a city, do it with a view.",
				completionNarration: "Follow-up is him panicking in 3/4 time.",
			},
		},
	},
	{
		tag: "FOLLOW-UP INTERCEPT",
		intro: []string{
			"Same primer, same panic, worse penmanship.",
			"He wants something burned. If it's the libretto, someone finally reviewed it.",
		},
		prompt: "Read the scare-note before he writes a sequel.",
		payloads: []choice{
			{
				message:             "Burn the libretto labeled %s.",
				solution:            "INFINITO",
				completionQuote:     "Amore Infinito. Please. Let me watch.",
				completionNarration: "If the score goes, the secrets go with it. Copy first.",
			},
			{
				message:             "Trust only the %s.",
				solution:            "PROMPTER",
				completionQuote:     "The prompter. The man in the box who already knows everyone's lines.",
				completionNarration: "Butts wants a word with the box. Quietly.",
			},
			{
				message:             "They cracked the chorus. Switch to %s.",
				solution:            "AUTOKEY",
				completionQuote:     "AUTOKEY. He wrote the method in the method. I love this job.",
				completionNarration: "If he's announcing upgrades, the finale is going to be a tantrum.",
			},
		},
	},
}

// scripts: beaufortIntercepts; expected: BeaufortCipher, reciprocal of the same family
var beaufortIntercepts = []interceptStage{
	{
		tag:      "INTERCEPTED TRANSMISSION",
		preamble: "*The Maestro is in the wings, sulking in a cape that has seen things.*\n",
		intro: []string{
			"He sang it " + quoteAccent("backwards", ", smartpants. That's the whole upgrade."),
			"",
			"Same family of repeating keys, just reciprocal, because he had a fight",
			"with the tenor and the pamphlet. This is the double-cross verse.",
		},
		prompt: "Reverse the aria and name the other buyer.",
		payloads: []choice{
			{
				message:             "The other buyer is %s.",
				solution:            "SPAIN",
				completionQuote:     "Spain. He sold the same page twice and called it a duet.",
				completionNarration: "Curtain note incoming. Try to enjoy the plot for once.",
			},
			{
				message:             "I sold the same page to %s.",
				solution:            "BOTH",
				completionQuote:     "BOTH. Honesty: finally, a motif I can hum.",
				completionNarration: "The council is going to need a bigger box.",
			},
			{
				message:             "Tell the council I am %s.",
				solution:            "LOYAL",
				completionQuote:     "Loyal. He encoded a lie and still missed the tempo.",
				completionNarration: "Final performance notes are on the back of the cape. I wish I were joking.",
			},
		},
	},
	{
		tag: "FOLLOW-UP INTERCEPT",
		intro: []string{
			"Last page of *Amore Infinito*. Thank every god that will listen.",
			"Same backwards sulk. This is where he takes a bow or a boat.",
		},
		prompt: "Bring the curtain down. Then never mention opera in this office.",
		payloads: []choice{
			{
				message:             "Final performance at %s.",
				solution:            "FENICE",
				completionQuote:     "La Fenice. If the house burns down, blame art, not us.",
				completionNarration: "Butts can have men at every door. I can have silence.",
			},
			{
				message:             "If I am caught I will %s.",
				solution:            "IMPROVISE",
				completionQuote:     "Improvise. That's what he called the first four hours too.",
				completionNarration: "Nice work, smartpants. If he improvises, we already have the score.",
			},
			{
				message:             "The real key is %s.",
				solution:            "MOTHER",
				completionQuote:     "Mother. She was the only good review. Of course it was her nickname.",
				completionNarration: "Recitativo can take a bow. Then he can take a cell.",
			},
		},
	},
}
