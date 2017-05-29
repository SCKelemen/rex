package router

import (
     "github.com/urfave/cli"
     aes "github.com/SCKelemen/rex/AES"
)

type Router struct {
    context *cli.Context
    aesRouter AESRouter
}
type AESRouter struct {
    CTR aes.CTR
}
func NewRouter(c *cli.Context) *Router {
    return &Router{context: c}
} 
func (r *Router) NewAESRouter() *AESRouter {
    return &AESRouter{ CTR: aes.CTR{} }
}

