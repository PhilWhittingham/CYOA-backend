package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	// AdventItem represents a single item in the advent calendar
	//   NB this is added as a struct for extensibility reasons
	AdventItem struct {
		Day  int    `json:"day"`
		Text string `json:"text"`
	}
)

type (
	// PlayerDetails represents the data for a player during a single
	// playthrough of CYOAdvent.
	PlayerDetails struct {
		VisitedForest  bool   `json:"visitedForest"`
		VisitedVillage bool   `json:"visitedVillage"`
		VisitedTree    bool   `json:"visitedTree"`
		VisitedHotel   bool   `json:"visitedHotel"`
		VisitedShop    bool   `json:"visitedShop"`
		WizardFriend   bool   `json:"wizardFriend"`
		Cursed         bool   `json:"cursed"`
		Rested         bool   `json:"rested"`
		Location       string `json:"location"`
	}
)

var (
	adventItems = map[int]*AdventItem{
		1: &AdventItem{
			Day:  1,
			Text: "Cereals best friend to help the day go by milky smooth",
		},
		2: &AdventItem{
			Day:  2,
			Text: "Grizzly friends. Perfect for a movie night",
		},
		3: &AdventItem{
			Day:  3,
			Text: "You can't spell \"functional\" without \"fun\"",
		},
		4: &AdventItem{
			Day:  4,
			Text: "Samin would be proud",
		},
		5: &AdventItem{
			Day:  5,
			Text: "Schmeckt den Regenbogen",
		},
		6: &AdventItem{
			Day:  6,
			Text: "Schmeckt den Regenbogen mit leckeres Pesto",
		},
		7: &AdventItem{
			Day:  7,
			Text: "Fast but not furious",
		},
		8: &AdventItem{
			Day:  8,
			Text: "It may be time to defrost the pizza dough",
		},
		9: &AdventItem{
			Day:  9,
			Text: "A british tradition",
		},
		10: &AdventItem{
			Day:  10,
			Text: "Ginger sold separately",
		},
		11: &AdventItem{
			Day:  11,
			Text: "Enemies of the heir... beware",
		},
		12: &AdventItem{
			Day:  12,
			Text: "I might have fudged this one up",
		},
		13: &AdventItem{
			Day:  13,
			Text: "But what if bread... Was sauce...",
		},
		14: &AdventItem{
			Day:  14,
			Text: "Anything from the trolly?",
		},
		15: &AdventItem{
			Day:  15,
			Text: "Best (only) Caribbean food in town.",
		},
		16: &AdventItem{
			Day:  16,
			Text: "Test test",
		},
		17: &AdventItem{
			Day:  17,
			Text: "Do you want to build a snowman?",
		},
		18: &AdventItem{
			Day:  18,
			Text: "We have Nutella at home",
		},
		19: &AdventItem{
			Day:  19,
			Text: "Many hugs in many mugs",
		},
		20: &AdventItem{
			Day:  20,
			Text: "Jelly? Gummy? Bears? Babies? Cola?",
		},
		21: &AdventItem{
			Day:  21,
			Text: "Probably doesn't melt very well",
		},
		22: &AdventItem{
			Day:  22,
			Text: "I searched for \"vegan buchers\" and this is what came up",
		},
		23: &AdventItem{
			Day:  23,
			Text: "Laces? Lances? Belts?",
		},
		24: &AdventItem{
			Day:  24,
			Text: "Getting board of this",
		},
	}
	playerDetails = PlayerDetails{
		VisitedForest:  false,
		VisitedVillage: false,
		VisitedTree:    false,
		VisitedShop:    false,
		VisitedHotel:   false,
		WizardFriend:   false,
		Cursed:         false,
		Rested:         false,
		Location:       "Start",
	}
)

//----------
// Handlers - Advent Items
//----------

func getAllAdventItems(c echo.Context) error {
	return c.JSON(http.StatusOK, adventItems)
}

func getAdventItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, adventItems[id])
}

func updateAdventItem(c echo.Context) error {
	u := new(AdventItem)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	adventItems[id].Text = u.Text
	return c.JSON(http.StatusOK, adventItems[id])
}

func deleteAdventItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(adventItems, id)
	return c.NoContent(http.StatusNoContent)
}

//----------
// Handlers - Player
//----------

func getPlayerDetails(c echo.Context) error {
	return c.JSON(http.StatusOK, playerDetails)
}

func updateVisitedForest(c echo.Context) error {
	playerDetails.VisitedForest = true
	return c.JSON(http.StatusOK, playerDetails.VisitedForest)
}

func updateVisitedVillage(c echo.Context) error {
	playerDetails.VisitedVillage = true
	return c.JSON(http.StatusOK, playerDetails.VisitedVillage)
}

func updateVisitedTree(c echo.Context) error {
	playerDetails.VisitedTree = true
	return c.JSON(http.StatusOK, playerDetails.VisitedTree)
}

func updateVisitedShop(c echo.Context) error {
	playerDetails.VisitedShop = true
	return c.JSON(http.StatusOK, playerDetails.VisitedShop)
}

func updateVisitedHotel(c echo.Context) error {
	playerDetails.VisitedHotel = true
	return c.JSON(http.StatusOK, playerDetails.VisitedHotel)
}

func updateWizardFriend(c echo.Context) error {
	playerDetails.WizardFriend = true
	return c.JSON(http.StatusOK, playerDetails.WizardFriend)
}

func updateCursed(c echo.Context) error {
	playerDetails.Cursed = true
	return c.JSON(http.StatusOK, playerDetails.Cursed)
}

func updateRested(c echo.Context) error {
	playerDetails.Rested = true
	return c.JSON(http.StatusOK, playerDetails.Rested)
}

func updateLocation(c echo.Context) error {
	p := new(PlayerDetails)
	if err := c.Bind(p); err != nil {
		return err
	}
	playerDetails.Location = p.Location
	return c.JSON(http.StatusOK, playerDetails.Location)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Routes
	e.GET("/adventItems", getAllAdventItems)
	e.GET("/adventItems/:id", getAdventItem)
	e.PUT("/adventItems/:id", updateAdventItem)
	e.DELETE("/adventItems/:id", deleteAdventItem)

	e.GET("/player", getPlayerDetails)
	e.POST("/player/forest", updateVisitedForest)
	e.POST("/player/village", updateVisitedVillage)
	e.POST("/player/tree", updateVisitedTree)
	e.POST("/player/shop", updateVisitedShop)
	e.POST("/player/hotel", updateVisitedHotel)
	e.POST("/player/wizard", updateWizardFriend)
	e.POST("/player/cursed", updateCursed)
	e.POST("/player/rested", updateRested)
	e.PUT("/player/location", updateLocation)

	// Start server
	e.Logger.Fatal(e.Start(":1333"))
}
