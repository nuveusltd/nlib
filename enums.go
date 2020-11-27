package nlib


type DBType int

const (
	mysql = iota
	postgres
	sqlite
	sqlserver
)
func (d DBType) String() string {
  return [...]string{"Mysql", "PostgreSQL", "SQLite", "Sql Server"}[d]
}



type JoinMethodType int

const (
	open = iota
	approve
	close
)

func (d JoinMethodType) String() string {
  return [...]string{"Everyone can join without approval", "Needs Admins Approve", "Closed to Join"}[d]
}

type JoinSourceType int

const (
	db  = iota
	ldap
)

func (d JoinSourceType) String() string {
  return [...]string{"Database", "Ldap"}[d]
}

type ConnectionType int

const (
	unsecure = iota
	secure
)

func (c ConnectionType) String() string {
	return [...]string{"UnSecure","Secure"}[c]
}
