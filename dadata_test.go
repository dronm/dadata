package dadata

import(
	"fmt"
	"testing"
	"encoding/json"
	"os"
)

const(
	ENV_VAR_DADATA_KEY = "DADATA_KEY"
	ENV_VAR_DADATA_TEST_INN = "DADATA_TEST_INN"
	ENV_VAR_DADATA_TEST_IP = "DADATA_TEST_IP"
)

func PrintStruct(str interface{}) error {
	res, err := json.Marshal(str)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}



func TestIPLocate(t *testing.T) {
	dadata := Dadata{Key: os.Getenv(ENV_VAR_DADATA_KEY)}	
	loc, err := dadata.IPLocate(os.Getenv(ENV_VAR_DADATA_TEST_IP),"")
	if err != nil {
		t.Fatalf("%v", err)
	}
	
	if err := PrintStruct(loc); err != nil {
		t.Fatalf("%v", err)
	}
}


func TestFindById(t *testing.T) {
	dadata := Dadata{Key: os.Getenv(ENV_VAR_DADATA_KEY)}	
	loc, err := dadata.FindById(os.Getenv(ENV_VAR_DADATA_TEST_INN), "", BY_ID_SEARCH_BRANCH_TYPE_ALL, BY_ID_SEARCH_TYPE_ALL, 0)
	if err := PrintStruct(loc); err != nil {
		t.Fatalf("%v", err)
	}

	if err != nil {
		t.Fatalf("%v", err)
	}
}

