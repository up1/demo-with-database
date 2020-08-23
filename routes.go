package main

const BudgetPath = "/budgets"

func (s *server) routes() {
	s.mux.HandleFunc(BudgetPath, s.getBudgets).Methods("GET")
}
