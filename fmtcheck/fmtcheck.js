// Returns true if card is not legal in harmony
function nonlegalHarmony(card) {
	const banned = ["eo103", "pr69", "pr83", "cn42", "pr175", "cg136", "cn47", "pr16", "cn65", "hm150", "sb134", "de55", "add1"];
	if (banned.indexOf(card) >= 0) {
		return true;
	}
	return false;
}

// Returns true if card is not legal in core
function nonlegalCore(card) {
	const banned = ["eo2", "eo103", "hm150", "sb134"];
	const sets = ["eo", "hm", "mt", "de", "sb", "ff"];
	if (banned.indexOf(card) >= 0) {
		return true;
	}
	for (const set of sets) {
		if (card.startsWith(set)) {
			return false;
		}
	}
	return true;
}

function check(urlstr) {
	const url = new URL(urlstr);
	const decklist = url.searchParams.get("v1code");
	const cards = decklist.split("-").map(c => c.split("x")[0]);

	const illegalHarmonyCards= cards.filter(nonlegalHarmony);
	const illegalCoreCards = cards.filter(nonlegalCore);

	return {
		"Harmony": {
			legal: illegalHarmonyCards.length === 0,
			illegalcards: illegalHarmonyCards
		},
		"Core": {
			legal: illegalCoreCards.length === 0,
			illegalcards: illegalCoreCards
		}
	};
}