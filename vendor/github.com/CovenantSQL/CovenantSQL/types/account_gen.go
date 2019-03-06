package types

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	hsp "github.com/CovenantSQL/HashStablePack/marshalhash"
)

// MarshalHash marshals for hash
func (z *Account) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 4
	o = append(o, 0x84)
	if oTemp, err := z.Address.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	if oTemp, err := z.NextNonce.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = hsp.AppendFloat64(o, z.Rating)
	o = hsp.AppendArrayHeader(o, uint32(SupportTokenNumber))
	for za0001 := range z.TokenBalance {
		o = hsp.AppendUint64(o, z.TokenBalance[za0001])
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Account) Msgsize() (s int) {
	s = 1 + 8 + z.Address.Msgsize() + 10 + z.NextNonce.Msgsize() + 7 + hsp.Float64Size + 13 + hsp.ArrayHeaderSize + (int(SupportTokenNumber) * (hsp.Uint64Size))
	return
}

// MarshalHash marshals for hash
func (z *MinerInfo) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 9
	o = append(o, 0x89)
	if oTemp, err := z.Address.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = hsp.AppendUint64(o, z.Deposit)
	o = hsp.AppendString(o, z.EncryptionKey)
	o = hsp.AppendString(o, z.Name)
	if oTemp, err := z.NodeID.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = hsp.AppendUint64(o, z.PendingIncome)
	o = hsp.AppendUint64(o, z.ReceivedIncome)
	o = hsp.AppendInt32(o, int32(z.Status))
	o = hsp.AppendArrayHeader(o, uint32(len(z.UserArrears)))
	for za0001 := range z.UserArrears {
		if z.UserArrears[za0001] == nil {
			o = hsp.AppendNil(o)
		} else {
			// map header, size 2
			o = append(o, 0x82)
			if oTemp, err := z.UserArrears[za0001].User.MarshalHash(); err != nil {
				return nil, err
			} else {
				o = hsp.AppendBytes(o, oTemp)
			}
			o = hsp.AppendUint64(o, z.UserArrears[za0001].Arrears)
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MinerInfo) Msgsize() (s int) {
	s = 1 + 8 + z.Address.Msgsize() + 8 + hsp.Uint64Size + 14 + hsp.StringPrefixSize + len(z.EncryptionKey) + 5 + hsp.StringPrefixSize + len(z.Name) + 7 + z.NodeID.Msgsize() + 14 + hsp.Uint64Size + 15 + hsp.Uint64Size + 7 + hsp.Int32Size + 12 + hsp.ArrayHeaderSize
	for za0001 := range z.UserArrears {
		if z.UserArrears[za0001] == nil {
			s += hsp.NilSize
		} else {
			s += 1 + 5 + z.UserArrears[za0001].User.Msgsize() + 8 + hsp.Uint64Size
		}
	}
	return
}

// MarshalHash marshals for hash
func (z *ProviderProfile) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 9
	o = append(o, 0x89)
	o = hsp.AppendUint64(o, z.Deposit)
	o = hsp.AppendUint64(o, z.GasPrice)
	o = hsp.AppendFloat64(o, z.LoadAvgPerCPU)
	o = hsp.AppendUint64(o, z.Memory)
	if oTemp, err := z.NodeID.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	if oTemp, err := z.Provider.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = hsp.AppendUint64(o, z.Space)
	o = hsp.AppendArrayHeader(o, uint32(len(z.TargetUser)))
	for za0001 := range z.TargetUser {
		if oTemp, err := z.TargetUser[za0001].MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	if oTemp, err := z.TokenType.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ProviderProfile) Msgsize() (s int) {
	s = 1 + 8 + hsp.Uint64Size + 9 + hsp.Uint64Size + 14 + hsp.Float64Size + 7 + hsp.Uint64Size + 7 + z.NodeID.Msgsize() + 9 + z.Provider.Msgsize() + 6 + hsp.Uint64Size + 11 + hsp.ArrayHeaderSize
	for za0001 := range z.TargetUser {
		s += z.TargetUser[za0001].Msgsize()
	}
	s += 10 + z.TokenType.Msgsize()
	return
}

// MarshalHash marshals for hash
func (z *SQLChainProfile) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 11
	o = append(o, 0x8b)
	if oTemp, err := z.Address.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = hsp.AppendBytes(o, z.EncodedGenesis)
	o = hsp.AppendUint64(o, z.GasPrice)
	if oTemp, err := z.ID.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = hsp.AppendUint32(o, z.LastUpdatedHeight)
	if oTemp, err := z.Meta.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = hsp.AppendArrayHeader(o, uint32(len(z.Miners)))
	for za0001 := range z.Miners {
		if z.Miners[za0001] == nil {
			o = hsp.AppendNil(o)
		} else {
			if oTemp, err := z.Miners[za0001].MarshalHash(); err != nil {
				return nil, err
			} else {
				o = hsp.AppendBytes(o, oTemp)
			}
		}
	}
	if oTemp, err := z.Owner.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = hsp.AppendUint64(o, z.Period)
	if oTemp, err := z.TokenType.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = hsp.AppendArrayHeader(o, uint32(len(z.Users)))
	for za0002 := range z.Users {
		if z.Users[za0002] == nil {
			o = hsp.AppendNil(o)
		} else {
			if oTemp, err := z.Users[za0002].MarshalHash(); err != nil {
				return nil, err
			} else {
				o = hsp.AppendBytes(o, oTemp)
			}
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SQLChainProfile) Msgsize() (s int) {
	s = 1 + 8 + z.Address.Msgsize() + 15 + hsp.BytesPrefixSize + len(z.EncodedGenesis) + 9 + hsp.Uint64Size + 3 + z.ID.Msgsize() + 18 + hsp.Uint32Size + 5 + z.Meta.Msgsize() + 7 + hsp.ArrayHeaderSize
	for za0001 := range z.Miners {
		if z.Miners[za0001] == nil {
			s += hsp.NilSize
		} else {
			s += z.Miners[za0001].Msgsize()
		}
	}
	s += 6 + z.Owner.Msgsize() + 7 + hsp.Uint64Size + 10 + z.TokenType.Msgsize() + 6 + hsp.ArrayHeaderSize
	for za0002 := range z.Users {
		if z.Users[za0002] == nil {
			s += hsp.NilSize
		} else {
			s += z.Users[za0002].Msgsize()
		}
	}
	return
}

// MarshalHash marshals for hash
func (z SQLChainRole) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	o = hsp.AppendByte(o, byte(z))
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z SQLChainRole) Msgsize() (s int) {
	s = hsp.ByteSize
	return
}

// MarshalHash marshals for hash
func (z *SQLChainUser) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 6
	o = append(o, 0x86)
	if oTemp, err := z.Address.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = hsp.AppendUint64(o, z.AdvancePayment)
	o = hsp.AppendUint64(o, z.Arrears)
	o = hsp.AppendUint64(o, z.Deposit)
	if z.Permission == nil {
		o = hsp.AppendNil(o)
	} else {
		// map header, size 2
		o = append(o, 0x82)
		o = hsp.AppendInt32(o, int32(z.Permission.Role))
		o = hsp.AppendArrayHeader(o, uint32(len(z.Permission.Patterns)))
		for za0001 := range z.Permission.Patterns {
			o = hsp.AppendString(o, z.Permission.Patterns[za0001])
		}
	}
	o = hsp.AppendInt32(o, int32(z.Status))
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SQLChainUser) Msgsize() (s int) {
	s = 1 + 8 + z.Address.Msgsize() + 15 + hsp.Uint64Size + 8 + hsp.Uint64Size + 8 + hsp.Uint64Size + 11
	if z.Permission == nil {
		s += hsp.NilSize
	} else {
		s += 1 + 5 + hsp.Int32Size + 9 + hsp.ArrayHeaderSize
		for za0001 := range z.Permission.Patterns {
			s += hsp.StringPrefixSize + len(z.Permission.Patterns[za0001])
		}
	}
	s += 7 + hsp.Int32Size
	return
}

// MarshalHash marshals for hash
func (z Status) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	o = hsp.AppendInt32(o, int32(z))
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Status) Msgsize() (s int) {
	s = hsp.Int32Size
	return
}

// MarshalHash marshals for hash
func (z *UserArrears) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 2
	o = append(o, 0x82)
	o = hsp.AppendUint64(o, z.Arrears)
	if oTemp, err := z.User.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UserArrears) Msgsize() (s int) {
	s = 1 + 8 + hsp.Uint64Size + 5 + z.User.Msgsize()
	return
}

// MarshalHash marshals for hash
func (z *UserPermission) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 2
	o = append(o, 0x82)
	o = hsp.AppendArrayHeader(o, uint32(len(z.Patterns)))
	for za0001 := range z.Patterns {
		o = hsp.AppendString(o, z.Patterns[za0001])
	}
	o = hsp.AppendInt32(o, int32(z.Role))
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UserPermission) Msgsize() (s int) {
	s = 1 + 9 + hsp.ArrayHeaderSize
	for za0001 := range z.Patterns {
		s += hsp.StringPrefixSize + len(z.Patterns[za0001])
	}
	s += 5 + hsp.Int32Size
	return
}

// MarshalHash marshals for hash
func (z UserPermissionRole) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	o = hsp.AppendInt32(o, int32(z))
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z UserPermissionRole) Msgsize() (s int) {
	s = hsp.Int32Size
	return
}