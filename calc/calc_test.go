package main

import (
	"strings"
	"testing"
)

func TestCheckSuccess1(t *testing.T) {
	var inputString = "1*(6*8*6)*(18/2/3/3*9)+4*17"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 2660 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess2(t *testing.T) {
	var inputString = "(2+2)*2"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 8 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess3(t *testing.T) {
	var inputString = "19+6*8*6+18/2/3/3*9+4"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 320 {
		t.Fatalf("fail")
	}
}
func TestCheckSuccess4(t *testing.T) {
	var inputString = "(10-5)*2"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 10 {
		t.Fatalf("Incorrect result")
	}
}
func TestCheckSuccess5(t *testing.T) {
	var inputString = "(3*6)/3"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 6 {
		t.Fatalf("Incorrect result")
	}
}
func TestCheckSuccess6(t *testing.T) {
	var inputString = "(8+12)-10"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 10 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess7(t *testing.T) {
	var inputString = "(15-7)+5"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 13 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess8(t *testing.T) {
	var inputString = "(4 * 9) / 2"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 18 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess9(t *testing.T) {
	var inputString = "(25 / 5) - 1"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 4 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess10(t *testing.T) {
	var inputString = "(6 + 3) * 2"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 18 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess11(t *testing.T) {
	var inputString = "((2 + 2) * 2) / 2 + 1 - 1"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 4 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess12(t *testing.T) {
	var inputString = "((10 - 5) * 2) + 5 - 3"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 12 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess13(t *testing.T) {
	var inputString = "((3 * 6) / 3) + 2 - 1"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 7 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess14(t *testing.T) {
	var inputString = "((8 + 12) - 10) * 3 / 2 "
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 15 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess15(t *testing.T) {
	var inputString = "((15 - 7) + 5) * 2 - 3"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 23 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess16(t *testing.T) {
	var inputString = "((4 * 9) / 2) + 3 - 1"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 20 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess17(t *testing.T) {
	var inputString = "((25 / 5) - 1) * 3 + 2"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 14 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess18(t *testing.T) {
	var inputString = "((25 / 5) - 1) * 3 + 2"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 14 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess19(t *testing.T) {
	var inputString = "((6 + 3) * 2) / 3 + 1"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 7 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess20(t *testing.T) {
	var inputString = "((18 - 10) + 3) * 2 - 4"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 18 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess21(t *testing.T) {
	var inputString = "(((18 - 10) + 3) * 2 - 4)*2 - 2*(((18 - 10) + 3) * 2 - 4)"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 0 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess22(t *testing.T) {
	var inputString = "((((18 - 10) + 3) * 2 - 4)*2 - 2*(((18 - 10) + 3) * 2 - 4))*101+7*(((25 / 5) - 1) * 3 + 2)"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 98 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess23(t *testing.T) {
	var inputString = "(1*(1*(1*(1*1))))"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 1 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess24(t *testing.T) {
	var inputString = "(1+(1*(1*(1+1))))"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 3 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess25(t *testing.T) {
	var inputString = "19/19 + 20/20 - 30/30 - 10/10"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 0 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess26(t *testing.T) {
	var inputString = "228*15-347584+(1919-999*0-(16*2))"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != -342277 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess27(t *testing.T) {
	var inputString = "8*(90+999)-189*11+(111-19)"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 6725 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess28(t *testing.T) {
	var inputString = "0"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 0 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess29(t *testing.T) {
	var inputString = "-100"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != -100 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess30(t *testing.T) {
	var inputString = "-1111111111"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != -1111111111 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess31(t *testing.T) {
	var inputString = "(-1111111111)*4"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != -4444444444 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess32(t *testing.T) {
	var inputString = "(-1111111111)+(-111111111)"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != -1222222222 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess33(t *testing.T) {
	var inputString = "(-1111111111)*4"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != -4444444444 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess34(t *testing.T) {
	var inputString = "4*(-100)"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != -400 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess35(t *testing.T) {
	var inputString = "(-100)*4"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != -400 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess36(t *testing.T) {
	var inputString = "(-100)"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != -100 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess37(t *testing.T) {
	var inputString = "(+100)"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if answer != 100 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckSuccess38(t *testing.T) {
	var inputString = "-1111111111-1111111111"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if answer != -2222222222 {
		t.Fatalf("Incorrect result")
	}
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCheckSuccess39(t *testing.T) {
	var inputString = "-1111111111*4"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if answer != -4444444444 {
		t.Fatalf("Incorrect result")
	}
	if err != nil {
		t.Fatalf("%s", err)
	}

}

func TestCheckSuccess40(t *testing.T) {
	var inputString = "-100*4"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if answer != -400 {
		t.Fatalf("Incorrect result")
	}
	if err != nil {
		t.Fatalf("%s", err)
	}

}

func TestCheckSuccess41(t *testing.T) {
	var inputString = "-200+100*4"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if answer != 200 {
		t.Fatalf("Incorrect result")
	}
	if err != nil {
		t.Fatalf("%s", err)
	}

}

func TestCheckSuccess42(t *testing.T) {
	var inputString = "1/2 + 1/2"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if answer != 1 {
		t.Fatalf("Incorrect result")
	}
	if err != nil {
		t.Fatalf("%s", err)
	}

}

func TestCheckSuccess43(t *testing.T) {
	var inputString = "1/4 + 1/4"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	answer, err := Calculator(stringToCalc)
	if answer != 1 { // 1/2!!!
		t.Fatalf("Incorrect result")
	}
	if err != nil {
		t.Fatalf("%s", err)
	}

}

func TestCheckFail(t *testing.T) {
	var inputString = "4*-100"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCheckFail1(t *testing.T) {
	var inputString = "1000--1000"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}

}

func TestCheckFail2(t *testing.T) {
	var inputString = "222**222"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}

}

func TestCheckFail3(t *testing.T) {
	var inputString = "339//222"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}

}

func TestCheckFail4(t *testing.T) {
	var inputString = "*200"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCheckFail5(t *testing.T) {
	var inputString = "(/200)+100*4"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCheckFail6(t *testing.T) {
	var inputString = "(-200)-"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCheckFail7(t *testing.T) {
	var inputString = "200/"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCheckFail8(t *testing.T) {
	var inputString = "200+200*"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCheckFail9(t *testing.T) {
	var inputString = "200*"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCheckFail10(t *testing.T) {
	var inputString = "-200+++200"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCheckFail11(t *testing.T) {
	var inputString = "-200+++++200"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCheckFail12(t *testing.T) {
	var inputString = "200/-(-200)"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCheckFail13(t *testing.T) {
	var inputString = "1000+6*200+(-100&(-100))"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCheckFail14(t *testing.T) {
	var inputString = "(-100)&10+6%200@12"
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	_, err := Calculator(stringToCalc)
	if err != nil {
		t.Fatalf("%s", err)
	}
}
