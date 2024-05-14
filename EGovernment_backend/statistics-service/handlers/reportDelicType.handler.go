package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"net/url"
	"statistics-service/domain"
	errorMessage "statistics-service/error"
	"statistics-service/services"
	"strconv"
	"time"
)

type ReportDelicTypeHandler struct {
	service services.ReportDelicTypeService
	DB      *mongo.Collection
}

func NewReportDelicTypeHandler(service services.ReportDelicTypeService, db *mongo.Collection) ReportDelicTypeHandler {
	return ReportDelicTypeHandler{
		service: service,
		DB:      db,
	}
}

func (r *ReportDelicTypeHandler) CreateDelictsReport(c *gin.Context) {
	rw := c.Writer
	h := c.Request

	delictType := c.Param("delictType")
	yearStr := c.Param("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		errorMsg := map[string]string{"error": "Invalid year parameter"}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

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
		Title string `json:"title" binding:"required"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	totalNumber, yearFromService, err := r.GetDelictsCountAndYearPoliceService(token, domain.DelictType(delictType), year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to obtain delicts information. Try again later."})
		return
	}

	currentDateTime := primitive.NewDateTimeFromTime(time.Now())
	id := primitive.NewObjectID()
	newReport := &domain.ReportDelict{
		ID:          id,
		Type:        domain.DelictType(delictType),
		Title:       requestBody.Title,
		Date:        currentDateTime,
		TotalNumber: totalNumber,
		Year:        yearFromService,
	}

	err, _ = r.service.Create(newReport)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create report"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Report successfully saved", "report": newReport})
}

func (r *ReportDelicTypeHandler) GetAll(c *gin.Context) {
	rw := c.Writer

	reports, err := r.service.GetAll()
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve reports from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
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

func (r *ReportDelicTypeHandler) GetByID(c *gin.Context) {
	rw := c.Writer
	id := c.Param("id")
	report, err := r.service.GetById(id)
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

func (r *ReportDelicTypeHandler) GetAllByDelictType(c *gin.Context) {
	rw := c.Writer
	delictType := c.Param("delictType")

	reports, err := r.service.GetAllByDelictType(domain.DelictType(delictType))
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve reports from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

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

func (r *ReportDelicTypeHandler) GetAllByDelictTypeAndYear(c *gin.Context) {
	rw := c.Writer
	delictType := c.Param("delictType")

	yearStr := c.Param("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		errorMsg := map[string]string{"error": "Invalid year parameter"}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	reports, err := r.service.GetAllByDelictTypeAndYear(domain.DelictType(delictType), year)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve reports from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

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

func (r *ReportDelicTypeHandler) GetCurrentUserFromAuthService(token string, rw http.ResponseWriter) (*domain.User, error) {
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

func (r *ReportDelicTypeHandler) GetDelictsCountAndYearPoliceService(token string, delictType domain.DelictType, year int) (int, int, error) {
	baseURL := fmt.Sprintf("http://police-service:8084/api/delict/get/delictType/%s/year/%d", url.QueryEscape(string(delictType)), year)

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := r.performAuthorizationRequestWithContext("GET", ctx, token, baseURL)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return 0, 0, errors.New("police service is not available")
		}
		return 0, 0, errors.New("error performing request")
	}
	defer resp.Body.Close()

	var delicts []domain.Delict
	if err := json.NewDecoder(resp.Body).Decode(&delicts); err != nil {
		return 0, 0, errors.New("failed to retrieve delicts")
	}

	return len(delicts), year, nil
}

func (r *ReportDelicTypeHandler) performAuthorizationRequestWithContext(method string, ctx context.Context, token string, url string) (*http.Response, error) {
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
