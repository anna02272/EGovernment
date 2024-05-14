package handlers

import (
	"context"
	"court-service/data"
	"court-service/domain"
	errorMessage "court-service/error"
	"court-service/services"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"net/http"
	"time"
)

type CitizenHandler struct {
	service services.CitizenService
	DB      *mongo.Collection
}

func NewCitizenHandler(service services.CitizenService, db *mongo.Collection) CitizenHandler {
	return CitizenHandler{
		service: service,
		DB:      db,
	}
}

func (s *CitizenHandler) AddCitizen(c *gin.Context) {
	rw := c.Writer
	h := c.Request

	// Provera autorizacije
	token := h.Header.Get("Authorization")
	url := "http://auth-service:8085/api/users/currentUser"
	timeout := 5 * time.Second
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

	// Dekodiranje odgovora
	var responseUser struct {
		LoggedInUser struct {
			Username string        `json:"username"`
			Email    string        `json:"email"`
			UserRole data.UserRole `json:"userRole"`
		} `json:"user"`
	}

	if err := decoder.Decode(&responseUser); err != nil {
		errorMsg := map[string]string{"error": "User object was not valid."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	//Provera uloge korisnika
	//if responseUser.LoggedInUser.UserRole != data.Citizen {
	//	errorMsg := map[string]string{"Unauthorized": " You are not a policeman."}
	//	errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
	//	return
	//}

	// Parsiranje podataka o građaninu iz zahteva
	var citizen *domain.Citizen
	if err := c.ShouldBindJSON(&citizen); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// Insertovanje građanina u bazu podataka
	insertedCitizen, _, err := s.service.InsertCitizen(citizen)
	if err != nil {
		errorMsg := map[string]string{"error": "Database problem."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	// Vraćanje odgovora sa statusom 201 Created
	rw.WriteHeader(http.StatusCreated)
	jsonResponse, err := json.Marshal(insertedCitizen)
	if err != nil {
		errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error marshaling JSON: %s", err), http.StatusInternalServerError)
		return
	}
	rw.Write(jsonResponse)
}

func (s *CitizenHandler) performAuthorizationRequestWithContext(method string, ctx context.Context, token string, url string) (*http.Response, error) {
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
func (s *CitizenHandler) GetAllCitizens(c *gin.Context) {
	rw := c.Writer
	h := c.Request

	// Provera autorizacije
	token := h.Header.Get("Authorization")
	url := "http://auth-service:8085/api/users/currentUser"
	timeout := 5 * time.Second
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

	// Dekodiranje odgovora
	var responseUser struct {
		LoggedInUser struct {
			Username string        `json:"username"`
			Email    string        `json:"email"`
			UserRole data.UserRole `json:"userRole"`
		} `json:"user"`
	}

	if err := decoder.Decode(&responseUser); err != nil {
		errorMsg := map[string]string{"error": "User object was not valid."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	// Provera uloge korisnika
	//if responseUser.LoggedInUser.UserRole != data.Policeman {
	//	errorMsg := map[string]string{"Unauthorized": " You are not a policeman."}
	//	errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
	//	return
	//}

	// Dobavljanje svih građana iz baze
	citizens, err := s.service.GetAllCitizens()
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve citizens from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	// Konverzija građana u JSON
	jsonResponse, err := json.Marshal(citizens)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (s *CitizenHandler) GetCitizenByID(c *gin.Context) {
	rw := c.Writer
	h := c.Request

	token := h.Header.Get("Authorization")
	url := "http://auth-service:8085/api/users/currentUser"

	timeout := 5 * time.Second
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
	var responseUser struct {
		LoggedInUser struct {
			UserRole data.UserRole `json:"userRole"`
		} `json:"user"`
	}
	if err := decoder.Decode(&responseUser); err != nil {
		errorMsg := map[string]string{"error": "User object was not valid."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	//if responseUser.LoggedInUser.UserRole != data.Policeman {
	//	errorMsg := map[string]string{"error": "Unauthorized. You are not a policeman."}
	//	errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
	//	return
	//}

	// Dobavljanje ID-ja građanina iz URL parametra
	jmbg := c.Param("jmbg")

	// Dobavljanje građanina po ID-ju iz baze
	citizen, err := s.service.GetCitizenByID(jmbg)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve citizen from the database. No such citizen."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	// Konvertovanje građanina u JSON
	jsonResponse, err := json.Marshal(citizen)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func ExtractTraceInfoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := otel.GetTextMapPropagator().Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
