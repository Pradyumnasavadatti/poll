package seed

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	st "github.com/Pradyumnasavadatti/poll/internals/store"
)

func Seeding(db *sql.DB, store *st.Store)error{
	ctx,cancel := context.WithTimeout(context.Background(), time.Minute * 3)
	defer cancel()
	for _,item := range(SeedPolls){
		var opts []string
		for _,i := range(item.Options){
			opts = append(opts, i.Name)
		}
		err := store.Poll.CreatePoll(ctx, &st.PollBody{
			Name: item.Name,
			Options: opts,
		})

		if err!=nil{
			return fmt.Errorf("Something went wrong %w",err)
		}
	}
	return nil
}