import { GetDraftPortfolioDocument, GetPortfolioDocument, GetPortfolioForResourceDocument} from "$lib/graphql/generated/sdk";
import type { Portfolio, PortfolioWeekSummary, Schedule, ProjectActivityWeek, ProjectActivity, Project} from "$lib/graphql/generated/sdk";
import { getApolloClient } from "$lib/graphql/gqlclient";
import { dateCompare } from "$lib/utils/check";
import { formatDateNoYear } from "$lib/utils/format";
import { deepCopy, isOneOf } from "$lib/utils/helpers";


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
 * get a list of all projects that are currently scheduled or in-flight
 * @returns a portfolio consisting of all currently scheduled projects
 */
export const getDraftPortfolio = async (additionalID:string): Promise<Portfolio> => {
    const client = getApolloClient();
    return client
        .query({ query: GetDraftPortfolioDocument, variables: { additionalID } })
        .then((res) => {
            if (res) {
                return res.data.getDraftPortfolio;
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
 * take a portfolio object and convert it to a UI friendly shape
 * @param res the portfolio to transform
 * @returns a portfolio table
 */
export const buildPortfolioTable = (res:Portfolio, startDate:Date|undefined, endDate: Date|undefined, highlightProjectID:string|undefined):ScheduleTable => {
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
        row.highlight = (row.project.id === highlightProjectID)

        for (let j = 0; j < res.weekSummary.length; j++) {
            const thisWeek = res.weekSummary[j] 
            const thisWeekEnd = new Date(thisWeek?.end)
            const thisWeekStart = new Date(thisWeek?.begin)

            if(thisWeekEnd < startDate || thisWeekStart > endDate)
                continue;

            let cell = {
                active:false, 
                isPastDue: false,
                activities:[], 
                end: thisWeek?.end,
                orgCapacity: thisWeek?.orgCapacity, 
                risks: [] as string[] 
            } as ProjectRowCell

            const paw = getWeekActivities(schedule.projectActivityWeeks as ProjectActivityWeek[], thisWeekEnd)
            cell.orgCapacity = paw.orgCapacity

            if (paw.activities && paw.activities.length > 0) {
                rowHasActivity = true

                cell.isPastDue = findWeekHasPastDueTasks(paw.activities)
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
    const today = new Date()
    if (!activities || !activities.length) {
        return out
    }

    for(let i = 0; i < activities.length; i++) {
        const act = activities[i]
        const res = act.resource

        if(act.status === "done" || act.status === "removed") {
            continue
        }

        if(resourceID && resourceID != res?.id)
            continue

        if(act.taskEndDate && new Date(act.taskEndDate) < today && (act.status == "new" || act.status == "accepted")) {
            out.push('Task "' + act.taskName + '" is past due')
        }

        if(res?.status != "inhouse") {
            out.push(res?.name + " is not currently in house and may not be available for this week")
        }
    }

    return out
}

/**
 * return whether or not any activities in the array are past due
 * @param activities a list of activities for the week
 * @returns true if any activities are past due
 */
export const findWeekHasPastDueTasks = (activities:ProjectActivity[]):boolean => {
    if (!activities || !activities.length) {
        return false
    }

    const today = new Date()
    for(let i = 0; i < activities.length; i++) {
        const act = activities[i]

        if(act.taskEndDate && new Date(act.taskEndDate) < today && (act.status == "new" || act.status == "accepted")) {
            console.log(act.taskEndDate, act.status)
            return true
        }
    }

    return false
}

export const flattenPortfolio = (port:Portfolio, startDate:Date, endDate:Date):FlatPortfolioItem[] => {
    let out:FlatPortfolioItem[] = []

    if(!port || !port.schedule) {
        return out;
    }

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

/**
 * take the risks and date info and return what the cell color should be
 * @param risks the risks calculated for the schedule
 * @param isPastDue a boolean for whether or not the task is past due
 * @returns a color specification
 */
export const getCellColor = (risks:string[], isPastDue:boolean):"red"|"yellow"|"green"|undefined => {
    if(isPastDue) {
        return "red"
    } else if (risks && risks.length > 0) {
        return "yellow"
    }

    return "green"
}


/**
 * 
 * @param port 
 * @returns 
 */
export const findOpenPastDueTasksInPortfolio = (port:Portfolio):ProjectActivity[] => {
    let out:ProjectActivity[] = []

    if(!port || !port.schedule || port.schedule.length === 0) {
        return out
    }

    const today = new Date()

    for (let i = 0; i < port.schedule.length; i++) {
        const sch = port.schedule[i]

        if (!sch.projectActivityWeeks) {
            continue
        }

        for (let j = 0; j < sch.projectActivityWeeks?.length; j++) {
            const paw = sch.projectActivityWeeks[j]

            if(!paw.activities || paw.activities.length === 0) {
                continue
            }

            for (let k = 0; k < paw.activities.length; k++) {
                let act = paw.activities[k]

                if(isOneOf(act.status, ["new", "accepted"]) && new Date(act.taskEndDate) < today) {
                    let outAct = deepCopy(act)
                    outAct.project = deepCopy(sch.project)
                    //console.log("sch", sch)
                    out.push(outAct)
                }
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
    highlight: Boolean
    project: Project
    weeks: ProjectRowCell[]
}
export interface ProjectRowCell {
    active: boolean
    isPastDue: boolean
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