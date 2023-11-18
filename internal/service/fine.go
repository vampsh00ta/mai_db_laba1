package service

import rep "TgDbMai/internal/repository"

type Fine interface {
	IssueFine(passport int, amount int, reason string) (*rep.Person, error)
	GetFines(passport int) (*rep.Person, error)
}

func (s service) IssueFine(passport int, amount int, reason string) (*rep.Person, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	person, err := s.rep.IssueFine(tx, passport, amount, reason)
	if err != nil || tx.Error != nil {
		return nil, err
	}

	return person, nil
}
func (s service) GetFines(passport int) (*rep.Person, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	person, err := s.rep.GetFines(tx, passport)
	if err != nil || tx.Error != nil {
		return nil, err
	}

	return person, nil
}
