package template

import _ "embed"

//go:embed admin.sql
var AdminSqlTemplate string

//go:embed robot.sql
var RobotSqlTemplate string
