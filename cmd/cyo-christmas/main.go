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
		VisitedWood    bool   `json:"visitedWood"`
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
	adventItems   = map[int]*AdventItem{}
	playerDetails = PlayerDetails{
		VisitedVillage: false,
		VisitedTree:    false,
		VisitedShop:    false,
		VisitedHotel:   false,
		WizardFriend:   false,
		Cursed:         false,
		Rested:         false,
		Location:       "Forest",
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

func getRemainingDays() int {
	return 1
}

//----------
// Handlers - Player
//----------

func getPlayerDetails(c echo.Context) error {
	return c.JSON(http.StatusOK, playerDetails)
}

func updateVisitedWood(c echo.Context) error {
	playerDetails.VisitedWood = true
	return c.JSON(http.StatusOK, playerDetails.VisitedWood)
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

	// Routes
	//e.POST("/adventItems", createAdventItem)
	e.GET("/adventItems", getAllAdventItems)
	e.GET("/adventItems/:id", getAdventItem)
	e.PUT("/adventItems/:id", updateAdventItem)
	e.DELETE("/adventItems/:id", deleteAdventItem)

	e.GET("/player", getPlayerDetails)
	e.POST("/player/wood", updateVisitedWood)
	e.POST("/player/village", updateVisitedVillage)
	e.POST("/player/tree", updateVisitedTree)
	e.POST("/player/shop", updateVisitedShop)
	e.POST("/player/hotel", updateVisitedHotel)
	e.POST("/player/wizard", updateWizardFriend)
	e.POST("/player/cursed", updateCursed)
	e.POST("/player/rested", updateRested)
	e.PUT("/player/location", updateLocation)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
