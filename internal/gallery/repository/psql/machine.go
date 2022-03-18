package psql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/devstackq/bazar/internal/gallery/models"
)

type MachineRepository struct {
	db *sql.DB
}

func MachineRepoInit(db *sql.DB) *MachineRepository {
	return &MachineRepository{
		db: db,
	}
}

func (mr MachineRepository) GetByID(ctx context.Context, id int) (*models.Machine,  error) {

	var result models.Machine
	var err error

	query := `SELECT
		u.phone, u.first_name,
		vin, title, description, year, price, city_id, m.country_id, category_id,
		state_id, brand_id, model_id, creator_id, fuel_id, drive_unit_id,
		trans_type_id, body_type_id, color_id, odometer, horse_power, volume,
	FROM bazar_machine AS m LEFT JOIN bazar_user AS u  ON m.creator_id = u.user_id  WHERE machine_id = $1`

	err = mr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Creator.Phone,
		&result.Creator.FirstName,
		&result.VIN,
		&result.Title,
		&result.Description,
		&result.Year,
		&result.Price,
		&result.City.ID,
		&result.Country.ID,
		&result.Category.ID,
		&result.State.ID,
		&result.Brand.ID,
		&result.Brand.Model.ID,
		&result.Creator.ID,
		&result.Fuel.ID,
		&result.DriveUnit.ID,
		&result.Transmission.ID,
		&result.BodyType.ID,
		&result.Color.ID,
		&result.Odometer,
		&result.HorsePower,
		&result.Volume,
	)
	log.Println(err,1)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (mr MachineRepository) GetListByUserID(ctx context.Context, id int)([]*models.Machine,  error) {
	
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
	volume
	FROM bazar_machine where creator_id = $1`
	
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

func (mr MachineRepository) GetList(ctx context.Context)([]*models.Machine,  error) {
	
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
	FROM bazar_machine`
	result := []*models.Machine{}

	rows, err := mr.db.QueryContext(ctx, query)
	
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

		result = append(result, &temp)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return result, nil
}

func (mr MachineRepository) Create(ctx context.Context, item *models.Machine)(id int, err error) {
	sqlQuery := `INSERT INTO bazar_machine(
		vin, title, description, year, price, created_at, updated_at, city_id, country_id, category_id,
		state_id, brand_id, model_id, creator_id, fuel_id, drive_unit_id,
		trans_type_id, body_type_id, color_id, odometer, horse_power, volume)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22) RETURNING machine_id`
		
	row := mr.db.QueryRowContext(ctx, sqlQuery, item.VIN, item.Title, item.Description, item.Year, item.Price, time.Now(),  time.Now(),
		 item.City.ID, item.Country.ID, item.Category.ID, item.State.ID, item.Brand.ID,
		item.Brand.Model.ID, item.Creator.ID, item.Fuel.ID, item.DriveUnit.ID,
		item.Transmission.ID, item.BodyType.ID, item.Color.ID, item.Odometer, item.HorsePower, item.Volume)
	
		err = row.Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}
