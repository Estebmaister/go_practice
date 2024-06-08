package config_test

import (
	"testing"

	"battle-of-monsters/app/config"
	utilstests "battle-of-monsters/app/tests/utils"
)

func TestLoad(t *testing.T) {
	utilstests.LoadEnv()

	if config.ENV.Port != "4000" {
		t.Errorf("Port expected as 4000 but got %v", config.ENV.Port)
	}

	if config.ENV.DBDriver != "sqlite" {
		t.Errorf("Driver expected as sqlite but got %v", config.ENV.DBDriver)
	}

	if config.ENV.DBName != "db/db.test.sqlite" {
		t.Errorf("Database name expected as db/db.test.sqlite but got %v", config.ENV.DBName)
	}
}
