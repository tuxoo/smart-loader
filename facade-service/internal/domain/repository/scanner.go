package repository

import (
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
)

func scanUser(row pgx.Row) (user model.User, err error) {
	if err = row.Scan(&user.Id, &user.Name, &user.LoginEmail, &user.RegisteredAt, &user.VisitedAt); err != nil {
		return user, err
	}
	return
}

func scanJobs(rows pgx.Rows) (jobs []model.Job, err error) {
	for rows.Next() {
		job, err := scanJob(rows)
		if err != nil {
			return jobs, err
		}
		jobs = append(jobs, job)
	}

	return
}

func scanJob(row pgx.Row) (job model.Job, err error) {
	if err = row.Scan(
		&job.Id,
		&job.Size,
		&job.Status,
		&job.CreatedAt,
	); err != nil {
		return
	}
	return
}

func scanDownloads(rows pgx.Rows) (downloads []model.Download, err error) {
	for rows.Next() {
		download, err := scanDownload(rows)
		if err != nil {
			return downloads, err
		}
		downloads = append(downloads, download)
	}

	return
}

func scanDownload(row pgx.Row) (download model.Download, err error) {
	if err = row.Scan(
		&download.Id,
		&download.Hash,
		&download.DownloadedAt,
		&download.Size,
	); err != nil {
		return
	}
	return
}
