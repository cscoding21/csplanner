import { GetPortfolioDocument, GetPortfolioForResourceDocument} from "$lib/graphql/generated/sdk";
import type { Portfolio, PortfolioWeekSummary, Schedule, ProjectActivityWeek, ProjectActivity, Project} from "$lib/graphql/generated/sdk";
import { getApolloClient } from "$lib/graphql/gqlclient";
import { deepCopy } from "$lib/utils/helpers";
import { dateCompare } from "$lib/utils/check";
import { formatDate, formatDateNoYear } from "$lib/utils/format";


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
    const client = getApolloClient();
    return client
        .query({ query: GetPortfolioForResourceDocument, variables: { resourceID } })
        .then((res) => {
            if (res) {
                return res.data.getPortfolioForResource;
            }
        })
        .catch(err => {
            return err;
        });
}


/**
 * extract a single week from an array based on the end date
 * @param paWeeks an array of project activity weeks
 * @param week the week end of the date to extract
 * @returns the project activity week object relevant to the passed in week
 */
export const getWeekActivities = (paWeeks:ProjectActivityWeek[], week:Date):ProjectActivityWeek  => {
    if (paWeeks && paWeeks.length > 0) {
        for (let i = 0; i < paWeeks.length; i++) {
            const paw = paWeeks[i]

            if(dateCompare(new Date(paw.end), week)) {
                return paw as ProjectActivityWeek
            }
        }
    }

    return {} as ProjectActivityWeek
}

/**
 * 
 * @param res the portfolio to transform
 * @returns a portfolio table
 */
export const buildPortfolioTable = (res:Portfolio, startDate:Date|undefined, endDate: Date|undefined):ScheduleTable => {
    let portfolioTable = {startDate: startDate, endDate: endDate, header: [], body: [], footer: []} as ScheduleTable

    if (!res || !res.weekSummary) {
        return portfolioTable
    }

    if(!startDate) {
        startDate = new Date(res.weekSummary[0]?.begin)
    }

    if(!endDate) {
        endDate = new Date(res.weekSummary[res.weekSummary.length-1]?.end)
    }

    portfolioTable.startDate = startDate
    portfolioTable.endDate = endDate

    let headerNames = []
    let footerObjs = []

    for(let i = 0; i < res.weekSummary.length; i++) {
        const thisHeaderWeek = res.weekSummary[i] 
        const thisHeaderWeekEnd = new Date(thisHeaderWeek?.end)
        const thisHeaderWeekStart = new Date(thisHeaderWeek?.begin)

        if(thisHeaderWeekEnd < startDate || thisHeaderWeekStart > endDate)
            continue;

        headerNames.push(formatDateNoYear(thisHeaderWeekEnd))
        footerObjs.push(thisHeaderWeek)
    }

    portfolioTable.header = ["Project", ...headerNames]
    portfolioTable.footer = footerObjs as PortfolioWeekSummary[]

    for(let i = 0; i < res.schedule.length; i++) {
        let schedule = res.schedule[i]
        let row = {weeks: [] as ProjectRowCell[]} as ProjectRow
        let rowHasActivity = false

        row.label = schedule.project.projectBasics.name
        row.project = schedule.project

        for (let j = 0; j < res.weekSummary.length; j++) {
            const thisWeek = res.weekSummary[j] 
            const thisWeekEnd = new Date(thisWeek?.end)
            const thisWeekStart = new Date(thisWeek?.begin)

            if(thisWeekEnd < startDate || thisWeekStart > endDate)
                continue;

            let cell = {active:false, activities:[], end: thisWeek?.end, orgCapacity: thisWeek?.orgCapacity, risks: [] as string[] } as ProjectRowCell

            const paw = getWeekActivities(schedule.projectActivityWeeks as ProjectActivityWeek[], thisWeekEnd)
            cell.orgCapacity = paw.orgCapacity

            if (paw.activities && paw.activities.length > 0) {
                rowHasActivity = true
                cell.active = true
                cell.activities = paw.activities as ProjectActivity[]
            }

            cell.risks = findWeekRisks(paw.activities as ProjectActivity[], undefined)

            row.weeks.push(cell)
        }

        if(rowHasActivity) {
            portfolioTable.body.push(row)
        }
    }

    return portfolioTable
}

export const findWeekRisks = (activities:ProjectActivity[], resourceID: string|undefined):string[] => {
    let out = [] as string[]
    if (!activities || !activities.length) {
        return out
    }

    for(let i = 0; i < activities.length; i++) {
        const res = activities[i].resource

        if(resourceID && resourceID != res?.id)
            continue

        if(res?.status != "inhouse") {
            out.push(res?.name + " is not currently in house and may not be available for this week")
        }
    }

    return out
}

export const flattenPortfolio = (port:Portfolio, startDate:Date, endDate:Date):FlatPortfolioItem[] => {
    let out:FlatPortfolioItem[] = []

    for(let i = 0; i < port.schedule.length; i++) {
        let thisSchedule:Schedule = port.schedule[i]

        if(!thisSchedule.projectActivityWeeks)
            continue

        for(let m = 0; m < thisSchedule.projectActivityWeeks?.length; m++) {
            let thisWeek = thisSchedule.projectActivityWeeks[m]
            let thisWeekStart = new Date(thisWeek?.begin)
            let thisWeekEnd = new Date(thisWeek?.end)

            if (!thisWeek.activities)
                continue
            
            for (let a = 0; a < thisWeek.activities?.length; a++) {
                if(thisWeekEnd < startDate || thisWeekStart > endDate)
                    continue;

                let thisActivity = thisWeek.activities[a]
                let thisItem:FlatPortfolioItem = {} as FlatPortfolioItem

                thisItem.projectID = thisSchedule.projectID
                thisItem.projectName = thisSchedule.project.projectBasics.name
                thisItem.weekEnd = thisWeek.end

                thisItem.milestoneID = thisActivity.milestoneID
                thisItem.milestoneName = thisActivity.milestoneName
                thisItem.taskID = thisActivity.taskID
                thisItem.taskName = thisActivity.taskName
                thisItem.actualizedHours = thisActivity.hoursSpent as number
                thisItem.hourlyRate = thisActivity.resource?.calculated.hourlyCost as number
                thisItem.costForLine = thisItem.actualizedHours * thisItem.hourlyRate
                thisItem.requiredSkill = thisActivity.requiredSkillID

                out.push(thisItem)
            }
        }
    }


    return out
}




export interface ScheduleTable {
    startDate: Date
    endDate: Date
    header: string[]
    body: ProjectRow[]
    footer: PortfolioWeekSummary[]
}
export interface ProjectRow {
    label: string
    project: Project
    weeks: ProjectRowCell[]
}
export interface ProjectRowCell {
    active: boolean
    risks: string[]
    end: Date
    orgCapacity: number
    activities: ProjectActivity[]
}
export interface FlatPortfolioItem {
    weekEnd: Date
    projectID: string
    projectName: string
    resourceID: string
    resourceName: string
    requiredSkill: string
    estimatedHours: number
    actualizedHours: number
    hourlyRate: number
    costForLine: number
    milestoneID: string
    milestoneName: string
    taskID: string
    taskName: string
}