// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package idl

import (
	"time"
)

type Activity struct {
	ID            string         `json:"id"`
	Type          string         `json:"type"`
	Summary       string         `json:"summary"`
	Detail        *string        `json:"detail,omitempty"`
	Context       string         `json:"context"`
	TargetID      *string        `json:"targetID,omitempty"`
	Resource      *Resource      `json:"resource,omitempty"`
	ControlFields *ControlFields `json:"controlFields"`
}

type ActivityResults struct {
	Paging  *Pagination `json:"paging"`
	Filters *Filters    `json:"filters"`
	Results []*Activity `json:"results,omitempty"`
}

type Artifact struct {
	ID            string         `json:"id"`
	FileType      string         `json:"fileType"`
	FileName      string         `json:"fileName"`
	URL           string         `json:"url"`
	ControlFields *ControlFields `json:"controlFields,omitempty"`
}

type Comment struct {
	ID            string         `json:"id"`
	ProjectID     string         `json:"projectId"`
	Text          string         `json:"text"`
	Resource      *Resource      `json:"resource"`
	Replies       []*Comment     `json:"replies,omitempty"`
	ControlFields *ControlFields `json:"controlFields"`
	Likes         []string       `json:"likes,omitempty"`
	Loves         []string       `json:"loves,omitempty"`
	Dislikes      []string       `json:"dislikes,omitempty"`
	LaughsAt      []string       `json:"laughsAt,omitempty"`
	Acknowledges  []string       `json:"acknowledges,omitempty"`
	IsEdited      bool           `json:"isEdited"`
}

type ControlFields struct {
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	CreatedBy    string     `json:"createdBy"`
	CreateByUser *User      `json:"createByUser,omitempty"`
	UpdatedBy    string     `json:"updatedBy"`
	UpdateByUser *User      `json:"updateByUser,omitempty"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
	DeletedBy    *string    `json:"deletedBy,omitempty"`
	DeleteByUser *User      `json:"deleteByUser,omitempty"`
}

type CreateOrganizationResult struct {
	Status       *Status       `json:"status"`
	Organization *Organization `json:"organization,omitempty"`
}

type CreateProjectCommentResult struct {
	Status  *Status  `json:"status,omitempty"`
	Comment *Comment `json:"comment,omitempty"`
}

type CreateProjectResult struct {
	Status  *Status  `json:"status,omitempty"`
	Project *Project `json:"project,omitempty"`
}

type CreateResourceResult struct {
	Status   *Status   `json:"status"`
	Resource *Resource `json:"resource,omitempty"`
}

type CreateUserResult struct {
	Status *Status `json:"status,omitempty"`
	User   *User   `json:"user,omitempty"`
}

type Filter struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Operation string `json:"operation"`
}

type Filters struct {
	Filters []*Filter `json:"filters,omitempty"`
}

type ImplementationTemplate struct {
	ID            string                         `json:"id"`
	Name          string                         `json:"name"`
	ControlFields *ControlFields                 `json:"controlFields"`
	Phases        []*ImplementationTemplatePhase `json:"phases"`
}

type ImplementationTemplatePhase struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Order       int    `json:"order"`
	Description string `json:"description"`
}

type InputFilter struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Operation string `json:"operation"`
}

type InputFilters struct {
	Filters []*InputFilter `json:"filters,omitempty"`
}

type InputPagination struct {
	ResultsPerPage *int `json:"resultsPerPage,omitempty"`
	PageNumber     *int `json:"pageNumber,omitempty"`
}

type List struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Values        []*ListItem    `json:"values"`
	ControlFields *ControlFields `json:"controlFields,omitempty"`
}

type ListItem struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type ListResults struct {
	Paging  *Pagination `json:"paging,omitempty"`
	Filters *Filters    `json:"filters"`
	Results []*List     `json:"results,omitempty"`
}

type LoginResult struct {
	Token    *string   `json:"token,omitempty"`
	Status   *Status   `json:"status,omitempty"`
	User     *User     `json:"user,omitempty"`
	Resource *Resource `json:"resource,omitempty"`
}

type Mutation struct {
}

type Notification struct {
	ID                    string         `json:"id"`
	UserID                string         `json:"userId"`
	UserName              string         `json:"userName"`
	RecipientIsBot        bool           `json:"recipientIsBot"`
	Type                  int            `json:"type"`
	ContextID             string         `json:"contextId"`
	Text                  *string        `json:"text,omitempty"`
	IsRead                bool           `json:"isRead"`
	InitiatorName         string         `json:"initiatorName"`
	InitiatorID           string         `json:"initiatorId"`
	InitiatorProfileImage *string        `json:"initiatorProfileImage,omitempty"`
	ControlFields         *ControlFields `json:"controlFields"`
}

type NotificationResults struct {
	Paging  *Pagination     `json:"paging,omitempty"`
	Filters *Filters        `json:"filters"`
	Results []*Notification `json:"results"`
}

type Organization struct {
	ID            *string               `json:"id,omitempty"`
	Name          string                `json:"name"`
	Defaults      *OrganizationDefaults `json:"defaults"`
	ControlFields *ControlFields        `json:"controlFields"`
}

type OrganizationDefaults struct {
	DiscountRate float64 `json:"discountRate"`
	HoursPerWeek int     `json:"hoursPerWeek"`
}

type PageAndFilter struct {
	Paging  *InputPagination `json:"paging,omitempty"`
	Filters *InputFilters    `json:"filters,omitempty"`
}

type Pagination struct {
	TotalResults   *int    `json:"totalResults,omitempty"`
	ResultsPerPage *int    `json:"resultsPerPage,omitempty"`
	PageNumber     *int    `json:"pageNumber,omitempty"`
	After          *string `json:"after,omitempty"`
}

type PortfolioSnapshot struct {
	Projects []*SnapshotProject `json:"projects"`
}

type Project struct {
	ID                *string             `json:"id,omitempty"`
	ProjectBasics     *ProjectBasics      `json:"projectBasics"`
	ProjectValue      *ProjectValue       `json:"projectValue"`
	ProjectCost       *ProjectCost        `json:"projectCost"`
	ProjectDaci       *ProjectDaci        `json:"projectDaci"`
	ControlFields     *ControlFields      `json:"controlFields"`
	ProjectFeatures   []*ProjectFeature   `json:"projectFeatures,omitempty"`
	ProjectMilestones []*ProjectMilestone `json:"projectMilestones,omitempty"`
}

type ProjectBasics struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	OwnerID     *string    `json:"ownerID,omitempty"`
}

type ProjectCost struct {
	Ongoing      *float64 `json:"ongoing,omitempty"`
	BlendedRate  *float64 `json:"blendedRate,omitempty"`
	InitialCost  *float64 `json:"initialCost,omitempty"`
	HourEstimate int      `json:"hourEstimate"`
}

type ProjectDaci struct {
	Driver      []*string `json:"driver,omitempty"`
	Approver    []*string `json:"approver,omitempty"`
	Contributor []*string `json:"contributor,omitempty"`
	Informed    []*string `json:"informed,omitempty"`
}

type ProjectFeature struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Priority    float64 `json:"priority"`
	Status      string  `json:"status"`
}

type ProjectFilters struct {
	Status *string `json:"status,omitempty"`
}

type ProjectMilestone struct {
	ID        string                  `json:"id"`
	StartDate *time.Time              `json:"startDate,omitempty"`
	EndDate   *time.Time              `json:"endDate,omitempty"`
	Phase     *ProjectMilestonePhase  `json:"phase"`
	Tasks     []*ProjectMilestoneTask `json:"tasks"`
}

type ProjectMilestonePhase struct {
	ID          string `json:"id"`
	Order       int    `json:"order"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectMilestoneTask struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	HourEstimate     int         `json:"hourEstimate"`
	Description      *string     `json:"description,omitempty"`
	RequiredSkillIDs []string    `json:"requiredSkillIDs,omitempty"`
	Skills           []*Skill    `json:"skills,omitempty"`
	Resources        []*Resource `json:"resources,omitempty"`
	ResourceIDs      []string    `json:"resourceIDs,omitempty"`
	Status           string      `json:"status"`
	StartDate        *time.Time  `json:"startDate,omitempty"`
	EndDate          *time.Time  `json:"endDate,omitempty"`
}

type ProjectResults struct {
	Paging  *Pagination `json:"paging,omitempty"`
	Filters *Filters    `json:"filters"`
	Results []*Project  `json:"results,omitempty"`
}

type ProjectValue struct {
	FundingSource        *string  `json:"fundingSource,omitempty"`
	DiscountRate         *float64 `json:"discountRate,omitempty"`
	YearOneValue         *float64 `json:"yearOneValue,omitempty"`
	YearTwoValue         *float64 `json:"yearTwoValue,omitempty"`
	YearThreeValue       *float64 `json:"yearThreeValue,omitempty"`
	YearFourValue        *float64 `json:"yearFourValue,omitempty"`
	YearFiveValue        *float64 `json:"yearFiveValue,omitempty"`
	NetPresentValue      *float64 `json:"netPresentValue,omitempty"`
	InternalRateOfReturn *float64 `json:"internalRateOfReturn,omitempty"`
}

type Query struct {
}

type RelateArtifact struct {
	ArtifactID string `json:"artifactID"`
	LinkID     string `json:"linkID"`
}

type Resource struct {
	ID            *string        `json:"id,omitempty"`
	Name          string         `json:"name"`
	Role          string         `json:"role"`
	User          *User          `json:"user,omitempty"`
	ProfileImage  *string        `json:"profileImage,omitempty"`
	Skills        []*Skill       `json:"skills,omitempty"`
	IsBot         bool           `json:"isBot"`
	ControlFields *ControlFields `json:"controlFields"`
}

type ResourceResults struct {
	Paging  *Pagination `json:"paging,omitempty"`
	Filters *Filters    `json:"filters"`
	Results []*Resource `json:"results,omitempty"`
}

type ResourceSnapshot struct {
	ScheduledResources []*ScheduledResource `json:"scheduledResources"`
}

type ScheduledResource struct {
	ProjectID          string     `json:"projectId"`
	ProjectName        string     `json:"projectName"`
	ProjectStatus      string     `json:"projectStatus"`
	MilestoneName      string     `json:"milestoneName"`
	MilestoneStartDate *time.Time `json:"milestoneStartDate,omitempty"`
	MilestoneEndDate   *time.Time `json:"milestoneEndDate,omitempty"`
	TaskName           string     `json:"taskName"`
	TaskStartDate      *time.Time `json:"taskStartDate,omitempty"`
	TaskEndDate        *time.Time `json:"taskEndDate,omitempty"`
	TaskHourEstimate   int        `json:"taskHourEstimate"`
	ResourceID         string     `json:"resourceId"`
	ResourceName       string     `json:"resourceName"`
}

type Skill struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Proficiency *float64 `json:"proficiency,omitempty"`
}

type SnapshotProject struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Npv       float64    `json:"npv"`
	Status    string     `json:"status"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate   *time.Time `json:"endDate,omitempty"`
}

type Status struct {
	Success bool     `json:"success"`
	Message []string `json:"message"`
}

// `Subscription` is where all the subscriptions your clients can
// request. You can use Schema Directives like normal to restrict
// access.
type Subscription struct {
}

type TimeResponse struct {
	UnixTime  int    `json:"unixTime"`
	TimeStamp string `json:"timeStamp"`
}

type UpdateComment struct {
	ProjectID string  `json:"projectId"`
	ID        *string `json:"id,omitempty"`
	Text      string  `json:"text"`
}

type UpdateCommentEmote struct {
	CommentID string `json:"commentID"`
	EmoteType string `json:"emoteType"`
}

type UpdateCommentReply struct {
	ParentCommentID string `json:"parentCommentID"`
	Text            string `json:"text"`
}

type UpdateLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateOrganization struct {
	ID       *string                     `json:"id,omitempty"`
	Name     string                      `json:"name"`
	Defaults *UpdateOrganizationDefaults `json:"defaults"`
}

type UpdateOrganizationDefaults struct {
	DiscountRate float64 `json:"discountRate"`
	HoursPerWeek int     `json:"hoursPerWeek"`
}

type UpdateProject struct {
	ID                *string                   `json:"id,omitempty"`
	ProjectBasics     *UpdateProjectBasics      `json:"projectBasics,omitempty"`
	ProjectValue      *UpdateProjectValue       `json:"projectValue,omitempty"`
	ProjectCost       *UpdateProjectCost        `json:"projectCost,omitempty"`
	ProjectDaci       *UpdateProjectDaci        `json:"projectDaci,omitempty"`
	ProjectFeatures   []*UpdateProjectFeature   `json:"projectFeatures,omitempty"`
	ProjectMilestones []*UpdateProjectMilestone `json:"projectMilestones,omitempty"`
}

type UpdateProjectBasics struct {
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Status      *string    `json:"status,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	OwnerID     *string    `json:"ownerID,omitempty"`
}

type UpdateProjectCost struct {
	Ongoing     *float64 `json:"ongoing,omitempty"`
	BlendedRate *float64 `json:"blendedRate,omitempty"`
}

type UpdateProjectDaci struct {
	Driver      []*string `json:"driver,omitempty"`
	Approver    []*string `json:"approver,omitempty"`
	Contributor []*string `json:"contributor,omitempty"`
	Informed    []*string `json:"informed,omitempty"`
}

type UpdateProjectFeature struct {
	ProjectID   string  `json:"projectID"`
	ID          *string `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	Priority    int     `json:"priority"`
}

type UpdateProjectMilestone struct {
	ID        string                        `json:"id"`
	StartDate *time.Time                    `json:"startDate,omitempty"`
	Phase     *UpdateProjectMilestonePhase  `json:"phase"`
	Tasks     []*UpdateProjectMilestoneTask `json:"tasks,omitempty"`
}

type UpdateProjectMilestonePhase struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Order       int    `json:"order"`
	Description string `json:"description"`
}

type UpdateProjectMilestoneTask struct {
	ID               *string    `json:"id,omitempty"`
	Name             string     `json:"name"`
	Description      *string    `json:"description,omitempty"`
	HourEstimate     int        `json:"hourEstimate"`
	ResourceIDs      []string   `json:"resourceIDs,omitempty"`
	RequiredSkillIDs []string   `json:"requiredSkillIDs,omitempty"`
	StartDate        *time.Time `json:"startDate,omitempty"`
	EndDate          *time.Time `json:"endDate,omitempty"`
	Status           string     `json:"status"`
	ProjectID        string     `json:"projectID"`
	MilestoneID      string     `json:"milestoneID"`
}

type UpdateProjectMilestoneTemplate struct {
	ProjectID  string `json:"projectId"`
	TemplateID string `json:"templateId"`
}

type UpdateProjectValue struct {
	FundingSource  *string  `json:"fundingSource,omitempty"`
	DiscountRate   *float64 `json:"discountRate,omitempty"`
	YearOneValue   *float64 `json:"yearOneValue,omitempty"`
	YearTwoValue   *float64 `json:"yearTwoValue,omitempty"`
	YearThreeValue *float64 `json:"yearThreeValue,omitempty"`
	YearFourValue  *float64 `json:"yearFourValue,omitempty"`
	YearFiveValue  *float64 `json:"yearFiveValue,omitempty"`
}

type UpdateResource struct {
	ID           *string        `json:"id,omitempty"`
	Name         string         `json:"name"`
	Role         *string        `json:"role,omitempty"`
	UserID       *string        `json:"userID,omitempty"`
	Email        *string        `json:"email,omitempty"`
	ProfileImage *string        `json:"profileImage,omitempty"`
	Skills       []*UpdateSkill `json:"skills,omitempty"`
}

type UpdateSkill struct {
	ResourceID  string  `json:"resourceID"`
	ID          string  `json:"id"`
	Proficiency float64 `json:"proficiency"`
}

type UpdateUser struct {
	Name            string  `json:"name"`
	Email           string  `json:"email"`
	Password        string  `json:"password"`
	ConfirmPassword string  `json:"confirmPassword"`
	ProfileImage    *string `json:"profileImage,omitempty"`
}

type User struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	ProfileImage  *string        `json:"profileImage,omitempty"`
	ControlFields *ControlFields `json:"controlFields,omitempty"`
}

type UserResults struct {
	Paging  *Pagination `json:"paging,omitempty"`
	Filters *Filters    `json:"filters"`
	Results []*User     `json:"results,omitempty"`
}