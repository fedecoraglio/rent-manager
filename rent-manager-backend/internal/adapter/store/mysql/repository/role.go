package repository

import (
	"context"
	"rent-manager-backend/internal/adapter/store/mysql"
	"rent-manager-backend/internal/core/domain"
)

type RoleRepository struct {
	db *mysql.DB
}

func NewRoleRepository(db *mysql.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (rr *RoleRepository) ListRoles(ctx context.Context, page uint64, limit uint64) ([]domain.Role, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := rr.db.QueryBuilder.
		Select(
			"rol_id",
			"rol_code",
			"rol_name",
			"rol_created_at",
		).
		From("roles").
		OrderBy("rol_id").
		Limit(limit).
		Offset(offset)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := rr.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []domain.Role

	for rows.Next() {
		var role domain.Role

		err := rows.Scan(
			&role.ID,
			&role.Code,
			&role.Name,
			&role.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}
