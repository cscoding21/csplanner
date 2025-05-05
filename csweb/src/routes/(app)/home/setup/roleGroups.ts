import type { UpdateRole } from "$lib/graphql/generated/sdk";
import { newID } from "$lib/utils/id";

export interface RoleGroup {
    name: string;
    id: string;
    roles: UpdateRole[]
}

export const roleGroups:RoleGroup[] = [
    {
        name: "SaaS Engineering" ,
        id: "saasRoleGroup",
        roles: [
            {
                id: "role:fse",
                description: "",
                name: "Full Stack Engineer",
                hourlyRate: 140.0,
                defaultSkills: [
                 { id: newID(), skillID: "Golang", parentID: "role:fse", proficiency: 2.0 },
                 { id: newID(), skillID: "GraphQL", parentID: "role:fse", proficiency: 2.0 },
                 { id: newID(), skillID: "Python", parentID: "role:fse", proficiency: 2.0 },
                 { id: newID(), skillID: "Database", parentID: "role:fse", proficiency: 2.0 }
                 ]
             },
             {
                id: "role:webdev",
                description: "",
                name: "Web Developer",
                hourlyRate: 100.0,
                defaultSkills: [
                 { id: newID(), skillID: "Svelte", parentID: "role:webdev", proficiency: 2.0 },
                 { id: newID(), skillID: "CSS", parentID: "role:webdev", proficiency: 2.0 },
                 { id: newID(), skillID: "JavaScript", parentID: "role:webdev", proficiency: 2.0 },
                 { id: newID(), skillID: "HTML", parentID: "role:webdev", proficiency: 2.0 }
                 ]
             },
             {
                id: "role:devops",
                description: "",
                name: "DevOps Engineer/SRE",
                hourlyRate: 140.0,
                defaultSkills: [
                 { id: newID(), skillID: "Kubernetes", parentID: "role:devops", proficiency: 2.0 },
                 { id: newID(), skillID: "AWS", parentID: "role:devops", proficiency: 2.0 },
                 { id: newID(), skillID: "Keycloak", parentID: "role:devops", proficiency: 2.0 }
                 ]
             },
             {
                id: "role:em",
                description: "",
                name: "Engineering Manager",
                hourlyRate: 180.0,
                defaultSkills: [
                 { id: newID(), skillID: "Golang", parentID: "role:em", proficiency: 2.0 },
                 { id: newID(), skillID: "JavaScript", parentID: "role:em", proficiency: 2.0 },
                 { id: newID(), skillID: "Leadership", parentID: "role:em", proficiency: 2.0 },
                 { id: newID(), skillID: "Agile", parentID: "role:em", proficiency: 2.0 }
                 ]
             },
        ]
    },
    {
        name: "Product Management" ,
        id: "projectManagementRoleGroup",
        roles: [
            {
                id: "role:pm",
                description: "",
                name: "Project Manager",
                hourlyRate: 120.0,
                defaultSkills: [
                 { id: newID(), skillID: "Business Analysis", parentID: "role:pm", proficiency: 2.0 },
                 { id: newID(), skillID: "Communications", parentID: "role:pm", proficiency: 2.0 }
                 ]
             },
             {
                id: "role:pdm",
                description: "",
                name: "Product Manager",
                hourlyRate: 100.0,
                defaultSkills: [
                 { id: newID(), skillID: "Business Analysis", parentID: "role:pdm", proficiency: 2.0 },
                 { id: newID(), skillID: "Product Development", parentID: "role:pdm", proficiency: 2.0 }
                 ]
             },
             {
                id: "role:po",
                description: "",
                name: "Product Owner",
                hourlyRate: 110.0,
                defaultSkills: [
                 { id: newID(), skillID: "Scrum", parentID: "role:po", proficiency: 2.0 },
                 { id: newID(), skillID: "Agile", parentID: "role:po", proficiency: 2.0 }
                 ]
             },
             {
                id: "role:pd",
                description: "",
                name: "Product Designer",
                hourlyRate: 130.0,
                defaultSkills: [
                 { id: newID(), skillID: "UI", parentID: "role:pd", proficiency: 2.0 },
                 { id: newID(), skillID: "UX", parentID: "role:pd", proficiency: 2.0 },
                 { id: newID(), skillID: "Product Design", parentID: "role:pd", proficiency: 2.0 }
                 ]
             },
        ]
    },
]