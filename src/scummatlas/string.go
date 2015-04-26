package scummatlas

func parsePrintOpcode(data []byte, offset int) (say string, length int) {
	// Process subinstructions
	var subinst byte
	subinst = 0x0f
	for subinst != 0x00 && offset < len(data) {
		subinst := data[offset]
		switch subinst {
		case 0x0F:
			say, length = parseString(data, offset+1)
			offset += length
		case 0x01, 0x02:
			//TODO encode color, right
			offset += 2
		case 0x00, 0x03, 0x08:
			//TODO encode position, width, offset
			offset += 4
		case 0x04, 0x06, 0x07:
			offset++
		default:
			offset++
			break
		}
	}
	return say, length
}

func parseString(data []byte, offset int) (say string, length int) {
	originalOffset := offset - 2
	i := 0
	say = ""
	for {
		i++
		if i > 200 {
			break
		}
		currChar := data[offset]
		if currChar == 0xff {
			escapeChar := data[offset+1]
			switch {
			case 0x01 <= escapeChar && escapeChar <= 0x03:
				offset += 1
			case 0x04 <= escapeChar && escapeChar <= 0x0e:
				offset += 4
			}
		} else if currChar >= 0x20 && currChar <= 0x7e { //printable ascii char
			offset++
			say += string(currChar)
		} else if currChar >= 0x00 {
			offset++
			break
		} else {
			panic("Invalid character in print")
		}
	}
	return say, offset - originalOffset
}