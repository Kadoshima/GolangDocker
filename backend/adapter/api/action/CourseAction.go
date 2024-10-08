package action

import (
	"backend/adapter/api/reqres"
	"backend/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

// GetAllCourseInfoAction GetCourseHandler godoc
// @Summary      コース情報をすべて取得します
// @Description  全てのコース情報を取得します
// @Tags         courses
// @Accept       json
// @Produce      json
// @Success      200      {object}  []domain.Course
// @Failure      404      {object}  map[string]string
// @Router       /api/courses [get]
func GetAllCourseInfoAction(w http.ResponseWriter, r *http.Request, useCase usecase.CourseUseCase) {
	//メソッドチェック
	if r.Method != http.MethodGet {
		reqres.WriteJSONErrorResponse(w, "Method error")
		return
	}

	// Course情報を取ってくる
	courses, err := useCase.GetAllCourseInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// スライスの長さとnilチェック
	if len(courses) <= 0 {
		println("courses slice is either too short or contains nil elements")
	}

	courseMap := make(map[int]string)
	for _, course := range courses {
		courseMap[course.ID] = course.Name
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(courses); err != nil {
		reqres.WriteJSONErrorResponse(w, "Json encode error")
	}

	return
}

// GetCourseInfoAction GetCourseHandler godoc
// @Summary      コース情報を取得します
// @Description  指定したコースの情報を取得します
// @Tags         course
// @Accept       json
// @Produce      json
// @Param        courseID  query     string  true  "コースID"
// @Success      200      {object}  domain.Course
// @Failure      404      {object}  map[string]string
// @Router       /api/course [get]
func GetCourseInfoAction(w http.ResponseWriter, r *http.Request, useCase usecase.CourseUseCase) {
	//メソッドチェック
	if r.Method != http.MethodGet {
		reqres.WriteJSONErrorResponse(w, "Invalid request method")
		return
	}

	courseID := make(map[string]string)

	if err := json.NewDecoder(r.Body).Decode(&courseID); err != nil {
		println(err.Error())
		return
	}

	courseIDInt, err := strconv.Atoi(courseID["courseID"])

	// Course情報を取ってくる
	course, err := useCase.GetCourseInfo(courseIDInt)
	if err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}

	courseMap := make(map[int]string) // マップの初期化

	// courseに対してnilチェック
	if course != nil {
		courseMap[course.ID] = course.Name
	} else {
		println("nil course found, skipping...")
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(courseMap); err != nil {
		println(err.Error())
	}

	return
}
