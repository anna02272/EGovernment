package handlers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
	errorMessage "vehicles-service/error"
	"vehicles-service/services"
)

type VehicleHandler struct {
	service services.VehicleService
	DB      *mongo.Collection
}

func NewVehicleHandler(service services.VehicleService, db *mongo.Collection) VehicleHandler {
	return VehicleHandler{
		service: service,
		DB:      db,
	}

}

type Method string

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
)

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

func (s *VehicleHandler) CreateVehicleDriver(rw http.ResponseWriter, h *http.Request) {
	token := h.Header.Get("Authorization")
	url := "https://auth-service:8085/api/users/currentUser"

	timeout := 5 * time.Second // Adjust the timeout duration as needed
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := s.performAuthorizationRequestWithContext("POST", ctx, token, url)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			errorMessage.ReturnJSONError(rw, "Authorization service is not available.", http.StatusBadRequest)
			return
		}

		errorMessage.ReturnJSONError(rw, "Error performing authorization request", http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode != 200 {
		errorMsg := map[string]string{"error": "Unauthorized."}
		errorMessage.ReturnJSONError(rw, errorMsg, http.StatusUnauthorized)
		return
	}

	//decoder := json.NewDecoder(resp.Body)
	//
	//// Define a struct to represent the JSON structure
	//var responseUser struct {
	//	LoggedInUser struct {
	//		username string        `json:"username"`
	//		email    string        `json:"email"`
	//		UserRole data.UserRole `json:"userRole"`
	//	} `json:"user"`
	//}
	//
	//notification, exists := c.Get("notification")
	//if !exists {
	//	s.logger.WithFields(logrus.Fields{"path": "notification/createNotification"}).Error("Notification not found in context")
	//	span.SetStatus(codes.Error, "Notification not found in context")
	//	error2.ReturnJSONError(rw, "Notification not found in context", http.StatusBadRequest)
	//	return
	//}
	//
	//notif, ok := notification.(domain.NotificationCreate)
	//if !ok {
	//	fmt.Println(notif)
	//	errorMsg := map[string]string{"error": "Invalid type for notification."}
	//	error2.ReturnJSONError(rw, errorMsg, http.StatusBadRequest)
	//	return
	//}
	//
	//insertedNotif, _, err := s.notificationService.InsertNotification(&notif, spanCtx)
	//if err != nil {
	//	span.SetStatus(codes.Error, err.Error())
	//	error.ReturnJSONError(rw, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//rw.WriteHeader(http.StatusCreated)
	//jsonResponse, err1 := json.Marshal(insertedNotif)
	//if err1 != nil {
	//	error.ReturnJSONError(rw, fmt.Sprintf("Error marshaling JSON: %s", err1), http.StatusInternalServerError)
	//	return
	//}
	//rw.Write(jsonResponse)

}
