package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"hair-studio-redmond/service/menu-service/internal/domain"
	pb "hair-studio-redmond/shared/proto/menu"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCHandler struct {
	pb.UnimplementedMenuServiceServer
	service domain.MenuService
}

func NewGRPCHandler(server *grpc.Server, service domain.MenuService) *gRPCHandler {
	handler := &gRPCHandler{
		service: service,
	}

	pb.RegisterMenuServiceServer(server, handler)

	return handler
}

func (h *gRPCHandler) GetMenuItems(ctx context.Context, req *emptypb.Empty) (*pb.GetMenuItemsResponse, error) {
	log.Printf("GetMenuItems")

	// 1. Fetch your domain items slice from the service layer
	menuModels, err := h.service.GetMenuItems(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get menu items: %v", err)
	}

	// 2. Map your array/slice of domain models back into your protobuf nested array
	var pbItems []*pb.MenuItem
	for _, m := range menuModels {
		pbItems = append(pbItems, &pb.MenuItem{
			Id:                  m.ID,
			Name:                m.Name,
			PreviewDescription:  m.PreviewDescription,
			DetailedDescription: m.DetailedDescription,
			Time:                m.Time,
			Price:               m.Price,
			Category:            m.Category,
		})
	}

	return &pb.GetMenuItemsResponse{
		Items: pbItems,
	}, nil
}

func (h *gRPCHandler) CreateMenuItem(ctx context.Context, req *pb.CreateMenuItemRequest) (*emptypb.Empty, error) {
	log.Printf("CreateMenuItemRequest %+v", req)

	// 1. Safety check: Verify the incoming nested ProfileInfo object is not nil
	menuItem := req.GetItem()
	if menuItem == nil {
		return nil, status.Error(codes.InvalidArgument, "Menu item payload cannot be empty")
	}

	model := &domain.MenuCreate{
		Name:                menuItem.GetName(),
		PreviewDescription:  menuItem.GetPreviewDescription(),
		DetailedDescription: menuItem.GetDetailedDescription(),
		Time:                menuItem.GetTime(),
		Price:               menuItem.GetPrice(),
		Category:            menuItem.GetCategory(),
	}

	// 3. Pass the pointer to your domain model into your service layer
	if err := h.service.CreateMenuItem(ctx, model); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update menu item: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (h *gRPCHandler) UpdateMenuItem(ctx context.Context, req *pb.UpdateMenuItemRequest) (*emptypb.Empty, error) {
	log.Printf("UpdateMenuItemRequest %+v", req)

	// 1. Safety check: Verify the incoming nested ProfileInfo object is not nil
	menuItem := req.GetItem()
	if menuItem == nil {
		return nil, status.Error(codes.InvalidArgument, "menu item payload cannot be empty")
	}

	model := &domain.MenuModel{
		ID:                  menuItem.GetId(),
		Name:                menuItem.GetName(),
		PreviewDescription:  menuItem.GetPreviewDescription(),
		DetailedDescription: menuItem.GetDetailedDescription(),
		Time:                menuItem.GetTime(),
		Price:               menuItem.GetPrice(),
		Category:            menuItem.GetCategory(),
	}

	// 3. Pass the pointer to your domain model into your service layer
	if err := h.service.UpdateMenuItem(ctx, model); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update menu item: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (h *gRPCHandler) DeleteMenuItem(ctx context.Context, req *pb.DeleteMenuItemRequest) (*emptypb.Empty, error) {
	log.Printf("DeleteMenuItem %+v", req)

	// 1. Safety check: Verify the incoming nested ProfileInfo object is not nil
	menuId := req.GetId()
	if menuId == "" {
		return nil, status.Error(codes.InvalidArgument, "menu id cannot be empty")
	}

	model := &domain.MenuDelete{
		ID: menuId,
	}

	// 3. Pass the pointer to your domain model into your service layer
	if err := h.service.DeleteMenuItem(ctx, model); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete menu item: %v", err)
	}

	return &emptypb.Empty{}, nil
}
