package slug

import (
	"context"
	"fmt"
	"strconv"

	"github.com/arieffian/simple-commerces-monorepo/internal/database"
	slug_pkg "github.com/gosimple/slug"
)

type SlugGeneratorService interface {
	Slug(ctx context.Context, str string) string
	GenerateUniqueSlug(ctx context.Context, str string, table string, field string) (string, error)
}

var _ SlugGeneratorService = (*slugGenerator)(nil)

type slugGenerator struct {
	db *database.DbInstance
}

type NewSlugGeneratorParams struct {
	Db *database.DbInstance
}

func (s *slugGenerator) Slug(ctx context.Context, str string) string {
	return slug_pkg.Make(str)
}

func (s *slugGenerator) GenerateUniqueSlug(ctx context.Context, str string, table string, field string) (string, error) {

	slug := slug_pkg.Make(str)

	if table == "" {
		return slug, nil
	}

	slugField := "slug"
	if field != "" {
		slugField = field
	}

	stmt := fmt.Sprintf(`SELECT count(%s) FROM %s WHERE %s = $1`, slugField, table, slugField)

	var count int
	err := s.db.Db.Raw(stmt, slug).Scan(&count).Error

	if err != nil {
		return "", err
	}

	if count == 0 {
		return slug, nil
	}

	// // @note: this is for removing slug with number
	// // exp := regexp.MustCompile(`-[0-9]+$`)
	// // slug = exp.ReplaceAllString(slug, "")

	pattern := `^(` + slug + `)(-[0-9]*)?$`

	stmt = fmt.Sprintf(`SELECT count(%s) FROM %s WHERE %s ~ $1`, slugField, table, slugField)

	err = s.db.Db.Raw(stmt, pattern).Scan(&count).Error

	if err != nil {
		return "", err
	}

	if count == 0 {
		return slug, nil
	}

	slug = slug + "-" + strconv.Itoa(count)

	return slug, nil
}

func NewSlugGeneratorService(p NewSlugGeneratorParams) *slugGenerator {
	return &slugGenerator{
		db: p.Db,
	}
}
