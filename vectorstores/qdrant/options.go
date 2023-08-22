package qdrant

import (
	"errors"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/embeddings"
)

const (
	_qdrantEnvVrName = "QDRANT_API_KEY"
	_defaultTextKey  = "text"
)

// ErrInvalidOptions is returned when the options given are invalid.
var ErrInvalidOptions = errors.New("invalid options")

// Option is a function type that can be used to modify the client.
type Option func(q *Store)

// WithEmbedder is an option for setting the embedder to use. Must be set.
func WithEmbedder(e embeddings.Embedder) Option {
	return func(q *Store) {
		q.embedder = e
	}
}

// WithIndexName is an option for specifying the index name. Must be set.
func WithIndexName(name string) Option {
	return func(q *Store) {
		q.indexName = name
	}
}

// WithCollctionName is an option for setting the collection used in qdrant. Must be set.
func WithCollectionName(collectionName string) Option {
	return func(q *Store) {
		q.collectionName = collectionName
	}
}

// WithAPIKey is an option for setting the api key. If the option is not set
// the api key is read from the PINECONE_API_KEY environment variable. If the
// variable is not present, an error will be returned.
func WithAPIKey(apiKey string) Option {
	return func(q *Store) {
		q.apiKey = apiKey
	}
}

// WithTextKey is an option for setting the text key in the metadata to the vectors
// in the index. The text key stores the text of the document the vector represents.
func WithTextKey(textKey string) Option {
	return func(q *Store) {
		q.textKey = textKey
	}
}

// Host is an option for setting the host to upsert and query the vectors
// from. Must be set.
func WithHost(host string) Option {
	return func(q *Store) {
		q.host = host
	}
}

// withGrpc is an official method to connect to the server.
func withGrpc() Option {
	return func(q *Store) {
		q.useGRPC = true
	}
}

func applyClientOptions(opts ...Option) (Store, error) {
	q := &Store{
		textKey: _defaultTextKey,
	}

	for _, opt := range opts {
		opt(q)
	}

	if q.indexName == "" {
		return Store{}, fmt.Errorf("%w: missing index name", ErrInvalidOptions)
	}

	if q.host == "" {
		return Store{}, fmt.Errorf("%w: missing host", ErrInvalidOptions)
	}

	if q.embedder == nil {
		return Store{}, fmt.Errorf("%w: missing embedder", ErrInvalidOptions)
	}

	if q.collectionName == "" {
		return Store{}, fmt.Errorf("%w: missing collection name", ErrInvalidOptions)
	}

	if q.apiKey == "" {
		q.apiKey = os.Getenv(_qdrantEnvVrName)
		if q.apiKey == "" {
			return Store{}, fmt.Errorf(
				"%w: missing api key. Pass it as an option or set the %s environment variable",
				ErrInvalidOptions,
				_qdrantEnvVrName,
			)
		}
	}

	return *q, nil
}
