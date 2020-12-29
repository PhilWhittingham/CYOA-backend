package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/PhilWhittingham/CYOA-backend/item"
	"github.com/PhilWhittingham/CYOA-backend/player"
)

var playerDetails = new(player.Details)

//----------
// Handlers - Advent Items
//----------

func getAllAdventItems(c echo.Context) error {
	return c.JSON(http.StatusOK, item.ItemsList)
}

func getAdventItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, item.ItemsList[id])
}

func updateAdventItem(c echo.Context) error {
	u := new(item.Detail)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	item.ItemsList[id].Text = u.Text
	return c.JSON(http.StatusOK, item.ItemsList[id])
}

func deleteAdventItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(item.ItemsList, id)
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
	p := new(player.Details)
	if err := c.Bind(p); err != nil {
		return err
	}
	playerDetails.Location = p.Location
	return c.JSON(http.StatusOK, playerDetails.Location)
}

// Initialise the api with routing and variables as necessary
func Initialise(pd *player.Details) {
	playerDetails = pd
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
