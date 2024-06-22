package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/gomail.v2"
	"log"
	"net/http"
	"police-service/data"
	"police-service/domain"
	errorMessage "police-service/error"
	"police-service/services"
	"strconv"
	"time"
)

var (
	smtpServer     = "smtp.office365.com"
	smtpServerPort = 587
	smtpEmail      = "EGovernmentPolice@outlook.com"
	smtpPassword   = "amhrxqinoamvtcss"
)

type DelictHandler struct {
	service       services.DelictService
	reportService services.ReportService
	DB            *mongo.Collection
}

func NewDelictHandler(service services.DelictService, db *mongo.Collection, reportService services.ReportService) DelictHandler {
	return DelictHandler{
		service:       service,
		reportService: reportService,
		DB:            db,
	}
}
func (s *DelictHandler) CreateDelict(c *gin.Context) {
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
	policemanID := responseUser.LoggedInUser.ID.Hex()
	delict, exists := c.Get("delict")
	if !exists {
		errorMsg := map[string]string{"Error": " delict object was not valid."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}
	delictInsert, ok := delict.(domain.DelictCreate)
	if !ok {
		errorMsg := map[string]string{"error": "Invalid type for delict."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}
	if !isValidDelictType(delictInsert.DelictType) {
		errorMessage.ReturnJSONError(rw, map[string]string{"error": "Invalid delict type."}, http.StatusBadRequest)
		return
	}
	if !isValidDelictStatus(delictInsert.DelictStatus) {
		errorMessage.ReturnJSONError(rw, map[string]string{"error": "Invalid delict status."}, http.StatusBadRequest)
		return
	}
	delictInsertDB, _, err := s.service.InsertDelict(&delictInsert, policemanID)
	if err != nil {
		errorMsg := map[string]string{"error": "Database problem."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	/*err = s.sendDelictMail(delictInsertDB.Description, delictInsertDB.DriverEmail)
	if err != nil {
		errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error sending email: %s", err), http.StatusInternalServerError)
		return
	}*/

	if delictInsert.NumberOfPenaltyPoints > 0 {
		driverID := delictInsert.DriverIdentificationNumber
		points := delictInsert.NumberOfPenaltyPoints
		updatePointsURL := fmt.Sprintf("http://vehicles-service:8080/api/driver/updatePenaltyPoints/%s", driverID)

		updatePointsReqBody := struct {
			Points int64 `json:"points"`
		}{
			Points: points,
		}

		updatePointsReqBodyJSON, err := json.Marshal(updatePointsReqBody)
		if err != nil {
			errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error marshaling JSON: %s", err), http.StatusInternalServerError)
			return
		}

		updatePointsReq, err := http.NewRequest("PATCH", updatePointsURL, bytes.NewBuffer(updatePointsReqBodyJSON))
		if err != nil {
			errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error creating update points request: %s", err), http.StatusInternalServerError)
			return
		}
		updatePointsReq.Header.Set("Content-Type", "application/json")
		updatePointsReq.Header.Set("Authorization", token)

		updatePointsResp, err := http.DefaultClient.Do(updatePointsReq)
		if err != nil || updatePointsResp.StatusCode != http.StatusOK {
			errorMessage.ReturnJSONError(rw, "Failed to update penalty points in vehicle driver service.", http.StatusInternalServerError)
			return
		}
		defer updatePointsResp.Body.Close()
	}

	if delictInsert.DelictStatus == domain.SentToCourt {

		citizenURL := fmt.Sprintf("http://court-service:8083/api/citizen/get/%s", delictInsert.DriverJmbg)
		citizenReq, err := http.NewRequest("GET", citizenURL, nil)
		if err != nil {
			errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error creating citizen request: %s", err), http.StatusInternalServerError)
			return
		}
		citizenReq.Header.Set("Authorization", token)
		citizenResp, err := http.DefaultClient.Do(citizenReq)
		if err != nil || citizenResp.StatusCode != http.StatusOK {
			errorMessage.ReturnJSONError(rw, "Failed to get citizen information.", http.StatusInternalServerError)
			return
		}
		defer citizenResp.Body.Close()

		var citizen domain.Citizen
		if err := json.NewDecoder(citizenResp.Body).Decode(&citizen); err != nil {
			errorMessage.ReturnJSONError(rw, "Failed to decode citizen information.", http.StatusInternalServerError)
			return
		}

		log.Printf("Fetched Citizen: %+v\n", citizen)

		courtURL := "http://court-service:8083/api/subject/create"
		subject := struct {
			ViolationID string         `json:"violation_id"`
			Accused     domain.Citizen `json:"accused"`
		}{
			ViolationID: delictInsertDB.ID.Hex(),
			Accused:     citizen,
		}
		log.Printf("Subject to be sent: %+v\n", subject)
		subjectJSON, err := json.Marshal(subject)
		if err != nil {
			errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error marshaling JSON: %s", err), http.StatusInternalServerError)
			return
		}
		log.Printf("Subject JSON: %s\n", string(subjectJSON))

		courtReq, err := http.NewRequest("POST", courtURL, bytes.NewBuffer(subjectJSON))
		if err != nil {
			errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error creating court request: %s", err), http.StatusInternalServerError)
			return
		}
		courtReq.Header.Set("Content-Type", "application/json")
		courtReq.Header.Set("Authorization", token)
		courtResp, err := http.DefaultClient.Do(courtReq)
		if err != nil || courtResp.StatusCode != http.StatusCreated {
			errorMessage.ReturnJSONError(rw, "Failed to create subject in court service.", http.StatusInternalServerError)
			return
		}
		defer courtResp.Body.Close()
		/*courtURL := "http://court-service:8083/api/subject/create"
		subject := struct {
			ViolationID string `json:"violation_id"`
		}{
			ViolationID: delictInsertDB.ID.Hex(),
		}
		subjectJSON, err := json.Marshal(subject)
		if err != nil {
			errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error marshaling JSON: %s", err), http.StatusInternalServerError)
			return
		}
		courtReq, err := http.NewRequest("POST", courtURL, bytes.NewBuffer(subjectJSON))
		if err != nil {
			errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error creating court request: %s", err), http.StatusInternalServerError)
			return
		}
		courtReq.Header.Set("Content-Type", "application/json")
		courtReq.Header.Set("Authorization", token)
		courtResp, err := http.DefaultClient.Do(courtReq)
		if err != nil || courtResp.StatusCode != http.StatusCreated {
			errorMessage.ReturnJSONError(rw, "Failed to create subject in court service.", http.StatusInternalServerError)
			return
		}
		defer courtResp.Body.Close()*/
	}

	rw.WriteHeader(http.StatusCreated)
	jsonResponse, err1 := json.Marshal(delictInsertDB)
	if err1 != nil {
		errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error marshaling JSON: %s", err1), http.StatusInternalServerError)
		return
	}
	rw.Write(jsonResponse)
}

func isValidDelictType(delictType domain.DelictType) bool {
	switch delictType {
	case domain.Speeding, domain.DrivingUnderTheInfluenceOfAlcohol, domain.DrivingUnderTheInfluence, domain.ImproperOvertaking, domain.ImproperParking, domain.FailureTooComplyWithTrafficLightsAndSigns, domain.ImproperUseOfSeatBeltsAndChildSeats, domain.UsingMobilePhoneWhileDriving, domain.ImproperUseOfMotorVehicle:
		return true
	default:
		return false
	}
}

func isValidDelictStatus(delictStatus domain.DelictStatus) bool {
	switch delictStatus {
	case domain.FineAwarded, domain.FinePaid, domain.SentToCourt:
		return true
	default:
		return false
	}
}

func (s *DelictHandler) PayFine(c *gin.Context) {
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
			Email    string        `json:"email"`
			UserRole data.UserRole `json:"userRole"`
		} `json:"user"`
	}
	if err := decoder.Decode(&responseUser); err != nil {
		errorMsg := map[string]string{"error": "User object was not valid."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	if responseUser.LoggedInUser.UserRole != data.Citizen {
		errorMsg := map[string]string{"Unauthorized": "You are not a citizen."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	delictId := c.Param("id")
	delict, err := s.service.GetDelictById(delictId, ctx)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve delict from the database. No such delict."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	if delict.DelictStatus != domain.FineAwarded {
		errorMsg := map[string]string{"error": "Delict is not in a payable status."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	delict.DelictStatus = domain.FinePaid
	if err := s.service.UpdateDelictStatus(delict); err != nil {
		errorMsg := map[string]string{"error": "Failed to update delict status."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	if err := s.sendPaymentConfirmationEmail(delict.Description, delict.DriverEmail); err != nil {
		errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error sending email: %s", err), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"message": "Fine paid successfully."}`))
}

func (s *DelictHandler) sendPaymentConfirmationEmail(Description, driverEmail string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", smtpEmail)
	m.SetHeader("To", driverEmail)
	m.SetHeader("Subject", "Potvrda isplate prekrsaja")

	bodyString := fmt.Sprintf("Vasa novcana kazna je uspesno isplacena.\n Opis isplacenog prekrsaja:\n %s ", Description)
	m.SetBody("text", bodyString)

	client := gomail.NewDialer(smtpServer, smtpServerPort, smtpEmail, smtpPassword)

	if err := client.DialAndSend(m); err != nil {
		log.Fatalf("Failed to send mail because of: %s", err)
		return err
	}
	return nil
}

func (s *DelictHandler) sendDelictMail(Description, email string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", smtpEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "EUprava obavestenje")

	bodyString := fmt.Sprintf("Za Vas je kreiran prekrsaj sa opisom:\n %s \nStanje vaseg prekrsaja mozete pratiti na portalu EUprave https://localhost:4200/", Description)
	m.SetBody("text", bodyString)

	client := gomail.NewDialer(smtpServer, smtpServerPort, smtpEmail, smtpPassword)

	if err := client.DialAndSend(m); err != nil {
		log.Fatalf("Failed to send mail because of: %s", err)
		return err
	}

	return nil
}

func (s *DelictHandler) GetAllDelicts(c *gin.Context) {
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
	delicts, err := s.service.GetAllDelicts()
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve delicts from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(delicts)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}
func (s *DelictHandler) GetAllDelictsByDelictType(c *gin.Context) {
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
	delicts, err := s.service.GetAllDelictsByDelictType(domain.DelictType(delictType))
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve delicts from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	// Convert delicts to JSON
	jsonResponse, err := json.Marshal(delicts)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (s *DelictHandler) GetDelictsByPolicemanID(c *gin.Context) {
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
			ID       primitive.ObjectID `json:"id"`
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

	policemanID := responseUser.LoggedInUser.ID.Hex()
	delicts, err := s.service.GetAllDelictsByPolicemanID(policemanID)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve delicts from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(delicts)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (s *DelictHandler) GetDelictsByDriver(c *gin.Context) {
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
			Email    string        `json:"email"`
			UserRole data.UserRole `json:"userRole"`
		} `json:"user"`
	}
	if err := decoder.Decode(&responseUser); err != nil {
		errorMsg := map[string]string{"error": "User object was not valid."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	if responseUser.LoggedInUser.UserRole != data.Citizen {
		errorMsg := map[string]string{"Unauthorized": " You are not citizen."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	driverEmail := responseUser.LoggedInUser.Email
	log.Printf("driverEmail: %+v\n", driverEmail)
	delicts, err := s.service.GetAllDelictsByDriver(driverEmail)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve delicts from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(delicts)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (s *DelictHandler) CheckDriverAlcoholDelicts(c *gin.Context) {
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

	driverID := c.Param("driverId")

	delicts, err := s.service.GetAllDelictsForDriverByDelictType(driverID)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve delicts from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	// Convert delicts to JSON
	jsonResponse, err := json.Marshal(delicts)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (s *DelictHandler) performAuthorizationRequestWithContext(method string, ctx context.Context, token string, url string) (*http.Response, error) {
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
func (s *DelictHandler) GetDelictByID(c *gin.Context) {
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

	// Check if the user role is TrafficPoliceman, Judge, or Employee
	allowedRoles := map[data.UserRole]bool{
		data.TrafficPoliceman: true,
		data.Judge:            true,
		data.Employee:         true,
		data.Policeman:        true,
		data.Citizen:          true,
	}

	if !allowedRoles[responseUser.LoggedInUser.UserRole] {
		errorMsg := map[string]string{"error": "Unauthorized. You do not have access to this resource."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}
	delictId := c.Param("id")
	delict, err := s.service.GetDelictById(delictId, ctx)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve delict from the database.No such delict."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(delict)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (s *DelictHandler) GetAllDelictsByDelictTypeAndYear(c *gin.Context) {
	rw := c.Writer
	h := c.Request
	token := h.Header.Get("Authorization")
	url := "http://auth-service:8085/api/users/currentUser"
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := s.performAuthorizationRequestWithContext("GET", ctx, token, url)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
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
	yearStr := c.Param("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		errorMsg := map[string]string{"error": "Invalid year parameter"}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	delicts, err := s.service.GetAllDelictsByDelictTypeAndYear(domain.DelictType(delictType), year)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve delicts from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	// Convert delicts to JSON
	jsonResponse, err := json.Marshal(delicts)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}
