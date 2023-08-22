package qdrant

import (
	"net/http"

	qdrant "github.com/qdrant/go-client/qdrant"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/vectorstores"

	"google.golang.org/grpc"
)

type Store struct {
	embedder          embeddings.Embedder
	grpcConn          *grpc.ClientConn
	collectionsClient qdrant.CollectionsClient
	pointsClient      qdrant.PointsClient

	textKey        string
	nameSpaceKey   string
	host           string
	schema         string
	collectionName string

	indexName        string
	apiKey           string
	connectionClient *http.Client
	queryAttrs       []string
	useGRPC          bool
}

var _ vectorstores.VectorStore = Store{}

// New creates a new Store with options.
// When using qdrant,
// the properties in the Class of qdrant must have properties with the values set by textKey and nameSpaceKey.
func New(opts ...Option) (Store, error) {
	s, err := applyClientOptions(opts...)
	if err != nil {
		return Store{}, err
	}

	s.collectionsClient = qdrant.NewCollectionsClient()

	return s, nil
}
