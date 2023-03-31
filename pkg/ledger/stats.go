package ledger

import (
	"context"

	"github.com/formancehq/ledger/pkg/storage"
	"github.com/pkg/errors"
)

type Stats struct {
	Transactions uint64 `json:"transactions"`
	Accounts     uint64 `json:"accounts"`
}

func (l *Ledger) Stats(ctx context.Context) (Stats, error) {
	var stats Stats

	transactions, err := l.store.CountTransactions(ctx, storage.TransactionsQuery{})
	if err != nil {
		return stats, errors.Wrap(err, "counting transactions")
	}

	accounts, err := l.store.CountAccounts(ctx, storage.AccountsQuery{})
	if err != nil {
		return stats, errors.Wrap(err, "counting accounts")
	}

	return Stats{
		Transactions: transactions,
		Accounts:     accounts,
	}, nil
}
