package main

import (
	"fmt"
	"net/http"

	"github.com/Pradyumnasavadatti/poll/internals/store"
)

type PollData struct{
	Name string	`json:"name"`	
	Options []string `json:"options"`
	Id int64 `json:"id"`
}

type Message struct{
	Message string `json:"message"`
}

type Vote struct{
	Option Option `json:"option"`
	VotesCount int64 `json:"votesCount"`
}

type Option struct{
	OptionName string `json:"optionName"`
	OptionId int64	`json:"optionId"`
	PollId int64	`json:"pollId"`
}

func (app * application) createPoll(w http.ResponseWriter, r *http.Request){
	var pollData PollData
	if err := ReadJson(w,r,&pollData); err!=nil{
		fmt.Printf("Error: while getting the body: %s",err.Error())
		wb := &Message{
			Message: "Unable to scan the body",
		}
		WriteJson(w,wb)
	}
	ctx := r.Context()
	payload := &store.PollBody{
		Name: pollData.Name,
		Options: pollData.Options,
	}
	if err:= app.store.Poll.CreatePoll(ctx, payload); err!=nil{
		fmt.Printf("Error: while executing the query: %s",err.Error())
		WriteJson(w,&Message{
			Message: "Unable to create the poll, Please try again later",
		})
		return 
	}
	WriteJson(w,&Message{
		Message: "Poll Created Successfully!",
	})

}
func (app * application) getPoll(w http.ResponseWriter, r *http.Request){
	var id PollData;
	if err:=ReadJson(w,r,&id); err!=nil{
		fmt.Printf("Unable to read the id %s \n",err.Error())
		WriteJson(w,&Message{
			Message: "Unable to read the id",
		})
		return
	}
	payload := &store.PollBody{
		Id: id.Id,
	}
	ctx := r.Context()
	result, err := app.store.Poll.GetPoll(ctx,payload)
	if err!=nil{
		fmt.Printf("Unable get the poll %s \n",err.Error())
		WriteJson(w,&Message{
			Message: "Unable get the poll",
		})
	}
	WriteJson(w,&result)
}
func (app * application) votePoll(w http.ResponseWriter, r *http.Request){
	var option Vote
	if err := ReadJson(w,r,&option); err!=nil{
		WriteError(w,err,"Unable to read the json")
		return
	}

	ctx := r.Context()

	payload := &store.Vote{
		Option: &store.PollOptions{
			PollId: option.Option.PollId,
			OptionId: option.Option.OptionId,
		},
		VotesCount: int(option.VotesCount),
	}

	err := app.store.Poll.VotePoll(ctx,payload)

	if err!=nil{
		WriteError(w,err,"Unable to vote now")
		return
	}

	WriteJson(w,&Message{
		Message: "Successfully voted!",
	})
}