package rethink

import (
	rdb "gopkg.in/gorethink/gorethink.v3"

	"github.com/dakstudios/trace-srv/db"
	proto "github.com/dakstudios/trace-srv/proto/trace"
)

type rethink struct {
	db *rbd.Session
}

var (
	Url = "localhost:28015"
)

func init() {
	db.Register(new(rethink))
}

func (r *rethink) Init() error {
	rdb.SetTags("gerethink", "json")

	var err error
	r.db, err = rdb.Connect(rdb.ConnectOpts{
		Address: Url,
	})

	if err != nil {
		return err
	}

	if _, err = rdb.DBCreate("dsadmin").RunWrite(r.db); err != nil {
		return err
	}

	if _, err = rdb.DB("dsadmin").TableCreate("tracing").RunWrite(r.db); err != nil {
		return err
	}

	if err = r.db.Close(); err != nil {
		return err
	}

	r.db, err = rdb.Connect(rdb.ConnectOpts{
		Address:  Url,
		Database: "dsadmin",
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *rethink) Create(span *proto.Span) error {
	if _, err := rdb.Table("tracing").Insert(span).RunWrite(r.db); err != nil {
		return err
	}

	return nil
}

func (r *rethink) Read(traceId string) ([]*proto.Span, error) {
	rows, err := rdb.Table("tracing").Filter(rdb.Row.Field("trace_id").Eq(traceId)).Run(r.db)
	if err != nil {
		return nil, err
	}

	spans := []*proto.Span{}
	if err = rows.All(&spans); err != nil {
		return nil, err
	}

	return spans, nil
}

func (r *rethink) Delete(traceId string) error {
	if _, err := rdb.Table("tracing").Filter(rdb.Row.Field("trace_id").Eq(traceId)).Delete().RunWrite(r.db); err != nil {
		return err
	}

	return nil
}

func (r *rethink) Search(name string, limit, offset int64, reverse bool) ([]*proto.Span, error) {
	query, err := rdb.Table("tracing").Filter(rdb.Row.Field("name").Eq(name)).Skip(offset).Limit(limit)
	if reverse {
		query = query.OrderBy(rdb.OrderByOpts{
			Index: rdb.Asc("timestamp"),
		})
	} else {
		query = query.OrderBy(rdb.OrderByOpts{
			Index: rdb.Desc("timestamp"),
		})
	}

	rows, err = query.Run(r.db)
	if err != nil {
		return nil, err
	}

	spans := []*proto.Span{}
	if err = rows.All(&spans); err != nil {
		return nil, err
	}

	return spans, nil
}

func (r *rethink) Live(name string, ch chan *proto.Span) error {

}
