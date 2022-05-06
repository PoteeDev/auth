package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/PoteeDev/auth/storage"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s took %v\n", what, time.Since(start))
	}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (h *profileHandler) Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"datail": "invalid json provided"})
		return
	}
	//compare the user from the request, with the one we defined:
	team, err := storage.GetAuthTeam(u.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"datail": "invalid credentials"})
		return
	}
	if !CheckPasswordHash(u.Password, team.Hash) {
		c.JSON(http.StatusUnauthorized, gin.H{"datail": "invalid credentials"})
		return
	}
	genErr, tokens := h.generateTokens(team.Login)
	if genErr != nil {
		c.JSON(http.StatusUnprocessableEntity, genErr.Error())
		return
	}
	c.JSON(http.StatusOK, tokens)
}
