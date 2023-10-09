package model

type CalcValue struct {
	InitialCost  float64 `json:"initialCost"`  // 初期費用
	RunningCost  float64 `json:"runningCost"`  // ランニングコスト
	ToolTerm     float64 `json:"toolTerm"`     // ツール使用期間
	ToolDays     float64 `json:"toolDays"`     // 1週間あたりのツール使用日数
	ToolHours    float64 `json:"toolHours"`    // 1日あたりのツール使用時間
	Productivity float64 `json:"productivity"` // 生産性
	Income       float64 `json:"income"`       // 年収
	WorkDays     float64 `json:"workDays"`     // 1週間あたりの労働日数
	WorkHours    float64 `json:"workHours"`    // 1日あたりの労働時間
}

// getter

func (c *CalcValue) GetInitialCost() float64 {
	return c.InitialCost
}

func (c *CalcValue) GetRunningCost() float64 {
	return c.RunningCost
}

func (c *CalcValue) GetToolTerm() float64 {
	return c.ToolTerm
}

func (c *CalcValue) GetToolDays() float64 {
	return c.ToolDays
}

func (c *CalcValue) GetToolHours() float64 {
	return c.ToolHours
}

func (c *CalcValue) GetProductivity() float64 {
	return c.Productivity
}

func (c *CalcValue) GetIncome() float64 {
	return c.Income
}

func (c *CalcValue) GetWorkDays() float64 {
	return c.WorkDays
}

func (c *CalcValue) GetWorkHours() float64 {
	return c.WorkHours
}

// setter

func (c *CalcValue) SetInitialCost(initialCost float64) {
	c.InitialCost = initialCost
}

func (c *CalcValue) SetRunningCost(runningCost float64) {
	c.RunningCost = runningCost
}

func (c *CalcValue) SetToolTerm(toolTerm float64) {
	c.ToolTerm = toolTerm
}

func (c *CalcValue) SetTooldays(toolDays float64) {
	c.ToolDays = toolDays
}

func (c *CalcValue) SetToolHours(toolHours float64) {
	c.ToolHours = toolHours
}

func (c *CalcValue) SetProductivity(productivity float64) {
	c.Productivity = productivity
}

func (c *CalcValue) SetIncome(income float64) {
	c.Income = income
}

func (c *CalcValue) SetWorkDays(workDays float64) {
	c.WorkDays = workDays
}

func (c *CalcValue) SetWorkHours(workHours float64) {
	c.WorkHours = workHours
}
