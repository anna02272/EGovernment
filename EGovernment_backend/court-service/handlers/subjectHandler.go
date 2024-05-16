package handlers

import (
	"context"
	"court-service/domain"
	"court-service/services"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"net/url"
	"time"
)

type SubjectHandler struct {
	service services.SubjectService
	DB      *mongo.Collection
}

func NewSubjectHandler(service services.SubjectService, db *mongo.Collection) SubjectHandler {
	return SubjectHandler{
		service: service,
		DB:      db,
	}
}

func (sh *SubjectHandler) CreateSubject(c *gin.Context) {
	var subject *domain.Subject

	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	createdSubject, err := sh.service.CreateSubject(subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subject"})
		return
	}

	c.JSON(http.StatusCreated, createdSubject)
}
func (sh *SubjectHandler) GetDelict(c *gin.Context) {
	h := c.Request
	token := h.Header.Get("Authorization")
	id := c.Param("id")
	delict, err := sh.GetDelictPoliceService(token, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to obtain delicts information. Try again later."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delict successfully gotten", "delict": delict})

}

func (sh *SubjectHandler) GetDelictPoliceService(token string, id string) (domain.Delict, error) {
	baseURL := fmt.Sprintf("http://police-service:8084/api/delict/get/%s", url.QueryEscape(id))

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := sh.performAuthorizationRequestWithContext("GET", ctx, token, baseURL)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return domain.Delict{}, errors.New("police service request timed out")
		}
		return domain.Delict{}, fmt.Errorf("error performing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Delict{}, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	var delict domain.Delict
	if err := json.NewDecoder(resp.Body).Decode(&delict); err != nil {
		return domain.Delict{}, fmt.Errorf("failed to decode response body: %v", err)
	}

	return delict, nil
}

func (sh *SubjectHandler) performAuthorizationRequestWithContext(method string, ctx context.Context, token string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}

	return resp, nil
}

//delictID := c.Param("id")
//log.Print(delictID)
//// Get the token from the request header
//token := c.GetHeader("Authorization")
//if token == "" {
//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
//	return
//}
//
//// Make HTTP GET request to the other service
//client := &http.Client{}
//req, err := http.NewRequest("GET", "http://police-service:8084/api/delict/get/"+delictID, nil)
//if err != nil {
//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
//	return
//}
//
//// Set the token for authentication in the request header
//req.Header.Set("Authorization", token)
//
//resp, err := client.Do(req)
//if err != nil {
//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get delict"})
//	return
//}
//defer resp.Body.Close()
//
//// Check the response status
//if resp.StatusCode != http.StatusOK {
//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get delict"})
//	return
//}

// Decode the response into Delict struct
//var delict domain.Delict
//if err := json.NewDecoder(resp.Body).Decode(&delict); err != nil {
//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse delict response"})
//	return
//}

//c.JSON(http.StatusOK, delict)
