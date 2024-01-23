// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	context "context"

	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// TeamStore is an autogenerated mock type for the TeamStore type
type TeamStore struct {
	mock.Mock
}

// AnalyticsGetTeamCountForScheme provides a mock function with given fields: schemeID
func (_m *TeamStore) AnalyticsGetTeamCountForScheme(schemeID string) (int64, error) {
	ret := _m.Called(schemeID)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(schemeID)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(schemeID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(schemeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnalyticsTeamCount provides a mock function with given fields: opts
func (_m *TeamStore) AnalyticsTeamCount(opts *model.TeamSearch) (int64, error) {
	ret := _m.Called(opts)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.TeamSearch) (int64, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*model.TeamSearch) int64); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*model.TeamSearch) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClearAllCustomRoleAssignments provides a mock function with given fields:
func (_m *TeamStore) ClearAllCustomRoleAssignments() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClearCaches provides a mock function with given fields:
func (_m *TeamStore) ClearCaches() {
	_m.Called()
}

// Get provides a mock function with given fields: id
func (_m *TeamStore) Get(id string) (*model.Team, error) {
	ret := _m.Called(id)

	var r0 *model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Team, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Team); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveMemberCount provides a mock function with given fields: teamID, restrictions
func (_m *TeamStore) GetActiveMemberCount(teamID string, restrictions *model.ViewUsersRestrictions) (int64, error) {
	ret := _m.Called(teamID, restrictions)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *model.ViewUsersRestrictions) (int64, error)); ok {
		return rf(teamID, restrictions)
	}
	if rf, ok := ret.Get(0).(func(string, *model.ViewUsersRestrictions) int64); ok {
		r0 = rf(teamID, restrictions)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, *model.ViewUsersRestrictions) error); ok {
		r1 = rf(teamID, restrictions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *TeamStore) GetAll() ([]*model.Team, error) {
	ret := _m.Called()

	var r0 []*model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.Team, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.Team); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllForExportAfter provides a mock function with given fields: limit, afterID
func (_m *TeamStore) GetAllForExportAfter(limit int, afterID string) ([]*model.TeamForExport, error) {
	ret := _m.Called(limit, afterID)

	var r0 []*model.TeamForExport
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string) ([]*model.TeamForExport, error)); ok {
		return rf(limit, afterID)
	}
	if rf, ok := ret.Get(0).(func(int, string) []*model.TeamForExport); ok {
		r0 = rf(limit, afterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TeamForExport)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(limit, afterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllPage provides a mock function with given fields: offset, limit, opts
func (_m *TeamStore) GetAllPage(offset int, limit int, opts *model.TeamSearch) ([]*model.Team, error) {
	ret := _m.Called(offset, limit, opts)

	var r0 []*model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, *model.TeamSearch) ([]*model.Team, error)); ok {
		return rf(offset, limit, opts)
	}
	if rf, ok := ret.Get(0).(func(int, int, *model.TeamSearch) []*model.Team); ok {
		r0 = rf(offset, limit, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, *model.TeamSearch) error); ok {
		r1 = rf(offset, limit, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllPrivateTeamListing provides a mock function with given fields:
func (_m *TeamStore) GetAllPrivateTeamListing() ([]*model.Team, error) {
	ret := _m.Called()

	var r0 []*model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.Team, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.Team); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTeamListing provides a mock function with given fields:
func (_m *TeamStore) GetAllTeamListing() ([]*model.Team, error) {
	ret := _m.Called()

	var r0 []*model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.Team, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.Team); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmptyInviteID provides a mock function with given fields:
func (_m *TeamStore) GetByEmptyInviteID() ([]*model.Team, error) {
	ret := _m.Called()

	var r0 []*model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.Team, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.Team); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByInviteId provides a mock function with given fields: inviteID
func (_m *TeamStore) GetByInviteId(inviteID string) (*model.Team, error) {
	ret := _m.Called(inviteID)

	var r0 *model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Team, error)); ok {
		return rf(inviteID)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Team); ok {
		r0 = rf(inviteID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(inviteID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: name
func (_m *TeamStore) GetByName(name string) (*model.Team, error) {
	ret := _m.Called(name)

	var r0 *model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Team, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Team); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByNames provides a mock function with given fields: name
func (_m *TeamStore) GetByNames(name []string) ([]*model.Team, error) {
	ret := _m.Called(name)

	var r0 []*model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]*model.Team, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func([]string) []*model.Team); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChannelUnreadsForAllTeams provides a mock function with given fields: excludeTeamID, userID
func (_m *TeamStore) GetChannelUnreadsForAllTeams(excludeTeamID string, userID string) ([]*model.ChannelUnread, error) {
	ret := _m.Called(excludeTeamID, userID)

	var r0 []*model.ChannelUnread
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]*model.ChannelUnread, error)); ok {
		return rf(excludeTeamID, userID)
	}
	if rf, ok := ret.Get(0).(func(string, string) []*model.ChannelUnread); ok {
		r0 = rf(excludeTeamID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ChannelUnread)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(excludeTeamID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChannelUnreadsForTeam provides a mock function with given fields: teamID, userID
func (_m *TeamStore) GetChannelUnreadsForTeam(teamID string, userID string) ([]*model.ChannelUnread, error) {
	ret := _m.Called(teamID, userID)

	var r0 []*model.ChannelUnread
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]*model.ChannelUnread, error)); ok {
		return rf(teamID, userID)
	}
	if rf, ok := ret.Get(0).(func(string, string) []*model.ChannelUnread); ok {
		r0 = rf(teamID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ChannelUnread)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(teamID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommonTeamIDsForMultipleUsers provides a mock function with given fields: userIDs
func (_m *TeamStore) GetCommonTeamIDsForMultipleUsers(userIDs []string) ([]string, error) {
	ret := _m.Called(userIDs)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]string, error)); ok {
		return rf(userIDs)
	}
	if rf, ok := ret.Get(0).(func([]string) []string); ok {
		r0 = rf(userIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(userIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommonTeamIDsForTwoUsers provides a mock function with given fields: userID, otherUserID
func (_m *TeamStore) GetCommonTeamIDsForTwoUsers(userID string, otherUserID string) ([]string, error) {
	ret := _m.Called(userID, otherUserID)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]string, error)); ok {
		return rf(userID, otherUserID)
	}
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(userID, otherUserID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, otherUserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMany provides a mock function with given fields: ids
func (_m *TeamStore) GetMany(ids []string) ([]*model.Team, error) {
	ret := _m.Called(ids)

	var r0 []*model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]*model.Team, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]string) []*model.Team); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMember provides a mock function with given fields: ctx, teamID, userID
func (_m *TeamStore) GetMember(ctx context.Context, teamID string, userID string) (*model.TeamMember, error) {
	ret := _m.Called(ctx, teamID, userID)

	var r0 *model.TeamMember
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*model.TeamMember, error)); ok {
		return rf(ctx, teamID, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.TeamMember); ok {
		r0 = rf(ctx, teamID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TeamMember)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, teamID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMembers provides a mock function with given fields: teamID, offset, limit, teamMembersGetOptions
func (_m *TeamStore) GetMembers(teamID string, offset int, limit int, teamMembersGetOptions *model.TeamMembersGetOptions) ([]*model.TeamMember, error) {
	ret := _m.Called(teamID, offset, limit, teamMembersGetOptions)

	var r0 []*model.TeamMember
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int, *model.TeamMembersGetOptions) ([]*model.TeamMember, error)); ok {
		return rf(teamID, offset, limit, teamMembersGetOptions)
	}
	if rf, ok := ret.Get(0).(func(string, int, int, *model.TeamMembersGetOptions) []*model.TeamMember); ok {
		r0 = rf(teamID, offset, limit, teamMembersGetOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TeamMember)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int, *model.TeamMembersGetOptions) error); ok {
		r1 = rf(teamID, offset, limit, teamMembersGetOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMembersByIds provides a mock function with given fields: teamID, userIds, restrictions
func (_m *TeamStore) GetMembersByIds(teamID string, userIds []string, restrictions *model.ViewUsersRestrictions) ([]*model.TeamMember, error) {
	ret := _m.Called(teamID, userIds, restrictions)

	var r0 []*model.TeamMember
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string, *model.ViewUsersRestrictions) ([]*model.TeamMember, error)); ok {
		return rf(teamID, userIds, restrictions)
	}
	if rf, ok := ret.Get(0).(func(string, []string, *model.ViewUsersRestrictions) []*model.TeamMember); ok {
		r0 = rf(teamID, userIds, restrictions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TeamMember)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string, *model.ViewUsersRestrictions) error); ok {
		r1 = rf(teamID, userIds, restrictions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamMembersForExport provides a mock function with given fields: userID
func (_m *TeamStore) GetTeamMembersForExport(userID string) ([]*model.TeamMemberForExport, error) {
	ret := _m.Called(userID)

	var r0 []*model.TeamMemberForExport
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*model.TeamMemberForExport, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) []*model.TeamMemberForExport); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TeamMemberForExport)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamsByScheme provides a mock function with given fields: schemeID, offset, limit
func (_m *TeamStore) GetTeamsByScheme(schemeID string, offset int, limit int) ([]*model.Team, error) {
	ret := _m.Called(schemeID, offset, limit)

	var r0 []*model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]*model.Team, error)); ok {
		return rf(schemeID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.Team); ok {
		r0 = rf(schemeID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(schemeID, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamsByUserId provides a mock function with given fields: userID
func (_m *TeamStore) GetTeamsByUserId(userID string) ([]*model.Team, error) {
	ret := _m.Called(userID)

	var r0 []*model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*model.Team, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) []*model.Team); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamsForUser provides a mock function with given fields: ctx, userID, excludeTeamID, includeDeleted
func (_m *TeamStore) GetTeamsForUser(ctx context.Context, userID string, excludeTeamID string, includeDeleted bool) ([]*model.TeamMember, error) {
	ret := _m.Called(ctx, userID, excludeTeamID, includeDeleted)

	var r0 []*model.TeamMember
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) ([]*model.TeamMember, error)); ok {
		return rf(ctx, userID, excludeTeamID, includeDeleted)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) []*model.TeamMember); ok {
		r0 = rf(ctx, userID, excludeTeamID, includeDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TeamMember)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, bool) error); ok {
		r1 = rf(ctx, userID, excludeTeamID, includeDeleted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamsForUserWithPagination provides a mock function with given fields: userID, page, perPage
func (_m *TeamStore) GetTeamsForUserWithPagination(userID string, page int, perPage int) ([]*model.TeamMember, error) {
	ret := _m.Called(userID, page, perPage)

	var r0 []*model.TeamMember
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]*model.TeamMember, error)); ok {
		return rf(userID, page, perPage)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.TeamMember); ok {
		r0 = rf(userID, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TeamMember)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(userID, page, perPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalMemberCount provides a mock function with given fields: teamID, restrictions
func (_m *TeamStore) GetTotalMemberCount(teamID string, restrictions *model.ViewUsersRestrictions) (int64, error) {
	ret := _m.Called(teamID, restrictions)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *model.ViewUsersRestrictions) (int64, error)); ok {
		return rf(teamID, restrictions)
	}
	if rf, ok := ret.Get(0).(func(string, *model.ViewUsersRestrictions) int64); ok {
		r0 = rf(teamID, restrictions)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, *model.ViewUsersRestrictions) error); ok {
		r1 = rf(teamID, restrictions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserTeamIds provides a mock function with given fields: userID, allowFromCache
func (_m *TeamStore) GetUserTeamIds(userID string, allowFromCache bool) ([]string, error) {
	ret := _m.Called(userID, allowFromCache)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) ([]string, error)); ok {
		return rf(userID, allowFromCache)
	}
	if rf, ok := ret.Get(0).(func(string, bool) []string); ok {
		r0 = rf(userID, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(userID, allowFromCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupSyncedTeamCount provides a mock function with given fields:
func (_m *TeamStore) GroupSyncedTeamCount() (int64, error) {
	ret := _m.Called()

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvalidateAllTeamIdsForUser provides a mock function with given fields: userID
func (_m *TeamStore) InvalidateAllTeamIdsForUser(userID string) {
	_m.Called(userID)
}

// MigrateTeamMembers provides a mock function with given fields: fromTeamID, fromUserID
func (_m *TeamStore) MigrateTeamMembers(fromTeamID string, fromUserID string) (map[string]string, error) {
	ret := _m.Called(fromTeamID, fromUserID)

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (map[string]string, error)); ok {
		return rf(fromTeamID, fromUserID)
	}
	if rf, ok := ret.Get(0).(func(string, string) map[string]string); ok {
		r0 = rf(fromTeamID, fromUserID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(fromTeamID, fromUserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDelete provides a mock function with given fields: teamID
func (_m *TeamStore) PermanentDelete(teamID string) error {
	ret := _m.Called(teamID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(teamID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveAllMembersByTeam provides a mock function with given fields: teamID
func (_m *TeamStore) RemoveAllMembersByTeam(teamID string) error {
	ret := _m.Called(teamID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(teamID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveAllMembersByUser provides a mock function with given fields: userID
func (_m *TeamStore) RemoveAllMembersByUser(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveMember provides a mock function with given fields: teamID, userID
func (_m *TeamStore) RemoveMember(teamID string, userID string) error {
	ret := _m.Called(teamID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(teamID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveMembers provides a mock function with given fields: teamID, userIds
func (_m *TeamStore) RemoveMembers(teamID string, userIds []string) error {
	ret := _m.Called(teamID, userIds)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(teamID, userIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetAllTeamSchemes provides a mock function with given fields:
func (_m *TeamStore) ResetAllTeamSchemes() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: team
func (_m *TeamStore) Save(team *model.Team) (*model.Team, error) {
	ret := _m.Called(team)

	var r0 *model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Team) (*model.Team, error)); ok {
		return rf(team)
	}
	if rf, ok := ret.Get(0).(func(*model.Team) *model.Team); ok {
		r0 = rf(team)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Team) error); ok {
		r1 = rf(team)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveMember provides a mock function with given fields: member, maxUsersPerTeam
func (_m *TeamStore) SaveMember(member *model.TeamMember, maxUsersPerTeam int) (*model.TeamMember, error) {
	ret := _m.Called(member, maxUsersPerTeam)

	var r0 *model.TeamMember
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.TeamMember, int) (*model.TeamMember, error)); ok {
		return rf(member, maxUsersPerTeam)
	}
	if rf, ok := ret.Get(0).(func(*model.TeamMember, int) *model.TeamMember); ok {
		r0 = rf(member, maxUsersPerTeam)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TeamMember)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.TeamMember, int) error); ok {
		r1 = rf(member, maxUsersPerTeam)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveMultipleMembers provides a mock function with given fields: members, maxUsersPerTeam
func (_m *TeamStore) SaveMultipleMembers(members []*model.TeamMember, maxUsersPerTeam int) ([]*model.TeamMember, error) {
	ret := _m.Called(members, maxUsersPerTeam)

	var r0 []*model.TeamMember
	var r1 error
	if rf, ok := ret.Get(0).(func([]*model.TeamMember, int) ([]*model.TeamMember, error)); ok {
		return rf(members, maxUsersPerTeam)
	}
	if rf, ok := ret.Get(0).(func([]*model.TeamMember, int) []*model.TeamMember); ok {
		r0 = rf(members, maxUsersPerTeam)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TeamMember)
		}
	}

	if rf, ok := ret.Get(1).(func([]*model.TeamMember, int) error); ok {
		r1 = rf(members, maxUsersPerTeam)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchAll provides a mock function with given fields: opts
func (_m *TeamStore) SearchAll(opts *model.TeamSearch) ([]*model.Team, error) {
	ret := _m.Called(opts)

	var r0 []*model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.TeamSearch) ([]*model.Team, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*model.TeamSearch) []*model.Team); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.TeamSearch) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchAllPaged provides a mock function with given fields: opts
func (_m *TeamStore) SearchAllPaged(opts *model.TeamSearch) ([]*model.Team, int64, error) {
	ret := _m.Called(opts)

	var r0 []*model.Team
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(*model.TeamSearch) ([]*model.Team, int64, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*model.TeamSearch) []*model.Team); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.TeamSearch) int64); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(*model.TeamSearch) error); ok {
		r2 = rf(opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SearchOpen provides a mock function with given fields: opts
func (_m *TeamStore) SearchOpen(opts *model.TeamSearch) ([]*model.Team, error) {
	ret := _m.Called(opts)

	var r0 []*model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.TeamSearch) ([]*model.Team, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*model.TeamSearch) []*model.Team); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.TeamSearch) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchPrivate provides a mock function with given fields: opts
func (_m *TeamStore) SearchPrivate(opts *model.TeamSearch) ([]*model.Team, error) {
	ret := _m.Called(opts)

	var r0 []*model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.TeamSearch) ([]*model.Team, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*model.TeamSearch) []*model.Team); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.TeamSearch) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: team
func (_m *TeamStore) Update(team *model.Team) (*model.Team, error) {
	ret := _m.Called(team)

	var r0 *model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Team) (*model.Team, error)); ok {
		return rf(team)
	}
	if rf, ok := ret.Get(0).(func(*model.Team) *model.Team); ok {
		r0 = rf(team)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Team) error); ok {
		r1 = rf(team)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLastTeamIconUpdate provides a mock function with given fields: teamID, curTime
func (_m *TeamStore) UpdateLastTeamIconUpdate(teamID string, curTime int64) error {
	ret := _m.Called(teamID, curTime)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(teamID, curTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMember provides a mock function with given fields: member
func (_m *TeamStore) UpdateMember(member *model.TeamMember) (*model.TeamMember, error) {
	ret := _m.Called(member)

	var r0 *model.TeamMember
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.TeamMember) (*model.TeamMember, error)); ok {
		return rf(member)
	}
	if rf, ok := ret.Get(0).(func(*model.TeamMember) *model.TeamMember); ok {
		r0 = rf(member)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TeamMember)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.TeamMember) error); ok {
		r1 = rf(member)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMembersRole provides a mock function with given fields: teamID, userIDs
func (_m *TeamStore) UpdateMembersRole(teamID string, userIDs []string) error {
	ret := _m.Called(teamID, userIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(teamID, userIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMultipleMembers provides a mock function with given fields: members
func (_m *TeamStore) UpdateMultipleMembers(members []*model.TeamMember) ([]*model.TeamMember, error) {
	ret := _m.Called(members)

	var r0 []*model.TeamMember
	var r1 error
	if rf, ok := ret.Get(0).(func([]*model.TeamMember) ([]*model.TeamMember, error)); ok {
		return rf(members)
	}
	if rf, ok := ret.Get(0).(func([]*model.TeamMember) []*model.TeamMember); ok {
		r0 = rf(members)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TeamMember)
		}
	}

	if rf, ok := ret.Get(1).(func([]*model.TeamMember) error); ok {
		r1 = rf(members)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserBelongsToTeams provides a mock function with given fields: userID, teamIds
func (_m *TeamStore) UserBelongsToTeams(userID string, teamIds []string) (bool, error) {
	ret := _m.Called(userID, teamIds)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string) (bool, error)); ok {
		return rf(userID, teamIds)
	}
	if rf, ok := ret.Get(0).(func(string, []string) bool); ok {
		r0 = rf(userID, teamIds)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(userID, teamIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTeamStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewTeamStore creates a new instance of TeamStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTeamStore(t mockConstructorTestingTNewTeamStore) *TeamStore {
	mock := &TeamStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
