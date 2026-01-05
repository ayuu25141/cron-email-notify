package dbconnection
import(
	   "context"
    "fmt"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"
)
var Pool *pgxpool.Pool
func Connectiontopostgres() error  {
	dsn := os.Getenv("Dburl")
	pool,err :=  pgxpool.New(context.Background(),dsn)
	if err != nil {
return fmt.Errorf("eror to connect db") 
	}
	Pool = pool
	return  nil

}