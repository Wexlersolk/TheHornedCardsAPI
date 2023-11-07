package services

import (
	"context"
)

type Group struct {
	ID         string `json:"id"`
	Group_name string `json:"group_name"`
	Group_info string `json:"group_info"`
}

func (g *Group) GetAllGroups() ([]*Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, group_name, group_info FROM thehorned_groups_table`
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var groups []*Group
	for rows.Next() {
		var group Group
		err := rows.Scan(
			&group.ID,
			&group.Group_name,
			&group.Group_info,
		)

		if err != nil {
			return nil, err
		}
		groups = append(groups, &group)
	}

	return groups, nil
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

func (g *Group) DeleteGroupById(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM thehorned_groups_table WHERE id = $1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (g *Group) DeleteGroupByName(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM thehorned_groups_table WHERE group_name = $1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (g *Group) DeleteAllGroups() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM thehorned_groups_table`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
