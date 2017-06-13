package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"qiniupkg.com/x/log.v7"
)

type Config struct {
	MysqlHost      string `json:"mysql_host"`
	MysqlPort      string `json:"mysql_port"`
	MysqlUser      string `json:"mysql_user"`
	MysqlPassword  string `json:"mysql_password"`
	DatabaseName   string `json:"database_name"`
	AppConfigTable string `json:"app_config_table"`
}

type UserInfo struct {
	UserID   string
	UserName string
}

type AppConfig struct {
	AppKey             string `json:"app_key"`
	AppName            string `json:"app_name"`
	UserID             string `json:"user_id"`
	Platform           int64  `json:"platform"`
	HttpMonitorEnabled bool   `json:"http_monitor_enabled"`
	CrashReportEnabled bool   `json:"crash_report_enabled"`
	TelemetryEnabled   bool   `json:"telemetry_enabled"`
}

var (
	db     *sql.DB
	config *Config
)

func InitMysql(c *Config) {
	config = c
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.MysqlUser, config.MysqlPassword, config.MysqlHost, config.MysqlPort))
	if err != nil {
		log.Error(err.Error())
		return
	}
}

func CreateDadabase() error {
	_, err := db.Exec(`
	CREATE DATABASE presniff;
	`)
	return err
}

func CreateAppConfigTable() error {
	_, err := db.Exec(fmt.Sprintf("use %s", config.DatabaseName))
	if err != nil {
		log.Error(err)
	}
	_, err = db.Exec(`
	CREATE TABLE appconfig (
	app_id INT(64) UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
	app_key VARCHAR(40) NOT NULL,
	secret_key VARCHAR(40) NOT NULL,
	app_name VARCHAR(40) NOT NULL,
	user_id INT(64) NOT NULL,
	platform TINYINT NOT NULL,
	http_monitor_enabled BOOL NOT NULL,
	crash_report_enabled BOOL NOT NULL,
	telemetry_enabled BOOL NOT NULL
	);
	`)
	return err
}

func QueryAppConfig(appKey string) (*AppConfig, error) {
	appconfig := &AppConfig{}
	_, err := db.Exec(fmt.Sprintf("use %s", config.DatabaseName))
	if err != nil {
		log.Error(err)
	}
	err = db.QueryRow(fmt.Sprintf(`select * from %s where app_key="%s"`, config.AppConfigTable, appKey)).Scan(&appconfig.AppKey, &appconfig.UserID, &appconfig.Platform, &appconfig.HttpMonitorEnabled, &appconfig.CrashReportEnabled, &appconfig.TelemetryEnabled, &appconfig.AppName)
	return appconfig, err
}

func InsertAppConfig(appconfig *AppConfig) error {
	_, err := db.Exec(fmt.Sprintf("use %s", config.DatabaseName))
	if err != nil {
		log.Error(err)
	}
	_, err = db.Exec(fmt.Sprintf(`insert into %s values ("%s", "%s", %d, %t, %t, %t)`, config.AppConfigTable, appconfig.AppKey, appconfig.UserID, appconfig.Platform, appconfig.HttpMonitorEnabled, appconfig.CrashReportEnabled, appconfig.TelemetryEnabled))
	if err != nil {
		log.Error(err)
	}
	return err
}

func UpdateAppConfig(appconfig *AppConfig) error {
	_, err := db.Exec(fmt.Sprintf("use %s", config.DatabaseName))
	if err != nil {
		log.Error(err)
	}
	_, err = db.Exec(fmt.Sprintf(`update %s set user_id="%s", platform=%d, http_monitor_enabled=%t, crash_report_enabled=%t, telemetry_enabled=%t where app_key="%s"`, config.AppConfigTable, appconfig.UserID, appconfig.Platform, appconfig.HttpMonitorEnabled, appconfig.CrashReportEnabled, appconfig.TelemetryEnabled, appconfig.AppKey))
	if err != nil {
		log.Error(err)
	}
	return err
}
