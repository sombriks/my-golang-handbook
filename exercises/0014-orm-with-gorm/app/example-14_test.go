package app

import (
	"log"
	"os"
	"testing"
)

var example14 *Example14

var testId uint

func TestMain(m *testing.M) {
	log.Println("setup tests")
	example14 = NewExample14()
	testId, _ = example14.Insert("Mr Test", "Test Road 7")
	result := m.Run()
	log.Println("teardown tests")
	example14.database.Close()
	os.Exit(result)
}

func TestExample14_ListByAddress(t *testing.T) {
	result, err := example14.ListByAddress("ROAD")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil || len(*result) == 0 {
		t.Fail()
	}
}

func TestExample14_ListByName(t *testing.T) {
	result, err := example14.ListByName("mr")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil || len(*result) == 0 {
		t.Fail()
	}

}

func TestExample14_Find(t *testing.T) {
	result, err := example14.Find(testId)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fail()
	}
}

func TestExample14_Insert(t *testing.T) {
	result, err := example14.Insert("Robert", "Red Avenue 10")
	if err != nil {
		t.Fatal(err)
	}
	if result == 0 {
		t.Fail()
	}
}

func TestExample14_Update(t *testing.T) {
	result, err := example14.Find(testId)
	if err != nil {
		t.Fatal(err)
	}
	result.Name = "Javier"
	err = example14.Update(result)
	if err != nil {
		t.Fatal(err)
	}
	check, err := example14.ListByName("javier")
	if result == nil || len(*check) == 0 {
		t.Fail()
	}
}

func TestExample14_Delete(t *testing.T) {
	err := example14.Delete(testId)
	if err != nil {
		t.Fatal(err)
	}
	check, err := example14.Find(testId)
	if err == nil {
		t.Fail()
	}
	if check.ID != 0 {
		t.Fail()
	}
}
