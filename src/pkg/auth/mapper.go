package auth

func ModelToDomain(model *SessionModel) *Session {
	return &Session{
		SessID:       model.SessID,
		StudentID:    model.StudentID,
		Expiration:   model.Expiration,
		AccessToken:  model.AccessToken,
		RefreshToken: model.RefreshToken,
	}
}

func DomainToModel(domain *Session) *SessionModel {
	return &SessionModel{
		SessID:       domain.SessID,
		StudentID:    domain.StudentID,
		Expiration:   domain.Expiration,
		AccessToken:  domain.AccessToken,
		RefreshToken: domain.RefreshToken,
	}
}
