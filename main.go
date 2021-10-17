package main

import (
	"time"

	"github.com/devstackq/ozon/cache"
	"github.com/devstackq/ozon/controller"
	router "github.com/devstackq/ozon/http"
	"github.com/devstackq/ozon/repository"
	"github.com/devstackq/ozon/service"
)

//global : type(Interface)  = package.Func - return new struct.Interface(methods)
var (
	urlRepository repository.UrlRepository = repository.NewPSqlRepository()
	urlService    service.UrlService       = service.NewUrlService(urlRepository)
	urlCache      cache.UrlCache           = cache.NewRedisCache("localhost:6379", 1, 24*time.Hour)
	urlController controller.UrlController = controller.NewUrlController(urlService, urlCache)
	httpRouter    router.Router            = router.NewMuxRouter()
)

func main() {
	httpRouter.GET("/", urlController.GetLinkByShortLink)
	httpRouter.POST("/", urlController.GenerateNewLink)
	httpRouter.Serve("8000")
}
