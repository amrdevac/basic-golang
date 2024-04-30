package loket

type Create struct {
	NamaLoket string `validate:"required,min=12,max=13,alphanum"  binding:"required" json:"nama_loket"`
	Bambang   string `validate:"required,noDot,noWhiteSpace" binding:"required" json:"ini_bambang"`
}

// type User struct {
// 	FirstName      string     `validate:"required"`
// 	LastName       string     `validate:"required"`
// 	Age            uint8      `validate:"gte=0,lte=130"`
// 	Email          string     `validate:"required,email"`
// 	Gender         string     `validate:"oneof=male female prefer_not_to"`
// 	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
// 	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
// }
