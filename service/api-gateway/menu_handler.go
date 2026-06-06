package main

import (
	"encoding/json"
	"hair-studio-redmond/shared/contracts"
	"log"
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *ApiGatewayHandlers) handleGetMenuItems(w http.ResponseWriter, r *http.Request) {

	resp, err := h.menuClient.Client.GetMenuItems(r.Context(), &emptypb.Empty{})
	if err != nil {
		log.Printf("Failed to get menu items: %v", err)
		http.Error(w, "Failed to get menu items", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{
		Data: resp.Items,
	}

	writeJSON(w, http.StatusOK, response)
}
func (h *ApiGatewayHandlers) handleUpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	var reqBody updateMenuItemRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	profile, err := h.menuClient.Client.UpdateMenuItem(r.Context(), reqBody.toProto())
	if err != nil {
		log.Printf("Failed to update menu item: %v", err)
		http.Error(w, "Failed to update menu item", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: profile}

	writeJSON(w, http.StatusCreated, response)
}
func (h *ApiGatewayHandlers) handleCreateMenuItem(w http.ResponseWriter, r *http.Request) {
	var reqBody createMenuItemRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Printf("failed to parse JSON data: %v", err)
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	profile, err := h.menuClient.Client.CreateMenuItem(r.Context(), reqBody.toProto())
	if err != nil {
		log.Printf("Failed to create menu item: %v", err)
		http.Error(w, "Failed to create menu item", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: profile}

	writeJSON(w, http.StatusCreated, response)
}
func (h *ApiGatewayHandlers) handleDeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	var reqBody deleteMenuItemRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	profile, err := h.menuClient.Client.DeleteMenuItem(r.Context(), reqBody.toProto())
	if err != nil {
		log.Printf("Failed to delete menu item: %v", err)
		http.Error(w, "Failed to delete menu item", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: profile}
	writeJSON(w, http.StatusCreated, response)
}
