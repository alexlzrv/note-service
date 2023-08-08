package client

import (
	"context"
	"log"

	desc "github.com/alexlzrv/note-service/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const address = "localhost:50051"

func RunClient() {
	con, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("didn`t connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)

	//createNote(client)
	//getNote(client)
	//updateNote(client)
	deleteNote(client)
}

func createNote(client desc.NoteV1Client) {
	res, err := client.Create(context.Background(), &desc.CreateRequest{
		Note: &desc.Note{
			Title:  "new title",
			Text:   "new text",
			Author: "new author",
		},
	})

	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Id: ", res.GetId())
}

func getNote(client desc.NoteV1Client) {
	res, err := client.Get(context.Background(), &desc.GetRequest{
		Id: 2,
	})

	if err != nil {
		log.Println(err.Error())
	}

	log.Printf("Title: %s\nText: %s\nAuthor: %s\n\n", res.Note.GetTitle(),
		res.Note.GetText(), res.Note.GetAuthor())
}

func updateNote(client desc.NoteV1Client) {
	note := desc.UpdateNoteInfo{
		Title:  wrapperspb.String("title2"),
		Text:   wrapperspb.String("text2"),
		Author: wrapperspb.String("author2"),
	}

	_, err := client.Update(context.Background(), &desc.UpdateRequest{
		Id:       2,
		NoteInfo: &note,
	})

	if err != nil {
		log.Println(err.Error())
	}

	log.Print("Update note")
}

func deleteNote(client desc.NoteV1Client) {
	_, err := client.Delete(context.Background(), &desc.DeleteRequest{
		Id: 2,
	})

	if err != nil {
		log.Println(err.Error())
	}

	log.Print("Delete note")
}
