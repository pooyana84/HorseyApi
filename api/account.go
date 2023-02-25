package api 

import "net/http"
import "encoding/json"
import "io/ioutil"

func (profile *Profile) GetProfile (token string, client *http.Client) (string) {
	request, err := NewRequest(token, "GET", "LichessUrl" + "/api/account", nil)
	checkErr(err)
	response, err := client.Do(request)
	checkErr(err)
	defer response.Body.Close()
	responseBody , err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(responseBody, profile)
	checkErr(err)
	return response.Status
} 
