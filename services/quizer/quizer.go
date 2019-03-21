package quizer

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/gersakbogdan/fasttrackquiz/proto"

	"google.golang.org/grpc"
)

var mutex = &sync.Mutex{}

// Implements quizer.QuizServer
type quizer struct {
	questions map[int64]*pb.ServerQuestion // one time loading during server initialization
	results   []int32                      // number of correct answers for all quizers (used for `better than`)
}

// server-to-client streaming RPC
// Questions are streamed rather than returned at once
func (s *quizer) ListQuestions(req *pb.Empty, stream pb.Quizer_ListQuestionsServer) error {
	for _, question := range s.questions {
		if err := stream.Send(question.Question); err != nil {
			return err
		}
	}

	return nil
}

// client-to-server streaming RPC
// Accepts a stream of Answers and
// returns the Result after all the answers were submitted
func (s *quizer) RecordAnswers(stream pb.Quizer_RecordAnswersServer) error {
	var correctAnswers int32

	startTime := time.Now()

	for {
		answer, err := stream.Recv()

		if err == io.EOF {
			endTime := time.Now()

			mutex.Lock()
			s.results[correctAnswers]++ // add last submission
			mutex.Unlock()

			return stream.SendAndClose(&pb.Result{
				CorrectAnswers: correctAnswers,
				BetterThen:     s.betterThen(correctAnswers),
				ElapsedTime:    float32(endTime.Sub(startTime).Seconds()),
			})
		}

		if err != nil {
			return err
		}

		if s.questions[answer.QuestionId].CorrectAnswer == answer.OptionId {
			correctAnswers++
		}
	}
}

func (s *quizer) load(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	_errCheck(&err, "Failed to load quiz questions: %v")

	err = json.Unmarshal(data, &s.questions)
	_errCheck(&err, "Failed to load quiz questions: %v")

	log.Println("Questions loaded successfully...")
}

func (s *quizer) betterThen(correctAnswers int32) int32 {
	var betterThen, total int32

	for i, k := range s.results {
		total += k
		if int32(i) < correctAnswers {
			betterThen += k
		}
	}

	total -= 1 // remove current submission from total
	if total == 0 {
		return 100
	}

	if betterThen == 0 {
		return 0
	}

	return (betterThen * 100) / total
}

func (s *quizer) Run(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	_errCheck(&err, "Failed to listen: %v")

	grpcServer := grpc.NewServer()
	pb.RegisterQuizerServer(grpcServer, s)

	return grpcServer.Serve(lis)
}

func NewServer(dbFile string) *quizer {
	quizer := &quizer{}
	quizer.load(dbFile)

	// Init with zero, no quiz results yet
	quizer.results = make([]int32, len(quizer.questions)+1)

	return quizer

}

func _errCheck(err *error, msg string) {
	if *err != nil {
		log.Fatalf(msg, *err)
	}
}
