package app

import (
	"errors"

	"github.com/opensourceways/xihe-server/domain"
)

type CompetitorApplyCmd domain.CompetitorInfo

func (cmd *CompetitorApplyCmd) Validate() error {
	b := cmd.Account != nil &&
		cmd.Name != nil &&
		cmd.Email != nil &&
		cmd.Identity != nil

	if !b {
		return errors.New("invalid cmd")
	}

	return nil
}

func (cmd *CompetitorApplyCmd) toCompetitor() *domain.CompetitorInfo {
	return (*domain.CompetitorInfo)(cmd)
}

type ChallengeCompetitorInfoDTO struct {
	Score        int  `json:"score"`
	IsCompetitor bool `json:"is_competitor"`

	AIQuestionInfo struct {
		RemainingTimes int  `json:"remaining_times"`
		InProgress     bool `json:"in_progress"`
	} `json:"ai_question"`
}

type ChoiceQuestionDTO struct {
	Desc    string   `json:"desc"`
	Options []string `json:"options"`
}

type CompletionQuestionDTO struct {
	Desc string `json:"desc"`
	Info string `json:"info"`
}

type AIQuestionDTO struct {
	Times       int                     `json:"times"`
	Answer      string                  `json:"answer"`
	Choices     []ChoiceQuestionDTO     `json:"choices"`
	Completions []CompletionQuestionDTO `json:"completions"`
}

type AIQuestionAnswerSubmitCmd struct {
	Times  int
	Result []string
	Answer string
}

type ChallengeRankingDTO struct {
	Score      int    `json:"score"`
	Competitor string `json:"competitor"`
}
