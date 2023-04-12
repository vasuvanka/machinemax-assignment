package definitions

type RegistrationState string

const (
	RegistrationStateSuccess   RegistrationState = "success"
	RegistrationStateFailure   RegistrationState = "failure"
	RegistrationStateDuplicate RegistrationState = "duplicate"
)

type DevEUIRegistrationRequestDTO struct {
	DevEUI DevEUI `json:"deveui"`
}

type DevEUIRegistrationResponseDTO struct {
	DevEUI  DevEUI            `json:"deveui"`
	State   RegistrationState `json:"state"`
	Message string            `json:"message,omitempty"`
}

type DevEUI string

func (entity *DevEUI) String() string {
	return string(*entity)
}

func (entity *DevEUI) ShortCode() string {
	uid := *entity
	code := uid[len(uid)-5:]
	return string(code)
}
