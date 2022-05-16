package psql

import (
	"context"
	"database/sql"
	"log"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
	"github.com/lib/pq"
)

type FilterRepository struct {
	db     *sql.DB
	bridge gallery.BridgeRepoInterface
}

func FilterRepoInit(db *sql.DB, bridge gallery.BridgeRepoInterface) gallery.FilterRepoInterface {
	return &FilterRepository{
		db:     db,
		bridge: bridge,
	}
}

//carById - fieldName - change; pagination; check - filter - priceTo - priceFrom
//distinct on - delete duplicate

func (fr FilterRepository) GetListMachineByFilter(ctx context.Context, keys *models.QueryParams, pageNum int) ([]*models.Machine, error) {

	var result []*models.Machine
	var err error
	// query := `SELECT DISTINCT ON (mch.machine_id)
	// query += " GROUP BY mch.machine_id"

	query := `SELECT 
		mch.machine_id, usr.phone, usr.first_name, usr.company, vin, title,
		phone, description, year, price, odometer,
		horse_power, volume, ctgr.name, mdl.name,
		brd.name, ctr.name, ct.name, st.name, fl.name,
		drut.name, trns.name, bt.name, cr.name, img.paths_img,
		mch.created_at
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
		LEFT JOIN bazar_machine_image AS img ON img.machine_id =  mch.machine_id `

	query += prepareQuery(keys)
	log.Println(query, "query")

	rows, err := fr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := models.Machine{}

		if err = rows.Scan(
			&temp.ID,
			&temp.Creator.Phone,
			&temp.Creator.FirstName,
			&temp.Creator.Company,
			&temp.VIN,
			&temp.Title,
			&temp.Creator.Phone,
			&temp.Description,
			&temp.Year,
			&temp.Price,
			&temp.Odometer,
			&temp.HorsePower,
			&temp.Volume,
			&temp.Category.Name,
			&temp.Model.Name,
			&temp.Brand.Name,
			&temp.Country.Name,
			&temp.City.Name,
			&temp.State.Name,
			&temp.Fuel.Name,
			&temp.DriveUnit.Name,
			&temp.Transmission.Name,
			&temp.BodyType.Name,
			&temp.Color.Name,
			pq.Array(&temp.Images),
			&temp.CreatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, &temp)
	}
	// for _, v := range result {
	// 	log.Println(v.ID, "id car")
	// }
	return result, nil

}
