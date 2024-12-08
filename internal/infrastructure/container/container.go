package container

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/infrastructure/mysql"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const currentfilepath = "internal/infrastructure/container/container.go"

var v *viper.Viper

type (
	Container struct {
		Mysqldb  *gorm.DB
		Apps     *Apps
		AuthUsc  usecase.AuthsUseCase
		UserUsc  usecase.UserUseCase
		ProvinceCityUsc usecase.ProvinceCityUseCase
		TokoUsc  usecase.TokoUseCase
		CategoryUsc usecase.CategoryUseCase
		ProductUsc usecase.ProductUseCase
		TrxUsc usecase.TrxUseCase
	}

	Apps struct {
		Name      string `mapstructure:"name"`
		Host      string `mapstructure:"host"`
		Version   string `mapstructure:"version"`
		Address   string `mapstructure:"address"`
		HttpPort  int    `mapstructure:"httpport"`
		SecretJwt string `mapstructure:"secretJwt"`
	}
)

func loadEnv() {
	projectDirName := "go-example-cruid"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	v.SetConfigFile(string(rootPath) + `/.env`)
}

func init() {
	v = viper.New()

	v.AutomaticEnv()
	loadEnv()

	path, err := os.Executable()
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("os.Executable panic : %s", err.Error()))
	}

	dir := filepath.Dir(path)
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed read config : %s", err.Error()))
	}


	err = v.ReadInConfig()
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed init config : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed read configuration file")
}

func AppsInit(v *viper.Viper) (apps Apps) {
	err := v.Unmarshal(&apps)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprint("Error when unmarshal configuration file : ", err.Error()))
	}
	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed when unmarshal configuration file")
	return
}

func InitContainer() (cont *Container) {
	apps := AppsInit(v)
	utils.InitJWT(apps.SecretJwt)
	mysqldb := mysql.DatabaseInit(v)

	userRepo := repository.NewUsersRepository(mysqldb)
	provinceCityRepo := repository.NewProvinceCityRepository()
	tokoRepo := repository.NewTokoRepository(mysqldb)
	categoryRepo := repository.NewCategoryRepository(mysqldb)
	productRepo := repository.NewProductRepository(mysqldb)
	trxRepo := repository.NewTrxRepository(mysqldb)

	authUsc := usecase.NewAuthUseCase(userRepo, provinceCityRepo, tokoRepo)
	userUsc := usecase.NewUserUseCase(userRepo, provinceCityRepo)
	provinceCityUsc := usecase.NewProvinceCityUseCase(provinceCityRepo)
	tokoUsc := usecase.NewTokoUseCase(tokoRepo)
	categoryUsc := usecase.NewCategoryUseCase(categoryRepo)
	productUsc := usecase.NewProductUseCase(productRepo, tokoRepo, userRepo)
	trxUsc := usecase.NewTrxUseCase(trxRepo, userRepo, productRepo)

	return &Container{
		Apps:     &apps,
		Mysqldb:  mysqldb,
		AuthUsc:  authUsc,
		UserUsc:  userUsc,
		ProvinceCityUsc: provinceCityUsc,
		TokoUsc:  tokoUsc,
		CategoryUsc: categoryUsc,
		ProductUsc: productUsc,
		TrxUsc:  trxUsc,
	}

}
