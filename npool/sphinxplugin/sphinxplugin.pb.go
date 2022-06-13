// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/sphinxplugin/sphinxplugin.proto

package sphinxplugin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CoinType int32

const (
	CoinType_CoinTypeUnKnow CoinType = 0
	// mainnet
	CoinType_CoinTypefilecoin  CoinType = 1
	CoinType_CoinTypebitcoin   CoinType = 2
	CoinType_CoinTypeethereum  CoinType = 3
	CoinType_CoinTypeusdterc20 CoinType = 4
	CoinType_CoinTypespacemesh CoinType = 5
	CoinType_CoinTypesolana    CoinType = 6
	CoinType_CoinTypeusdttrc20 CoinType = 7
	CoinType_CoinTypebsc       CoinType = 8
	CoinType_CoinTypetron      CoinType = 9
	// testnet
	CoinType_CoinTypetfilecoin  CoinType = 100
	CoinType_CoinTypetbitcoin   CoinType = 101
	CoinType_CoinTypetethereum  CoinType = 102
	CoinType_CoinTypetusdterc20 CoinType = 103
	CoinType_CoinTypetspacemesh CoinType = 104
	CoinType_CoinTypetsolana    CoinType = 105
	CoinType_CoinTypetusdttrc20 CoinType = 106
	CoinType_CoinTypetbsc       CoinType = 107
	CoinType_CoinTypettron      CoinType = 108
)

// Enum value maps for CoinType.
var (
	CoinType_name = map[int32]string{
		0:   "CoinTypeUnKnow",
		1:   "CoinTypefilecoin",
		2:   "CoinTypebitcoin",
		3:   "CoinTypeethereum",
		4:   "CoinTypeusdterc20",
		5:   "CoinTypespacemesh",
		6:   "CoinTypesolana",
		7:   "CoinTypeusdttrc20",
		8:   "CoinTypebsc",
		9:   "CoinTypetron",
		100: "CoinTypetfilecoin",
		101: "CoinTypetbitcoin",
		102: "CoinTypetethereum",
		103: "CoinTypetusdterc20",
		104: "CoinTypetspacemesh",
		105: "CoinTypetsolana",
		106: "CoinTypetusdttrc20",
		107: "CoinTypetbsc",
		108: "CoinTypettron",
	}
	CoinType_value = map[string]int32{
		"CoinTypeUnKnow":     0,
		"CoinTypefilecoin":   1,
		"CoinTypebitcoin":    2,
		"CoinTypeethereum":   3,
		"CoinTypeusdterc20":  4,
		"CoinTypespacemesh":  5,
		"CoinTypesolana":     6,
		"CoinTypeusdttrc20":  7,
		"CoinTypebsc":        8,
		"CoinTypetron":       9,
		"CoinTypetfilecoin":  100,
		"CoinTypetbitcoin":   101,
		"CoinTypetethereum":  102,
		"CoinTypetusdterc20": 103,
		"CoinTypetspacemesh": 104,
		"CoinTypetsolana":    105,
		"CoinTypetusdttrc20": 106,
		"CoinTypetbsc":       107,
		"CoinTypettron":      108,
	}
)

func (x CoinType) Enum() *CoinType {
	p := new(CoinType)
	*p = x
	return p
}

func (x CoinType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoinType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_sphinxplugin_sphinxplugin_proto_enumTypes[0].Descriptor()
}

func (CoinType) Type() protoreflect.EnumType {
	return &file_npool_sphinxplugin_sphinxplugin_proto_enumTypes[0]
}

func (x CoinType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoinType.Descriptor instead.
func (CoinType) EnumDescriptor() ([]byte, []int) {
	return file_npool_sphinxplugin_sphinxplugin_proto_rawDescGZIP(), []int{0}
}

type TransactionType int32

const (
	TransactionType_Invalid        TransactionType = 0
	TransactionType_WalletNew      TransactionType = 1 // proxy -> sign
	TransactionType_TransactionNew TransactionType = 2
	TransactionType_Sign           TransactionType = 3 // proxy -> sign
	TransactionType_Balance        TransactionType = 4 // proxy -> plugin
	TransactionType_PreSign        TransactionType = 5 // proxy -> pluign get nonce
	TransactionType_Broadcast      TransactionType = 6 // proxy -> plugin mpool push
	TransactionType_RegisterCoin   TransactionType = 7 // plugin -> proxy
	TransactionType_SyncMsgState   TransactionType = 8 // plugin -> proxy
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "Invalid",
		1: "WalletNew",
		2: "TransactionNew",
		3: "Sign",
		4: "Balance",
		5: "PreSign",
		6: "Broadcast",
		7: "RegisterCoin",
		8: "SyncMsgState",
	}
	TransactionType_value = map[string]int32{
		"Invalid":        0,
		"WalletNew":      1,
		"TransactionNew": 2,
		"Sign":           3,
		"Balance":        4,
		"PreSign":        5,
		"Broadcast":      6,
		"RegisterCoin":   7,
		"SyncMsgState":   8,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_sphinxplugin_sphinxplugin_proto_enumTypes[1].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_npool_sphinxplugin_sphinxplugin_proto_enumTypes[1]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_npool_sphinxplugin_sphinxplugin_proto_rawDescGZIP(), []int{1}
}

// fil
type UnsignedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint64  `protobuf:"varint,100,opt,name=Version,proto3" json:"Version,omitempty"`
	To      string  `protobuf:"bytes,110,opt,name=To,proto3" json:"To,omitempty"`
	From    string  `protobuf:"bytes,120,opt,name=From,proto3" json:"From,omitempty"`
	Value   float64 `protobuf:"fixed64,130,opt,name=Value,proto3" json:"Value,omitempty"`
	// fil/eth/erc20/bsc
	Nonce    uint64 `protobuf:"varint,140,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	GasLimit int64  `protobuf:"varint,150,opt,name=GasLimit,proto3" json:"GasLimit,omitempty"`
	// eth/erc20/bsc
	GasPrice   int64  `protobuf:"varint,151,opt,name=GasPrice,proto3" json:"GasPrice,omitempty"`
	ChainID    int64  `protobuf:"varint,152,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	ContractID string `protobuf:"bytes,153,opt,name=ContractID,proto3" json:"ContractID,omitempty"`
	// fil
	GasFeeCap  uint64 `protobuf:"varint,160,opt,name=GasFeeCap,proto3" json:"GasFeeCap,omitempty"`
	GasPremium uint64 `protobuf:"varint,170,opt,name=GasPremium,proto3" json:"GasPremium,omitempty"`
	Method     uint64 `protobuf:"varint,180,opt,name=Method,proto3" json:"Method,omitempty"`
	Params     []byte `protobuf:"bytes,190,opt,name=Params,proto3" json:"Params,omitempty"`
	// btc
	Unspent []*Unspent `protobuf:"bytes,200,rep,name=Unspent,proto3" json:"Unspent,omitempty"`
	// sol
	RecentBhash string `protobuf:"bytes,210,opt,name=RecentBhash,proto3" json:"RecentBhash,omitempty"`
	// trc20
	TxData []byte `protobuf:"bytes,220,opt,name=TxData,proto3" json:"TxData,omitempty"`
}

func (x *UnsignedMessage) Reset() {
	*x = UnsignedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsignedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsignedMessage) ProtoMessage() {}

func (x *UnsignedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsignedMessage.ProtoReflect.Descriptor instead.
func (*UnsignedMessage) Descriptor() ([]byte, []int) {
	return file_npool_sphinxplugin_sphinxplugin_proto_rawDescGZIP(), []int{0}
}

func (x *UnsignedMessage) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *UnsignedMessage) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *UnsignedMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *UnsignedMessage) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *UnsignedMessage) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *UnsignedMessage) GetGasLimit() int64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *UnsignedMessage) GetGasPrice() int64 {
	if x != nil {
		return x.GasPrice
	}
	return 0
}

func (x *UnsignedMessage) GetChainID() int64 {
	if x != nil {
		return x.ChainID
	}
	return 0
}

func (x *UnsignedMessage) GetContractID() string {
	if x != nil {
		return x.ContractID
	}
	return ""
}

func (x *UnsignedMessage) GetGasFeeCap() uint64 {
	if x != nil {
		return x.GasFeeCap
	}
	return 0
}

func (x *UnsignedMessage) GetGasPremium() uint64 {
	if x != nil {
		return x.GasPremium
	}
	return 0
}

func (x *UnsignedMessage) GetMethod() uint64 {
	if x != nil {
		return x.Method
	}
	return 0
}

func (x *UnsignedMessage) GetParams() []byte {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *UnsignedMessage) GetUnspent() []*Unspent {
	if x != nil {
		return x.Unspent
	}
	return nil
}

func (x *UnsignedMessage) GetRecentBhash() string {
	if x != nil {
		return x.RecentBhash
	}
	return ""
}

func (x *UnsignedMessage) GetTxData() []byte {
	if x != nil {
		return x.TxData
	}
	return nil
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignType string `protobuf:"bytes,100,opt,name=SignType,proto3" json:"SignType,omitempty"` // secp256k1
	Data     []byte `protobuf:"bytes,110,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_npool_sphinxplugin_sphinxplugin_proto_rawDescGZIP(), []int{1}
}

func (x *Signature) GetSignType() string {
	if x != nil {
		return x.SignType
	}
	return ""
}

func (x *Signature) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// ============================= btc
type Unspent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxID          string  `protobuf:"bytes,100,opt,name=TxID,proto3" json:"TxID,omitempty"`
	Vout          uint32  `protobuf:"varint,110,opt,name=Vout,proto3" json:"Vout,omitempty"`
	Address       string  `protobuf:"bytes,120,opt,name=Address,proto3" json:"Address,omitempty"`
	Account       string  `protobuf:"bytes,130,opt,name=Account,proto3" json:"Account,omitempty"`
	ScriptPubKey  string  `protobuf:"bytes,140,opt,name=ScriptPubKey,proto3" json:"ScriptPubKey,omitempty"`
	RedeemScript  string  `protobuf:"bytes,150,opt,name=RedeemScript,proto3" json:"RedeemScript,omitempty"`
	Amount        float64 `protobuf:"fixed64,160,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Confirmations int64   `protobuf:"varint,170,opt,name=Confirmations,proto3" json:"Confirmations,omitempty"`
	Spendable     bool    `protobuf:"varint,180,opt,name=Spendable,proto3" json:"Spendable,omitempty"`
}

func (x *Unspent) Reset() {
	*x = Unspent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unspent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unspent) ProtoMessage() {}

func (x *Unspent) ProtoReflect() protoreflect.Message {
	mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unspent.ProtoReflect.Descriptor instead.
func (*Unspent) Descriptor() ([]byte, []int) {
	return file_npool_sphinxplugin_sphinxplugin_proto_rawDescGZIP(), []int{2}
}

func (x *Unspent) GetTxID() string {
	if x != nil {
		return x.TxID
	}
	return ""
}

func (x *Unspent) GetVout() uint32 {
	if x != nil {
		return x.Vout
	}
	return 0
}

func (x *Unspent) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Unspent) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Unspent) GetScriptPubKey() string {
	if x != nil {
		return x.ScriptPubKey
	}
	return ""
}

func (x *Unspent) GetRedeemScript() string {
	if x != nil {
		return x.RedeemScript
	}
	return ""
}

func (x *Unspent) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Unspent) GetConfirmations() int64 {
	if x != nil {
		return x.Confirmations
	}
	return 0
}

func (x *Unspent) GetSpendable() bool {
	if x != nil {
		return x.Spendable
	}
	return false
}

type MsgTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  int32    `protobuf:"varint,100,opt,name=Version,proto3" json:"Version,omitempty"`
	TxIn     []*TxIn  `protobuf:"bytes,110,rep,name=TxIn,proto3" json:"TxIn,omitempty"`
	TxOut    []*TxOut `protobuf:"bytes,120,rep,name=TxOut,proto3" json:"TxOut,omitempty"`
	LockTime uint32   `protobuf:"varint,130,opt,name=LockTime,proto3" json:"LockTime,omitempty"`
}

func (x *MsgTx) Reset() {
	*x = MsgTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgTx) ProtoMessage() {}

func (x *MsgTx) ProtoReflect() protoreflect.Message {
	mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgTx.ProtoReflect.Descriptor instead.
func (*MsgTx) Descriptor() ([]byte, []int) {
	return file_npool_sphinxplugin_sphinxplugin_proto_rawDescGZIP(), []int{3}
}

func (x *MsgTx) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *MsgTx) GetTxIn() []*TxIn {
	if x != nil {
		return x.TxIn
	}
	return nil
}

func (x *MsgTx) GetTxOut() []*TxOut {
	if x != nil {
		return x.TxOut
	}
	return nil
}

func (x *MsgTx) GetLockTime() uint32 {
	if x != nil {
		return x.LockTime
	}
	return 0
}

type TxIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousOutPoint *OutPoint `protobuf:"bytes,100,opt,name=PreviousOutPoint,proto3" json:"PreviousOutPoint,omitempty"`
	SignatureScript  []byte    `protobuf:"bytes,110,opt,name=SignatureScript,proto3" json:"SignatureScript,omitempty"`
	Witness          [][]byte  `protobuf:"bytes,120,rep,name=Witness,proto3" json:"Witness,omitempty"`
	Sequence         uint32    `protobuf:"varint,130,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
}

func (x *TxIn) Reset() {
	*x = TxIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxIn) ProtoMessage() {}

func (x *TxIn) ProtoReflect() protoreflect.Message {
	mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxIn.ProtoReflect.Descriptor instead.
func (*TxIn) Descriptor() ([]byte, []int) {
	return file_npool_sphinxplugin_sphinxplugin_proto_rawDescGZIP(), []int{4}
}

func (x *TxIn) GetPreviousOutPoint() *OutPoint {
	if x != nil {
		return x.PreviousOutPoint
	}
	return nil
}

func (x *TxIn) GetSignatureScript() []byte {
	if x != nil {
		return x.SignatureScript
	}
	return nil
}

func (x *TxIn) GetWitness() [][]byte {
	if x != nil {
		return x.Witness
	}
	return nil
}

func (x *TxIn) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

type OutPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// byte len 32
	Hash  []byte `protobuf:"bytes,100,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Index uint32 `protobuf:"varint,110,opt,name=Index,proto3" json:"Index,omitempty"`
}

func (x *OutPoint) Reset() {
	*x = OutPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutPoint) ProtoMessage() {}

func (x *OutPoint) ProtoReflect() protoreflect.Message {
	mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutPoint.ProtoReflect.Descriptor instead.
func (*OutPoint) Descriptor() ([]byte, []int) {
	return file_npool_sphinxplugin_sphinxplugin_proto_rawDescGZIP(), []int{5}
}

func (x *OutPoint) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *OutPoint) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type TxOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value    int64  `protobuf:"varint,100,opt,name=Value,proto3" json:"Value,omitempty"`
	PkScript []byte `protobuf:"bytes,110,opt,name=PkScript,proto3" json:"PkScript,omitempty"`
}

func (x *TxOut) Reset() {
	*x = TxOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxOut) ProtoMessage() {}

func (x *TxOut) ProtoReflect() protoreflect.Message {
	mi := &file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxOut.ProtoReflect.Descriptor instead.
func (*TxOut) Descriptor() ([]byte, []int) {
	return file_npool_sphinxplugin_sphinxplugin_proto_rawDescGZIP(), []int{6}
}

func (x *TxOut) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TxOut) GetPkScript() []byte {
	if x != nil {
		return x.PkScript
	}
	return nil
}

var File_npool_sphinxplugin_sphinxplugin_proto protoreflect.FileDescriptor

var file_npool_sphinxplugin_sphinxplugin_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2f, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xd7, 0x03, 0x0a, 0x0f, 0x55, 0x6e,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x6e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18,
	0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x15, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x15, 0x0a, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x8c, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x47, 0x61, 0x73,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x47, 0x61,
	0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x08, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x97, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x47, 0x61, 0x73, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x98,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x1f,
	0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x18, 0x99, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x12,
	0x1d, 0x0a, 0x09, 0x47, 0x61, 0x73, 0x46, 0x65, 0x65, 0x43, 0x61, 0x70, 0x18, 0xa0, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x47, 0x61, 0x73, 0x46, 0x65, 0x65, 0x43, 0x61, 0x70, 0x12, 0x1f,
	0x0a, 0x0a, 0x47, 0x61, 0x73, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x18, 0xaa, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x47, 0x61, 0x73, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x12,
	0x17, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0xb4, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x17, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0xbe, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x34, 0x0a, 0x07, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x18, 0xc8, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x55, 0x6e, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x65, 0x6e,
	0x74, 0x42, 0x68, 0x61, 0x73, 0x68, 0x18, 0xd2, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x68, 0x61, 0x73, 0x68, 0x12, 0x17, 0x0a, 0x06, 0x54, 0x78,
	0x44, 0x61, 0x74, 0x61, 0x18, 0xdc, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x54, 0x78, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x3b, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x8f, 0x02, 0x0a, 0x07, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x78, 0x49, 0x44, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x78, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x56, 0x6f, 0x75, 0x74, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x56, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19,
	0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0c, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x23,
	0x0a, 0x0c, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x96,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0xa0, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xaa, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x09, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0xb4, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x61, 0x62,
	0x6c, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x54, 0x78, 0x12, 0x18, 0x0a, 0x07,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x54, 0x78, 0x49, 0x6e, 0x18, 0x6e,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x49, 0x6e, 0x52, 0x04, 0x54, 0x78,
	0x49, 0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x54, 0x78, 0x4f, 0x75, 0x74, 0x18, 0x78, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x4f, 0x75, 0x74, 0x52, 0x05, 0x54, 0x78, 0x4f, 0x75,
	0x74, 0x12, 0x1b, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x82, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xaf,
	0x01, 0x0a, 0x04, 0x54, 0x78, 0x49, 0x6e, 0x12, 0x46, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x10, 0x50,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x28, 0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x18, 0x78, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x57, 0x69, 0x74, 0x6e,
	0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x08, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x82, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x22, 0x34, 0x0a, 0x08, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x39, 0x0a, 0x05, 0x54, 0x78, 0x4f, 0x75, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6b, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x50, 0x6b, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x2a, 0xa1, 0x03, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x0e, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x4b, 0x6e, 0x6f, 0x77,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x66, 0x69,
	0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x10, 0x02, 0x12, 0x14, 0x0a,
	0x10, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x75,
	0x73, 0x64, 0x74, 0x65, 0x72, 0x63, 0x32, 0x30, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x10,
	0x05, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x6f, 0x6c,
	0x61, 0x6e, 0x61, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x75, 0x73, 0x64, 0x74, 0x74, 0x72, 0x63, 0x32, 0x30, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x62, 0x73, 0x63, 0x10, 0x08, 0x12, 0x10, 0x0a,
	0x0c, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x74, 0x72, 0x6f, 0x6e, 0x10, 0x09, 0x12,
	0x15, 0x0a, 0x11, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x74, 0x66, 0x69, 0x6c, 0x65,
	0x63, 0x6f, 0x69, 0x6e, 0x10, 0x64, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x74, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x10, 0x65, 0x12, 0x15, 0x0a, 0x11,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x74, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x10, 0x66, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x74,
	0x75, 0x73, 0x64, 0x74, 0x65, 0x72, 0x63, 0x32, 0x30, 0x10, 0x67, 0x12, 0x16, 0x0a, 0x12, 0x43,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x74, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x10, 0x68, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x74,
	0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x10, 0x69, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x74, 0x75, 0x73, 0x64, 0x74, 0x74, 0x72, 0x63, 0x32, 0x30, 0x10, 0x6a,
	0x12, 0x10, 0x0a, 0x0c, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x74, 0x62, 0x73, 0x63,
	0x10, 0x6b, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x74, 0x74,
	0x72, 0x6f, 0x6e, 0x10, 0x6c, 0x2a, 0x98, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x4e, 0x65, 0x77, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x65, 0x77, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x69, 0x67,
	0x6e, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x04,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x10, 0x05, 0x12, 0x0d, 0x0a,
	0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e, 0x10, 0x07, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x08,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e,
	0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x70, 0x68, 0x69, 0x6e,
	0x78, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_sphinxplugin_sphinxplugin_proto_rawDescOnce sync.Once
	file_npool_sphinxplugin_sphinxplugin_proto_rawDescData = file_npool_sphinxplugin_sphinxplugin_proto_rawDesc
)

func file_npool_sphinxplugin_sphinxplugin_proto_rawDescGZIP() []byte {
	file_npool_sphinxplugin_sphinxplugin_proto_rawDescOnce.Do(func() {
		file_npool_sphinxplugin_sphinxplugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_sphinxplugin_sphinxplugin_proto_rawDescData)
	})
	return file_npool_sphinxplugin_sphinxplugin_proto_rawDescData
}

var file_npool_sphinxplugin_sphinxplugin_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_npool_sphinxplugin_sphinxplugin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_sphinxplugin_sphinxplugin_proto_goTypes = []interface{}{
	(CoinType)(0),           // 0: sphinx.plugin.v1.CoinType
	(TransactionType)(0),    // 1: sphinx.plugin.v1.TransactionType
	(*UnsignedMessage)(nil), // 2: sphinx.plugin.v1.UnsignedMessage
	(*Signature)(nil),       // 3: sphinx.plugin.v1.Signature
	(*Unspent)(nil),         // 4: sphinx.plugin.v1.Unspent
	(*MsgTx)(nil),           // 5: sphinx.plugin.v1.MsgTx
	(*TxIn)(nil),            // 6: sphinx.plugin.v1.TxIn
	(*OutPoint)(nil),        // 7: sphinx.plugin.v1.OutPoint
	(*TxOut)(nil),           // 8: sphinx.plugin.v1.TxOut
}
var file_npool_sphinxplugin_sphinxplugin_proto_depIdxs = []int32{
	4, // 0: sphinx.plugin.v1.UnsignedMessage.Unspent:type_name -> sphinx.plugin.v1.Unspent
	6, // 1: sphinx.plugin.v1.MsgTx.TxIn:type_name -> sphinx.plugin.v1.TxIn
	8, // 2: sphinx.plugin.v1.MsgTx.TxOut:type_name -> sphinx.plugin.v1.TxOut
	7, // 3: sphinx.plugin.v1.TxIn.PreviousOutPoint:type_name -> sphinx.plugin.v1.OutPoint
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_sphinxplugin_sphinxplugin_proto_init() }
func file_npool_sphinxplugin_sphinxplugin_proto_init() {
	if File_npool_sphinxplugin_sphinxplugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsignedMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unspent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgTx); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutPoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_npool_sphinxplugin_sphinxplugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_sphinxplugin_sphinxplugin_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_sphinxplugin_sphinxplugin_proto_goTypes,
		DependencyIndexes: file_npool_sphinxplugin_sphinxplugin_proto_depIdxs,
		EnumInfos:         file_npool_sphinxplugin_sphinxplugin_proto_enumTypes,
		MessageInfos:      file_npool_sphinxplugin_sphinxplugin_proto_msgTypes,
	}.Build()
	File_npool_sphinxplugin_sphinxplugin_proto = out.File
	file_npool_sphinxplugin_sphinxplugin_proto_rawDesc = nil
	file_npool_sphinxplugin_sphinxplugin_proto_goTypes = nil
	file_npool_sphinxplugin_sphinxplugin_proto_depIdxs = nil
}
