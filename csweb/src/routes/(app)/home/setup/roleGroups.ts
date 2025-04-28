import type { UpdateRole } from "$lib/graphql/generated/sdk";

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
                id: "",
                description: "",
                name: "Full Stack Engineer",
                hourlyRate: 140.0,
                defaultSkills: [
                 { id: "", resourceID: "", proficiency: 2.0 },
                 { id: "", resourceID: "", proficiency: 2.0 },
                 { id: "", resourceID: "", proficiency: 2.0 },
                 { id: "", resourceID: "", proficiency: 2.0 }
                 ]
             },
             {
                id: "",
                description: "",
                name: "Web Developer",
                hourlyRate: 100.0,
                defaultSkills: [
                 { id: "", resourceID: "Svelte", proficiency: 2.0 },
                 { id: "", resourceID: "CSS", proficiency: 2.0 },
                 { id: "", resourceID: "JavaScript", proficiency: 2.0 },
                 { id: "", resourceID: "HTML", proficiency: 2.0 }
                 ]
             },
             {
                id: "",
                description: "",
                name: "DevOps Engineer/SRE",
                hourlyRate: 140.0,
                defaultSkills: [
                 { id: "", resourceID: "Kubernetes", proficiency: 2.0 },
                 { id: "", resourceID: "AWS", proficiency: 2.0 },
                 { id: "", resourceID: "Keycloak", proficiency: 2.0 }
                 ]
             },
             {
                id: "",
                description: "",
                name: "Engineering Manager",
                hourlyRate: 180.0,
                defaultSkills: [
                 { id: "", resourceID: "Golang", proficiency: 2.0 },
                 { id: "", resourceID: "JavaScript", proficiency: 2.0 },
                 { id: "", resourceID: "Leadership", proficiency: 2.0 },
                 { id: "", resourceID: "Agile", proficiency: 2.0 }
                 ]
             },
        ]
    },
    {
        name: "Product Management" ,
        id: "pdmRoleGroup",
        roles: [
            {
                id: "",
                description: "",
                name: "Project Manager",
                hourlyRate: 120.0,
                defaultSkills: [
                 { id: "", resourceID: "Business Analysis", proficiency: 2.0 },
                 { id: "", resourceID: "Communications", proficiency: 2.0 }
                 ]
             },
             {
                id: "",
                description: "",
                name: "Product Manager",
                hourlyRate: 100.0,
                defaultSkills: [
                 { id: "", resourceID: "Business Analysis", proficiency: 2.0 },
                 { id: "", resourceID: "Product Development", proficiency: 2.0 }
                 ]
             },
             {
                id: "",
                description: "",
                name: "Product Owner",
                hourlyRate: 110.0,
                defaultSkills: [
                 { id: "", resourceID: "Scrum", proficiency: 2.0 },
                 { id: "", resourceID: "Agile", proficiency: 2.0 }
                 ]
             },
             {
                id: "",
                description: "",
                name: "Product Designer",
                hourlyRate: 130.0,
                defaultSkills: [
                 { id: "", resourceID: "UI", proficiency: 2.0 },
                 { id: "", resourceID: "UX", proficiency: 2.0 },
                 { id: "", resourceID: "Product Design", proficiency: 2.0 }
                 ]
             },
        ]
    },
]