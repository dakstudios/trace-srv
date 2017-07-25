package rethink

import (
	rdb "gopkg.in/gorethink/gorethink.v3"

	"github.com/dakstudios/trace-srv/db"
	proto "github.com/dakstudios/trace-srv/proto/trace"
)

type rethink struct {
	//db *mgo.Database
}

func init() {
	db.Register(new(rethink))
}

func (r *rethink) Init() error {
	return nil
}

func (r *rethink) Create(span *proto.Span) error {
	return nil
}

func (r *rethink) Read(traceId string) ([]*proto.Span, error) {
	return nil, nil
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
