package action

import (
	"backend/adapter/api/reqres"
	"backend/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

// GetAllDepartmentAction GetDepartmentsHandler godoc
// @Summary      学部情報を取得します
// @Description  全ての学部の情報を取得します
// @Tags         department
// @Accept       json
// @Produce      json
// @Success      200      {object}  []domain.Department
// @Failure      500      {object}  map[string]string
// @Router       /api/departments/get [get]
func GetAllDepartmentAction(w http.ResponseWriter, r *http.Request, useCase usecase.DepartmentUseCase) {

	// メソッドチェック
	if r.Method != http.MethodGet {
		reqres.WriteJSONErrorResponse(w, "Invalid request method")
		return
	}

	// Department情報を取得
	departments, err := useCase.GetAllDepartments()
	if err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}

	// スライスの長さとnilチェック
	if len(departments) <= 0 {
		reqres.WriteJSONErrorResponse(w, "unexpected error")
		return
	}

	departmentMap := make(map[int]string)
	for _, department := range departments {
		departmentMap[department.ID] = department.Name
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(departmentMap); err != nil {
		println(err.Error())
	}

	return
}

// GetDepartmentAction GetDepartmentsHandler godoc
// @Summary      学部情報を取得します
// @Description  全ての学部の情報を取得します
// @Tags         department
// @Accept       json
// @Produce      json
// @Success      200      {object}  domain.Department
// @Failure      500      {object}  map[string]string
// @Router       /api/department/get [get]
func GetDepartmentAction(w http.ResponseWriter, r *http.Request, useCase usecase.DepartmentUseCase) {

	// メソッドチェック
	if r.Method != http.MethodGet {
		reqres.WriteJSONErrorResponse(w, "Invalid request method")
		return
	}

	departmentIDMap := make(map[string]string)

	if err := json.NewDecoder(r.Body).Decode(&departmentIDMap); err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}

	departmentIDInt, err := strconv.Atoi(departmentIDMap["departmentID"])
	if err != nil {
		reqres.WriteJSONErrorResponse(w, "unexpected error")
		return
	}

	// Department情報を取得
	department, err := useCase.GetDepartmentInfo(departmentIDInt)
	if err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}

	departmentMap := make(map[int]string)

	// departmentに対してnilチェック
	if department != nil {
		departmentMap[department.ID] = department.Name
	} else {
		println("nil department found, skipping...")
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(departmentMap); err != nil {
		println(err.Error())
	}

	return
}
