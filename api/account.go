package api 

import "net/http"
import "encoding/json"
import "io/ioutil"

func (account *Account) GetProfile (token string, client *http.Client) (string) {
	request, err := NewRequest(token, "GET", LichessUrl + "/api/account", nil)
	checkErr(err)
	response, err := client.Do(request)
	checkErr(err)
	defer response.Body.Close()
	responseBody , err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(responseBody, account)
	checkErr(err)
	return response.Status
} 
