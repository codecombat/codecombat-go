// This file was auto-generated by Fern from our API Definition.

package classrooms

import (
	context "context"
	fmt "fmt"
	codecombatgo "github.com/codecombat/codecombat-go"
	core "github.com/codecombat/codecombat-go/core"
	http "net/http"
	url "net/url"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

// Returns the classroom details for a class code.
func (c *Client) Get(ctx context.Context, request *codecombatgo.ClassroomsGetRequest) (*codecombatgo.ClassroomResponseWithCode, error) {
	baseURL := "https://codecombat.com/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "classrooms"

	queryParams := make(url.Values)
	queryParams.Add("code", fmt.Sprintf("%v", request.Code))
	if request.RetMemberLimit != nil {
		queryParams.Add("retMemberLimit", fmt.Sprintf("%v", *request.RetMemberLimit))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *codecombatgo.ClassroomResponseWithCode
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Creates a new empty `Classroom`.
func (c *Client) Create(ctx context.Context, request *codecombatgo.ClassroomsCreateRequest) (*codecombatgo.ClassroomResponseWithCode, error) {
	baseURL := "https://codecombat.com/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "classrooms"

	var response *codecombatgo.ClassroomResponseWithCode
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Upserts a user into the classroom.
//
// The document's `_id` or `slug`.
func (c *Client) UpsertMember(ctx context.Context, handle string, request *codecombatgo.ClassroomsUpsertMemberRequest) (*codecombatgo.ClassroomResponse, error) {
	baseURL := "https://codecombat.com/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"classrooms/%v/members", handle)

	var response *codecombatgo.ClassroomResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPut,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Remove a user from the classroom.
//
// The document's `_id` or `slug`.
func (c *Client) RemoveMember(ctx context.Context, handle string, request *codecombatgo.ClassroomsRemoveMemberRequest) (*codecombatgo.ClassroomResponse, error) {
	baseURL := "https://codecombat.com/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"classrooms/%v/members", handle)

	var response *codecombatgo.ClassroomResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodDelete,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Enrolls a user in a course in a classroom.
// If the course is paid, user must have an active license.
// User must be a member of the classroom.
//
// The classroom's `_id`.
// The course's `_id`.
func (c *Client) EnrollUserInCourse(ctx context.Context, classroomHandle string, courseHandle string, request *codecombatgo.ClassroomsEnrollUserInCourseRequest) (*codecombatgo.ClassroomResponse, error) {
	baseURL := "https://codecombat.com/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"classrooms/%v/courses/%v/enrolled", classroomHandle, courseHandle)

	queryParams := make(url.Values)
	if request.RetMemberLimit != nil {
		queryParams.Add("retMemberLimit", fmt.Sprintf("%v", *request.RetMemberLimit))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *codecombatgo.ClassroomResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPut,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Removes an enrolled user from a course in a classroom.
//
// The classroom's `_id`.
// The course's `_id`.
func (c *Client) RemoveEnrolledUser(ctx context.Context, classroomHandle string, courseHandle string, request *codecombatgo.ClassroomsRemoveEnrolledUserRequest) (*codecombatgo.ClassroomResponse, error) {
	baseURL := "https://codecombat.com/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"classrooms/%v/courses/%v/remove-enrolled", classroomHandle, courseHandle)

	queryParams := make(url.Values)
	if request.RetMemberLimit != nil {
		queryParams.Add("retMemberLimit", fmt.Sprintf("%v", *request.RetMemberLimit))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *codecombatgo.ClassroomResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPut,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Returns a list of all members stats for the classroom.
//
// The classroom's `_id`.
func (c *Client) GetMembersStats(ctx context.Context, classroomHandle string, request *codecombatgo.ClassroomsGetMembersStatsRequest) ([]*codecombatgo.ClassroomsGetMembersStatsResponseItem, error) {
	baseURL := "https://codecombat.com/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"classrooms/%v/stats", classroomHandle)

	queryParams := make(url.Values)
	if request.Project != nil {
		queryParams.Add("project", fmt.Sprintf("%v", *request.Project))
	}
	if request.MemberLimit != nil {
		queryParams.Add("memberLimit", fmt.Sprintf("%v", *request.MemberLimit))
	}
	if request.MemberSkip != nil {
		queryParams.Add("memberSkip", fmt.Sprintf("%v", *request.MemberSkip))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response []*codecombatgo.ClassroomsGetMembersStatsResponseItem
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Returns a list of all levels played by the user for the classroom.
//
// The classroom's `_id`.
// The classroom member's `_id`.
func (c *Client) GetLevelsPlayed(ctx context.Context, classroomHandle string, memberHandle string) ([]*codecombatgo.LevelSessionResponse, error) {
	baseURL := "https://codecombat.com/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"classrooms/%v/members/%v/sessions", classroomHandle, memberHandle)

	var response []*codecombatgo.LevelSessionResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}
