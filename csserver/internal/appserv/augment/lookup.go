package augment

import (
	"context"
	"csserver/internal/appserv/csmap"
	"csserver/internal/appserv/factory"
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/services/iam/user"
	"csserver/internal/services/list"
	"csserver/internal/services/resource"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

// ---TODO: impelement better caching
var (
	_skillCache    *[]list.ListItem
	_resourceCache *[]resource.Resource
	_userCache     *[]user.User

	_skillCacheLock    sync.Mutex
	_resourceCacheLock sync.Mutex
	_userCacheLock     sync.Mutex
)

func getSkills() *[]list.ListItem {
	if _skillCache != nil {
		return _skillCache
	}

	_skillCacheLock.Lock()
	defer _skillCacheLock.Unlock()

	ctx := context.Background()
	ss := factory.GetListService()

	skillList, err := ss.GetList(ctx, "Skills")
	if err != nil {
		log.Error(err)
		return nil
	}

	_skillCache = &skillList.Values

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

func getResourceById(resources []resource.Resource, id string) *idl.Resource {
	for _, r := range resources {
		if r.ID == id {
			out := csmap.ResourceResourceToIdl(r)

			if r.UserEmail != nil {
				out.User = getUserByEmail(*r.UserEmail)
			}

			return &out
		}
	}

	return nil
}

func findResources() *[]resource.Resource {
	if _resourceCache != nil {
		return _resourceCache
	}

	_resourceCacheLock.Lock()
	defer _resourceCacheLock.Unlock()

	ctx := context.Background()
	rs := factory.GetResourceService()

	resources, err := rs.FindAllResources(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	_resourceCache = &resources.Results

	return _resourceCache
}

func findUsers() *[]user.User {
	if _userCache != nil {
		return _userCache
	}

	_userCacheLock.Lock()
	defer _userCacheLock.Unlock()

	ctx := context.Background()
	us := factory.GetIAMAdminService()

	users, err := us.FindAllUsers(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	_userCache = &users.Results

	return _userCache
}

func getUserByEmail(email string) *idl.User {
	users := findUsers()

	for _, u := range *users {
		if u.Email == email {
			out := csmap.UserUserToIdl(u)
			return &out
		}
	}

	return nil
}
