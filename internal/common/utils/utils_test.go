package utils

import (
	"okgopro/internal/common/utils/snowflake"
	"testing"

	"github.com/google/uuid"
)

func TestGenAppKeyOrSecret(t *testing.T) {
	u1 := uuid.New()
	u2, _ := uuid.NewRandom()
	t.Log("u1", u1)
	t.Log("u2", u2)

}
func TestGenAppKeyOrSecret2(t *testing.T) {
	u2, _ := snowflake.NewSnowflake(1, 1)

	t.Log("u2", u2.NextVal())

}
