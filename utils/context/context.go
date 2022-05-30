package context

import "yamlrest/services/database"

type AppContext struct {
	Database database.InMemDB
}
