package report

import (
	"fmt"
	"time"
)

type Report struct {
	ID        string
	Type      string
	Content   string
	Status    string
	CreatedAt time.Time
}

func GenerateReport(reportID string, reportType string) *Report {
	report := &Report{
		ID:        reportID,
		Type:      reportType,
		Status:    "in progress",
		CreatedAt: time.Now(),
	}

	time.Sleep(5 * time.Second)
	report.Content = fmt.Sprintf("Este es el contenido del reporte tipo '%s'", reportType)
	report.Status = "completed"

	return report
}

func (r *Report) GetReportURL() string {
	return fmt.Sprintf("https://path/to/your/report/%s", r.ID)
}
