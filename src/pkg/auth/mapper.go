package auth

func ModelToDomain(model *SessionModel) *Session {
	// if access token parses error access token is nil
	accessToken, _ := DecodeAccessToken(model.AccessToken)
	// if refresh token parses error refresh token is nil
	refreshToken, _ := DecodeRefreshToken(model.RefreshToken)

	return &Session{
		SessID:       model.SessID,
		StudentID:    model.StudentID,
		Expiration:   model.Expiration,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func DomainToModel(domain *Session) *SessionModel {
	return &SessionModel{
		SessID:       domain.SessID,
		StudentID:    domain.StudentID,
		Expiration:   domain.Expiration,
		AccessToken:  domain.AccessToken.Token,
		RefreshToken: domain.RefreshToken.Token,
	}
}
