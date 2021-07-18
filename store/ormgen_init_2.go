// DON'T EDIT *** generated by ormgen *** DON'T EDIT
package store





import "fmt"
const (  
	UserTypeTable = `
	create table if not exists %[1]s.registration_user_type(
	  ut_id int generated by default as identity primary key,
	  name varchar(64) unique not null
	);`
 
	UserTable = `
	create table if not exists %[1]s.registration_user(
	  u_id int generated by default as identity primary key,
	  updated timestamptz not null default current_timestamp,
	  name varchar(64) not null,
	  firstname varchar(64) not null,
	  registration_user_type_ut_id int not null,
	  registration_user_type_name varchar(64) not null
	);`

)

var tables = []string{ 
	UserTypeTable,
	UserTable,
}

var indexes = []string{ 
}

var schema = "config"

//TODO enable whitelisting
func New(st []string) []string {
	tstrings := make([]string, len(tables)+len(indexes))
	for i, table := range tables {
		tstrings[i] = fmt.Sprintf(table, schema)
	}
	for j, index := range indexes {
		tstrings[len(tables) + j] = index
	}
	return tstrings
}

var Setup []string

func InitSchema() []string {
	if schema != "" {
		for idx, query := range Setup {
			Setup[idx] = fmt.Sprintf(query, schema)
		}
	}
        return Setup
}
