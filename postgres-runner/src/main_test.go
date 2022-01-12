package src

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_connectAndExec(t *testing.T) {
	const connstring = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	tts := []struct {
		name       string
		command    string
		connstring string
		wantErr    error
	}{
		{
			name:       "happy path - multiple statements",
			command:    "SELECT current_database(); SELECT current_database(); select * from pg_stat_activity;",
			connstring: connstring,
		},
		{
			name:       "error - returns error",
			command:    "SELECT bad syntax",
			connstring: connstring,
			wantErr:    fmt.Errorf(`execing sql query: 42703, column "bad" does not exist`),
		},
		{
			name:       "error - bad connection string returns error",
			command:    "SELECT current_database();",
			connstring: "bad conn string",
			wantErr:    fmt.Errorf("%w", fmt.Errorf(`connecting to database:`)),
		},
		{
			name:       "error - can't login to database returns error",
			command:    "SELECT current_database();",
			connstring: "postgres://postgres:postgres@locallllhost:5432/postgres?sslmode=disable",
			wantErr:    fmt.Errorf(`connecting to database:`),
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			err := connectAndExec(tt.connstring, tt.command)
			if tt.wantErr == nil || err == nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			}
		})
	}
}
