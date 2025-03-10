package tests

import (
	"csserver/internal/appserv/factory"
	"csserver/internal/calendar"
	"csserver/internal/config"
	"csserver/internal/services/portfolio"
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

	ram, err := scheduleService.GetInitialResourceAllocationMap(rm)
	if err != nil {
		t.Error(err)
	}

	proj, err := projectService.GetProjectByID(ctx, pid)
	if err != nil {
		t.Error(err)
		return
	}

	result, err := scheduleService.CalculateProjectSchedule(ctx, &proj.Data, startDate, ram)
	if err != nil {
		t.Error(err)
	}

	drawScheduleResult(result)
}

func TestGetPortfolio(t *testing.T) {
	ctx := getTestContext()

	portfolioService := factory.GetPortfolioService()

	result, err := portfolioService.GetUnbalancedPortfolio(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	if result == nil {
		t.Error("no results returned")
		return
	}

	for _, s := range result.Schedule {
		fmt.Println("###########################################################")
		fmt.Println("-----------")
		drawScheduleResult(s)
		fmt.Println("-----------")
		fmt.Println("###########################################################")
	}
}

func TestBalancePortfolio(t *testing.T) {
	ctx := getTestContext()
	ps := factory.GetPortfolioService()

	port, err := ps.GetUnbalancedPortfolio(ctx)
	if err != nil {
		t.Error(err)
	}

	err = ps.BalancePortfolio(ctx, port)
	if err != nil {
		t.Error(err)
	}

	drawPortfolio(*port)

	drawResourceSummary(port.Schedule)
}

func TestGetPortfolioForResource(t *testing.T) {
	ctx := getTestContext()
	ps := factory.GetPortfolioService()

	port, err := ps.GetBalancedPortfolio(ctx)
	if err != nil {
		t.Error(err)
	}

	resPort, err := ps.GetPortfolioForResource(ctx, port, "resource:barret")
	if err != nil {
		t.Error(err)
	}

	drawPortfolio(*resPort)
}

func drawPortfolio(port portfolio.Portfolio) {

	fmt.Printf("------ Portfolio info ------\n")
	fmt.Printf("Project count: %v\n", len(port.ProjectMap))
	fmt.Printf("Resource count: %v\n", len(port.ResourceMap))
	if port.Validator != nil {
		fmt.Printf("Iterations: %v\n", port.Validator.Iterations)
	}
	for _, s := range port.Schedule {
		drawScheduleResult(s)

		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++")
	}
}

func drawScheduleResult(result schedule.Schedule) {
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

func drawResourceSummary(scheduleList []schedule.Schedule) {
	type rss struct {
		week      calendar.CSWeek
		hours     int
		taskCount int
		projectID string
	}
	resMap := make(map[string][]rss)

	for _, r := range scheduleList {
		for _, w := range r.ProjectActivityWeeks {
			for _, a := range w.Activities {
				_, ok := resMap[a.ResourceID]
				if ok {
					resMap[a.ResourceID] = append(resMap[a.ResourceID], rss{week: calendar.GetWeek(w.Begin), hours: a.HoursSpent, taskCount: 1, projectID: a.ProjectID})
				} else {
					resMap[a.ResourceID] = []rss{{week: calendar.GetWeek(w.Begin), hours: a.HoursSpent, taskCount: 1, projectID: a.ProjectID}}
				}
			}
		}
	}

	for k, v := range resMap {
		fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
		totalHours := 0
		totalTasks := 0

		fmt.Printf("Resource: %s\n", k)
		for _, a := range v {
			totalHours += a.hours
			totalTasks++
			fmt.Printf(" - %s | WE: %v | hours: %v\n", a.projectID, a.week.End, a.hours)
		}

		fmt.Println("--------")
		fmt.Printf("Total tasks %v | Total hours: %v\n", totalTasks, totalHours)
	}
}
