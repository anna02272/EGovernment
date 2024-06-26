package handlers

import (
	"bufio"
	"bytes"
	"context"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/jung-kurt/gofpdf"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/gomail.v2"
	"io"

	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"police-service/data"
	"police-service/domain"
	errorMessage "police-service/error"
	"police-service/services"
	"police-service/storage"
	"strconv"
	"strings"
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
	logger        *log.Logger
	storage       *storage.FileStorage
}

func NewDelictHandler(service services.DelictService, db *mongo.Collection, reportService services.ReportService, l *log.Logger, s *storage.FileStorage) DelictHandler {
	return DelictHandler{
		service:       service,
		reportService: reportService,
		DB:            db,
		logger:        l,
		storage:       s,
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

	pdfFilePath, err := s.GenerateDelictPDF(delictInsertDB)
	if err != nil {
		log.Printf("Error generating PDF: %v\n", err)
		errorMsg := map[string]string{"error": "Failed to generate PDF report."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	log.Printf("Generated PDF saved at: %s", pdfFilePath)

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

		/*citizenURL := fmt.Sprintf("http://court-service:8083/api/citizen/get/%s", delictInsert.DriverJmbg)
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
		defer courtResp.Body.Close()*/
		courtURL := "http://court-service:8083/api/subject/create"
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
		defer courtResp.Body.Close()
	}

	rw.WriteHeader(http.StatusCreated)
	jsonResponse, err1 := json.Marshal(delictInsertDB)
	if err1 != nil {
		errorMessage.ReturnJSONError(rw, fmt.Sprintf("Error marshaling JSON: %s", err1), http.StatusInternalServerError)
		return
	}
	rw.Write(jsonResponse)
}

func (s *DelictHandler) GenerateTestPDF(ctx *gin.Context) {
	// Example data
	delictID := "test_delict_id"
	description := "Test Description"

	// File paths
	filePath := "/home/flower/Desktop/EUpravaProject/EGovernment/EGovernment_backend/pdf_reports"
	txtFilename := "delict_report_" + delictID + ".txt"
	txtFilePath := filepath.Join(filePath, txtFilename)
	pdfFilename := "delict_report_" + delictID + ".pdf"
	pdfFilePath := filepath.Join(filePath, pdfFilename)

	// Step 1: Generate text file
	file, err := os.Create(txtFilePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating text file: %v", err)})
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Delict Report\n"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error writing to text file: %v", err)})
		return
	}
	_, err = file.WriteString(fmt.Sprintf("Delict ID: %s\n", delictID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error writing to text file: %v", err)})
		return
	}
	_, err = file.WriteString(fmt.Sprintf("Description: %s\n", description))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error writing to text file: %v", err)})
		return
	}

	log.Println("Text file created successfully")

	// Step 2: Convert text file to PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 14)

	txtFile, err := os.Open(txtFilePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error opening text file: %v", err)})
		return
	}
	defer txtFile.Close()

	scanner := bufio.NewScanner(txtFile)
	for scanner.Scan() {
		line := scanner.Text()
		pdf.Cell(40, 10, line)
		pdf.Ln(10)
	}

	if err := scanner.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error reading text file: %v", err)})
		return
	}

	err = pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error generating PDF: %v", err)})
		return
	}

	log.Printf("PDF generated successfully: %s\n", pdfFilePath)
	ctx.JSON(http.StatusOK, gin.H{"message": "PDF generated successfully", "pdfPath": pdfFilePath})
}

func (s *DelictHandler) GenerateDelictPDF(delict *domain.Delict) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, 10, 190, 12, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 12, "Izvestaj o prekrsaju", "", 0, "C", true, 0, "")
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 14)

	pdf.SetFillColor(255, 255, 255)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, "Detalji izvestaja", "", 0, "C", false, 0, "")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, fmt.Sprintf("ID prekrsaja: %s", delict.ID.Hex()), "", 0, "", false, 0, "")
	pdf.Ln(10)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, fmt.Sprintf("ID policajca: %s", delict.PolicemanID), "", 0, "", false, 0, "")
	pdf.Ln(10)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, fmt.Sprintf("Broj vozacke dozvole: %s", delict.VehicleLicenceNumber), "", 0, "", false, 0, "")
	pdf.Ln(10)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, fmt.Sprintf("Email vozaca: %s", delict.DriverEmail), "", 0, "", false, 0, "")
	pdf.Ln(10)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, fmt.Sprintf("JMBG vozaca: %s", delict.DriverIdentificationNumber), "", 0, "", false, 0, "")
	pdf.Ln(10)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, fmt.Sprintf("Datum: %s", delict.Date.Time().Format("02.01.2006. 15:04:05")), "", 0, "", false, 0, "")
	pdf.Ln(10)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, fmt.Sprintf("Lokacija: %s", delict.Location), "", 0, "", false, 0, "")
	pdf.Ln(10)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, fmt.Sprintf("Opis: %s", delict.Description), "", 0, "", false, 0, "")
	pdf.Ln(10)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, fmt.Sprintf("Tip prekrsaja: %s", delict.DelictType), "", 0, "", false, 0, "")
	pdf.Ln(10)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, fmt.Sprintf("Status prekrsaja: %s", delict.DelictStatus), "", 0, "", false, 0, "")
	pdf.Ln(10)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, fmt.Sprintf("Novcana kazna: %.2f", delict.PriceOfFine), "", 0, "", false, 0, "")
	pdf.Ln(10)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, fmt.Sprintf("Broj kaznenih poena: %d", delict.NumberOfPenaltyPoints), "", 0, "", false, 0, "")
	pdf.Ln(10)

	pdf.SetFooterFunc(func() {
		// Footer
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 10)
		pdf.CellFormat(0, 10, "Generisano od strane eUprave", "", 0, "C", false, 0, "")
	})

	pdfDir := os.Getenv("FILE_PATH")

	pdfFilename := "delict_report_" + delict.ID.Hex() + ".pdf"
	pdfFilePath := filepath.Join(pdfDir, pdfFilename)

	err := pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		log.Println("Error generating PDF:", err)
		return "", err
	}

	log.Printf("Generated PDF saved at: %s", pdfFilePath)
	return pdfFilePath, nil
}

func (h *DelictHandler) ServeDelictPDF(c *gin.Context) {
	delictID := c.Param("id")

	pdfDir := os.Getenv("FILE_PATH")

	pdfFilename := fmt.Sprintf("delict_report_%s.pdf", delictID)
	pdfFilePath := filepath.Join(pdfDir, pdfFilename)

	pdfFile, err := os.Open(pdfFilePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	defer pdfFile.Close()

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename="+pdfFilename)

	http.ServeFile(c.Writer, c.Request, pdfFilePath)
}

func isValidDelictType(delictType domain.DelictType) bool {
	switch delictType {
	case domain.Speeding, domain.DrivingUnderTheInfluenceOfAlcohol, domain.DrivingUnderTheInfluence, domain.ImproperOvertaking, domain.ImproperParking, domain.FailureTooComplyWithTrafficLightsAndSigns, domain.ImproperUseOfSeatBeltsAndChildSeats, domain.UsingMobilePhoneWhileDriving, domain.ImproperUseOfMotorVehicle, domain.Other:
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

func (s *DelictHandler) GetImageURLS(c *gin.Context) {
	rw := c.Writer
	//h := c.Request
	/*token := h.Header.Get("Authorization")
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
	}*/

	folderName := c.Param("folderName")
	imageURLs, err := s.storage.GetImageURLS(folderName)
	if err != nil {
		errorMessage.ReturnJSONError(rw, map[string]string{"error": "Error getting image URLs"}, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rw).Encode(imageURLs)
}

func (s *DelictHandler) GetImageContent(c *gin.Context) {
	rw := c.Writer
	//h := c.Request
	/*token := h.Header.Get("Authorization")
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
	}*/

	folderName := c.Param("folderName")
	imageName := c.Param("imageName")
	imagePath := path.Join(folderName, imageName)
	imagePath = strings.TrimPrefix(imagePath, "/")

	imageType := mime.TypeByExtension(filepath.Ext(imagePath))
	if imageType == "" {
		errorMessage.ReturnJSONError(rw, map[string]string{"error": "Error retrieving image type"}, http.StatusInternalServerError)
		return
	}

	imageContent, err := s.storage.GetImageContent(imagePath)
	if err != nil {
		errorMessage.ReturnJSONError(rw, map[string]string{"error": "Error retrieving image content"}, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", imageType)
	rw.WriteHeader(http.StatusOK)
	rw.Write(imageContent)
}

func (s *DelictHandler) UploadImages(c *gin.Context) {
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

	folderName := c.Param("folderName")

	err = h.ParseMultipartForm(40 << 20)
	if err != nil {
		errorMessage.ReturnJSONError(rw, map[string]string{"error": "Error parsing form"}, http.StatusBadRequest)
		return
	}

	files := h.MultipartForm.File["images"]

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			errorMessage.ReturnJSONError(rw, map[string]string{"error": "Error opening file"}, http.StatusInternalServerError)
			return
		}
		defer src.Close()

		imageContent, err := io.ReadAll(src)
		if err != nil {
			errorMessage.ReturnJSONError(rw, map[string]string{"error": "Error reading file"}, http.StatusInternalServerError)
			return
		}

		err = s.storage.SaveImage(folderName, file.Filename, imageContent)
		if err != nil {
			errorMessage.ReturnJSONError(rw, map[string]string{"error": "Error saving file"}, http.StatusInternalServerError)
			return
		}
		log.Printf("[UploadImages] File %s successfully saved in folder %s", file.Filename, folderName)
	}

	rw.WriteHeader(http.StatusOK)
}
