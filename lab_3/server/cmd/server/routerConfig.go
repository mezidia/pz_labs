package main

var ROUTER_CONFIG = map[string]HttpHandlerFunc{
	"/channels": HttpHandlerFunc(ComposeForumsHandler()),
}
