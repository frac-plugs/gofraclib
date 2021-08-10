package gofraclib

import (
	"encoding/json"
	"net/http"
)

func (u Client) GetUserByID(uid string) (*User, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://example.com/users/"+uid, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", u.Token)
	resp, err := client.Do(req)

	var user User

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u Client) GetProviderAccountByID(uid string) (*ProviderAccount, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://example.com/admin/"+uid, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", u.Token)
	resp, err := client.Do(req)

	var account *ProviderAccount

	err = json.NewDecoder(resp.Body).Decode(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
