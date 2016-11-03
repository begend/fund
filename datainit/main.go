package main

import (
	"fmt"

	"time"

	"bitbucket.org/jixiuf/fund/db"
	"bitbucket.org/jixiuf/fund/defs"
	"bitbucket.org/jixiuf/fund/dt"
	"bitbucket.org/jixiuf/fund/eastmoney"
)

func main() {
	initFundValueHistory()
}

// 初始化历史净值数据
func initFundValueHistory() {
	dbT, _ := dt.NewDatabaseTemplateWithConfig(defs.DBConfig, true)
	db.FundValueHistoryCreateTable(dbT)
	stockList := eastmoney.GetFundIdList(eastmoney.FundTypeAll)
	for idx, fb := range stockList {
		if idx < 950 {
			continue
		}

		fmt.Printf("%d/%d id=%s\n", idx, len(stockList), fb.Id)
		fd, err := eastmoney.GetFund(fb.Id, true)
		if err != nil {
			fmt.Println(err)
			continue
		}
		db.FundValueHistoryInsertAll(dbT, fd)
		time.Sleep(time.Second)
	}
}

func dailyUpdateeFundValue() {
	dbT, _ := dt.NewDatabaseTemplateWithConfig(defs.DBConfig, true)
	db.FundValueHistoryCreateTable(dbT)
	stockList := eastmoney.GetFundIdList(eastmoney.FundTypeAll)
	for idx, fb := range stockList {
		fmt.Printf("%d/%d id=%s\n", idx, len(stockList), fb.Id)
		fd, err := eastmoney.GetFundDetail(fb.Id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		db.FundValueHistoryInsertLast(dbT, fd)
		time.Sleep(time.Second)
	}

}
