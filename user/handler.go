package user

import (
    "encoding/json"
    "net/http"
    "time"
)

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request, service UserService) {
    var payload UpdateProfilePayload
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    userID, ok := r.Context().Value("userID").(string)
    if !ok {
        http.Error(w, "User ID not found in context", http.StatusUnauthorized)
        return
    }

    user, err := service.GetUserByEmail(userID)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    user.Username = payload.Username
    user.ProfileImage = payload.ProfileImage
    user.UpdatedAt = time.Now()

    err = service.UpdateUser(user)
    if err != nil {
        http.Error(w, "Failed to update profile", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})
}
