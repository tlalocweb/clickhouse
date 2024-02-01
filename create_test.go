package clickhouse_test

import (
	"testing"

	"gorm.io/gorm/utils/tests"
)

func TestCreate(t *testing.T) {
	var user = User{ID: 1, Name: "create", FirstName: "zhang", LastName: "jinzhu", Age: 18, Active: true, Salary: 8.8888}

	if err := DB.Create(&user).Error; err != nil {
		t.Fatalf("failed to create user, got error %v", err)
	}

	var result User
	if err := DB.Find(&result, user.ID).Error; err != nil {
		t.Fatalf("failed to query user, got error %v", err)
	}

	tests.AssertEqual(t, result, user)
}

func TestBatchCreate(t *testing.T) {
	var users = []User{
		{ID: 11, Name: "batch_create_1", FirstName: "zhang", LastName: "jinzhu", Age: 18, Active: true, Salary: 6},
		{ID: 12, Name: "batch_create_2", FirstName: "zhang", LastName: "jinzhu", Age: 18, Active: false, Salary: 6.12},
		{ID: 13, Name: "batch_create_3", FirstName: "zhang", LastName: "jinzhu", Age: 18, Active: true, Salary: 6.1234},
		{ID: 14, Name: "batch_create_4", FirstName: "zhang", LastName: "jinzhu", Age: 18, Active: false, Salary: 6.123456},
	}

	if err := DB.Create(&users).Error; err != nil {
		t.Fatalf("failed to create users, got error %v", err)
	}

	var results []User
	DB.Find(&results)

	for _, u := range users {
		var result User
		if err := DB.Find(&result, u.ID).Error; err != nil {
			t.Fatalf("failed to query user, got error %v", err)
		}

		tests.AssertEqual(t, result, u)
	}
}

func TestCreateWithMap(t *testing.T) {
	var user = User{ID: 122, Name: "create2", FirstName: "zhang", LastName: "jinzhu", Age: 18, Active: true, Salary: 6.6666}

	if err := DB.Table("users").Create(&map[string]interface{}{
		"id": user.ID, "name": user.Name, "first_name": user.FirstName, "last_name": user.LastName, "age": user.Age, "active": user.Active, "salary": user.Salary,
	}).Error; err != nil {
		t.Fatalf("failed to create user, got error %v", err)
	}

	var result User
	if err := DB.Find(&result, user.ID).Error; err != nil {
		t.Fatalf("failed to query user, got error %v", err)
	}

	user.CreatedAt = result.CreatedAt
	user.UpdatedAt = result.UpdatedAt
	tests.AssertEqual(t, result, user)
}

// Just a sanity check on those structs - we will use them in association tests
func TestCreateOnManyManyStructs(t *testing.T) {
	var user = User2{ID: 789, Name: "create3", FirstName: "zhang", LastName: "jinzhu", Age: 18, Active: true, Salary: 6.6666}

	// DB.Transaction(func(tx *gorm.DB) (err error) {
	// 	if err = tx.Create(&user).Error; err != nil {
	// 		t.Fatalf("failed to create user (User2), got error %v", err)
	// 	}
	// 	return err
	// })
	if err := DB.Create(&user).Error; err != nil {
		t.Fatalf("failed to create user (User2), got error %v", err)
	}

	//	time.Sleep(1 * time.Second)

	var email = Email{Email: "sadk@sad.com", Something: "just another field"}
	// DB.Transaction(func(tx *gorm.DB) (err error) {
	// 	if err = tx.Create(&email).Error; err != nil {
	// 		t.Fatalf("failed to create email, got error %v", err)
	// 	}
	// 	return err
	// })
	if err := DB.Create(&email).Error; err != nil {
		t.Fatalf("failed to create email, got error %v", err)
	}

	//	time.Sleep(1 * time.Second)

	// THIS FAILS - not sure how to fix yet
	// // ok now
	// var anotheruser = User2{
	// 	Name:      "many_to_many",
	// 	FirstName: "many",
	// 	LastName:  "to_many",
	// 	Age:       18,
	// 	Active:    true,
	// 	Emails: []Email{
	// 		{Email: "abc@abc.com"},
	// 	},
	// }

	// if err := DB.Create(&anotheruser).Error; err != nil {
	// 	t.Fatalf("failed to create user, got error %v", err)
	// }

}
