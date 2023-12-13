package services

import (
	"context"
)

type Group struct {
	ID         int    `json:"group_id"`
	Group_name string `json:"group_name"`
}

func (g *Group) GetAllGroups() ([]*Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT group_id, group_name FROM thehorned_groups_table`
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
			INSERT INTO thehorned_groups_table (group_name)
			VALUES ($1) returning group_id, group_name
		`

	err := db.QueryRowContext(
		ctx,
		query,
		group.Group_name,
	).Scan(
		&group.ID,
		&group.Group_name,
	)

	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (g *Group) DeleteGroupById(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM thehorned_groups_table WHERE group_id = $1`
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

func (g *Group) GetAllCardsFromGroup(groupID int) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT card_id
		FROM thehorned_cards_table
		WHERE group_id = $1
	`

	rows, err := db.QueryContext(ctx, query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cardIDs []string
	for rows.Next() {
		var cardID string
		err := rows.Scan(&cardID)
		if err != nil {
			return nil, err
		}
		cardIDs = append(cardIDs, cardID)
	}

	return cardIDs, nil
}
