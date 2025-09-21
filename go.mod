// Go Release Tour - Interactive Go language features learning platform
//
// This module provides a web-based interactive tutorial for learning
// new features introduced in Go versions 1.18 through 1.25.
// Each lesson includes official documentation references and
// executable code examples.
//
// Project: https://github.com/[username]/go-release-tour
// License: MIT
// Go compatibility: 1.24+
//
// Features covered:
// - Go 1.18: Generics, Type Parameters, Workspace Mode
// - Go 1.19: Memory Arenas, Atomic Types
// - Go 1.20: Comparable Types, Slice to Array Conversion, errors.Join
// - Go 1.21: Built-in Functions, slices Package, maps Package
// - Go 1.22: For-Range Integers, Loop Variables, math/rand/v2, Enhanced HTTP Routing
// - Go 1.23: Structured Logging, Iterators, Timer Reset, slices.Concat, cmp.Or, maps.Collect
// - Go 1.24: Generic Type Aliases, Swiss Tables Maps, crypto/mlkem, Testing Loop, os.Root, Weak Pointers
// - Go 1.25: Container-aware GOMAXPROCS, Trace Flight Recorder, testing/synctest, go.mod ignore, JSON v2, go doc -http
module go-release-tour

go 1.24

// Development dependencies
// Air for hot reload development: go run github.com/air-verse/air@latest
// Docker Compose for containerized development environment
