export interface RoleGroup {
    name: string;
    id: string;
    roles: RoleTemplate[]
}
export interface RoleTemplate {
    name: string;
    hourlyRate: number;
    skills: SkillTemplate[]
}
export interface SkillTemplate {
    name: string;
    proficiency: number;
}

export const roleGroups:RoleGroup[] = [
    {
        name: "SaaS Engineering" ,
        id: "saasRoleGroup",
        roles: [
            {
                name: "Full Stack Engineer",
                hourlyRate: 140.0,
                skills: [
                 { name: "Golang", proficiency: 2.0 },
                 { name: "Database", proficiency: 2.0 },
                 { name: "JavaScript", proficiency: 2.0 },
                 { name: "TypeScript", proficiency: 2.0 }
                 ]
             },
             {
                name: "Web Developer",
                hourlyRate: 100.0,
                skills: [
                 { name: "Svelte", proficiency: 2.0 },
                 { name: "CSS", proficiency: 2.0 },
                 { name: "JavaScript", proficiency: 2.0 },
                 { name: "HTML", proficiency: 2.0 }
                 ]
             },
             {
                name: "DevOps Engineer/SRE",
                hourlyRate: 140.0,
                skills: [
                 { name: "Kubernetes", proficiency: 2.0 },
                 { name: "AWS", proficiency: 2.0 },
                 { name: "Keycloak", proficiency: 2.0 }
                 ]
             },
             {
                name: "Engineering Manager",
                hourlyRate: 180.0,
                skills: [
                 { name: "Golang", proficiency: 2.0 },
                 { name: "JavaScript", proficiency: 2.0 },
                 { name: "Leadership", proficiency: 2.0 },
                 { name: "Agile", proficiency: 2.0 }
                 ]
             },
        ]
    },
    {
        name: "Product Management" ,
        id: "pdmRoleGroup",
        roles: [
            {
                name: "Project Manager",
                hourlyRate: 120.0,
                skills: [
                 { name: "Business Analysis", proficiency: 2.0 },
                 { name: "Communications", proficiency: 2.0 }
                 ]
             },
             {
                name: "Product Manager",
                hourlyRate: 100.0,
                skills: [
                 { name: "Business Analysis", proficiency: 2.0 },
                 { name: "Product Development", proficiency: 2.0 }
                 ]
             },
             {
                name: "Product Owner",
                hourlyRate: 110.0,
                skills: [
                 { name: "Scrum", proficiency: 2.0 },
                 { name: "Agile", proficiency: 2.0 }
                 ]
             },
             {
                name: "Product Designer",
                hourlyRate: 130.0,
                skills: [
                 { name: "UI", proficiency: 2.0 },
                 { name: "UX", proficiency: 2.0 },
                 { name: "Product Design", proficiency: 2.0 }
                 ]
             },
        ]
    },
]