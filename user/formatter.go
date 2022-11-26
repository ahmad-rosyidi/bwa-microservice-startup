package user

type UserFormatter struct {
	ID         int    `json:"id"` // [1] comment aja, jika tidak ingin menampilkan id pada saat response json
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID, // [1] comment aja, jika tidak ingin menampilkan id pada saat response json
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return formatter
}
