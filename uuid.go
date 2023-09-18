package uuid

import (
    "database/sql/driver"
    googleUUID "github.com/google/uuid"
)

func New() UUID {
    u := googleUUID.New()

    return UUID{
        UUID: u,
        Str:  u.String(),
    }
}

func Parse(s string) (UUID, error) {
    u := UUID{}
    var err error

    u.UUID, err = googleUUID.Parse(s)

    if err == nil {
        u.Str = u.UUID.String()
    }

    return u, err
}

type UUID struct {
    UUID googleUUID.UUID

    Str string
}

func (u UUID) String() string {
    return u.Str
}

func (u UUID) IsNil() bool {
    return u.UUID == googleUUID.Nil
}

func (u UUID) URN() string {
    return u.UUID.URN()
}

func (u *UUID) Scan(src any) error {
    err := u.UUID.Scan(src)
    u.Str = u.UUID.String()
    return err
}

func (u UUID) Value() (driver.Value, error) {
    return u.UUID.Value()
}

func (u UUID) MarshalText() ([]byte, error) {
    return u.UUID.MarshalText()
}

func (u *UUID) UnmarshalText(data []byte) error {
    err := u.UUID.UnmarshalText(data)
    u.Str = u.UUID.String()
    return err
}

func (u UUID) MarshalBinary() ([]byte, error) {
    return u.UUID.MarshalBinary()
}

func (u *UUID) UnmarshalBinary(data []byte) error {
    err := u.UUID.UnmarshalBinary(data)
    u.Str = u.UUID.String()
    return err
}
