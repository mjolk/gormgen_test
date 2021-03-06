// DON'T EDIT *** generated by ormgen *** DON'T EDIT //
package store




import (
	"database/sql" 
	"github.com/lib/pq"
	"github.com/gofrs/uuid"
	"fmt"
)












func UserFrom_lvl1() string { 
	return `FROM registration_user regu`
}

func UserSelect_lvl1() string { 
	return `SELECT regu.u_id, regu.updated, regu.name, regu.firstname, regu.registration_user_type_ut_id, regu.registration_user_type_name`
}

func LoadUser_lvl1(rs *sql.Rows) (*User, error) {
	result, err := LoadUsers_lvl1(rs)
	if err != nil {
		return nil, err
	}
	if len(result) > 1 {
		return nil, fmt.Errorf("More than one result, %d results found", len(result)) 
	} else if len(result) == 0 {
		return nil, sql.ErrNoRows
	}
	return result[0], nil
}

func LoadUsers_lvl1(rs *sql.Rows) ([]*User, error) {
	var err error
	defer rs.Close()	
	result := make([]*User, 0)
	var recv0 sql.NullInt64
	var recv1 pq.NullTime
	var recv2 sql.NullString
	var recv3 sql.NullString
	var recv4 sql.NullInt64
	var recv5 sql.NullString
	dest := []interface{}{ 
		&recv0,
		&recv1,
		&recv2,
		&recv3,
		&recv4,
		&recv5,
	}
	for rs.Next() {
		if err = rs.Scan(dest...); err != nil {
			return nil, err
		}
		s := newProxyUser()
		s.UserType = new(UserType)
		s.ID = checkSqlInt64Value(dest[0])
		s.Updated = checkSqlTimeValue(dest[1])
		s.Name = checkSqlStringValue(dest[2])
		s.FirstName = checkSqlStringValue(dest[3])
		s.UserType.ID = checkSqlInt64Value(dest[4])
		s.UserType.Name = checkSqlStringValue(dest[5])
		entity := s.User
		var newID bool = true
		for _, ent := range result {
			if ent.ID == s.ID {
				entity = ent
				newID = false
			}
		}
		
		

				
				

		
                if newID {
			result = append(result, entity)
		}
	}
	if cErr := rs.Close(); cErr != nil {
		return nil, cErr
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return result, nil
}


