package depend

// CorpToken a method of get corp token,user is xdchenadmin@admin
func corpToken() (cid, token string, err error) {
	email := "xdchenadmin@admin"
	password := "12345678"
	j, err := Login(email, "", password)
	if err != nil {
		return "", "", err
	}
	cid, _ = j.Get("corp_id").String()
	token, _ = j.Get("token").String()
	return cid, token, nil
}

// SuperToken a method of get super user token
func superToken() (token string, err error) {
	email := "admin@admin"
	password := "abc123"
	j, err := Login(email, "", password)
	if err != nil {
		return "", err
	}
	token, _ = j.Get("token").String()
	return token, nil
}

// Auth return header of the Authorization
func Auth(token string) map[string]string {
	auth := map[string]string{}
	if token == "" {
		return nil
	}
	auth["Authorization"] = token
	return auth
}

// CorpAuth return CorpId and header of Auth
func CorpAuth() (corpID string, auth map[string]string) {
	id, token, err := corpToken()
	if err != nil {
		return "", nil
	}
	auth = Auth(token)
	return id, auth
}

// SuperAuth return header of Auth
func SuperAuth() (auth map[string]string) {
	token, err := superToken()
	if err != nil {
		return nil
	}
	auth = Auth(token)
	return auth
}
