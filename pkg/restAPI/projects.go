package restAPI

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a APIRouter) GetProjects(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	intNum, err := strconv.Atoi(mux.Vars(r)["num"])
	if err != nil {
		a.logger.WithContext(ctx).Error("error in reading route variables")
		a.sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	projects, err := a.projectService.GetProjectNamesWithForkCount(ctx, intNum)
	if err != nil {
		a.logger.WithContext(ctx).Error("error processing request: ", err.Error())
		a.sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.sendResponseData(w, projects)
}
