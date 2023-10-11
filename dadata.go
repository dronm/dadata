package dadata

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"errors"
	"bytes"
)

const(
	API_URL = "https://suggestions.dadata.ru/suggestions/api/4_1/rs"
	
	DEF_LANG = "ru"
	
	DEF_BY_ID_SEARCH_COUNT = 10
	
	ORG_TYPE_ENTERPRIZE = "LEGAL"
	ORG_TYPE_PERSON = "INDIVIDUAL"
)

const (
	BY_ID_SEARCH_TYPE_ALL = iota
	BY_ID_SEARCH_TYPE_ORG
	BY_ID_SEARCH_TYPE_IP
)
type ByIDSearchType int

const (
	BY_ID_SEARCH_BRANCH_TYPE_ALL = iota
	BY_ID_SEARCH_BRANCH_TYPE_MAIN
	BY_ID_SEARCH_BRANCH_TYPE_BRANCH
)

type ByIDSearchBranchType int

type Dadata struct {
	Key string `json:"key"`
}

func NewDadata(key string) *Dadata{
	return &Dadata{Key: key}
}
func (d *Dadata) AddHeaders(req *http.Request) error {
	if d.Key == "" {
		return errors.New("Key is not set")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+d.Key)
	return nil
}

//query ИНН или ОГРН
//count Количество результатов (максимум — 300)
//kpp КПП для поиска по филиалам
//branch_type Головная организация (MAIN) или филиал (BRANCH)
//serachType BY_ID_SEARCH_TYPE_ALL/BY_ID_SEARCH_TYPE_ORG/BY_ID_SEARCH_TYPE_IP
func (d *Dadata) FindById(query, kpp string, branchType ByIDSearchBranchType, serachType ByIDSearchType, count int) (*FindByIdResponse, error){
	if count == 0 {
		count = DEF_BY_ID_SEARCH_COUNT
	}
	client := &http.Client{}
	query_params := fmt.Sprintf(`"query": "%s"`, query)
	
	if kpp != "" {
		query_params+= ", " + fmt.Sprintf(`"kpp": "%s"`, kpp)
	}
	
	//
	if branchType != BY_ID_SEARCH_BRANCH_TYPE_ALL {
		var br_type string
		if branchType == BY_ID_SEARCH_BRANCH_TYPE_BRANCH {
			br_type = "BRANCH"
		}else{
			br_type = "MAIN"
		}
		query_params+= ", " + fmt.Sprintf(`"branch_type": "%s"`, br_type)
	}
	
	//
	if serachType != BY_ID_SEARCH_TYPE_ALL {
		var s_type string
		if serachType == BY_ID_SEARCH_TYPE_IP {
			s_type = ORG_TYPE_PERSON
		}else{
			s_type = ORG_TYPE_ENTERPRIZE
		}
		query_params+= ", " + fmt.Sprintf(`"type": "%s"`, s_type)
	}
	
	//
	if count > 0 {
		query_params+= ", " + fmt.Sprintf(`"count": %d`, count)
	}

	req, _ := http.NewRequest("POST", API_URL+"/findById/party", bytes.NewBuffer([]byte("{" + query_params + "}")))	
	
	if err := d.AddHeaders(req); err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := client.Do(req)	
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkForError(resp, body); err != nil {
		return nil, err
	}
//fmt.Println("body=", string(body))		
	var suggest FindByIdResponse
	if err := json.Unmarshal(body, &suggest); err != nil {
		return nil, err
	}
	return &suggest, nil		
	
}

func (d *Dadata) IPLocate(ip, lang string) (*IPLocateResponse, error){
	if lang == "" {
		lang = DEF_LANG
	}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/iplocate/address?ip=%s&language=%s", API_URL, ip, lang), nil)
	
	if err := d.AddHeaders(req); err != nil {
		return nil, err
	}
	resp, err := client.Do(req)	
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := checkForError(resp, body); err != nil {
		return nil, err
	}
	var loc IPLocateResponse
	if err := json.Unmarshal(body, &loc); err != nil {
		return nil, err
	}
	return &loc, nil		
}

//
func checkForError(resp *http.Response, body []byte) error{
	if resp.StatusCode != 200 {
//fmt.Println("body", string(body))	
		err_fields := struct {
			Family string `json:"family"`
			Reason string `json:"reason"`
			Message string `json:"message"`
		}{}
	
		if err := json.Unmarshal(body, &err_fields); err != nil {
			return err
		}
		return errors.New(fmt.Sprintf("HTTP error: %d, reason: %s, text: %s", resp.StatusCode, err_fields.Reason, err_fields.Message))
		
		//return errors.New(fmt.Sprintf("HTTP error: %d", resp.StatusCode))
	}
	return nil
}


