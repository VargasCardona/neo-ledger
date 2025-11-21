package guitars

import (
	"context"
  "fmt"
	"strings"
	"time"
	"github.com/jackc/pgx/v5"
	repo "github.com/vargascardona/neo-ledger/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListGuitars(ctx context.Context) ([]repo.Guitar, error)
	FindGuitarByID(ctx context.Context, id int) (repo.Guitar, error)
	CreateGuitar(ctx context.Context, tempGuitar Guitar) (repo.Guitar, error)
	UpdateGuitar(ctx context.Context, tempGuitar Guitar) (repo.Guitar, error)
	DeleteGuitar(ctx context.Context, id int) error
  DisplayName(ctx context.Context, g repo.Guitar) string
  IsVintage(ctx context.Context, g repo.Guitar) bool
}

type svc struct {
	repo *repo.Queries
	db *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db: db,
	}
}

func (s *svc) ListGuitars(ctx context.Context) ([]repo.Guitar, error) {
	return s.repo.ListGuitars(ctx)
}

func (s *svc) FindGuitarByID(ctx context.Context, id int) (repo.Guitar, error) {
	return s.repo.FindGuitarByID(ctx, int64(id))
}

func (s *svc) CreateGuitar(ctx context.Context, tempGuitar Guitar) (repo.Guitar, error) {
  tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Guitar{}, err
	}
  defer  tx.Rollback(ctx)
	qtx := s.repo.WithTx(tx)

  params := repo.CreateGuitarParams{
    Brand:           tempGuitar.Brand,
    Model:           tempGuitar.Model,
    Year:            intToInt16Ptr(tempGuitar.Year),              // *int -> *int16
    Notes:           tempGuitar.Notes,
    }

  guitar, err := qtx.CreateGuitar(ctx, params)

  if err != nil {
    return repo.Guitar{}, err
  }

	tx.Commit(ctx)
  return guitar, nil
}

func (s *svc) UpdateGuitar(ctx context.Context, tempGuitar Guitar) (repo.Guitar, error) {
  tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Guitar{}, err
	}
  defer  tx.Rollback(ctx)
	qtx := s.repo.WithTx(tx)

  params := repo.UpdateGuitarParams{
    ID:           tempGuitar.ID,
    Brand:           tempGuitar.Brand,
    Model:           tempGuitar.Model,
    Year:            intToInt16Ptr(tempGuitar.Year),              // *int -> *int16
    Notes:           tempGuitar.Notes,
    }

  guitar, err := qtx.UpdateGuitar(ctx, params)

  if err != nil {
    return repo.Guitar{}, err
  }

	tx.Commit(ctx)
  return guitar, nil
}

func (s *svc) DeleteGuitar(ctx context.Context, id int) (error) {
	return s.repo.DeleteGuitar(ctx, int64(id))
}

func (s *svc) DisplayName(ctx context.Context, g repo.Guitar) string {
	parts := []string{}

	year := 0
	if g.Year != nil {
		year = int(*g.Year)
	}
	if year >= 1900 && year <= time.Now().Year() {
		parts = append(parts, fmt.Sprintf("%d", year))
	}

	brand := strings.TrimSpace(g.Brand)
	model := strings.TrimSpace(g.Model)
	parts = append(parts, brand, model)

	if year > 0 && time.Now().Year()-year >= 40 {
		parts = append(parts, "(Vintage)")
	}

	if len(parts) == 2 {
		return brand + " " + model
	}
	return strings.Join(parts, " ")
}

func (s *svc) IsVintage(_ context.Context, g repo.Guitar) bool {
	year := int(*g.Year)
	currentYear := time.Now().Year()
	return year > 0 && currentYear-year >= 40
}

func intToInt16Ptr(i *int) *int16 {
    if i == nil {
        return nil
    }
    v := int16(*i)
    return &v
}
