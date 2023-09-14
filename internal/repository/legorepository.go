package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"legoapi/internal/app"
)

type Repository interface {
	GetAllLegoSets() ([]app.LegoSet, error)
	GetLegoSetByCode(code string) (app.LegoSet, error)
	CreateLegoSet(legoSet app.LegoSet) error
	UpdateLegoSet(code string, legoSet app.LegoSet) error
	DeleteLegoSet(code string) error
}

type LegoSetRepository struct {
	DB *sql.DB
}

func NewLegoSetRepository(db *sql.DB) *LegoSetRepository {
	return &LegoSetRepository{
		DB: db,
	}
}

func (r *LegoSetRepository) GetAllLegoSets() ([]app.LegoSet, error) {
	rows, err := r.DB.Query("SELECT code, name, piece_count, image_url, price FROM lego_sets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var legoSets []app.LegoSet
	for rows.Next() {
		var legoSet app.LegoSet
		err := rows.Scan(&legoSet.Code, &legoSet.Name, &legoSet.PieceCount, &legoSet.ImageURL, &legoSet.Price)
		if err != nil {
			return nil, err
		}
		legoSet.CostPerPiece = legoSet.Price / float64(legoSet.PieceCount)
		legoSets = append(legoSets, legoSet)
	}

	return legoSets, nil
}

func (r *LegoSetRepository) GetLegoSetByCode(code string) (app.LegoSet, error) {
	var legoSet app.LegoSet
	row := r.DB.QueryRow("SELECT code, name, piece_count, image_url, price FROM lego_sets WHERE code = $1", code)
	err := row.Scan(&legoSet.Code, &legoSet.Name, &legoSet.PieceCount, &legoSet.ImageURL, &legoSet.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return app.LegoSet{}, fmt.Errorf("LEGO seti bulunamadı: %s", code)
		}
		return app.LegoSet{}, err
	}
	legoSet.CostPerPiece = legoSet.Price / float64(legoSet.PieceCount)
	return legoSet, nil
}

func (r *LegoSetRepository) CreateLegoSet(legoSet app.LegoSet) error {
	_, err := r.DB.Exec("INSERT INTO lego_sets (code, name, piece_count, image_url, price) VALUES ($1, $2, $3, $4, $5)",
		legoSet.Code, legoSet.Name, legoSet.PieceCount, legoSet.ImageURL, legoSet.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *LegoSetRepository) UpdateLegoSet(code string, legoSet app.LegoSet) error {
	_, err := r.DB.Exec("UPDATE lego_sets SET name = $1, piece_count = $2, image_url = $3, price = $4 WHERE code = $5",
		legoSet.Name, legoSet.PieceCount, legoSet.ImageURL, legoSet.Price, code)
	if err != nil {
		return err
	}
	return nil
}

func (r *LegoSetRepository) DeleteLegoSet(code string) error {
	_, err := r.DB.Exec("DELETE FROM lego_sets WHERE code = $1", code)
	if err != nil {
		return err
	}
	return nil
}
