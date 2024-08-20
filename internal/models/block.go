package models

import "github.com/dougmendes/blockchain-property-registry/internal/models"

type Block struct {
	Index     int
	Timestamp string
	Data      models.Property
	PrevHash  string
	Hash      string
}
