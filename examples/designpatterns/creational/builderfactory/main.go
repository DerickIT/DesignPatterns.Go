package main

import "fmt"

type User struct {
	Name     string
	Email    string
	Age      int
	IsActive bool
}

type UserBuilder struct {
	user User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{user: User{}}
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) WithAge(age int) *UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilder) WithIsActive(isActive bool) *UserBuilder {
	b.user.IsActive = isActive
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}

func main() {
	user := NewUserBuilder().
		WithName("John Doe").
		WithEmail("john.doe@example.com").
		WithAge(30).
		WithIsActive(true).
		Build()
	user1 := NewUserBuilder().
		WithName("alan").
		WithEmail("3333@lld.com").
		WithAge(23).
		WithIsActive(true).
		Build()
	fmt.Println(user)
	fmt.Println(user1)
}
