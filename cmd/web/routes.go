package main

import (
	"github.com/priyam314/Bookings/pkg/config"
	"github.com/priyam314/Bookings/pkg/handler"

	//"github.com/bmizerany/pat"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// using third pary library called "pat" for mux
func routes(app *config.AppConfig) http.Handler {
	/*
		// Pat
		mux := pat.New()
		// takes two parameters
		// 1. pattern
		// 2. handler func
		mux.Get("/", http.HandlerFunc(handler.Repo.Home))
		mux.Get("/about", http.HandlerFunc(handler.Repo.About))
	*/

	// Chi
	// MiddleWare
	// 		middleware is a loosely defined term for any software or service
	// 		that enables the *part* of a system to communicate and manage data.
	//		it is the software that handles communication between components
	//		and i/o, so dev can focus on specific purpose of their application.
	//
	// 		middleware helps to proccess some request that comes in your web
	// 		application and forms some action on it
	mux := chi.NewRouter()

	// gracefully absorbs panics and prints the stackstrace
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)

	// this function provides the functionality to serve the entire file
	// system directory with indexes, takes input of "FileSystem interface"
	// so gave that directory. As an output it returns Handler.
	// content will be served from that directory, which here we say root dir
	fileserver := http.FileServer(http.Dir("./static/"))

	// Handle registers a new route with a matcher for the URL
	// http.StripPrefix forwards the handling of the request to one you specify
	// as its parameter, but before that it modifies the requested URL by
	// striping off the specified prefix
	// for e.g. if browser request he resource "/static/example.txt"
	// StripPrefix will cut "/static/" and forward th modified request to the
	// handler returned by http.FileServer(), so request resource is `/example.txt`
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))

	return mux
}
