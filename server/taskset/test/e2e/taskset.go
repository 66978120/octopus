package e2e

import (
	typeJob "volcano.sh/volcano/pkg/apis/batch/v1alpha1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TaskSet Controller E2E Test", func() {
	It("Successed TaskSet Test", func() {

		context := initTestContext(TestNameSpace)

		defer cleanupTestContext(context)

		taskset := buildSuccessedTaskSet("successedtaskset")

		created, err := context.tsclient.OctopusV1alpha1().TaskSets(TestNameSpace).Create(taskset)

		Expect(err).NotTo(HaveOccurred())

		defer saveTaskSet(context, taskset, OutputFilePath+taskset.Name+".yaml")

		var ts *typeJob.Job

		err, ts = waitJobStates(context, created, []TaskSetPhase{TaskSetAttemptPreparing, TaskSetAttemptRunning, TaskSetCompleted}, twoMinute)

		Expect(err).NotTo(HaveOccurred())

		Expect(ts.Status.State).Should(Equal(typeJob.Succeeded))
		Expect(ts.Status.Phase).Should(Equal(typeJob.Completed))
	})

	It("Failed TaskSet Test", func() {
		context := initTestContext(TestNameSpace)
		defer cleanupTestContext(context)

		taskset := buildFailedTaskSet("failedtaskset")

		created, err := context.tsclient.OctopusV1alpha1().TaskSets(TestNameSpace).Create(taskset)
		Expect(err).NotTo(HaveOccurred())
		defer saveTaskSet(context, taskset, OutputFilePath+taskset.Name+".yaml")
		var ts *typeJob.Job
		err, ts = waitJobStates(context, created, []TaskSetPhase{TaskSetAttemptPreparing, TaskSetAttemptRunning, TaskSetCompleted}, twoMinute)
		Expect(err).NotTo(HaveOccurred())
		Expect(ts.Status.State).Should(Equal(typeTaskSetController.FAILED))
		Expect(ts.Status.Phase).Should(Equal(typeTaskSetController.TaskSetCompleted.GetName()))
	})

	It("Random Failed TaskSet Test", func() {
		context := initTestContext(TestNameSpace)
		defer cleanupTestContext(context)
		// case2: random failed job
		taskset := buildRandomFailedTaskSet("randomfailedtaskset")

		created, err := context.tsclient.OctopusV1alpha1().TaskSets(TestNameSpace).Create(taskset)
		Expect(err).NotTo(HaveOccurred())
		defer saveTaskSet(context, taskset, OutputFilePath+taskset.Name+".yaml")
		var ts *typeJob.Job
		err, ts = waitJobStates(context, created, []TaskSetPhase{TaskSetAttemptPreparing, TaskSetAttemptRunning, TaskSetCompleted}, twoMinute)
		Expect(err).NotTo(HaveOccurred())
		Expect(ts.Status.Phase).Should(Equal(typeTaskSetController.TaskSetCompleted.GetName()))

	})

	It("Two Role Retry TaskSet Test", func() {
		context := initTestContext(TestNameSpace)
		defer cleanupTestContext(context)

		taskset := emptyTaskSet("tworolefailedretrytaskset")

		created, err := context.tsclient.OctopusV1alpha1().TaskSets(TestNameSpace).Create(taskset)
		Expect(err).NotTo(HaveOccurred())
		defer saveTaskSet(context, taskset, OutputFilePath+taskset.Name+".yaml")
		var ts *typeJob.Job
		err = waitJobStates(context, created, []TaskSetPhase{TaskSetCompleted}, twoMinute)
		Expect(err).NotTo(HaveOccurred())

	})

})