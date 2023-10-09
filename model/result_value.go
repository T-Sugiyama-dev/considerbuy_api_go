package model

type ResultValue struct {
	ToolTerm          int `json:"toolTerm"`          // ツール使用年数
	WorkHours         int `json:"workHours"`         // 1日あたりの労働時間
	MonthlyWorkDays   int `json:"monthlyWorkDays"`   // 1月あたりの労働日数
	AnnualWorkDays    int `json:"annualWorkDays"`    // 1年あたりの労働日数
	HourlyWage        int `json:"hourlyWage"`        // 時給
	DailySavedMinutes int `json:"dailySavedMinutes"` // 1日あたりの節約時間
	DailyProfit       int `json:"dailyProfit"`       // 1日あたりの利益
	SavedTime         int `json:"savedTime"`         // 節約時間
	Profit            int `json:"profit"`            // 利益
	TotalCost         int `json:"totalCost"`         // 合計費用
	Roi               int `json:"roi"`               // ROI
	PaybackDays       int `json:"paybackDays"`       // 回収期間
}

// getter

func (r *ResultValue) GetToolTerm() int {
	return r.ToolTerm
}

func (r *ResultValue) GetWorkHours() int {
	return r.WorkHours
}

func (r *ResultValue) GetMonthlyWorkDays() int {
	return r.MonthlyWorkDays
}

func (r *ResultValue) GetAnnualWorkDays() int {
	return r.AnnualWorkDays
}

func (r *ResultValue) GetHourlyWage() int {
	return r.HourlyWage
}

func (r *ResultValue) GetDailySavedMinutes() int {
	return r.DailySavedMinutes
}

func (r *ResultValue) GetDailyProfit() int {
	return r.DailyProfit
}

func (r *ResultValue) GetSavedTime() int {
	return r.SavedTime
}

func (r *ResultValue) GetProfit() int {
	return r.Profit
}

func (r *ResultValue) GetTotalCost() int {
	return r.TotalCost
}

func (r *ResultValue) GetRoi() int {
	return r.Roi
}

func (r *ResultValue) GetPaybackDays() int {
	return r.PaybackDays
}

// setter

func (r *ResultValue) SetToolTerm(toolTerm int) {
	r.ToolTerm = toolTerm
}

func (r *ResultValue) SetWorkHours(workHours int) {
	r.WorkHours = workHours
}

func (r *ResultValue) SetMonthlyWorkDays(monthlyWorkDays int) {
	r.MonthlyWorkDays = monthlyWorkDays
}

func (r *ResultValue) SetAnnualWorkDays(annualWorkDays int) {
	r.AnnualWorkDays = annualWorkDays
}

func (r *ResultValue) SetHourlyWage(hourlyWage int) {
	r.HourlyWage = hourlyWage
}

func (r *ResultValue) SetDailySavedMinutes(dailySavedMinutes int) {
	r.DailySavedMinutes = dailySavedMinutes
}

func (r *ResultValue) SetDailyProfit(dailyProfit int) {
	r.DailyProfit = dailyProfit
}

func (r *ResultValue) SetSavedTime(savedTime int) {
	r.SavedTime = savedTime
}

func (r *ResultValue) SetProfit(profit int) {
	r.Profit = profit
}

func (r *ResultValue) SetTotalCost(totalCost int) {
	r.TotalCost = totalCost
}

func (r *ResultValue) SetRoi(roi int) {
	r.Roi = roi
}

func (r *ResultValue) SetPaybackDays(paybackDays int) {
	r.PaybackDays = paybackDays
}
