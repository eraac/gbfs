package gbfs

// Error type
type Error string

// List of errors the gbfs.Client can return
const (
	ErrBaseURLMissing Error = "base url is missing"
	ErrFeedNotExist   Error = "feed not exist"
	ErrInvalidFeed    Error = "invalid feed"
)

// Error return the error formatted in string
func (e Error) Error() string {
	return string(e)
}
