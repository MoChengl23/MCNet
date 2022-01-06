package timer


type Callback func(a int)
type Timer struct{
	
}

func(timer *Timer)  AddTask(delay int){

	go timer.UpdateTask()
}

func(timer *Timer) UpdateTask(){



}