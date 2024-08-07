package user


type UserPayload struct {
    Username     string `json:"username"`
    Email        string `json:"email"`
    AuthProvider string `json:"auth_provider"`
    ProfileImage string `json:"profile_image"`
    IsValidated  bool   `json:"is_validated"`
}

// UpdateProfilePayload contient les données pour la mise à jour du profil utilisateur.
type UpdateProfilePayload struct {
    Username      string `json:"username"`
    ProfileImage  string `json:"profile_image"`
}
