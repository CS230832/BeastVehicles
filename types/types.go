package types

type UserRole string

const (
	Root    UserRole = "root"
	CEO     UserRole = "ceo"
	Manager UserRole = "manager"
)

type (
	UserStore interface {
		AddUser(user *UserRegisterPayload) error
		RemoveUser(username string) error
		GetUserByUserName(username string) (*UserPayload, error)

		AddLoginToken(username string, token string) error
		RemoveLoginToken(username string, token string) error
		HasLoginToken(username string, token string) (bool, error)
		RemoveAllLoginTokens(username string) error
	}
)

type (
	UserPayload struct {
		ID          int      `json:"-"`
		UserName    string   `json:"username,omitempty"`
		Password    string   `json:"-"`
		Role        UserRole `json:"role,omitempty"`
		FirstName   string   `json:"first_name,omitempty"`
		LastName    string   `json:"last_name,omitempty"`
		ParkingName string   `json:"parking,omitempty"`
	}

	UserRegisterPayload struct {
		UserName    string   `json:"username"`
		Password    string   `json:"password"`
		Role        UserRole `json:"role"`
		FirstName   string   `json:"first_name"`
		LastName    string   `json:"last_name"`
		ParkingName string   `json:"parking"`
	}

	UserLoginPayload struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
)
