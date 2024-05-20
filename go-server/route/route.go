package route

import (
	"Go_simpleWMS/handler/auth"
	"Go_simpleWMS/handler/test"
	"Go_simpleWMS/handler/upload"
	"Go_simpleWMS/route/group"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/semaphore"
)

func Route(ginServer *gin.Engine, sem *semaphore.Weighted) {
	// 设置静态资源路径
	ginServer.Static("/static", "./static")

	ginApi := ginServer.Group("/api", utils.SemaphoreMiddleware(sem))

	ginApi.GET("/ping", utils.SemaphoreMiddleware(sem), func(c *gin.Context) {
		test.Ping(c)
	})
	//鉴权接口
	ginApi.GET("/auth", func(c *gin.Context) {
		auth.Auth(c)
	})
	ginApi.POST("/upload/goods_img", func(c *gin.Context) {
		upload.GoodsImageUpload(c)
	})
	// 分组路由
	group.UserGroup(ginApi)          // 用户路由
	group.WarehouseGroup(ginApi)     // 仓库路由
	group.GoodsType(ginApi)          // 货品类型路由
	group.DepartmentGroup(ginApi)    // 部门路由
	group.InventoryTypeGroup(ginApi) // 库存类型路由
	group.StaffGroup(ginApi)         // 员工路由
	group.GoodsGroup(ginApi)         // 货品路由
	group.InventoryGroup(ginApi)     // 出入库路由
	group.UnitGroup(ginApi)          // 单位路由
}
