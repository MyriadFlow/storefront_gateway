package db

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	log "github.com/sirupsen/logrus"

	"github.com/TheLazarusNetwork/marketplace-engine/config/creatify"
	"github.com/TheLazarusNetwork/marketplace-engine/models"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

func InitDB() {

	var (
		host     = os.Getenv("DB_HOST")
		username = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
		port     = os.Getenv("DB_PORT")
	)

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s",
		host, username, password, dbname, port)
	var err error
	Db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	if err = Db.DB().Ping(); err != nil {
		log.Fatal("failed to ping database", err)
	}
	if err := Db.AutoMigrate(&models.FlowId{}, &models.User{}, &models.Role{}).Error; err != nil {
		log.Fatal(err)
	}
	//Create user_roles table
	Db.Exec(`create table if not exists user_roles (
			wallet_address text,
			role_id text,
			unique (wallet_address,role_id)
			)`)

	//Create flow id
	Db.Exec(`
	DO $$ BEGIN
		CREATE TYPE flow_id_type AS ENUM (
			'AUTH',
			'ROLE');
	EXCEPTION
    	WHEN duplicate_object THEN null;
	END $$;`)

	creatorRoleId, err := creatify.GetRole(creatify.CREATOR_ROLE)
	if err != nil {
		logwrapper.Fatal(err)
	}
	operatorRoleId, err := creatify.GetRole(creatify.OPERATOR_ROLE)
	if err != nil {
		logwrapper.Fatal(err)
	}

	// TODO: create role only if they does not exist
	rolesToBeAdded := []models.Role{
		{Name: "Creator Role", RoleId: hexutil.Encode(creatorRoleId[:]), Eula: "TODO Creator EULA"},
		{Name: "Operator Role", RoleId: hexutil.Encode(operatorRoleId[:]), Eula: "TODO Operator EULA"}}
	for _, role := range rolesToBeAdded {
		if err := Db.Model(&models.Role{}).FirstOrCreate(&role).Error; err != nil {
			log.Fatal(err)
		}
	}

}
