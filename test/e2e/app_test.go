package e2e

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"testing"
)

type UserOutput struct {
	Id    int         `json:"id,string"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Role  *RoleOutput `json:"role"`
}

type RoleOutput struct {
	Id   int    `json:"id,string"`
	Name string `json:"name"`
}

const testAddress string = "http://localhost:8080"

func Test_CreateRole(t *testing.T) {
	testData := []struct {
		Id   int
		Name string
	}{
		{1, "admin"},
		{2, "editor"},
		{3, "deletable"},
	}

	for _, test := range testData {
		t.Run("create role", func(t *testing.T) {
			body := map[string]string{
				"name": test.Name,
			}
			jsonBody, _ := json.Marshal(body)

			url := testAddress + "/role"
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Errorf("failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "ication/json")

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("get request failed: %v", err)
			}

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("failed to read response body: %v", err)
			}

			roleOutput := &RoleOutput{}
			if err := json.Unmarshal(resBody, roleOutput); err != nil {
				t.Errorf("failed to unmarshal response json: %v", err)
			}

			if roleOutput.Name != test.Name {
				t.Errorf(
					"returned role name of %v does not match posted role name of %v",
					roleOutput.Name,
					test.Name,
				)
			}

			if roleOutput.Id != test.Id {
				t.Errorf(
					"returned role id of %v does not match posted role id of %v",
					roleOutput.Id,
					test.Id,
				)
			}
		})
	}
}

func Test_CreateUser(t *testing.T) {
	testData := []struct {
		Id       int
		Name     string
		Email    string
		Password string
		RoleId   int
	}{
		{1, "user 1", "user1@test.com", "pass1", 1},
		{2, "user 2", "user2@test.com", "pass2", 2},
		{3, "user 3", "user3@test.com", "pass3", 3},
	}

	for _, test := range testData {
		t.Run("create user", func(t *testing.T) {
			body := map[string]string{
				"name":     test.Name,
				"email":    test.Email,
				"password": test.Password,
				"role_id":  strconv.Itoa(test.RoleId),
			}
			jsonBody, _ := json.Marshal(body)

			url := testAddress + "/user"
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Errorf("failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "ication/json")

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("get request failed: %v", err)
			}

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("failed to read response body: %v", err)
			}

			userOutput := &UserOutput{}
			if err := json.Unmarshal(resBody, userOutput); err != nil {
				t.Errorf("failed to unmarshal response json: %v", err)
			}

			if userOutput.Id != test.Id {
				t.Errorf(
					"returned user id of %v does not match posted user id of %v",
					userOutput.Id,
					test.Id,
				)
			}

			if userOutput.Name != test.Name {
				t.Errorf(
					"returned user name of %v does not match posted user name of %v",
					userOutput.Name,
					test.Name,
				)
			}

			if userOutput.Email != test.Email {
				t.Errorf(
					"returned user email of %v does not match posted user email of %v",
					userOutput.Email,
					test.Email,
				)
			}
		})
	}
}

func Test_UpdateUser(t *testing.T) {
	testData := []struct {
		Id     int
		Name   string
		Email  string
		RoleId int
	}{
		{1, "user 1 update", "user1-update@test.com", 2},
		{2, "user 2 update", "user2-update@test.com", 1},
	}

	for _, test := range testData {
		t.Run("update user", func(t *testing.T) {
			body := map[string]string{
				"name":    test.Name,
				"email":   test.Email,
				"role_id": strconv.Itoa(test.RoleId),
			}
			jsonBody, _ := json.Marshal(body)

			url := testAddress + "/user/" + strconv.Itoa(test.Id)
			req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Errorf("failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "ication/json")

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("get request failed: %v", err)
			}

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("failed to read response body: %v", err)
			}

			userOutput := &UserOutput{}
			if err := json.Unmarshal(resBody, userOutput); err != nil {
				t.Errorf("failed to unmarshal response json: %v", err)
			}

			if userOutput.Id != test.Id {
				t.Errorf(
					"returned user id of %v does not match posted user id of %v",
					userOutput.Id,
					test.Id,
				)
			}

			if userOutput.Name != test.Name {
				t.Errorf(
					"returned user name of %v does not match posted user name of %v",
					userOutput.Name,
					test.Name,
				)
			}

			if userOutput.Email != test.Email {
				t.Errorf(
					"returned user email of %v does not match posted user email of %v",
					userOutput.Email,
					test.Email,
				)
			}
		})
	}
}

func Test_UpdateRole(t *testing.T) {
	testData := []struct {
		Id   int
		Name string
	}{
		{1, "admin"},
		{2, "editor"},
	}

	for _, test := range testData {
		t.Run("update role", func(t *testing.T) {
			body := map[string]string{
				"name": test.Name,
			}
			jsonBody, _ := json.Marshal(body)

			url := testAddress + "/role/" + strconv.Itoa(test.Id)
			req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Errorf("failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "ication/json")

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("get request failed: %v", err)
			}

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("failed to read response body: %v", err)
			}

			roleOutput := &RoleOutput{}
			if err := json.Unmarshal(resBody, roleOutput); err != nil {
				t.Errorf("failed to unmarshal response json: %v", err)
			}

			if roleOutput.Name != test.Name {
				t.Errorf(
					"returned role name of %v does not match posted role name of %v",
					roleOutput.Name,
					test.Name,
				)
			}

			if roleOutput.Id != test.Id {
				t.Errorf(
					"returned role id of %v does not match posted role id of %v",
					roleOutput.Id,
					test.Id,
				)
			}
		})
	}
}

func Test_GetManyRoles(t *testing.T) {
	url := testAddress + "/roles"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "ication/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("get request failed: %v", err)
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("failed to read response body: %v", err)
	}

	roleOutput := []RoleOutput{}
	if err := json.Unmarshal(resBody, &roleOutput); err != nil {
		t.Errorf("failed to unmarshal response json: %v", err)
	}

	if len(roleOutput) != 3 {
		t.Error("expected 3 roles to be returned")
	}
}

func Test_GetManyUsers(t *testing.T) {
	url := testAddress + "/users"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "ication/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("get request failed: %v", err)
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("failed to read response body: %v", err)
	}

	userOutput := []UserOutput{}
	if err := json.Unmarshal(resBody, &userOutput); err != nil {
		t.Errorf("failed to unmarshal response json: %v", err)
	}

	if len(userOutput) != 3 {
		t.Error("expected 3 roles to be returned")
	}
}

func Test_UpdatePassword(t *testing.T) {
	body := map[string]string{
		"password": "newpass",
	}
	jsonBody, _ := json.Marshal(body)

	url := testAddress + "/user/1/password"
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "ication/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("get request failed: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Error("expected a 200 status code to be returned")
	}
}

func Test_DeleteUser(t *testing.T) {
	url := testAddress + "/user/3"
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "ication/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("get request failed: %v", err)
	}

	url = testAddress + "/users"
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "ication/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("get request failed: %v", err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("failed to read response body: %v", err)
	}

	userOutput := []UserOutput{}
	if err := json.Unmarshal(resBody, &userOutput); err != nil {
		t.Errorf("failed to unmarshal response json: %v: %v", err, string(resBody))
	}

	if len(userOutput) != 2 {
		t.Error("expected 2 user to be returned")
	}
}

func Test_DeleteRole(t *testing.T) {
	url := testAddress + "/role/3"
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "ication/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("get request failed: %v", err)
	}

	url = testAddress + "/roles"
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "ication/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("get request failed: %v", err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("failed to read response body: %v", err)
	}

	roleOutput := []RoleOutput{}
	if err := json.Unmarshal(resBody, &roleOutput); err != nil {
		t.Errorf("failed to unmarshal response json: %v", err)
	}

	if len(roleOutput) != 2 {
		t.Error("expected 2 role to be returned")
	}
}
