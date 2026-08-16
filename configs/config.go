package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	Db DbConfig
}

type DbConfig struct {
	DSN string
}

func InitConfig()*Config{
	err := godotenv.Load()
	if err != nil{
		log.Println("error load .env file")
	}
	return  &Config{
		DbConfig{
			DSN: os.Getenv("DSN"),
		},	
	}
}