package handlers

/* --- This package handles connections and queries for the albums dataset -- */

import (
	"net/http"
	"strconv"

	"example.com/data"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Add a new album to the Database
func InsertAlbum(c *gin.Context) {
	var newAlbum data.AlbumInsert
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	sqlStatement := `
    INSERT INTO album (title, artists, price)
    VALUES ($1, $2, $3)
    RETURNING id;
  `
	db := data.OpenConn()
	defer db.Close()

	// Create a place to store the newly assigned ID
	var id int64

	// Insert the row and return the values in the new row to print for debugging
	r := db.QueryRow(sqlStatement, newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	err := r.Scan(&id)
	if err != nil {
		panic(err)
	}

	// Create Album from AlbumInsert and return id
	albumwithID := data.Album{
		ID:     id,
		Title:  newAlbum.Title,
		Artist: newAlbum.Artist,
		Price:  newAlbum.Price,
	}

	// return JSON response to client
	c.IndentedJSON(http.StatusCreated, albumwithID)
}

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	// Read all values in Album Table from Postgres
	sqlStatement := `SELECT * FROM album`
	db := data.OpenConn()
	defer db.Close()
	// Store rows from database in a list by iterating through the SQLResponse
	var albums []data.Album

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb data.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			panic(err)
		}
		albums = append(albums, alb)
	}
	// Return the JSON response to the client
	c.IndentedJSON(http.StatusOK, albums)
}

// Matches Artist name to existing list and returns all matches
func GetAlbumsByAuthorName(c *gin.Context) {
	// Take an argument from REST API POST
	name := c.Param("name")
	// Use argument to search database for matching entries
	// (SQL Injection? Look for alternatives or suggest parsing)
	sqlStatement := `SELECT * FROM album WHERE artists ILIKE $1`

	db := data.OpenConn()
	defer db.Close()
	// Store rows from return query by iterating through SQL Response object
	var albums []data.Album

	rows, err := db.Query(sqlStatement, name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb data.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			panic(err)
		}
		albums = append(albums, alb)
	}

	// Return JSON response to client
	c.IndentedJSON(http.StatusOK, albums)
}

// Grab Albums where the prices are less than value passed
func GetAlbumsPriceLessThan(c *gin.Context) {
	// Parse argument provided into float and use to filter database results
	price_temp := c.Param("price")
	price, err := strconv.ParseFloat(price_temp, 64)
	if err != nil {
		panic(err)
	}

	sqlStatement := `SELECT * FROM album WHERE price < ?`

	db := data.OpenConn()
	defer db.Close()

	// Iterate through response and store in array
	var albums []data.Album

	rows, err := db.Query(sqlStatement, price)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb data.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			panic(err)
		}
		albums = append(albums, alb)
	}

	// Return JSON response to client
	c.IndentedJSON(http.StatusOK, albums)
}
