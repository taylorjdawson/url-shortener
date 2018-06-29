package urlutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Redirect struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Type        int    `json:"type`
}

type Firebase struct {
	Functions struct {
		Predeploy []string `json:"predeploy"`
	} `json:"functions"`
	Hosting struct {
		Public    string     `json:"public"`
		Ignore    []string   `json:"ignore"`
		Redirects []Redirect `json:"redirects"`
		Rewrites  []struct {
			Source   string `json:"source"`
			Function string `json:"function"`
		} `json:"rewrites"`
	} `json:"hosting"`
}

func getFirebaseJson(file string) Firebase {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fb Firebase
	json.Unmarshal(raw, &fb)
	return fb
}

// Helper function
func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func writeJSONToFile(json_data interface{}) error {
	data, _ := json.MarshalIndent(json_data, "", "  ")
	err := ioutil.WriteFile("output.json", data, 0644)
	return err
}

func (fb Firebase) listUrls() {

	urls := fb.Hosting.Redirects

	for _, element := range urls {
		fmt.Printf("%s\t->\t%s\n", element.Source, element.Destination)
	}
}

// Consider checking for dup source and dup dest

// Check for duplicates sources

// TODO: return fb.indexOf() != -1 DRY
func (fb *Firebase) containsUrl(source_url string) bool {

	//containsUrl := false

	/*urls := fb.Hosting.Redirects

	  for _ , url := range urls {
	    if url.Source == source_url {
	      return true
	    }
	  }*/

	return fb.indexOf(source_url) != -1
}

func (fb *Firebase) createUrlRedirect(source_url string, destination_url string) error {

	// Consider what else could 'go' wrong?

	var err error = nil

	if fb.containsUrl(source_url) {
		err = fmt.Errorf("Source url already exists '%s'", source_url)
	} else {
		fb.Hosting.Redirects = append(fb.Hosting.Redirects, Redirect{
			Source:      source_url,
			Destination: destination_url,
			Type:        302, //TODO: make Const
		})
	}
	return err
}

func (fb *Firebase) removeUrlRedirect(source_url string) error {
	var err error = nil

	i := fb.indexOf(source_url)

	if i == -1 {
		err = fmt.Errorf("url doesn't exists '%s'", source_url)
	} else {
		fb.Hosting.Redirects = append(fb.Hosting.Redirects[:i], fb.Hosting.Redirects[i+1:]...)
	}
	return err
}

// -1 if element doesn't exist
func (fb *Firebase) indexOf(source_url string) int {

	urls := fb.Hosting.Redirects

	for index, url := range urls {
		if url.Source == source_url {
			return index
		}
	}
	return -1 // TODO: Make const
}
