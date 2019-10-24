package services

type Account struct {
	AccessKeyID string
}

func GetAccountInfo() Account {
	session := GetConfig()
	cred, err := session.Config.Credentials.Get()
	if err != nil {
		panic("Could not get your credentials, " + err.Error())
	}
	account := Account{AccessKeyID: cred.AccessKeyID}
	return account
}
