package postgres

import (
	"database/sql"
	"fmt"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/external"

	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/contract-tracker/resources"
)

const (
	tasksTable = "tasks"

	tasksId               = "id"
	tasksAccount          = "account"
	tasksSignature        = "signature"
	tasksFileIpfsHash     = "file_ipfs_hash"
	tasksMetadataIpfsHash = "metadata_ipfs_hash"
	tasksTokenId          = "token_id"
	tasksStatus           = "status"
)

type tasksQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewTasksQ(database *pgdb.DB) external.TasksQ {
	return &tasksQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", tasksTable)).From(tasksTable),
	}
}

func (q *tasksQ) New() external.TasksQ {
	return NewTasksQ(q.database.Clone())
}

func (q *tasksQ) Page(page pgdb.OffsetPageParams) external.TasksQ {
	q.selector = page.ApplyTo(q.selector, "id")
	return q
}

func (q *tasksQ) Sort(sort pgdb.Sorts) external.TasksQ {
	q.selector = sort.ApplyTo(q.selector, map[string]string{
		"id":         "id",
		"created_at": "created_at",
	})

	return q
}

func (q *tasksQ) Select(selector external.TaskSelector) (tasks []external.Task, err error) {
	return q.selectByQuery(applyTasksSelector(q.selector, selector))
}

func (q *tasksQ) GetById(id int64) (*external.Task, error) {
	var task external.Task

	err := q.database.Get(&task, q.selector.Where(squirrel.Eq{tasksId: id}))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &task, err
}

func (q *tasksQ) GetByHash(ipfsHash string) (*external.Task, error) {
	var task external.Task

	err := q.database.Get(&task, q.selector.Where(squirrel.Eq{tasksMetadataIpfsHash: ipfsHash}))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &task, err
}

func (q *tasksQ) Insert(task external.Task) (id int64, err error) {
	statement := squirrel.Insert(tasksTable).
		Suffix("returning id").
		SetMap(structs.Map(&task))

	err = q.database.Get(&id, statement)
	return id, err
}

func (q *tasksQ) Delete(id int64) error {
	statement := squirrel.Delete(tasksTable).Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
}

func (q *tasksQ) Update(task external.Task, id int64) error {
	statement := squirrel.Update(tasksTable).
		SetMap(structs.Map(&task)).
		Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
}

func (q *tasksQ) UpdateFileIpfsHash(newIpfsHash string, id int64) error {
	return q.updateHash(tasksFileIpfsHash, newIpfsHash, id)
}

func (q *tasksQ) UpdateMetadataIpfsHash(newIpfsHash string, id int64) error {
	return q.updateHash(tasksMetadataIpfsHash, newIpfsHash, id)
}

func (q *tasksQ) updateHash(fieldName, newIpfsHash string, id int64) error {
	statement := squirrel.Update(tasksTable).
		Set(fieldName, newIpfsHash).
		Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
}

func (q *tasksQ) UpdateTokenId(newTokenId, id int64) error {
	statement := squirrel.Update(tasksTable).
		Set(tasksTokenId, newTokenId).
		Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
}

func (q *tasksQ) UpdateStatus(newStatus resources.TaskStatus, id int64) error {
	statement := squirrel.Update(tasksTable).
		Set(tasksStatus, newStatus).
		Where(squirrel.Eq{tasksId: id})
	return q.database.Exec(statement)
}

func (q *tasksQ) Transaction(fn func(q external.TasksQ) error) (err error) {
	return q.database.Transaction(func() error {
		return fn(q)
	})
}

func (q *tasksQ) selectByQuery(query squirrel.Sqlizer) (subtasks []external.Task, err error) {
	err = q.database.Select(&subtasks, query)
	return subtasks, err
}

func applyTasksSelector(sql squirrel.SelectBuilder, selector external.TaskSelector) squirrel.SelectBuilder {
	sql = selector.PageParams.ApplyTo(sql, tasksId)

	if selector.Account != nil {
		sql = sql.Where(squirrel.Eq{tasksAccount: *selector.Account})
	}
	if selector.IpfsHash != nil {
		sql = sql.Where(squirrel.Eq{tasksMetadataIpfsHash: *selector.IpfsHash})
	}
	if selector.Status != nil {
		sql = sql.Where(squirrel.Eq{tasksStatus: *selector.Status})
	}

	return sql
}
