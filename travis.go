package githubbot

import (
	"fmt"

	"github.com/cpalone/travishook"
)

func (s *Session) travisServer(port int) {
	server := travishook.NewServer(port, "/travishook")
	server.GoListenAndServe()
	for {
		p := <-server.Out
		fmt.Printf("Received payload with status: %s\n", p.Status)
		s.sendMessage(fmt.Sprintf(
			"[travis | %s | %s ] Status: %s",
			p.Repository.Name, p.Branch, p.Status), "")
	}
}