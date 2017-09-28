package util

import (
    "github.com/BurntSushi/toml"
)

type Config struct {
    
}

type StoreConfig struct {
    DbType  string
    DbName    string
    Permissions string
}

func LoadConfig() Config {}
