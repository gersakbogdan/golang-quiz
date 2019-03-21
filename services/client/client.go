package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"

	pb "github.com/gersakbogdan/fasttrackquiz/proto"
)

// Retrieve questions from rpc server
// and send them to `questionsChannel` channel in order to process them
func getQuestions(client pb.QuizerClient, questionsChannel chan *pb.Question) {
	stream, err := client.ListQuestions(context.Background(), &pb.Empty{})
	_errCheck(&err, "Failed to retrieve questions: %v")

	for {
		question, err := stream.Recv()
		if err == io.EOF {
			close(questionsChannel)
			break
		}
		_errCheck(&err, "ListQuestions(_) = _, %v")

		questionsChannel <- question
	}
}

// Read questions from `questionsChannel` channel
// and send to server one by one
func runRecordAnswers(client pb.QuizerClient, questionsChannel chan *pb.Question, done chan bool) {
	stream, err := client.RecordAnswers(context.Background())
	_errCheck(&err, "Failed to save answers: %v")

	for {
		question, ok := <-questionsChannel
		if ok {
			err := stream.Send(prompt(question))
			_errCheck(&err, "stream.Send(point) error: %v")
		} else {
			break // nothing to do here anymore
		}
	}

	result, err := stream.CloseAndRecv()
	_errCheck(&err, "stream.CloseAndRecv error: %v")

	displaySummary(result)
	done <- true
}

func Run(host string, port int) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	_errCheck(&err, "Failed to dial: %v")
	defer conn.Close()

	client := pb.NewQuizerClient(conn)

	log.Println("CLI Client started...")

	// Retrieve questions and save answers in a producer/consumer way
	// display one question, answer question and record (sent to server)
	// `questions` channels is used for coordination
	var done = make(chan bool)
	var questions = make(chan *pb.Question)

	// Retrieve questions in a new goroutine
	go getQuestions(client, questions)

	// Save answers
	go runRecordAnswers(client, questions, done)

	<-done // waits till last question is answered
}

func displaySummary(result *pb.Result) {
	fmt.Printf("\n\nCorrect answers: %d\n", result.CorrectAnswers)
	fmt.Printf("You scored better than %d%% of all quizers\n", result.BetterThen)
	fmt.Printf("Quiz solved in %.2fsecs\n\n", result.ElapsedTime)
}

func _errCheck(err *error, msg string) {
	if *err != nil {
		panic(fmt.Errorf(msg, *err))
	}
}
