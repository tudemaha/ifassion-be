package dto

type DatabaseInsert struct {
	True  []string
	False []string
}

type ResultData struct {
	Time     string
	Database DatabaseInsert
	Status   bool
}
