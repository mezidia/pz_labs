package main

var ROUTER_CONFIG = map[string]HttpHandlerFunc{
	"/users":  HttpHandlerFunc(ComposeUsersHandler()),
	"/forums": HttpHandlerFunc(ComposeForumsHandler()),
}
