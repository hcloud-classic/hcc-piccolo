package pbtomodel

import (
	"github.com/hcloud-classic/pb"
	"hcc/piccolo/model"
)

// PbTaskToModelTask : Change task of proto type to model
func PbTaskToModelTask(task *pb.Task) *model.Task {
	var children []model.Task
	var threads []model.Task

	var emptyTaskList []model.Task

	modelProcess := &model.Task{
		CMD:        task.GetCMD(),
		State:      task.GetState(),
		PID:        int(task.GetPID()),
		PPID:       int(task.GetPPID()),
		PGID:       int(task.GetPGID()),
		SID:        int(task.GetSID()),
		Priority:   int(task.GetPriority()),
		Nice:       int(task.GetNice()),
		NumThreads: int(task.GetNumThreads()),
		StartTime:  task.GetStartTime(),
		Children:   nil,
		Threads:    nil,
		CPUUsage:   task.GetCPUUsage(),
		MemUsage:   task.GetMemUsage(),
		EPMType:    task.GetEPMType(),
		EPMSource:  int(task.GetEPMSource()),
		EPMTarget:  int(task.GetEPMTarget()),
	}

	for _, child := range task.GetChildren() {
		children = append(children, model.Task{
			CMD:        child.GetCMD(),
			State:      child.GetState(),
			PID:        int(child.GetPID()),
			PPID:       int(child.GetPPID()),
			PGID:       int(child.GetPGID()),
			SID:        int(child.GetSID()),
			Priority:   int(child.GetPriority()),
			Nice:       int(child.GetNice()),
			NumThreads: int(child.GetNumThreads()),
			StartTime:  child.GetStartTime(),
			Children:   emptyTaskList,
			Threads:    emptyTaskList,
			CPUUsage:   child.GetCPUUsage(),
			MemUsage:   child.GetMemUsage(),
			EPMType:    child.GetEPMType(),
			EPMSource:  int(child.GetEPMSource()),
			EPMTarget:  int(child.GetEPMTarget()),
		})
	}
	modelProcess.Children = children

	for _, thread := range task.GetThreads() {
		threads = append(threads, model.Task{
			CMD:        thread.GetCMD(),
			State:      thread.GetState(),
			PID:        int(thread.GetPID()),
			PPID:       int(thread.GetPPID()),
			PGID:       int(thread.GetPGID()),
			SID:        int(thread.GetSID()),
			Priority:   int(thread.GetPriority()),
			Nice:       int(thread.GetNice()),
			NumThreads: int(thread.GetNumThreads()),
			StartTime:  thread.GetStartTime(),
			Children:   emptyTaskList,
			Threads:    emptyTaskList,
			CPUUsage:   thread.GetCPUUsage(),
			MemUsage:   thread.GetMemUsage(),
			EPMType:    thread.GetEPMType(),
			EPMSource:  int(thread.GetEPMSource()),
			EPMTarget:  int(thread.GetEPMTarget()),
		})
	}
	modelProcess.Threads = threads

	return modelProcess
}
