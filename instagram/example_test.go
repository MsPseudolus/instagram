package instagram

import (
	"fmt"
)

// ExampleNew sets up the whole instagram API
func ExampleNew() {
	apiAuthenticatedUser := New("client_key", "secret", "", true)
	if ok, err := apiAuthenticatedUser.VerifyCredentials(); !ok {
		panic(err)
	}
	fmt.Println("Successfully created instagram.Api with user credentials")
}

// ExampleApi_IterateMedia shows how to use iteration on a channel to avoid the complex pagination calls
func ExampleApi_IterateMedia() {
	// *** or ***
	api := New("client_id", "client_secret", "access_token", true)

	mediasResponse, err := api.GetRecentMedia(nil)
	if err != nil {
		panic(err)
	}

	// Stop 30 days ago
	doneChan := make(chan bool)

	mediaIter, errChan := api.IterateMedia(mediasResponse, doneChan /* optional */)
	for media := range mediaIter {
		processMedia(media)

		if isDone(media) {
			close(doneChan) // Signal to iterator to quit
			break
		}
	}

	// When mediaIter is closed, errChan will either have a single error on it or it will have been closed so this is safe.
	if err := <-errChan; err != nil {
		panic(err)
	}
}

func processMedia(m *Media) {}
func isDone(m *Media) bool {
	return false
}
func processUser(u *User) {}
