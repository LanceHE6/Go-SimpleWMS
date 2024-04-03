package route

import (
	"Go_simpleWMS/handler/auth"
	"Go_simpleWMS/handler/department"
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
)

func Route(ginServer *gin.Engine) {
	// 设置静态资源路径
	ginServer.Static("/res", "./static/res")

	ginServer.GET("/ping", func(context *gin.Context) {
		test.Ping(context)
	})
	// 鉴权接口
	ginServer.GET("/auth", utils.AuthMiddleware(), func(context *gin.Context) {
		auth.AuthByHeader(context)
	})
	ginServer.POST("/upload", func(context *gin.Context) {
		upload.UploadFile(context)
	})
	// 路由分组
	userGroup := ginServer.Group("/user")

	userGroup.POST("/register", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		user.Register(context)
	})
	userGroup.POST("/upload", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		user.UploadUsers(context)
	})
	userGroup.POST("/login", func(context *gin.Context) {
		user.Login(context)
	})
	userGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		user.DeleteUser(context)
	})
	userGroup.PUT("/update", utils.AuthMiddleware(), func(context *gin.Context) {
		user.UpdateUser(context)
	})
	userGroup.GET("/list", utils.AuthMiddleware(), func(context *gin.Context) {
		user.ListUsers(context)
	})

	warehouseGroup := ginServer.Group("/warehouse")
	warehouseGroup.POST("/add", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		warehouse.AddWarehouse(context)
	})
	warehouseGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		warehouse.DeleteWarehouse(context)
	})
	warehouseGroup.PUT("/update", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		warehouse.UpdateWarehouse(context)
	})
	warehouseGroup.GET("/list", utils.AuthMiddleware(), func(context *gin.Context) {
		warehouse.ListWarehouse(context)
	})

	goodsTypeGroup := ginServer.Group("/gt")
	goodsTypeGroup.POST("/add", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		goodsType.AddGoodsType(context)
	})
	goodsTypeGroup.PUT("/update", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		goodsType.UpdateGoodsType(context)
	})
	goodsTypeGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		goodsType.DeleteGoodsType(context)
	})
	goodsTypeGroup.GET("/list", utils.AuthMiddleware(), func(context *gin.Context) {
		goodsType.ListGoodsType(context)
	})

	departmentGroup := ginServer.Group("/dept")
	departmentGroup.POST("/add", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		department.AddDepartment(context)
	})
	departmentGroup.PUT("/update", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		department.UpdateDepartment(context)
	})
	departmentGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		department.DeleteDepartment(context)
	})
	departmentGroup.GET("/list", utils.AuthMiddleware(), func(context *gin.Context) {
		department.ListDepartment(context)
	})

	staffGroup := ginServer.Group("/staff")
	staffGroup.POST("/add", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		staff.AddStaff(context)
	})
	staffGroup.PUT("/update", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		staff.UpdateStaff(context)
	})
	staffGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		staff.DeleteStaff(context)
	})
	staffGroup.GET("/list", utils.AuthMiddleware(), func(context *gin.Context) {
		staff.ListStaff(context)
	})

	inventoryTypeGroup := ginServer.Group("/invt")
	inventoryTypeGroup.POST("/add", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		inventoryType.AddInventoryType(context)
	})
	inventoryTypeGroup.PUT("/update", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		inventoryType.UpdateInventoryType(context)
	})
	inventoryTypeGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		inventoryType.DeleteInventoryType(context)
	})
	inventoryTypeGroup.GET("/list", utils.AuthMiddleware(), func(context *gin.Context) {
		inventoryType.ListInventoryType(context)
	})

	unitGroup := ginServer.Group("/unit")
	unitGroup.POST("/add", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		unit.AddUnit(context)
	})
	unitGroup.DELETE("/delete", utils.AuthMiddleware(), utils.IsSuperAdminMiddleware(), func(context *gin.Context) {
		unit.DeleteUnit(context)
	})
	unitGroup.GET("/list", utils.AuthMiddleware(), func(context *gin.Context) {
		unit.ListUnit(context)
	})
}
