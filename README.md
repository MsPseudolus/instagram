# instagram

An [instagram](http://instagram.com) [API](http://instagram.com/developer) wrapper.

## Features

Based on [yanatan16/golang-instagram](https://github.com/yanatan16/golang-instagram).
Modernized and updated to match changes to the Instagram API.

Changes from the original:

* [x] Removed deprecated API endpoints (which are many)
* [x] Added support for new media type "carousel"
* [ ] Added support for OAuth flow to retrieve access token
* [x] Added go module
* [x] Added support for non-default `http.Client`
* [x] Added support for `context.Context`

## Testing

For the API, there are fake tests and integration tests. The fake tests swap
fixtures JSON responses via an `httptest` server to cover specific caess. The
integration tests use provided credentials to exercise the API in general ways
(not specific to the authorized user).

```
# Run unit and fake tests
make test

# Run unit, fake, and integration tests (requires authentication, see below).
make test-integration
```

#### Authenticating

Set the following environment variables to run the integration tests and
command line tool.

```
export TEST_INSTAGRAM_CLIENT_ID=<client id>
export TEST_INSTAGRAM_CLIENT_SECRET=<client secret>
export TEST_INSTAGRAM_ACCESS_TOKEN=<access token>
```

#### Adding a test fixture

The `cmd` package contains a simple command line tools to query endpoints. You
can use it to return sample data, then save it as a new fixture. Here's the command used
to add the 'video' fixture.

```
go run cmd/list/main.go -count 1 -maxid 1979319391662961209_11073382793 -raw | jq . > testdata/video.json
```

## License

MIT. See LICENSE file.
