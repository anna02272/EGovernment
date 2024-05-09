package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
	"vehicles-service/data"
	"vehicles-service/domain"
	errorMessage "vehicles-service/error"
	"vehicles-service/services"
)

type VehicleDriverHandler struct {
	service services.VehicleDriverService
	DB      *mongo.Collection
}

func NewVehicleDriverHandler(service services.VehicleDriverService, db *mongo.Collection) VehicleDriverHandler {
	return VehicleDriverHandler{
		service: service,
		DB:      db,
	}

}

func (s *VehicleDriverHandler) CreateVehicleDriver(c *gin.Context) {
	rw := c.Writer
	h := c.Request

	token := h.Header.Get("Authorization")
	url := "http://auth-service:8085/api/users/currentUser"

	timeout := 5 * time.Second // Adjust the timeout duration as needed
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := s.performAuthorizationRequestWithContext("GET", ctx, token, url)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			errorMsg := map[string]string{"error": "Authorization service is not available."}
			errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
			return
		}

		errorMsg := map[string]string{"error": "Error performing authorization request."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return

	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode != 200 {
		errorMsg := map[string]string{"error": "Unauthorized."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(resp.Body)

	// Define a struct to represent the JSON structure
	var responseUser struct {
		LoggedInUser struct {
			username string        `json:"username"`
			email    string        `json:"email"`
			UserRole data.UserRole `json:"userRole"`
		} `json:"user"`
	}

	if err := decoder.Decode(&responseUser); err != nil {
		errorMsg := map[string]string{"error": "User object was not valid."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	if responseUser.LoggedInUser.UserRole != data.Policeman {
		errorMsg := map[string]string{"Unauthorized": " You are not policeman."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	vehicleDriver, exists := c.Get("vehicleDriver")
	if !exists {
		errorMsg := map[string]string{"Error": " vehicleDriver object was not valid."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}
	vehicleDriverInsert, ok := vehicleDriver.(domain.VehicleDriverCreate)
	if !ok {
		errorMsg := map[string]string{"error": "Invalid type for vehicle driver."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	vehicleDriverInsertDB, _, err := s.service.InsertVehicleDriver(&vehicleDriverInsert)
	if err != nil {
		errorMsg := map[string]string{"error": "Database problem."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	jsonResponse, err1 := json.Marshal(vehicleDriverInsertDB)
	if err1 != nil {
		errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error marshaling JSON: %s", err1), http.StatusInternalServerError)
		return
	}
	rw.Write(jsonResponse)

}

func (s *VehicleDriverHandler) performAuthorizationRequestWithContext(method string, ctx context.Context, token string, url string) (*http.Response, error) {
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
