package trialbalance

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

var err error

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Get(c echo.Context) error {
	cc := c.(*abstraction.Context)
	payload := new(dto.TrialBalanceGetRequest)

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
	payload := new(dto.TrialBalanceGetByIDRequest)

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
	payload := new(dto.TrialBalanceCreateRequest)
	fmt.Println("Masuk 1")
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	fmt.Println("Masuk 2")
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}
	fmt.Println("Masuk 3")
	result, err := h.service.Create(cc, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

func (h *handler) Update(c echo.Context) error {
	cc := c.(*abstraction.Context)
	payload := new(dto.TrialBalanceUpdateRequest)

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
	payload := new(dto.TrialBalanceDeleteRequest)

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
	payload := new(dto.TrialBalanceGetRequest)

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

	result, err := h.service.Export(cc)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

func (h *handler) Import(c echo.Context) error {
	cc := c.(*abstraction.Context)
	payload := new(dto.TrialBalanceImportRequest)

	file, err := c.FormFile("file")
	if err != nil {
		// return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
		return nil
	}

	src, err := file.Open()
	if err != nil {
		return nil
		// return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	defer src.Close()

	f, err := excelize.OpenReader(src)
	if err != nil {
		return nil
		// return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheet := f.GetSheetName(f.GetActiveSheetIndex())
	rows, err := f.GetRows(sheet)
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	head1, err := f.GetCellValue(sheet, "B6")
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	head2, err := f.GetCellValue(sheet, "C6")
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	head3, err := f.GetCellValue(sheet, "F6")
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	head4, err := f.GetCellValue(sheet, "G7")
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	head5, err := f.GetCellValue(sheet, "H6")
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	if strings.ToLower(head1) != "no akun" || strings.ToLower(head2) != "keterangan" || strings.ToLower(head3) != "wp reff" || strings.ToLower(head4) != "unaudited" || strings.ToLower(head5) != "adjustment journal entry" {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	rows = rows[8:][:]

	datas := []model.TrialBalanceDetailEntity{}

	for _, row := range rows {
		if len(row) > 1 && len(row[1]) == 9 {
			if len(row) < 6 {
				continue
			}
			nominal, err := strconv.ParseFloat(row[6], 64)
			if err != nil {
				continue
			}
			coa := row[1]
			nominalBeforeAje := nominal
			data := model.TrialBalanceDetailEntity{
				Code:            coa,
				AmountBeforeAje: nominalBeforeAje,
				Description:     row[4],
			}
			datas = append(datas, data)
		}
	}

	payload = &dto.TrialBalanceImportRequest{
		UserId: cc.Auth.ID,
	}

	result, err := h.service.Import(cc, payload, &datas)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}
