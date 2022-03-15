package models

//new model, brand add admin

type Car struct {
	ID       string `json:"id"`
	VIN string `json:"vin"`
	Name string `json:"name"`
	Description string `json:"description"`
	Year int `json:"year"`
	Price float64 `json:"price"`
	PhotoPath string  `json:"photo_src"` // src - save - file server
	Filter `json:"relation"`
	TechSpec `json:"tech_spec"`
}

//add tables
type TechSpec struct {
	ID       string `json:"id"`
	HorsePower float32 `json:"horse_power"`
	DriveUnit string `json:"drive_unit"`
	EngineType string `json:"engine_type"`
	TypeBody string `json:"type_body"`
	Color string `json:"color"`
	Odometer float64 `json:"odometer"`
}

//search by descr || title
type Search struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

//embeded?
type Filter struct {
	ID int `json:"id"`
	CategoryID int `json:"category_id"`  //auto/ moto/ yacht
	YearID int `json:"year_id"`
	ModelID int `json:"model_id"`
	BrandID int `json:"brand_id"`
	StateCarID int  `json:"state_car_id"` // new/bu/damage
	SalerID int  `json:"user_id"`
	CountryID int `json:"country_id"`
	CityID int  `json:"city_id"`
}