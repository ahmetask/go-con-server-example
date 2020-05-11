package main

import gocon "github.com/ahmetask/go-con"

func main() {

	// Initialize DB
	var configs []StaticDBConfig

	type CustomConfig struct {
		A string
		B bool
	}

	configs = append(configs, StaticDBConfig{
		AppName:   "go-con-client",
		NameSpace: "default", //k8s namespace
		Port:      "8081",    //go-con-client config port
		Content: CustomConfig{
			A: "A",
			B: true,
		},
	})

	type CustomConfigInner struct {
		A string
	}

	type CustomConfig2 struct {
		C int
		D bool
		E CustomConfigInner
	}

	var ExampleConfig2 = CustomConfig2{
		C: 1,
		D: true,
		E: CustomConfigInner{
			A: "A",
		},
	}

	configs = append(configs, StaticDBConfig{
		AppName:   "go-con-client-2",
		NameSpace: "default", //k8s namespace
		Port:      "8082", //go-con-client-2 config port
		Content:   ExampleConfig2,
	})

	staticDB := NewStaticDB(configs)

	configManager := gocon.ConfigManager{
		DbInstance:       staticDB,
		ConfigServerPort: "8080", // Server port
		BroadcastModuleConfig: gocon.BroadcastModuleConfig{
			Available: true, // Broadcast available
			IsLocal:   false, //Run in k8s env
		},
	}

	err := configManager.Initialize()
	panic(err)
}
