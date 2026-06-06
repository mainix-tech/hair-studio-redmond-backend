package grpc

import (
	"context"
	"hair-studio-redmond/service/catalog-service/internal/domain"
	pb "hair-studio-redmond/shared/proto/catalog"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCHandler struct {
	pb.UnimplementedCatalogServiceServer
	service domain.CatalogService
}

func NewGRPCHandler(server *grpc.Server, service domain.CatalogService) *gRPCHandler {
	handler := &gRPCHandler{
		service: service,
	}

	pb.RegisterCatalogServiceServer(server, handler)

	return handler
}

func (h *gRPCHandler) GetCatalogItems(ctx context.Context, req *emptypb.Empty) (*pb.GetCatalogItemsResponse, error) {
	log.Printf("GetCatalogItems")

	// 1. Fetch your domain items slice from the service layer
	catalogModels, err := h.service.GetCatalogItems(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get catalog items: %v", err)
	}

	// 2. Map your array/slice of domain models back into your protobuf nested array
	var pbItems []*pb.CatalogItem
	for _, m := range catalogModels {
		pbItems = append(pbItems, &pb.CatalogItem{
			Id:       m.ID,
			Title:    m.Title,
			ImgUrl:   m.ImgUrl,
			Category: m.Category,
		})
	}

	return &pb.GetCatalogItemsResponse{
		Items: pbItems,
	}, nil
}

func (h *gRPCHandler) CreateCatalogItem(ctx context.Context, req *pb.CreateCatalogItemRequest) (*emptypb.Empty, error) {
	log.Printf("CreateCatalogItem %+v", req)

	// 1. Safety check: Verify the incoming nested ProfileInfo object is not nil
	catalogItem := req.GetItem()
	if catalogItem == nil {
		return nil, status.Error(codes.InvalidArgument, "Catalog item payload cannot be empty")
	}

	model := &domain.CatalogCreate{
		Title:    catalogItem.GetTitle(),
		Category: catalogItem.GetCategory(),
		ImgUrl:   catalogItem.GetImgUrl(),
	}

	// 3. Pass the pointer to your domain model into your service layer
	if err := h.service.CreateCatalogItem(ctx, model); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update menu item: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (h *gRPCHandler) UpdateCatalogItem(ctx context.Context, req *pb.UpdateCatalogItemRequest) (*emptypb.Empty, error) {
	log.Printf("UpdateCatalogItemRequest %+v", req)

	// 1. Safety check: Verify the incoming nested ProfileInfo object is not nil
	catalogItem := req.GetItem()
	if catalogItem == nil {
		return nil, status.Error(codes.InvalidArgument, "catalog item payload cannot be empty")
	}

	model := &domain.CatalogModel{
		ID:       catalogItem.GetId(),
		Title:    catalogItem.GetTitle(),
		Category: catalogItem.GetCategory(),
		ImgUrl:   catalogItem.GetImgUrl(),
	}

	// 3. Pass the pointer to your domain model into your service layer
	if err := h.service.UpdateCatalogItem(ctx, model); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update catalog item: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (h *gRPCHandler) DeleteCatalogItem(ctx context.Context, req *pb.DeleteCatalogItemRequest) (*emptypb.Empty, error) {
	log.Printf("DeleteCatalogItem %+v", req)

	// 1. Safety check: Verify the incoming nested ProfileInfo object is not nil
	catalogId := req.GetId()
	if catalogId == "" {
		return nil, status.Error(codes.InvalidArgument, "catalog id cannot be empty")
	}

	model := &domain.CatalogDelete{
		ID: catalogId,
	}

	// 3. Pass the pointer to your domain model into your service layer
	if err := h.service.DeleteCatalogItem(ctx, model); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete catalog item: %v", err)
	}

	return &emptypb.Empty{}, nil
}
