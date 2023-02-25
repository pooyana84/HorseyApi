package api 

import "net/http"
import "encoding/json"
import "io/ioutil"
import "strconv"

func (account *Account) GetProfile (token string, client *http.Client) (string) {
	request, err := NewRequest(token, "GET", LichessUrl + "/api/account", nil)
	checkErr(err)
	response, err := client.Do(request)
	checkErr(err)
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	err = json.Unmarshal(responseBody, account)
	checkErr(err)
	return response.Status
}

func (email *Email) GetEmail (token string, client *http.Client) (string) {
	request, err := NewRequest(token, "GET", LichessUrl + "/api/account/email", nil)
	checkErr(err)
	response, err := client.Do(request)
	checkErr(err)
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	err = json.Unmarshal(responseBody, email)
	checkErr(err)
	return response.Status
}


func (prefs *Prefs) GetPreferences (token string, client *http.Client) (string) {
	request, err := NewRequest(token, "GET", LichessUrl + "/api/account/preferences", nil)
	checkErr(err)
	response, err := client.Do(request)
	checkErr(err)
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	err = json.Unmarshal(responseBody, prefs)
	checkErr(err)
	return response.Status
}


func (kid *Kid) GetKidMode (token string, client *http.Client) (string) {
	request, err := NewRequest(token, "GET", LichessUrl + "/api/account/kid", nil)
	checkErr(err)
	response, err := client.Do(request)
	checkErr(err)
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	err = json.Unmarshal(responseBody, kid)
	checkErr(err)
	return response.Status
}

func SetKidMode (v bool, token string, client *http.Client) (string) {
	request, err := NewRequest(token, "POST", LichessUrl + "/api/account/kid", nil)
	checkErr(err)
	urlData := request.URL.Query()
	urlData.Add("v", strconv.FormatBool(v))
	request.URL.RawQuery = urlData.Encode()
	response, err := client.Do(request)
	checkErr(err)
	response.Body.Close()
	return response.Status
}
