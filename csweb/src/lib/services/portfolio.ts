import { GetPortfolioDocument} from "$lib/graphql/generated/sdk";
import type { Portfolio, PortfolioWeekSummary, Schedule, ProjectActivityWeek, ProjectActivity } from "$lib/graphql/generated/sdk";
import { getApolloClient } from "$lib/graphql/gqlclient";
import { deepCopy } from "$lib/utils/helpers";


/**
 * get a list of all projects that are currently scheduled or in-flight
 * @returns a portfolio consisting of all currently scheduled projects
 */
export const getPortfolio = async (): Promise<Portfolio> => {
    const client = getApolloClient();
    return client
        .query({ query: GetPortfolioDocument })
        .then((res) => {
            if (res) {
                return res.data.getPortfolio;
            }
        })
        .catch(err => {
            return err;
        });
};


/**
 * return a single project from the portfolio schedule
 * @param projectID the ID of the project to return
 * @returns a Schedule object for the given project
 */
export const getScheduledProjectFromPortfolio = async(projectID: string): Promise<Schedule> => {
    return getPortfolio()
        .then(res => {
            let schedule = res.schedule.filter(s => s.projectID === projectID)

            if (schedule.length > 0) {
                return schedule[0]
            }

            return Promise.reject("project " + projectID + " not found in portfolio schedule")
        })
        .catch(err => {
            return err
        })
}

/**
 * return a portfolio object with all work stripped out except for a single resource
 * @param resourceID the resource ID to retrieve the work for
 * @returns a portfolio object with only work for the specified resource
 */
export const findScheduledWorkForResource = async(resourceID:string): Promise<Portfolio> => {
    return getPortfolio()
        .then(res => {
            let portfolio:Portfolio = { weekSummary: [] as PortfolioWeekSummary[], schedule: [] as Schedule[] }

            for (let i = 0; i < res.schedule.length; i++) {
                const sch = res.schedule[i]
                let tasks = [] as ProjectActivity[]
                let weeks = [] as ProjectActivityWeek[]
                if(!sch.projectActivityWeeks)
                    continue

                for (let j = 0; j < sch.projectActivityWeeks?.length; j++) {
                    const paw = sch.projectActivityWeeks[j] as ProjectActivityWeek

                    if (!paw.activities) 
                        continue

                    let newWeek = deepCopy(paw)
                    tasks = paw.activities.filter(t => t.resourceID === resourceID)

                    if (tasks.length > 0) {
                        newWeek.activities = tasks

                        weeks.push(newWeek)
                    }
                }

                if(weeks.length > 0) {
                    let newSched = deepCopy(sch)
                    newSched.projectActivityWeeks = weeks

                    portfolio.schedule.push(newSched)
                }
            }

            return portfolio
        })
        .catch(err => {
            return err
        })
}