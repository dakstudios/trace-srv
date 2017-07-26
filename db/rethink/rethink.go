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
	rdb.Table("tracing").GetAll(
}

func (r *rethink) Delete(traceId string) error {
	return nil
}

func (r *rethink) Search(name string, limit, offset int64, reverse bool) ([]*proto.Span, error) {
	return nil, nil
}

func (r *rethink) Live(name string, ch chan *proto.Span) error {
	return nil
}
