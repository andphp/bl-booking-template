package boot

import (
	"fmt"
	"log"
	"os"

	"github.com/andphp/go-gin/goby/translator"
	"github.com/andphp/go-gin/goby/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
)

var (
	GLOBAL_VIPER               *viper.Viper
	GLOBAL_LOG                 *zap.Logger
	GLOBAL_Concurrency_Control = &singleflight.Group{}
)

func InitConfig() {
	GLOBAL_VIPER = utils.Viper() // 初始化Viper
	GLOBAL_LOG = utils.Zap()

	err := os.Setenv("GIN_MODE", "release")
	if err != nil {
		log.Fatal("Setenv GIN_MODE error ", err)
	}
	if err := translator.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
}
