//go:generate stringer -type=CloseReason -output ext.close_reason_string.go

package errors

type CloseReason Code

const (
	CloseNormal                          CloseReason = iota // both. 正常关闭连接.不发送遗嘱.
	CloseByPeer                          CloseReason = 1    // both.
	CloseUnknownError                    CloseReason = 2    // both. 未知原因的错误
	CloseWithWillMessage                 CloseReason = 4    // client. 包含遗嘱消息的断开. 客户端希望断开但也需要服务端发布它的遗嘱消息.
	CloseUnspecifiedReason               CloseReason = 128  // both. 未指定错误 / unspecified error reason. 连接被关闭, 但发送端不愿意透露原因, 或者没有其他适用的原因码.
	CloseMalformedPacket                 CloseReason = 129  // client or server. 无效的报文. 收到的报文不符合本规范.
	CloseWrongProtocol                   CloseReason = 130  // both. 协议错误. 收到意外的或无序的报文.
	CloseNotImplement                    CloseReason = 131  // both. 收到的报文有效, 但根据实现无法进行处理.
	CloseNotAuthorized                   CloseReason = 135  // server. 请求没有被授权
	CloseBusy                            CloseReason = 137  // client. 服务端正忙且不能继续处理此客户端的请求
	CloseClosing                         CloseReason = 139  // client. 服务正在关闭
	CloseKeepAliveTimeout                CloseReason = 141  // both/server. 连接因为在超过1.5倍的保持连接时间内没有收到任何报文而关闭.
	CloseDupSession                      CloseReason = 143  // server. 另一个使用了相同的客户标识符的连接已建立, 导致此连接关闭.
	CloseTopicFilterInvalid              CloseReason = 143  // server. 主题过滤器格式正确, 但不被服务端所接受.
	CloseTopicNameInvalid                CloseReason = 144  // client or server. 主题名格式正确, 但不被客户端或服务端所接受.
	CloseBufferOverflow                  CloseReason = 147  // client or server. 客户端或服务端收到了数量超过接收最大值的未发送PUBACK或PUBCOMP的发布消息.
	CloseTopicAliasInvalid               CloseReason = 148  // client or server. 客户端或服务端收到的PUBLISH报文包含的主题别名大于其在CONNECT或CONNACK中发送的主题别名最大值.
	ClosePacketTooLong                   CloseReason = 149  // client or server. 报文过大. 报文长度大于此客户端或服务端的最大报文长度.
	ClosePacketRateTooFast               CloseReason = 150  // both. 消息速率过高. 收到的数据速率太高.
	CloseQuoteOverflow                   CloseReason = 151  // both. 超出配额. 已超出实现限制或管理限制
	CloseAdministrativeOperate           CloseReason = 152  // both. 管理操作. 连接因为管理操作被关闭
	ClosePayloadFormatInvalid            CloseReason = 153  // client or server. 载荷格式无效. 载荷格式与指定的载荷格式指示符不匹配
	CloseUnsupportedRetain               CloseReason = 154  // server. 不支持保留. 服务端不支持保留消息
	CloseUnsupportedQoS                  CloseReason = 155  // server. 不支持的QoS等级. 客户端指定的QoS等级大于CONNACK报文中指定的最大QoS等级.
	CloseUseOtherServer                  CloseReason = 156  // server. (临时)使用其他服务端. 客户端应该临时使用其他服务端
	CloseServerMovedPermanent            CloseReason = 157  // server. 服务端已(永久)移动. 服务端已移动且客户端应该永久使用其他服务端.
	CloseUnsupportedSharedSubscription   CloseReason = 158  // server. 不支持共享订阅. 服务端不支持共享订阅.
	CloseConnectionMadeTooFreq           CloseReason = 159  // server. 超出连接速率限制. 此连接因为连接速率过高而被关闭
	CloseConnectionTimeout               CloseReason = 160  // server. 最大连接时间. 超出为此连接授予的最大连接时间
	CloseUnsupportedSubscriptionId       CloseReason = 161  // server. 不支持订阅标识符. 服务端不支持订阅标识符；订阅未被接受.
	CloseUnsupportedWildcardSubscription CloseReason = 162  // server. 不支持通配符订阅. 服务端不支持通配符订阅；订阅未被接受.
)

// --- PASS: TestCloseReasonConstants (0.00s)
// lib_test.go:13:   - 0 (0) => "CloseNormal" [true]
// lib_test.go:13:   - 2 (2) => "CloseUnknownError" [true]
// lib_test.go:13:   - 4 (4) => "CloseWithWillMessage" [true]
// lib_test.go:13:   - 80 (128) => "CloseUnspecifiedReason" [true]
// lib_test.go:13:   - 81 (129) => "CloseMalformedPacket" [true]
// lib_test.go:13:   - 82 (130) => "CloseWrongProtocol" [true]
// lib_test.go:13:   - 83 (131) => "CloseNotImplement" [true]
// lib_test.go:13:   - 87 (135) => "CloseNotAuthorized" [true]
// lib_test.go:13:   - 89 (137) => "CloseBusy" [true]
// lib_test.go:13:   - 8b (139) => "CloseClosing" [true]
// lib_test.go:13:   - 8d (141) => "CloseKeepAliveTimeout" [true]
// lib_test.go:13:   - 8f (143) => "CloseDupSession" [true]
// lib_test.go:13:   - 90 (144) => "CloseTopicNameInvalid" [true]
// lib_test.go:13:   - 93 (147) => "CloseBufferOverflow" [true]
// lib_test.go:13:   - 94 (148) => "CloseTopicAliasInvalid" [true]
// lib_test.go:13:   - 95 (149) => "ClosePacketTooLong" [true]
// lib_test.go:13:   - 96 (150) => "ClosePacketRateTooFast" [true]
// lib_test.go:13:   - 97 (151) => "CloseQuoteOverflow" [true]
// lib_test.go:13:   - 98 (152) => "CloseAdministrativeOperate" [true]
// lib_test.go:13:   - 99 (153) => "ClosePayloadFormatInvalid" [true]
// lib_test.go:13:   - 9a (154) => "CloseUnsupportedRetain" [true]
// lib_test.go:13:   - 9b (155) => "CloseUnsupportedQoS" [true]
// lib_test.go:13:   - 9c (156) => "CloseUseOtherServer" [true]
// lib_test.go:13:   - 9d (157) => "CloseServerMovedPermanent" [true]
// lib_test.go:13:   - 9e (158) => "CloseUnsupportedSharedSubscription" [true]
// lib_test.go:13:   - 9f (159) => "CloseConnectionMadeTooFreq" [true]
// lib_test.go:13:   - a0 (160) => "CloseConnectionTimeout" [true]
// lib_test.go:13:   - a1 (161) => "CloseUnsupportedSubscriptionId" [true]
// lib_test.go:13:   - a2 (162) => "CloseUnsupportedWildcardSubscription" [true]

func (i CloseReason) New(msg string, args ...interface{}) *MqttError {
	return New(msg, args...).Reason(i)
}

func (i CloseReason) IsValidReason() bool {
	switch {
	case 0 <= i && i <= 2:
	case i == 4:
	case 128 <= i && i <= 131:
	case i == 135:
	case i == 137:
	case i == 139:
	case i == 141:
	case 143 <= i && i <= 144:
	case 147 <= i && i <= 162:
	default:
		return false
	}
	return true
}

func (i CloseReason) Parse(s string) CloseReason {
	switch s {
	// case 0 <= i && i <= 2:
	// 	return _CloseReason_name_0[_CloseReason_index_0[i]:_CloseReason_index_0[i+1]]
	case _CloseReason_name_1:
		return 4
	// case 128 <= i && i <= 131:
	// 	i -= 128
	// 	return _CloseReason_name_2[_CloseReason_index_2[i]:_CloseReason_index_2[i+1]]
	case _CloseReason_name_3:
		return 135
	case _CloseReason_name_4:
		return 137
	case _CloseReason_name_5:
		return 139
	case _CloseReason_name_6:
		return 141
	// case 143 <= i && i <= 144:
	// 	i -= 143
	// 	return _CloseReason_name_7[_CloseReason_index_7[i]:_CloseReason_index_7[i+1]]
	// case 147 <= i && i <= 162:
	// 	i -= 147
	// 	return _CloseReason_name_8[_CloseReason_index_8[i]:_CloseReason_index_8[i+1]]

	default:
		if c, bad := _ParseCloseReason4Uint8(s, _CloseReason_name_0, _CloseReason_index_0, 0); !bad {
			return c
		}
		if c, bad := _ParseCloseReason5Uint8(s, _CloseReason_name_2, _CloseReason_index_2, 147); !bad {
			return c
		}
		if c, bad := _ParseCloseReason3Uint8(s, _CloseReason_name_7, _CloseReason_index_7, 143); !bad {
			return c
		}
		if c, bad := _ParseCloseReasonUint16(s, _CloseReason_name_8, _CloseReason_index_8, 147); !bad {
			return c
		}
	}
	return 0
}

func _ParseCloseReason4Uint8(s string, _m string, _i [4]uint8, base CloseReason) (c CloseReason, bad bool) {
	for t := 0; t < len(_i); t++ {
		b := _i[t]
		e := len(_m)
		if t < len(_i)-1 {
			e = int(_i[t+1])
		}
		n := _m[b:e]
		if n == s {
			return CloseReason(t + int(base)), false
		}
	}
	return 0, true
}

func _ParseCloseReason5Uint8(s string, _m string, _i [5]uint8, base CloseReason) (c CloseReason, bad bool) {
	for t := 0; t < len(_i); t++ {
		b := _i[t]
		e := len(_m)
		if t < len(_i)-1 {
			e = int(_i[t+1])
		}
		n := _m[b:e]
		if n == s {
			return CloseReason(t + int(base)), false
		}
	}
	return 0, true
}

func _ParseCloseReason3Uint8(s string, _m string, _i [3]uint8, base CloseReason) (c CloseReason, bad bool) {
	for t := 0; t < len(_i); t++ {
		b := _i[t]
		e := len(_m)
		if t < len(_i)-1 {
			e = int(_i[t+1])
		}
		n := _m[b:e]
		if n == s {
			return CloseReason(t + int(base)), false
		}
	}
	return 0, true
}

func _ParseCloseReasonUint16(s string, _m string, _i [17]uint16, base CloseReason) (c CloseReason, bad bool) {
	for t := 0; t < len(_i); t++ {
		b := _i[t]
		e := len(_m)
		if t < len(_i)-1 {
			e = int(_i[t+1])
		}
		n := _m[b:e]
		if n == s {
			return CloseReason(t + int(base)), false
		}
	}
	return 0, true
}

// func (n CloseReason) String() string {
// 	if x, ok := CloseReasonToString[n]; ok {
// 		return x
// 	}
// 	return CloseReasonToString[CloseUnknownErr]
// }
//
// func (n CloseReason) FromString(s string) CloseReason {
// 	if x, ok := CloseReasonFromString[s]; ok {
// 		return x
// 	}
// 	return CloseUnknownErr
// }
