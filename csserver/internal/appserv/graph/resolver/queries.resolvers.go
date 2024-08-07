package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"csserver/internal/appserv/csmap"
	"csserver/internal/appserv/factory"
	"csserver/internal/appserv/graph"
	"csserver/internal/appserv/graph/idl"

	"fmt"
)

// CurrentUser is the resolver for the currentUser field.
func (r *queryResolver) CurrentUser(ctx context.Context) (*idl.User, error) {
	//ch := factory.GetContextHelper()
	us := factory.GetUserService()

	// email := ch.GetUserEmailFromContext(ctx)

	user, err := us.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	out := csmap.UserUserToIdl(*user)

	return &out, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*idl.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

// FindProjects is the resolver for the findProjects field.
func (r *queryResolver) FindProjects(ctx context.Context, pageAndFilter idl.PageAndFilter) (*idl.ProjectResults, error) {
	panic(fmt.Errorf("not implemented: FindProjects - findProjects"))
}

// GetProject is the resolver for the getProject field.
func (r *queryResolver) GetProject(ctx context.Context, id string) (*idl.Project, error) {
	panic(fmt.Errorf("not implemented: GetProject - getProject"))
}

// PortfolioSnapshot is the resolver for the portfolioSnapshot field.
func (r *queryResolver) PortfolioSnapshot(ctx context.Context) (*idl.PortfolioSnapshot, error) {
	panic(fmt.Errorf("not implemented: PortfolioSnapshot - portfolioSnapshot"))
}

// ResourceSnapshot is the resolver for the resourceSnapshot field.
func (r *queryResolver) ResourceSnapshot(ctx context.Context) (*idl.ResourceSnapshot, error) {
	panic(fmt.Errorf("not implemented: ResourceSnapshot - resourceSnapshot"))
}

// FindProjectComments is the resolver for the findProjectComments field.
func (r *queryResolver) FindProjectComments(ctx context.Context, projectID string) ([]*idl.Comment, error) {
	panic(fmt.Errorf("not implemented: FindProjectComments - findProjectComments"))
}

// GetCommentThread is the resolver for the getCommentThread field.
func (r *queryResolver) GetCommentThread(ctx context.Context, id string) (*idl.Comment, error) {
	panic(fmt.Errorf("not implemented: GetCommentThread - getCommentThread"))
}

// FindActivity is the resolver for the findActivity field.
func (r *queryResolver) FindActivity(ctx context.Context, pageAndFilter idl.PageAndFilter) (*idl.ActivityResults, error) {
	panic(fmt.Errorf("not implemented: FindActivity - findActivity"))
}

// FindAllProjectTemplates is the resolver for the findAllProjectTemplates field.
func (r *queryResolver) FindAllProjectTemplates(ctx context.Context) ([]*idl.ImplementationTemplate, error) {
	panic(fmt.Errorf("not implemented: FindAllProjectTemplates - findAllProjectTemplates"))
}

// GetOrganization is the resolver for the getOrganization field.
func (r *queryResolver) GetOrganization(ctx context.Context) (*idl.Organization, error) {
	panic(fmt.Errorf("not implemented: GetOrganization - getOrganization"))
}

// FindAllUsers is the resolver for the findAllUsers field.
func (r *queryResolver) FindAllUsers(ctx context.Context) (*idl.UserResults, error) {
	panic(fmt.Errorf("not implemented: FindAllUsers - findAllUsers"))
}

// FindAllResources is the resolver for the findAllResources field.
func (r *queryResolver) FindAllResources(ctx context.Context) (*idl.ResourceResults, error) {
	panic(fmt.Errorf("not implemented: FindAllResources - findAllResources"))
}

// FindUserNotifications is the resolver for the findUserNotifications field.
func (r *queryResolver) FindUserNotifications(ctx context.Context, pageAndFilter *idl.PageAndFilter) (*idl.NotificationResults, error) {
	panic(fmt.Errorf("not implemented: FindUserNotifications - findUserNotifications"))
}

// GetResource is the resolver for the getResource field.
func (r *queryResolver) GetResource(ctx context.Context, id string) (*idl.Resource, error) {
	panic(fmt.Errorf("not implemented: GetResource - getResource"))
}

// FindAllLists is the resolver for the findAllLists field.
func (r *queryResolver) FindAllLists(ctx context.Context) (*idl.ListResults, error) {
	panic(fmt.Errorf("not implemented: FindAllLists - findAllLists"))
}

// GetList is the resolver for the getList field.
func (r *queryResolver) GetList(ctx context.Context, nameOrID string) (*idl.List, error) {
	panic(fmt.Errorf("not implemented: GetList - getList"))
}

// GetArtifact is the resolver for the getArtifact field.
func (r *queryResolver) GetArtifact(ctx context.Context, id *string) (*idl.Artifact, error) {
	panic(fmt.Errorf("not implemented: GetArtifact - getArtifact"))
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
