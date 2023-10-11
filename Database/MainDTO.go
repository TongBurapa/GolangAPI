package Database

type Config struct {
	User       string
	Password   string
	Host       string
	Port       int
	Name       string
	DisableTLS bool
}

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	Role     string `json:"Role"`
}

type UsersArr []struct {
	ID       int    `json:"id"`
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	Role     string `json:"Role"`
}

type UserDelete struct {
	ID []int `json:"id"`
}

type LoginAuth struct {
	Username string
	Password string
}
