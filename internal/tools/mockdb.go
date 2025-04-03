package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]loginDetails{
	"Alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *loginDetails {
	time.Sleep(time.Second * 1)

	var clientData = loginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {

	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
