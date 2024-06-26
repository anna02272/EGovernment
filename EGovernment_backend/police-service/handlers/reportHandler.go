package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"police-service/data"
	"police-service/domain"
	errorMessage "police-service/error"
	"police-service/services"
	"time"
)

type ReportHandler struct {
	service services.ReportService
	DB      *mongo.Collection
}

func NewReportHandler(service services.ReportService, db *mongo.Collection) ReportHandler {
	return ReportHandler{
		service: service,
		DB:      db,
	}

}

func (s *ReportHandler) CreateReport(c *gin.Context) {
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
			ID       primitive.ObjectID `json:"id"`
			username string             `json:"username"`
			email    string             `json:"email"`
			UserRole data.UserRole      `json:"userRole"`
		} `json:"user"`
	}

	if err := decoder.Decode(&responseUser); err != nil {
		errorMsg := map[string]string{"error": "User object was not valid."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	if responseUser.LoggedInUser.UserRole != data.TrafficPoliceman {
		errorMsg := map[string]string{"Unauthorized": " You are not traffic policeman."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	report, exists := c.Get("report")
	if !exists {
		errorMsg := map[string]string{"Error": " report object was not valid."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	reportInsert, ok := report.(domain.ReportCreate)
	if !ok {
		errorMsg := map[string]string{"error": "Invalid type for delict."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	if !isValidDelictType(reportInsert.DelictType) {
		errorMessage.ReturnJSONError(rw, map[string]string{"error": "Invalid delict type."}, http.StatusBadRequest)
		return
	}
	log.Println("Inserting report...")
	reportInsertDB, _, err := s.service.InsertReport(&reportInsert, responseUser.LoggedInUser.ID.Hex(), "66409febc7cbf41296293b92", "66409febc7cbf41296293b92")
	if err != nil {
		errorMsg := map[string]string{"error": "Database problem."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	log.Println("Report inserted successfully.")

	rw.WriteHeader(http.StatusCreated)
	jsonResponse, err1 := json.Marshal(reportInsertDB)
	if err1 != nil {
		errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error marshaling JSON: %s", err1), http.StatusInternalServerError)
		return
	}
	rw.Write(jsonResponse)

}

func (s *ReportHandler) GetAllReports(c *gin.Context) {
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

	reports, err := s.service.GetAllReport()
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve reports from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	// Convert vehicles to JSON
	jsonResponse, err := json.Marshal(reports)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (s *ReportHandler) GetAllReportssByDelictType(c *gin.Context) {
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

	delictType := c.Param("delictType")
	reports, err := s.service.GetAllReportsByDelictType(domain.DelictType(delictType))
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve reports from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	// Convert reports to JSON
	jsonResponse, err := json.Marshal(reports)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (s *ReportHandler) performAuthorizationRequestWithContext(method string, ctx context.Context, token string, url string) (*http.Response, error) {
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

func (s *ReportHandler) GetReportByID(c *gin.Context) {
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

	if responseUser.LoggedInUser.UserRole != data.TrafficPoliceman {
		errorMsg := map[string]string{"error": "Unauthorized. You are not a traffic policeman."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	reportId := c.Param("id")

	report, err := s.service.GetReportById(reportId, ctx)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve report from the database.No such report."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(report)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}
