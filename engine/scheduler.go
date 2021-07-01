package engine

import (
	"chasqi/engine/agent"
	"chasqi/rules"
)

type Scheduler struct {
	navigationTree  rules.NavigationTree
	amountOfAgents  int
	agents          []agent.Agent
	completedAgents []agent.Agent
	failedAgents    []agent.Agent
}

func NewScheduler(tree *rules.NavigationTree) *Scheduler {
	s := Scheduler{}
	s.navigationTree = *tree
	s.amountOfAgents = tree.AmountOfAgents
	return &s
}

func (s *Scheduler) Start() {
	// TODO iterate through scheduler and start all agents
	go s.agents[0].Start()
}

/**
Creates new agents based on the configurations and initializes
them
*/
func (s *Scheduler) Schedule(
	debugChannel chan string,
) {
	agentList := make([]agent.Agent, s.navigationTree.AmountOfAgents)
	for i := 0; i < s.amountOfAgents; i++ {
		a := new(agent.Agent)
		a.Init(
			s.navigationTree,
			debugChannel,
			i,
		)
		agentList[i] = *a
	}
	s.agents = agentList
}
