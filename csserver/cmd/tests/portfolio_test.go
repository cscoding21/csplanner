package tests

import (
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
	"csserver/internal/services/portfolio"
	"fmt"
	"testing"
	"time"
)

func init() {
	config.InitConfig()
}

func TestPortfolioGetResourceUtilizationTable(t *testing.T) {
	service := factory.GetPortfolioService()
	ctx := getTestContext()

	table, err := service.GetResourceUtilizationTable(ctx)
	if err != nil {
		t.Error(err)
	}

	t.Log(table)
}

func TestGetSchedule(t *testing.T) {
	portfolioService := factory.GetPortfolioService()
	projectService := factory.GetProjectService()

	ctx := getTestContext()
	pid := "project:1"
	startDate := time.Now()

	proj, err := projectService.GetProjectByID(ctx, pid)
	if err != nil {
		t.Error(err)
	}

	result, err := portfolioService.ScheduleProject(ctx, proj, startDate)
	if err != nil {
		t.Error(err)
	}

	drawResult(result)
}

func drawResult(result portfolio.ProjectSchedule) {
	fmt.Printf("%s (%s)\n", result.ProjectName, result.ProjectID)
	fmt.Printf("%v - %v\n", result.Begin.Format("2006-01-02"), result.End.Format("2006-01-02"))

	if len(result.Exceptions) > 0 {
		fmt.Println("************** Exceptions **************")
		for _, e := range result.Exceptions {
			fmt.Printf(" - %s: %s\n", e.Scope, e.Message)
		}
	}

	fmt.Println("-------------- Weeks --------------")
	for _, paw := range result.ProjectActivityWeeks {
		fmt.Printf("%s - %s\n", paw.Begin.Format("2006-01-02"), paw.End.Format("2006-01-02"))
		for _, a := range paw.Activities {
			fmt.Printf(" - %v hours - %s: %s\n", a.HoursSpent, a.ResourceName, a.TaskName)
		}
	}

}
