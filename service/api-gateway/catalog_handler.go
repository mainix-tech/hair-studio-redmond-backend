package main

import (
	"encoding/json"
	"hair-studio-redmond/shared/contracts"
	"log"
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *ApiGatewayHandlers) handleGetCatalogItems(w http.ResponseWriter, r *http.Request) {

	resp, err := h.catalogClient.Client.GetCatalogItems(r.Context(), &emptypb.Empty{})
	if err != nil {
		log.Printf("Failed to get catalog items: %v", err)
		http.Error(w, "Failed to get catalog items", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{
		Data: resp.Items,
	}

	writeJSON(w, http.StatusOK, response)
}
func (h *ApiGatewayHandlers) handleUpdateCatalogItem(w http.ResponseWriter, r *http.Request) {
	var reqBody updateCatalogItemRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	catalog, err := h.catalogClient.Client.UpdateCatalogItem(r.Context(), reqBody.toProto())
	if err != nil {
		log.Printf("Failed to update catalog item: %v", err)
		http.Error(w, "Failed to update catalog item", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: catalog}

	writeJSON(w, http.StatusCreated, response)
}
func (h *ApiGatewayHandlers) handleCreateCatalogItem(w http.ResponseWriter, r *http.Request) {
	var reqBody createCatalogItemRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Printf("failed to parse JSON data: %v", err)
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	catalog, err := h.catalogClient.Client.CreateCatalogItem(r.Context(), reqBody.toProto())
	if err != nil {
		log.Printf("Failed to create catalog item: %v", err)
		http.Error(w, "Failed to create catalog item", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: catalog}

	writeJSON(w, http.StatusCreated, response)
}
func (h *ApiGatewayHandlers) handleDeleteCatalogItem(w http.ResponseWriter, r *http.Request) {
	var reqBody deleteCatalogItemRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	catalog, err := h.catalogClient.Client.DeleteCatalogItem(r.Context(), reqBody.toProto())
	if err != nil {
		log.Printf("Failed to delete catalog item: %v", err)
		http.Error(w, "Failed to delete catalog item", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: catalog}
	writeJSON(w, http.StatusCreated, response)
}
