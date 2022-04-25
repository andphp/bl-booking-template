package boot

import (
	"fmt"
	"log"
	"os"

	"github.com/andphp/go-gin/goby/translator"
	"github.com/andphp/go-gin/goby/utils"
	"golang.org/x/sync/singleflight"
)

var (
	GLOBAL_Concurrency_Control = &singleflight.Group{}
)

func InitConfig() {
	utils.Viper() // 初始化Viper
	utils.Zap()

	err := os.Setenv("GIN_MODE", "release")
	if err != nil {
		log.Fatal("Setenv GIN_MODE error ", err)
	}
	if err := translator.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
}
