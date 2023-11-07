package services

import (
	"context"
	"time"
)

type Card struct {
	ID           string    `json:"id"`
	Group_name   string    `json:"group_name"`
	Card_hint    string    `json:"card_hint"`
	Display_word string    `json:"display_word"`
	Hidden_word  string    `json:"hidden_word"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (c *Card) GetAllCards() ([]*Card, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, group_name, card_hint, display_word, hidden_word, created_at, updated_at FROM thehorned_cards_table`
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var cards []*Card
	for rows.Next() {
		var card Card
		err := rows.Scan(
			&card.ID,
			&card.Group_name, // Fixed: Missing group_name in the SELECT query
			&card.Card_hint,  // Fixed: Missing card_hint in the SELECT query
			&card.Display_word,
			&card.Hidden_word,
			&card.CreatedAt,
			&card.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		cards = append(cards, &card)
	}

	return cards, nil
}
func (c *Card) GetCardById(id string) (*Card, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
        SELECT id, group_name, card_hint, display_word, hidden_word, created_at, updated_at 
        FROM thehorned_cards_table
        WHERE id = $1
    `
	var card Card

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&card.ID,
		&card.Group_name,
		&card.Card_hint,
		&card.Display_word,
		&card.Hidden_word,
		&card.CreatedAt,
		&card.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &card, nil
}
func (c *Card) CreateCard(card Card) (*Card, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        INSERT INTO thehorned_cards_table (group_name, card_hint, display_word, hidden_word, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6) returning *
    `

	err := db.QueryRowContext(
		ctx,
		query,
		card.Group_name,
		card.Card_hint,
		card.Display_word,
		card.Hidden_word,
		time.Now(),
		time.Now(),
	).Scan(
		&card.ID,
		&card.Group_name,
		&card.Card_hint,
		&card.Display_word,
		&card.Hidden_word,
		&card.CreatedAt,
		&card.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (c *Card) UpdateCard(id string, body Card) (*Card, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        UPDATE thehorned_cards_table
        SET
            group_name = $2,
            card_hint = $3,
            display_word = $4,
            hidden_word = $5,
            updated_at = $6
        WHERE id = $1
        returning *
    `

	err := db.QueryRowContext(
		ctx,
		query,
		id,
		body.Group_name,
		body.Card_hint,
		body.Display_word,
		body.Hidden_word,
		time.Now(),
	).Scan(
		&body.ID,
		&body.Group_name,
		&body.Card_hint,
		&body.Display_word,
		&body.Hidden_word,
		&body.CreatedAt,
		&body.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *Card) DeleteCard(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM thehorned_cards_table WHERE id = $1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *Card) DeleteAllCards() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM thehorned_cards_table`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
