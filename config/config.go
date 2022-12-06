package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port        string `envconfig:"PORT"`
	AutoMigrate bool   `envconfig:"AUTO_MIGRATE"`
	Jwt         struct {
		JwtSecretKey string `envconfig:"JWT_SECRET_KEY"`
		Expires      int    `envconfig:"JWT_EXPRIRES"`
	}

	Mysql struct {
		Host     string `envconfig:"DB_HOST"`
		Port     string `envconfig:"DB_PORT"`
		User     string `envconfig:"DB_USER"`
		Pass     string `envconfig:"DB_PASS"`
		Database string `envconfig:"DB_DATABASE"`
	}
	Cloud struct {
		Name      string `envconfig:"CLD_NAME"`
		ApiKey    string `envconfig:"CLD_API_KEY"`
		SecretKey string `envconfig:"CLD_SECRET_KEY"`
	}
	Mailer struct {
		Sender string `envconfig:"MAIL_SENDER"`
		Pass   string `envconfig:"MAIL_PASS"`
		Port   string `envconfig:"MAIL_PORT"`
		Host   string `envconfig:"MAIL_HOST"`
	}
	Link struct {
		LinkRenderict string `envconfig:"LINK_RENDERICT"`
		HostName      string `envconfig:"HOST_NAME"`
		HostFontName  string `envconfig:"HOST_FONT_NAME"`
	}
	Redis struct {
		Host     string `envconfig:"REDIS_HOST"`
		Port     string `envconfig:"REDIS_PORT"`
		Password string `envconfig:"REDIS_PASS"`
	}
}

var config *Config

func init() {
	var err error
	fmt.Println("App run on mode: ", os.Getenv("MODE"))

	if os.Getenv("MODE") == "production" {
		fmt.Println("load product env")
		err = godotenv.Load(".env.production")
	} else {
		fmt.Println("load dev env")
		err = godotenv.Load(".env")
	}

	if err != nil {
		log.Fatal("false get err", err)
	}

	config = &Config{}

	err = envconfig.Process("", config)

	if err != nil {
		panic(fmt.Sprintf("Failed to decode config env: %v", err))
	}
}

func GetConfig() *Config {
	return config
}
