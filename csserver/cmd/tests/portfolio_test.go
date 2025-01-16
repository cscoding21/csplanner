package tests

import (
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
	"csserver/internal/services/schedule"
	"fmt"
	"testing"
	"time"
)

func init() {
	config.InitConfig()
}

func TestGetSchedule(t *testing.T) {
	scheduleService := factory.GetScheduleService()
	projectService := factory.GetProjectService()
	resourceService := factory.GetResourceService()

	ctx := getTestContext()
	pid := "project:1"
	startDate := time.Now()

	rm, err := resourceService.GetResourceMap(ctx, true)
	if err != nil {
		t.Error(err)
	}

	proj, err := projectService.GetProjectByID(ctx, pid)
	if err != nil {
		t.Error(err)
	}

	result, err := scheduleService.CalculateProjectSchedule(ctx, proj, startDate, rm)
	if err != nil {
		t.Error(err)
	}

	drawResult(result)
}

func TestGetPortfolio(t *testing.T) {
	portfolioService := factory.GetPortfolioService()
	ctx := getTestContext()

	result, err := portfolioService.GetCurrentPortfolio(ctx)
	if err != nil {
		t.Error(err)
	}

	for _, s := range result.Schedule {
		fmt.Println("###########################################################")
		fmt.Println("-----------")
		drawResult(s)
		fmt.Println("-----------")
		fmt.Println("###########################################################")
	}
}

func drawResult(result schedule.Schedule) {
	fmt.Printf("%s (%s)\n", result.ProjectName, result.ProjectID)
	fmt.Printf("%v - %v\n", result.Begin.Format("2006-01-02"), result.End.Format("2006-01-02"))
	fmt.Printf("Hash - %v\n", result.Hash)

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
