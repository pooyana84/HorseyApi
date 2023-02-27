package api 

import "encoding/json"
import "io/ioutil"
import "strconv"

func (session *Session) GetProfile (account *Account) (string) {
	request, err := NewRequest(session.Token, "GET", LichessUrl + "/api/account", nil)
	checkErr(err)
	response, err := session.Client.Do(request)
	checkErr(err)
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	err = json.Unmarshal(responseBody, account)
	checkErr(err)
	return response.Status
}
func (session *Session) GetEmail (email *Email) (string) {
	request, err := NewRequest(session.Token, "GET", LichessUrl + "/api/account/email", nil)
	checkErr(err)
	response, err := session.Client.Do(request)
	checkErr(err)
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	err = json.Unmarshal(responseBody, email)
	checkErr(err)
	return response.Status
}


func (session *Session) GetPreferences (prefs *Prefs) (string) {
	request, err := NewRequest(session.Token, "GET", LichessUrl + "/api/account/preferences", nil)
	checkErr(err)
	response, err := session.Client.Do(request)
	checkErr(err)
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	err = json.Unmarshal(responseBody, prefs)
	checkErr(err)
	return response.Status
}


func (session *Session) GetKidMode (kid *Kid) (string) {
	request, err := NewRequest(session.Token, "GET", LichessUrl + "/api/account/kid", nil)
	checkErr(err)
	response, err := session.Client.Do(request)
	checkErr(err)
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	err = json.Unmarshal(responseBody, kid)
	checkErr(err)
	return response.Status
}

func (session *Session) SetKidMode (v bool) (string) {
	request, err := NewRequest(session.Token, "POST", LichessUrl + "/api/account/kid", nil)
	checkErr(err)
	urlData := request.URL.Query()
	urlData.Add("v", strconv.FormatBool(v))
	request.URL.RawQuery = urlData.Encode()
	response, err := session.Client.Do(request)
	checkErr(err)
	response.Body.Close()
	return response.Status
}
