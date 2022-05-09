package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
)

type BridgeRepository struct {
	db *sql.DB
}

func MachineBridgeInit(db *sql.DB) gallery.BridgeRepoInterface {
	return &BridgeRepository{
		db: db,
	}
}

func (ur BridgeRepository) GetListSrc(ctx context.Context, machineID int) ([]string, error) {
	query := `SELECT path FROM bazar_machine_image WHERE machine_id = $1`

	result := []string{}

	rows, err := ur.db.QueryContext(ctx, query, machineID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := ""
		if err = rows.Scan(
			&temp,
		); err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return result, nil
}

func (mr BridgeRepository) GetByID(ctx context.Context, id int) (*models.Machine, error) {

	var result models.Machine
	var err error

	var tempImg sql.NullString

	query := `SELECT
		usr.phone, usr.first_name, vin, title, phone, 
		description, year, price, odometer,
		horse_power, volume, ctgr.name, mdl.name,
		brd.name, ctr.name, ct.name, st.name, fl.name,
		drut.name, trns.name, bt.name, cr.name, img.path
	FROM bazar_machine AS mch
		LEFT JOIN bazar_user AS usr ON usr.user_id = mch.creator_id   
		LEFT JOIN bazar_category AS ctgr ON  ctgr.id = mch.category_id  
		LEFT JOIN bazar_model AS mdl ON mdl.id =  mch.model_id 
		LEFT JOIN bazar_brand AS brd ON  brd.id= mch.brand_id 
		LEFT JOIN bazar_country AS ctr ON  ctr.id = mch.country_id
		LEFT JOIN bazar_city AS ct ON  ct.id = mch.city_id 
		LEFT JOIN bazar_state AS st ON  st.id = mch.state_id
		LEFT JOIN bazar_fuel AS fl ON  fl.id = mch.fuel_id 
		LEFT JOIN bazar_drive_unit AS drut ON  drut.id = mch.drive_unit_id
		LEFT JOIN bazar_transmission AS trns ON trns.id =  mch.trans_type_id
		LEFT JOIN bazar_body_type AS bt ON bt.id = mch.body_type_id
		LEFT JOIN bazar_color AS cr ON cr.id =  mch.color_id
		LEFT JOIN bazar_machine_image AS img ON img.machine_id =  mch.machine_id

	WHERE mch.machine_id = $1`

	err = mr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Creator.Phone,
		&result.Creator.FirstName,
		&result.VIN,
		&result.Title,
		&result.Creator.Phone,
		&result.Description,
		&result.Year,
		&result.Price,
		&result.Odometer,
		&result.HorsePower,
		&result.Volume,
		&result.Category.Name,
		&result.Model.Name,
		&result.Brand.Name,
		&result.Country.Name,
		&result.City.Name,
		&result.State.Name,
		&result.Fuel.Name,
		&result.DriveUnit.Name,
		&result.Transmission.Name,
		&result.BodyType.Name,
		&result.Color.Name,
		&tempImg,
	)

	if err != nil {
		return nil, err
	}

	if tempImg.Valid {
		result.MainImage = tempImg.String
	}

	images, err := mr.GetListSrc(ctx, id)
	result.Images = images

	if err != nil {
		return nil, err
	}

	return &result, nil
}
