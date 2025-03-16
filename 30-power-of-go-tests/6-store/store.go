package store

import (
	"errors"
	"os"
)

var ErrUnopenable error = errors.New("can't open store file")

type Store struct{}

func Open(path string) (*Store, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, ErrUnopenable // perte d'informations sur err
	}

	f.Close()
	return &Store{}, nil
}

// func Open(path string) (*Store, error) {
// 	 f, err := os.Open(path)
// 	 if err != nil {
// 		return nil, err
// 	 }

// 	 f.Close()
// 	 return &Store{}, nil
// }
