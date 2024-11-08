package service

import (
	"mini/entity"
	"mini/repository"
)

type LoanService struct {
    loanRepo repository.LoanRepository
}

func NewLoanService(loanRepo repository.LoanRepository) *LoanService {
    return &LoanService{loanRepo: loanRepo}
}

func (s *LoanService) GetAllLoans() ([]entity.Loan, error) {
    return s.loanRepo.GetAllLoans()
}

func (s *LoanService) GetLoanByID(id uint) (entity.Loan, error) {
    return s.loanRepo.GetLoanByID(id)
}
