package api

import (
	"errors"
	"hrm/db"
	"strconv"
	"strings"
)

type PartsRowInfo struct {
	Id                uint64  `json:"id"`
	VehicleName       string  `json:"vehicle_name"`
	PartsName         string  `json:"parts_name"`
	PartsNo           string  `json:"parts_no"`
	PartsCount        uint64  `json:"parts_count"`
	PartsPrice        float64 `json:"parts_price"`
	PartsSellingPrice float64 `json:"parts_selling_price"`
}

type AllPartsRowInfo struct {
	RowCount uint64         `json:"row_count"`
	RowInfos []PartsRowInfo `json:"row_infos"`
}

type CountInfo struct {
	Count uint64 `json:"count"`
}

type SearchInfo struct {
	KeyWord string `json:"key_word"`
}

type InsertInfo struct {
	NewId uint64 `json:"new_id"`
}

type PagingInfo struct {
	LimitPagingNum int `json:"limit_paging_num"`
	Index          int `json:"index"`
}

func AllPartsRow(isAdmin bool) (*AllPartsRowInfo, error) {
	pgInfo, err := db.New()
	if nil != err {
		return nil, err
	}
	defer pgInfo.Close()

	rows, err := pgInfo.Conn.Query("SELECT * FROM t_parts")
	if nil != err {
		return nil, err
	}
	defer rows.Close()

	allRowInfo := &AllPartsRowInfo{}
	var rowInfos []PartsRowInfo
	row := &PartsRowInfo{}
	price := ""
	sellingPrice := ""
	for rows.Next() {
		err = rows.Scan(&row.Id, &row.VehicleName, &row.PartsName, &row.PartsNo, &row.PartsCount, &price, &sellingPrice)
		if nil != err {
			return nil, err
		}

		// 去除单价前面的$、￥等不同的货币字符
		p := string(([]rune(price))[1:])
		row.PartsPrice, err = strconv.ParseFloat(strings.Replace(p, ",", "", -1), 64)
		if nil != err {
			return nil, err
		}

		sp := string(([]rune(sellingPrice))[1:])
		row.PartsSellingPrice, err = strconv.ParseFloat(strings.Replace(sp, ",", "", -1), 64)
		if nil != err {
			return nil, err
		}

		if !isAdmin {
			row.PartsPrice = 0.0
		}

		rowInfos = append(rowInfos, *row)
	}

	allRowInfo.RowCount = uint64(len(rowInfos))
	allRowInfo.RowInfos = rowInfos
	return allRowInfo, nil
}

func Count() (*CountInfo, error) {
	pgInfo, err := db.New()
	if nil != err {
		return nil, err
	}
	defer pgInfo.Close()

	rows, err := pgInfo.Conn.Query("SELECT COUNT(*) FROM t_parts")
	if nil != err {
		return nil, err
	}
	defer rows.Close()

	ci := &CountInfo{}
	if !rows.Next() {
		ci.Count = 0
	}

	err = rows.Scan(&ci.Count)
	if nil != err {
		return nil, err
	}

	return ci, nil
}

func UpdatePartsRow(row *PartsRowInfo) error {
	pgInfo, err := db.New()
	if nil != err {
		return err
	}
	defer pgInfo.Close()

	sql := `UPDATE t_parts
	SET vehicle_name=$1, parts_name=$2, parts_no=$3, parts_count=$4, parts_price=$5, parts_selling_price=$6
	WHERE parts_id=$7;`
	stmt, err := pgInfo.Conn.Prepare(sql)
	if nil != err {
		return err
	}
	defer stmt.Close()

	ret, err := stmt.Exec(
		row.VehicleName,
		row.PartsName,
		row.PartsNo,
		row.PartsCount,
		row.PartsPrice,
		row.PartsSellingPrice,
		row.Id)

	if nil != err {
		return err
	}

	ar, err := ret.RowsAffected()
	if -1 == ar {
		return errors.New("Update fail, affected row error")
	}

	return err
}

func InsertPartsRow(row *PartsRowInfo) (*InsertInfo, error) {
	pgInfo, err := db.New()
	if nil != err {
		return nil, err
	}
	defer pgInfo.Close()

	sql := `INSERT INTO t_parts(parts_id, vehicle_name, parts_name, parts_no, parts_count, parts_price, parts_selling_price)
	VALUES (nextval('t_parts_parts_id_seq'), $1, $2, $3, $4, $5, $6) RETURNING parts_id;`

	insertInfo := &InsertInfo{}
	err = pgInfo.Conn.QueryRow(sql,
		row.VehicleName,
		row.PartsName,
		row.PartsNo,
		row.PartsCount,
		row.PartsPrice,
		row.PartsSellingPrice).Scan(&insertInfo.NewId)
	if nil != err {
		return nil, err
	}

	return insertInfo, nil
}

func DeletePartsRow(row *PartsRowInfo) error {
	pgInfo, err := db.New()
	if nil != err {
		return err
	}
	defer pgInfo.Close()

	sql := `DELETE FROM t_parts
	WHERE parts_id=$1;`

	stmt, err := pgInfo.Conn.Prepare(sql)
	if nil != err {
		return err
	}
	defer stmt.Close()

	ret, err := stmt.Exec(row.Id)
	if nil != err {
		return err
	}

	ra, err := ret.RowsAffected()
	if -1 == ra {
		return errors.New("Delete fail, affected row error")
	}

	return err
}

func Search(sInfo *SearchInfo, isAdmin bool) (*AllPartsRowInfo, error) {
	pgInfo, err := db.New()
	if nil != err {
		return nil, err
	}
	defer pgInfo.Close()

	sql := `SELECT * FROM t_parts WHERE vehicle_name like $1 OR parts_name like $2 OR parts_no like $3;`
	stmt, err := pgInfo.Conn.Prepare(sql)
	if nil != err {
		return nil, err
	}
	defer stmt.Close()

	key := "%" + sInfo.KeyWord + "%"
	rows, err := stmt.Query(key, key, key)
	if nil != err {
		return nil, err
	}
	defer rows.Close()

	allRowInfo := &AllPartsRowInfo{}
	var partsRows []PartsRowInfo
	row := &PartsRowInfo{}
	price := ""
	sellingPrice := ""
	for rows.Next() {
		err = rows.Scan(&row.Id, &row.VehicleName, &row.PartsName, &row.PartsNo, &row.PartsCount, &price, &sellingPrice)
		if nil != err {
			return nil, err
		}

		// 去除单价前面的$、￥等不同的货币字符
		p := string(([]rune(price))[1:])
		row.PartsPrice, err = strconv.ParseFloat(strings.Replace(p, ",", "", -1), 64)
		if nil != err {
			return nil, err
		}

		sp := string(([]rune(price))[1:])
		row.PartsSellingPrice, err = strconv.ParseFloat(strings.Replace(sp, ",", "", -1), 64)
		if nil != err {
			return nil, err
		}

		if !isAdmin {
			row.PartsPrice = 0.0
		}

		partsRows = append(partsRows, *row)
	}

	allRowInfo.RowCount = uint64(len(partsRows))
	allRowInfo.RowInfos = partsRows
	return allRowInfo, nil
}

func Paging(isAdmin bool, info *PagingInfo) (*AllPartsRowInfo, error) {
	pgInfo, err := db.New()
	if nil != err {
		return nil, err
	}
	defer pgInfo.Close()

	sql := `SELECT * FROM t_parts LIMIT $1 OFFSET $2;`
	stmt, err := pgInfo.Conn.Prepare(sql)
	if nil != err {
		return nil, err
	}
	defer stmt.Close()

	offset := info.LimitPagingNum * (info.Index - 1)
	rows, err := stmt.Query(info.LimitPagingNum, offset)
	if nil != err {
		return nil, err
	}
	defer rows.Close()

	allRowInfo := &AllPartsRowInfo{}
	var rowInfos []PartsRowInfo
	row := &PartsRowInfo{}
	price := ""
	sellingPrice := ""
	for rows.Next() {
		err = rows.Scan(&row.Id, &row.VehicleName, &row.PartsName, &row.PartsNo, &row.PartsCount, &price, &sellingPrice)
		if nil != err {
			return nil, err
		}

		// 去除单价前面的$、￥等不同的货币字符
		p := string(([]rune(price))[1:])
		row.PartsPrice, err = strconv.ParseFloat(strings.Replace(p, ",", "", -1), 64)
		if nil != err {
			return nil, err
		}

		sp := string(([]rune(sellingPrice))[1:])
		row.PartsSellingPrice, err = strconv.ParseFloat(strings.Replace(sp, ",", "", -1), 64)
		if nil != err {
			return nil, err
		}

		if !isAdmin {
			row.PartsPrice = 0.0
		}

		rowInfos = append(rowInfos, *row)
	}

	allRowInfo.RowCount = uint64(len(rowInfos))
	allRowInfo.RowInfos = rowInfos
	return allRowInfo, nil
}
