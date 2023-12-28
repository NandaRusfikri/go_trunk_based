package pkg

import (
	"fmt"
	"github.com/joho/godotenv"
	"go_trunk_based/constant"
	"go_trunk_based/dto"
	"os"
	"strconv"
)

func LoadEnvironment(path string) (config dto.ConfigEnvironment) {
	err := godotenv.Load(path)
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Read environment variables from docker-compose.yml
	config.DbHost = os.Getenv("DB_HOST")
	config.DbPort = os.Getenv("DB_PORT")
	config.DbUser = os.Getenv("DB_USER")
	config.DbName = os.Getenv("DB_NAME")
	config.DbPass = os.Getenv("DB_PASS")
	config.DbSslmode = os.Getenv("DB_SSLMODE")

	config.SmtpHost = os.Getenv("SMTP_HOST")
	config.SmtpPort = os.Getenv("SMTP_PORT")
	config.SmtpName = os.Getenv("SMTP_NAME")
	config.SmtpEmail = os.Getenv("SMTP_EMAIL")
	config.SmtpPassword = os.Getenv("SMTP_PASSWORD")

	config.Timezone = os.Getenv("TIMEZONE")
	config.Version = os.Getenv("VERSION")
	config.RestPort = os.Getenv("REST_PORT")
	config.GoEnv = os.Getenv("GO_ENV")
	config.SwaggerHost = os.Getenv("SWAGGER_HOST")
	config.JwtSecret = os.Getenv("JWT_SECRET")

	return config
}

func LoadFeature(path string) error {

	if err := godotenv.Load(path); err != nil {
		return err
	}

	constant.FEATURE_MODULE_PRODUCT, _ = strconv.ParseBool(os.Getenv("FEATURE_MODULE_PRODUCT"))
	constant.FEATURE_API_USERS, _ = strconv.ParseBool(os.Getenv("FEATURE_API_USERS"))
	constant.FEATURE_USER_CREATE_VALIDATE_EMAIL, _ = strconv.ParseBool(os.Getenv("FEATURE_USER_CREATE_VALIDATE_EMAIL"))

	return nil
}
