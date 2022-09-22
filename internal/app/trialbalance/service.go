package trialbalance

import (
	"fmt"
	"quran/internal/abstraction"
	"quran/internal/dto"
	"quran/internal/factory"
	"quran/internal/model"
	"quran/internal/repository"
	"quran/pkg/util/helper"
	"quran/pkg/util/response"
	"quran/pkg/util/trxmanager"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type service struct {
	Repository          repository.TrialBalance
	TBDetailRepository  repository.TrialBalanceDetail
	FormatterRepository repository.Formatter
	CoaRepository       repository.Coa
	Db                  *gorm.DB
}

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.TrialBalanceGetRequest) (*dto.TrialBalanceGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.TrialBalanceGetByIDRequest) (*dto.TrialBalanceGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.TrialBalanceCreateRequest) (*dto.TrialBalanceCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.TrialBalanceUpdateRequest) (*dto.TrialBalanceUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.TrialBalanceDeleteRequest) (*dto.TrialBalanceDeleteResponse, error)
	Import(ctx *abstraction.Context, payload *dto.TrialBalanceImportRequest, datas *[]model.TrialBalanceDetailEntity) (*dto.TrialBalanceImportResponse, error)
	Export(ctx *abstraction.Context) (*dto.TrialBalanceExportResponse, error)
}

func NewService(f *factory.Factory) *service {
	repository := f.TrialBalanceRepository
	TBDetailRepository := f.TrialBalanceDetailRepository
	FormatterRepository := f.FormatterRepository
	CoaRepository := f.CoaRepository
	db := f.Db
	return &service{
		Repository:          repository,
		TBDetailRepository:  TBDetailRepository,
		FormatterRepository: FormatterRepository,
		CoaRepository:       CoaRepository,
		Db:                  db,
	}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.TrialBalanceGetRequest) (*dto.TrialBalanceGetResponse, error) {
	data, info, err := s.Repository.Find(ctx, &payload.TrialBalanceFilterModel, &payload.Pagination)
	if err != nil {
		return &dto.TrialBalanceGetResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.TrialBalanceGetResponse{
		Datas:          *data,
		PaginationInfo: *info,
	}
	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.TrialBalanceGetByIDRequest) (*dto.TrialBalanceGetByIDResponse, error) {
	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		return &dto.TrialBalanceGetByIDResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.TrialBalanceGetByIDResponse{
		TrialBalanceEntityModel: *data,
	}
	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.TrialBalanceCreateRequest) (*dto.TrialBalanceCreateResponse, error) {
	var data model.TrialBalanceEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data.Context = ctx
		data.TrialBalanceEntity = payload.TrialBalanceEntity
		result, err := s.Repository.Create(ctx, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.TrialBalanceCreateResponse{}, err
	}
	result := &dto.TrialBalanceCreateResponse{
		TrialBalanceEntityModel: data,
	}
	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.TrialBalanceUpdateRequest) (*dto.TrialBalanceUpdateResponse, error) {
	var data model.TrialBalanceEntityModel

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		if _, err := s.Repository.FindByID(ctx, &payload.ID); err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err)
		}
		data.Context = ctx
		data.TrialBalanceEntity = payload.TrialBalanceEntity
		result, err := s.Repository.Update(ctx, &payload.ID, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.TrialBalanceUpdateResponse{}, err
	}
	result := &dto.TrialBalanceUpdateResponse{
		TrialBalanceEntityModel: data,
	}
	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.TrialBalanceDeleteRequest) (*dto.TrialBalanceDeleteResponse, error) {
	var data model.TrialBalanceEntityModel
	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		if _, err := s.Repository.FindByID(ctx, &payload.ID); err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err)
		}
		data.Context = ctx
		result, err := s.Repository.Delete(ctx, &payload.ID, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.TrialBalanceDeleteResponse{}, err
	}
	result := &dto.TrialBalanceDeleteResponse{
		TrialBalanceEntityModel: data,
	}
	return result, nil
}

func (s *service) Import(ctx *abstraction.Context, payload *dto.TrialBalanceImportRequest, datas *[]model.TrialBalanceDetailEntity) (*dto.TrialBalanceImportResponse, error) {
	var dataTB model.TrialBalanceEntityModel
	currentYear, currentMonth, _ := time.Now().Date()
	currentLocation := time.Now().Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	period := lastOfMonth.Format("2006-01-02")

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		//cek company berdasarkan user
		//belum ada
		//skip
		nolimit := 1000
		criteriaTB := dto.TrialBalanceGetRequest{}
		criteriaTB.Pagination.PageSize = &nolimit
		criteriaTB.Period = &period

		getTrialBalance, _, err := s.Repository.Find(ctx, &criteriaTB.TrialBalanceFilterModel, &criteriaTB.Pagination)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err)
		}
		version := len(*getTrialBalance) + 1

		dataTB.Context = ctx
		dataTB.TrialBalanceEntity = model.TrialBalanceEntity{
			Versions:    version,
			Period:      period,
			FormatterID: 3,
			CompanyID:   1,
		}
		resultTB, err := s.Repository.Create(ctx, &dataTB)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}

		for _, v := range *datas {
			dataTBD := model.TrialBalanceDetailEntityModel{
				Context:                  ctx,
				TrialBalanceDetailEntity: v,
			}
			dataTBD.TrialBalanceID = resultTB.ID
			_, err := s.TBDetailRepository.Create(ctx, &dataTBD)
			if err != nil {
				return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
			}
		}

		dataTB = *resultTB
		return nil
	}); err != nil {
		return &dto.TrialBalanceImportResponse{}, err
	}
	result := &dto.TrialBalanceImportResponse{
		Data: dataTB,
	}

	return result, nil
}

func (s *service) Export(ctx *abstraction.Context) (*dto.TrialBalanceExportResponse, error) {
	fmt.Println("Masuk 1")
	var (
		criteria dto.FormatterGetRequest
	)
	fmt.Println("Masuk 1.1")
	TB := "TRIAL-BALANCE"
	criteria.FormatterFilterModel.FormatterFor = &TB
	fmt.Println("Masuk 1.2")

	data, err := s.FormatterRepository.FindWithDetail(ctx, &criteria.FormatterFilterModel)
	if err != nil {
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	f := excelize.NewFile()
	sheet := f.GetSheetName(f.GetActiveSheetIndex())
	arrStyleColWidth := []map[string]interface{}{
		{"COL": "A", "WIDTH": 0.83},
		{"COL": "B", "WIDTH": 15.38},
		{"COL": "C", "WIDTH": 2.14},
		{"COL": "D", "WIDTH": 2.14},
		{"COL": "E", "WIDTH": 57.45},
		{"COL": "F", "WIDTH": 6.43},
		{"COL": "G", "WIDTH": 17.65},
		{"COL": "H", "WIDTH": 10.71},
		{"COL": "I", "WIDTH": 16.83},
		{"COL": "J", "WIDTH": 10.10},
		{"COL": "K", "WIDTH": 17.65},
		{"COL": "L", "WIDTH": 22.14},
	}
	for _, v := range arrStyleColWidth {
		tmpColWidth := fmt.Sprintf("%f", v["WIDTH"])
		colWidth, _ := strconv.ParseFloat(tmpColWidth, 64)
		err = f.SetColWidth(sheet, fmt.Sprintf("%s", v["COL"]), fmt.Sprintf("%s", v["COL"]), colWidth)
		if err != nil {
			panic(err)
			return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
	}

	styleDefault, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Family: "Arial",
			Size:   10,
		},
	})
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.SetColStyle(sheet, "A:Z", styleDefault)
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	stylingBorderRightOnly, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	stylingBorderTopOnly, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	stylingHeader, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#fac090"},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	stylingHeader2, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Color: "#cc00d1",
			Bold:  true,
		},
		Alignment: &excelize.Alignment{
			WrapText:   true,
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#dbdbdb"},
		},
	})
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	stylingSubTotal, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold: true,
		},
	})
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	stylingTotal, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold: true,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#ccff33"},
		},
	})
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	stylingTotalControl, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold: true,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#3ada24"},
		},
	})
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	err = f.MergeCell(sheet, "B6", "B8")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.MergeCell(sheet, "C6", "E8")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.MergeCell(sheet, "F6", "F8")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.MergeCell(sheet, "H6", "K7")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	stylingCurrency, err := f.NewStyle(&excelize.Style{
		NumFmt: 7,
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	err = f.SetCellStyle(sheet, "B6", "L8", stylingHeader)
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.SetCellStyle(sheet, "F6", "F8", stylingHeader2)
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	err = f.SetCellValue(sheet, "B6", "No Akun")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.SetCellValue(sheet, "C6", "Keterangan")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.SetCellValue(sheet, "F6", "WP Reff")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.SetCellValue(sheet, "G6", "PT xxx")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.SetCellValue(sheet, "G7", "Unaudited")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.SetCellValue(sheet, "G8", "31-Dec-21")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.SetCellValue(sheet, "H6", "Adjustment Journal Entry")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.SetCellValue(sheet, "I8", "Debet")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.SetCellValue(sheet, "K8", "Kredit")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.SetCellFormula(sheet, "L6", "=G6")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.SetCellFormula(sheet, "L7", "=G7")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	err = f.SetCellFormula(sheet, "L8", "=G8")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	row := 10
	var summary []map[string]interface{}
	var total []map[string]interface{}
	for _, v := range data.FormatterDetail {
		// if i == 25 {
		// 	break
		// }
		// var codeCoa string
		if !(v.IsTotal != nil && *v.IsTotal == true) {
			f.SetCellValue(sheet, fmt.Sprintf("B%d", row), v.Description)
		}
		if *v.IsCoa {
			row++
			codeCoa := fmt.Sprintf("%v", v.Code)
			coas, err := s.CoaRepository.FindWithCode(ctx, &codeCoa)
			if err != nil {
				return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
			}
			rowBefore := row
			for _, coa := range *coas {
				// cari di tb
				var (
					desc               string
					nominal_before_aje float64
					nominal_debet_aje  float64
					nominal_kredit_aje float64
					// nominal_after_aje  float64
				)
				codeCoa := fmt.Sprintf("%v", coa.Code)
				tbdetails, err := s.TBDetailRepository.FindWithCode(ctx, &codeCoa)
				if err != nil {
					return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
				}
				for _, tbd := range *tbdetails {
					desc = tbd.Description
					nominal_before_aje = tbd.AmountBeforeAje
					nominal_debet_aje = tbd.AmountAjeDr
					nominal_kredit_aje = tbd.AmountAjeCr
					// nominal_after_aje = tbd.AmountAfterAje
				}
				f.SetCellValue(sheet, fmt.Sprintf("B%d", row), coa.Code)
				// f.SetCellValue(sheet, fmt.Sprintf("E%d", row), desc)
				fmt.Println(desc)
				f.SetCellValue(sheet, fmt.Sprintf("E%d", row), coa.Name)
				f.SetCellValue(sheet, fmt.Sprintf("G%d", row), nominal_before_aje)
				f.SetCellValue(sheet, fmt.Sprintf("I%d", row), nominal_debet_aje)
				f.SetCellValue(sheet, fmt.Sprintf("K%d", row), nominal_kredit_aje)
				// f.SetCellValue(sheet, fmt.Sprintf("L%d", row), nominal_after_aje)
				f.SetCellFormula(sheet, fmt.Sprintf("L%d", row), fmt.Sprintf("=G%d+I%d-K%d", row, row, row))
				row++
			}
			rowAfter := row - 1
			if *v.AutoSummary {
				var tmp = map[string]interface{}{"code": v.Code, "row": row}
				summary = append(summary, tmp)
				f.SetCellValue(sheet, fmt.Sprintf("D%d", row), "Subtotal")
				f.SetCellFormula(sheet, fmt.Sprintf("G%d", row), fmt.Sprintf("=SUM(G%d:G%d)", rowBefore, rowAfter))
				f.SetCellFormula(sheet, fmt.Sprintf("I%d", row), fmt.Sprintf("=SUM(I%d:I%d)", rowBefore, rowAfter))
				f.SetCellFormula(sheet, fmt.Sprintf("K%d", row), fmt.Sprintf("=SUM(K%d:K%d)", rowBefore, rowAfter))
				f.SetCellFormula(sheet, fmt.Sprintf("L%d", row), fmt.Sprintf("=SUM(L%d:L%d)", rowBefore, rowAfter))
				f.SetCellStyle(sheet, fmt.Sprintf("D%d", row), fmt.Sprintf("L%d", row), stylingSubTotal)
				row++
			}
		}

		if v.IsTotal != nil && *v.IsTotal == true {
			fxSummary := regexp.MustCompile("[\\*\\-\\+\\/\\s]+").Split(v.FxSummary, -1)
			if len(fxSummary) == 0 {
				if len(v.FxSummary) == 0 {
					f.SetCellValue(sheet, fmt.Sprintf("A%d", row), "ERROR NO FORMULA")
					continue
				}
				f.SetCellValue(sheet, fmt.Sprintf("A%d", row), "FORMULA NOT PROVIDED")
				continue
			}
			formulaTotalAmountBeforeAje := v.FxSummary
			formulaTotalAmountAjeDr := v.FxSummary
			formulaTotalAmountAjeCr := v.FxSummary
			formulaTotalAmountAfterAje := v.FxSummary
			var tmpArrFx = make(map[string][]string)
			var tmpArrFx2 = make(map[string][]string)
			var tmpArrFx3 = make(map[string][]string)
			var tmpArrFx4 = make(map[string][]string)
			// var tmpFx map[string][]string
			for _, fx := range fxSummary {
				for _, sum := range summary {
					if strings.Index(fmt.Sprintf("%v", sum["code"]), fx) >= 0 {
						tmpArrFx[fx] = append(tmpArrFx[fx], fmt.Sprintf("G%v", sum["row"]))
						tmpArrFx2[fx] = append(tmpArrFx2[fx], fmt.Sprintf("I%v", sum["row"]))
						tmpArrFx3[fx] = append(tmpArrFx3[fx], fmt.Sprintf("K%v", sum["row"]))
						tmpArrFx4[fx] = append(tmpArrFx4[fx], fmt.Sprintf("L%v", sum["row"]))
					}
				}
				formulaTotalAmountBeforeAje = helper.ReplaceWholeWord(formulaTotalAmountBeforeAje, fx, strings.Join(tmpArrFx[fx], "+"))
				formulaTotalAmountAjeDr = helper.ReplaceWholeWord(formulaTotalAmountAjeDr, fx, strings.Join(tmpArrFx2[fx], "+"))
				formulaTotalAmountAjeCr = helper.ReplaceWholeWord(formulaTotalAmountAjeCr, fx, strings.Join(tmpArrFx3[fx], "+"))
				formulaTotalAmountAfterAje = helper.ReplaceWholeWord(formulaTotalAmountAfterAje, fx, strings.Join(tmpArrFx4[fx], "+"))
			}
			var tmp = map[string]interface{}{"code": v.Code, "row": row}
			total = append(total, tmp)
			f.SetCellValue(sheet, fmt.Sprintf("C%d", row), v.Description)
			f.SetCellFormula(sheet, fmt.Sprintf("G%d", row), fmt.Sprintf("=%s", formulaTotalAmountBeforeAje))
			f.SetCellFormula(sheet, fmt.Sprintf("I%d", row), fmt.Sprintf("=%s", formulaTotalAmountAjeDr))
			f.SetCellFormula(sheet, fmt.Sprintf("K%d", row), fmt.Sprintf("=%s", formulaTotalAmountAjeCr))
			f.SetCellFormula(sheet, fmt.Sprintf("L%d", row), fmt.Sprintf("=%s", formulaTotalAmountAfterAje))

			if v.IsControl != nil && *v.IsControl == true {
				f.SetCellStyle(sheet, fmt.Sprintf("B%d", row), fmt.Sprintf("L%d", row), stylingTotalControl)
			} else {
				f.SetCellStyle(sheet, fmt.Sprintf("B%d", row), fmt.Sprintf("L%d", row), stylingTotal)
			}
			row++
		}
		row++
	}

	if err = f.SetCellStyle(sheet, fmt.Sprintf("B%d", row), fmt.Sprintf("L%d", row), stylingBorderTopOnly); err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	if err = f.SetCellStyle(sheet, "G9", fmt.Sprintf("L%d", row-1), stylingCurrency); err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	if err = f.SetCellStyle(sheet, "A9", fmt.Sprintf("A%d", row-1), stylingBorderRightOnly); err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	if err = f.SetSheetFormatPr(sheet, excelize.DefaultRowHeight(12.85)); err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	f.SetDefaultFont("Arial")

	err = f.SaveAs("assets/TrialBalance.xlsx")
	if err != nil {
		panic(err)
		return &dto.TrialBalanceExportResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.TrialBalanceExportResponse{
		File: "assets/TrialBalance.xlsx",
	}
	return result, nil
}
