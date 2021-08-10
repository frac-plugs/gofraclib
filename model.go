package gofraclib

type Type string

type Status string

const (
	ADMIN Type = "ADMIN"
	FIELD Type = "FIELD"

	VERIFIED             Status = "VERIFIED"
	PENDING_VERIFICATION Status = "PENDING_VERIFICATION"
)

type ProviderAccount struct {
	UID        string `json:"uid" firestore:"uid"`
	Type       Type   `json:"type" firestore:"type"`
	Status     Status `json:"status" firestore:"status"`
	ProviderID string `json:"providerId" firestore:"providerId"`
}

type ClientAccount struct {
	UID               string   `json:"uid" firestore:"uid"`
	Status            Status   `json:"status" firestore:"status"`
	VerifiedProviders []string `json:"verifiedProviders" firestore:"verifiedProviders"`
	VerifiedContacts  []string `json:"verifiedContacts" firestore:"verifiedContacts"`
}

type User struct {
	UID         string `json:"uid" firestore:"uid"`
	PhoneNumber string `json:"phoneNumber" firestore:"phoneNumber"`
	Email       string `json:"email" firestore:"email"`
	Name        string `json:"name" firestore:"name"`
}
