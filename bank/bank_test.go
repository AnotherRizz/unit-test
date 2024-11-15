package bank

import (
    "testing"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func TestAccount(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Account Suite")
}

var _ = Describe("Account", func() {
    var account Account

    BeforeEach(func() {
        account = Account{}
    })

    Describe("Deposit", func() {
        It("should increase the balance with a valid amount", func() {
            Expect(account.Deposit(100)).To(Succeed())
            Expect(account.GetBalance()).To(Equal(100.0))
        })

        It("should return an error for zero or negative deposit amount", func() {
            Expect(account.Deposit(0)).To(MatchError("deposit amount must be greater than zero"))
            Expect(account.Deposit(-50)).To(MatchError("deposit amount must be greater than zero"))
        })
    })

    Describe("Withdraw", func() {
        BeforeEach(func() {
            account.Deposit(100)
        })

        It("should decrease the balance with a valid amount", func() {
            Expect(account.Withdraw(50)).To(Succeed())
            Expect(account.GetBalance()).To(Equal(50.0))
        })

        It("should return an error for zero or negative withdraw amount", func() {
            Expect(account.Withdraw(0)).To(MatchError("withdrawal amount must be greater than zero"))
            Expect(account.Withdraw(-50)).To(MatchError("withdrawal amount must be greater than zero"))
        })

        It("should return an error if withdrawal amount is greater than balance", func() {
            Expect(account.Withdraw(200)).To(MatchError("insufficient funds"))
        })
    })

    Describe("GetBalance", func() {
        It("should return the current balance", func() {
            account.Deposit(100)
            Expect(account.GetBalance()).To(Equal(100.0))
        })
    })
})
