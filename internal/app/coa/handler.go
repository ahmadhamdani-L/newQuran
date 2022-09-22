package coa

import (
	"fmt"
	"net/http"
	"quran/internal/abstraction"
	"quran/internal/dto"
	"quran/internal/factory"
	"quran/internal/model"
	"quran/pkg/util/response"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

type handler struct {
	service *service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

var err error

func (h *handler) Get(c echo.Context) error {
	cc := c.(*abstraction.Context)
	payload := new(dto.CoaGetRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Find(cc, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}
	return response.CustomSuccessBuilder(http.StatusOK, result.Datas, "Get Data Success", &result.PaginationInfo).Send(c)
}

func (h *handler) GetByID(c echo.Context) error {
	cc := c.(*abstraction.Context)
	payload := new(dto.CoaGetByIDRequest)

	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.FindByID(cc, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

func (h *handler) Create(c echo.Context) error {
	cc := c.(*abstraction.Context)
	payload := new(dto.CoaCreateRequest)

	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Create(cc, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

func (h *handler) Update(c echo.Context) error {
	cc := c.(*abstraction.Context)
	payload := new(dto.CoaUpdateRequest)

	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Update(cc, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

func (h *handler) Delete(c echo.Context) error {
	cc := c.(*abstraction.Context)
	payload := new(dto.CoaDeleteRequest)

	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Delete(cc, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

func (h *handler) Export(c echo.Context) error {
	cc := c.(*abstraction.Context)
	payload := new(dto.CoaGetRequest)

	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	nolimit := 100000
	page := 1
	payload.PageSize = &nolimit
	payload.Page = &page

	if payload.Sort == nil {
		coa := "code"
		asc := "asc"
		payload.SortBy = &coa
		payload.Sort = &asc
	}

	result, err := h.service.Find(cc, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	f := excelize.NewFile()
	sheetCoa := "COA"
	f.SetSheetName(f.GetSheetName(0), sheetCoa)

	f.SetCellValue(sheetCoa, "A1", "COA")
	f.SetCellValue(sheetCoa, "B1", "Account Name")
	f.SetCellValue(sheetCoa, "H1", "Grouping")

	row := 2
	for _, v := range result.Datas {
		coaSplit := strings.Split(v.Name, "-")
		roww := fmt.Sprintf("%d", row)
		f.SetCellValue(sheetCoa, "A"+roww, v.Code)
		f.SetCellValue(sheetCoa, "B"+roww, v.Name)
		f.SetCellValue(sheetCoa, "C"+roww, nil)
		f.SetCellValue(sheetCoa, "D"+roww, "-")
		if len(coaSplit) > 0 && coaSplit[0] != "" {
			f.SetCellValue(sheetCoa, "E"+roww, coaSplit[0])
		}
		if len(coaSplit) > 1 && coaSplit[1] != "" {
			f.SetCellValue(sheetCoa, "F"+roww, coaSplit[1])
		}
		if len(coaSplit) > 2 && coaSplit[2] != "" {
			f.SetCellValue(sheetCoa, "G"+roww, coaSplit[2])
		}

		f.SetCellValue(sheetCoa, "H"+roww, v.CoaGroup.Name)
		row += 1
	}

	err = f.AutoFilter(sheetCoa, "A1", "H1", "")
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	err = f.SaveAs("assets/COA.xlsx")
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	file := dto.CoaExportResponse{
		File: "assets/COA.xlsx",
	}

	return response.SuccessResponse(&file).Send(c)
}

func (h *handler) Import(c echo.Context) error {
	// cc := c.(*abstraction.Context)
	// payload := new(dto.CoaImportRequest)

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("error")
		// return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
		return nil
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println("error2")
		return nil
		// return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	defer src.Close()

	f, err := excelize.OpenReader(src)
	if err != nil {
		fmt.Println("error3")
		return nil
		// return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("OKE")
	sheet := f.GetSheetName(f.GetActiveSheetIndex())
	rows, err := f.GetRows(sheet)
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	headA1, err := f.GetCellValue(sheet, "A1")
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	headB1, err := f.GetCellValue(sheet, "B1")
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	headC1, err := f.GetCellValue(sheet, "C1")
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	if strings.ToLower(headA1) != "coa" || strings.ToLower(headB1) != "account name" || strings.ToLower(headC1) != "coa group code" {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	fmt.Println("oke", strings.ToLower(headA1) != "coa")
	rows = rows[1:][:]

	datas := []model.CoaEntity{}

	for y, row := range rows {
		coaGroupId, err := strconv.Atoi(row[2])
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
		}
		data := model.CoaEntity{
			Code:       row[0],
			Name:       row[1],
			CoaGroupId: coaGroupId,
		}
		// for x, colCell := range row {
		// cellCoordinates, err := excelize.CoordinatesToCellName(x+1, y+1)
		// if err != nil {
		// 	return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
		// }
		// }
		datas = append(datas, data)
		if y > 10 {
			break
		}
	}

	fmt.Println(datas)

	return response.SuccessResponse("CEKI").Send(c)
}
