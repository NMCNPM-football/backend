package models

import (
	"github.com/NMCNPM-football/backend/config"
	"github.com/NMCNPM-football/backend/internal/must"
	"log"
	"testing"
	"time"
)

func TestUpdatePlayerType(t *testing.T) {
	// Connect to the database

	cfg := config.ReadConfigAndArg()

	logger, sentry, err := must.NewLogger(cfg.SentryDSN, cfg.ServiceName+"-app")
	if err != nil {
		log.Fatalf("logger: %v", err)
	}

	defer logger.Sync()
	defer sentry.Flush(2 * time.Second)

	db := must.ConnectDb(cfg.Db)

	// Get all players
	var players []Player
	err = db.Raw("SELECT * FROM `players`").Scan(&players).Error
	if err != nil {
		t.Fatalf("Failed to select players: %v", err)
	}

	// Update TypePlayer for each player
	for _, player := range players {
		if player.Nationality == "VIE" {
			player.TypePlayer = "National Player"
		} else {
			player.TypePlayer = "Naturalization Player"
		}
		err = db.Model(&player).Updates(Player{TypePlayer: player.TypePlayer}).Error
		if err != nil {
			t.Fatalf("Failed to update player: %v", err)
		}
	}
}
