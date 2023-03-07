package db

import (
	"fmt"
	"testing"
)

func init() {

}

func Test_Query(t *testing.T) {
	fmt.Println("Test db Query function")
	conn, err := Open(1)
	if nil != err {
		panic(err)
	}

	rows, err := conn.Query("select * from t_vehicle_name")
	if nil != err {
		panic(err)
	}

	for rows.Next() {
		vehicleName := ""
		err = rows.Scan(&vehicleName)
		if nil != err {
			panic(err)
		}

		fmt.Printf("Vehicle Name: %s\n", vehicleName)
	}
}
