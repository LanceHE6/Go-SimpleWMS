package route

import (
	"Go_simpleWMS/handler/auth"
	"Go_simpleWMS/handler/department"
	"Go_simpleWMS/handler/goods"
	"Go_simpleWMS/handler/goodsType"
	"Go_simpleWMS/handler/inventoryType"
	"Go_simpleWMS/handler/staff"
	"Go_simpleWMS/handler/test"
	"Go_simpleWMS/handler/unit"
	"Go_simpleWMS/handler/upload"
	"Go_simpleWMS/handler/user"
	"Go_simpleWMS/handler/warehouse"
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/semaphore"
)

func Route(ginServer *gin.Engine, sem *semaphore.Weighted) {
	// 设置静态资源路径
	ginServer.Static("/res", "./static/res")

	ginApi := ginServer.Group("/api", utils.SemaphoreMiddleware(sem))

	ginApi.GET("/ping", utils.SemaphoreMiddleware(sem), func(c *gin.Context) {
		test.Ping(c)
	})
	//鉴权接口
	ginApi.GET("/auth", func(c *gin.Context) {
		auth.Auth(c)
	})
	ginApi.POST("/upload", func(c *gin.Context) {
		upload.UploadFile(c)
	})
	// 路由分组
	userGroup := ginApi.Group("/user")

	userGroup.POST("/register", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		user.Register(c)
	})
	userGroup.POST("/upload", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		user.UploadUsers(c)
	})
	userGroup.POST("/login", func(c *gin.Context) {
		user.Login(c)
	})
	userGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		user.DeleteUser(c)
	})
	userGroup.PUT("/update", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		user.UpdateUser(c)
	})
	userGroup.GET("/list", utils.AuthMiddleware(), func(c *gin.Context) {
		user.ListUsers(c)
	})

	warehouseGroup := ginApi.Group("/warehouse", utils.AuthMiddleware())
	warehouseGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		warehouse.AddWarehouse(c)
	})
	warehouseGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		warehouse.DeleteWarehouse(c)
	})
	warehouseGroup.PUT("/update", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		warehouse.UpdateWarehouse(c)
	})
	warehouseGroup.GET("/list", func(c *gin.Context) {
		warehouse.ListWarehouse(c)
	})

	goodsTypeGroup := ginApi.Group("/gt", utils.AuthMiddleware())
	goodsTypeGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		goodsType.AddGoodsType(c)
	})
	goodsTypeGroup.PUT("/update", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		goodsType.UpdateGoodsType(c)
	})
	goodsTypeGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		goodsType.DeleteGoodsType(c)
	})
	goodsTypeGroup.GET("/list", func(c *gin.Context) {
		goodsType.ListGoodsType(c)
	})

	departmentGroup := ginApi.Group("/dept", utils.AuthMiddleware())
	departmentGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		department.AddDepartment(c)
	})
	departmentGroup.PUT("/update", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		department.UpdateDepartment(c)
	})
	departmentGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		department.DeleteDepartment(c)
	})
	departmentGroup.GET("/list", func(c *gin.Context) {
		department.ListDepartment(c)
	})

	staffGroup := ginApi.Group("/staff", utils.AuthMiddleware())
	staffGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		staff.AddStaff(c)
	})
	staffGroup.PUT("/update", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		staff.UpdateStaff(c)
	})
	staffGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		staff.DeleteStaff(c)
	})
	staffGroup.GET("/list", func(c *gin.Context) {
		staff.ListStaff(c)
	})

	inventoryTypeGroup := ginApi.Group("/invt", utils.AuthMiddleware())
	inventoryTypeGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		inventoryType.AddInventoryType(c)
	})
	inventoryTypeGroup.PUT("/update", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		inventoryType.UpdateInventoryType(c)
	})
	inventoryTypeGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		inventoryType.DeleteInventoryType(c)
	})
	inventoryTypeGroup.GET("/list", func(c *gin.Context) {
		inventoryType.ListInventoryType(c)
	})

	unitGroup := ginApi.Group("/unit", utils.AuthMiddleware())
	unitGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		unit.AddUnit(c)
	})
	unitGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		unit.DeleteUnit(c)
	})
	unitGroup.GET("/list", func(c *gin.Context) {
		unit.ListUnit(c)
	})

	goodsGroup := ginApi.Group("/goods", utils.AuthMiddleware())
	goodsGroup.POST("/add", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		goods.AddGoods(c)
	})
	goodsGroup.POST("/update", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		goods.UpdateGoods(c)
	})
	goodsGroup.DELETE("/delete", utils.IsSuperAdminMiddleware(), func(c *gin.Context) {
		goods.DeleteGoods(c)
	})
	goodsGroup.GET("/search", func(c *gin.Context) {
		goods.SearchGoods(c)
	})
}
