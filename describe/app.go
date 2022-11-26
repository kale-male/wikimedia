package describe

type App struct {
	Config          Config
	WikimediaClient WikimediaClient
}

func MakeApp() App {
	var app App
	app.Config = ReadConfig()
	app.WikimediaClient = MakeCachedWikimediaClient(MakeHttpClient(&app.Config))
	return app
}
