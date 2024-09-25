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
  context: Scalars['String']['output'];
  detail?: Maybe<Scalars['String']['output']>;
  id: Scalars['String']['output'];
  resource?: Maybe<Resource>;
  summary: Scalars['String']['output'];
  targetID?: Maybe<Scalars['String']['output']>;
  type: Scalars['String']['output'];
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

export type Comment = {
  __typename?: 'Comment';
  acknowledges?: Maybe<Array<Scalars['String']['output']>>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['String']['output'];
  dislikes?: Maybe<Array<Scalars['String']['output']>>;
  id: Scalars['String']['output'];
  isEdited: Scalars['Boolean']['output'];
  laughsAt?: Maybe<Array<Scalars['String']['output']>>;
  likes?: Maybe<Array<Scalars['String']['output']>>;
  loves?: Maybe<Array<Scalars['String']['output']>>;
  projectId: Scalars['String']['output'];
  replies?: Maybe<Array<Comment>>;
  text: Scalars['String']['output'];
  updatedAt: Scalars['Time']['output'];
  user: User;
};

export type CommentResults = {
  __typename?: 'CommentResults';
  filters: Filters;
  paging: Pagination;
  results?: Maybe<Array<Maybe<Comment>>>;
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

export type CreateResourceResult = {
  __typename?: 'CreateResourceResult';
  resource?: Maybe<Resource>;
  status: Status;
};

export type CreateUserResult = {
  __typename?: 'CreateUserResult';
  status?: Maybe<Status>;
  user?: Maybe<User>;
};

export type Filter = {
  __typename?: 'Filter';
  key: Scalars['String']['output'];
  operation: Scalars['String']['output'];
  value: Scalars['String']['output'];
};

export type Filters = {
  __typename?: 'Filters';
  filters?: Maybe<Array<Maybe<Filter>>>;
};

export type InputFilter = {
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
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  values: Array<ListItem>;
};

export type ListItem = {
  __typename?: 'ListItem';
  name: Scalars['String']['output'];
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
  createOrganization: CreateOrganizationResult;
  createProject: CreateProjectResult;
  createProjectComment: CreateProjectCommentResult;
  createProjectCommentReply: CreateProjectCommentResult;
  createResource: CreateResourceResult;
  createUser: CreateUserResult;
  deleteProject: Status;
  deleteProjectComment: Status;
  deleteProjectFeature: CreateProjectResult;
  deleteProjectTask: CreateProjectResult;
  deleteResource: Status;
  deleteResourceSkill: Status;
  setNotificationsRead: Status;
  setProjectMilestonesFromTemplate: CreateProjectResult;
  toggleEmote: Status;
  updateOrganization: CreateOrganizationResult;
  updateProject: CreateProjectResult;
  updateProjectComment: CreateProjectCommentResult;
  updateProjectFeature: CreateProjectResult;
  updateProjectTask: CreateProjectResult;
  updateResource: CreateResourceResult;
  updateResourceSkill: Status;
  updateUser: CreateUserResult;
};


export type MutationCreateOrganizationArgs = {
  input: UpdateOrganization;
};


export type MutationCreateProjectArgs = {
  input: UpdateProject;
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


export type MutationCreateUserArgs = {
  input: UpdateUser;
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


export type MutationDeleteResourceArgs = {
  id: Scalars['String']['input'];
};


export type MutationDeleteResourceSkillArgs = {
  resourceID: Scalars['String']['input'];
  skillID: Scalars['String']['input'];
};


export type MutationSetNotificationsReadArgs = {
  input: Array<Scalars['String']['input']>;
};


export type MutationSetProjectMilestonesFromTemplateArgs = {
  input?: InputMaybe<UpdateProjectMilestoneTemplate>;
};


export type MutationToggleEmoteArgs = {
  input: UpdateCommentEmote;
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


export type MutationUpdateResourceArgs = {
  input: UpdateResource;
};


export type MutationUpdateResourceSkillArgs = {
  input: UpdateSkill;
};


export type MutationUpdateUserArgs = {
  input: UpdateUser;
};

export type Notification = {
  __typename?: 'Notification';
  contextId: Scalars['String']['output'];
  createdAt?: Maybe<Scalars['Time']['output']>;
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
};

export type OrganizationDefaults = {
  __typename?: 'OrganizationDefaults';
  discountRate: Scalars['Float']['output'];
  hoursPerWeek: Scalars['Int']['output'];
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

export type PortfolioSnapshot = {
  __typename?: 'PortfolioSnapshot';
  projects: Array<SnapshotProject>;
};

export type Project = {
  __typename?: 'Project';
  createdAt?: Maybe<Scalars['Time']['output']>;
  createdBy?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  projectBasics: ProjectBasics;
  projectCost: ProjectCost;
  projectDaci: ProjectDaci;
  projectFeatures?: Maybe<Array<ProjectFeature>>;
  projectMilestones?: Maybe<Array<ProjectMilestone>>;
  projectValue: ProjectValue;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['String']['output']>;
};

export type ProjectBasics = {
  __typename?: 'ProjectBasics';
  description: Scalars['String']['output'];
  name: Scalars['String']['output'];
  ownerID?: Maybe<Scalars['String']['output']>;
  startDate?: Maybe<Scalars['Time']['output']>;
  status: Scalars['String']['output'];
};

export type ProjectCost = {
  __typename?: 'ProjectCost';
  blendedRate?: Maybe<Scalars['Float']['output']>;
  hourEstimate: Scalars['Int']['output'];
  initialCost?: Maybe<Scalars['Float']['output']>;
  ongoing?: Maybe<Scalars['Float']['output']>;
};

export type ProjectDaci = {
  __typename?: 'ProjectDaci';
  approver?: Maybe<Array<Maybe<Resource>>>;
  contributor?: Maybe<Array<Maybe<Resource>>>;
  driver?: Maybe<Array<Maybe<Resource>>>;
  informed?: Maybe<Array<Maybe<Resource>>>;
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
  endDate?: Maybe<Scalars['Time']['output']>;
  id: Scalars['String']['output'];
  phase: ProjectMilestonePhase;
  startDate?: Maybe<Scalars['Time']['output']>;
  tasks: Array<ProjectMilestoneTask>;
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
  description?: Maybe<Scalars['String']['output']>;
  endDate?: Maybe<Scalars['Time']['output']>;
  hourEstimate: Scalars['Int']['output'];
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  requiredSkillIDs?: Maybe<Array<Scalars['String']['output']>>;
  resourceIDs?: Maybe<Array<Scalars['String']['output']>>;
  resources?: Maybe<Array<Resource>>;
  skills?: Maybe<Array<Skill>>;
  startDate?: Maybe<Scalars['Time']['output']>;
  status: Scalars['String']['output'];
};

export type ProjectResults = {
  __typename?: 'ProjectResults';
  filters: Filters;
  paging?: Maybe<Pagination>;
  results?: Maybe<Array<Project>>;
};

export type ProjectValue = {
  __typename?: 'ProjectValue';
  discountRate?: Maybe<Scalars['Float']['output']>;
  fundingSource?: Maybe<Scalars['String']['output']>;
  internalRateOfReturn?: Maybe<Scalars['Float']['output']>;
  netPresentValue?: Maybe<Scalars['Float']['output']>;
  yearFiveValue?: Maybe<Scalars['Float']['output']>;
  yearFourValue?: Maybe<Scalars['Float']['output']>;
  yearOneValue?: Maybe<Scalars['Float']['output']>;
  yearThreeValue?: Maybe<Scalars['Float']['output']>;
  yearTwoValue?: Maybe<Scalars['Float']['output']>;
};

export type Projecttemplate = {
  __typename?: 'Projecttemplate';
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
};

export type ProjecttemplateResults = {
  __typename?: 'ProjecttemplateResults';
  filters: Filters;
  paging?: Maybe<Pagination>;
  results?: Maybe<Array<Projecttemplate>>;
};

export type Query = {
  __typename?: 'Query';
  currentUser?: Maybe<User>;
  findActivity: ActivityResults;
  findAllLists: ListResults;
  findAllProjectTemplates: ProjecttemplateResults;
  findAllResources: ResourceResults;
  findAllUsers: UserResults;
  findProjectComments: CommentResults;
  findProjects: ProjectResults;
  findResources: ResourceResults;
  findUserNotifications: NotificationResults;
  getCommentThread: Comment;
  getList?: Maybe<List>;
  getOrganization: Organization;
  getProject: Project;
  getResource: Resource;
  getUser: User;
  portfolioSnapshot: PortfolioSnapshot;
  resourceSnapshot: ResourceSnapshot;
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


export type QueryGetListArgs = {
  nameOrID: Scalars['String']['input'];
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
  createdAt?: Maybe<Scalars['Time']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  isBot: Scalars['Boolean']['output'];
  name: Scalars['String']['output'];
  profileImage?: Maybe<Scalars['String']['output']>;
  role: Scalars['String']['output'];
  skills?: Maybe<Array<Skill>>;
  type: Scalars['String']['output'];
  user?: Maybe<User>;
};

export type ResourceResults = {
  __typename?: 'ResourceResults';
  filters: Filters;
  paging?: Maybe<Pagination>;
  results?: Maybe<Array<Resource>>;
};

export type ResourceSnapshot = {
  __typename?: 'ResourceSnapshot';
  scheduledResources: Array<ScheduledResource>;
};

export type ScheduledResource = {
  __typename?: 'ScheduledResource';
  milestoneEndDate?: Maybe<Scalars['Time']['output']>;
  milestoneName: Scalars['String']['output'];
  milestoneStartDate?: Maybe<Scalars['Time']['output']>;
  projectId: Scalars['String']['output'];
  projectName: Scalars['String']['output'];
  projectStatus: Scalars['String']['output'];
  resourceId: Scalars['String']['output'];
  resourceName: Scalars['String']['output'];
  taskEndDate?: Maybe<Scalars['Time']['output']>;
  taskHourEstimate: Scalars['Int']['output'];
  taskName: Scalars['String']['output'];
  taskStartDate?: Maybe<Scalars['Time']['output']>;
};

export type Skill = {
  __typename?: 'Skill';
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  proficiency?: Maybe<Scalars['Float']['output']>;
};

export type SnapshotProject = {
  __typename?: 'SnapshotProject';
  endDate?: Maybe<Scalars['Time']['output']>;
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  npv: Scalars['Float']['output'];
  startDate?: Maybe<Scalars['Time']['output']>;
  status: Scalars['String']['output'];
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
};

export type UpdateCommentReply = {
  parentCommentID: Scalars['String']['input'];
  text: Scalars['String']['input'];
};

export type UpdateOrganization = {
  defaults: UpdateOrganizationDefaults;
  id?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
};

export type UpdateOrganizationDefaults = {
  discountRate: Scalars['Float']['input'];
  hoursPerWeek: Scalars['Int']['input'];
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
  name: Scalars['String']['input'];
  ownerID?: InputMaybe<Scalars['String']['input']>;
  startDate?: InputMaybe<Scalars['Time']['input']>;
  status?: InputMaybe<Scalars['String']['input']>;
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
  startDate?: InputMaybe<Scalars['Time']['input']>;
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
  endDate?: InputMaybe<Scalars['Time']['input']>;
  hourEstimate: Scalars['Int']['input'];
  id?: InputMaybe<Scalars['String']['input']>;
  milestoneID: Scalars['String']['input'];
  name: Scalars['String']['input'];
  projectID: Scalars['String']['input'];
  requiredSkillIDs?: InputMaybe<Array<Scalars['String']['input']>>;
  resourceIDs?: InputMaybe<Array<Scalars['String']['input']>>;
  startDate?: InputMaybe<Scalars['Time']['input']>;
  status: Scalars['String']['input'];
};

export type UpdateProjectMilestoneTemplate = {
  projectId: Scalars['String']['input'];
  templateId: Scalars['String']['input'];
};

export type UpdateProjectValue = {
  discountRate?: InputMaybe<Scalars['Float']['input']>;
  fundingSource?: InputMaybe<Scalars['String']['input']>;
  yearFiveValue?: InputMaybe<Scalars['Float']['input']>;
  yearFourValue?: InputMaybe<Scalars['Float']['input']>;
  yearOneValue?: InputMaybe<Scalars['Float']['input']>;
  yearThreeValue?: InputMaybe<Scalars['Float']['input']>;
  yearTwoValue?: InputMaybe<Scalars['Float']['input']>;
};

export type UpdateResource = {
  email?: InputMaybe<Scalars['String']['input']>;
  id?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
  profileImage?: InputMaybe<Scalars['String']['input']>;
  role?: InputMaybe<Scalars['String']['input']>;
  skills?: InputMaybe<Array<UpdateSkill>>;
  type: Scalars['String']['input'];
  userID?: InputMaybe<Scalars['String']['input']>;
};

export type UpdateSkill = {
  id: Scalars['String']['input'];
  proficiency: Scalars['Float']['input'];
  resourceID: Scalars['String']['input'];
};

export type UpdateUser = {
  confirmPassword: Scalars['String']['input'];
  email: Scalars['String']['input'];
  name: Scalars['String']['input'];
  password: Scalars['String']['input'];
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
  firstName
  lastName
  email
  id
  profileImage
}
    `;
export const CommentFragmentFragmentDoc = gql`
    fragment commentFragment on Comment {
  id
  text
  isEdited
  projectId
  user {
    ...userFragment
  }
  createdAt
  updatedAt
  replies {
    id
    text
    isEdited
    user {
      ...userFragment
    }
    likes
    loves
    dislikes
    laughsAt
    acknowledges
  }
  likes
  loves
  dislikes
  laughsAt
  acknowledges
}
    ${UserFragmentFragmentDoc}`;
export const ControlFieldsFragmentFragmentDoc = gql`
    fragment controlFieldsFragment on ControlFields {
  createdBy
  createdAt
  updatedAt
  updatedBy
  deletedAt
  deletedBy
  createByUser {
    firstName
    lastName
    email
  }
  updateByUser {
    firstName
    lastName
    email
  }
  deleteByUser {
    firstName
    lastName
    email
  }
}
    `;
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
  updatedAt
}
    `;
export const ResourceFragmentFragmentDoc = gql`
    fragment resourceFragment on Resource {
  name
  id
  user {
    ...userFragment
  }
  role
  profileImage
  isBot
  type
  createdAt
  skills {
    id
    name
    proficiency
  }
}
    ${UserFragmentFragmentDoc}`;
export const ProjectFragmentFragmentDoc = gql`
    fragment projectFragment on Project {
  id
  createdAt
  updatedAt
  createdBy
  updatedBy
  projectBasics {
    name
    description
    status
    startDate
    ownerID
  }
  projectValue {
    fundingSource
    discountRate
    yearOneValue
    yearTwoValue
    yearThreeValue
    yearFourValue
    yearFiveValue
    netPresentValue
    internalRateOfReturn
  }
  projectCost {
    ongoing
    blendedRate
    initialCost
    hourEstimate
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
      requiredSkillIDs
      skills {
        id
        name
      }
      resourceIDs
      resources {
        ...resourceFragment
      }
      startDate
      endDate
    }
  }
}
    ${ResourceFragmentFragmentDoc}`;
export const FindActivityDocument = gql`
    query findActivity($input: PageAndFilter!) {
  findActivity(pageAndFilter: $input) {
    paging {
      ...pagingFragment
    }
    results {
      id
      summary
      detail
      context
      targetID
      type
      resource {
        name
        id
        profileImage
        user {
          email
        }
      }
    }
  }
}
    ${PagingFragmentFragmentDoc}`;
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
      ...commentFragment
    }
  }
}
    ${PagingFragmentFragmentDoc}
${CommentFragmentFragmentDoc}`;
export const GetCommentThreadDocument = gql`
    query getCommentThread($commentID: String!) {
  getCommentThread(id: $commentID) {
    ...commentFragment
  }
}
    ${CommentFragmentFragmentDoc}`;
export const FindAllListsDocument = gql`
    query findAllLists {
  findAllLists {
    paging {
      ...pagingFragment
    }
    results {
      id
      name
      values {
        value
        name
      }
    }
  }
}
    ${PagingFragmentFragmentDoc}`;
export const GetListDocument = gql`
    query getList($nameOrID: String!) {
  getList(nameOrID: $nameOrID) {
    id
    name
    values {
      value
      name
    }
  }
}
    `;
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
export const PortfolioSnapshotDocument = gql`
    query portfolioSnapshot {
  portfolioSnapshot {
    projects {
      id
      name
      npv
      status
      startDate
      endDate
    }
  }
}
    `;
export const ResourceSnapshotDocument = gql`
    query resourceSnapshot {
  resourceSnapshot {
    scheduledResources {
      projectId
      projectName
      projectStatus
      milestoneName
      milestoneStartDate
      milestoneEndDate
      taskName
      taskStartDate
      taskEndDate
      taskHourEstimate
      resourceId
      resourceName
    }
  }
}
    `;
export const FindProjectsDocument = gql`
    query findProjects($input: PageAndFilter!) {
  findProjects(pageAndFilter: $input) {
    paging {
      ...pagingFragment
    }
    results {
      ...projectFragment
    }
  }
}
    ${PagingFragmentFragmentDoc}
${ProjectFragmentFragmentDoc}`;
export const GetProjectDocument = gql`
    query getProject($id: String!) {
  getProject(id: $id) {
    ...projectFragment
  }
}
    ${ProjectFragmentFragmentDoc}`;
export const FindAllProjectTemplatesDocument = gql`
    query findAllProjectTemplates {
  findAllProjectTemplates {
    results {
      name
      id
      phases {
        id
        name
        order
        description
      }
    }
  }
}
    `;
export const CreateProjectDocument = gql`
    mutation createProject($project: UpdateProject!) {
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
export const CreateUserDocument = gql`
    mutation createUser($info: UpdateUser!) {
  createUser(input: $info) {
    status {
      success
      message
    }
  }
}
    `;
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
      success
      message
    }
    resource {
      ...resourceFragment
    }
  }
}
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
    success
    message
  }
}
    `;
export const GetResourceDocument = gql`
    query getResource($id: String!) {
  getResource(id: $id) {
    ...resourceFragment
  }
}
    ${ResourceFragmentFragmentDoc}`;
export const FindAllResourcesDocument = gql`
    query findAllResources {
  findAllResources {
    paging {
      ...pagingFragment
    }
    results {
      ...resourceFragment
    }
  }
}
    ${PagingFragmentFragmentDoc}
${ResourceFragmentFragmentDoc}`;
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
    findUserNotifications(variables: FindUserNotificationsQueryVariables, options?: C): Promise<FindUserNotificationsQuery> {
      return requester<FindUserNotificationsQuery, FindUserNotificationsQueryVariables>(FindUserNotificationsDocument, variables, options) as Promise<FindUserNotificationsQuery>;
    },
    setNotificationsRead(variables: SetNotificationsReadMutationVariables, options?: C): Promise<SetNotificationsReadMutation> {
      return requester<SetNotificationsReadMutation, SetNotificationsReadMutationVariables>(SetNotificationsReadDocument, variables, options) as Promise<SetNotificationsReadMutation>;
    },
    portfolioSnapshot(variables?: PortfolioSnapshotQueryVariables, options?: C): Promise<PortfolioSnapshotQuery> {
      return requester<PortfolioSnapshotQuery, PortfolioSnapshotQueryVariables>(PortfolioSnapshotDocument, variables, options) as Promise<PortfolioSnapshotQuery>;
    },
    resourceSnapshot(variables?: ResourceSnapshotQueryVariables, options?: C): Promise<ResourceSnapshotQuery> {
      return requester<ResourceSnapshotQuery, ResourceSnapshotQueryVariables>(ResourceSnapshotDocument, variables, options) as Promise<ResourceSnapshotQuery>;
    },
    findProjects(variables: FindProjectsQueryVariables, options?: C): Promise<FindProjectsQuery> {
      return requester<FindProjectsQuery, FindProjectsQueryVariables>(FindProjectsDocument, variables, options) as Promise<FindProjectsQuery>;
    },
    getProject(variables: GetProjectQueryVariables, options?: C): Promise<GetProjectQuery> {
      return requester<GetProjectQuery, GetProjectQueryVariables>(GetProjectDocument, variables, options) as Promise<GetProjectQuery>;
    },
    findAllProjectTemplates(variables?: FindAllProjectTemplatesQueryVariables, options?: C): Promise<FindAllProjectTemplatesQuery> {
      return requester<FindAllProjectTemplatesQuery, FindAllProjectTemplatesQueryVariables>(FindAllProjectTemplatesDocument, variables, options) as Promise<FindAllProjectTemplatesQuery>;
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
    deleteProjectTask(variables: DeleteProjectTaskMutationVariables, options?: C): Promise<DeleteProjectTaskMutation> {
      return requester<DeleteProjectTaskMutation, DeleteProjectTaskMutationVariables>(DeleteProjectTaskDocument, variables, options) as Promise<DeleteProjectTaskMutation>;
    },
    updateProjectFeature(variables: UpdateProjectFeatureMutationVariables, options?: C): Promise<UpdateProjectFeatureMutation> {
      return requester<UpdateProjectFeatureMutation, UpdateProjectFeatureMutationVariables>(UpdateProjectFeatureDocument, variables, options) as Promise<UpdateProjectFeatureMutation>;
    },
    deleteProjectFeature(variables: DeleteProjectFeatureMutationVariables, options?: C): Promise<DeleteProjectFeatureMutation> {
      return requester<DeleteProjectFeatureMutation, DeleteProjectFeatureMutationVariables>(DeleteProjectFeatureDocument, variables, options) as Promise<DeleteProjectFeatureMutation>;
    },
    createUser(variables: CreateUserMutationVariables, options?: C): Promise<CreateUserMutation> {
      return requester<CreateUserMutation, CreateUserMutationVariables>(CreateUserDocument, variables, options) as Promise<CreateUserMutation>;
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
    findAllUsers(variables?: FindAllUsersQueryVariables, options?: C): Promise<FindAllUsersQuery> {
      return requester<FindAllUsersQuery, FindAllUsersQueryVariables>(FindAllUsersDocument, variables, options) as Promise<FindAllUsersQuery>;
    },
    currentUser(variables?: CurrentUserQueryVariables, options?: C): Promise<CurrentUserQuery> {
      return requester<CurrentUserQuery, CurrentUserQueryVariables>(CurrentUserDocument, variables, options) as Promise<CurrentUserQuery>;
    }
  };
}
export type Sdk = ReturnType<typeof getSdk>;