// Steve Phillips / elimisteve
// 2012.10.08

package webhooks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	VERBOSE = false
)

var ChPayloads = make(chan *GitHubPayload)

//
// Accept GitHub Webhook data
//
func WebhookHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		info := fmt.Sprintf("req:\n\n%+#v\n", req)
		fmt.Printf("%v\n", info)
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Error in WebhookHandler: %v\n", err)
		return
	}
	defer req.Body.Close()

	payload, err := BodyToPayload(body)
	if err != nil {
		fmt.Printf("Returned from BodyToCommit: %v\n", err)
		return
	}
	ChPayloads <- payload
	return
}

// WebhookListener call WebhookHandler, which takes JSON HTTP POSTs,
// parses the relevant data, then sends it over the given channel
func WebhookListener(port string, payloads chan *GitHubPayload) {
	// FIXME: Can using global variables possibly be a non-bad idea?
	ChPayloads = payloads
	http.HandleFunc("/webhook", WebhookHandler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error in WebhookListener ListenAndServe: %v\n", err)
	}
}

func webhookDataToPayload(data string) (*GitHubPayload, error) {
	payload := GitHubPayload{}
	err := json.Unmarshal([]byte(data), &payload)
	if err != nil {
		msg := "Couldn't unmarshal WebHook data: " + data
		log.Printf(msg)
		return nil, err
	}
	return &payload, nil
}

func PayloadToCommit(payload *GitHubPayload) *GitCommit {
	// Parse these pieces:
	// commit.Author    = payload[pusher][name]
	// commit.Email     = payload[pusher][email]
	// // commit.Author    = payload[head_commit][author][username]
	// // commit.Email     = payload[head_commit][author][email]
	// commit.Message   = payload[head_commit][message]
	// commit.Repo      = payload[repository][name]
	// commit.RepoOwner = payload[repository][owner][name]

	cleanMessage := strings.Replace(payload.HeadCommit.Message,
		"\n", "    ", -1)

	commit := GitCommit{
		Author:    *payload.Pusher.Name,
		Email:     *payload.Pusher.Email,
		Message:   strings.Replace(cleanMessage, "\n", "    ", -1),
		Repo:      payload.Repository.Name,
		RepoOwner: *payload.Repository.Owner.Name,
	}
	return &commit
}

func bodyToWebhookData(body []byte) (string, error) {
	values, err := url.ParseQuery(string(body))
	if err != nil {
		return "", fmt.Errorf("Error parsing body '%s': %v\n", body, err)
	}
	// Parsing GitHub-specific format
	data := ""
	_, contained := values["payload"]
	if contained {
		data = values["payload"][0]
	} else {
		return "", fmt.Errorf("Poorly-formed webhook callback: '%+v'\n", values)
	}
	return data, nil
}

func BodyToPayload(body []byte) (*GitHubPayload, error) {
	data, err := bodyToWebhookData(body)
	if err != nil {
		return nil, err
	}
	payload, err := webhookDataToPayload(data)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func BodyToCommit(body []byte) (*GitCommit, error) {
	payload, err := BodyToPayload(body)
	if err != nil {
		return nil, err
	}
	commit := PayloadToCommit(payload)
	return commit, nil
}
