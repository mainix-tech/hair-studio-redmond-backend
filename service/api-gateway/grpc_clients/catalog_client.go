package grpc_clients

import (
	pb "hair-studio-redmond/shared/proto/catalog"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CatalogServiceClient struct {
	Client pb.CatalogServiceClient
	conn   *grpc.ClientConn
}

func NewCatalogServiceClient() (*CatalogServiceClient, error) {
	catalogServiceURL := os.Getenv("CATALOG_SERVICE_URL")
	if catalogServiceURL == "" {
		catalogServiceURL = "catalog-service:9095"
	}

	conn, err := grpc.NewClient(catalogServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewCatalogServiceClient(conn)

	return &CatalogServiceClient{
		Client: client,
		conn:   conn,
	}, nil
}

func (c *CatalogServiceClient) Close() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}
