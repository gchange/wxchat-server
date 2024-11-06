package httpserver

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

// Wrap adds several routes from package `net/http/pprof` to *gin.Engine object
func WrapPProf(rg gin.IRoutes) {
	wrapPProfGroup(rg)
}

// WrapPProfGroup adds several routes from package `net/http/pprof` to *gin.RouterGroup object
func wrapPProfGroup(rg gin.IRoutes) {
	rg.GET("/debug/pprof/", IndexHandler())
	rg.GET("/debug/pprof/heap", HeapHandler())
	rg.GET("/debug/pprof/goroutine", GoroutineHandler())
	rg.GET("/debug/pprof/allocs", AllocsHandler())
	rg.GET("/debug/pprof/block", BlockHandler())
	rg.GET("/debug/pprof/threadcreate", ThreadCreateHandler())
	rg.GET("/debug/pprof/cmdline", CmdlineHandler())
	rg.GET("/debug/pprof/profile", ProfileHandler())
	rg.GET("/debug/pprof/symbol", SymbolHandler())
	rg.POST("/debug/pprof/symbol", SymbolHandler())
	rg.GET("/debug/pprof/trace", TraceHandler())
	rg.GET("/debug/pprof/mutex", MutexHandler())
}

// IndexHandler will pass the call from /debug/pprof to pprof
func IndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Index(ctx.Writer, ctx.Request)
	}
}

// HeapHandler will pass the call from /debug/pprof/heap to pprof
func HeapHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("heap").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// GoroutineHandler will pass the call from /debug/pprof/goroutine to pprof
func GoroutineHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("goroutine").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// AllocsHandler will pass the call from /debug/pprof/allocs to pprof
func AllocsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("allocs").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// BlockHandler will pass the call from /debug/pprof/block to pprof
func BlockHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("block").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// ThreadCreateHandler will pass the call from /debug/pprof/threadcreate to pprof
func ThreadCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("threadcreate").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// CmdlineHandler will pass the call from /debug/pprof/cmdline to pprof
func CmdlineHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Cmdline(ctx.Writer, ctx.Request)
	}
}

// ProfileHandler will pass the call from /debug/pprof/profile to pprof
func ProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Profile(ctx.Writer, ctx.Request)
	}
}

// SymbolHandler will pass the call from /debug/pprof/symbol to pprof
func SymbolHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Symbol(ctx.Writer, ctx.Request)
	}
}

// TraceHandler will pass the call from /debug/pprof/trace to pprof
func TraceHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Trace(ctx.Writer, ctx.Request)
	}
}

// MutexHandler will pass the call from /debug/pprof/mutex to pprof
func MutexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("mutex").ServeHTTP(ctx.Writer, ctx.Request)
	}
}
