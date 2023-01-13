package repository

import (
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
)

func scanBriefJobStages(rows pgx.Rows) ([]model.BriefJobStage, error) {
	var stages []model.BriefJobStage

	for rows.Next() {
		var stage model.BriefJobStage

		if err := scanBriefJobStage(&stage, rows); err != nil {
			return nil, err
		}

		stages = append(stages, stage)
	}

	return stages, nil
}

func scanBriefJobStage(stage *model.BriefJobStage, row pgx.Row) error {
	if err := row.Scan(
		&stage.Id,
		&stage.Urls,
		&stage.Status,
	); err != nil {
		return err
	}

	return nil
}
