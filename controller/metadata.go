package controller

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"yamlrest/utils/context"
	"yamlrest/utils/model"
	"yamlrest/utils/validation"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type Server struct {
	Context *context.AppContext
	Routers *mux.Router
}

// Create server instance and route handlers
func CreateServer(ctx *context.AppContext) *Server {
	server := &Server{
		Context: ctx,
		Routers: mux.NewRouter(),
	}
	server.paths()
	return server
}

func (s *Server) paths() {
	s.Routers.HandleFunc("/appData", s.createAppData).Methods("POST")
	s.Routers.HandleFunc("/appData", s.getAppData).Methods("GET")
}

func (s *Server) getAppData(res http.ResponseWriter, req *http.Request) {

	queryStr := req.URL.Query()
	data := s.Context.Database.RetrieveAllData()
	var searchOutput []interface{}

	// No search parameter, return all records
	if len(queryStr) == 0 {
		for _, data := range data {
			searchOutput = append(searchOutput, data)
		}
		yaml.NewEncoder(res).Encode(searchOutput)
		return
	}

	// If search parameter matches our index, use that
	if i, ok := queryStr["source"]; ok {
		yaml.NewEncoder(res).Encode(append(searchOutput, s.Context.Database.Read(i[0])))
		return
	}

	// Do manual search by taking each record and checking with given query param(s)
	for _, data := range data {
		if manualSearch(data, queryStr) {
			searchOutput = append(searchOutput, data)
		}
	}
	yaml.NewEncoder(res).Encode(searchOutput)
	return
}

func (s *Server) createAppData(res http.ResponseWriter, req *http.Request) {

	if req.Body == nil {
		s.returnError(res, "Cannot work with empty payload")
		return
	}

	var bodyBytes []byte
	bodyBytes, _ = io.ReadAll(req.Body)

	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	bodyString := string(bodyBytes)

	var bodyMetadata = model.Metadata{}
	err := yaml.Unmarshal([]byte(bodyString), &bodyMetadata)

	if err != nil {
		s.returnError(res, "Error in unmarshalling YAML body")
		return
	}

	if isValid, errorStr := validation.ValidateMetadata(bodyMetadata); !isValid {
		s.returnError(res, errorStr)
		return
	}

	// For the key-value data store, Source is chosen as the key as it is the closest piece of information to being unique. Thus, we effectively built an index around source
	s.Context.Database.Create(bodyMetadata.Source, bodyMetadata)
	fmt.Println("Record created successfully")
	res.WriteHeader(http.StatusOK)
	yaml.NewEncoder(res).Encode(`Record created successfully`)
	return
}

func (s *Server) returnError(res http.ResponseWriter, errString string) {
	fmt.Println("Invalid request received: ", errString)
	res.WriteHeader(http.StatusBadRequest)
	yaml.NewEncoder(res).Encode(errString)
	return
}

func manualSearch(data interface{}, urlQuerystr map[string][]string) bool {

	searchSpace := data.(model.Metadata)

	// Goal is to take every param and match it against current record to see if its a match
	for param, value := range urlQuerystr {
		switch queryParam := strings.ToLower(strings.TrimSpace(param)); queryParam {

		// Partial query search for title and description, full match required for remaining fields
		case "title":
			if !strings.Contains(strings.ToLower(searchSpace.Title), strings.ToLower(value[0])) {
				return false
			}
		case "version":
			if strings.ToLower(searchSpace.Version) != strings.ToLower(value[0]) {
				return false
			}
		case "maintainers.name":
			for _, name := range value {
				hit := false
				for i := range searchSpace.Maintainers {
					if strings.ToLower(searchSpace.Maintainers[i].Name) == strings.ToLower(name) {
						hit = true
						break
					}
				}
				if hit == false {
					return false
				}
			}
		case "maintainers.email":
			for _, email := range value {
				hit := false
				for i := range searchSpace.Maintainers {
					if strings.ToLower(searchSpace.Maintainers[i].Email) == strings.ToLower(email) {
						hit = true
						break
					}
				}
				if hit == false {
					return false
				}
			}
		case "company":
			if strings.ToLower(searchSpace.Company) != strings.ToLower(value[0]) {
				return false
			}
		case "website":
			if strings.ToLower(searchSpace.Website) != strings.ToLower(value[0]) {
				return false
			}
		case "license":
			if strings.ToLower(searchSpace.License) != strings.ToLower(value[0]) {
				return false
			}
		case "description":
			if !strings.Contains(strings.ToLower(searchSpace.Description), strings.ToLower(value[0])) {
				return false
			}
		}
	}
	return true
}
