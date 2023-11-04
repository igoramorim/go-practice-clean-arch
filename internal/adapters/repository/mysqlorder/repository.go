package mysqlorder

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
)

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// TODO: Unit tests.

type Repository struct {
	db *sql.DB
}

var _ dorder.Repository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, order *dorder.Order) error {
	const query = `insert into orders(id, price, tax, final_price, created_at) values (?,?,?,?,?)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	persistable := PersistableFromDomain(order)
	_, err = stmt.ExecContext(ctx, persistable.ID, persistable.Price, persistable.Tax, persistable.FinalPrice, persistable.CreatedAt)
	if err != nil {
		var mySQLErr *mysql.MySQLError
		if errors.As(err, &mySQLErr) && mySQLErr.Number == 1062 {
			return dorder.ErrOrderAlreadyExists
		}
		return err
	}

	return nil
}

func (r *Repository) FindAllByPage(ctx context.Context, page, limit int, sort string) ([]*dorder.Order, int64, error) {
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = defaultSort
	}

	if page == 0 {
		page = defaultPage
	}

	if limit == 0 {
		limit = defaultLimit
	}

	total, err := r.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	var initialSeq = (page - 1) * limit
	var finalSeq = page * limit
	const query = `
		select o.id, o.price, o.tax, o.final_price, o.created_at
		from orders o
		where seq > %d and seq <= %d
		order by o.seq %s;`
	stmt, err := r.db.Prepare(fmt.Sprintf(query, initialSeq, finalSeq, sort))
	if err != nil {
		return nil, 0, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	orders := make([]*dorder.Order, 0, limit)
	for rows.Next() {
		var persistable PersistableOrder
		if err = rows.Scan(&persistable.ID, &persistable.Price, &persistable.Tax, &persistable.FinalPrice,
			&persistable.CreatedAt); err != nil {
			return nil, 0, err
		}
		orders = append(orders, PersistableToDomain(persistable))
	}

	return orders, total, nil
}

func (r *Repository) Count(ctx context.Context) (int64, error) {
	stmt, err := r.db.Prepare("select max(seq) from orders")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx)
	if err != nil {
		return 0, err
	}
	if err = row.Err(); err != nil {
		return 0, err
	}

	var count int64
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

const (
	defaultSort  = "asc"
	defaultPage  = 1
	defaultLimit = 10
)
