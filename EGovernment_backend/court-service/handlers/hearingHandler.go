package handlers

import (
	"context"
	"court-service/data"
	"court-service/domain"
	errorMessage "court-service/error"
	"court-service/services"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/smtp"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type HearingHandler struct {
	service services.HearingService
	serv    services.SubjectService

	DB *mongo.Collection
}

func NewHearingHandler(service services.HearingService, db *mongo.Collection, serv services.SubjectService) HearingHandler {
	return HearingHandler{
		service: service,
		serv:    serv,
		DB:      db,
	}
}

//	func (hh *HearingHandler) CreateHearing(c *gin.Context) {
//		var hearing domain.Hearing
//
//		if err := c.ShouldBindJSON(&hearing); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
//			return
//		}
//
//		createdHearing, err := hh.service.CreateHearing(&hearing)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hearing"})
//			return
//		}
//
//		c.JSON(http.StatusCreated, createdHearing)
//	}
func (hh *HearingHandler) CreateHearing(c *gin.Context) {
	rw := c.Writer
	req := c.Request

	// Dobavljanje korisnika i njegovih podataka
	token := req.Header.Get("Authorization")
	url := "http://auth-service:8085/api/users/currentUser"

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := hh.performAuthorizationRequestWithContext("GET", ctx, token, url)
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
			ID       primitive.ObjectID `json:"id"`
			Username string             `json:"username"`
			Email    string             `json:"email"`
			UserRole data.UserRole      `json:"userRole"`
		} `json:"user"`
	}

	if err := decoder.Decode(&responseUser); err != nil {
		errorMsg := map[string]string{"error": "User object was not valid."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	userID := responseUser.LoggedInUser.ID
	userEmail := responseUser.LoggedInUser.Email

	var hearingInsert domain.Hearing
	if err := c.ShouldBindJSON(&hearingInsert); err != nil {
		errorMsg := map[string]string{"error": "Invalid request payload"}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	hearingInsert.JudgeID = userID

	if hearingInsert.SubjectID.IsZero() {
		errorMsg := map[string]string{"error": "Subject ID is required"}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	dateParsed, err := time.Parse("2006-01-02", hearingInsert.Date)
	if err != nil {
		errorMsg := map[string]string{"error": "Neispravan format datuma, koristite format YYYY-MM-DD"}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	hearingInsert.Date = dateParsed.Format("2006-01-02")

	// Dobavljanje subjekta na osnovu subject_id koristeći SubjectService
	//subjectID := hearingInsert.SubjectID
	//subject, err := hh.serv.GetSubjectByID(subjectID)
	//if err != nil {
	//	errorMsg := map[string]string{"error": "Failed to retrieve subject"}
	//	errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
	//	return
	//}

	// Dobavljanje emaila vozača iz prekršaja
	violationID := "66895db22f5753983f9e9f75"
	offense, err := hh.GetDelict(c, violationID)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve offense"}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	driverEmail := offense.DriverEmail

	hearingInsertDB, err := hh.service.CreateHearing(&hearingInsert)
	if err != nil {
		errorMsg := map[string]string{"error": "Database problem."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	if err := hh.sendEmailNotification(driverEmail, hearingInsertDB); err != nil {
		errorMessage.ReturnJSONError(rw, fmt.Sprintf("Failed to send email: %s", err), http.StatusInternalServerError)
		return
	}

	if err := hh.sendEmailNotification(userEmail, hearingInsertDB); err != nil {
		errorMessage.ReturnJSONError(rw, fmt.Sprintf("Failed to send email: %s", err), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	jsonResponse, err1 := json.Marshal(hearingInsertDB)
	if err1 != nil {
		errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error marshaling JSON: %s", err1), http.StatusInternalServerError)
		return
	}
	rw.Write(jsonResponse)
}

func (hh *HearingHandler) GetDelict(c *gin.Context, id string) (domain.Delict, error) {
	token := c.Request.Header.Get("Authorization")
	delict, err := hh.GetDelictPoliceService(token, id)
	if err != nil {
		return domain.Delict{}, fmt.Errorf("failed to obtain delict information: %v", err)
	}
	return delict, nil
}
func (hh *HearingHandler) GetHearingByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid ID format"})
		return
	}

	hearing, err := hh.service.GetHearingByID(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to get hearing"})
		return
	}

	if hearing == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Hearing not found"})
		return
	}

	c.JSON(http.StatusOK, hearing)
}

func (h *HearingHandler) performAuthorizationRequestWithContext(method string, ctx context.Context, token string, url string) (*http.Response, error) {
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
func (hh *HearingHandler) GetDelictPoliceService(token string, id string) (domain.Delict, error) {
	baseURL := fmt.Sprintf("http://police-service:8084/api/delict/get/%s", url.QueryEscape(id))

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := hh.performAuthorizationRequestWithContext("GET", ctx, token, baseURL)
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

func (hh *HearingHandler) performAuthorizationRequestWithContextt(method string, ctx context.Context, token string, url string) (*http.Response, error) {
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
func (hh *HearingHandler) sendEmailNotification(email string, hearing *domain.Hearing) error {
	// Konfiguracija SMTP kredencijala za slanje emaila preko Mailtrap servisa
	smtpHost := "smtp.mailtrap.io"
	smtpPort := "465"
	smtpUsername := "ae3a005d56dc9c" // Promeniti sa svojim Mailtrap korisničkim imenom
	smtpPassword := "90d61363fd1a08" // Promeniti sa svojim Mailtrap lozinkom

	// Formatiranje datuma ročišta
	hearingDate, err := time.Parse("2006-01-02", hearing.Date)
	if err != nil {
		return fmt.Errorf("failed to parse hearing date: %v", err)
	}
	formattedDate := hearingDate.Format("02.01.2006") // Primer formata: 15.06.2024

	// Generisanje linka za potvrdu prisustva
	confirmationLink := fmt.Sprintf("http://yourdomain.com/confirm?hearing=%s", hearing.ID.Hex())

	// Formiranje poruke emaila
	subject := "Obaveštenje o zakazanom ročištu"
	body := fmt.Sprintf("Poštovani,\n\nZakazano je ročište za vaš prekršaj.\nDatum ročišta: %s\n\nMolimo vas da potvrdite svoje prisustvo klikom na sledeći link: %s\n\nS poštovanjem,\nSudski sistem", formattedDate, confirmationLink)

	message := "To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUsername, []string{email}, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
func (hh *HearingHandler) GetJudgeHearings(c *gin.Context) {
	rw := c.Writer
	req := c.Request

	// Dobavljanje korisnika i njegovih podataka
	token := req.Header.Get("Authorization")
	url := "http://auth-service:8085/api/users/currentUser"

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := hh.performAuthorizationRequestWithContext("GET", ctx, token, url)
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
			ID       primitive.ObjectID `json:"id"`
			Username string             `json:"username"`
			Email    string             `json:"email"`
			UserRole data.UserRole      `json:"userRole"`
		} `json:"user"`
	}

	if err := decoder.Decode(&responseUser); err != nil {
		errorMsg := map[string]string{"error": "User object was not valid."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	judgeID := responseUser.LoggedInUser.ID

	// Dohvatanje ročišta za sudiju na osnovu njegovog ID-a
	hearings, err := hh.service.GetJudgeHearings(judgeID)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve judge hearings"}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, hearings)
}
