package pingdom

import (
	"albcsaba/terraform-provider-pingdom/pkg/pingdom/models"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetAllChecks(t *testing.T) {

	/* Load static testfiles */
	absolutePath, _ := filepath.Abs("../static_testfiles/getAllChecks.json")
	testFile, err := os.Open(absolutePath)
	if err != nil {
		log.Fatalf("Cannot open testfile: %v", err)
	}
	defer testFile.Close()

	byteValue, _ := ioutil.ReadAll(testFile)
	var getAllCheckResult models.Checks
	err = json.Unmarshal(byteValue, &getAllCheckResult)
	if err != nil {
		log.Fatalf("Cannot unmarshall static testfile: %v", err)
	}

	/* Mock the client */
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockClient := NewMockClient(mockController)

	mockClient.EXPECT().Do(GET_ALL_CHECKS, http.MethodGet, nil).Return(
		&Response{
			Body:       byteValue,
			StatusCode: http.StatusOK,
		}, nil).AnyTimes()

	/* Test cases */
	type args struct {
		client Client
	}

	tests := []struct {
		name    string
		args    args
		want    []models.Check
		wantErr bool
	}{
		{
			name:    "GetAllProjects",
			args:    args{client: mockClient},
			want:    getAllCheckResult.Checks,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllChecks(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllChecks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllChecks() got = %v, want %v", got, tt.want)
			}
		})
	}
}
