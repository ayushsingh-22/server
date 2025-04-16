package models

type Admin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Query struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Service         string `json:"service"`
	Message         string `json:"message"`
	SubmittedAt     string `json:"submitted_at"`
	NumGuards       string `json:"numGuards"`
	DurationType    string `json:"durationType"`
	DurationValue   string `json:"durationValue"`
	CameraRequired  bool   `json:"cameraRequired"`
	VehicleRequired bool   `json:"vehicleRequired"`
	Status          string `json:"status"` // optional: pending/approved/rejected
}
