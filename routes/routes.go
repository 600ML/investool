// 在这个文件中注册 URL handler

package routes

import "github.com/gin-gonic/gin"

// Routes 注册 API URL 路由
func Routes(app *gin.Engine) {
	app.GET("/invest/", StockIndex)
	app.POST("/invest/selector", StockSelector)
	app.POST("/invest/checker", StockChecker)
	app.GET("/invest/fund", FundIndex)
	app.GET("/invest/fund/filter", FundFilter)
	app.POST("/invest/fund/check", FundCheck)
	app.GET("/invest/about", About)
	app.GET("/invest/comment", Comment)
	app.GET("/invest/fund/similarity", FundSimilarity)
	app.GET("/invest/materials", Materials)
	app.POST("/invest/fund/query_by_stock", QueryFundByStock)
	app.GET("/invest/fund/managers", FundManagers)
}
