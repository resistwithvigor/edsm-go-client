package edsm_go_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/resistwithvigor/edsm-go-client/logging"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//var (
//	urlSphereSystems = func(systemName string, radius float64) string {
//		return fmt.Sprintf("%s?coords=1&showid=1&radius=%s&systemName=%s", EDSMSystemsSphere, url.QueryEscape(strconv.FormatFloat(radius, 'f', 2, 64)), url.QueryEscape(systemName))
//	}
//
//	urlSystem = func(systemName string) string {
//		return fmt.Sprintf("%s?coords=1&systemName=%s", EDSMSystemsSingle, url.QueryEscape(systemName))
//	}
//)

func NewEDSMClient(debug bool, userAgent string) EDSMClient {
	return EDSMClient{debug, userAgent}
}

func (client *EDSMClient) request(method string, url string, b []byte) (response []byte, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))

	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", client.UserAgent)

	httpClient := &http.Client{Timeout: 120 * time.Second}

	res, err := httpClient.Do(req)

	if err != nil {
		logging.Warning.Println("Request error", err)
		err = errors.New("http request returned an error")
		return
	}

	defer res.Body.Close()

	response, err = ioutil.ReadAll(res.Body)

	if err != nil {
		logging.Warning.Println("Request error", err)
		err = errors.New("error while reading body")
		return
	}

	if client.Debug {
		logging.Trace.Printf("API REQUEST\tURL :: %s\n", url)
		logging.Trace.Printf("API RESPONSE\tSTATUS :: %s\n", res.Status)
		for k, v := range res.Header {
			logging.Trace.Printf("API RESPONSE\tHEADER :: [%s] = %+v\n", k, v)
		}
		logging.Trace.Printf("API RESPONSE\tBODY :: [%s]\n", response)
	}
	return
}

type EDSMServerStatusDate time.Time

func (sd *EDSMServerStatusDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 03:04:05", s)
	if err != nil {
		return err
	}
	*sd = EDSMServerStatusDate(t)
	return nil
}

func (sd EDSMServerStatusDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(sd)
}

func (sd EDSMServerStatusDate) Format(s string) string {
	t := time.Time(sd)
	return t.Format(s)
}
