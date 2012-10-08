// Steve Phillips / elimisteve
// 2012.10.08

package main

import (
	"fmt"
	"github.com/elimisteve/go.github/webhooks"
)

const CURL_COMMAND = `    $ curl -X POST -d 'payload=%7B%22pusher%22%3A%7B%22name%22%3A%22elimisteve%22%2C%22email%22%3A%22elimisteve%40gmail.com%22%7D%2C%22repository%22%3A%7B%22name%22%3A%22go-helpers%22%2C%22created_at%22%3A%222012-01-27T22%3A55%3A30-08%3A00%22%2C%22size%22%3A396%2C%22has_wiki%22%3Atrue%2C%22private%22%3Afalse%2C%22watchers%22%3A2%2C%22language%22%3A%22Go%22%2C%22url%22%3A%22https%3A%2F%2Fgithub.com%2Fprototypemagic%2Fgo-helpers%22%2C%22fork%22%3Afalse%2C%22id%22%3A3289113%2C%22pushed_at%22%3A%222012-10-08T13%3A20%3A47-07%3A00%22%2C%22has_downloads%22%3Atrue%2C%22open_issues%22%3A0%2C%22has_issues%22%3Atrue%2C%22homepage%22%3A%22prototypemagic.com%22%2C%22description%22%3A%22ProtoType%20Magic%27s%20Go%20helper%20functions%22%2C%22organization%22%3A%22prototypemagic%22%2C%22stargazers%22%3A2%2C%22forks%22%3A0%2C%22owner%22%3A%7B%22name%22%3A%22prototypemagic%22%2C%22email%22%3Anull%7D%7D%2C%22forced%22%3Afalse%2C%22after%22%3A%22e7b5ff46b3157d741f90c8158a121ea669e35d7a%22%2C%22head_commit%22%3A%7B%22added%22%3A%5B%5D%2C%22modified%22%3A%5B%22README.md%22%5D%2C%22timestamp%22%3A%222012-10-08T13%3A20%3A55-07%3A00%22%2C%22author%22%3A%7B%22name%22%3A%22Steve%20Phillips%22%2C%22username%22%3A%22elimisteve%22%2C%22email%22%3A%22elimisteve%40gmail.com%22%7D%2C%22removed%22%3A%5B%5D%2C%22url%22%3A%22https%3A%2F%2Fgithub.com%2Fprototypemagic%2Fgo-helpers%2Fcommit%2Fe7b5ff46b3157d741f90c8158a121ea669e35d7a%22%2C%22id%22%3A%22e7b5ff46b3157d741f90c8158a121ea669e35d7a%22%2C%22distinct%22%3Atrue%2C%22message%22%3A%225th%20trivial%20commit%20to%20README.md%22%2C%22committer%22%3A%7B%22name%22%3A%22Steve%20Phillips%22%2C%22username%22%3A%22elimisteve%22%2C%22email%22%3A%22elimisteve%40gmail.com%22%7D%7D%2C%22deleted%22%3Afalse%2C%22ref%22%3A%22refs%2Fheads%2Fmaster%22%2C%22commits%22%3A%5B%7B%22added%22%3A%5B%5D%2C%22modified%22%3A%5B%22README.md%22%5D%2C%22timestamp%22%3A%222012-10-08T13%3A20%3A55-07%3A00%22%2C%22author%22%3A%7B%22name%22%3A%22Steve%20Phillips%22%2C%22username%22%3A%22elimisteve%22%2C%22email%22%3A%22elimisteve%40gmail.com%22%7D%2C%22removed%22%3A%5B%5D%2C%22url%22%3A%22https%3A%2F%2Fgithub.com%2Fprototypemagic%2Fgo-helpers%2Fcommit%2Fe7b5ff46b3157d741f90c8158a121ea669e35d7a%22%2C%22id%22%3A%22e7b5ff46b3157d741f90c8158a121ea669e35d7a%22%2C%22distinct%22%3Atrue%2C%22message%22%3A%225th%20trivial%20commit%20to%20README.md%22%2C%22committer%22%3A%7B%22name%22%3A%22Steve%20Phillips%22%2C%22username%22%3A%22elimisteve%22%2C%22email%22%3A%22elimisteve%40gmail.com%22%7D%7D%5D%2C%22before%22%3A%225b5bf85f6d8336b2b7e3a1a1574e70e101535f5f%22%2C%22compare%22%3A%22https%3A%2F%2Fgithub.com%2Fprototypemagic%2Fgo-helpers%2Fcompare%2F5b5bf85f6d83...e7b5ff46b315%22%2C%22created%22%3Afalse%7D' localhost:7777/webhook`

func main() {
	payloads := make(chan *webhooks.GitHubPayload)
	go webhooks.WebhookListener("7777", payloads)

	fmt.Printf("Run this to simulate a WebHook callback from GitHub:\n%s\n\n",
		CURL_COMMAND)

	payload := <-payloads
	fmt.Printf("Unmarshal'd callback data:\n%+v\n\n", payload)

	commit := webhooks.PayloadToCommit(payload)
	fmt.Printf("Sample use of simplified commit data:\n\n")
	fmt.Printf("%s just pushed to %s/%s on GitHub: '%s'\n",
		commit.Author, commit.RepoOwner, commit.Repo, commit.Message)
}
