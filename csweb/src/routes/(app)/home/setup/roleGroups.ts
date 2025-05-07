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
                 { id: newID(), skillID: "skill:go", parentID: "role:fse", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:graphql", parentID: "role:fse", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:python", parentID: "role:fse", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:db", parentID: "role:fse", proficiency: 2.0 }
                 ]
             },
             {
                id: "role:webdev",
                description: "",
                name: "Web Developer",
                hourlyRate: 100.0,
                defaultSkills: [
                 { id: newID(), skillID: "skill:svelte", parentID: "role:webdev", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:css", parentID: "role:webdev", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:js", parentID: "role:webdev", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:html", parentID: "role:webdev", proficiency: 2.0 }
                 ]
             },
             {
                id: "role:devops",
                description: "",
                name: "DevOps Engineer/SRE",
                hourlyRate: 140.0,
                defaultSkills: [
                 { id: newID(), skillID: "skill:k8s", parentID: "role:devops", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:aws", parentID: "role:devops", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:keycloak", parentID: "role:devops", proficiency: 2.0 }
                 ]
             },
             {
                id: "role:em",
                description: "",
                name: "Engineering Manager",
                hourlyRate: 180.0,
                defaultSkills: [
                 { id: newID(), skillID: "skill:go", parentID: "role:em", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:js", parentID: "role:em", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:leadership", parentID: "role:em", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:agile", parentID: "role:em", proficiency: 2.0 }
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
                 { id: newID(), skillID: "skill:ba", parentID: "role:pm", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:comms", parentID: "role:pm", proficiency: 2.0 }
                 ]
             },
             {
                id: "role:pdm",
                description: "",
                name: "Product Manager",
                hourlyRate: 100.0,
                defaultSkills: [
                 { id: newID(), skillID: "skill:ba", parentID: "role:pdm", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:pdm", parentID: "role:pdm", proficiency: 2.0 }
                 ]
             },
             {
                id: "role:po",
                description: "",
                name: "Product Owner",
                hourlyRate: 110.0,
                defaultSkills: [
                 { id: newID(), skillID: "skill:scrum", parentID: "role:po", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:agile", parentID: "role:po", proficiency: 2.0 }
                 ]
             },
             {
                id: "role:pd",
                description: "",
                name: "Product Designer",
                hourlyRate: 130.0,
                defaultSkills: [
                 { id: newID(), skillID: "skill:ui", parentID: "role:pd", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:ux", parentID: "role:pd", proficiency: 2.0 },
                 { id: newID(), skillID: "skill:proddes", parentID: "role:pd", proficiency: 2.0 }
                 ]
             },
        ]
    },
]