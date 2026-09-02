package content

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
				message:             "Meet me in the %s at midnight.",
				solution:            "CATACOMBS",
				completionQuote:     "The catacombs at midnight. Creepy, but it checks out.",
				completionNarration: "Hold on — another packet just hit the wire.",
			},
			{
				message:             "The password is %s.",
				solution:            "ILOVEROMANS",
				completionQuote:     "ILOVEROMANS. Not exactly subtle, Caesar.",
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
				message:             "The new tablets read %s when held to the glass.",
				solution:            "LOOKBACK",
				completionQuote:     "LOOKBACK. Cute. They're literally telling us how the cipher works.",
				completionNarration: atbashStageOneNarration,
			},
			{
				message:             "From now on, every dispatch is written in %s.",
				solution:            "MIRRORHAND",
				completionQuote:     "MIRRORHAND. So that's how they're writing now.",
				completionNarration: atbashStageOneNarration,
			},
			{
				message:             "The scribes now speak only in %s.",
				solution:            "REVERSAL",
				completionQuote:     "REVERSAL. Flip the alphabet, flip the message. Got it.",
				completionNarration: atbashStageOneNarration,
			},
		},
	},
	{
		tag:      "FOLLOW-UP INTERCEPT",
		preamble: "*A few hours later, Butts's minion barges into your office*\n",
		intro: []string{
			"Sergeant Butts intercepted an urgent despatch meant for their garrison.",
			"Same chaotic scribbles, but the handwriting is hurried and messy.",
			"We believe this to be immediate retreat coordinates — decode it fast",
			"so we can cut them off!",
		},
		prompt: "You've managed to crack the first one. Now let's see if you can do it again.",
		payloads: []choice{
			{
				message:             "Fall back to %s before sunrise. Burn the maps.",
				solution:            "RAVENNA",
				completionQuote:     "Ravenna. That's their bolt-hole.\nNice work, smartpants.",
				completionNarration: "If they move now, they can cut the garrison off at the gates.",
			},
			{
				message:             "Abandon the camp and rally at %s.",
				solution:            "AQUILEIA",
				completionQuote:     "Aquileia — that's a long march for a hurried retreat.",
				completionNarration: "Butts can have the road blocked before they ever get there.",
			},
			{
				message:             "Immediate withdrawal. Coordinates: %s ford.",
				solution:            "TIBER",
				completionQuote:     "The Tiber ford. I wouldn't mind getting a little wet.",
				completionNarration: "It doesn't look like they're going to make it.",
			},
		},
	},
}
