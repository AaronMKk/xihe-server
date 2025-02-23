package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/opensourceways/xihe-server/infrastructure/repositories"
)

func (col finetuneCol) toFinetuneDoc(do *repositories.UserFinetuneDO) (bson.M, error) {
	doc := finetuneItem{
		Id:              do.Id,
		Name:            do.Name,
		Task:            do.Task,
		Model:           do.Model,
		CreatedAt:       do.CreatedAt,
		Hyperparameters: do.Hyperparameters,
	}

	return genDoc(&doc)
}

func (col finetuneCol) toFinetuneSummaryDO(doc *finetuneItem, do *repositories.FinetuneSummaryDO) {
	detail := &doc.JobDetail

	*do = repositories.FinetuneSummaryDO{
		Id:        doc.Id,
		Name:      doc.Name,
		CreatedAt: doc.CreatedAt,
		FinetuneJobDetailDO: repositories.FinetuneJobDetailDO{
			Error:    detail.Error,
			Status:   detail.Status,
			Duration: detail.Duration,
		},
	}
}

func (col finetuneCol) toFinetuneDetailDO(doc *finetuneItem, do *repositories.FinetuneDetailDO) {
	job := &doc.Job
	detail := &doc.JobDetail

	*do = repositories.FinetuneDetailDO{
		Id:        doc.Id,
		CreatedAt: doc.CreatedAt,
		FinetuneConfigDO: repositories.FinetuneConfigDO{
			Name:            doc.Name,
			Task:            doc.Task,
			Model:           doc.Model,
			Hyperparameters: doc.Hyperparameters,
		},
		Job: repositories.FinetuneJobInfoDO{
			JobId:    job.JobId,
			Endpoint: job.Endpoint,
		},
		JobDetail: repositories.FinetuneJobDetailDO{
			Error:    detail.Error,
			Status:   detail.Status,
			Duration: detail.Duration,
		},
	}
}
