package ie2awsrdslib

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	ie2datatypes "github.com/insightengine2/ie2-utilities/types"
	ie2utilities "github.com/insightengine2/ie2-utilities/utils"
	"github.com/jackc/pgx/v5"
)

const ENV_RDS_HOST = "IE2_RDS_HOST"
const ENV_RDS_DBNAME = "IE2_RDS_DBNAME"
const ENV_RDS_PORT = "IE2_RDS_PORT"
const ENV_REGION = "AWS_REGION"
const ENV_USERNAME = "IE2_RDS_UNAME"
const ENV_SECRETKEY = "IE2_RDS_PWD_KEY"

func getRDSParams() (*ie2datatypes.RDSParams, error) {

	res := ie2datatypes.RDSParams{}

	// get username
	uname, err := ie2utilities.IE2GetEnv(ENV_USERNAME)

	if err != nil {
		return nil, err
	}

	res.DBUserName = uname

	// get host
	host, err := ie2utilities.IE2GetEnv(ENV_RDS_HOST)

	if err != nil {
		return nil, err
	}

	res.DBHost = host

	// get port
	port, err := ie2utilities.IE2GetEnv(ENV_RDS_PORT)

	if err != nil {
		return nil, err
	}

	res.DBPort = port

	// get region
	region, err := ie2utilities.IE2GetEnv(ENV_REGION)

	if err != nil {
		return nil, err
	}

	res.DBRegion = region

	// get database name
	dbname, err := ie2utilities.IE2GetEnv(ENV_RDS_DBNAME)

	if err != nil {
		return nil, err
	}

	res.DBName = dbname

	return &res, nil
}

func getRDSPWD() (string, error) {

	secretKey := os.Getenv(ENV_SECRETKEY)

	if len(secretKey) <= 0 {
		msg := fmt.Sprintf("missing environment variable: %s", ENV_SECRETKEY)
		log.Print(msg)
		return "", errors.New(msg)
	}

	config, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Printf(err.Error())
		return "", err
	}

	sm := secretsmanager.NewFromConfig(config)

	if sm == nil {
		msg := "Failed to create secretsmanager client"
		log.Print(msg)
		return "", errors.New(msg)
	}

	val, err := sm.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(ENV_SECRETKEY),
		VersionStage: aws.String("AWSCURRENT"),
	})

	if err != nil {
		log.Print(err)
		return "", err
	}

	return *val.SecretString, nil
}

func IE2RDSPostgresConnection() (*pgx.Conn, error) {

	rdsParams, err := getRDSParams()

	if err != nil {
		return nil, err
	}

	pwd, err := getRDSPWD()

	if err != nil {
		return nil, err
	}

	// connection string - assemble!
	connString := fmt.Sprintf("postgres://%s:%s/@%s:%s/%s", rdsParams.DBUserName, pwd, rdsParams.DBHost, rdsParams.DBPort, rdsParams.DBName)

	db, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
