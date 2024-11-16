package middleware

import "encoding/json"

type AuthData struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
}

func (authData *AuthData) LoadFromMap(m map[string]interface{}) error {
	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, authData)
	}
	return err
}
