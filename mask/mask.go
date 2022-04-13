package mask

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scan(&input)
	cntAllMasks, _ := strconv.Atoi(input)

	res := countPacks(cntAllMasks)
	println(res)
}

func countPacks(inp int) string {
	var cntMascksSec, cntMascksT int
	var allPrice, maxBoxPrice, maxPacksPrice, allPriceBox, allPricePack float64
	var priceBox float64 = 160
	var pricePack float64 = 11
	priceOne := 0.55
	cntPack := 24
	cntBox := 16 * cntPack

	cntBoxesMy := int(inp / cntBox)
	if cntBoxesMy > 0 {
		cntMascksSec = inp - cntBoxesMy*cntBox
		allPriceBox = float64(cntBoxesMy) * priceBox
	} else {
		cntMascksSec = inp
	}

	cntPacksMy := int(cntMascksSec / cntPack)
	if cntPacksMy > 0 {
		cntMascksT = cntMascksSec - cntPacksMy*cntPack
		allPricePack = float64(cntBoxesMy) * priceBox
	} else {
		cntMascksT = cntMascksSec
	}

	allPrice = allPriceBox + allPricePack + float64(cntMascksT)*priceOne

	maxBoxPrice = (float64(cntBoxesMy) + 1) * priceBox
	if maxBoxPrice < allPrice {
		cntBoxesMy = cntBoxesMy + 1
		cntPacksMy = 0
		cntMascksT = 0

		return strconv.Itoa(cntMascksT) + " " + strconv.Itoa(cntPacksMy) + " " + strconv.Itoa(cntBoxesMy)
	}

	maxPacksPrice = allPriceBox + (float64(cntPacksMy)+1)*pricePack
	if maxPacksPrice < allPrice {
		cntPacksMy = cntPacksMy + 1
		cntMascksT = 0

		return strconv.Itoa(cntMascksT) + " " + strconv.Itoa(cntPacksMy) + " " + strconv.Itoa(cntBoxesMy)
	}

	return strconv.Itoa(cntMascksT) + " " + strconv.Itoa(cntPacksMy) + " " + strconv.Itoa(cntBoxesMy)
}
