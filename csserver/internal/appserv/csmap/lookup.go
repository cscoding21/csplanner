package csmap

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/appserv/orgmap"
	"csserver/internal/common"
	"csserver/internal/services/iam/appuser"
	"csserver/internal/services/list"
	"csserver/internal/services/project"
	"csserver/internal/services/projecttemplate"
	"csserver/internal/services/resource"
	"csserver/internal/utils"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

// ---TODO: impelement better caching
var (
	_skillCache    *[]list.ListItem
	_resourceCache *[]resource.Resource
	_userCache     *[]appuser.Appuser
	_projectCache  *[]project.Project
	_roleCache     *[]resource.Role
	_templateCache *[]projecttemplate.Projecttemplate
	_listCache     *[]list.List

	_skillCacheLock    sync.Mutex
	_resourceCacheLock sync.Mutex
	_userCacheLock     sync.Mutex
	_projectCacheLock  sync.Mutex
	_roleCacheLock     sync.Mutex
	_templateCacheLock sync.Mutex
	_listCacheLock     sync.Mutex
)

func ClearAllCaches() {
	ExpireProjectCache()
	ExpireResourceCache()
	ExpireTemplateCache()
	ExpireRoleCache()
	ExpireListCache()
	ExpireUserCache()

}

func ExpireProjectCache() {
	_projectCacheLock.Lock()
	defer _projectCacheLock.Unlock()
	_projectCache = nil
}

func ExpireResourceCache() {
	_resourceCacheLock.Lock()
	defer _resourceCacheLock.Unlock()
	_resourceCache = nil
}

func ExpireUserCache() {
	_userCacheLock.Lock()
	defer _userCacheLock.Unlock()
	_userCache = nil
}

func ExpireRoleCache() {
	_roleCacheLock.Lock()
	defer _roleCacheLock.Unlock()
	_roleCache = nil
}

func ExpireTemplateCache() {
	_templateCacheLock.Lock()
	defer _templateCacheLock.Unlock()
	_templateCache = nil
}

func ExpireListCache() {
	_listCacheLock.Lock()
	defer _listCacheLock.Unlock()
	_listCache = nil
}

func getSkills(ctx context.Context) *[]list.ListItem {
	if _skillCache != nil && len(*_skillCache) > 0 {
		return _skillCache
	}

	_skillCacheLock.Lock()
	defer _skillCacheLock.Unlock()

	ss := factory.GetListService(ctx)

	skillList, err := ss.GetList(ctx, list.ListNameSkills)
	if err != nil {
		log.Error(err)
		return nil
	}

	sc := utils.RefToValSlice(skillList.Data.Values)
	_skillCache = &sc

	return _skillCache
}

func getSkillById(items []list.ListItem, id string) *list.ListItem {
	for _, item := range items {
		if strings.EqualFold(item.Value, id) {
			return &item
		}
	}

	return nil
}

func getListById(lists []list.List, name string) *list.List {
	for _, l := range lists {
		if strings.EqualFold(l.Name, name) {
			return &l
		}
	}

	return nil
}

func getResourceById(ctx context.Context, resources []resource.Resource, id string) *idl.Resource {
	for _, r := range resources {
		if r.ID == id {
			out := ResourceResourceToIdl(r)
			AugmentResource(ctx, &out)

			return &out
		}
	}

	return nil
}

func getRoleById(ctx context.Context, id string) *idl.Role {
	roles := findRoles(ctx)
	for _, p := range *roles {
		if p.ID == id {
			out := RoleResourceToIdl(p)

			return &out
		}
	}

	return nil
}

func getProjectById(ctx context.Context, projects []project.Project, id string) *idl.Project {
	for _, p := range projects {
		if p.ID == id {
			out := ProjectProjectToIdl(p)

			AugmentProject(ctx, &p, &out)

			return &out
		}
	}

	return nil
}

func findProjects(ctx context.Context) *[]project.Project {
	if _projectCache != nil && len(*_projectCache) > 0 {
		return _projectCache
	}

	_projectCacheLock.Lock()
	defer _projectCacheLock.Unlock()

	rs := factory.GetProjectService(ctx)

	projects, err := rs.FindAllProjects(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	_projectCache = utils.ValToRef(common.ExtractDataFromBase(projects.Results))

	return _projectCache
}

func findResources(ctx context.Context) *[]resource.Resource {
	if _resourceCache != nil && len(*_resourceCache) > 0 {
		return _resourceCache
	}

	_resourceCacheLock.Lock()
	defer _resourceCacheLock.Unlock()

	rs := factory.GetResourceService(ctx)

	resources, err := rs.FindAllResources(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	_resourceCache = utils.ValToRef(common.ExtractDataFromBase(resources.Results))

	return _resourceCache
}

func findUsers(ctx context.Context) *[]appuser.Appuser {
	if _userCache != nil && len(*_userCache) > 0 {
		return _userCache
	}

	_userCacheLock.Lock()
	defer _userCacheLock.Unlock()

	orgInfo, err := orgmap.GetSaaSOrg(ctx)
	if err != nil {
		log.Error(err)
	}

	us := factory.GetIAMAdminService(ctx)

	users, err := us.FindAllUsers(ctx, orgInfo.Info.Org.Realm)
	if err != nil {
		log.Error(err)
		return nil
	}

	_userCache = &users.Results

	return _userCache
}

func findLists(ctx context.Context) *[]list.List {
	if _listCache != nil && len(*_listCache) > 0 {
		return _listCache
	}

	_listCacheLock.Lock()
	defer _listCacheLock.Unlock()

	us := factory.GetListService(ctx)

	lists, err := us.FindAllLists(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	_listCache = utils.ValToRef(common.ExtractDataFromBase(lists.Results))

	return _listCache
}

func findRoles(ctx context.Context) *[]resource.Role {
	if _roleCache != nil && len(*_roleCache) > 0 {
		return _roleCache
	}

	_roleCacheLock.Lock()
	defer _roleCacheLock.Unlock()

	us := factory.GetResourceService(ctx)

	roles, err := us.FindAllRoles(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	_roleCache = utils.ValToRef(common.ExtractDataFromBase(roles.Results))

	return _roleCache
}

func findTemplates(ctx context.Context) *[]projecttemplate.Projecttemplate {
	if _templateCache != nil && len(*_templateCache) > 0 {
		return _templateCache
	}

	_templateCacheLock.Lock()
	defer _templateCacheLock.Unlock()

	us := factory.GetProjectTemplateService(ctx)

	templates, err := us.FindAllProjecttemplates(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	_templateCache = utils.ValToRef(common.ExtractDataFromBase(templates.Results))

	return _templateCache
}

func getUserByEmail(ctx context.Context, email string) *idl.User {
	users := findUsers(ctx)

	for _, u := range *users {
		if u.Email == email {
			out := AppuserAppuserToIdl(u)
			return &out
		}
	}

	return nil
}

func getStatusTransitionDetails(ctx context.Context, p *project.Project) {
	service := factory.GetProjectService(ctx)

	service.GetStatusTransitionDetails(p)
}
