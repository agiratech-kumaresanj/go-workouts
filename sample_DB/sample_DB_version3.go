package main

import (
    "encoding/csv"
    "fmt"
    "encoding/json"
    "database/sql"
    _"github.com/lib/pq"
    "log"
     "os"
    )

type declare struct{

	policyID int
 	statecode, county, eq_site_limit, hu_site_limit, fl_site_limit, fr_site_limit, tiv_2011, tiv_2012, eq_site_deductible, hu_site_deductible, fl_site_deductible, fr_site_deductible, point_latitude, point_longitude, line, construction string
	point_granularity int
}

func main() {
    //var con string
    //fmt.Println("enter: 1.yes or 2.no")
	
	//fmt.Scanf("%s", &con)
	db_Connection()
	//db_Search()
	/*if(con == "yes"){
    }else{
    	main()
    }*/
}


func db_Connection(){
    
//===================================INSERT QUERY in one veriable=====================
    var sStmt string = "insert into detail_db (policyID, statecode, county, eq_site_limit, hu_site_limit, fl_site_limit, fr_site_limit, tiv_2011, tiv_2012, eq_site_deductible, hu_site_deductible, fl_site_deductible, fr_site_deductible, point_latitude, point_longitude, line, construction, point_granularity) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)"
    
    var sStmt1 string = "truncate table detail_db"

//===================================DATABASE CONNECTION============================

	db, err := sql.Open("postgres", "user=postgres password='root' dbname='sample' sslmode=disable")
	if err != nil {
		log.Fatal(err)
		
		fmt.Println("there are errors")

	}else {

		fmt.Println("================database db_Connection connected========")
	}

	    
//==================================file format open && Read=============================

    file_cs, err := os.Open("sample.csv")
  
    reader := csv.NewReader(file_cs)
    //reader.FieldsPerRecord = -1 

    rawcsvdata,_ :=reader.ReadAll()
   // rawcsvdata.Comma = ';'
//=====================================truncate the file===========================
   db.Exec(sStmt1)

 for _, each := range rawcsvdata {
 
   fmt.Println(each[0],each[1],each[2],each[3],each[4],each[5],each[6],each[7],each[8],each[9],each[10],each[11],each[12],each[13],each[14],each[15],each[16],each[17])
   

   db.Exec(sStmt, each[0],each[1],each[2],each[3],each[4],each[5],each[6],each[7],each[8],each[9],each[10],each[11],each[12],each[13],each[14],each[15],each[16],each[17])
 
    if err != nil {
 
    log.Fatal(err)

  }


}
   fmt.Println("================completed==================") 


//===================================table row reading and write===============================
   var id int
   fmt.Println("Enter search id:")
   fmt.Scanf("%d",&id)
   fmt.Println("hi",id)
 
    rows, err2 := db.Query("SELECT * FROM detail_db WHERE policyid = $1",id)
	if err2 != nil {
	    log.Fatal(err2)
	}
    var d declare
    var d_store []declare
   	
   	for rows.Next() {
 	
    if err2 := rows.Scan(&d.policyID, &d.statecode, &d.county, &d.eq_site_limit, &d.hu_site_limit, &d.fl_site_limit, &d.fr_site_limit, &d.tiv_2011, &d.tiv_2012, &d.eq_site_deductible, &d.hu_site_deductible, &d.fl_site_deductible, &d.fr_site_deductible, &d.point_latitude, &d.point_longitude, &d.line, &d.construction, &d.point_granularity); err2 != nil {
	        log.Fatal(err2)
	    }
    fmt.Println(d)

    d_store = append(d_store,d)    
	}
//========================================json=======================================
	d_json := declare{
		policyID: d.policyID, 
		statecode: d.statecode, 
		county: d.county, 
		eq_site_limit: d.eq_site_limit, 
		hu_site_limit: d.hu_site_limit, 
		fl_site_limit: d.fl_site_limit, 
		fr_site_limit: d.fr_site_limit, 
		tiv_2011: d.tiv_2011, 
		tiv_2012: d.tiv_2012, 
		eq_site_deductible: d.eq_site_deductible, 
		hu_site_deductible: d.hu_site_deductible, 
		fl_site_deductible: d.fl_site_deductible, 
		fr_site_deductible: d.fr_site_deductible, 
		point_latitude: d.point_latitude, 
		point_longitude: d.point_longitude, 
		line: d.line, 
		construction: d.construction, 
		point_granularity: d.point_granularity, 
	}
    fmt.Println(d_store)
    //fmt.Println(d_json)

    b, err3 := json.Marshal(d_json)
	if err3 != nil {
		fmt.Println("error:", err3)
	}
	//fmt.Println(b)
	os.Stdout.Write(b)
	fmt.Println("")
	if err := rows.Err(); err != nil {
	    log.Fatal(err)
	}


	defer db.Close()

}
