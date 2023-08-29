// This file was auto-generated by Fern from our API Definition.

package api

import (
	fmt "fmt"
	core "github.com/codecombat/codecombat-go/core"
)

type UsersCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	// `"student"` or `"teacher"`. If unset, a home user will be created, unable to join classrooms.
	Role              *core.Optional[UsersCreateRequestRole]       `json:"role,omitempty"`
	PreferredLanguage *core.Optional[string]                       `json:"preferredLanguage,omitempty"`
	HeroConfig        *core.Optional[UsersCreateRequestHeroConfig] `json:"heroConfig,omitempty"`
	Birthday          *core.Optional[string]                       `json:"birthday,omitempty"`
}

type UsersGetRequest struct {
	// Set to non-empty string to include stats.playTime in response
	IncludePlayTime *string `json:"-"`
}

type UsersGetClassroomsRequest struct {
	// limit the return number of members for each classroom
	RetMemberLimit *float64 `json:"-"`
}

type UsersGrantLicenseRequest struct {
	Ends DatetimeString `json:"ends"`
}

type UsersGrantPremiumSubscriptionRequest struct {
	Ends DatetimeString `json:"ends"`
}

type UsersSetAceConfigRequest struct {
	// controls whether autocompletion snippets show up, the default value is true
	LiveCompletion *core.Optional[bool] `json:"liveCompletion,omitempty"`
	// controls whether things like automatic parenthesis and quote completion happens, the default value is false
	Behaviors *core.Optional[bool] `json:"behaviors,omitempty"`
	// only for home users, should be one of ["python", "javascript", "cpp", "lua", "coffeescript"] right now
	Language *core.Optional[string] `json:"language,omitempty"`
}

type UsersSetHeroRequest struct {
	ThangType *core.Optional[ObjectIdString] `json:"thangType,omitempty"`
}

type UsersShortenLicenseRequest struct {
	Ends DatetimeString `json:"ends"`
}

type UsersShortenSubscriptionRequest struct {
	Ends DatetimeString `json:"ends"`
}

type UsersCreateRequestHeroConfig struct {
	ThangType *ObjectIdString `json:"thangType,omitempty"`
}

// `"student"` or `"teacher"`. If unset, a home user will be created, unable to join classrooms.
type UsersCreateRequestRole string

const (
	UsersCreateRequestRoleStudent UsersCreateRequestRole = "student"
	UsersCreateRequestRoleTeacher UsersCreateRequestRole = "teacher"
)

func NewUsersCreateRequestRoleFromString(s string) (UsersCreateRequestRole, error) {
	switch s {
	case "student":
		return UsersCreateRequestRoleStudent, nil
	case "teacher":
		return UsersCreateRequestRoleTeacher, nil
	}
	var t UsersCreateRequestRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UsersCreateRequestRole) Ptr() *UsersCreateRequestRole {
	return &u
}

type DatetimeString = string

type UsersUpdateRequest struct {
	// Set to new name string
	Name string `json:"name"`
	// Set the birthday
	Birthday *core.Optional[string] `json:"birthday,omitempty"`
}