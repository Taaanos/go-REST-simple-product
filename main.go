package main

func main() {
	cfg := LoadConfig()
	dbCfg := cfg.Database
	a := App{}
	a.Initialize(dbCfg.Dialect(), dbCfg.ConnectionInfo())
	a.Run(cfg.Port)
}
