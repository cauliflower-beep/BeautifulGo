package main


type RawRpc struct {}

func(r *RawRpc) Hello(req string,rep *string) error{
	*rep = "hello" + req
	return nil
}