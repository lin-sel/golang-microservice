package controller

import (
	"net/http"

	"bitbucket.com/lin-sel/golang-microservice/service"
	"github.com/gorilla/mux"
)

// ContactController use to Edit, Add, Delete, Get and GetAll
type ContactController struct {
	ContactService *service.ContactService
}

// NewContactController Return New ContactController Object
func NewContactController(ser *service.ContactService) *ContactController {
	return &ContactController{
		ContactService: ser,
	}
}

// RouteSpecifier Register Endpoint to Server
func (contactController *ContactController) RouteSpecifier(route *mux.Router) {
	route.HandleFunc("/contact", contactController.AddContact).Methods(http.MethodPost)
	route.HandleFunc("/contact", contactController.AddContact).Methods(http.MethodGet)
}

// AddContact Add New Contact
func (contactController *ContactController) AddContact(w http.ResponseWriter, r *http.Request) {
	// Write Code Here
}

// GetContacts Get All Contact
func (contactController *ContactController) GetContacts(w http.ResponseWriter, r *http.Request) {
	// Write Code Here
}
