package tp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	baseURLString  = ".tpondemand.com/api/v1/"
	libraryVersion = "0.1"
	userAgent      = "tipi/" + libraryVersion
)

var (
	subdomain string
	token     string
)

// Client struct
type Client struct {

	// HTTP client
	client    *http.Client
	UserAgent string

	Bugs        *BugService
	Comments    *CommentService
	Entities    *EntityService
	Features    *FeatureService
	Tasks       *TaskService
	Users       *UserService
	UserStories *UserStoryService
}

// NewClient func
func NewClient(dmn, tkn string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	subdomain = dmn
	token = tkn

	c := &Client{client: httpClient, UserAgent: userAgent}

	c.Bugs = &BugService{client: c}
	c.Comments = &CommentService{client: c}
	c.Entities = &EntityService{client: c}
	c.Features = &FeatureService{client: c}
	c.Tasks = &TaskService{client: c}
	c.Users = &UserService{client: c}
	c.UserStories = &UserStoryService{client: c}

	return c
}

// NewRequest func
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {

	var url string
	if strings.Contains(urlStr, "?") {
		url = fmt.Sprintf("https://%s%s%s&format=json&token=%s", subdomain, baseURLString, urlStr, token)
	} else {
		url = fmt.Sprintf("https://%s%s%s?format=json&token=%s", subdomain, baseURLString, urlStr, token)
	}

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

// Do function
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return response, err
}

// Response struct
type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// ErrorResponse struct
type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"error,omitempty"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method,
		r.Response.Request.URL,
		r.Response.StatusCode,
		r.Message)
}

// CheckResponse struct
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}
