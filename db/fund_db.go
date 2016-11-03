package db

import (
	"fmt"

	"bitbucket.org/jixiuf/fund/dt"
	"bitbucket.org/jixiuf/fund/eastmoney"
)

func FundValueHistoryCreateTable(d dt.DatabaseTemplate) {
	sql := ` create table if not exists fund_value_history(
fundId varchar(6) default '',
name varchar(64) default '',
time timestamp default CURRENT_TIMESTAMP,
value float default 0 comment '净值',
totalValue float default 0 comment '累计净值',
dayRatio float default 0 comment '日增比率',
fenHongType tinyint default 0 comment '1.每份基金份额折算1.012175663份 2.每份派现金0.2150元',
fenHongRatio float default 0 comment '分红比率,如每份基金份额折算1.012175663份',
primary key(fundId,time)
)`
	d.ExecDDL(sql)

}

func FundValueHistoryInsertAll(d dt.DatabaseTemplate, fd eastmoney.Fund) {
	sql := "insert into fund_value_history (fundId,name,time,value,totalValue,dayRatio,fenHongType,fenHongRatio) values"
	for idx, fv := range fd.FundValueList {
		sql += fmt.Sprintf("('%s','%s','%s',%f,%f,%f,%d,%f)", fd.Id, fd.Name, fv.Time.Format("2006-01-02 00:00:00"), fv.Value, fv.TotalValue, fv.DayRatio, fv.FenHongType, fv.FenHongRatio)
		if idx != len(fd.FundValueList)-1 {
			sql += ",\n"
		}
	}
	sql += " ON DUPLICATE KEY UPDATE dayRatio=values(dayRatio),value=values(value),totalValue=values(totalValue),name=values(name),fenHongRatio=values(fenHongRatio),fenHongType=values(fenHongType)"
	d.ExecDDL(sql)
}

func FundValueHistoryInsertLast(d dt.DatabaseTemplate, fd eastmoney.Fund) {
	sql := "insert into fund_value_history (fundId,name,time,value,totalValue,dayRatio,fenHongRatio) values"
	sql += fmt.Sprintf("('%s','%s','%s',%f,%f,%f,%f)", fd.Id, fd.Name, fd.FundValueLastUpdateTime.Format("2006-01-02 00:00:00"), fd.FundValueLast, fd.TotalFundValueLast, fd.DayRatioLast, 0.0)
	sql += " ON DUPLICATE KEY UPDATE dayRatio=values(dayRatio),value=values(value),totalValue=values(totalValue),name=values(name)"
	d.ExecDDL(sql)
}
