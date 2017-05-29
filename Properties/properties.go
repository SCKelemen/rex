package Properties

import (
)



type FlagValue struct {
    Value string
    Filepath string
    isFile bool
}

type Context struct {
    K FlagValue
    IV FlagValue
    P FlagValue
    C FlagValue
    M FlagValue
}