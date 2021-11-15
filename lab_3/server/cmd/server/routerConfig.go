package main

var ROUTER_CONFIG = map[string]HttpHandlerFunc{
	"/forums": HttpHandlerFunc(ComposeForumsHandler()),
}
