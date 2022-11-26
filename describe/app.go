package describe

type App struct {
	Config          Config
	WikimediaClient WikimediaClient
}

func MakeApp() App {
	var app App
	app.Config = ReadConfig()
	app.WikimediaClient = MakeCachedWikimediaClient(&app.Config, MakeHttpClient(&app.Config))
	return app
}
