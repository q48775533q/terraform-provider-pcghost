package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/TyunTech/terraform-petstore/action/pets"
	"github.com/TyunTech/terraform-petstore/model/pet"
)

var db *gorm.DB

func init() {
	initializeRDSConn()
	validateRDS()
}

func initializeRDSConn() {
	user := os.Getenv("rds_user")
	password := os.Getenv("rds_password")
	host := os.Getenv("rds_host")
	port := os.Getenv("rds_port")
	database := os.Getenv("rds_database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

func validateRDS() {
	//If the pets table does not already exist, create it
	if !db.HasTable("pets") {
		db.CreateTable(&pet.Pet{})
	}
}

func optionsPetHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE")
	c.Header("Access-Control-Allow-Headers", "origin, content-type, accept")
}

func main() {
	r := gin.Default()
	r.POST("/api/pets", createPetHandler)
	r.GET("/api/pets/:id", getPetHandler)
	r.GET("/api/pets", listPetsHandler)
	r.PATCH("/api/pets/:id", updatePetHandler)
	r.DELETE("/api/pets/:id", deletePetHandler)
	r.OPTIONS("/api/pets", optionsPetHandler)
	r.OPTIONS("/api/pets/:id", optionsPetHandler)

	r.Run(":8000")
}

func createPetHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var req pets.CreatePetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := pets.CreatePet(db, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func listPetsHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	limit := 10
	if c.Query("limit") != "" {
		newLimit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			limit = 10
		} else {
			limit = newLimit
		}
	}
	if limit > 50 {
		limit = 50
	}
	req := pets.ListPetsRequest{Limit: uint(limit)}
	res, _ := pets.ListPets(db, &req)
	c.JSON(http.StatusOK, res)
}

func getPetHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	id := c.Param("id")
	req := pets.GetPetRequest{ID: id}
	res, _ := pets.GetPet(db, &req)
	if res == nil {
		c.JSON(http.StatusNotFound, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func updatePetHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var req pets.UpdatePetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	req.ID = id
	res, err := pets.UpdatePet(db, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func deletePetHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	id := c.Param("id")
	req := pets.DeletePetRequest{ID: id}
	err := pets.DeletePet(db, &req)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Status(http.StatusOK)
}
