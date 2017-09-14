package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	wol "github.com/voowoo/go-wol"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	error := parseArgs()
	if error != nil {
		os.Exit(1)
	}
	r := NewServer()
	r.Run()
}

// NewServer creates the applications gin server
func NewServer() *gin.Engine {
	r := gin.Default()
	r.POST("/wake", wake)
	return r
}

func wake(c *gin.Context) {
	broadcastIP := c.DefaultQuery("broadcastIP", options.BroadcastIP)
	mac := c.Query("mac")

	if mac == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "query parameter mac is missing",
		})
		return
	}

	macAdress, error := wol.NewMacAdressFrom(mac)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("failed to parse %s as a MAC adress", mac),
		})
		return
	}

	broadcastAdress, error := wol.NewBroadcastAdressFrom(broadcastIP)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("failed to parse %s as a IPv4 adress", broadcastIP),
		})
		return
	}

	error = wol.NewMagicPacket(macAdress).Send(broadcastAdress)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("failed to send a magic packet for %s to %s", macAdress, broadcastAdress),
		})
	} else {
		log.Printf("Sent magic packet for %s to %s", mac, broadcastIP)
		c.Status(204)
	}

	return
}
