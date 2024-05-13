
package pkg

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

func NewConn() (*pgxpool.Pool, error) {

	//new conenction
	
	fmt.Println("success connection")

	return connPool, nil
}
