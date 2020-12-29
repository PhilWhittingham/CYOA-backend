package item

// Detail represents a single item in the advent calendar
//   NB this is added as a struct for extensibility reasons
type Detail struct {
	Day  int    `json:"day"`
	Text string `json:"text"`
}

// ItemsList is a full list of all 24 advent items
var ItemsList = map[int]*Detail{
	1: {
		Day:  1,
		Text: "Cereals best friend to help the day go by milky smooth",
	},
	2: {
		Day:  2,
		Text: "Grizzly friends. Perfect for a movie night",
	},
	3: {
		Day:  3,
		Text: "You can't spell \"functional\" without \"fun\"",
	},
	4: {
		Day:  4,
		Text: "Samin would be proud",
	},
	5: {
		Day:  5,
		Text: "Schmeckt den Regenbogen",
	},
	6: {
		Day:  6,
		Text: "Schmeckt den Regenbogen mit leckeres Pesto",
	},
	7: {
		Day:  7,
		Text: "Fast but not furious",
	},
	8: {
		Day:  8,
		Text: "It may be time to defrost the pizza dough",
	},
	9: {
		Day:  9,
		Text: "A british tradition",
	},
	10: {
		Day:  10,
		Text: "Ginger sold separately",
	},
	11: {
		Day:  11,
		Text: "Enemies of the heir... beware",
	},
	12: {
		Day:  12,
		Text: "I might have fudged this one up",
	},
	13: {
		Day:  13,
		Text: "But what if bread... Was sauce...",
	},
	14: {
		Day:  14,
		Text: "Anything from the trolly?",
	},
	15: {
		Day:  15,
		Text: "Best (only) Caribbean food in town.",
	},
	16: {
		Day:  16,
		Text: "Doing 24 things is hard",
	},
	17: {
		Day:  17,
		Text: "Do you want to build a snowman?",
	},
	18: {
		Day:  18,
		Text: "We have Nutella at home",
	},
	19: {
		Day:  19,
		Text: "Many hugs in many mugs",
	},
	20: {
		Day:  20,
		Text: "Jelly? Gummy? Bears? Babies? Cola?",
	},
	21: {
		Day:  21,
		Text: "Probably doesn't melt very well",
	},
	22: {
		Day:  22,
		Text: "I searched for \"vegan buchers\" and this is what came up",
	},
	23: {
		Day:  23,
		Text: "Laces? Lances? Belts?",
	},
	24: {
		Day:  24,
		Text: "Getting board of this",
	},
}
