package main
import ("fmt"
       "log"
 //      "string"
)
func main(){
txn, err := db.Begin()
if err != nil {
	log.Fatal(err)
}

stmt, err := txn.Prepare(pq.CopyIn("users", "name", "age"))
if err != nil {
	log.Fatal(err)
}

for _, user := range users {
	_, err = stmt.Exec(user.Name, int64(user.Age))
	if err != nil {
		log.Fatal(err)
	}
}

_, err = stmt.Exec()
if err != nil {
	log.Fatal(err)
}

err = stmt.Close()
if err != nil {
	log.Fatal(err)
}

err = txn.Commit()
if err != nil {

	log.Fatal(err)
}

}