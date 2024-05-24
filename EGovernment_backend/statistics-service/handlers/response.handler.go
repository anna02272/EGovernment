package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"statistics-service/domain"
	errorMessage "statistics-service/error"
	"statistics-service/services"
	"time"
)

type ResponseHandler struct {
	service services.ResponseService
	DB      *mongo.Collection
}

func NewResponseHandler(service services.ResponseService, db *mongo.Collection) ResponseHandler {
	return ResponseHandler{
		service: service,
		DB:      db,
	}
}

func (r *ResponseHandler) Create(c *gin.Context) {
	rw := c.Writer
	h := c.Request

	token := h.Header.Get("Authorization")
	user, err := r.GetCurrentUserFromAuthService(token, rw)
	if err != nil {
		return
	}

	if user.UserRole != domain.Employee {
		errorMessage.ReturnJSONError(rw, map[string]string{"error": "Unauthorized. You are not an employee."}, http.StatusUnauthorized)
		return
	}
	var requestBody struct {
		Text       string `json:"text" binding:"required"`
		Attachment string `json:"attachment"`
		Accepted   bool   `json:"accepted" binding:"required"`
		SendTo     string `json:"send_to" binding:"required"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	currentDateTime := primitive.NewDateTimeFromTime(time.Now())
	id := primitive.NewObjectID()
	newResponse := &domain.Response{
		ID:         id,
		Text:       requestBody.Text,
		Attachment: requestBody.Attachment,
		Accepted:   requestBody.Accepted,
		SendTo:     requestBody.SendTo,
		Date:       currentDateTime,
	}

	err, _ = r.service.Create(newResponse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create response"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Response successfully saved", "response": newResponse})
}

func (r *ResponseHandler) GetAll(c *gin.Context) {
	rw := c.Writer
	h := c.Request

	token := h.Header.Get("Authorization")
	user, err := r.GetCurrentUserFromAuthService(token, rw)
	if err != nil {
		return
	}

	if user.UserRole != domain.Employee {
		errorMessage.ReturnJSONError(rw, map[string]string{"error": "Unauthorized. You are not an employee."}, http.StatusUnauthorized)
		return
	}

	responses, err := r.service.GetAll()
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve responses from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(responses)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (r *ResponseHandler) GetByID(c *gin.Context) {
	rw := c.Writer
	h := c.Request
	id := c.Param("id")

	token := h.Header.Get("Authorization")
	user, err := r.GetCurrentUserFromAuthService(token, rw)
	if err != nil {
		return
	}

	if user.UserRole != domain.Employee {
		errorMessage.ReturnJSONError(rw, map[string]string{"error": "Unauthorized. You are not an employee."}, http.StatusUnauthorized)
		return
	}

	response, err := r.service.GetById(id)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve response from the database.No such response."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (r *ResponseHandler) GetCurrentUserFromAuthService(token string, rw http.ResponseWriter) (*domain.User, error) {
	url := "http://auth-service:8085/api/users/currentUser"

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := r.performAuthorizationRequestWithContext("GET", ctx, token, url)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			errorMessage.ReturnJSONError(rw, map[string]string{"error": "Authorization service is not available."}, http.StatusBadRequest)
			return nil, err
		}
		errorMessage.ReturnJSONError(rw, map[string]string{"error": "Error performing authorization request."}, http.StatusBadRequest)
		return nil, err
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode != http.StatusOK {
		errorMessage.ReturnJSONError(rw, map[string]string{"error": "Unauthorized."}, http.StatusUnauthorized)
		return nil, errors.New("Unauthorized")
	}

	decoder := json.NewDecoder(resp.Body)
	var responseUser struct {
		LoggedInUser struct {
			UserRole domain.UserRole `json:"userRole"`
		} `json:"user"`
	}
	if err := decoder.Decode(&responseUser); err != nil {
		errorMessage.ReturnJSONError(rw, map[string]string{"error": "User object was not valid."}, http.StatusUnauthorized)
		return nil, err
	}

	return &domain.User{UserRole: responseUser.LoggedInUser.UserRole}, nil
}

func (r *ResponseHandler) performAuthorizationRequestWithContext(method string, ctx context.Context, token string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
