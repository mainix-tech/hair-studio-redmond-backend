package grpc_clients

import (
	pb "hair-studio-redmond/shared/proto/menu"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MenuServiceClient struct {
	Client pb.MenuServiceClient
	conn   *grpc.ClientConn
}

func NewMenuServiceClient() (*MenuServiceClient, error) {
	menuServiceURL := os.Getenv("MENU_SERVICE_URL")
	if menuServiceURL == "" {
		menuServiceURL = "menu-service:9094"
	}

	conn, err := grpc.NewClient(menuServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewMenuServiceClient(conn)

	return &MenuServiceClient{
		Client: client,
		conn:   conn,
	}, nil
}

func (c *MenuServiceClient) Close() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}
