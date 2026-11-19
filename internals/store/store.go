package store

import (
	"context"
	"database/sql"
)

type Store struct{
	Poll interface{
		CreatePoll(context.Context, *PollBody) error
		GetPoll(context.Context, *PollBody) (*Poll,error)
		VotePoll(context.Context,*Vote)error
	}
}

func NewStorage(db *sql.DB) *Store{
	return &Store{
		Poll: &PollStore{db: db},
	}
}

