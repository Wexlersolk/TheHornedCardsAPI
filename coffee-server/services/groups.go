package services

import (
	"context"
)

type Group struct {
	ID         string `json:"id"`
	Group_name string `json:"group_name"`
	Group_info string `json:"group_info"`
}

func (g *Group) CreateGroup(group Group) (*Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        INSERT INTO thehorned_groups_table (group_name, group_info)
        VALUES ($1, $2) returning *
    `

	err := db.QueryRowContext(
		ctx,
		query,
		group.Group_name,
		group.Group_info,
	).Scan(
		&group.ID,
		&group.Group_name,
		&group.Group_info,
	)

	if err != nil {
		return nil, err
	}

	return &group, nil
}
