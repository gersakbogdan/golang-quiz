package client

import (
	pb "github.com/gersakbogdan/fasttrackquiz/proto"
	"github.com/manifoldco/promptui"
	"log"
)

type questionOption struct {
	Id   int64
	Text string
}

// Display question in a nice way using `promptui`
func prompt(question *pb.Question) *pb.Answer {
	// Prepare options, map to slice (random order)
	items := []questionOption{}
	for _, option := range question.Options {
		items = append(items, questionOption{option.Id, option.Text})
	}

	prompt := promptui.Select{
		Label: question.Text,
		Items: items,
		Templates: &promptui.SelectTemplates{
			Active:   `👉 {{ .Text | cyan }}`,
			Inactive: `   {{ .Text | white }}`,
			Selected: `{{ .Text }}`,
		},
		//HideSelected: true,
		//HideHelp:     true,
	}

	itemId, _, err := prompt.Run()

	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return nil
	}

	return &pb.Answer{
		QuestionId: question.Id,
		OptionId:   int64(question.Options[items[itemId].Id].Id),
	}
}
