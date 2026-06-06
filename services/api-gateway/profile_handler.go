package main

import (
	"encoding/json"
	"hair-studio-redmond/shared/contracts"
	"log"
	"net/http"
)

func (h *ApiGatewayHandlers) handleUpdateProfileInfo(w http.ResponseWriter, r *http.Request) {
	var reqBody updateProfileInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	profile, err := h.profileClient.Client.UpdateProfileInfo(r.Context(), reqBody.toProto())
	if err != nil {
		log.Printf("Failed to update profile: %v", err)
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: profile}

	writeJSON(w, http.StatusCreated, response)
}
