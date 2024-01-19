package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kovalyov-valentin/profiles-service/internal/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	var input models.Users

	if err := ctx.BindJSON(&input); err != nil {
		logrus.Error("Failed to bind JSON:", err)
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if input.Name == "" || input.Surname == "" {
		errMsg := "Field Name and Surname are required fields"
		logrus.Error(errMsg)
		newErrorResponse(ctx, http.StatusBadRequest, errors.New(errMsg).Error())
		return
	}

	id, err := h.services.CreateUser(ctx, input)
	if err != nil {
		logrus.WithError(err).Error("Failed to create record:", err)
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Infof("Record created successfully. ID: %d", id)

	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		logrus.WithError(err).Error("Failed to convert ID to integer")
		ctx.JSON(http.StatusInternalServerError, errors.New("invalid user ID").Error())
	}

	record, err := h.services.GetUser(ctx, id)
	if err != nil {
		logrus.WithError(err).Error("Failed to convert ID to integer")
		newErrorResponse(ctx, http.StatusBadRequest, "error: no record with you id")
		return
	}
	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, record)
}

func (h *Handler) GetUsers(ctx *gin.Context) {
	listsRecords, err := h.services.GetUsers(ctx)
	if err != nil {
		logrus.WithError(err).Error("Failed to get all records")
		newErrorResponse(ctx, http.StatusBadRequest, "error: failed to get all records")
		return
	}

	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, listsRecords)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		logrus.WithError(err).Error("Failed to convert user ID to integer")
		ctx.JSON(http.StatusInternalServerError, errors.New("invalid user ID").Error())
	}

	var input models.Users
	if err := ctx.BindJSON(&input); err != nil {
		logrus.WithError(err).Error("Failed to bind JSON input")
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateUser(ctx, id, input); err != nil {
		logrus.WithError(err).Error("Failed to update record")
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		logrus.WithError(err).Error("Invalid user ID")
		ctx.JSON(http.StatusInternalServerError, errors.New("invalid user ID").Error())
	}

	err = h.services.DeleteUser(ctx, id)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "error: no record with you id")
		return
	}

	logrus.WithField("ID", id).Info("Record deleted successfully")
	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
