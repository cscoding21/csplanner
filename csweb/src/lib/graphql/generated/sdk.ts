import type { DocumentNode } from 'graphql';
import gql from 'graphql-tag';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Time: { input: any; output: any; }
  Upload: { input: any; output: any; }
};

export type Activity = {
  __typename?: 'Activity';
  activityDate?: Maybe<Scalars['Time']['output']>;
  context: Scalars['String']['output'];
  detail: Scalars['String']['output'];
  id: Scalars['String']['output'];
  link?: Maybe<Scalars['String']['output']>;
  user: User;
  userEmail: Scalars['String']['output'];
};

export type ActivityResults = {
  __typename?: 'ActivityResults';
  filters: Filters;
  paging: Pagination;
  results?: Maybe<Array<Maybe<Activity>>>;
};

export type Artifact = {
  __typename?: 'Artifact';
  fileName: Scalars['String']['output'];
  fileType: Scalars['String']['output'];
  id: Scalars['String']['output'];
  url: Scalars['String']['output'];
};

export type BaseModel = {
  __typename?: 'BaseModel';
  createByUser?: Maybe<User>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['String']['output'];
  deleteByUser?: Maybe<User>;
  deletedAt?: Maybe<Scalars['Time']['output']>;
  deletedBy?: Maybe<Scalars['String']['output']>;
  id: Scalars['String']['output'];
  updateByUser?: Maybe<User>;
  updatedAt: Scalars['Time']['output'];
  updatedBy: Scalars['String']['output'];
};

export type CsWeek = {
  __typename?: 'CSWeek';
  begin: Scalars['Time']['output'];
  end: Scalars['Time']['output'];
  weekNumber: Scalars['Int']['output'];
  year: Scalars['Int']['output'];
};

export type CalculatedScheduleInfo = {
  __typename?: 'CalculatedScheduleInfo';
  isComplete: Scalars['Boolean']['output'];
};

export type Comment = {
  __typename?: 'Comment';
  acknowledges?: Maybe<Array<Scalars['String']['output']>>;
  context?: Maybe<Scalars['String']['output']>;
  dislikes?: Maybe<Array<Scalars['String']['output']>>;
  id: Scalars['String']['output'];
  isActivity: Scalars['Boolean']['output'];
  isEdited: Scalars['Boolean']['output'];
  laughsAt?: Maybe<Array<Scalars['String']['output']>>;
  likes?: Maybe<Array<Scalars['String']['output']>>;
  loves?: Maybe<Array<Scalars['String']['output']>>;
  projectId: Scalars['String']['output'];
  replies?: Maybe<Array<CommentEnvelope>>;
  text: Scalars['String']['output'];
  user: User;
};

export type CommentEnvelope = {
  __typename?: 'CommentEnvelope';
  data: Comment;
  meta: BaseModel;
};

export type CommentResults = {
  __typename?: 'CommentResults';
  filters: Filters;
  paging: Pagination;
  results?: Maybe<Array<Maybe<CommentEnvelope>>>;
};

export type ControlFields = {
  __typename?: 'ControlFields';
  createByUser?: Maybe<User>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['String']['output'];
  deleteByUser?: Maybe<User>;
  deletedAt?: Maybe<Scalars['Time']['output']>;
  deletedBy?: Maybe<Scalars['String']['output']>;
  updateByUser?: Maybe<User>;
  updatedAt: Scalars['Time']['output'];
  updatedBy: Scalars['String']['output'];
};

export type CreateListResult = {
  __typename?: 'CreateListResult';
  list?: Maybe<List>;
  status: Status;
};

export type CreateOrganizationResult = {
  __typename?: 'CreateOrganizationResult';
  organization?: Maybe<Organization>;
  status: Status;
};

export type CreateProjectCommentResult = {
  __typename?: 'CreateProjectCommentResult';
  comment?: Maybe<Comment>;
  status?: Maybe<Status>;
};

export type CreateProjectResult = {
  __typename?: 'CreateProjectResult';
  project?: Maybe<Project>;
  status?: Maybe<Status>;
};

export type CreateProjectTemplateResult = {
  __typename?: 'CreateProjectTemplateResult';
  status: Status;
  template?: Maybe<Projecttemplate>;
};

export type CreateResourceResult = {
  __typename?: 'CreateResourceResult';
  resource?: Maybe<Resource>;
  status: Status;
};

export type CreateRoleResult = {
  __typename?: 'CreateRoleResult';
  role?: Maybe<Role>;
  status: Status;
};

export type CreateUserResult = {
  __typename?: 'CreateUserResult';
  status?: Maybe<Status>;
  user?: Maybe<User>;
};

export type Filter = {
  __typename?: 'Filter';
  customName?: Maybe<Scalars['String']['output']>;
  key: Scalars['String']['output'];
  operation: Scalars['String']['output'];
  value: Scalars['String']['output'];
};

export type Filters = {
  __typename?: 'Filters';
  filters?: Maybe<Array<Maybe<Filter>>>;
};

export type InputFilter = {
  customName?: InputMaybe<Scalars['String']['input']>;
  key: Scalars['String']['input'];
  operation: Scalars['String']['input'];
  value: Scalars['String']['input'];
};

export type InputFilters = {
  filters?: InputMaybe<Array<InputMaybe<InputFilter>>>;
};

export type InputPagination = {
  pageNumber?: InputMaybe<Scalars['Int']['input']>;
  resultsPerPage?: InputMaybe<Scalars['Int']['input']>;
};

export type List = {
  __typename?: 'List';
  description: Scalars['String']['output'];
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  values: Array<ListItem>;
};

export type ListItem = {
  __typename?: 'ListItem';
  name: Scalars['String']['output'];
  sortOrder: Scalars['Int']['output'];
  value: Scalars['String']['output'];
};

export type ListResults = {
  __typename?: 'ListResults';
  filters: Filters;
  paging?: Maybe<Pagination>;
  results?: Maybe<Array<List>>;
};

export type Mutation = {
  __typename?: 'Mutation';
  createProject: CreateProjectResult;
  createProjectComment: CreateProjectCommentResult;
  createProjectCommentReply: CreateProjectCommentResult;
  createResource: CreateResourceResult;
  deleteProject: Status;
  deleteProjectComment: Status;
  deleteProjectFeature: CreateProjectResult;
  deleteProjectTask: CreateProjectResult;
  deleteProjectTemplate: Status;
  deleteProjectValueLine: CreateProjectResult;
  deleteResource: Status;
  deleteResourceSkill: Status;
  deleteRole: Status;
  runProcesses: Status;
  setNotificationsRead: Status;
  setProjectMilestonesFromTemplate: CreateProjectResult;
  setProjectStatus: CreateProjectResult;
  toggleEmote: Status;
  updateAllRoles: CreateRoleResult;
  updateList: CreateListResult;
  updateOrganization: CreateOrganizationResult;
  updateProject: CreateProjectResult;
  updateProjectComment: CreateProjectCommentResult;
  updateProjectFeature: CreateProjectResult;
  updateProjectTask: CreateProjectResult;
  updateProjectTaskStatus: CreateProjectResult;
  updateProjectTemplate: CreateProjectTemplateResult;
  updateProjectValueLine: CreateProjectResult;
  updateResource: CreateResourceResult;
  updateResourceSkill: Status;
  updateRole: CreateRoleResult;
  updateUser: CreateUserResult;
};


export type MutationCreateProjectArgs = {
  input: UpdateNewProject;
};


export type MutationCreateProjectCommentArgs = {
  input: UpdateComment;
};


export type MutationCreateProjectCommentReplyArgs = {
  input: UpdateCommentReply;
};


export type MutationCreateResourceArgs = {
  input: UpdateResource;
};


export type MutationDeleteProjectArgs = {
  id: Scalars['String']['input'];
};


export type MutationDeleteProjectCommentArgs = {
  id: Scalars['String']['input'];
};


export type MutationDeleteProjectFeatureArgs = {
  featureID: Scalars['String']['input'];
  projectID: Scalars['String']['input'];
};


export type MutationDeleteProjectTaskArgs = {
  milestoneID: Scalars['String']['input'];
  projectID: Scalars['String']['input'];
  taskID: Scalars['String']['input'];
};


export type MutationDeleteProjectTemplateArgs = {
  id: Scalars['String']['input'];
};


export type MutationDeleteProjectValueLineArgs = {
  projectID: Scalars['String']['input'];
  valueLineID: Scalars['String']['input'];
};


export type MutationDeleteResourceArgs = {
  id: Scalars['String']['input'];
};


export type MutationDeleteResourceSkillArgs = {
  resourceID: Scalars['String']['input'];
  skillID: Scalars['String']['input'];
};


export type MutationDeleteRoleArgs = {
  id: Scalars['String']['input'];
};


export type MutationSetNotificationsReadArgs = {
  input: Array<Scalars['String']['input']>;
};


export type MutationSetProjectMilestonesFromTemplateArgs = {
  input?: InputMaybe<UpdateProjectMilestoneTemplate>;
};


export type MutationSetProjectStatusArgs = {
  newStatus: Scalars['String']['input'];
  projectID: Scalars['String']['input'];
};


export type MutationToggleEmoteArgs = {
  input: UpdateCommentEmote;
};


export type MutationUpdateAllRolesArgs = {
  input?: InputMaybe<Array<InputMaybe<UpdateRole>>>;
};


export type MutationUpdateListArgs = {
  input: UpdateList;
};


export type MutationUpdateOrganizationArgs = {
  input: UpdateOrganization;
};


export type MutationUpdateProjectArgs = {
  input: UpdateProject;
};


export type MutationUpdateProjectCommentArgs = {
  input: UpdateComment;
};


export type MutationUpdateProjectFeatureArgs = {
  input: UpdateProjectFeature;
};


export type MutationUpdateProjectTaskArgs = {
  input: UpdateProjectMilestoneTask;
};


export type MutationUpdateProjectTaskStatusArgs = {
  milestoneID: Scalars['String']['input'];
  projectID: Scalars['String']['input'];
  status: Scalars['String']['input'];
  taskID: Scalars['String']['input'];
};


export type MutationUpdateProjectTemplateArgs = {
  input?: InputMaybe<UpdateProjecttemplate>;
};


export type MutationUpdateProjectValueLineArgs = {
  input: UpdateProjectValueLine;
};


export type MutationUpdateResourceArgs = {
  input: UpdateResource;
};


export type MutationUpdateResourceSkillArgs = {
  input: UpdateSkill;
};


export type MutationUpdateRoleArgs = {
  input: UpdateRole;
};


export type MutationUpdateUserArgs = {
  input: UpdateUser;
};

export type Notification = {
  __typename?: 'Notification';
  contextId: Scalars['String']['output'];
  id: Scalars['String']['output'];
  initiatorEmail: Scalars['String']['output'];
  initiatorName: Scalars['String']['output'];
  initiatorProfileImage?: Maybe<Scalars['String']['output']>;
  isRead: Scalars['Boolean']['output'];
  recipientIsBot: Scalars['Boolean']['output'];
  text?: Maybe<Scalars['String']['output']>;
  type: Scalars['Int']['output'];
  updatedAt?: Maybe<Scalars['Time']['output']>;
  userEmail: Scalars['String']['output'];
  userName: Scalars['String']['output'];
};

export type NotificationResults = {
  __typename?: 'NotificationResults';
  filters: Filters;
  paging?: Maybe<Pagination>;
  results: Array<Notification>;
};

export type Organization = {
  __typename?: 'Organization';
  defaults: OrganizationDefaults;
  id?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
  setup: OrganizationSetup;
  urlKey: Scalars['String']['output'];
};

export type OrganizationDefaults = {
  __typename?: 'OrganizationDefaults';
  commsCoefficient: Scalars['Float']['output'];
  discountRate: Scalars['Float']['output'];
  focusFactor: Scalars['Float']['output'];
  genericBlendedHourlyRate: Scalars['Float']['output'];
  hoursPerWeek: Scalars['Int']['output'];
  workingHoursPerYear: Scalars['Float']['output'];
};

export type OrganizationSetup = {
  __typename?: 'OrganizationSetup';
  hasFundingSources: Scalars['Boolean']['output'];
  hasResources: Scalars['Boolean']['output'];
  hasReviewedOrgSettings: Scalars['Boolean']['output'];
  hasRoles: Scalars['Boolean']['output'];
  hasSkills: Scalars['Boolean']['output'];
  hasTemplates: Scalars['Boolean']['output'];
  hasValueCategories: Scalars['Boolean']['output'];
  isReadyForProjects: Scalars['Boolean']['output'];
};

export type PageAndFilter = {
  filters?: InputMaybe<InputFilters>;
  paging?: InputMaybe<InputPagination>;
};

export type Pagination = {
  __typename?: 'Pagination';
  after?: Maybe<Scalars['String']['output']>;
  pageNumber?: Maybe<Scalars['Int']['output']>;
  resultsPerPage?: Maybe<Scalars['Int']['output']>;
  totalResults?: Maybe<Scalars['Int']['output']>;
};

export type Portfolio = {
  __typename?: 'Portfolio';
  begin?: Maybe<Scalars['Time']['output']>;
  calculated: PortfolioCalculatedData;
  end?: Maybe<Scalars['Time']['output']>;
  schedule: Array<Schedule>;
  weekSummary: Array<Maybe<PortfolioWeekSummary>>;
};

export type PortfolioCalculatedData = {
  __typename?: 'PortfolioCalculatedData';
  countInFlight: Scalars['Int']['output'];
  countScheduled: Scalars['Int']['output'];
  totalCount: Scalars['Int']['output'];
  totalInFlight: Scalars['Int']['output'];
  totalValue: Scalars['Float']['output'];
  valueInFlight: Scalars['Float']['output'];
  valueScheduled: Scalars['Float']['output'];
};

export type PortfolioWeekSummary = {
  __typename?: 'PortfolioWeekSummary';
  allocatedHours: Scalars['Int']['output'];
  begin: Scalars['Time']['output'];
  end: Scalars['Time']['output'];
  orgCapacity: Scalars['Int']['output'];
  weekNumber: Scalars['Int']['output'];
  year: Scalars['Int']['output'];
};

export type Project = {
  __typename?: 'Project';
  calculated?: Maybe<ProjectCalculatedData>;
  id?: Maybe<Scalars['String']['output']>;
  projectBasics: ProjectBasics;
  projectCost: ProjectCost;
  projectDaci: ProjectDaci;
  projectFeatures?: Maybe<Array<ProjectFeature>>;
  projectMilestones?: Maybe<Array<ProjectMilestone>>;
  projectStatusBlock: ProjectStatusBlock;
  projectValue: ProjectValue;
};

export type ProjectActivity = {
  __typename?: 'ProjectActivity';
  hoursSpent: Scalars['Int']['output'];
  milestoneID: Scalars['String']['output'];
  milestoneName: Scalars['String']['output'];
  project?: Maybe<Project>;
  projectID: Scalars['String']['output'];
  requiredSkillID: Scalars['String']['output'];
  resource?: Maybe<Resource>;
  resourceID: Scalars['String']['output'];
  status: Scalars['String']['output'];
  taskEndDate?: Maybe<Scalars['Time']['output']>;
  taskID: Scalars['String']['output'];
  taskName: Scalars['String']['output'];
};

export type ProjectActivityWeek = {
  __typename?: 'ProjectActivityWeek';
  activities?: Maybe<Array<ProjectActivity>>;
  begin: Scalars['Time']['output'];
  end: Scalars['Time']['output'];
  orgCapacity: Scalars['Int']['output'];
  weekNumber: Scalars['Int']['output'];
  year: Scalars['Int']['output'];
};

export type ProjectBasics = {
  __typename?: 'ProjectBasics';
  description: Scalars['String']['output'];
  isCapitalized: Scalars['Boolean']['output'];
  name: Scalars['String']['output'];
  owner?: Maybe<User>;
  ownerID?: Maybe<Scalars['String']['output']>;
  startDate?: Maybe<Scalars['Time']['output']>;
};

export type ProjectCalculatedData = {
  __typename?: 'ProjectCalculatedData';
  completedCost?: Maybe<Scalars['Float']['output']>;
  completedHours?: Maybe<Scalars['Int']['output']>;
  completedTasks?: Maybe<Scalars['Int']['output']>;
  featureStatusAcceptedCount?: Maybe<Scalars['Int']['output']>;
  featureStatusDoneCount?: Maybe<Scalars['Int']['output']>;
  featureStatusProposedCount?: Maybe<Scalars['Int']['output']>;
  featureStatusRemovedCount?: Maybe<Scalars['Int']['output']>;
  healthyTasks?: Maybe<Scalars['Int']['output']>;
  projectPercentComplete?: Maybe<Scalars['Float']['output']>;
  remainingCost?: Maybe<Scalars['Float']['output']>;
  remainingHours?: Maybe<Scalars['Int']['output']>;
  remainingTasks?: Maybe<Scalars['Int']['output']>;
  teamMembers?: Maybe<Array<Maybe<Resource>>>;
  totalTasks?: Maybe<Scalars['Int']['output']>;
  unhealthyTasks?: Maybe<Scalars['Int']['output']>;
};

export type ProjectCost = {
  __typename?: 'ProjectCost';
  blendedRate?: Maybe<Scalars['Float']['output']>;
  calculated?: Maybe<ProjectCostCalculatedData>;
  ongoing?: Maybe<Scalars['Float']['output']>;
};

export type ProjectCostCalculatedData = {
  __typename?: 'ProjectCostCalculatedData';
  hourEstimate: Scalars['Int']['output'];
  hoursActualized: Scalars['Int']['output'];
  initialCost?: Maybe<Scalars['Float']['output']>;
};

export type ProjectDaci = {
  __typename?: 'ProjectDaci';
  approver?: Maybe<Array<Maybe<Resource>>>;
  contributor?: Maybe<Array<Maybe<Resource>>>;
  driver?: Maybe<Array<Maybe<Resource>>>;
  informed?: Maybe<Array<Maybe<Resource>>>;
};

export type ProjectEnvelope = {
  __typename?: 'ProjectEnvelope';
  data: Project;
  meta: BaseModel;
};

export type ProjectFeature = {
  __typename?: 'ProjectFeature';
  description: Scalars['String']['output'];
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  priority: Scalars['String']['output'];
  status: Scalars['String']['output'];
};

export type ProjectFilters = {
  __typename?: 'ProjectFilters';
  status?: Maybe<Scalars['String']['output']>;
};

export type ProjectMilestone = {
  __typename?: 'ProjectMilestone';
  calculated?: Maybe<ProjectMilestoneCalculatedData>;
  id: Scalars['String']['output'];
  phase: ProjectMilestonePhase;
  tasks: Array<ProjectMilestoneTask>;
};

export type ProjectMilestoneCalculatedData = {
  __typename?: 'ProjectMilestoneCalculatedData';
  completedTasks?: Maybe<Scalars['Int']['output']>;
  estimatedEndDate?: Maybe<Scalars['Time']['output']>;
  estimatedStartDate?: Maybe<Scalars['Time']['output']>;
  hoursRemaining?: Maybe<Scalars['Int']['output']>;
  isComplete?: Maybe<Scalars['Boolean']['output']>;
  isInFlight?: Maybe<Scalars['Boolean']['output']>;
  percentDone?: Maybe<Scalars['Float']['output']>;
  removedHours?: Maybe<Scalars['Int']['output']>;
  totalHours?: Maybe<Scalars['Int']['output']>;
  totalTasks?: Maybe<Scalars['Int']['output']>;
  unhealthyTasks?: Maybe<Scalars['Int']['output']>;
};

export type ProjectMilestonePhase = {
  __typename?: 'ProjectMilestonePhase';
  description: Scalars['String']['output'];
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  order: Scalars['Int']['output'];
};

export type ProjectMilestoneTask = {
  __typename?: 'ProjectMilestoneTask';
  calculated?: Maybe<ProjectTaskCalculatedData>;
  description?: Maybe<Scalars['String']['output']>;
  hourEstimate: Scalars['Int']['output'];
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  requiredSkillID: Scalars['String']['output'];
  resourceIDs?: Maybe<Array<Scalars['String']['output']>>;
  resources?: Maybe<Array<Resource>>;
  skills?: Maybe<Array<Skill>>;
  status: Scalars['String']['output'];
};

export type ProjectResults = {
  __typename?: 'ProjectResults';
  filters: Filters;
  paging?: Maybe<Pagination>;
  results?: Maybe<Array<ProjectEnvelope>>;
};

export type ProjectScheduleResult = {
  __typename?: 'ProjectScheduleResult';
  schedule: Schedule;
};

export type ProjectStatusBlock = {
  __typename?: 'ProjectStatusBlock';
  allowedNextStates?: Maybe<Array<ProjectStatusTransition>>;
  status: Scalars['String']['output'];
};

export type ProjectStatusTransition = {
  __typename?: 'ProjectStatusTransition';
  canEnter: Scalars['Boolean']['output'];
  checkResult?: Maybe<ValidationResult>;
  nextState: Scalars['String']['output'];
};

export type ProjectTaskCalculatedData = {
  __typename?: 'ProjectTaskCalculatedData';
  actualizedCost?: Maybe<Scalars['Float']['output']>;
  actualizedHoursToComplete?: Maybe<Scalars['Int']['output']>;
  averageHourlyRate?: Maybe<Scalars['Float']['output']>;
  commsHourAdjustment?: Maybe<Scalars['Int']['output']>;
  exceptions?: Maybe<Array<Scalars['String']['output']>>;
  portfolioEstimatedCompleteDate?: Maybe<Scalars['Time']['output']>;
  resourceContention?: Maybe<Scalars['Float']['output']>;
  skillsHourAdjustment?: Maybe<Scalars['Int']['output']>;
};

export type ProjectTemplateTask = {
  __typename?: 'ProjectTemplateTask';
  description?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
  requiredSkillID?: Maybe<Scalars['String']['output']>;
};

export type ProjectValue = {
  __typename?: 'ProjectValue';
  calculated?: Maybe<ProjectValueCalculatedData>;
  discountRate?: Maybe<Scalars['Float']['output']>;
  projectValueLines?: Maybe<Array<ProjectValueLine>>;
};

export type ProjectValueCalculatedData = {
  __typename?: 'ProjectValueCalculatedData';
  fiveYearGross?: Maybe<Scalars['Float']['output']>;
  internalRateOfReturn?: Maybe<Scalars['Float']['output']>;
  netPresentValue?: Maybe<Scalars['Float']['output']>;
  yearFiveValue?: Maybe<Scalars['Float']['output']>;
  yearFourValue?: Maybe<Scalars['Float']['output']>;
  yearOneValue?: Maybe<Scalars['Float']['output']>;
  yearThreeValue?: Maybe<Scalars['Float']['output']>;
  yearTwoValue?: Maybe<Scalars['Float']['output']>;
};

export type ProjectValueLine = {
  __typename?: 'ProjectValueLine';
  description?: Maybe<Scalars['String']['output']>;
  fundingSource: Scalars['String']['output'];
  id: Scalars['String']['output'];
  valueCategory: Scalars['String']['output'];
  yearFiveValue?: Maybe<Scalars['Float']['output']>;
  yearFourValue?: Maybe<Scalars['Float']['output']>;
  yearOneValue?: Maybe<Scalars['Float']['output']>;
  yearThreeValue?: Maybe<Scalars['Float']['output']>;
  yearTwoValue?: Maybe<Scalars['Float']['output']>;
};

export type Projecttemplate = {
  __typename?: 'Projecttemplate';
  description: Scalars['String']['output'];
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  phases: Array<ProjecttemplatePhase>;
};

export type ProjecttemplatePhase = {
  __typename?: 'ProjecttemplatePhase';
  description: Scalars['String']['output'];
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  order: Scalars['Int']['output'];
  tasks?: Maybe<Array<ProjectTemplateTask>>;
};

export type ProjecttemplateResults = {
  __typename?: 'ProjecttemplateResults';
  filters: Filters;
  paging?: Maybe<Pagination>;
  results?: Maybe<Array<Projecttemplate>>;
};

export type Query = {
  __typename?: 'Query';
  calculateProjectSchedule: ProjectScheduleResult;
  checkProjectStatus: ValidationResult;
  currentUser?: Maybe<User>;
  findActivity: ActivityResults;
  findAllLists: ListResults;
  findAllProjectTemplates: ProjecttemplateResults;
  findAllResources: ResourceResults;
  findAllRoles: RoleResults;
  findAllUsers: UserResults;
  findProjectComments: CommentResults;
  findProjects: ProjectResults;
  findResources: ResourceResults;
  findUserNotifications: NotificationResults;
  getCommentThread: CommentEnvelope;
  getDraftPortfolio: Portfolio;
  getList?: Maybe<List>;
  getOrganization: Organization;
  getPortfolio: Portfolio;
  getPortfolioForResource: Portfolio;
  getProject: ProjectEnvelope;
  getResource: ResourceEnvelope;
  getUser: User;
};


export type QueryCalculateProjectScheduleArgs = {
  projectID: Scalars['String']['input'];
  startDate: Scalars['Time']['input'];
};


export type QueryCheckProjectStatusArgs = {
  newStatus: Scalars['String']['input'];
  projectID: Scalars['String']['input'];
};


export type QueryFindActivityArgs = {
  pageAndFilter: PageAndFilter;
};


export type QueryFindProjectCommentsArgs = {
  projectID: Scalars['String']['input'];
};


export type QueryFindProjectsArgs = {
  pageAndFilter: PageAndFilter;
};


export type QueryFindResourcesArgs = {
  pageAndFilter?: InputMaybe<PageAndFilter>;
};


export type QueryFindUserNotificationsArgs = {
  pageAndFilter?: InputMaybe<PageAndFilter>;
};


export type QueryGetCommentThreadArgs = {
  id: Scalars['String']['input'];
};


export type QueryGetDraftPortfolioArgs = {
  additionalID: Scalars['String']['input'];
};


export type QueryGetListArgs = {
  nameOrID: Scalars['String']['input'];
};


export type QueryGetPortfolioForResourceArgs = {
  resourceID: Scalars['String']['input'];
};


export type QueryGetProjectArgs = {
  id: Scalars['String']['input'];
};


export type QueryGetResourceArgs = {
  id: Scalars['String']['input'];
};


export type QueryGetUserArgs = {
  id: Scalars['String']['input'];
};

export type RelateArtifact = {
  __typename?: 'RelateArtifact';
  artifactID: Scalars['String']['output'];
  linkID: Scalars['String']['output'];
};

export type Resource = {
  __typename?: 'Resource';
  annualizedCost?: Maybe<Scalars['Float']['output']>;
  availableHoursPerWeek?: Maybe<Scalars['Int']['output']>;
  calculated: ResourceCalculatedData;
  id?: Maybe<Scalars['String']['output']>;
  initialCost?: Maybe<Scalars['Float']['output']>;
  name: Scalars['String']['output'];
  profileImage?: Maybe<Scalars['String']['output']>;
  role?: Maybe<Role>;
  roleID?: Maybe<Scalars['String']['output']>;
  skills?: Maybe<Array<Skill>>;
  status: Scalars['String']['output'];
  type: Scalars['String']['output'];
  user?: Maybe<User>;
  userEmail?: Maybe<Scalars['String']['output']>;
};

export type ResourceAllocationGrid = {
  __typename?: 'ResourceAllocationGrid';
  weekActivities?: Maybe<Array<Maybe<ProjectActivity>>>;
};

export type ResourceCalculatedData = {
  __typename?: 'ResourceCalculatedData';
  hourlyCost?: Maybe<Scalars['Float']['output']>;
  hourlyCostMethod: Scalars['String']['output'];
};

export type ResourceEnvelope = {
  __typename?: 'ResourceEnvelope';
  data: Resource;
  meta: BaseModel;
};

export type ResourceResults = {
  __typename?: 'ResourceResults';
  filters: Filters;
  paging?: Maybe<Pagination>;
  results?: Maybe<Array<ResourceEnvelope>>;
};

export type Role = {
  __typename?: 'Role';
  defaultSkills?: Maybe<Array<Skill>>;
  description: Scalars['String']['output'];
  hourlyRate?: Maybe<Scalars['Float']['output']>;
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
};

export type RoleEnvelope = {
  __typename?: 'RoleEnvelope';
  data: Role;
  meta: BaseModel;
};

export type RoleResults = {
  __typename?: 'RoleResults';
  filters: Filters;
  paging?: Maybe<Pagination>;
  results?: Maybe<Array<RoleEnvelope>>;
};

export type Schedule = {
  __typename?: 'Schedule';
  begin?: Maybe<Scalars['Time']['output']>;
  calculated?: Maybe<CalculatedScheduleInfo>;
  end?: Maybe<Scalars['Time']['output']>;
  exceptions?: Maybe<Array<ScheduleException>>;
  project: Project;
  projectActivityWeeks?: Maybe<Array<ProjectActivityWeek>>;
  projectID: Scalars['String']['output'];
  projectName: Scalars['String']['output'];
};

export type ScheduleException = {
  __typename?: 'ScheduleException';
  level: Scalars['Int']['output'];
  message: Scalars['String']['output'];
  scope: Scalars['String']['output'];
  type: Scalars['String']['output'];
};

export type Skill = {
  __typename?: 'Skill';
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  proficiency?: Maybe<Scalars['Float']['output']>;
  skillID?: Maybe<Scalars['String']['output']>;
};

export type Status = {
  __typename?: 'Status';
  message: Array<Scalars['String']['output']>;
  success: Scalars['Boolean']['output'];
  validationResult?: Maybe<ValidationResult>;
};

/**
 * `Subscription` is where all the subscriptions your clients can
 * request. You can use Schema Directives like normal to restrict
 * access.
 */
export type Subscription = {
  __typename?: 'Subscription';
  /** `currentTime` will return a stream of `Time` objects. */
  currentTime: TimeResponse;
  notificationUpdate: Scalars['String']['output'];
};

export type TimeResponse = {
  __typename?: 'TimeResponse';
  timeStamp: Scalars['String']['output'];
  unixTime: Scalars['Int']['output'];
};

export type UpdateComment = {
  id?: InputMaybe<Scalars['String']['input']>;
  projectId: Scalars['String']['input'];
  text: Scalars['String']['input'];
};

export type UpdateCommentEmote = {
  commentID: Scalars['String']['input'];
  emoteType: Scalars['String']['input'];
  parentID?: InputMaybe<Scalars['String']['input']>;
  projectID: Scalars['String']['input'];
};

export type UpdateCommentReply = {
  parentCommentID: Scalars['String']['input'];
  text: Scalars['String']['input'];
};

export type UpdateList = {
  description: Scalars['String']['input'];
  id: Scalars['String']['input'];
  values: Array<UpdateListItem>;
};

export type UpdateListItem = {
  name: Scalars['String']['input'];
  sortOrder?: InputMaybe<Scalars['Int']['input']>;
  value: Scalars['String']['input'];
};

export type UpdateNewProject = {
  description: Scalars['String']['input'];
  isCapitalized: Scalars['Boolean']['input'];
  name: Scalars['String']['input'];
  ownerID: Scalars['String']['input'];
  templateID: Scalars['String']['input'];
};

export type UpdateOrganization = {
  defaults: UpdateOrganizationDefaults;
  id?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
};

export type UpdateOrganizationDefaults = {
  commsCoefficient: Scalars['Float']['input'];
  discountRate: Scalars['Float']['input'];
  focusFactor?: InputMaybe<Scalars['Float']['input']>;
  genericBlendedHourlyRate: Scalars['Float']['input'];
  hoursPerWeek: Scalars['Int']['input'];
  workingHoursPerYear: Scalars['Float']['input'];
};

export type UpdateProject = {
  id?: InputMaybe<Scalars['String']['input']>;
  projectBasics?: InputMaybe<UpdateProjectBasics>;
  projectCost?: InputMaybe<UpdateProjectCost>;
  projectDaci?: InputMaybe<UpdateProjectDaci>;
  projectFeatures?: InputMaybe<Array<InputMaybe<UpdateProjectFeature>>>;
  projectMilestones?: InputMaybe<Array<InputMaybe<UpdateProjectMilestone>>>;
  projectValue?: InputMaybe<UpdateProjectValue>;
};

export type UpdateProjectBasics = {
  description?: InputMaybe<Scalars['String']['input']>;
  isCapitalized: Scalars['Boolean']['input'];
  name: Scalars['String']['input'];
  ownerID?: InputMaybe<Scalars['String']['input']>;
  startDate?: InputMaybe<Scalars['Time']['input']>;
};

export type UpdateProjectCost = {
  blendedRate?: InputMaybe<Scalars['Float']['input']>;
  ongoing?: InputMaybe<Scalars['Float']['input']>;
};

export type UpdateProjectDaci = {
  approverIDs?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>;
  contributorIDs?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>;
  driverIDs?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>;
  informedIDs?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>;
};

export type UpdateProjectFeature = {
  description: Scalars['String']['input'];
  id?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
  priority: Scalars['String']['input'];
  projectID: Scalars['String']['input'];
  status: Scalars['String']['input'];
};

export type UpdateProjectMilestone = {
  id: Scalars['String']['input'];
  phase: UpdateProjectMilestonePhase;
  tasks?: InputMaybe<Array<UpdateProjectMilestoneTask>>;
};

export type UpdateProjectMilestonePhase = {
  description: Scalars['String']['input'];
  id: Scalars['String']['input'];
  name: Scalars['String']['input'];
  order: Scalars['Int']['input'];
};

export type UpdateProjectMilestoneTask = {
  description?: InputMaybe<Scalars['String']['input']>;
  hourEstimate: Scalars['Int']['input'];
  id?: InputMaybe<Scalars['String']['input']>;
  milestoneID: Scalars['String']['input'];
  name: Scalars['String']['input'];
  projectID: Scalars['String']['input'];
  requiredSkillID: Scalars['String']['input'];
  resourceIDs?: InputMaybe<Array<Scalars['String']['input']>>;
  status: Scalars['String']['input'];
};

export type UpdateProjectMilestoneTemplate = {
  projectId: Scalars['String']['input'];
  templateId: Scalars['String']['input'];
};

export type UpdateProjectTemplateTask = {
  description?: InputMaybe<Scalars['String']['input']>;
  id?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
  requiredSkillID?: InputMaybe<Scalars['String']['input']>;
};

export type UpdateProjectValue = {
  discountRate: Scalars['Float']['input'];
};

export type UpdateProjectValueLine = {
  description?: InputMaybe<Scalars['String']['input']>;
  fundingSource: Scalars['String']['input'];
  id?: InputMaybe<Scalars['String']['input']>;
  projectID: Scalars['String']['input'];
  valueCategory: Scalars['String']['input'];
  yearFiveValue?: InputMaybe<Scalars['Float']['input']>;
  yearFourValue?: InputMaybe<Scalars['Float']['input']>;
  yearOneValue?: InputMaybe<Scalars['Float']['input']>;
  yearThreeValue?: InputMaybe<Scalars['Float']['input']>;
  yearTwoValue?: InputMaybe<Scalars['Float']['input']>;
};

export type UpdateProjecttemplate = {
  description: Scalars['String']['input'];
  id?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
  phases: Array<UpdateProjecttemplatePhase>;
};

export type UpdateProjecttemplatePhase = {
  description: Scalars['String']['input'];
  id?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
  order: Scalars['Int']['input'];
  tasks?: InputMaybe<Array<UpdateProjectTemplateTask>>;
};

export type UpdateResource = {
  annualizedCost?: InputMaybe<Scalars['Float']['input']>;
  availableHoursPerWeek?: InputMaybe<Scalars['Int']['input']>;
  email?: InputMaybe<Scalars['String']['input']>;
  id?: InputMaybe<Scalars['String']['input']>;
  initialCost?: InputMaybe<Scalars['Float']['input']>;
  name: Scalars['String']['input'];
  profileImage?: InputMaybe<Scalars['String']['input']>;
  roleID?: InputMaybe<Scalars['String']['input']>;
  skills?: InputMaybe<Array<UpdateSkill>>;
  status: Scalars['String']['input'];
  type: Scalars['String']['input'];
  userID?: InputMaybe<Scalars['String']['input']>;
};

export type UpdateRole = {
  defaultSkills?: InputMaybe<Array<UpdateSkill>>;
  description: Scalars['String']['input'];
  hourlyRate?: InputMaybe<Scalars['Float']['input']>;
  id?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
};

export type UpdateSkill = {
  id?: InputMaybe<Scalars['String']['input']>;
  parentID: Scalars['String']['input'];
  proficiency: Scalars['Float']['input'];
  skillID: Scalars['String']['input'];
};

export type UpdateUser = {
  email: Scalars['String']['input'];
  firstName: Scalars['String']['input'];
  id: Scalars['String']['input'];
  lastName: Scalars['String']['input'];
  profileImage?: InputMaybe<Scalars['String']['input']>;
};

export type User = {
  __typename?: 'User';
  email: Scalars['String']['output'];
  firstName: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  lastName: Scalars['String']['output'];
  profileImage?: Maybe<Scalars['String']['output']>;
};

export type UserResults = {
  __typename?: 'UserResults';
  filters: Filters;
  paging?: Maybe<Pagination>;
  results?: Maybe<Array<User>>;
};

export type ValidationMessage = {
  __typename?: 'ValidationMessage';
  field: Scalars['String']['output'];
  message: Scalars['String']['output'];
};

export type ValidationResult = {
  __typename?: 'ValidationResult';
  messages?: Maybe<Array<Maybe<ValidationMessage>>>;
  pass: Scalars['Boolean']['output'];
};

export const UserFragmentFragmentDoc = gql`
    fragment userFragment on User {
  id
  firstName
  lastName
  email
  profileImage
}
    `;
export const BaseModelFragmentFragmentDoc = gql`
    fragment baseModelFragment on BaseModel {
  id
  createdBy
  createdAt
  updatedAt
  updatedBy
  deletedAt
  deletedBy
  createByUser {
    ...userFragment
  }
  updateByUser {
    ...userFragment
  }
  deleteByUser {
    ...userFragment
  }
}
    ${UserFragmentFragmentDoc}`;
export const CommentFragmentFragmentDoc = gql`
    fragment commentFragment on Comment {
  id
  text
  isEdited
  isActivity
  context
  projectId
  replies {
    meta {
      ...baseModelFragment
    }
    data {
      id
      text
      isEdited
      projectId
      likes
      loves
      dislikes
      laughsAt
      acknowledges
    }
  }
  likes
  loves
  dislikes
  laughsAt
  acknowledges
}
    ${BaseModelFragmentFragmentDoc}`;
export const ControlFieldsFragmentFragmentDoc = gql`
    fragment controlFieldsFragment on ControlFields {
  createdBy
  createdAt
  updatedAt
  updatedBy
  deletedAt
  deletedBy
  createByUser {
    ...userFragment
  }
  updateByUser {
    ...userFragment
  }
  deleteByUser {
    ...userFragment
  }
}
    ${UserFragmentFragmentDoc}`;
export const StatusFragmentFragmentDoc = gql`
    fragment statusFragment on Status {
  success
  message
}
    `;
export const PagingFragmentFragmentDoc = gql`
    fragment pagingFragment on Pagination {
  pageNumber
  totalResults
  resultsPerPage
  after
}
    `;
export const ValidationResultFragmentFragmentDoc = gql`
    fragment validationResultFragment on ValidationResult {
  pass
  messages {
    field
    message
  }
}
    `;
export const ListFragmentFragmentDoc = gql`
    fragment listFragment on List {
  id
  name
  description
  values {
    value
    name
    sortOrder
  }
}
    `;
export const NotificationFragmentFragmentDoc = gql`
    fragment notificationFragment on Notification {
  id
  userEmail
  userName
  text
  recipientIsBot
  isRead
  recipientIsBot
  type
  contextId
  initiatorName
  initiatorEmail
  initiatorProfileImage
}
    `;
export const OrganizationFragmentFragmentDoc = gql`
    fragment organizationFragment on Organization {
  name
  id
  urlKey
  defaults {
    discountRate
    commsCoefficient
    focusFactor
    hoursPerWeek
    workingHoursPerYear
    genericBlendedHourlyRate
  }
  setup {
    hasSkills
    hasFundingSources
    hasValueCategories
    hasRoles
    hasTemplates
    hasResources
    hasReviewedOrgSettings
    isReadyForProjects
  }
}
    `;
export const SkillFragmentFragmentDoc = gql`
    fragment skillFragment on Skill {
  id
  skillID
  name
  proficiency
}
    `;
export const RoleFragmentFragmentDoc = gql`
    fragment roleFragment on Role {
  id
  name
  description
  hourlyRate
  defaultSkills {
    ...skillFragment
  }
}
    ${SkillFragmentFragmentDoc}`;
export const ResourceFragmentFragmentDoc = gql`
    fragment resourceFragment on Resource {
  name
  id
  user {
    ...userFragment
  }
  roleID
  role {
    ...roleFragment
  }
  profileImage
  userEmail
  initialCost
  annualizedCost
  type
  status
  availableHoursPerWeek
  skills {
    ...skillFragment
  }
  calculated {
    hourlyCost
    hourlyCostMethod
  }
}
    ${UserFragmentFragmentDoc}
${RoleFragmentFragmentDoc}
${SkillFragmentFragmentDoc}`;
export const ProjectFragmentFragmentDoc = gql`
    fragment projectFragment on Project {
  id
  projectBasics {
    name
    description
    startDate
    ownerID
    isCapitalized
    owner {
      ...userFragment
    }
  }
  projectStatusBlock {
    status
    allowedNextStates {
      nextState
      canEnter
      checkResult {
        pass
        messages {
          field
          message
        }
      }
    }
  }
  projectValue {
    discountRate
    projectValueLines {
      id
      fundingSource
      valueCategory
      yearOneValue
      yearTwoValue
      yearThreeValue
      yearFourValue
      yearFiveValue
      description
    }
    calculated {
      netPresentValue
      internalRateOfReturn
      yearOneValue
      yearTwoValue
      yearThreeValue
      yearFourValue
      yearFiveValue
      fiveYearGross
    }
  }
  projectCost {
    ongoing
    blendedRate
    calculated {
      initialCost
      hourEstimate
      hoursActualized
    }
  }
  projectDaci {
    driver {
      ...resourceFragment
    }
    approver {
      ...resourceFragment
    }
    contributor {
      ...resourceFragment
    }
    informed {
      ...resourceFragment
    }
  }
  projectFeatures {
    id
    name
    description
    priority
    status
  }
  projectMilestones {
    id
    calculated {
      hoursRemaining
      totalHours
      removedHours
      completedTasks
      totalTasks
      isComplete
      isInFlight
      unhealthyTasks
      percentDone
    }
    phase {
      name
      description
      id
      order
    }
    tasks {
      id
      name
      description
      status
      hourEstimate
      requiredSkillID
      skills {
        id
        name
      }
      resourceIDs
      resources {
        ...resourceFragment
      }
      calculated {
        resourceContention
        actualizedCost
        actualizedHoursToComplete
        commsHourAdjustment
        skillsHourAdjustment
        averageHourlyRate
        exceptions
        portfolioEstimatedCompleteDate
      }
    }
  }
  calculated {
    teamMembers {
      ...resourceFragment
    }
    featureStatusAcceptedCount
    featureStatusProposedCount
    featureStatusDoneCount
    featureStatusRemovedCount
    unhealthyTasks
    healthyTasks
    totalTasks
    completedTasks
    remainingTasks
    completedCost
    remainingCost
    completedHours
    remainingHours
    projectPercentComplete
  }
}
    ${UserFragmentFragmentDoc}
${ResourceFragmentFragmentDoc}`;
export const ScheduleFragmentFragmentDoc = gql`
    fragment scheduleFragment on Schedule {
  projectID
  project {
    ...projectFragment
  }
  begin
  end
  exceptions {
    level
    type
    message
    scope
  }
  projectActivityWeeks {
    end
    begin
    year
    weekNumber
    orgCapacity
    activities {
      taskName
      taskID
      resourceID
      hoursSpent
      requiredSkillID
      taskEndDate
      status
      resource {
        ...resourceFragment
      }
    }
  }
  calculated {
    isComplete
  }
}
    ${ProjectFragmentFragmentDoc}
${ResourceFragmentFragmentDoc}`;
export const PortfolioFragmentFragmentDoc = gql`
    fragment portfolioFragment on Portfolio {
  begin
  end
  weekSummary {
    weekNumber
    year
    begin
    end
    orgCapacity
    allocatedHours
  }
  schedule {
    ...scheduleFragment
  }
  calculated {
    countInFlight
    countScheduled
    totalCount
    valueInFlight
    valueScheduled
    totalValue
  }
}
    ${ScheduleFragmentFragmentDoc}`;
export const ScheduleResultFragmentFragmentDoc = gql`
    fragment scheduleResultFragment on ProjectScheduleResult {
  schedule {
    ...scheduleFragment
  }
}
    ${ScheduleFragmentFragmentDoc}`;
export const ProjectTemplateFragmentFragmentDoc = gql`
    fragment projectTemplateFragment on Projecttemplate {
  name
  description
  id
  phases {
    id
    name
    order
    description
    tasks {
      name
      description
      requiredSkillID
    }
  }
}
    `;
export const FindActivityDocument = gql`
    query findActivity($input: PageAndFilter!) {
  findActivity(pageAndFilter: $input) {
    paging {
      ...pagingFragment
    }
    results {
      id
      detail
      link
      context
      userEmail
      activityDate
      user {
        ...userFragment
      }
    }
  }
}
    ${PagingFragmentFragmentDoc}
${UserFragmentFragmentDoc}`;
export const CreateProjectCommentDocument = gql`
    mutation createProjectComment($comment: UpdateComment!) {
  createProjectComment(input: $comment) {
    status {
      ...statusFragment
    }
    comment {
      id
      text
    }
  }
}
    ${StatusFragmentFragmentDoc}`;
export const UpdateProjectCommentDocument = gql`
    mutation updateProjectComment($comment: UpdateComment!) {
  updateProjectComment(input: $comment) {
    status {
      ...statusFragment
    }
    comment {
      id
      text
    }
  }
}
    ${StatusFragmentFragmentDoc}`;
export const CreateProjectCommentReplyDocument = gql`
    mutation createProjectCommentReply($reply: UpdateCommentReply!) {
  createProjectCommentReply(input: $reply) {
    status {
      ...statusFragment
    }
    comment {
      id
      text
    }
  }
}
    ${StatusFragmentFragmentDoc}`;
export const DeleteProjectCommentDocument = gql`
    mutation deleteProjectComment($id: String!) {
  deleteProjectComment(id: $id) {
    ...statusFragment
  }
}
    ${StatusFragmentFragmentDoc}`;
export const FindProjectCommentsDocument = gql`
    query findProjectComments($projectID: String!) {
  findProjectComments(projectID: $projectID) {
    paging {
      ...pagingFragment
    }
    results {
      meta {
        ...baseModelFragment
      }
      data {
        ...commentFragment
      }
    }
  }
}
    ${PagingFragmentFragmentDoc}
${BaseModelFragmentFragmentDoc}
${CommentFragmentFragmentDoc}`;
export const GetCommentThreadDocument = gql`
    query getCommentThread($commentID: String!) {
  getCommentThread(id: $commentID) {
    meta {
      ...baseModelFragment
    }
    data {
      ...commentFragment
    }
  }
}
    ${BaseModelFragmentFragmentDoc}
${CommentFragmentFragmentDoc}`;
export const FindAllListsDocument = gql`
    query findAllLists {
  findAllLists {
    paging {
      ...pagingFragment
    }
    results {
      ...listFragment
    }
  }
}
    ${PagingFragmentFragmentDoc}
${ListFragmentFragmentDoc}`;
export const GetListDocument = gql`
    query getList($nameOrID: String!) {
  getList(nameOrID: $nameOrID) {
    ...listFragment
  }
}
    ${ListFragmentFragmentDoc}`;
export const UpdateListDocument = gql`
    mutation updateList($input: UpdateList!) {
  updateList(input: $input) {
    status {
      ...statusFragment
    }
    list {
      ...listFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ListFragmentFragmentDoc}`;
export const FindUserNotificationsDocument = gql`
    query findUserNotifications($input: PageAndFilter!) {
  findUserNotifications(pageAndFilter: $input) {
    paging {
      ...pagingFragment
    }
    results {
      ...notificationFragment
    }
  }
}
    ${PagingFragmentFragmentDoc}
${NotificationFragmentFragmentDoc}`;
export const SetNotificationsReadDocument = gql`
    mutation setNotificationsRead($ids: [String!]!) {
  setNotificationsRead(input: $ids) {
    success
    message
  }
}
    `;
export const GetOrganizationDocument = gql`
    query getOrganization {
  getOrganization {
    ...organizationFragment
  }
}
    ${OrganizationFragmentFragmentDoc}`;
export const UpdateOrganizationDocument = gql`
    mutation updateOrganization($input: UpdateOrganization!) {
  updateOrganization(input: $input) {
    status {
      ...statusFragment
    }
    organization {
      ...organizationFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${OrganizationFragmentFragmentDoc}`;
export const GetPortfolioDocument = gql`
    query getPortfolio {
  getPortfolio {
    ...portfolioFragment
  }
}
    ${PortfolioFragmentFragmentDoc}`;
export const GetDraftPortfolioDocument = gql`
    query getDraftPortfolio($additionalID: String!) {
  getDraftPortfolio(additionalID: $additionalID) {
    ...portfolioFragment
  }
}
    ${PortfolioFragmentFragmentDoc}`;
export const GetPortfolioForResourceDocument = gql`
    query getPortfolioForResource($resourceID: String!) {
  getPortfolioForResource(resourceID: $resourceID) {
    ...portfolioFragment
  }
}
    ${PortfolioFragmentFragmentDoc}`;
export const FindProjectsDocument = gql`
    query findProjects($input: PageAndFilter!) {
  findProjects(pageAndFilter: $input) {
    paging {
      ...pagingFragment
    }
    results {
      meta {
        ...baseModelFragment
      }
      data {
        ...projectFragment
      }
    }
  }
}
    ${PagingFragmentFragmentDoc}
${BaseModelFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const GetProjectDocument = gql`
    query getProject($id: String!) {
  getProject(id: $id) {
    meta {
      ...baseModelFragment
    }
    data {
      ...projectFragment
    }
  }
}
    ${BaseModelFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const CheckProjectStatusDocument = gql`
    query checkProjectStatus($projectID: String!, $newStatus: String!) {
  checkProjectStatus(projectID: $projectID, newStatus: $newStatus) {
    ...validationResultFragment
  }
}
    ${ValidationResultFragmentFragmentDoc}`;
export const CreateProjectDocument = gql`
    mutation createProject($project: UpdateNewProject!) {
  createProject(input: $project) {
    status {
      ...statusFragment
    }
    project {
      ...projectFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const UpdateProjectDocument = gql`
    mutation updateProject($project: UpdateProject!) {
  updateProject(input: $project) {
    status {
      ...statusFragment
    }
    project {
      ...projectFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const DeleteProjectDocument = gql`
    mutation deleteProject($id: String!) {
  deleteProject(id: $id) {
    ...statusFragment
  }
}
    ${StatusFragmentFragmentDoc}`;
export const ToggleEmoteDocument = gql`
    mutation toggleEmote($input: UpdateCommentEmote!) {
  toggleEmote(input: $input) {
    ...statusFragment
  }
}
    ${StatusFragmentFragmentDoc}`;
export const SetProjectMilestonesFromTemplateDocument = gql`
    mutation setProjectMilestonesFromTemplate($input: UpdateProjectMilestoneTemplate!) {
  setProjectMilestonesFromTemplate(input: $input) {
    status {
      ...statusFragment
    }
    project {
      ...projectFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const UpdateProjectTaskDocument = gql`
    mutation updateProjectTask($input: UpdateProjectMilestoneTask!) {
  updateProjectTask(input: $input) {
    status {
      ...statusFragment
    }
    project {
      ...projectFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const UpdateProjectTaskStatusDocument = gql`
    mutation updateProjectTaskStatus($projectID: String!, $milestoneID: String!, $taskID: String!, $status: String!) {
  updateProjectTaskStatus(
    projectID: $projectID
    milestoneID: $milestoneID
    taskID: $taskID
    status: $status
  ) {
    status {
      ...statusFragment
    }
    project {
      ...projectFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const DeleteProjectTaskDocument = gql`
    mutation deleteProjectTask($projectID: String!, $milestoneID: String!, $taskID: String!) {
  deleteProjectTask(
    projectID: $projectID
    milestoneID: $milestoneID
    taskID: $taskID
  ) {
    status {
      ...statusFragment
    }
    project {
      ...projectFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const UpdateProjectFeatureDocument = gql`
    mutation updateProjectFeature($input: UpdateProjectFeature!) {
  updateProjectFeature(input: $input) {
    status {
      ...statusFragment
    }
    project {
      ...projectFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const DeleteProjectFeatureDocument = gql`
    mutation deleteProjectFeature($projectID: String!, $featureID: String!) {
  deleteProjectFeature(projectID: $projectID, featureID: $featureID) {
    status {
      ...statusFragment
    }
    project {
      ...projectFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const UpdateProjectValueLineDocument = gql`
    mutation updateProjectValueLine($input: UpdateProjectValueLine!) {
  updateProjectValueLine(input: $input) {
    status {
      ...statusFragment
    }
    project {
      ...projectFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const DeleteProjectValueLineDocument = gql`
    mutation deleteProjectValueLine($projectID: String!, $valueLineID: String!) {
  deleteProjectValueLine(projectID: $projectID, valueLineID: $valueLineID) {
    status {
      ...statusFragment
    }
    project {
      ...projectFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const SetProjectStatusDocument = gql`
    mutation setProjectStatus($projectID: String!, $newStatus: String!) {
  setProjectStatus(projectID: $projectID, newStatus: $newStatus) {
    status {
      ...statusFragment
    }
    project {
      ...projectFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const UpdateResourceSkillDocument = gql`
    mutation updateResourceSkill($input: UpdateSkill!) {
  updateResourceSkill(input: $input) {
    success
    message
  }
}
    `;
export const DeleteResourceSkillDocument = gql`
    mutation deleteResourceSkill($resourceID: String!, $skillID: String!) {
  deleteResourceSkill(resourceID: $resourceID, skillID: $skillID) {
    success
    message
  }
}
    `;
export const CreateResourceDocument = gql`
    mutation createResource($input: UpdateResource!) {
  createResource(input: $input) {
    status {
      ...statusFragment
    }
    resource {
      ...resourceFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ResourceFragmentFragmentDoc}`;
export const UpdateResourceDocument = gql`
    mutation updateResource($input: UpdateResource!) {
  updateResource(input: $input) {
    status {
      ...statusFragment
    }
    resource {
      ...resourceFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ResourceFragmentFragmentDoc}`;
export const DeleteResourceDocument = gql`
    mutation deleteResource($id: String!) {
  deleteResource(id: $id) {
    ...statusFragment
  }
}
    ${StatusFragmentFragmentDoc}`;
export const GetResourceDocument = gql`
    query getResource($id: String!) {
  getResource(id: $id) {
    meta {
      ...baseModelFragment
    }
    data {
      ...resourceFragment
    }
  }
}
    ${BaseModelFragmentFragmentDoc}
${ResourceFragmentFragmentDoc}`;
export const FindAllResourcesDocument = gql`
    query findAllResources {
  findAllResources {
    paging {
      ...pagingFragment
    }
    results {
      meta {
        ...baseModelFragment
      }
      data {
        ...resourceFragment
      }
    }
  }
}
    ${PagingFragmentFragmentDoc}
${BaseModelFragmentFragmentDoc}
${ResourceFragmentFragmentDoc}`;
export const FindResourcesDocument = gql`
    query findResources($input: PageAndFilter!) {
  findResources(pageAndFilter: $input) {
    paging {
      ...pagingFragment
    }
    results {
      meta {
        ...baseModelFragment
      }
      data {
        ...resourceFragment
      }
    }
  }
}
    ${PagingFragmentFragmentDoc}
${BaseModelFragmentFragmentDoc}
${ResourceFragmentFragmentDoc}`;
export const FindAllRolesDocument = gql`
    query findAllRoles {
  findAllRoles {
    paging {
      ...pagingFragment
    }
    results {
      meta {
        ...baseModelFragment
      }
      data {
        ...roleFragment
      }
    }
  }
}
    ${PagingFragmentFragmentDoc}
${BaseModelFragmentFragmentDoc}
${RoleFragmentFragmentDoc}`;
export const UpdateRoleDocument = gql`
    mutation updateRole($input: UpdateRole!) {
  updateRole(input: $input) {
    status {
      ...statusFragment
    }
    role {
      ...roleFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${RoleFragmentFragmentDoc}`;
export const UpdateAllRolesDocument = gql`
    mutation updateAllRoles($input: [UpdateRole]!) {
  updateAllRoles(input: $input) {
    status {
      ...statusFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}`;
export const DeleteRroleDocument = gql`
    mutation deleteRrole($id: String!) {
  deleteRole(id: $id) {
    ...statusFragment
  }
}
    ${StatusFragmentFragmentDoc}`;
export const CalculateProjectScheduleDocument = gql`
    query calculateProjectSchedule($projectID: String!, $startDate: Time!) {
  calculateProjectSchedule(projectID: $projectID, startDate: $startDate) {
    ...scheduleResultFragment
  }
}
    ${ScheduleResultFragmentFragmentDoc}`;
export const FindAllProjectTemplatesDocument = gql`
    query findAllProjectTemplates {
  findAllProjectTemplates {
    results {
      ...projectTemplateFragment
    }
  }
}
    ${ProjectTemplateFragmentFragmentDoc}`;
export const UpdateProjectTemplateDocument = gql`
    mutation updateProjectTemplate($input: UpdateProjecttemplate!) {
  updateProjectTemplate(input: $input) {
    status {
      ...statusFragment
    }
    template {
      ...projectTemplateFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${ProjectTemplateFragmentFragmentDoc}`;
export const DeleteProjectTemplateDocument = gql`
    mutation deleteProjectTemplate($id: String!) {
  deleteProjectTemplate(id: $id) {
    ...statusFragment
  }
}
    ${StatusFragmentFragmentDoc}`;
export const FindAllUsersDocument = gql`
    query findAllUsers {
  findAllUsers {
    paging {
      ...pagingFragment
    }
    results {
      ...userFragment
    }
  }
}
    ${PagingFragmentFragmentDoc}
${UserFragmentFragmentDoc}`;
export const CurrentUserDocument = gql`
    query currentUser {
  currentUser {
    ...userFragment
  }
}
    ${UserFragmentFragmentDoc}`;
export const UpdateUserDocument = gql`
    mutation updateUser($input: UpdateUser!) {
  updateUser(input: $input) {
    status {
      ...statusFragment
    }
    user {
      ...userFragment
    }
  }
}
    ${StatusFragmentFragmentDoc}
${UserFragmentFragmentDoc}`;
export type Requester<C = {}> = <R, V>(doc: DocumentNode, vars?: V, options?: C) => Promise<R> | AsyncIterable<R>
export function getSdk<C>(requester: Requester<C>) {
  return {
    findActivity(variables: FindActivityQueryVariables, options?: C): Promise<FindActivityQuery> {
      return requester<FindActivityQuery, FindActivityQueryVariables>(FindActivityDocument, variables, options) as Promise<FindActivityQuery>;
    },
    createProjectComment(variables: CreateProjectCommentMutationVariables, options?: C): Promise<CreateProjectCommentMutation> {
      return requester<CreateProjectCommentMutation, CreateProjectCommentMutationVariables>(CreateProjectCommentDocument, variables, options) as Promise<CreateProjectCommentMutation>;
    },
    updateProjectComment(variables: UpdateProjectCommentMutationVariables, options?: C): Promise<UpdateProjectCommentMutation> {
      return requester<UpdateProjectCommentMutation, UpdateProjectCommentMutationVariables>(UpdateProjectCommentDocument, variables, options) as Promise<UpdateProjectCommentMutation>;
    },
    createProjectCommentReply(variables: CreateProjectCommentReplyMutationVariables, options?: C): Promise<CreateProjectCommentReplyMutation> {
      return requester<CreateProjectCommentReplyMutation, CreateProjectCommentReplyMutationVariables>(CreateProjectCommentReplyDocument, variables, options) as Promise<CreateProjectCommentReplyMutation>;
    },
    deleteProjectComment(variables: DeleteProjectCommentMutationVariables, options?: C): Promise<DeleteProjectCommentMutation> {
      return requester<DeleteProjectCommentMutation, DeleteProjectCommentMutationVariables>(DeleteProjectCommentDocument, variables, options) as Promise<DeleteProjectCommentMutation>;
    },
    findProjectComments(variables: FindProjectCommentsQueryVariables, options?: C): Promise<FindProjectCommentsQuery> {
      return requester<FindProjectCommentsQuery, FindProjectCommentsQueryVariables>(FindProjectCommentsDocument, variables, options) as Promise<FindProjectCommentsQuery>;
    },
    getCommentThread(variables: GetCommentThreadQueryVariables, options?: C): Promise<GetCommentThreadQuery> {
      return requester<GetCommentThreadQuery, GetCommentThreadQueryVariables>(GetCommentThreadDocument, variables, options) as Promise<GetCommentThreadQuery>;
    },
    findAllLists(variables?: FindAllListsQueryVariables, options?: C): Promise<FindAllListsQuery> {
      return requester<FindAllListsQuery, FindAllListsQueryVariables>(FindAllListsDocument, variables, options) as Promise<FindAllListsQuery>;
    },
    getList(variables: GetListQueryVariables, options?: C): Promise<GetListQuery> {
      return requester<GetListQuery, GetListQueryVariables>(GetListDocument, variables, options) as Promise<GetListQuery>;
    },
    updateList(variables: UpdateListMutationVariables, options?: C): Promise<UpdateListMutation> {
      return requester<UpdateListMutation, UpdateListMutationVariables>(UpdateListDocument, variables, options) as Promise<UpdateListMutation>;
    },
    findUserNotifications(variables: FindUserNotificationsQueryVariables, options?: C): Promise<FindUserNotificationsQuery> {
      return requester<FindUserNotificationsQuery, FindUserNotificationsQueryVariables>(FindUserNotificationsDocument, variables, options) as Promise<FindUserNotificationsQuery>;
    },
    setNotificationsRead(variables: SetNotificationsReadMutationVariables, options?: C): Promise<SetNotificationsReadMutation> {
      return requester<SetNotificationsReadMutation, SetNotificationsReadMutationVariables>(SetNotificationsReadDocument, variables, options) as Promise<SetNotificationsReadMutation>;
    },
    getOrganization(variables?: GetOrganizationQueryVariables, options?: C): Promise<GetOrganizationQuery> {
      return requester<GetOrganizationQuery, GetOrganizationQueryVariables>(GetOrganizationDocument, variables, options) as Promise<GetOrganizationQuery>;
    },
    updateOrganization(variables: UpdateOrganizationMutationVariables, options?: C): Promise<UpdateOrganizationMutation> {
      return requester<UpdateOrganizationMutation, UpdateOrganizationMutationVariables>(UpdateOrganizationDocument, variables, options) as Promise<UpdateOrganizationMutation>;
    },
    getPortfolio(variables?: GetPortfolioQueryVariables, options?: C): Promise<GetPortfolioQuery> {
      return requester<GetPortfolioQuery, GetPortfolioQueryVariables>(GetPortfolioDocument, variables, options) as Promise<GetPortfolioQuery>;
    },
    getDraftPortfolio(variables: GetDraftPortfolioQueryVariables, options?: C): Promise<GetDraftPortfolioQuery> {
      return requester<GetDraftPortfolioQuery, GetDraftPortfolioQueryVariables>(GetDraftPortfolioDocument, variables, options) as Promise<GetDraftPortfolioQuery>;
    },
    getPortfolioForResource(variables: GetPortfolioForResourceQueryVariables, options?: C): Promise<GetPortfolioForResourceQuery> {
      return requester<GetPortfolioForResourceQuery, GetPortfolioForResourceQueryVariables>(GetPortfolioForResourceDocument, variables, options) as Promise<GetPortfolioForResourceQuery>;
    },
    findProjects(variables: FindProjectsQueryVariables, options?: C): Promise<FindProjectsQuery> {
      return requester<FindProjectsQuery, FindProjectsQueryVariables>(FindProjectsDocument, variables, options) as Promise<FindProjectsQuery>;
    },
    getProject(variables: GetProjectQueryVariables, options?: C): Promise<GetProjectQuery> {
      return requester<GetProjectQuery, GetProjectQueryVariables>(GetProjectDocument, variables, options) as Promise<GetProjectQuery>;
    },
    checkProjectStatus(variables: CheckProjectStatusQueryVariables, options?: C): Promise<CheckProjectStatusQuery> {
      return requester<CheckProjectStatusQuery, CheckProjectStatusQueryVariables>(CheckProjectStatusDocument, variables, options) as Promise<CheckProjectStatusQuery>;
    },
    createProject(variables: CreateProjectMutationVariables, options?: C): Promise<CreateProjectMutation> {
      return requester<CreateProjectMutation, CreateProjectMutationVariables>(CreateProjectDocument, variables, options) as Promise<CreateProjectMutation>;
    },
    updateProject(variables: UpdateProjectMutationVariables, options?: C): Promise<UpdateProjectMutation> {
      return requester<UpdateProjectMutation, UpdateProjectMutationVariables>(UpdateProjectDocument, variables, options) as Promise<UpdateProjectMutation>;
    },
    deleteProject(variables: DeleteProjectMutationVariables, options?: C): Promise<DeleteProjectMutation> {
      return requester<DeleteProjectMutation, DeleteProjectMutationVariables>(DeleteProjectDocument, variables, options) as Promise<DeleteProjectMutation>;
    },
    toggleEmote(variables: ToggleEmoteMutationVariables, options?: C): Promise<ToggleEmoteMutation> {
      return requester<ToggleEmoteMutation, ToggleEmoteMutationVariables>(ToggleEmoteDocument, variables, options) as Promise<ToggleEmoteMutation>;
    },
    setProjectMilestonesFromTemplate(variables: SetProjectMilestonesFromTemplateMutationVariables, options?: C): Promise<SetProjectMilestonesFromTemplateMutation> {
      return requester<SetProjectMilestonesFromTemplateMutation, SetProjectMilestonesFromTemplateMutationVariables>(SetProjectMilestonesFromTemplateDocument, variables, options) as Promise<SetProjectMilestonesFromTemplateMutation>;
    },
    updateProjectTask(variables: UpdateProjectTaskMutationVariables, options?: C): Promise<UpdateProjectTaskMutation> {
      return requester<UpdateProjectTaskMutation, UpdateProjectTaskMutationVariables>(UpdateProjectTaskDocument, variables, options) as Promise<UpdateProjectTaskMutation>;
    },
    updateProjectTaskStatus(variables: UpdateProjectTaskStatusMutationVariables, options?: C): Promise<UpdateProjectTaskStatusMutation> {
      return requester<UpdateProjectTaskStatusMutation, UpdateProjectTaskStatusMutationVariables>(UpdateProjectTaskStatusDocument, variables, options) as Promise<UpdateProjectTaskStatusMutation>;
    },
    deleteProjectTask(variables: DeleteProjectTaskMutationVariables, options?: C): Promise<DeleteProjectTaskMutation> {
      return requester<DeleteProjectTaskMutation, DeleteProjectTaskMutationVariables>(DeleteProjectTaskDocument, variables, options) as Promise<DeleteProjectTaskMutation>;
    },
    updateProjectFeature(variables: UpdateProjectFeatureMutationVariables, options?: C): Promise<UpdateProjectFeatureMutation> {
      return requester<UpdateProjectFeatureMutation, UpdateProjectFeatureMutationVariables>(UpdateProjectFeatureDocument, variables, options) as Promise<UpdateProjectFeatureMutation>;
    },
    deleteProjectFeature(variables: DeleteProjectFeatureMutationVariables, options?: C): Promise<DeleteProjectFeatureMutation> {
      return requester<DeleteProjectFeatureMutation, DeleteProjectFeatureMutationVariables>(DeleteProjectFeatureDocument, variables, options) as Promise<DeleteProjectFeatureMutation>;
    },
    updateProjectValueLine(variables: UpdateProjectValueLineMutationVariables, options?: C): Promise<UpdateProjectValueLineMutation> {
      return requester<UpdateProjectValueLineMutation, UpdateProjectValueLineMutationVariables>(UpdateProjectValueLineDocument, variables, options) as Promise<UpdateProjectValueLineMutation>;
    },
    deleteProjectValueLine(variables: DeleteProjectValueLineMutationVariables, options?: C): Promise<DeleteProjectValueLineMutation> {
      return requester<DeleteProjectValueLineMutation, DeleteProjectValueLineMutationVariables>(DeleteProjectValueLineDocument, variables, options) as Promise<DeleteProjectValueLineMutation>;
    },
    setProjectStatus(variables: SetProjectStatusMutationVariables, options?: C): Promise<SetProjectStatusMutation> {
      return requester<SetProjectStatusMutation, SetProjectStatusMutationVariables>(SetProjectStatusDocument, variables, options) as Promise<SetProjectStatusMutation>;
    },
    updateResourceSkill(variables: UpdateResourceSkillMutationVariables, options?: C): Promise<UpdateResourceSkillMutation> {
      return requester<UpdateResourceSkillMutation, UpdateResourceSkillMutationVariables>(UpdateResourceSkillDocument, variables, options) as Promise<UpdateResourceSkillMutation>;
    },
    deleteResourceSkill(variables: DeleteResourceSkillMutationVariables, options?: C): Promise<DeleteResourceSkillMutation> {
      return requester<DeleteResourceSkillMutation, DeleteResourceSkillMutationVariables>(DeleteResourceSkillDocument, variables, options) as Promise<DeleteResourceSkillMutation>;
    },
    createResource(variables: CreateResourceMutationVariables, options?: C): Promise<CreateResourceMutation> {
      return requester<CreateResourceMutation, CreateResourceMutationVariables>(CreateResourceDocument, variables, options) as Promise<CreateResourceMutation>;
    },
    updateResource(variables: UpdateResourceMutationVariables, options?: C): Promise<UpdateResourceMutation> {
      return requester<UpdateResourceMutation, UpdateResourceMutationVariables>(UpdateResourceDocument, variables, options) as Promise<UpdateResourceMutation>;
    },
    deleteResource(variables: DeleteResourceMutationVariables, options?: C): Promise<DeleteResourceMutation> {
      return requester<DeleteResourceMutation, DeleteResourceMutationVariables>(DeleteResourceDocument, variables, options) as Promise<DeleteResourceMutation>;
    },
    getResource(variables: GetResourceQueryVariables, options?: C): Promise<GetResourceQuery> {
      return requester<GetResourceQuery, GetResourceQueryVariables>(GetResourceDocument, variables, options) as Promise<GetResourceQuery>;
    },
    findAllResources(variables?: FindAllResourcesQueryVariables, options?: C): Promise<FindAllResourcesQuery> {
      return requester<FindAllResourcesQuery, FindAllResourcesQueryVariables>(FindAllResourcesDocument, variables, options) as Promise<FindAllResourcesQuery>;
    },
    findResources(variables: FindResourcesQueryVariables, options?: C): Promise<FindResourcesQuery> {
      return requester<FindResourcesQuery, FindResourcesQueryVariables>(FindResourcesDocument, variables, options) as Promise<FindResourcesQuery>;
    },
    findAllRoles(variables?: FindAllRolesQueryVariables, options?: C): Promise<FindAllRolesQuery> {
      return requester<FindAllRolesQuery, FindAllRolesQueryVariables>(FindAllRolesDocument, variables, options) as Promise<FindAllRolesQuery>;
    },
    updateRole(variables: UpdateRoleMutationVariables, options?: C): Promise<UpdateRoleMutation> {
      return requester<UpdateRoleMutation, UpdateRoleMutationVariables>(UpdateRoleDocument, variables, options) as Promise<UpdateRoleMutation>;
    },
    updateAllRoles(variables: UpdateAllRolesMutationVariables, options?: C): Promise<UpdateAllRolesMutation> {
      return requester<UpdateAllRolesMutation, UpdateAllRolesMutationVariables>(UpdateAllRolesDocument, variables, options) as Promise<UpdateAllRolesMutation>;
    },
    deleteRrole(variables: DeleteRroleMutationVariables, options?: C): Promise<DeleteRroleMutation> {
      return requester<DeleteRroleMutation, DeleteRroleMutationVariables>(DeleteRroleDocument, variables, options) as Promise<DeleteRroleMutation>;
    },
    calculateProjectSchedule(variables: CalculateProjectScheduleQueryVariables, options?: C): Promise<CalculateProjectScheduleQuery> {
      return requester<CalculateProjectScheduleQuery, CalculateProjectScheduleQueryVariables>(CalculateProjectScheduleDocument, variables, options) as Promise<CalculateProjectScheduleQuery>;
    },
    findAllProjectTemplates(variables?: FindAllProjectTemplatesQueryVariables, options?: C): Promise<FindAllProjectTemplatesQuery> {
      return requester<FindAllProjectTemplatesQuery, FindAllProjectTemplatesQueryVariables>(FindAllProjectTemplatesDocument, variables, options) as Promise<FindAllProjectTemplatesQuery>;
    },
    updateProjectTemplate(variables: UpdateProjectTemplateMutationVariables, options?: C): Promise<UpdateProjectTemplateMutation> {
      return requester<UpdateProjectTemplateMutation, UpdateProjectTemplateMutationVariables>(UpdateProjectTemplateDocument, variables, options) as Promise<UpdateProjectTemplateMutation>;
    },
    deleteProjectTemplate(variables: DeleteProjectTemplateMutationVariables, options?: C): Promise<DeleteProjectTemplateMutation> {
      return requester<DeleteProjectTemplateMutation, DeleteProjectTemplateMutationVariables>(DeleteProjectTemplateDocument, variables, options) as Promise<DeleteProjectTemplateMutation>;
    },
    findAllUsers(variables?: FindAllUsersQueryVariables, options?: C): Promise<FindAllUsersQuery> {
      return requester<FindAllUsersQuery, FindAllUsersQueryVariables>(FindAllUsersDocument, variables, options) as Promise<FindAllUsersQuery>;
    },
    currentUser(variables?: CurrentUserQueryVariables, options?: C): Promise<CurrentUserQuery> {
      return requester<CurrentUserQuery, CurrentUserQueryVariables>(CurrentUserDocument, variables, options) as Promise<CurrentUserQuery>;
    },
    updateUser(variables: UpdateUserMutationVariables, options?: C): Promise<UpdateUserMutation> {
      return requester<UpdateUserMutation, UpdateUserMutationVariables>(UpdateUserDocument, variables, options) as Promise<UpdateUserMutation>;
    }
  };
}
export type Sdk = ReturnType<typeof getSdk>;