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
	auth := map[string]string{"id-prefix": "ci"}
	if token == "" {
		return nil
	}
	auth["Authorization"] = token
	return auth
}

// CorpAuth return CorpId and header of Auth
func CorpAuth() (corpID string, auth map[string]string, err error) {
	cid, token, err := corpToken()
	if err != nil {
		return "", nil, err
	}
	auth = Auth(token)
	return cid, auth, nil
}

// SuperAuth return header of Auth
func SuperAuth() (auth map[string]string, err error) {
	token, err := superToken()
	if err != nil {
		return nil, err
	}
	auth = Auth(token)
	return auth, nil
}

// CommonAuth common user header of the corp "109777231634510875",uid is "125424776930931207"
func CommonAuth() (auth map[string]string, err error) {
	email := "common.user@xdchen.com"
	password := "12345678"
	j, err := Login(email, "", password)
	if err != nil {
		return nil, err
	}
	token, _ := j.Get("token").String()
	auth = Auth(token)
	return auth, nil
}

// OtherAuth rcm user header of the other corp "125425254762820103", uid is "125425255299691015"
func OtherAuth() (auth map[string]string, err error) {
	email := "othercorp@xdchen.com"
	password := "12345678"
	j, err := Login(email, "", password)
	if err != nil {
		return nil, err
	}
	token, _ := j.Get("token").String()
	auth = Auth(token)
	return auth, nil
}
