// DON'T EDIT *** generated by ormgen *** DON'T EDIT //
package store





import (
	"context"
	"fmt"
	"log"
	"os"
)

const (
	dbError = "Can't initialize database. error: %s\n"
	SCHEMA  = "bp_registration"
)

var cfg dbConfig

func createTables(ctx SQLContext, tables []string) error {
	for _, query := range New(tables) {
		if _, err := ctx.Exec(query); err != nil {
			log.Printf("error @ query : \n %s \n", query)
			return err
		}
	}
	for _, init := range InitSchema() {
		if _, err := ctx.Exec(init); err != nil {
			log.Printf("error @ query : \n %s \n", init)
			return err
		}
	}

	return nil
}

func Init(ctx context.Context) {
	cfg = loadConfig()	
	ok := false
	val := "no"
	val, ok = os.LookupEnv("INITDB")
	log.Printf("RECREATE DB: %s", val)
	if ok && val == "yes" {
		err := Transact(
			NewSQLContext(ctx),
			func(sctx SQLContext) error {
				if err := initDB(sctx); err != nil {
					return err
				}
				return nil
			},
		)
		if err != nil {
			log.Fatalf(dbError, err)
		}

	}
	

}

func createLtree(ctx SQLContext) error {
	if _, err := ctx.Exec(
		fmt.Sprintf(
			"create extension if not exists ltree with schema %s",
			cfg.Schema,
		),
	); err != nil {
		return err
	}
	return nil
}

func initDB(ctx SQLContext) error {
	var err error

	if err = dropSchema(ctx); err != nil {
		return err
	}

	if err = createSchema(ctx); err != nil {
		return err
	}

	if err = setSessionPath(ctx); err != nil {
		return err
	}

	if err = setUserSearchPath(ctx); err != nil {
		return err
	}

	/*	if err = createLtree(ctx); err != nil {
			return err
		}
	*/

	if err = createTables(ctx, cfg.Schema, UserSchema); err != nil {
		return err
	}
	return nil
}


func setSessionPath(ctx SQLContext) error {
	_, err := ctx.Exec(
		fmt.Sprintf("set search_path = %s", cfg.Schema),
	)
	if err != nil {
		return err
	}
	return nil
}

func setUserSearchPath(ctx SQLContext) error {
	_, err := ctx.Exec(
		fmt.Sprintf(
			"alter role %s set search_path = %s",
			cfg.Db,
			cfg.Schema,
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func dropSchema(ctx SQLContext) error {
	_, err := ctx.Exec(
		fmt.Sprintf(
			"drop schema if exists %s cascade",
			cfg.Schema,
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func createSchema(ctx SQLContext) error {
	_, err := ctx.Exec(
		fmt.Sprintf(
			"create schema if not exists %s",
			cfg.Schema,
		),
	)
	if err != nil {
		return err
	}
	return nil
}
