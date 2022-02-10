package main

/*
go原生的rpc服务默认使用gob编码，其他语言调用go写的rpc服务将比较困难。
要方便实现跨语言rpc调用，例如python客户端调用go写的rpc服务，
就需要采用其他编码方式，如json或者pb等
 */
