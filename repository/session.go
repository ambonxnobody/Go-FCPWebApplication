package repository

import (
	"a21hc3NpZ25tZW50/model"
	"time"

	"gorm.io/gorm"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionsRepo struct {
	db *gorm.DB
}

func NewSessionsRepo(db *gorm.DB) *sessionsRepo {
	return &sessionsRepo{db}
}

func (u *sessionsRepo) AddSessions(session model.Session) error {
	res := u.db.Create(&session)
	if res.Error != nil {
		return res.Error
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) DeleteSession(token string) error {
	var session model.Session
	res := u.db.Where("token = ?", token).First(&session)
	if res.Error != nil {
		return res.Error
	}

	res = u.db.Delete(&session)
	if res.Error != nil {
		return res.Error
	}

	return nil // TODO: replace this
}

func (u *sessionsRepo) UpdateSessions(session model.Session) error {
	var update model.Session
	res := u.db.Where("email = ?", session.Email).First(&update)
	if res.Error != nil {
		return res.Error
	}

	update.Token = session.Token
	update.Email = session.Email
	update.Expiry = session.Expiry
	res = u.db.Save(&update)
	if res.Error != nil {
		return res.Error
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	var session model.Session
	res := u.db.Where("email = ?", email).First(&session)
	if res.Error != nil {
		return model.Session{}, res.Error
	}

	return session, nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	res := u.db.Where("token = ?", token).First(&session)
	if res.Error != nil {
		return session, res.Error
	}

	return session, nil // TODO: replace this
}

func (u *sessionsRepo) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}

func (u *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
