package kaeufer

// Contact contains contact data from the Leads API
type Contact struct {
	Company     string `json:"company"`
	Salutation  string `json:"salutation"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	City        string `json:"city"`
	CountryCode string `json:"country_code"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Mobile      string `json:"mobil"` // really is mobil
	Fax         string `json:"fax"`
}

// Answer contains a string answer to a question.
type Answer struct {
	Answer string `json:"answer"`
}

// AnsweredQuestion data from API
type AnsweredQuestion struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

// Lead data from API
type Lead struct {
	ID                  int                `json:"lead_id"`
	VendorID            int                `json:"vendor_id"`
	ExternalAPIID       string             `json:"external_api_id"`
	StatusID            int                `json:"status_id"`
	StatusName          string             `json:"status_name"`
	Price               int                `json:"price"`
	BoughtAt            string             `json:"bought_at"`
	BoughtAtUnix        int                `json:"bought_at_unix"`
	ProductName         string             `json:"product_name"`
	ProductUUID         string             `json:"product_uuid"`
	Reachabiliy         string             `json:"reachability"`
	Talkability         string             `json:"talkability"`
	Appointment         string             `json:"appointment"`
	CallingNotes        string             `json:"calling_notes"`
	OfferContact        Contact            `json:"offer_contact"`
	InstallationContact Contact            `json:"installation_contact"`
	AnsweredQuestions   []AnsweredQuestion `json:"answered_questions"`
}
