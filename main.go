package main

func main() {
	// boolPtr := flag.Bool("prod", false, "Provide this flag in production. This ensures that a .config file is provided before the application starts.")
	// flag.Parse()
	// cfg := LoadConfig(*boolPtr)
	cfg := LoadEnvVarsConfig()
	dbCfg := cfg.Database
	a := App{}
	a.Initialize(dbCfg.Dialect(), dbCfg.ConnectionInfo())
	a.Run(cfg.Port)
}
