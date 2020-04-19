package sherlockanalyser

import (
	"testing"
)

/*
containsElement will check whether a given slice has an element or not.
*/
func containsElement(s []uint64, id uint64) bool {
	for _, i := range s {
		if i == id {
			return true
		}
	}
	return false
}

//=======================================================================
// TESTS
//=======================================================================

/*
TestQueueHasIDAlreadyInUse tests whether a id is already in use.
*/
func TestQueueHasIDAlreadyInUse(t *testing.T) {
	queue := NewAnalyserQueue()
	task := &analyserTaskRequest{}
	id := queue.AppendQueue(task)
	if id == 0 {
		t.Fatal("got a 0 id from the AppendQueue-Method")
	} else if contains := queue.ContainsTaskID(id); !contains {
		t.Fatalf("id is not in the queue altought it should be. ID: %d, Contains: %t", id, contains)
	} else {
		t.Log("successfully found ID in task queue")
	}

}

/*
TestQueueContainsTaskIDFailes tests trys
*/
func TestQueueContainsTaskIDFailes(t *testing.T) {
	queue := NewAnalyserQueue()
	task := &analyserTaskRequest{}
	id := queue.AppendQueue(task)
	if id == 0 {
		t.Fatal("got a 0 id from the AppendQueue-Method")
	} else if contains := queue.ContainsTaskID(id + 1); contains {
		t.Fatalf("id is not in the queue altought it should be. ID: %d, Contains: %t", id, contains)
	} else {
		t.Log("successfully found ID in taskqueue")
	}

}

/*
TestQueueAppendFailesBecauseOfNilPointer will try to make the Apppend Method fail because of a nullpointer.
*/
func TestQueueAppendFailesBecauseOfNilPointer(t *testing.T) {
	queue := NewAnalyserQueue()
	id := queue.AppendQueue(nil)
	if id == 0 {
		t.Log("got a zero returncode as expected", 0, id)
	} else {
		t.Fatalf("got a non-zero returncode. Expected: %d, Got: %d", 0, id)
	}

}

/*
TestQueueRemoveTask will append a task to the queue and remove it.
*/
func TestQueueRemoveTask(t *testing.T) {
	queue := NewAnalyserQueue()
	task := &analyserTaskRequest{}
	id := queue.AppendQueue(task)
	if id == 0 {
		t.Fatal("go a 0 id from the AppendQueue-Method")
	} else if contains := queue.ContainsTaskID(id); !contains {
		t.Fatalf("id is not in the queue altought it should be. ID: %d, Contains: %t", id, contains)
	} else {
		t.Log("successfully added task to the taskqueue")
	}
	if works := queue.RemoveFromQueue(id); !works {
		t.Fatalf("couldnt remove task with valid id Expected: %t, Got: %t, ID: %d", true, works, id)
	} else {
		t.Log("successfully removed task")
	}
}

/*
TestQueueRemoveTaskFailedBecauseOfZeroID will fail because of a zero id.
*/
func TestQueueRemoveTaskFailedBecauseOfZeroID(t *testing.T) {
	queue := NewAnalyserQueue()
	task := &analyserTaskRequest{}
	id := queue.AppendQueue(task)
	if id == 0 {
		t.Fatal("go a 0 id from the AppendQueue-Method")
	} else if contains := queue.ContainsTaskID(id); !contains {
		t.Fatalf("id is not in the queue altought it should be. ID: %d, Contains: %t", id, contains)
	} else {
		t.Log("successfully added task to the taskqueue")
	}
	if works := queue.RemoveFromQueue(0); works {
		t.Fatalf("could remove task with unvalid id Expected: %t, Got: %t, ID: %d", false, works, id)
	} else {
		t.Log("successfully removed task")
	}
}

/*
TestGetQueueStatus will test the status of the queue and show its status.
*/
func TestGetQueueStatus(t *testing.T) {
	message := "States of the Task currently in the Queue. \nUndone: 1, Processing: 1, CrawlerError: 1, Saving: 1, SendToCrawler: 1, Finished: 1"
	queue := NewAnalyserQueue()
	task1 := &analyserTaskRequest{}
	task2 := &analyserTaskRequest{}
	task3 := &analyserTaskRequest{}
	task4 := &analyserTaskRequest{}
	task5 := &analyserTaskRequest{}
	task6 := &analyserTaskRequest{}

	task1.SetState(UNDONE)
	task2.SetState(PROCESSING)
	task3.SetState(CRAWLERERROR)
	task4.SetState(SAVING)
	task5.SetState(SENDTOCRAWLER)
	task6.SetState(FINISHED)

	id1 := queue.AppendQueue(task1)
	id2 := queue.AppendQueue(task2)
	id3 := queue.AppendQueue(task3)
	id4 := queue.AppendQueue(task4)
	id5 := queue.AppendQueue(task5)
	id6 := queue.AppendQueue(task6)

	if id1 == 0 || id2 == 0 || id3 == 0 || id4 == 0 || id5 == 0 || id6 == 0{
		t.Fatalf("Missmatching ids excpeted non-zero value but got id1: %d, id2: %d, id3: %d, id4: %d, id5: %d, id6: %d", id1, id2, id3, id4, id5, id6)
	} else {
		t.Log("successfully created different tasks with different states")
	}

	status := queue.GetStatusOfQueue()

	if status != message {
		t.Fatalf("messages does not match Expected: %s, Got: %s", message, status)
	} else {
		t.Log("message go as expected.")
	}
}

/*
TestGetAllTaskIds if GetAllTaskIds will return all ids.
*/
func TestGetAllTaskIds(t *testing.T) {
	queue := NewAnalyserQueue()
	task1 := &analyserTaskRequest{}
	task2 := &analyserTaskRequest{}
	id1 := queue.AppendQueue(task1)
	id2 := queue.AppendQueue(task2)

	if id1 == 0 || id2 == 0 {
		t.Fatalf("Missmatching ids excpeted non-zero value but got id1: %d, id2: %d", id1, id2)
	} else {
		t.Log("successfully creates 4 different tasks with different states")
	}

	ids := queue.getAllTaskIds()

	if length := len(ids); length != 2 {
		t.Fatalf("The length of the slice of ids does not be match the real one. Expected: %d, Got: %d", 2, length)
	} else if contains := containsElement(ids, id1); !contains {
		t.Fatalf("element in should be in map but wasnt there. Coulndt find: %d", id1)
	} else if contains := containsElement(ids, id2); !contains {
		t.Fatalf("element in should be in map but wasnt there. Coulndt find: %d", id2)
	} else {
		t.Log("successfully found all elements")
	}
}
