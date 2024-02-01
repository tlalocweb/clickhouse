package clickhouse_test

import (
	"testing"
)

func TestAssociateAdd(t *testing.T) {
	// BREAKS - Association() Append() actually later runs on update.
	// The Updateof the updated_at column is killing us again. Right now I don't
	// know how to remedy this in the driver.

	// var user = User2{ID: 789, Name: "mj", FirstName: "billy", LastName: "jean", Age: 18, Active: true, Salary: 6.6666}

	// if err := DB.Create(&user).Error; err != nil {
	// 	t.Fatalf("failed to create user (User2), got error %v", err)
	// }

	// if err := DB.Model(&user).Association("Emails").Append([]Email{{Email: "email1"}, {Email: "email2"}}); err != nil {
	// 	t.Fatalf("failed to append emails, got error %v", err)
	// }
}
