package repository

import (
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
)

func scanBriefJobStages(rows pgx.Rows) (stages []model.BriefJobStage, err error) {
	for rows.Next() {
		stage, err := scanBriefJobStage(rows)
		if err != nil {
			return stages, err
		}
		stages = append(stages, stage)
	}

	return
}

func scanBriefJobStage(row pgx.Row) (stage model.BriefJobStage, err error) {
	if err = row.Scan(
		&stage.Id,
		&stage.Urls,
		&stage.Status,
	); err != nil {
		return
	}

	return
}
