package genesis

import (
	"encoding/json"
	"os"
	"time"
)

type Genesis struct {
	Date                 time.Time         `json:"date"`
	ChainId              uint16            `json:"chain_id"`
	TransactionsPerBlock uint16            `json:"transactions_per_block"`
	Difficulty           uint16            `json:"difficulty"`
	MiningReward         uint64            `json:"mining_reward"`
	GasPrice             uint64            `json:"gas_price"`
	Balances             map[string]uint64 `json:"balances"`
}

func Load() (Genesis, error) {
	path := "zblock/genesis"
	content, err := os.ReadFile(path)
	if err != nil {
		return Genesis{}, err
	}

	var genesis Genesis
	if err := json.Unmarshal(content, &genesis); err != nil {
		return Genesis{}, err
	}

	return genesis, nil
}
