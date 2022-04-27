package data_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
	"user/internal/biz"
	"user/internal/data"
)

var _ = Describe("User", func() {
	var ro biz.UserRepo
	var uD *biz.User
	BeforeEach(func() {
		ro = data.NewUserRepo(Db, nil)
		uD = &biz.User{
			ID:       1,
			Mobile:   "13803881388",
			Password: "admin123456",
			NickName: "aliliin",
			Role:     1,
			Birthday: nil,
		}
	})

	It("CreateUser", func() {
		u, err := ro.CreateUser(ctx, uD)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(u.Mobile).Should(Equal("13803881388"))
	})
	It("ListUser", func() {
		user, total, err := ro.ListUser(ctx, 1, 10)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(user).ShouldNot(BeEmpty())
		Ω(total).Should(Equal(1))
		Ω(len(user)).Should(Equal(1))
		Ω(user[0].Mobile).Should(Equal("13803881388"))
	})
	It("UpdateUser", func() {
		birthDay := time.Unix(int64(693646426), 0)
		uD.NickName = "xxx"
		uD.Birthday = &birthDay
		uD.Gender = "male"
		user, err := ro.UpdateUser(ctx, uD)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(user).Should(BeTrue())
	})
	It("CheckPassword", func() {
		p1 := "admin"
		encryptedPassword := "$pbkdf2-sha512$5p7doUNIS9I5mvhA$b18171ff58b04c02ed70ea4f39bda036029c107294bce83301a02fb53a1bcae0"
		password, err := ro.CheckPassword(ctx, p1, encryptedPassword)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(password).Should(BeTrue())

		encryptedPassword1 := "$pbkdf2-sha512$5p7doUNIS9I5mvhA$b18171ff58b04c02ed70ea4f39bda036029c107294bce83301a02fb53a1bcae"
		password, err = ro.CheckPassword(ctx, p1, encryptedPassword1)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(password).Should(BeFalse())
	})
	It("GetUserByID", func() {
		user, err := ro.GetUserById(ctx, uD.ID)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(user.ID).Should(Equal(uD.ID))
	})
})
