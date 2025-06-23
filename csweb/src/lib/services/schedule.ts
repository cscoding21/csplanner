import type { PortfolioWeekSummary, ProjectActivity, ProjectActivityWeek, Resource, Schedule } from "$lib/graphql/generated/sdk"
import { formatDateNoYear } from "$lib/utils/format"
import { findWeekHasPastDueTasks, findWeekRisks } from "./portfolio"


export interface ProjectScheduleTable {
    startDate: Date
    endDate: Date
    header: string[]
    body: ProjectScheduleRow[]
    footer: PortfolioWeekSummary[]
}
export interface ProjectScheduleRow {
    label: string
    resource?: Resource
    weeks: ProjectScheduleRowCell[]
}
export interface ProjectScheduleRowCell {
    active: boolean
    risks: string[]
    isPastDue: boolean
    showTasks: boolean
    end: Date
    orgCapacity: number
    activities: ProjectActivity[]
}

export const getScheduleTableByResource = (r: Schedule):ProjectScheduleTable => {
        let table = {header: [] as string[], body: [] as ProjectScheduleRow[]} as ProjectScheduleTable
        let head = []
        let resourceSet = new Set<Resource>()

        head.push("Resources")

        if (!r.projectActivityWeeks) {
            return {} as ProjectScheduleTable
        }

        for (let i = 0; i < r.projectActivityWeeks.length; i++) {
            const week = r.projectActivityWeeks[i]
            head.push(formatDateNoYear(week.end))

            if (week.activities) {
                for (let j = 0; j < week.activities?.length; j++) {
                    const act = week.activities[j]
                    resourceSet.add(act.resource as Resource)
                }
            }
        }

        table.header = head as string[]
        table.body = [] as ProjectScheduleRow[]
        const resourceArray = [...resourceSet]

        for (let x = 0; x < resourceArray.length; x++) {
            const thisResource = resourceArray[x] as Resource
            let row:ProjectScheduleRow = {label: thisResource.name, resource: thisResource, weeks: [] as ProjectScheduleRowCell[]}

            for (let i = 0; i < r.projectActivityWeeks.length; i++) {
                const week = r.projectActivityWeeks[i]
                let weeksActivities:ProjectScheduleRowCell = { 
                    showTasks: true,
                    isPastDue: false,
                    end: week.end,
                    active: true,
                    orgCapacity: 0,
                    activities: [] as ProjectActivity[], 
                    risks: [] as string[]
                }

                if (week.activities) {
                    for(let j = 0; j < week.activities.length; j++) {
                        const activity = week.activities[j]
                    
                        if (thisResource.name == activity.resource?.name) {
                            weeksActivities.activities.push(activity)
                            weeksActivities.risks = findWeekRisks(week.activities, activity.resourceID)
                        }
                    }

                    weeksActivities.isPastDue = findWeekHasPastDueTasks(weeksActivities.activities as ProjectActivity[])
                }
            
                row.weeks.push(weeksActivities)
            }
            
            table.body.push(row)
        }

        return table
    }


    export const getScheduleTableByTask = (r: Schedule):ProjectScheduleTable|undefined => {
        if(!r || !r.projectActivityWeeks) {
            return undefined
        }

        let table = {header: [] as string[], body: [] as ProjectScheduleRow[]} as ProjectScheduleTable
        let taskSet = new Set<string>()

        table.header.push("Task")

        if (!r.projectActivityWeeks) {
            return {} as ProjectScheduleTable
        }

        for (let i = 0; i < r.projectActivityWeeks.length; i++) {
            const week = r.projectActivityWeeks[i]
            table.header.push(formatDateNoYear(week.end))

            if (week.activities) {
                for (let j = 0; j < week.activities?.length; j++) {
                    const act = week.activities[j]
                    taskSet.add(act.taskName as string)
                }
            }
        }

        const taskArray = [...taskSet]

        for (let x = 0; x < taskArray.length; x++) {
            const thisTask = taskArray[x] as string
            let row:ProjectScheduleRow = {label: thisTask, resource: undefined, weeks: [] as ProjectScheduleRowCell[]}

            for (let i = 0; i < r.projectActivityWeeks.length; i++) {
                const week = r.projectActivityWeeks[i]
                let weeksActivities:ProjectScheduleRowCell = { 
                    showTasks: false,
                    end: week.end, 
                    isPastDue: false,
                    active: true,
                    orgCapacity: 0,
                    activities: [] as ProjectActivity[],
                    risks: [] as string[]
                }

                if (week.activities) {
                    for(let j = 0; j < week.activities.length; j++) {
                        const activity = week.activities[j]
                    
                        if (thisTask == activity.taskName) {
                            weeksActivities.activities.push(activity)
                        }
                    }

                    weeksActivities.risks = findWeekRisks(weeksActivities.activities, undefined)
                    weeksActivities.isPastDue = findWeekHasPastDueTasks(weeksActivities.activities as ProjectActivity[])
                }
            
                row.weeks.push(weeksActivities)
            }
            
            table.body.push(row)
        }

        return table
    }

