package context

import (
	"context"
	gcontext "github.com/gorilla/context"
	"net/http"
)

// @Author linjiabao
// @Date   2022/5/24

//MyContext context、gorilla使用方式
type MyContext struct {
	context.Context
	//something my own
	//todo
}

//下面是tomb 衍生使用context的方式
//import (
//"golang.org/x/net/context"
//tomb "gopkg.in/tomb.v2"
//)
//
//// NewContext returns a Context that is canceled either when parent is canceled
//// or when t is Killed.
//func NewContext(parent context.Context, t *tomb.Tomb) context.Context {
//	ctx, cancel := context.WithCancel(parent)
//	go func() {
//		select {
//		case <-t.Dying():
//			cancel()
//		case <-ctx.Done():
//		}
//	}()
//	return ctx
//}

// NewContext returns a Context whose Value method returns values associated
// with req using the Gorilla context package:
// http://www.gorillatoolkit.org/pkg/context
func NewContext(parent context.Context, req *http.Request) context.Context {
	return &wrapper{parent, req}
}

type wrapper struct {
	context.Context
	req *http.Request
}

type key int

const reqKey key = 0

// Value returns Gorilla's context package's value for this Context's request
// and key. It delegates to the parent Context if there is no such value.
func (ctx *wrapper) Value(key interface{}) interface{} {
	if key == reqKey {
		return ctx.req
	}
	if val, ok := gcontext.GetOk(ctx.req, key); ok {
		return val
	}
	return ctx.Context.Value(key)
}

// HTTPRequest returns the *http.Request associated with ctx using NewContext,
// if any.
func HTTPRequest(ctx context.Context) (*http.Request, bool) {
	// We cannot use ctx.(*wrapper).req to get the request because ctx may
	// be a Context derived from a *wrapper. Instead, we use Value to
	// access the request if it is anywhere up the Context tree.
	req, ok := ctx.Value(reqKey).(*http.Request)
	return req, ok
}
