package grpc_clients

import (
	pb "hair-studio-redmond/shared/proto/profile"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProfileServiceClient struct {
	Client pb.ProfileServiceClient
	conn   *grpc.ClientConn
}

func NewProfileServiceClient() (*ProfileServiceClient, error) {
	profileServiceURL := os.Getenv("PROFILE_SERVICE_URL")
	if profileServiceURL == "" {
		profileServiceURL = "profile-service:9093"
	}

	conn, err := grpc.NewClient(profileServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewProfileServiceClient(conn)

	return &ProfileServiceClient{
		Client: client,
		conn:   conn,
	}, nil
}

func (c *ProfileServiceClient) Close() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}
