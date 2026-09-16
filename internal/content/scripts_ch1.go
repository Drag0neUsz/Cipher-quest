package content

// Chapter 1 — Substitution Ciphers: Caesar's March
// Caesar is moving on the city. Each cracked dispatch forces his scribes
// onto a fancier monoalphabetic trick. Arc: muster → new route → target → launch.

var caesarIntercepts = []interceptStage{
	{
		tag: "INTERCEPTED TRANSMISSION",
		intro: []string{
			"Hey smartpants!",
			"",
			"Sergeant Butts needs you to decipher this message from",
			"the Roman Caesar to his army. All we know is that he LOVES the",
			"number " + quoteAccent("3", ", but we're not sure what to do with this..."),
		},
		prompt: "Do you think you can help us out?",
		payloads: []choice{
			{
				message:             "Muster the tenth at the %s.",
				solution:            "CAMPUS",
				completionQuote:     "The Campus Martius. That's where you park a legion if you want the city to notice.",
				completionNarration: "Hold on — another packet just hit the wire.",
			},
			{
				message:             "Assemble the cohorts outside the %s.",
				solution:            "CAPENA",
				completionQuote:     "Porta Capena. South gate, south problem.",
				completionNarration: "Sergeant Butts sent a second message to you.",
			},
			{
				message:             "The signal fire burns on the %s.",
				solution:            "PALATINE",
				completionQuote:     "The Palatine hill. Of course he would light it there.",
				completionNarration: "A second transmission is coming through...",
			},
		},
	},
	{
		tag: "FOLLOW-UP INTERCEPT",
		intro: []string{
			"You cracked the first one!",
			"",
			"Sergeant Butts snagged a second message from the same courier.",
			"Same hand, same obsession with the number " + quoteAccent("3", "."),
			"This one has to give us an upper hand...",
		},
		prompt: "Decode the missing word and we can warn the city.",
		payloads: []choice{
			{
				message:             "The watchword at the gate is %s.",
				solution:            "AQUILA",
				completionQuote:     "Aquila — the eagle. That's the watchword.",
				completionNarration: "The gate guards will be listening for a bird. We can get there first.",
			},
			{
				message:             "Grain wagons roll through %s.",
				solution:            "OSTIA",
				completionQuote:     "Ostia. He's feeding the march before he starts it.",
				completionNarration: "If the grain moves, the tenth is not far behind.",
			},
			{
				message:             "Issue the new shields on the %s.",
				solution:            "JANICULUM",
				completionQuote:     "The Janiculum. That's a lot of lumber for a 'routine drill.'",
				completionNarration: "Butts can have eyes on the hill before the paint dries.",
			},
		},
	},
}

var atbashStageOneNarration = "The soldier walks away, muttering about the new cipher:\n" +
	QuoteTextStyle.Render("'Seems legit. I'll let sarge know about this.'")

var atbashIntercepts = []interceptStage{
	{
		tag: "INTERCEPTED TRANSMISSION",
		intro: []string{
			"Bad news, smartpants!",
			"",
			"Caesar figured out we cracked his lucky shift.\nHis scribes switched tactics completely.",
			"They're using a completely new cipher — our intel tells us that it",
			"has something to do with " + quoteAccent("mirrors", "."),
			"Think you can reflect this back at them?",
		},
		prompt: "Could a mirror be the key to cracking this?",
		payloads: []choice{
			{
				message:             "March east. Avoid the %s road.",
				solution:            "APPIAN",
				completionQuote:     "The Appian Way. So they know we watch the obvious road. Rude.",
				completionNarration: atbashStageOneNarration,
			},
			{
				message:             "New heading: the %s way.",
				solution:            "SALARIA",
				completionQuote:     "Via Salaria. The salt road. At least they'll be well seasoned.",
				completionNarration: atbashStageOneNarration,
			},
			{
				message:             "From now on, every dispatch is written in %s.",
				solution:            "MIRRORHAND",
				completionQuote:     "MIRRORHAND. Cute. They're literally telling us how the cipher works.",
				completionNarration: atbashStageOneNarration,
			},
		},
	},
	{
		tag:      "FOLLOW-UP INTERCEPT",
		preamble: "*A few hours later, Butts's minion barges into your office*\n",
		intro: []string{
			"Sergeant Butts intercepted an urgent despatch meant for their column.",
			"Same chaotic scribbles, but the handwriting is hurried and messy.",
			"We believe this to be the revised march — decode it fast",
			"so we can cut them off!",
		},
		prompt: "You've managed to crack the first one. Now let's see if you can do it again.",
		payloads: []choice{
			{
				message:             "Column proceeds to %s before sunrise.",
				solution:            "RAVENNA",
				completionQuote:     "Ravenna. That's their new heading.\nNice work, smartpants.",
				completionNarration: "If they move now, they can cut the column off at the gates.\nCaesar's luck with the number 3 just ran out.",
			},
			{
				message:             "Rally the tenth at %s.",
				solution:            "AQUILEIA",
				completionQuote:     "Aquileia — that's a long march for a 'revised drill.'",
				completionNarration: "Butts can have the road blocked before they ever get there.\nCaesar's luck with the number 3 just ran out.",
			},
			{
				message:             "Ford the %s by nightfall.",
				solution:            "TIBER",
				completionQuote:     "The Tiber. I wouldn't mind getting a little wet.",
				completionNarration: "It doesn't look like they're going to make it.\nCaesar's luck with the number 3 just ran out.",
			},
		},
	},
}

var keywordIntercepts = []interceptStage{
	{
		tag: "INTERCEPTED TRANSMISSION",
		intro: []string{
			"He's done with shifts, smartpants.",
			"",
			"Caesar's personal tablets just hit the bag — not the army-wide",
			"nonsense. The alphabet itself is scrambled. Intel says he mixed it",
			"from " + quoteAccent("his own name", ". Of course he did."),
			"Vanity: the original key-management policy.",
		},
		prompt: "Unmix the alphabet and tell Butts who he's buying.",
		payloads: []choice{
			{
				message:             "Pay the consul at the %s.",
				solution:            "CURIA",
				completionQuote:     "The Curia. He's shopping for a senate the honest way: coin.",
				completionNarration: "Hold on — the sealed orders have a second page.",
			},
			{
				message:             "Silver for the guards of the %s.",
				solution:            "ESQUILINE",
				completionQuote:     "The Esquiline. Those guards are about to have a very confusing payday.",
				completionNarration: "A follow-up tablet just slid under the door.",
			},
			{
				message:             "Trust no one in the senate but %s.",
				solution:            "BRUTUS",
				completionQuote:     "Brutus? Yeah, I'd keep an eye on that one.",
				completionNarration: "History might have a few notes about this later.",
			},
		},
	},
	{
		tag:      "FOLLOW-UP INTERCEPT",
		preamble: "*Butts slaps a wax tablet on your desk. It still smells expensive.*\n",
		intro: []string{
			"Same mixed alphabet. Same enormous signature.",
			"This one is the actual target — which wall, which gate, which",
			"unfortunate brick. Decode it before the ram starts practicing.",
		},
		prompt: "Name the breach and we can be standing in it first.",
		payloads: []choice{
			{
				message:             "The battering ram kisses %s.",
				solution:            "COLLINE",
				completionQuote:     "Porta Collina. North wall, north headache.",
				completionNarration: "Butts is already shouting at a map. Try not to look smug.",
			},
			{
				message:             "First breach at the %s gate.",
				solution:            "CAPENA",
				completionQuote:     "Capena again. He really likes that door.",
				completionNarration: "If they hit the south gate, we can have it looking busy.",
			},
			{
				message:             "Ignore the walls. Take the %s.",
				solution:            "CLOACA",
				completionQuote:     "The sewer. I am not paid enough for this sentence.",
				completionNarration: "Butts says if Caesar wants the Cloaca Maxima, he can have it.",
			},
		},
	},
}
