package settings

import "database/sql"

// DB connection, filled in "init" func
var DB *sql.DB

// Count of users in one part
var USER_ROW_COUNT = 15
