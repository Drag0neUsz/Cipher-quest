package content

// Chapter 4 — Modern Ciphers: The Boinkers File
// A 1917 wireless hut. Lieutenant Boinkers invents four unbreakable systems
// a month and explains all of them at lunch. Arc: battery → new square →
// fake pad → he mails the real key with the telegram.

// scripts: playfairIntercepts; expected: PlayfairCipher, 5x5 square, key BOINK
var playfairIntercepts = []interceptStage{
	{
		tag:      "INTERCEPTED TRANSMISSION",
		preamble: "*The wireless hut smells like coffee, ozone, and unexplained confidence.*\n",
		intro: []string{
			"Welcome to the future, smartpants. It has worse coffee.",
			"",
			"Lieutenant Boinkers briefed the room on his 'pair system' at lunch.",
			"Letters travel in " + quoteAccent("pairs", ", like nervous lieutenants. 5 by 5 square."),
			"He hummed the keyword. It was not subtle.",
		},
		prompt: "Read the pairs and tell us where the battery goes.",
		payloads: []choice{
			{
				message:             "Relocate the battery to %s.",
				solution:            "YPRES",
				completionQuote:     "Ypres. He encoded a place everyone can already hear.",
				completionNarration: "Zero hour is in the next group. Boinkers is still talking.",
			},
			{
				message:             "Hold the ridge at %s.",
				solution:            "VERDUN",
				completionQuote:     "Verdun. If the ridge could vote, it would vote no.",
				completionNarration: "Timing group just came through. Same square, same lunch briefing.",
			},
			{
				message:             "Do not fire on %s.",
				solution:            "AMIENS",
				completionQuote:     "Amiens. A 'do not' that required a whole square. Impressive.",
				completionNarration: "There's a time hack in the next burst. Try to beat his explanation.",
			},
		},
	},
	{
		tag: "FOLLOW-UP INTERCEPT",
		intro: []string{
			"Same square, same pairs, slightly more shouting in Morse.",
			"This one is when the guns get ideas. If you miss it, the ridge finds out the hard way.",
		},
		prompt: "Name the hour. Then hide from Boinkers. He wants feedback.",
		payloads: []choice{
			{
				message:             "Barrage begins at %s.",
				solution:            "DAWN",
				completionQuote:     "Dawn. Classic. The sun remains available as a clock.",
				completionNarration: "We can have the line ready before the first bird complains.",
			},
			{
				message:             "Zero hour is %s.",
				solution:            "ZERO",
				completionQuote:     "ZERO. He encrypted the word zero. I need a transfer and a biscuit.",
				completionNarration: "Butts says if the hour is zero, the planning was too.",
			},
			{
				message:             "Wait for the %s.",
				solution:            "WHISTLE",
				completionQuote:     "The whistle. Nothing says modern warfare like a man with a pea.",
				completionNarration: "If they wait for a whistle, we can be the ones blowing it.",
			},
		},
	},
}

// scripts: bifidIntercepts; expected: BifidCipher, Polybius + split, period 5
var bifidIntercepts = []interceptStage{
	{
		tag: "INTERCEPTED TRANSMISSION",
		intro: []string{
			"He fractionated the alphabet, smartpants. I don't get paid enough.",
			"",
			"Boinkers 'improved' Playfair after we cracked it — letters become numbers,",
			"get shuffled, become letters again. Period of five. He drew it on a napkin.",
			"The napkin is in the file. The file is sticky.",
		},
		prompt: "Unsplit the grid and find the new square.",
		payloads: []choice{
			{
				message:             "New square is filed under %s.",
				solution:            "PIGEON",
				completionQuote:     "Pigeon. The codebook is in with the birds. Of course it is.",
				completionNarration: "Target coordinates are in the next burst. The pigeons look smug.",
			},
			{
				message:             "The grid starts with %s.",
				solution:            "FRANCE",
				completionQuote:     "FRANCE. He used the map as the key. Cartography weeps.",
				completionNarration: "There's a rail cut in the follow-up. Same sticky napkin energy.",
			},
			{
				message:             "Destroy yesterday's %s.",
				solution:            "SQUARE",
				completionQuote:     "Yesterday's square. He upgrades by throwing the last one in a bin.",
				completionNarration: "Decode the dump before he fractionates the furniture.",
			},
		},
	},
	{
		tag:      "FOLLOW-UP INTERCEPT",
		preamble: "*A pigeon lands on the set. Butts does not make eye contact.*\n",
		intro: []string{
			"Same split, same period, slightly more feathers.",
			"This is the rail they want cut. If the dump goes, the convoy goes hungry.",
		},
		prompt: "Name the cut and we can have a nicer railway.",
		payloads: []choice{
			{
				message:             "The dump is at %s.",
				solution:            "ARRAS",
				completionQuote:     "Arras. That's a lot of shells for a man who loses napkins.",
				completionNarration: "We can have the dump looking boring by tea.",
			},
			{
				message:             "Mine the rails at %s.",
				solution:            "SOMME",
				completionQuote:     "The Somme. He fractionated a river. Ambition, I'll give him that.",
				completionNarration: "Engineers can walk the line before anyone plants a hobby.",
			},
			{
				message:             "Balloons up over %s.",
				solution:            "MARNE",
				completionQuote:     "The Marne. If the balloons are up, so is the surprise.",
				completionNarration: "Butts wants the sky empty and Boinkers quieter.",
			},
		},
	},
}

// scripts: boinkersIntercepts; expected: repeating XOR / fake OTP, key BOINK
var boinkersIntercepts = []interceptStage{
	{
		tag: "INTERCEPTED TRANSMISSION",
		intro: []string{
			"He swears this is a one-time pad, smartpants.",
			"",
			"It is " + quoteAccent("five letters", ". He hummed them. Again. At lunch. Again."),
			"If you reuse a pad, it is not a pad. I told him. He promoted the napkin.",
		},
		prompt: "Treat it like a broken pad and read the boast.",
		payloads: []choice{
			{
				message:             "The pad is just %s.",
				solution:            "BOINK",
				completionQuote:     "BOINK. He put the key in the ciphertext. User friendly, he said.",
				completionNarration: "The actual convoy order is in the next group. Try to look surprised.",
			},
			{
				message:             "Never reuse %s.",
				solution:            "BOINK",
				completionQuote:     "Never reuse BOINK. He reused BOINK. I need a chaplain.",
				completionNarration: "Follow-up names the water. Same five letters. Obviously.",
			},
			{
				message:             "I wrote the key on my %s.",
				solution:            "CUFF",
				completionQuote:     "His cuff. The unbreakable system is machine-washable.",
				completionNarration: "Convoy details next. If his cuff goes to laundry, the war pauses.",
			},
		},
	},
	{
		tag: "FOLLOW-UP INTERCEPT",
		intro: []string{
			"Same five letters. Same lunch. Different water.",
			"This is the convoy he was hiding under the personality.",
		},
		prompt: "Name the water and we can have a quieter Channel.",
		payloads: []choice{
			{
				message:             "Convoy sails from %s.",
				solution:            "DOVER",
				completionQuote:     "Dover. He hid a cliff behind BOINK.",
				completionNarration: "Harbor can hold the sailing until the pad learns what 'one-time' means.",
			},
			{
				message:             "U-boats wait at %s.",
				solution:            "WIGHT",
				completionQuote:     "The Isle of Wight. Even the boats are embarrassed.",
				completionNarration: "If they're waiting there, we can wait better.",
			},
			{
				message:             "The listening post is %s.",
				solution:            "CALAIS",
				completionQuote:     "Calais. He encrypted a place you can see on a clear day.",
				completionNarration: "Butts wants the post deaf and the lieutenant reassigned.",
			},
		},
	},
}

// scripts: vernamIntercepts; expected: VernamCipher, true OTP
var vernamIntercepts = []interceptStage{
	{
		tag:      "INTERCEPTED TRANSMISSION",
		preamble: "*Two envelopes arrive stapled together. Someone has written 'DO NOT SEPARATE.'*\n",
		intro: []string{
			"He finally did it right, smartpants. That's the bad news.",
			"",
			"True random pad. Real Vernam. Unbreakable — until you notice he mailed",
			"the key in the " + quoteAccent("same envelope", "."),
			"Unbreakable, he said. He stapled the pad to the ciphertext and called it user friendly.",
		},
		prompt: "Open both envelopes. Try not to scream.",
		payloads: []choice{
			{
				message:             "Attached please find the %s.",
				solution:            "KEY",
				completionQuote:     "The key. Attached. Please find. I found it. I hate it.",
				completionNarration: "There's a personnel note in the second staple. Of course there is.",
			},
			{
				message:             "Do not lose the %s.",
				solution:            "PAD",
				completionQuote:     "Do not lose the pad. He glued it to the message. Problem solved.",
				completionNarration: "Reassignment orders are on the back. Someone upstairs has a sense of humor.",
			},
			{
				message:             "Second envelope is the %s.",
				solution:            "ENVELOPE",
				completionQuote:     "The second envelope is the envelope. I am requesting a medal.",
				completionNarration: "Decode the punchline. Then we can all go home, or to the zoo.",
			},
		},
	},
	{
		tag:      "FOLLOW-UP INTERCEPT",
		preamble: "*Butts reads the last slip twice. He almost smiles.*\n",
		intro: []string{
			"Same pad, same staple, mercifully last page.",
			"This is what they did with Boinkers after the envelopes.",
		},
		prompt: "Read the reassignment. Then lock the drawer. The drawer knows too much.",
		payloads: []choice{
			{
				message:             "Boinkers has been reassigned to %s.",
				solution:            "ZOO",
				completionQuote:     "The zoo. Secret, unofficial, and full of things that don't invent pads.",
				completionNarration: "If you ever find a giraffe with a codebook, you did not hear it from me.",
			},
			{
				message:             "Promote him to %s.",
				solution:            "ELSEWHERE",
				completionQuote:     "Elsewhere. That's not a rank. That's a prayer.",
				completionNarration: "Nice work, smartpants. The wireless can go back to being merely terrible.",
			},
			{
				message:             "The war will be won by %s.",
				solution:            "LUNCH",
				completionQuote:     "Lunch. He briefed four ciphers over soup and lost a war of envelopes.",
				completionNarration: "Butts is buying. I would not eat anything that comes with a napkin diagram.",
			},
		},
	},
}
