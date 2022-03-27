package psql

import (
	"context"
	"database/sql"
	"time"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
)

type MachineRepository struct {
	db *sql.DB
}

func MachineRepoInit(db *sql.DB) gallery.MachineRepoInterface {
	return &MachineRepository{
		db: db,
	}
}

func (mr MachineRepository) GetByID(ctx context.Context, id int) (*models.Machine, error) {
	var result models.Machine
	var err error

	query := `SELECT
		usr.phone, usr.first_name, vin, title,
		description, year, price, odometer,
		horse_power, volume, ctgr.name, mdl.name,
		brd.name, ctr.name, ct.name, st.name, fl.name,
		drut.name, trns.name, bt.name, cr.name
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
		LEFT JOIN bazar_trans AS trns ON trns.id =  mch.trans_type_id
		LEFT JOIN bazar_body_type AS bt ON bt.id = mch.body_type_id
		LEFT JOIN bazar_color AS cr ON cr.id =  mch.color_id  
	WHERE machine_id = $1`

	err = mr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Creator.Phone,
		&result.Creator.FirstName,
		&result.VIN,
		&result.Title,
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
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (mr MachineRepository) GetListByUserID(ctx context.Context, id int) ([]*models.Machine, error) {
	query := `SELECT
	machine_id,
	vin,
	title,
	description, 
	year,
	price,
	odometer,
	created_at,
	horse_power,
	volume,
	FROM bazar_machine
	WHERE creator_id = $1`

	result := []*models.Machine{}

	rows, err := mr.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.Machine{}

		if err = rows.Scan(
			&temp.ID,
			&temp.VIN,
			&temp.Title,
			&temp.Description,
			&temp.Year,
			&temp.Price,
			&temp.Odometer,
			&temp.CreatedAt,
			&temp.HorsePower,
			&temp.Volume,
		); err != nil {
			return nil, err
		}

		result = append(result, &temp)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return result, nil
}

func (mr MachineRepository) GetList(ctx context.Context, pageNum int) ([]*models.Machine, error) {

	var limit = 9
	var result = []*models.Machine{}

	query := `SELECT
	machine_id,
	vin,
	title,
	description, 
	year,
	price,
	odometer,
	volume, 
	created_at,
	horse_power
	FROM bazar_machine
	ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	pageNum = limit * (pageNum - 1)

	// log.Println(limit, pageNum)

	rows, err := mr.db.QueryContext(ctx, query, limit, pageNum)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.Machine{}

		if err = rows.Scan(
			&temp.ID,
			&temp.VIN,
			&temp.Title,
			&temp.Description,
			&temp.Year,
			&temp.Price,
			&temp.Odometer,
			&temp.Volume,
			&temp.CreatedAt,
			&temp.HorsePower,
		); err != nil {
			return nil, err
		}
		//get first photo
		srcQuery := `SELECT path FROM bazar_machine_image WHERE machine_id = $1`
		err = mr.db.QueryRowContext(ctx, srcQuery, temp.ID).Scan(&temp.MainImage)
		if err != nil && err.Error() != "sql: no rows in result set" {
			return nil, err
		}
		result = append(result, &temp)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return result, nil
}

func (mr MachineRepository) Create(ctx context.Context, item *models.Machine) (id int, err error) {
	sqlQuery := `INSERT INTO bazar_machine(
		vin, title, description, year, price, created_at, updated_at, city_id, country_id, category_id,
		state_id, brand_id, model_id, creator_id, fuel_id, drive_unit_id,
		trans_type_id, body_type_id, color_id, odometer, horse_power, volume)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22) RETURNING machine_id`

	row := mr.db.QueryRowContext(ctx, sqlQuery, item.VIN, item.Title, item.Description, item.Year, item.Price, time.Now(), time.Now(),
		item.City.ID, item.Country.ID, item.Category.ID, item.State.ID, item.Brand.ID,
		item.Brand.Model.ID, item.Creator.ID, item.Fuel.ID, item.DriveUnit.ID,
		item.Transmission.ID, item.BodyType.ID, item.Color.ID, item.Odometer, item.HorsePower, item.Volume)

	err = row.Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}
