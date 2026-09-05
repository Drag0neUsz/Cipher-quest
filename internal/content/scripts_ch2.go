package content

// Chapter 2 — Transposition Ciphers: The Scrambled Galleon
// Butts is on loan to the Admiralty and seasick. Captain Widdershins and
// quartermaster Patch rearrange every letter until the ransom note looks mean.
// Arc: find the galleon → stop the silver raid → beat them to the stash.

// scripts: scytaleIntercepts; expected: ScytaleCipher, staff width
var scytaleIntercepts = []interceptStage{
	{
		tag:      "INTERCEPTED TRANSMISSION",
		preamble: "*The Admiralty cabin pitches. Butts is the same color as the soup.*\n",
		intro: []string{
			"New assignment, smartpants. Try not to enjoy it.",
			"",
			"Captain Widdershins has a galleon and a quartermaster named Patch",
			"who wraps her orders around a " + quoteAccent("belaying pin", " and calls it 'filing.'"),
			"Every letter is present. They're just as organized as the crew.",
		},
		prompt: "Unwrap the pin and tell us where she's lurking.",
		payloads: []choice{
			{
				message:             "We lie up in %s cove.",
				solution:            "TORTUGA",
				completionQuote:     "Tortuga. Of course it's Tortuga. Subtle as a cannon.",
				completionNarration: "Hold on — Patch sent a second wrap. The pin is still warm.",
			},
			{
				message:             "The galleon sleeps at %s.",
				solution:            "SKULL",
				completionQuote:     "Skull Cove. She named it herself. I can tell.",
				completionNarration: "Butts wants the cove watched. After he finds the horizon again.",
			},
			{
				message:             "Anchor in the lee of %s.",
				solution:            "NASSAU",
				completionQuote:     "Nassau. That's not a hideout, that's a tourist pamphlet.",
				completionNarration: "A second strip of leather just hit the desk. Same pin, same mess.",
			},
		},
	},
	{
		tag: "FOLLOW-UP INTERCEPT",
		intro: []string{
			"Same wrap, same pin, same complete inability to write in a straight line.",
			"This one names the prize. If it's silver, the convoy office is going to cry.",
		},
		prompt: "Name the prey and we can spoil dinner.",
		payloads: []choice{
			{
				message:             "Hunt the ship flying %s.",
				solution:            "SILVER",
				completionQuote:     "Silver. The letters were all here. They were just drunk.",
				completionNarration: "The convoy can be warned before Widdershins finishes her rum.",
			},
			{
				message:             "The prize is the %s galleon.",
				solution:            "SPANISH",
				completionQuote:     "The Spanish galleon. Patch spelled it in a circle. Twice.",
				completionNarration: "Admiralty will love this. Butts will love a bucket.",
			},
			{
				message:             "Ignore fishermen. Take the %s.",
				solution:            "BULLION",
				completionQuote:     "Bullion. At least she's honest about the hobby.",
				completionNarration: "If they want the gold, we can be standing on it first.",
			},
		},
	},
}

// scripts: columnarIntercepts; expected: ColumnarCipher, keyword / column count
var columnarIntercepts = []interceptStage{
	{
		tag: "INTERCEPTED TRANSMISSION",
		intro: []string{
			"Patch upgraded, smartpants. Sort of.",
			"",
			"The next orders are written down the " + quoteAccent("ledger columns", " of his account book."),
			"He still can't spell 'quarantine,' but he can count to five if you don't rush him.",
		},
		prompt: "Read down the columns. Which port do they hit?",
		payloads: []choice{
			{
				message:             "We cut out of %s at slack tide.",
				solution:            "HAVANA",
				completionQuote:     "Havana. That's a customs house with a drinking problem.",
				completionNarration: "Second page of the ledger just fell out. Typical.",
			},
			{
				message:             "Raid the customs house at %s.",
				solution:            "KINGSTON",
				completionQuote:     "Kingston. The governor is going to need a new hobby.",
				completionNarration: "Butts says if you decode the warehouse next, he might keep lunch down.",
			},
			{
				message:             "The governor sleeps in %s.",
				solution:            "PORTROYAL",
				completionQuote:     "Port Royal. Loud, wet, and already on fire in spirit.",
				completionNarration: "There's a warehouse column on the back. Of course there is.",
			},
		},
	},
	{
		tag:      "FOLLOW-UP INTERCEPT",
		preamble: "*Patch has numbered the columns. Two of the numbers are 7.*\n",
		intro: []string{
			"Same ledger, same tide, slightly more ink than paper.",
			"This one is the building — warehouse, customs, or the place they keep the boom.",
		},
		prompt: "Name the door and we can lock it from the inside.",
		payloads: []choice{
			{
				message:             "Second warehouse on %s.",
				solution:            "WHARF",
				completionQuote:     "The wharf. Patch drew a little skull in the margin. Professional.",
				completionNarration: "Harbor watch can look busy near warehouse two.",
			},
			{
				message:             "The silver is under the %s.",
				solution:            "CUSTOMS",
				completionQuote:     "Under the customs house. They're stealing from the people who steal legally.",
				completionNarration: "Butts wants the floorboards counted. He is very tired.",
			},
			{
				message:             "Blow the lock on the %s.",
				solution:            "POWDER",
				completionQuote:     "The powder magazine. Subtle, Widdershins. Really blending in.",
				completionNarration: "If they want a bang, we can arrange a quieter evening.",
			},
		},
	},
}

// scripts: railFenceIntercepts; expected: RailFenceCipher, rail count
var railFenceIntercepts = []interceptStage{
	{
		tag: "INTERCEPTED TRANSMISSION",
		intro: []string{
			"New handwriting, same crime.",
			"",
			"Widdershins had Patch write this one zigzag — like a drunk sailor",
			"climbing the " + quoteAccent("ratlines", ". Up, down, up, down, leftover rum."),
			"It's a boarding signal. We'd like to not be boarded.",
		},
		prompt: "Follow the zigzag and give us the signal.",
		payloads: []choice{
			{
				message:             "Raise the black when you see %s.",
				solution:            "LANTERN",
				completionQuote:     "A lantern. Very mysterious. Also very on fire.",
				completionNarration: "There's a second zigzag about the watch. Patch is proud of himself.",
			},
			{
				message:             "Three lamps on the %s.",
				solution:            "MIZZEN",
				completionQuote:     "The mizzen. If you need three lamps, you are not hiding.",
				completionNarration: "Timing's in the next rail. Try not to get seasick on the letters.",
			},
			{
				message:             "The boarding cry is %s.",
				solution:            "CUTLASS",
				completionQuote:     "CUTLASS. She yelled the password. In the cipher. Incredible.",
				completionNarration: "Butts wants the hour next. He has a bucket and a dream.",
			},
		},
	},
	{
		tag: "FOLLOW-UP INTERCEPT",
		intro: []string{
			"Same rails, same ratlines, slightly more spelling.",
			"This one is when they board. If you get it, the convoy can look like a navy.",
		},
		prompt: "Name the watch and we can be wearing it.",
		payloads: []choice{
			{
				message:             "We board on the %s bell.",
				solution:            "DOG",
				completionQuote:     "The dog watch. Even the schedule is untrained.",
				completionNarration: "Night bells we can fake. Enthusiasm we cannot.",
			},
			{
				message:             "Wait for the %s.",
				solution:            "FOG",
				completionQuote:     "Fog. Nature's 'are you sure about this.'",
				completionNarration: "If they wait for weather, we can arrive as weather.",
			},
			{
				message:             "Strike after the %s.",
				solution:            "SALUTE",
				completionQuote:     "They salute their own fake navy flag. I need a transfer.",
				completionNarration: "Butts says if they want ceremony, we can give them an audience.",
			},
		},
	},
}

// scripts: routeIntercepts; expected: RouteCipher, spiral / deck plan
var routeIntercepts = []interceptStage{
	{
		tag:      "INTERCEPTED TRANSMISSION",
		preamble: "*Someone spilled grog on a deck plan and called it encryption.*\n",
		intro: []string{
			"Last trick from Widdershins, then I want dry land.",
			"",
			"Patch wrote this around the " + quoteAccent("deck from the helm", ", spiral in, curse outward."),
			"It's the stash. If we beat them to the X, she can rearrange someone else's vowels.",
		},
		prompt: "Walk the deck and tell us where the loot sits.",
		payloads: []choice{
			{
				message:             "Bury it under the %s.",
				solution:            "PALM",
				completionQuote:     "Under the palm. Original. A credit to literature.",
				completionNarration: "Escape heading is on the rim of the same map. Naturally.",
			},
			{
				message:             "Cache at %s rock.",
				solution:            "DEADMAN",
				completionQuote:     "Deadman's Rock. She really leaned into the genre.",
				completionNarration: "If that's the stash, the getaway is the next loop.",
			},
			{
				message:             "No maps. The X is %s.",
				solution:            "ELSEWHERE",
				completionQuote:     "ELSEWHERE. That's not a place, Patch. That's a shrug.",
				completionNarration: "Butts just laughed, which is worrying. Decode the run next.",
			},
		},
	},
	{
		tag: "FOLLOW-UP INTERCEPT",
		intro: []string{
			"Same spiral, starting at the wheel again.",
			"This is the coward's clause — where they run if the silver bites back.",
		},
		prompt: "Name the heading and we can be waiting on it.",
		payloads: []choice{
			{
				message:             "If it goes sour, run for %s.",
				solution:            "BARBADOS",
				completionQuote:     "Barbados. Warm, smug, and already overbooked.",
				completionNarration: "The Admiralty can have a welcoming committee. Butts wants a nap.",
			},
			{
				message:             "Scuttle the prize off %s.",
				solution:            "HATTERAS",
				completionQuote:     "Hatteras. Graveyard of ships, playground of bad ideas.",
				completionNarration: "If they dump the prize, we can still be the ones with a net.",
			},
			{
				message:             "Fair wind to %s.",
				solution:            "PROVIDENCE",
				completionQuote:     "Providence. She wants luck and a harbor. She can have neither.",
				completionNarration: "Nice work, smartpants. If I never see a belaying pin again, I can die happy.",
			},
		},
	},
}
