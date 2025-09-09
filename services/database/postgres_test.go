package database

import (
	"app/cmd/emails_sender_service/config"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPostgres(t *testing.T) {
	ctx := context.Background()
	conf := config.LoadConfig()
	fmt.Println(conf)
	db, err := NewPostgres(ctx, conf.Database)
	require.NoError(t, err)
	defer db.Close()

	require.NotNil(t, db.GetPool(), "Expected non-nil connection pool")
}
