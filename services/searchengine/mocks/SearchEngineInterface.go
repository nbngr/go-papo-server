// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make searchengine-mocks`.

package mocks

import (
	model "bitbucket.org/enesyteam/papo-server/model"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// SearchEngineInterface is an autogenerated mock type for the SearchEngineInterface type
type SearchEngineInterface struct {
	mock.Mock
}

// DataRetentionDeleteIndexes provides a mock function with given fields: cutoff
func (_m *SearchEngineInterface) DataRetentionDeleteIndexes(cutoff time.Time) *model.AppError {
	ret := _m.Called(cutoff)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(time.Time) *model.AppError); ok {
		r0 = rf(cutoff)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteChannel provides a mock function with given fields: channel
func (_m *SearchEngineInterface) DeleteChannel(channel *model.Channel) *model.AppError {
	ret := _m.Called(channel)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Channel) *model.AppError); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteChannelPosts provides a mock function with given fields: channelID
func (_m *SearchEngineInterface) DeleteChannelPosts(channelID string) *model.AppError {
	ret := _m.Called(channelID)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeletePost provides a mock function with given fields: post
func (_m *SearchEngineInterface) DeletePost(post *model.Post) *model.AppError {
	ret := _m.Called(post)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Post) *model.AppError); ok {
		r0 = rf(post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteUser provides a mock function with given fields: user
func (_m *SearchEngineInterface) DeleteUser(user *model.User) *model.AppError {
	ret := _m.Called(user)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.User) *model.AppError); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteUserPosts provides a mock function with given fields: userID
func (_m *SearchEngineInterface) DeleteUserPosts(userID string) *model.AppError {
	ret := _m.Called(userID)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *SearchEngineInterface) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetVersion provides a mock function with given fields:
func (_m *SearchEngineInterface) GetVersion() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// IndexChannel provides a mock function with given fields: channel
func (_m *SearchEngineInterface) IndexChannel(channel *model.Channel) *model.AppError {
	ret := _m.Called(channel)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Channel) *model.AppError); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// IndexPost provides a mock function with given fields: post, teamId
func (_m *SearchEngineInterface) IndexPost(post *model.Post, teamId string) *model.AppError {
	ret := _m.Called(post, teamId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Post, string) *model.AppError); ok {
		r0 = rf(post, teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// IndexUser provides a mock function with given fields: user, teamsIds, channelsIds
func (_m *SearchEngineInterface) IndexUser(user *model.User, teamsIds []string, channelsIds []string) *model.AppError {
	ret := _m.Called(user, teamsIds, channelsIds)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.User, []string, []string) *model.AppError); ok {
		r0 = rf(user, teamsIds, channelsIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// IsActive provides a mock function with given fields:
func (_m *SearchEngineInterface) IsActive() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsAutocompletionEnabled provides a mock function with given fields:
func (_m *SearchEngineInterface) IsAutocompletionEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsIndexingEnabled provides a mock function with given fields:
func (_m *SearchEngineInterface) IsIndexingEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsIndexingSync provides a mock function with given fields:
func (_m *SearchEngineInterface) IsIndexingSync() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsSearchEnabled provides a mock function with given fields:
func (_m *SearchEngineInterface) IsSearchEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PurgeIndexes provides a mock function with given fields:
func (_m *SearchEngineInterface) PurgeIndexes() *model.AppError {
	ret := _m.Called()

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func() *model.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// RefreshIndexes provides a mock function with given fields:
func (_m *SearchEngineInterface) RefreshIndexes() *model.AppError {
	ret := _m.Called()

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func() *model.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// SearchChannels provides a mock function with given fields: teamId, term
func (_m *SearchEngineInterface) SearchChannels(teamId string, term string) ([]string, *model.AppError) {
	ret := _m.Called(teamId, term)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(teamId, term)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string) *model.AppError); ok {
		r1 = rf(teamId, term)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// SearchPosts provides a mock function with given fields: channels, searchParams, page, perPage
func (_m *SearchEngineInterface) SearchPosts(channels *model.ChannelList, searchParams []*model.SearchParams, page int, perPage int) ([]string, model.PostSearchMatches, *model.AppError) {
	ret := _m.Called(channels, searchParams, page, perPage)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*model.ChannelList, []*model.SearchParams, int, int) []string); ok {
		r0 = rf(channels, searchParams, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 model.PostSearchMatches
	if rf, ok := ret.Get(1).(func(*model.ChannelList, []*model.SearchParams, int, int) model.PostSearchMatches); ok {
		r1 = rf(channels, searchParams, page, perPage)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(model.PostSearchMatches)
		}
	}

	var r2 *model.AppError
	if rf, ok := ret.Get(2).(func(*model.ChannelList, []*model.SearchParams, int, int) *model.AppError); ok {
		r2 = rf(channels, searchParams, page, perPage)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*model.AppError)
		}
	}

	return r0, r1, r2
}

// SearchUsersInChannel provides a mock function with given fields: teamId, channelId, restrictedToChannels, term, options
func (_m *SearchEngineInterface) SearchUsersInChannel(teamId string, channelId string, restrictedToChannels []string, term string, options *model.UserSearchOptions) ([]string, []string, *model.AppError) {
	ret := _m.Called(teamId, channelId, restrictedToChannels, term, options)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string, []string, string, *model.UserSearchOptions) []string); ok {
		r0 = rf(teamId, channelId, restrictedToChannels, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(string, string, []string, string, *model.UserSearchOptions) []string); ok {
		r1 = rf(teamId, channelId, restrictedToChannels, term, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 *model.AppError
	if rf, ok := ret.Get(2).(func(string, string, []string, string, *model.UserSearchOptions) *model.AppError); ok {
		r2 = rf(teamId, channelId, restrictedToChannels, term, options)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*model.AppError)
		}
	}

	return r0, r1, r2
}

// SearchUsersInTeam provides a mock function with given fields: teamId, restrictedToChannels, term, options
func (_m *SearchEngineInterface) SearchUsersInTeam(teamId string, restrictedToChannels []string, term string, options *model.UserSearchOptions) ([]string, *model.AppError) {
	ret := _m.Called(teamId, restrictedToChannels, term, options)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, []string, string, *model.UserSearchOptions) []string); ok {
		r0 = rf(teamId, restrictedToChannels, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, []string, string, *model.UserSearchOptions) *model.AppError); ok {
		r1 = rf(teamId, restrictedToChannels, term, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *SearchEngineInterface) Start() *model.AppError {
	ret := _m.Called()

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func() *model.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *SearchEngineInterface) Stop() *model.AppError {
	ret := _m.Called()

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func() *model.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// TestConfig provides a mock function with given fields: cfg
func (_m *SearchEngineInterface) TestConfig(cfg *model.Config) *model.AppError {
	ret := _m.Called(cfg)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Config) *model.AppError); ok {
		r0 = rf(cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// UpdateConfig provides a mock function with given fields: cfg
func (_m *SearchEngineInterface) UpdateConfig(cfg *model.Config) {
	_m.Called(cfg)
}
