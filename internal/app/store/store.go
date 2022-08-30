package store

import "github.com/captkakao/http-rest-api/internal/app/store/sqlstore"

type Store interface {
	User() sqlstore.UserRepository
}
