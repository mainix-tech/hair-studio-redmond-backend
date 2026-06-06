package main

import (
	"encoding/json"
	"hair-studio-redmond/shared/contracts"
	profilePb "hair-studio-redmond/shared/proto/profile"
	"log"
	"net/http"
)

func (h *ApiGatewayHandlers) handleUpdateProfileInfo(w http.ResponseWriter, r *http.Request) {
	var reqBody UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	profile, err := h.profileClient.Client.UpdateProfileInfo(r.Context(), reqBody.ToProto())
	if err != nil {
		log.Printf("Failed to update profile: %v", err)
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: profile}

	writeJSON(w, http.StatusCreated, response)
}

func (h *ApiGatewayHandlers) handleGetProfileInfo(w http.ResponseWriter, r *http.Request) {

	resp, err := h.profileClient.Client.GetProfileInfo(r.Context(), &profilePb.GetProfileInfoRequest{})
	if err != nil {
		log.Printf("Failed to get menu items: %v", err)
		http.Error(w, "Failed to get menu items", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{
		Data: resp,
	}

	writeJSON(w, http.StatusOK, response)
}
