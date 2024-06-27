package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"vehicles-service/data"
	"vehicles-service/domain"
	errorMessage "vehicles-service/error"
	"vehicles-service/services"
)

type VehicleHandler struct {
	service       services.VehicleService
	driverService services.VehicleDriverService
	DB            *mongo.Collection
}

func NewVehicleHandler(service services.VehicleService, db *mongo.Collection, driverService services.VehicleDriverService) VehicleHandler {
	return VehicleHandler{
		service:       service,
		DB:            db,
		driverService: driverService,
	}

}

func (s *VehicleHandler) performAuthorizationRequestWithContext(method string, ctx context.Context, token string, url string) (*http.Response, error) {
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

func (s *VehicleHandler) GenerateVehiclesReportPDF() (string, error) {
	vehicles, err := s.service.GetAllRegisteredVehicles()
	if err != nil {
		log.Println("Error retrieving registered vehicles:", err)
		return "", err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, 10, 190, 12, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 12, "Izvestaj o registrovanim vozilima", "", 0, "C", true, 0, "")
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 14)

	pdf.SetFillColor(255, 255, 255)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, "Detalji registrovanih vozila", "", 0, "C", false, 0, "")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)

	for _, vehicle := range vehicles {
		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("ID: %s", vehicle.ID.Hex()), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Registraciona tablica: %s", vehicle.RegistrationPlate), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Model vozila: %s", vehicle.VehicleModel), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Vlasnik vozila: %s", vehicle.VehicleOwner), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Datum registracije: %s", vehicle.RegistrationDate.Format("02.01.2006")), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Kategorija: %s", vehicle.Category), "", 0, "", false, 0, "")
		pdf.Ln(10)
	}

	pdf.SetFooterFunc(func() {
		// Footer
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 10)
		pdf.CellFormat(0, 10, "Generisano od strane eUprave", "", 0, "C", false, 0, "")
	})

	pdfDir := os.Getenv("FILE_PATH")

	pdfFilename := "registered_vehicles_report.pdf"
	pdfFilePath := filepath.Join(pdfDir, pdfFilename)

	err = pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		log.Println("Error generating PDF:", err)
		return "", err
	}

	log.Printf("Generated PDF saved at: %s", pdfFilePath)
	return pdfFilePath, nil
}
func (h *VehicleHandler) ServeVehiclesPDF(c *gin.Context) {
	pdfDir := os.Getenv("FILE_PATH")

	pdfFilename := "registered_vehicles_report.pdf"
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

func (s *VehicleHandler) CreateVehicle(c *gin.Context) {
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
		errorMsg := map[string]string{"error": "Unauthorized. You are not policeman"}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	vehicle, exists := c.Get("vehicle")
	if !exists {
		errorMsg := map[string]string{"error": "vehicle object was not valid"}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	vehicleInsert, ok := vehicle.(domain.VehicleCreate)

	registrationPlate := vehicleInsert.RegistrationPlate

	existingVehicle, err := s.service.GetVehicleByID(registrationPlate, ctx)

	if existingVehicle != nil {
		errorMsg := map[string]string{"error": "Vehicle with this registration plate already exists."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusConflict)
		return
	}

	vehicleDriverId := vehicleInsert.VehicleOwner
	vehicleDriver, _ := s.driverService.GetVehicleDriverByID(vehicleDriverId, ctx)

	if vehicleDriver == nil {
		errorMsg := map[string]string{"error": "There's no driver with that ID in database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	if !isValidRegistrationPlate(registrationPlate) {
		errorMsg := map[string]string{"error": "Invalid registration plate format."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	if !ok {
		errorMsg := map[string]string{"error": "Invalid type for vehicle."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	vehicleDriverInsertDB, _, err := s.service.InsertVehicle(&vehicleInsert)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR HERE")
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

func (s *VehicleHandler) GetAllVehicles(c *gin.Context) {
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

	allowedRoles := map[data.UserRole]bool{
		data.TrafficPoliceman: true,
		data.Policeman:        true,
		data.Employee:         false,
		data.Judge:            false,
	}

	if !allowedRoles[responseUser.LoggedInUser.UserRole] {
		errorMsg := map[string]string{"error": "Unauthorized. You do not have access to this resource."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	vehicles, err := s.service.GetAllVehicles()
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve vehicles from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(vehicles)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (s *VehicleHandler) GenerateAndServeVehiclesByCategoryReportPDF(c *gin.Context) {
	categoryParam := c.Param("category")
	if categoryParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category parameter is required"})
		return
	}

	category := domain.Category(categoryParam)

	vehicles, err := s.service.GetAllRegisteredVehiclesByCategory(category)
	fmt.Println(vehicles)
	fmt.Println("ALL REGISTERED VEHICLES BY CATEGORY")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving registered vehicles for category"})
		log.Println("Error retrieving registered vehicles for category:", err)
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, 10, 190, 12, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 12, "Izvestaj o registrovanim vozilima za kategoriju "+string(category), "", 0, "C", true, 0, "")
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 14)
	pdf.SetFillColor(255, 255, 255)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, "Detalji registrovanih vozila", "", 0, "C", false, 0, "")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)

	for _, vehicle := range vehicles {
		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("ID: %s", vehicle.ID.Hex()), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Registraciona tablica: %s", vehicle.RegistrationPlate), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Model vozila: %s", vehicle.VehicleModel), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Vlasnik vozila: %s", vehicle.VehicleOwner), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Datum registracije: %s", vehicle.RegistrationDate.Format("02.01.2006")), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Kategorija: %s", vehicle.Category), "", 0, "", false, 0, "")
		pdf.Ln(20) // Add extra space here between each vehicle's details
	}

	pdf.SetFooterFunc(func() {
		// Footer
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 10)
		pdf.CellFormat(0, 10, "Generisano od strane eUprave", "", 0, "C", false, 0, "")
	})

	// Serve PDF as downloadable file
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=registered_vehicles_report_"+string(category)+".pdf")
	c.Writer.Header().Set("Content-Type", "application/pdf")

	err = pdf.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating PDF"})
		log.Println("Error generating PDF:", err)
		return
	}

	log.Println("Generated PDF served successfully")
}

func (s *VehicleHandler) GenerateAndServeVehiclesReportPDF(c *gin.Context) {
	// Generate PDF
	vehicles, err := s.service.GetAllRegisteredVehicles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving registered vehicles"})
		log.Println("Error retrieving registered vehicles:", err)
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)

	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, 10, 190, 12, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 12, "Izvestaj o registrovanim vozilima", "", 0, "C", true, 0, "")
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 14)

	pdf.SetFillColor(255, 255, 255)
	pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 8, "Detalji registrovanih vozila", "", 0, "C", false, 0, "")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)

	for _, vehicle := range vehicles {
		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("ID: %s", vehicle.ID.Hex()), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Registraciona tablica: %s", vehicle.RegistrationPlate), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Model vozila: %s", vehicle.VehicleModel), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Vlasnik vozila: %s", vehicle.VehicleOwner), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Datum registracije: %s", vehicle.RegistrationDate.Format("02.01.2006")), "", 0, "", false, 0, "")
		pdf.Ln(10)

		pdf.SetFillColor(240, 240, 240)
		pdf.Rect(10, pdf.GetY()+2, 190, 8, "F")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 8, fmt.Sprintf("Kategorija: %s", vehicle.Category), "", 0, "", false, 0, "")
		pdf.Ln(20) // Add extra space here between each vehicle's details
	}

	pdf.SetFooterFunc(func() {
		// Footer
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 10)
		pdf.CellFormat(0, 10, "Generisano od strane eUprave", "", 0, "C", false, 0, "")
	})

	// Serve PDF as downloadable file
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=registered_vehicles_report.pdf")
	c.Writer.Header().Set("Content-Type", "application/pdf")

	err = pdf.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating PDF"})
		log.Println("Error generating PDF:", err)
		return
	}

	log.Println("Generated PDF served successfully")
}
func (s *VehicleHandler) GetAllRegisteredVehicles(c *gin.Context) {
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

	allowedRoles := map[data.UserRole]bool{
		data.TrafficPoliceman: true,
		data.Policeman:        true,
		data.Employee:         false,
		data.Judge:            false,
	}

	if !allowedRoles[responseUser.LoggedInUser.UserRole] {
		errorMsg := map[string]string{"error": "Unauthorized. You do not have access to this resource."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	vehicles, err := s.service.GetAllRegisteredVehicles()
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve registered vehicles from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	if len(vehicles) == 0 {
		errorMsg := map[string]string{"message": "No registered vehicles found."}
		jsonResponse, err := json.Marshal(errorMsg)
		if err != nil {
			errorMessage.ReturnJSONError(rw, map[string]string{"error": "Error marshaling JSON."}, http.StatusInternalServerError)
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
		rw.Write(jsonResponse)
		return
	}

	//go func() {
	//	_, err := s.GenerateVehiclesReportPDF()
	//	if err != nil {
	//		log.Printf("Error generating PDF: %v", err)
	//	}
	//}()

	jsonResponse, err := json.Marshal(vehicles)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (s *VehicleHandler) GetAllVehiclesByCategoryAndYear(c *gin.Context) {
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
	category := c.Param("category")
	yearStr := c.Param("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		errorMsg := map[string]string{"error": "Invalid year parameter"}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	vehicles, err := s.service.GetAllVehiclesByCategoryAndYear(domain.Category(category), year)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve registered vehicles from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(vehicles)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (s *VehicleHandler) GetAllRegisteredVehiclesByCategory(c *gin.Context) {
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
	category := c.Param("category")
	if err != nil {
		errorMsg := map[string]string{"error": "Invalid year parameter"}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
		return
	}

	vehicles, err := s.service.GetAllRegisteredVehiclesByCategory(domain.Category(category))
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve registered vehicles from the database."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(vehicles)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func (s *VehicleHandler) GetNumberOfRegisteredVehiclesByCategory(c *gin.Context) {
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
	category := c.Param("category")

	count, err := s.service.GetNumberOfRegisteredVehiclesByCategory(domain.Category(category))
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve number of registered vehicles by category."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	response := map[string]int64{"count": count}
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

func (s *VehicleHandler) GetVehicleByID(c *gin.Context) {
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

	allowedRoles := map[data.UserRole]bool{
		data.TrafficPoliceman: true,
		data.Policeman:        true,
		data.Employee:         false,
		data.Judge:            false,
	}

	if !allowedRoles[responseUser.LoggedInUser.UserRole] {
		errorMsg := map[string]string{"error": "Unauthorized. You do not have access to this resource."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	vehicleID := c.Param("id")

	vehicle, err := s.service.GetVehicleByID(vehicleID, ctx)
	if err != nil {
		errorMsg := map[string]string{"error": "Failed to retrieve vehicle from the database.No such vehicle."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(vehicle)
	if err != nil {
		errorMsg := map[string]string{"error": "Error marshaling JSON."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}

func isValidRegistrationPlate(registrationPlate string) bool {
	pattern := `^(NS|BG|BP)\d{3}[A-Z]{2}$`
	regex := regexp.MustCompile(pattern)
	registrationPlate = strings.ToUpper(registrationPlate)
	return regex.MatchString(registrationPlate)
}
