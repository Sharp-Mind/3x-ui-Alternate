package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"x-ui/database/model"
)

func (h *Handler) GetInboundClients(ctx *gin.Context) {
	// Получаем ID
	id := ctx.Param("id")
	inboundId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid inbound ID"})
		return
	}
	//	Получаем клиентов
	clients, err := h.services.InboundClient.GetInboundClients(inboundId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"clients": clients})
}

func (h *Handler) AddInboundClient(ctx *gin.Context) {
	//	Получаем ID
	id := ctx.Param("id")
	//id := 2
	inboundId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid inbound ID"})
		return
	}
	//	Получаем клиента
	var newClient model.Client
	err = ctx.ShouldBindJSON(&newClient)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Получаем логин и пароль из Basic Auth
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.JSON(400, gin.H{"error": "Missing Basic Auth credentials"})
		return
	}
	//log.Println(username, password)
	//	Проверяем пользователя
	if !h.services.User.CheckUser(username, password) {
		ctx.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	//	Добавляем клиента
	key, _, err := h.services.InboundClient.AddInboundClient(inboundId, &newClient)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"" +
		"message": "Client added successfully",
		"result": key,
	})
}
