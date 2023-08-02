package controllers

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/anuj0809/Backend_AS/database"
	"github.com/anuj0809/Backend_AS/models"
	"github.com/gofiber/fiber/v2"
)

// create players
func CratePlayer(c *fiber.Ctx) error {
	// checks if the specified content types are acceptable.
	c.Accepts("application/json")
	var player models.Players
	// parse the json data and check for error
	if err := c.BodyParser(&player); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to parse players data.",
		})
	}
	//insert a new player into the database
	result := database.DB.Create(&player)
	// check for error
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create a new player.",
		})
	}
	return c.Status(http.StatusCreated).JSON(player)
}

// get all players
func GetAllPlayers(c *fiber.Ctx) error {
	var players []models.Players
	// query to get all theplayers in decendinf order of thier scores
	result := database.DB.Order("score DESC").Find(&players)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch players data from the database.",
		})
	}
	return c.Status(http.StatusOK).JSON(&players)
}

// get a random player
func GetRandomPlayer(c *fiber.Ctx) error {
	var player models.Players
	var count int64
	// get the count of number of players in the database
	database.DB.Model(&models.Players{}).Count(&count)
	if count == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "No players found in the database.",
		})
	}
	// select a random number where max will be the count
	random := rand.Intn(int(count))
	// fetch a random player using index
	result := database.DB.Offset(random).Limit(1).First(&player)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch a random players data from the database.",
		})
	}
	return c.Status(http.StatusOK).JSON(&player)

}

// get player on rank
func GetPlayerByRank(c *fiber.Ctx) error {
	var player models.Players
	// get the rank value from the param
	// convert it into int from string
	rank, err := strconv.Atoi(c.Params("val"))
	if err != nil || rank <= 0 {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid rank value.",
		})
	}
	// query to get the first player based on the rank entered
	result := database.DB.Order("score DESC").Offset(rank - 1).Limit(1).First(&player)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Player with this rank does not exist.",
		})
	}
	return c.Status(http.StatusOK).JSON(&player)

}

// update player
func UpdatePlayer(c *fiber.Ctx) error {
	// get the rank value from the param
	// convert it into int from string
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid player ID.",
		})
	}
	var updatedPlayer models.Players
	c.Accepts("application/json")
	if err = c.BodyParser(&updatedPlayer); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to parse players data.",
		})
	}
	var player models.Players
	// get the first player based on the id
	result := database.DB.First(&player, id)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Player with this ID does not exist.",
		})
	}
	// if updated, update the following values in the database
	if updatedPlayer.Name != "" {
		player.Name = updatedPlayer.Name
	}
	if updatedPlayer.Score != 0 {
		player.Score = updatedPlayer.Score
	}
	// save the updated values to the database
	result = database.DB.Save(&player)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Updating player data failed.",
		})
	}
	return c.Status(http.StatusOK).JSON(&player)
}

func DeletePlayer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid player ID.",
		})
	}
	var player models.Players
	// query to delete the player based on the given id
	result := database.DB.First(&player, id)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Player with this ID does not exist.",
		})
	}
	result = database.DB.Delete(&player)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Deleting player data failed.",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Deleted Successfully.",
	})
}
