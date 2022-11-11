package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Client struct {
	ID        string   `json:"id"`
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

type Address struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
}

var clients = []Client{
	{ID: "1", FirstName: "Peter", LastName: "Parker", Address: &Address{Country: "US", State: "NY", City: "New York"}},
	{ID: "2", FirstName: "Lucifer", LastName: "Morningstar", Address: &Address{Country: "US", State: "CA", City: "Los Angeles"}},
	{ID: "3", FirstName: "Luis", LastName: "Silva", Address: &Address{Country: "BR", State: "DF", City: "Brasília"}},
}

func getClients(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, clients)
}

func getClientByID(c *gin.Context) {
	id := c.Param("id")

	for _, client := range clients {
		if client.ID == id {
			c.IndentedJSON(http.StatusOK, client)
           return 
		}
	}
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Cliente não encontrado"})
}

func createClient(c *gin.Context) {
	var newClient Client

	err := c.BindJSON(&newClient)
	if err != nil {
		return
	}

	clients = append(clients, newClient)
	c.IndentedJSON(http.StatusCreated, newClient)
}

func deleteClient(c *gin.Context) {
    clientToDelete := c.Param("id")
    for index, client := range clients{
        if client.ID == clientToDelete {
            clients = append(clients[:index], clients[index+1:]...)
            c.IndentedJSON(http.StatusOK, clients)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Cliente não encontrado"})
}

func main() {

	router := gin.Default()

	router.GET("/clients", getClients)
    router.GET("/clients/:id", getClientByID)
	router.POST("/clients", createClient)
    router.DELETE("/clients/:id", deleteClient)

	router.Run(":3000")
}
