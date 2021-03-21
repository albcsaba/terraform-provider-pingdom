package pingdom

import (
	"albcsaba/terraform-provider-pingdom/pkg/pingdom/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

const (
	GET_ALL_CHECKS = "/api/3.1/checks?include_tags=true&inclue_severity=true"
)

func GetAllChecks(client Client) ([]models.Check, error) {
	res, err := client.Do(GET_ALL_CHECKS, http.MethodGet, nil)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, errors.New("Cannot fetch checks. StatusCode: " + strconv.Itoa(res.StatusCode))
	}
	var unmarshalledResult models.Checks
	err = json.Unmarshal(res.Body, &unmarshalledResult)
	if err != nil {
		return nil, err
	}

	return unmarshalledResult.Checks, nil
}
