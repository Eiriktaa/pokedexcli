package internals

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internals/configs/context"
	"time"
)

func httpRequestUnmarshal[T any](res *http.Response, data *T) ([]byte, error) {
	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.Unmarshal(byteData, data); err != nil {
		return nil, err
	}

	return byteData, nil
}

func httpDoGetRequest[T any](r *context.ReplContext, url string, data *T) error {

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("request failed with %d code", res.StatusCode)
	}

	byteData, err := httpRequestUnmarshal(res, &data)
	if err != nil {
		return fmt.Errorf("failed to decode json object: %v", err)
	}

	r.Cache.AddData(url, byteData)
	return nil
}

// gets location data from
// preforms a get requests with caching, and handles unmarharling into the dataType param
// caller is responible to make sure object matches response struct
func HttpGetApiDataWithUnmarshal[T any](r *context.ReplContext, url string, dataType *T) error {

	if data, found := r.Cache.GetData(url); found {
		if err := json.Unmarshal(data, &dataType); err != nil {
			return err
		}
	} else {
		// not in cache
		if err := httpDoGetRequest(r, url, dataType); err != nil {
			return err
		}
	}
	return nil

}
