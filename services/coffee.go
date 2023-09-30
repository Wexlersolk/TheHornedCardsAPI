package services


import "time"

type Coffe struct {
	ID string 'json:"id"'
	Name string 'json:"name"'
	Roast string 'json:"roast "'
	Region string 'json:"region"'
	Price string 'json:"price"'
	GrindUnit string 'json:"grind_unit"'
	CreatedAt time.Time 'json:"created_at"'
	UpdatedAt time.Time 'json:"updated_at"'

}

func (c *Coffe) GetAllCoffes() ([]*Coffe, error){
	ctx, cansel := context.WithTimeout(context.Background(), dbTimeout)
	defer cansel()

	query := 'select * from coffees'

	rows, err := db.QueryContext(ctx, query)
	if err != nil{
		return nil, err
	}

	var coffees []*Coffe
	for rows.Next(){
		var coffe Coffe
		err := rows.Scan(
			&coffe.ID,
			&coffe.Name,
			&coffe.Roast,
			&coffe.Price,
			&coffe.GrindUnit,
			&coffe.CreatedAt,
			&coffe.UpdatedAt
		)

		if err!=nil{
			return nil, err
		}

		coffees = append(coffes, &coffee)

	} 
	return coffes, nil
}