package connection

import (
        "encoding/csv" 
        "fmt"
        "database/sql"
       _"github.com/lib/pq"
        "log"
        "os"
        "encoding/json"
       )
var db *sql.DB
var err error

type Declare struct{

	PolicyID int `json:"id"`
	Statecode, County, Eq_site_limit, Hu_site_limit, Fl_site_limit, Fr_site_limit, Tiv_2011, Tiv_2012, Eq_site_deductible, Hu_site_deductible, Fl_site_deductible, Fr_site_deductible, Point_latitude, Point_longitude, Line, Construction string
	Point_granularity int
}


func DB_con(){
    db, err = sql.Open("postgres", "user=postgres password='root' dbname='sample' sslmode=disable")
	if err != nil {
		log.Fatal(err)
		fmt.Println("there are errors")
	}else {
		fmt.Println("================database db_Connection connected========")
	}
}

func DB_insert() string{

    var sStmt string = "insert into detail_db (policyID, statecode, county, eq_site_limit, hu_site_limit, fl_site_limit, fr_site_limit, tiv_2011, tiv_2012, eq_site_deductible, hu_site_deductible, fl_site_deductible, fr_site_deductible, point_latitude, point_longitude, line, construction, point_granularity) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)"
    var sStmt1 string = "truncate table detail_db"
	file_cs, err := os.Open("sample.csv")
	    reader := csv.NewReader(file_cs)
	    rawcsvdata,_ :=reader.ReadAll()
	//=====================================truncate the file===========================
	   db.Exec(sStmt1)
	 for _, each := range rawcsvdata {
	   db.Exec(sStmt, each[0],each[1],each[2],each[3],each[4],each[5],each[6],each[7],each[8],each[9],each[10],each[11],each[12],each[13],each[14],each[15],each[16],each[17])
	    if err != nil {
	    log.Fatal(err)
	  }
	}
	var d1 []Declare
	   fmt.Println("================completed==================")
	   b,_ := db.Query("SELECT * FROM detail_db")
	   for b.Next() {
       var d Declare
       b.Scan(&d.PolicyID, &d.Statecode, &d.County, &d.Eq_site_limit, &d.Hu_site_limit, &d.Fl_site_limit, &d.Fr_site_limit, &d.Tiv_2011, &d.Tiv_2012, &d.Eq_site_deductible, &d.Hu_site_deductible, &d.Fl_site_deductible, &d.Fr_site_deductible, &d.Point_latitude, &d.Point_longitude, &d.Line, &d.Construction, &d.Point_granularity) 
       d1=append(d1,d)
	   }
	   d_json,_ := json.Marshal(d1)
	   fmt.Println(len(d_json))
return string(d_json)
}

func DB_Search(id int) string{

    rows, err2 := db.Query("SELECT * FROM detail_db WHERE policyid = $1",id)
	if err2 != nil {
	    log.Fatal(err2)
	}
    var d Declare
    var d_store []Declare
   	for rows.Next() {
     if err2 := rows.Scan(&d.PolicyID, &d.Statecode, &d.County, &d.Eq_site_limit, &d.Hu_site_limit, &d.Fl_site_limit, &d.Fr_site_limit, &d.Tiv_2011, &d.Tiv_2012, &d.Eq_site_deductible, &d.Hu_site_deductible, &d.Fl_site_deductible, &d.Fr_site_deductible, &d.Point_latitude, &d.Point_longitude, &d.Line, &d.Construction, &d.Point_granularity); err2 != nil {
	         log.Fatal(err2)
	     }
    d_store = append(d_store,d)    
	}
    if err := rows.Err(); err != nil {
      log.Fatal(err)
  }
    d_json,_ := json.Marshal(d_store)
  fmt.Println(string(d_json))
return string(d_json)

}

