package player

// Details represents the data for a player during a single
// playthrough of CYOAdvent.
type Details struct {
	VisitedForest  bool   `json:"visitedForest"`
	VisitedVillage bool   `json:"visitedVillage"`
	VisitedTree    bool   `json:"visitedTree"`
	VisitedHotel   bool   `json:"visitedHotel"`
	VisitedShop    bool   `json:"visitedShop"`
	WizardFriend   bool   `json:"wizardFriend"`
	Cursed         bool   `json:"cursed"`
	Rested         bool   `json:"rested"`
	Location       string `json:"location"`
	ID             string `json:"ID"`
}

// NewPlayer returns a new Details object for the player
func NewPlayer(id string) *Details {
	return &Details{
		ID:       id,
		Location: "Start",
	}
}
