package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)


type PollStore struct{
	db *sql.DB
}

type PollBody struct{
	Name string
	Options []string
	Id int64
}

type PollOptions struct{
	OptionName string
	VotesCount int64
	OptionId int64
	PollId int64
}

type Vote struct{
	VotesCount int
	Option *PollOptions
}

type Poll struct{
	Poll PollBody
	PollOptions []PollOptions
}

func (ps *PollStore) CreatePoll(ctx context.Context, pb *PollBody)error{
	if pb.Name == ""{
		return errors.New(`Error: poll name is required`)
	}
	if len(pb.Options) == 0 {
        return errors.New("poll must have at least one option")
    }
	tx, err := ps.db.BeginTx(ctx,nil)
	if err != nil{
		return fmt.Errorf(`Error while begining the stmt %w`,err)
	}
	defer tx.Rollback()

	var id int = 0
	err = tx.QueryRowContext(ctx, `INSERT INTO POLLS (NAME) VALUES ($1) RETURNING ID`,pb.Name).Scan(&id)
	if err != nil{
		return fmt.Errorf(`Error while creating the poll %w`,err)
	}

	stmt1,err := tx.PrepareContext(ctx, `INSERT INTO POLL_OPTIONS (OPTION_NAME,POLL_ID) VALUES ($1,$2)`)
	if err != nil{
		return fmt.Errorf(`Error while creating the prepared stmt %w`,err)
	}
	defer stmt1.Close()

	for _, item := range pb.Options{
		if _, err := stmt1.ExecContext(ctx, item, id); err!=nil{
			return fmt.Errorf(`Error while inserting the options %w`,err)
		}
	}
	return tx.Commit()
}

func (ps *PollStore) GetPoll(ctx context.Context,pb *PollBody) (*Poll,error){
	var pollResult PollBody;
	if err := ps.db.QueryRowContext(ctx,`SELECT ID, NAME FROM POLLS WHERE ID = $1`,pb.Id).Scan(&pollResult.Id,&pollResult.Name); err!=nil{
		return nil, fmt.Errorf("Error while querrying the poll %w",err)
	}

	var pollOptions []PollOptions;

	rows, err := ps.db.QueryContext(ctx,`SELECT OPTION_NAME, VOTES_COUNT FROM POLL_OPTIONS WHERE POLL_ID = $1`,pb.Id)
	if(err!=nil){
		return nil, fmt.Errorf("Error while fetching the options %w",err)
	}
	defer rows.Close()

	for(rows.Next()){
		var po PollOptions
		if err := rows.Scan(&po.OptionName,&po.VotesCount); err != nil{
			return nil,fmt.Errorf("Error : %w",err)
		}
		pollOptions = append(pollOptions,po)
		pollResult.Options = append(pollResult.Options, po.OptionName)
	}

	if err := rows.Err(); err!=nil{
		return nil,fmt.Errorf("Error while fetching the options: %w",err)
	}
	return &Poll{
		Poll: pollResult,
		PollOptions: pollOptions,
	},nil

}

func (ps *PollStore)VotePoll(ctx context.Context,vote *Vote)error{
	var voteCount int
	err := ps.db.QueryRowContext(ctx, `SELECT VOTES_COUNT FROM POLL_OPTIONS WHERE ID = $1 AND POLL_ID = $2`,vote.Option.OptionId,vote.Option.PollId).Scan(&voteCount)
	if(err!=nil){
		return fmt.Errorf("Unable to query the votes: %w",err)
	}

	finalCount := voteCount + int(vote.VotesCount)

	res, err := ps.db.ExecContext(ctx,`UPDATE POLL_OPTIONS SET VOTES_COUNT = $1 WHERE ID = $2 AND POLL_ID = $3`,finalCount,vote.Option.OptionId,vote.Option.PollId)

	if err!=nil{
		return fmt.Errorf("Unable to update the vote count: %w",err)
	}

	rowEffected, err := res.RowsAffected()

	if rowEffected == 0 || err!=nil {
		if err==nil{
			err = fmt.Errorf("Option not found")
		}
		return fmt.Errorf("Unable to update the vote: %w",err)
	}

	return nil
}