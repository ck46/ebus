package api

const (
	secret = "secretkey"
	decode = "decoded"
)

type User struct {
	ID        uint
	Email     string
	FirstName string
	LastName  string
	Password  string
	Phone     string
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code  int
	Error string
}

type Item struct {
	ID          uint     `json:"item_id"`
	Name        string   `json:"item_name"`
	StoreName   string   `json:"store_name"`
	PromoImage  string   `json:"default_image"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Stock       uint     `json:"stock"`
}

type Store struct {
	ID       uint
	Name     string
	Image    string
	Catalogs []*Item
}

type Response struct {
	Code   int
	Result interface{}
}
