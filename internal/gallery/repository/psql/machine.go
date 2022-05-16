package psql

import (
	"context"
	"database/sql"
	"time"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
	"github.com/lib/pq"
)

type MachineRepository struct {
	db     *sql.DB
	bridge gallery.BridgeRepoInterface
}

func MachineRepoInit(db *sql.DB, bridge gallery.BridgeRepoInterface) gallery.MachineRepoInterface {
	return &MachineRepository{
		db:     db,
		bridge: bridge,
	}
}

func (mr MachineRepository) GetByID(ctx context.Context, id int) (*models.Machine, error) {
	var result *models.Machine
	var err error
	result, err = mr.bridge.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// or company_id?

func (mr MachineRepository) GetListByUserID(ctx context.Context, id float64, pageNum int) ([]*models.Machine, error) {
	limit := 9
	query := `SELECT
	machine_id,
	title,
	price,
	created_at,
	mdl.name,
	brd.name
	FROM bazar_machine  AS mch
	LEFT JOIN bazar_model AS mdl ON mdl.id =  mch.model_id 
	LEFT JOIN bazar_brand AS brd ON  brd.id= mch.brand_id 
	WHERE creator_id = $1
	ORDER BY created_at DESC LIMIT $2 OFFSET $3`

	pageNum = limit * (pageNum - 1)

	result := []*models.Machine{}

	rows, err := mr.db.QueryContext(ctx, query, id, limit, pageNum)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.Machine{}

		if err = rows.Scan(
			&temp.ID,
			&temp.Title,
			&temp.Price,
			&temp.CreatedAt,
			&temp.Brand.Model.Name,
			&temp.Brand.Name,
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

// get list machines with pagination
func (mr MachineRepository) GetList(ctx context.Context, pageNum int) ([]*models.Machine, error) {
	limit := 9
	result := []*models.Machine{}

	query := `SELECT
	machine_id,
	vin,
	title,
	description,
	year,
	price,
	odometer,
	created_at,
	mdl.name,
	brd.name
	FROM bazar_machine  AS mch
	LEFT JOIN bazar_model AS mdl ON mdl.id =  mch.model_id 
	LEFT JOIN bazar_brand AS brd ON  brd.id= mch.brand_id 
	ORDER BY created_at DESC LIMIT $1 OFFSET $2 `

	pageNum = limit * (pageNum - 1)

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
			&temp.CreatedAt,
			&temp.Brand.Name,
			&temp.Brand.Model.Name,
		); err != nil {
			return nil, err
		}
		// get first photo
		srcQuery := `SELECT paths_img FROM bazar_machine_image WHERE machine_id = $1`
		err = mr.db.QueryRowContext(ctx, srcQuery, temp.ID).Scan(pq.Array(&temp.Images))
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
		item.Country.City.ID, item.Country.ID, item.Category.ID, item.State.ID, item.Brand.ID,
		item.Brand.Model.ID, item.Creator.ID, item.Fuel.ID, item.DriveUnit.ID,
		item.Transmission.ID, item.BodyType.ID, item.Color.ID, item.Odometer, item.HorsePower, item.Volume)

	err = row.Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}
