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
  //db_Connection()
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

/*func db_Search(){

}*/

func db_Connection(){
    

    var sStmt string = "insert into detail_db (policyID, statecode, county, eq_site_limit, hu_site_limit, fl_site_limit, fr_site_limit, tiv_2011, tiv_2012, eq_site_deductible, hu_site_deductible, fl_site_deductible, fr_site_deductible, point_latitude, point_longitude, line, construction, point_granularity) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)"

	db, err := sql.Open("postgres", "user=postgres password='root' dbname='sample' sslmode=disable")
	if err != nil {
		log.Fatal(err)
		
		fmt.Println("there are errors")

	}else {

		fmt.Println("================database db_Connection connected========")
	}

	/*var id int
	var name, position string
   fmt.Println("Enter 1.id, 2.name, 3. position")
   fmt.Scanf("%d %s %s", &id, &name, &position)
	_, err1 := db.Exec(sStmt,id,name,position)
    
    if err1 != nil {
    fmt.Println(err1)
    }*/
	    
  //file format open
    file_cs, err := os.Open("sample.csv")
  
    reader := csv.NewReader(file_cs)
    //reader.FieldsPerRecord = -1 

    rawcsvdata,_ :=reader.ReadAll()
   // rawcsvdata.Comma = ';'


 for _, each := range rawcsvdata {
 
   fmt.Println(each[0],each[1],each[2],each[3],each[4],each[5],each[6],each[7],each[8],each[9],each[10],each[11],each[12],each[13],each[14],each[15],each[16],each[17])
 
   db.Exec(sStmt, each[0],each[1],each[2],each[3],each[4],each[5],each[6],each[7],each[8],each[9],each[10],each[11],each[12],each[13],each[14],each[15],each[16],each[17])
  //_, err := db.Exec(sStmt, each[0],each[1],each[2],each[3],each[4],each[5],each[6],each[7],each[8],each[9],each[10],each[11],each[12],each[13],each[14],each[15],each[16],each[17])
 
    if err != nil {
 
    log.Fatal(err)

  }


}
   fmt.Println("================completed==================") 

 /*
 var sStmt1 string = "select * from csv where policyID = $1"
 db.Query(sStmt1)
// if sStmt1 == id {
 
 for i := 0; i < 18;i++{
   
   fmt.Println(db.Exec(each[i]))
 }*/
 /*var policyID int
 var statecode, county, eq_site_limit, hu_site_limit, fl_site_limit, fr_site_limit, tiv_2011, tiv_2012, eq_site_deductible, hu_site_deductible, fl_site_deductible, fr_site_deductible, point_latitude, point_longitude, line, construction string
 var point_granularity int
 */
/*var id int
err = db.QueryRow("select name from users where policyID = ?", 1).Scan(&id)
if err != nil {
	log.Fatal(err)
}
fmt.Println(id)
*/
 /*rows, err2 := db.QueryRow("SELECT * FROM details WHERE policyID = %d",id)
	if err2 != nil {
	    log.Fatal(err2)
	}
*/
	/*for rows.Next() {
	    //var name string
	    var no int
        var name1, pos string 
	    if err2 := rows.Scan(&policyID, &statecode, &county, &eq_site_limit, &hu_site_limit, &fl_site_limit, &fr_site_limit, &tiv_2011, &tiv_2012, &eq_site_deductible, &hu_site_deductible, &fl_site_deductible, &fr_site_deductible, &point_latitude, &point_longitude, &line, &construction, &point_granularity); err2 != nil {
	        log.Fatal(err2)
	    }
	    fmt.Printf("%d %s is %s\n",no, name1, pos)
	}
	
	if err := rows.Err(); err != nil {
	    log.Fatal(err)
	}
*/
/* } else{
 	fmt.Println("Its not in the list")
 }
*/   /*rows, err2 := db.Query("SELECT * FROM detail")

    if err1 != nil {
    fmt.Println(err2)
    }

	fmt.Println(rows)*/

//table row reading and write
   var id int
   fmt.Println("Enter search id:")
   fmt.Scanf("%d",&id)
   fmt.Println("hi",id)
 
    rows, err2 := db.Query("SELECT * FROM detail_db WHERE policyid = $1",id)
	if err2 != nil {
	    log.Fatal(err2)
	}
  // fmt.Println(rows)

   	for rows.Next() {
	    //var name string
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
