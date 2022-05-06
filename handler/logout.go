package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *profileHandler) Logout(c *gin.Context) {
	//If metadata is passed and the tokens valid, delete them from the redis store
	metadata, _ := h.tk.ExtractTokenMetadata(c.Request)
	if metadata != nil {
		deleteErr := h.rd.DeleteTokens(metadata)
		if deleteErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"detail": deleteErr.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"msg": "successfully logged out"})
}
