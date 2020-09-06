package service

// ContactService use to access Edit, Add, Delete, Get and GetAll
type ContactService struct {
	// Essential Parameter
}

// NewContactService Return New ContactController Object
func NewContactService() *ContactService {
	return &ContactService{}
}
