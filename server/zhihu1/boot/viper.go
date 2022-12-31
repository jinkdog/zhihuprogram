package boot

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	g "main/app/global" //在反序列化时能引用global包里的的//g为别名
	"os"
)

const ( //配置管理器的建立
	configEnv = "ZHIHU_CONFIG_PATH" // 预定义环境变量
	//configFile = "C:\\Users\\张丰毅\\go\\src_workplace\\zhihu1\\manifest\\config\\config.example.yaml" // 预定义配置文件位置(有用)
	//configFile = "zhihu1\\appmanifest\\config\\config.example.yaml" // 预定义配置文件位置（没用）
	//configFile = "./manifest/config/config.example.yaml" // 预定义配置文件位置（没用）
	//configFile = "./app/manifest/config/config.example.yaml" // 预定义配置文件位置(有用)
	//configFile = "./root/gopro/config.example.yaml" // 预定义配置文件位置
	configFile = "/root/gopro/config.example.yaml" // 预定义配置文件位置
)

func ViperSetup(path ...string) { // 获取配置文件路径，传入文件路径
	var configPath string

	// 优先级: 参数 > 命令行 > 环境变量 > 默认值
	if len(path) != 0 {
		// 参数
		configPath = path[0]
	} else {
		// 命令行
		flag.StringVar(&configPath, "c", "", "set config path")
		flag.Parse()

		if configPath == "" {
			if configPath = os.Getenv(configEnv); configPath != "" {
				// 环境变量
			} else {
				// 默认值
				configPath = configFile
			}
		}
	}

	fmt.Printf("get config path: %s", configPath)

	v := viper.New()            //新建一个配置管理器
	v.SetConfigFile(configPath) // 设置配置文件路径，找到其位置
	v.SetConfigType("yaml")     // 设置配置文件类型
	err := v.ReadInConfig()     // 读取配置文件
	if err != nil {
		panic(fmt.Errorf("get config file failed, err: %v", err))
	}
	//用panic好处，执行配置文件不成功，后续肯定不成功

	if err = v.Unmarshal(&g.Config); err != nil {
		//反序列化
		// 将配置文件反序列化（映射）到 Config 结构体（global中定义）
		panic(fmt.Errorf("unmarshal config failed, err: %v", err))
	}
}
