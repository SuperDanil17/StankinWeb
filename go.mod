// Require is a list of external dependencies required by this module.
// Each dependency is specified as a module path and version number.
// The go command uses this information to download the specified
// versions of the dependencies and build the module.
require (
    // Gorilla Mux is a powerful URL router and dispatcher
    // for matching incoming requests to their respective
    // handler. It is used in this module to handle HTTP
    // requests and dispatch them to the appropriate handler.
    github.com/gorilla/mux v1.8.1 // indirect

    // Gorilla SecureCookie is a package for securely signing
    // and encrypting cookies. It is used in this module to
    // sign and encrypt session cookies.
    github.com/gorilla/securecookie v1.1.2 // indirect

    // Gorilla Sessions is a package for managing user sessions.
    // It is used in this module to manage user sessions and
    // store session data.
    github.com/gorilla/sessions v1.2.2 // indirect

    // Go SQLite3 is a Go driver for SQLite3 databases. It is used
    // in this module to interact with SQLite3 databases.
    github.com/mattn/go-sqlite3 v1.14.18 // indirect

    // Go Crypto is a collection of cryptography algorithms
    // implemented in Go. It is used in this module for various
    // cryptography-related tasks.
    golang.org/x/crypto v0.15.0 // indirect
)
