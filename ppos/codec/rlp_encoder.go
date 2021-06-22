package codec

const (

	/**
	 * [0x80]
	 * If a string is 0-55 bytes long, the RLP encoding consists of a single
	 * byte with value 0x80 plus the length of the string followed by the
	 * string. The range of the first byte is thus [0x80, 0xb7].
	 */
	OFFSET_SHORT_STRING = 0x80

	/**
	 * [0xb7]
	 * If a string is more than 55 bytes long, the RLP encoding consists of a
	 * single byte with value 0xb7 plus the length of the length of the string
	 * in binary form, followed by the length of the string, followed by the
	 * string. For example, a length-1024 string would be encoded as
	 * \xb9\x04\x00 followed by the string. The range of the first byte is thus
	 * [0xb8, 0xbf].
	 */
	OFFSET_LONG_STRING = 0xb7

	/**
	 * [0xc0]
	 * If the total payload of a list (i.e. the combined length of all its
	 * items) is 0-55 bytes long, the RLP encoding consists of a single byte
	 * with value 0xc0 plus the length of the list followed by the concatenation
	 * of the RLP encodings of the items. The range of the first byte is thus
	 * [0xc0, 0xf7].
	 */
	OFFSET_SHORT_LIST = 0xc0

	/**
	 * [0xf7]
	 * If the total payload of a list is more than 55 bytes long, the RLP
	 * encoding consists of a single byte with value 0xf7 plus the length of the
	 * length of the list in binary form, followed by the length of the list,
	 * followed by the concatenation of the RLP encodings of the items. The
	 * range of the first byte is thus [0xf8, 0xff].
	 */
	OFFSET_LONG_LIST = 0xf7
)

type BytesSlice []byte

type ParamEncoder interface {
	GetEncodeData() BytesSlice
}

func ToByteArray(value int) []byte {
	return []byte{
		byte((value >> 24) & 0xff),
		byte((value >> 16) & 0xff),
		byte((value >> 8) & 0xff),
		byte(value & 0xff),
	}
}

func MinimalByteArray(value int) []byte {
	encoded := ToByteArray(value)
	for i, v := range encoded {
		if v != 0 {
			return encoded[i:]
		}
	}
	return []byte{}
}

func EncodeBytes(bs []byte, offset byte) []byte {
	if len(bs) == 1 && offset == OFFSET_SHORT_STRING && bs[0] >= 0x00 && bs[0] <= 0x7f {
		return bs
	} else if len(bs) <= 55 {
		r := []byte{offset + byte(len(bs))}
		r = append(r, bs...)
		return r
	} else {
		encodedStringLength := MinimalByteArray(len(bs))
		r := []byte{(offset + 0x37) + byte(len(encodedStringLength))}
		r = append(r, encodedStringLength...)
		r = append(r, bs...)
		return r
	}
}

func EncodeBytesSlice(list []BytesSlice) BytesSlice {
	if len(list) == 0 {
		return EncodeBytes([]byte{}, OFFSET_SHORT_LIST)
	} else {
		var r []byte
		for _, v := range list {
			r = append(r, EncodeBytes(v, OFFSET_SHORT_STRING)...)
		}
		return EncodeBytes(r, OFFSET_SHORT_LIST)
	}
}

type RlpEncoder struct{}

func encodeInterface(intf interface{}) []byte {
	switch intf.(type) {
	case RlpString:
		return encodeRlpString(intf.(RlpString))
	case RlpList:
		return encodeRlpList(intf.(RlpList))
	default:
		panic("not support parameter type for RlpEncoder")
	}
}

func encodeRlpString(v RlpString) []byte {
	return EncodeBytes(v.GetBytes(), OFFSET_SHORT_STRING)
}

func encodeRlpList(v RlpList) []byte {
	values := v.GetValues()
	if len(values) == 0 {
		return EncodeBytes([]byte{}, OFFSET_SHORT_LIST)
	} else {
		var result []byte
		for _, it := range values {
			result = append(result, encodeInterface(it)...)
		}
		return EncodeBytes(result, OFFSET_SHORT_LIST)
	}
}

func (re RlpEncoder) Encode(intf interface{}) []byte {
	return encodeInterface(intf)
}
