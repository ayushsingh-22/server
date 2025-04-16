package models

type Admin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Query struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Email           string  `json:"email"`
	Phone           string  `json:"phone"`
	Service         string  `json:"service"`
	Message         string  `json:"message"`
	SubmittedAt     string  `json:"submitted_at"`
	NumGuards       string  `json:"numGuards"`
	DurationType    string  `json:"durationType"`
	DurationValue   string  `json:"durationValue"`
	CameraRequired  bool    `json:"cameraRequired"`
	VehicleRequired bool    `json:"vehicleRequired"`
	FirstAid        bool    `json:"firstAid"`
	WalkieTalkie    bool    `json:"walkieTalkie"`
	BulletProof     bool    `json:"bulletProof"`
	FireSafety      bool    `json:"fireSafety"`
	Status          string  `json:"status"`
	Cost            float64 `json:"cost"`
}

type ServiceRevenue struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type TopService struct {
	Service string  `json:"service"`
	Revenue float64 `json:"revenue"`
}

type MonthlyRevenue struct {
	Month   string  `json:"month"`
	Revenue float64 `json:"revenue"`
	Growth  float64 `json:"growth"`
}