package service

import (
	"github.com/T-sugiyama-dev/consider_buy/model"
)

type CalcService struct{}

func (CalcService) Calculate(c *model.CalcValue) (model.ResultValue, error) {

	result := model.ResultValue{}

	monthlyWeeks := 4.35
	annualMonths := 12
	annualWeeks := float64(annualMonths) * monthlyWeeks
	workDays := c.GetWorkDays()
	workHours := c.GetWorkHours()
	toolTerm := c.GetToolTerm()
	toolDays := c.GetToolDays()
	toolHours := c.GetToolHours()
	initialCost := c.GetInitialCost()
	runningCost := c.GetRunningCost()
	income := c.GetIncome()
	productivity := c.GetProductivity() / 100

	// 計算ロジック
	monthlyWorkDays := monthlyWeeks * workDays
	annualWorkDays := float64(annualMonths) * monthlyWorkDays
	hourlyWage := income / float64(annualWorkDays) / workHours
	dailySavedMinutes := toolHours * productivity * 60
	dailyProfit := hourlyWage * dailySavedMinutes / 60
	savedTime := toolHours * productivity * toolDays * annualWeeks * toolTerm
	profit := savedTime * hourlyWage
	totalCost := initialCost + runningCost*toolTerm
	roi := (profit / totalCost) * 100
	paybackDays := totalCost / (hourlyWage * productivity * toolHours)

	result.SetToolTerm(int(toolTerm))
	result.SetWorkHours(int(workHours))
	result.SetMonthlyWorkDays(int(monthlyWorkDays))
	result.SetAnnualWorkDays(int(annualWorkDays))
	result.SetHourlyWage(int(hourlyWage))
	result.SetDailySavedMinutes(int(dailySavedMinutes))
	result.SetDailyProfit(int(dailyProfit))
	result.SetSavedTime(int(savedTime))
	result.SetProfit(int(profit))
	result.SetTotalCost(int(totalCost))
	result.SetRoi(int(roi))
	result.SetPaybackDays(int(paybackDays))

	return result, nil
}
