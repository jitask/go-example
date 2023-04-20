package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	h := &sayHandler{new(MySay)}

	typ := reflect.TypeOf(h)
	srv := reflect.ValueOf(h)

	req := &Request{
		Name: "json",
	}

	resp := &Response{}

	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mType := method.Type
		mName := method.Name

		num := mType.NumIn()
		contextType := mType.In(1)
		argType := mType.In(2)
		replyType := mType.In(3)

		ctx := reflect.Zero(contextType)
		argV := reflect.New(argType.Elem())
		replyV := reflect.New(replyType.Elem())
		setValue(argV.Interface())
		function := method.Func
		returnValues := function.Call([]reflect.Value{srv, ctx, reflect.ValueOf(argV.Interface()), reflect.ValueOf(replyV.Interface())})
		retValue := returnValues[0].Interface()
		if retValue == nil {
			if r, ok := replyV.Interface().(*Response); ok {
				fmt.Println(r.Msg)
			}
		}

		//另一种方式调用
		md := srv.MethodByName("Hello")
		returnValues = md.Call([]reflect.Value{reflect.ValueOf(context.Background()), reflect.ValueOf(req), reflect.ValueOf(resp)})
		retValue = returnValues[0].Interface()
		if retValue == nil {
			if r, ok := replyV.Interface().(*Response); ok {
				fmt.Println(r.Msg)
			}
		}

		fmt.Println(mName, num, contextType.Name(), argType.Elem().Name(), replyType.Elem().Name())
	}
}

func setValue(v interface{}) {
	r := &Request{
		Name: "json",
	}

	buf, err := json.Marshal(r)
	if err == nil {
		_ = json.Unmarshal(buf, v)
	}
}

type SayHandler interface {
	Hello(context.Context, *Request, *Response) error
}

type sayHandler struct {
	SayHandler
}

func (h *sayHandler) Hello(ctx context.Context, in *Request, out *Response) error {
	return h.SayHandler.Hello(ctx, in, out)
}

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

type MySay struct{}

func (s *MySay) Hello(ctx context.Context, req *Request, rsp *Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}
