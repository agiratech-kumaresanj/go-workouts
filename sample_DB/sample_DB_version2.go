package main

import (
    "encoding/csv"
    "fmt"
    "database/sql"
    _"github.com/lib/pq"
    "log"
     "os"
    )

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

   	for rows.Next() {
 	var policyID int
 	var statecode, county, eq_site_limit, hu_site_limit, fl_site_limit, fr_site_limit, tiv_2011, tiv_2012, eq_site_deductible, hu_site_deductible, fl_site_deductible, fr_site_deductible, point_latitude, point_longitude, line, construction string
	var point_granularity int 

    if err2 := rows.Scan(&policyID, &statecode, &county, &eq_site_limit, &hu_site_limit, &fl_site_limit, &fr_site_limit, &tiv_2011, &tiv_2012, &eq_site_deductible, &hu_site_deductible, &fl_site_deductible, &fr_site_deductible, &point_latitude, &point_longitude, &line, &construction, &point_granularity); err2 != nil {
	        log.Fatal(err2)
	    }

    fmt.Println(policyID, statecode, county, eq_site_limit, hu_site_limit, fl_site_limit, fr_site_limit, tiv_2011, tiv_2012, eq_site_deductible, hu_site_deductible, fl_site_deductible, fr_site_deductible, point_latitude, point_longitude, line, construction, point_granularity)
	}
	
	if err := rows.Err(); err != nil {
	    log.Fatal(err)
	}


	defer db.Close()

}
