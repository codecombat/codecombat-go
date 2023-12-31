// This file was auto-generated by Fern from our API Definition.

package api

import (
	core "github.com/codecombat/codecombat-go/core"
)

type PostUsersHandleOAuthIdentitiesRequest struct {
	// Your OAuth Provider ID.
	Provider string `json:"provider"`
	// Will be passed through your lookup URL to get the user ID. Required if no `code`.
	AccessToken *core.Optional[string] `json:"accessToken,omitempty"`
	// Will be passed to the OAuth token endpoint to get a token. Required if no `accessToken`.
	Code *core.Optional[string] `json:"code,omitempty"`
}

type ClassroomResponseCoursesItem struct {
	Id         *ObjectIdString  `json:"_id,omitempty"`
	Levels     []map[string]any `json:"levels,omitempty"`
	Enrolled   []ObjectIdString `json:"enrolled,omitempty"`
	InstanceId *ObjectIdString  `json:"instance_id,omitempty"`
}

// Subset of properties listed here
type ClassroomResponseWithCode struct {
	Id          *ObjectIdString                         `json:"_id,omitempty"`
	Name        *string                                 `json:"name,omitempty"`
	Members     []ObjectIdString                        `json:"members,omitempty"`
	OwnerId     *ObjectIdString                         `json:"ownerID,omitempty"`
	Description *string                                 `json:"description,omitempty"`
	Code        *string                                 `json:"code,omitempty"`
	CodeCamel   *string                                 `json:"codeCamel,omitempty"`
	Courses     []*ClassroomResponseWithCodeCoursesItem `json:"courses,omitempty"`
	ClanId      *ObjectIdString                         `json:"clanId,omitempty"`
}

type ClassroomResponseWithCodeCoursesItem struct {
	Id         *ObjectIdString  `json:"_id,omitempty"`
	Levels     []map[string]any `json:"levels,omitempty"`
	Enrolled   []ObjectIdString `json:"enrolled,omitempty"`
	InstanceId *ObjectIdString  `json:"instance_id,omitempty"`
}

type ClassroomsGetMembersStatsResponseItemStats struct {
	GamesCompleted *float64 `json:"gamesCompleted,omitempty"`
	// Total play time in seconds
	Playtime *float64 `json:"playtime,omitempty"`
}

type LevelSessionResponseLevel struct {
	// The id for the level.
	Original *string `json:"original,omitempty"`
}

type LevelSessionResponseState struct {
	Complete *bool `json:"complete,omitempty"`
}

// Subset of properties listed here
type UserResponse struct {
	Id              *ObjectIdString                    `json:"_id,omitempty"`
	Email           *string                            `json:"email,omitempty"`
	Name            *string                            `json:"name,omitempty"`
	Slug            *string                            `json:"slug,omitempty"`
	Role            *RoleString                        `json:"role,omitempty"`
	Stats           *UserResponseStats                 `json:"stats,omitempty"`
	OAuthIdentities []*UserResponseOAuthIdentitiesItem `json:"oAuthIdentities,omitempty"`
	Subscription    *UserResponseSubscription          `json:"subscription,omitempty"`
	License         *UserResponseLicense               `json:"license,omitempty"`
}

type UserResponseLicense struct {
	Ends   *DatetimeString `json:"ends,omitempty"`
	Active *bool           `json:"active,omitempty"`
}

type UserResponseOAuthIdentitiesItem struct {
	Provider *string `json:"provider,omitempty"`
	Id       *string `json:"id,omitempty"`
}

type UserResponseStats struct {
	GamesCompleted *float64           `json:"gamesCompleted,omitempty"`
	Concepts       map[string]float64 `json:"concepts,omitempty"`
	// Included only when specifically requested on the endpoint
	PlayTime *float64 `json:"playTime,omitempty"`
}

type UserResponseSubscription struct {
	Ends   *DatetimeString `json:"ends,omitempty"`
	Active *bool           `json:"active,omitempty"`
}

type ObjectIdString = string

// Usually either 'teacher' or 'student'
type RoleString = string
