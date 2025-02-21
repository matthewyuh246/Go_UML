package main

import "fmt"

// EmployeeData は従業員のデータを表します。
type EmployeeData struct {
	Name       string
	Department string
}

// PayCalculator は給与計算を担当します。
type PayCalculator struct{}

// getRegularHours は給与計算用の労働時間計算ロジックです。（プライベートメソッド）
func (pc *PayCalculator) getRegularHours() {
	fmt.Println("給与計算用の労働時間計算ロジック")
}

// CalculatePay は従業員の給料を計算します。
func (pc *PayCalculator) CalculatePay(employeeData EmployeeData) {
	pc.getRegularHours()
	fmt.Printf("%sの給料を計算しました\n", employeeData.Name)
}

// HourReporter は労働時間のレポートを担当します。
type HourReporter struct{}

// getRegularHours は労働時間レポート専用の労働時間計算ロジックです。（プライベートメソッド）
func (hr *HourReporter) getRegularHours() {
	fmt.Println("労働時間レポート専用の労働時間計算ロジック")
}

// ReportHours は従業員の労働時間をレポートします。
func (hr *HourReporter) ReportHours(employeeData EmployeeData) {
	hr.getRegularHours()
	fmt.Printf("%sの労働時間をレポートしました\n", employeeData.Name)
}

// EmployeeRepository は従業員データの保存を担当します。
type EmployeeRepository struct{}

// Save はデータの保存処理（現状は未実装）。
func (er *EmployeeRepository) Save() {
	// 保存処理の実装がここに来る場合は実装してください。
}

// func main() {
// 	employeeData := EmployeeData{Name: "鈴木", Department: "開発"}
// 	payCalculator := PayCalculator{}
// 	hourReporter := HourReporter{}

// 	fmt.Println("経理部門")
// 	payCalculator.CalculatePay(employeeData)
// 	fmt.Println("")
// 	fmt.Println("人事部門")
// 	hourReporter.ReportHours(employeeData)
// }
