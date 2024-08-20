package blockchain

import (
	"fmt"

	"github.com/dougmendes/blockchain-property-registry/internal/models"
)

var Blockchain []models.Block

func calculateHash(block models.Block) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data.ID, block.PrevHash, block.Data.Owner)
	return fmt.Sprintf("%x", record)
}
