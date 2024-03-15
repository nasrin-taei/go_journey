package main

import "fmt"

type mobile interface { //in order to implement this interface, an struct must implement all of this functions.
	call(number string)
	sendSms(number string)
}

type iphone struct { // implementation of mobile.
	modelNumber string
}

func (i iphone) call(phoneNumber string) {
	fmt.Printf("iphone calling %s ...\n", phoneNumber)
}

func (i *iphone) sendSms(number string) {
	fmt.Println("sending sms via iphone")
}

func (i iphone) scanByFaceID() {
	fmt.Println("Bring your face over my sensor.")
}

type samsung struct { // implementation of mobile.
	serialNumber string
}

func (s samsung) call(number string) {
	fmt.Printf("samsung is calling number %s\n", number)
}

func (s samsung) sendSms(number string) {
	fmt.Printf("sending SMS to %s via samsung\n", number)
}

func (s *samsung) useFingerPrint(fingerNumber int) {
	fmt.Printf("Bring your finger number %d to the finger print scanner\n", fingerNumber)
}

func main() {
	//first example
	var ourMobile mobile // it's nil
	ourMobile = samsung{ //Reference is not needed because despite the fact that useFingerPrint is pointer receiver function, it is not one of mobile interface functions
		serialNumber: "7987KLJ",
	}
	ourMobile.sendSms("5465456646")

	ourMobile = &iphone{ //reference of iphone passed because one of its functions has pointer receiver.
		modelNumber: "iphone 14",
	}
	ourMobile.sendSms("5465456646")

	// second example
	behradsPhone := iphone{
		modelNumber: "13",
	}

	behradsPhone.scanByFaceID()
	behradsPhone.call("+9812345678")

	nasrinsPhone := samsung{
		serialNumber: "987GKLKJKVK6587",
	}

	nasrinsPhone.useFingerPrint(5)
	nasrinsPhone.call("+9812345678")

	doEmergencyCall(nasrinsPhone)
	doEmergencyCall(&behradsPhone)

	assertType(nasrinsPhone)
	assertType2(&behradsPhone)
}

func doEmergencyCall(phoneDevice mobile) {
	phoneDevice.call("110")
}

func assertType(phoneDevice mobile) {
	switch v := phoneDevice.(type) {
	case *iphone:
		fmt.Printf("phone device is iphone %v\n", v)
		v.scanByFaceID()
	case samsung:
		fmt.Printf("phone device is samsung %v\n", v)
		v.useFingerPrint(5)
	default:
		fmt.Println("phone device is nil")
	}
}

func assertType2(phoneDevice mobile) {
	x, ok := phoneDevice.(*iphone)
	if ok {
		x.sendSms("8797698")
	}

	y, ok := phoneDevice.(samsung)
	if ok {
		y.useFingerPrint(10)
	}
}
