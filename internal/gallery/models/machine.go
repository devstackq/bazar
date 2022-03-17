package models

import "time"

//new model, brand add admin
//// src - save - file server; photoName, []byte

type Machine struct {
	ID       string `json:"id"`
	VIN string `json:"vin"`
	Title string `json:"title"`
	Description string `json:"description"`
	Year int `json:"year"`
	Price float64 `json:"price"`
	PhotoName string  `json:"photo_name"` 
	Photo []byte `json:"photo"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Odometer float64 `json:"odometer"` 
	HorsePower int `json:"horse_power"` 
	Transmission `json:"trans"` 
	Body `json:"body"`
	Color `json:"color"`
	Category  `json:"category"`
	Brand `json:"brand"`
	Model `json:"model"`
	State `json:"state"`
	Saler `json:"saler"`
	Country  `json:"country"`
	City  `json:"city"`
	Fuel  `json:"fuel"`
	DriveUnit  `json:"drive_unit"`
	// Filter `json:"relation"`
}


//search by descr || title
// type Search struct {
// 	ID int `json:"id"`
// 	Name string `json:"name"`
// 	Description string `json:"description"`
// }
//embeded?
type Filter struct {
	ID int `json:"id"`
	CategoryID int `json:"category_id"`  //auto/ moto/ yacht
	ModelID int `json:"model_id"`
	BrandID int `json:"brand_id"`
	StateCarID int  `json:"state_car_id"` // new/bu/damage
	SalerID int  `json:"saler_id"`
	CountryID int `json:"country_id"`
	CityID int  `json:"city_id"`
	FuelTypeID int `json:"fuel_type_id"` 
	DriveUnitID int `json:"drive_unit_id"` 
	TransmissionTypeID int `json:"trans_type_id"` 
	BodyTypeID int `json:"body_type_id"` 
	ColorID int `json:"color_id"` 
}

//todo: own endpoint, crud
type Transmission struct {
	ID int 
	Name string
}
type Body struct {
	ID int 
	Name string
}
type Color struct {
	ID int 
	Name string
}
type Category struct {
	ID int 
	Name string
}
type Model struct {
	ID int 
	Name string
}
type Brand struct {
	ID int 
	Name string
}
type State struct {
	ID int 
	Name string
}
type Saler struct {
	ID int 
	Name string
	Email string
}
type Country struct {
	ID int 
	Name string
}
type City struct {
	ID int 
	Name string
}
type Fuel struct {
	ID int 
	Name string
}
type DriveUnit struct {
	ID int 
	Name string
}