package instagram

import (
	"context"
)

// IterateMedia makes pagination easy by converting the repeated
// api.NextMedias() call to a channel of media.  Media will be passed in the
// reverse order of individual requests, for instance GetUserRecentMedia will
// go in reverse CreatedTime order.  Use context to cancel iteration.
func (api *Api) IterateMedia(ctx context.Context, res *PaginatedMediasResponse) (<-chan *Media, <-chan error) {
	mediaChan := make(chan *Media)
	errChan := make(chan error, 1)

	go func() {
		defer close(mediaChan)
		defer close(errChan)

		for {
			if res == nil {
				return
			}
			if len(res.Medias) == 0 {
				return
			}

			// Iterate backwards
			for i := len(res.Medias) - 1; i >= 0; i-- {
				select {
				case <-ctx.Done():
					return
				case mediaChan <- &res.Medias[i]:
				}
			}

			// Paginate to next response
			var err error
			res, err = api.NextMedias(ctx, res.Pagination)
			if err != nil {
				errChan <- err
				return
			}
		}
	}()

	return mediaChan, errChan
}
