package department

import (
	"Go_simpleWMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateDeptRequest struct {
	Did  string `json:"did" form:"did" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

func UpdateDepartment(context *gin.Context) {
	var data updateDeptRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "did and department name are required"})
		return
	}
	did := data.Did
	depName := data.Name

	tx, err := utils.GetDbConnection()

	if tx == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot begin transaction",
			"detail": err.Error(),
		})
		return
	}

	// 判断该部门是否已存在
	var registered int
	err = tx.QueryRow("SELECT count(name) FROM department WHERE did=?", did).Scan(&registered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot get the number of department for this did",
			"detail": err.Error(),
		})
		return
	}
	if registered == 0 {
		context.JSON(http.StatusForbidden, gin.H{"message": "The department does not exist"})
		return
	}

	// 更新部门
	_, err = tx.Exec("UPDATE department SET name=? WHERE did=?", depName, did)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot update the department",
			"detail": err.Error(),
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Cannot commit the transaction",
			"detail": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Department updated successfully"})
}
