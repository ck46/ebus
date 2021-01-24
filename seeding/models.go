package seeding

type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	Name        string   `json:"name"`
	Store       string   `json:"store_name"`
	PromoImage  string   `json:"promo_image"`
	Categories  []string `json:"categories"`
	Description string   `json:"description"`
	Stock       uint     `json:"stock"`
	Price       float32  `json:"price"`
	Images      []string `json:"item_images"`
}

type Departments struct {
	Departments []Department `json:"departments"`
}

type Department struct {
	Name       string   `json:"name"`
	Categories []string `json:"categories"`
}

type Categories struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Stores struct {
	Stores []Store `json:"stores"`
}

type Store struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

type Statuses struct {
	Statuses []Status `json:"statuses"`
}

type Status struct {
	Name                string `json:"name"`
	Description         string `json:"description"`
	CompleteTransaction bool   `json:"complete_trn"`
}

type PaymentOptions struct {
	PaymentOptions []PaymentOption `json:"payment_options"`
}

type PaymentOption struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
