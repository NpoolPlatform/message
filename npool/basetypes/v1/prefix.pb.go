// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/basetypes/v1/prefix.proto

package v1

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

type Prefix int32

const (
	Prefix_DefaultPrefix                           Prefix = 0
	Prefix_PrefixUserCode                          Prefix = 10
	Prefix_PrefixUserAccount                       Prefix = 20
	Prefix_PrefixUserLogin                         Prefix = 30
	Prefix_PrefixWalletAccountLock                 Prefix = 40
	Prefix_PrefixAPIRegister                       Prefix = 50
	Prefix_PrefixCreateCoin                        Prefix = 60
	Prefix_PrefixCreateAppCoinDescription          Prefix = 70
	Prefix_PrefixCreateAppCoin                     Prefix = 80
	Prefix_PrefixSetFiat                           Prefix = 90
	Prefix_PrefixCreateFiatCurrencyFeed            Prefix = 100
	Prefix_PrefixCreateFiatCurrency                Prefix = 110
	Prefix_PrefixCreateCoinCurrencyFeed            Prefix = 120
	Prefix_PrefixCreateCoinCurrency                Prefix = 130
	Prefix_PrefixUpdateCoinCurrency                Prefix = 131
	Prefix_PrefixCreateCoinFiat                    Prefix = 140
	Prefix_PrefixCreateCoinFiatCurrency            Prefix = 150
	Prefix_PrefixCreateAppCountry                  Prefix = 160
	Prefix_PrefixCreateCountry                     Prefix = 170
	Prefix_PrefixCreateAppLang                     Prefix = 180
	Prefix_PrefixCreateLang                        Prefix = 190
	Prefix_PrefixCreateMessage                     Prefix = 200
	Prefix_PrefixCreateNotifUser                   Prefix = 210
	Prefix_PrefixCreateNotif                       Prefix = 220
	Prefix_PrefixCreateEmailTemplate               Prefix = 230
	Prefix_PrefixCreateFrontendTemplate            Prefix = 240
	Prefix_PrefixCreateSMSTemplate                 Prefix = 250
	Prefix_PrefixCreateApp                         Prefix = 260
	Prefix_PrefixCreateUser                        Prefix = 270
	Prefix_PrefixCreateSubscriber                  Prefix = 280
	Prefix_PrefixCreateAppSubscribe                Prefix = 290
	Prefix_PrefixCreateAuth                        Prefix = 300
	Prefix_PrefixCreateKyc                         Prefix = 310
	Prefix_PrefixCreateRole                        Prefix = 320
	Prefix_PrefixCreateRoleUser                    Prefix = 330
	Prefix_PrefixCreateDepositAccount              Prefix = 340
	Prefix_PrefixCreateGoodBenefitAccount          Prefix = 350
	Prefix_PrefixCreatePaymentAccount              Prefix = 360
	Prefix_PrefixCreateUserTransfer                Prefix = 370
	Prefix_PrefixCreateUserAccount                 Prefix = 380
	Prefix_PrefixCreatePlatformAccount             Prefix = 390
	Prefix_PrefixCreateOAuthThirdParty             Prefix = 400
	Prefix_PrefixCreateAppOAuthThirdParty          Prefix = 410
	Prefix_PrefixOAuthLogin                        Prefix = 420
	Prefix_PrefixCreateAppEvent                    Prefix = 430
	Prefix_PrefixCreateInvitationCode              Prefix = 440
	Prefix_PrefixCreateRegistration                Prefix = 450
	Prefix_PrefixCreateCommission                  Prefix = 460
	Prefix_PrefixCloneCommission                   Prefix = 470
	Prefix_PrefixCreateInspireAchievementStatement Prefix = 480
	Prefix_PrefixCreateInspireAchievement          Prefix = 490
	Prefix_PrefixCreateDeviceInfo                  Prefix = 500
	Prefix_PrefixCreateVendorBrand                 Prefix = 510
	Prefix_PrefixCreateVendorLocation              Prefix = 520
	Prefix_PrefixCreateGood                        Prefix = 530
	Prefix_PrefixCreateRequiredGood                Prefix = 540
	Prefix_PrefixLikeGood                          Prefix = 550
	Prefix_PrefixCommentGood                       Prefix = 560
	Prefix_PrefixRecommendGood                     Prefix = 570
	Prefix_PrefixScoreGood                         Prefix = 580
	Prefix_PrefixCreateAppGood                     Prefix = 590
	Prefix_PrefixCreateAppDefaultGood              Prefix = 600
	Prefix_PrefixCreateTopMost                     Prefix = 610
	Prefix_PrefixCreateTopMostGood                 Prefix = 620
	Prefix_PrefixCreateLedger                      Prefix = 700
	Prefix_PrefixCreateLedgerStatement             Prefix = 710
	Prefix_PrefixCreateLedgerProfit                Prefix = 720
	Prefix_PrefixCreateGoodLedger                  Prefix = 730
	Prefix_PrefixCreateGoodLedgerStatement         Prefix = 740
	Prefix_PrefixCreateGoodLedgerUnsoldStatement   Prefix = 750
	Prefix_PrefixMigrate                           Prefix = 10000
	Prefix_PrefixScheduler                         Prefix = 20000
	Prefix_PrefixAccountLock                       Prefix = 30000
	Prefix_PrefixCreateOrder                       Prefix = 40000
	Prefix_PrefixCreateWithdraw                    Prefix = 40010
)

// Enum value maps for Prefix.
var (
	Prefix_name = map[int32]string{
		0:     "DefaultPrefix",
		10:    "PrefixUserCode",
		20:    "PrefixUserAccount",
		30:    "PrefixUserLogin",
		40:    "PrefixWalletAccountLock",
		50:    "PrefixAPIRegister",
		60:    "PrefixCreateCoin",
		70:    "PrefixCreateAppCoinDescription",
		80:    "PrefixCreateAppCoin",
		90:    "PrefixSetFiat",
		100:   "PrefixCreateFiatCurrencyFeed",
		110:   "PrefixCreateFiatCurrency",
		120:   "PrefixCreateCoinCurrencyFeed",
		130:   "PrefixCreateCoinCurrency",
		131:   "PrefixUpdateCoinCurrency",
		140:   "PrefixCreateCoinFiat",
		150:   "PrefixCreateCoinFiatCurrency",
		160:   "PrefixCreateAppCountry",
		170:   "PrefixCreateCountry",
		180:   "PrefixCreateAppLang",
		190:   "PrefixCreateLang",
		200:   "PrefixCreateMessage",
		210:   "PrefixCreateNotifUser",
		220:   "PrefixCreateNotif",
		230:   "PrefixCreateEmailTemplate",
		240:   "PrefixCreateFrontendTemplate",
		250:   "PrefixCreateSMSTemplate",
		260:   "PrefixCreateApp",
		270:   "PrefixCreateUser",
		280:   "PrefixCreateSubscriber",
		290:   "PrefixCreateAppSubscribe",
		300:   "PrefixCreateAuth",
		310:   "PrefixCreateKyc",
		320:   "PrefixCreateRole",
		330:   "PrefixCreateRoleUser",
		340:   "PrefixCreateDepositAccount",
		350:   "PrefixCreateGoodBenefitAccount",
		360:   "PrefixCreatePaymentAccount",
		370:   "PrefixCreateUserTransfer",
		380:   "PrefixCreateUserAccount",
		390:   "PrefixCreatePlatformAccount",
		400:   "PrefixCreateOAuthThirdParty",
		410:   "PrefixCreateAppOAuthThirdParty",
		420:   "PrefixOAuthLogin",
		430:   "PrefixCreateAppEvent",
		440:   "PrefixCreateInvitationCode",
		450:   "PrefixCreateRegistration",
		460:   "PrefixCreateCommission",
		470:   "PrefixCloneCommission",
		480:   "PrefixCreateInspireAchievementStatement",
		490:   "PrefixCreateInspireAchievement",
		500:   "PrefixCreateDeviceInfo",
		510:   "PrefixCreateVendorBrand",
		520:   "PrefixCreateVendorLocation",
		530:   "PrefixCreateGood",
		540:   "PrefixCreateRequiredGood",
		550:   "PrefixLikeGood",
		560:   "PrefixCommentGood",
		570:   "PrefixRecommendGood",
		580:   "PrefixScoreGood",
		590:   "PrefixCreateAppGood",
		600:   "PrefixCreateAppDefaultGood",
		610:   "PrefixCreateTopMost",
		620:   "PrefixCreateTopMostGood",
		700:   "PrefixCreateLedger",
		710:   "PrefixCreateLedgerStatement",
		720:   "PrefixCreateLedgerProfit",
		730:   "PrefixCreateGoodLedger",
		740:   "PrefixCreateGoodLedgerStatement",
		750:   "PrefixCreateGoodLedgerUnsoldStatement",
		10000: "PrefixMigrate",
		20000: "PrefixScheduler",
		30000: "PrefixAccountLock",
		40000: "PrefixCreateOrder",
		40010: "PrefixCreateWithdraw",
	}
	Prefix_value = map[string]int32{
		"DefaultPrefix":                           0,
		"PrefixUserCode":                          10,
		"PrefixUserAccount":                       20,
		"PrefixUserLogin":                         30,
		"PrefixWalletAccountLock":                 40,
		"PrefixAPIRegister":                       50,
		"PrefixCreateCoin":                        60,
		"PrefixCreateAppCoinDescription":          70,
		"PrefixCreateAppCoin":                     80,
		"PrefixSetFiat":                           90,
		"PrefixCreateFiatCurrencyFeed":            100,
		"PrefixCreateFiatCurrency":                110,
		"PrefixCreateCoinCurrencyFeed":            120,
		"PrefixCreateCoinCurrency":                130,
		"PrefixUpdateCoinCurrency":                131,
		"PrefixCreateCoinFiat":                    140,
		"PrefixCreateCoinFiatCurrency":            150,
		"PrefixCreateAppCountry":                  160,
		"PrefixCreateCountry":                     170,
		"PrefixCreateAppLang":                     180,
		"PrefixCreateLang":                        190,
		"PrefixCreateMessage":                     200,
		"PrefixCreateNotifUser":                   210,
		"PrefixCreateNotif":                       220,
		"PrefixCreateEmailTemplate":               230,
		"PrefixCreateFrontendTemplate":            240,
		"PrefixCreateSMSTemplate":                 250,
		"PrefixCreateApp":                         260,
		"PrefixCreateUser":                        270,
		"PrefixCreateSubscriber":                  280,
		"PrefixCreateAppSubscribe":                290,
		"PrefixCreateAuth":                        300,
		"PrefixCreateKyc":                         310,
		"PrefixCreateRole":                        320,
		"PrefixCreateRoleUser":                    330,
		"PrefixCreateDepositAccount":              340,
		"PrefixCreateGoodBenefitAccount":          350,
		"PrefixCreatePaymentAccount":              360,
		"PrefixCreateUserTransfer":                370,
		"PrefixCreateUserAccount":                 380,
		"PrefixCreatePlatformAccount":             390,
		"PrefixCreateOAuthThirdParty":             400,
		"PrefixCreateAppOAuthThirdParty":          410,
		"PrefixOAuthLogin":                        420,
		"PrefixCreateAppEvent":                    430,
		"PrefixCreateInvitationCode":              440,
		"PrefixCreateRegistration":                450,
		"PrefixCreateCommission":                  460,
		"PrefixCloneCommission":                   470,
		"PrefixCreateInspireAchievementStatement": 480,
		"PrefixCreateInspireAchievement":          490,
		"PrefixCreateDeviceInfo":                  500,
		"PrefixCreateVendorBrand":                 510,
		"PrefixCreateVendorLocation":              520,
		"PrefixCreateGood":                        530,
		"PrefixCreateRequiredGood":                540,
		"PrefixLikeGood":                          550,
		"PrefixCommentGood":                       560,
		"PrefixRecommendGood":                     570,
		"PrefixScoreGood":                         580,
		"PrefixCreateAppGood":                     590,
		"PrefixCreateAppDefaultGood":              600,
		"PrefixCreateTopMost":                     610,
		"PrefixCreateTopMostGood":                 620,
		"PrefixCreateLedger":                      700,
		"PrefixCreateLedgerStatement":             710,
		"PrefixCreateLedgerProfit":                720,
		"PrefixCreateGoodLedger":                  730,
		"PrefixCreateGoodLedgerStatement":         740,
		"PrefixCreateGoodLedgerUnsoldStatement":   750,
		"PrefixMigrate":                           10000,
		"PrefixScheduler":                         20000,
		"PrefixAccountLock":                       30000,
		"PrefixCreateOrder":                       40000,
		"PrefixCreateWithdraw":                    40010,
	}
)

func (x Prefix) Enum() *Prefix {
	p := new(Prefix)
	*p = x
	return p
}

func (x Prefix) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Prefix) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_prefix_proto_enumTypes[0].Descriptor()
}

func (Prefix) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_prefix_proto_enumTypes[0]
}

func (x Prefix) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Prefix.Descriptor instead.
func (Prefix) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_prefix_proto_rawDescGZIP(), []int{0}
}

var File_npool_basetypes_v1_prefix_proto protoreflect.FileDescriptor

var file_npool_basetypes_v1_prefix_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2a,
	0xe1, 0x10, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x10,
	0x0a, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x14, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x1e, 0x12, 0x1b, 0x0a,
	0x17, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x10, 0x28, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x41, 0x50, 0x49, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x10,
	0x32, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x69, 0x6e, 0x10, 0x3c, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x69, 0x6e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x46, 0x12, 0x17, 0x0a, 0x13, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x43, 0x6f,
	0x69, 0x6e, 0x10, 0x50, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x53, 0x65,
	0x74, 0x46, 0x69, 0x61, 0x74, 0x10, 0x5a, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x46, 0x65, 0x65, 0x64, 0x10, 0x64, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x10, 0x6e, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x46, 0x65, 0x65, 0x64, 0x10, 0x78, 0x12, 0x1d, 0x0a, 0x18, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x10, 0x82, 0x01, 0x12, 0x1d, 0x0a, 0x18, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x10, 0x83, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x46, 0x69, 0x61, 0x74, 0x10,
	0x8c, 0x01, 0x12, 0x21, 0x0a, 0x1c, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x46, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x10, 0x96, 0x01, 0x12, 0x1b, 0x0a, 0x16, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x10,
	0xa0, 0x01, 0x12, 0x18, 0x0a, 0x13, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x10, 0xaa, 0x01, 0x12, 0x18, 0x0a, 0x13,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x4c,
	0x61, 0x6e, 0x67, 0x10, 0xb4, 0x01, 0x12, 0x15, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x10, 0xbe, 0x01, 0x12, 0x18, 0x0a,
	0x13, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x10, 0xc8, 0x01, 0x12, 0x1a, 0x0a, 0x15, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72,
	0x10, 0xd2, 0x01, 0x12, 0x16, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x10, 0xdc, 0x01, 0x12, 0x1e, 0x0a, 0x19, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x10, 0xe6, 0x01, 0x12, 0x21, 0x0a, 0x1c, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x10, 0xf0, 0x01, 0x12, 0x1c,
	0x0a, 0x17, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x4d,
	0x53, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x10, 0xfa, 0x01, 0x12, 0x14, 0x0a, 0x0f,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x10,
	0x84, 0x02, 0x12, 0x15, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x10, 0x8e, 0x02, 0x12, 0x1b, 0x0a, 0x16, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x10, 0x98, 0x02, 0x12, 0x1d, 0x0a, 0x18, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x10, 0xa2, 0x02, 0x12, 0x15, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x10, 0xac, 0x02, 0x12, 0x14, 0x0a, 0x0f,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x10,
	0xb6, 0x02, 0x12, 0x15, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x10, 0xc0, 0x02, 0x12, 0x19, 0x0a, 0x14, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x10, 0xca, 0x02, 0x12, 0x1f, 0x0a, 0x1a, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x10, 0xd4, 0x02, 0x12, 0x23, 0x0a, 0x1e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0xde, 0x02, 0x12, 0x1f, 0x0a, 0x1a, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0xe8, 0x02, 0x12, 0x1d, 0x0a, 0x18, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x10, 0xf2, 0x02, 0x12, 0x1c, 0x0a, 0x17, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0xfc, 0x02, 0x12, 0x20, 0x0a, 0x1b, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x86, 0x03, 0x12, 0x20, 0x0a, 0x1b, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x10, 0x90, 0x03, 0x12, 0x23, 0x0a, 0x1e,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x10, 0x9a,
	0x03, 0x12, 0x15, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0xa4, 0x03, 0x12, 0x19, 0x0a, 0x14, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x10, 0xae, 0x03, 0x12, 0x1f, 0x0a, 0x1a, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x10, 0xb8, 0x03, 0x12, 0x1d, 0x0a, 0x18, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0xc2, 0x03, 0x12, 0x1b, 0x0a, 0x16, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0xcc, 0x03,
	0x12, 0x1a, 0x0a, 0x15, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0xd6, 0x03, 0x12, 0x2c, 0x0a, 0x27,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0xe0, 0x03, 0x12, 0x23, 0x0a, 0x1e, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0xea, 0x03, 0x12,
	0x1b, 0x0a, 0x16, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0xf4, 0x03, 0x12, 0x1c, 0x0a, 0x17,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x10, 0xfe, 0x03, 0x12, 0x1f, 0x0a, 0x1a, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x88, 0x04, 0x12, 0x15, 0x0a, 0x10, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x10,
	0x92, 0x04, 0x12, 0x1d, 0x0a, 0x18, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x10, 0x9c,
	0x04, 0x12, 0x13, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x69, 0x6b, 0x65, 0x47,
	0x6f, 0x6f, 0x64, 0x10, 0xa6, 0x04, 0x12, 0x16, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x10, 0xb0, 0x04, 0x12, 0x18,
	0x0a, 0x13, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x47, 0x6f, 0x6f, 0x64, 0x10, 0xba, 0x04, 0x12, 0x14, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x10, 0xc4, 0x04, 0x12, 0x18,
	0x0a, 0x13, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x70, 0x47, 0x6f, 0x6f, 0x64, 0x10, 0xce, 0x04, 0x12, 0x1f, 0x0a, 0x1a, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x10, 0xd8, 0x04, 0x12, 0x18, 0x0a, 0x13, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x4d, 0x6f, 0x73, 0x74,
	0x10, 0xe2, 0x04, 0x12, 0x1c, 0x0a, 0x17, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x4d, 0x6f, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x10, 0xec,
	0x04, 0x12, 0x17, 0x0a, 0x12, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x10, 0xbc, 0x05, 0x12, 0x20, 0x0a, 0x1b, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0xc6, 0x05, 0x12, 0x1d, 0x0a, 0x18,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x10, 0xd0, 0x05, 0x12, 0x1b, 0x0a, 0x16, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x4c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x10, 0xda, 0x05, 0x12, 0x24, 0x0a, 0x1f, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0xe4, 0x05, 0x12, 0x2a,
	0x0a, 0x25, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f,
	0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x55, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0xee, 0x05, 0x12, 0x12, 0x0a, 0x0d, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x10, 0x90, 0x4e, 0x12, 0x15,
	0x0a, 0x0f, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x10, 0xa0, 0x9c, 0x01, 0x12, 0x17, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x10, 0xb0, 0xea, 0x01, 0x12, 0x17,
	0x0a, 0x11, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x10, 0xc0, 0xb8, 0x02, 0x12, 0x1a, 0x0a, 0x14, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x10,
	0xca, 0xb8, 0x02, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_npool_basetypes_v1_prefix_proto_rawDescOnce sync.Once
	file_npool_basetypes_v1_prefix_proto_rawDescData = file_npool_basetypes_v1_prefix_proto_rawDesc
)

func file_npool_basetypes_v1_prefix_proto_rawDescGZIP() []byte {
	file_npool_basetypes_v1_prefix_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_v1_prefix_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_v1_prefix_proto_rawDescData)
	})
	return file_npool_basetypes_v1_prefix_proto_rawDescData
}

var file_npool_basetypes_v1_prefix_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_basetypes_v1_prefix_proto_goTypes = []interface{}{
	(Prefix)(0), // 0: basetypes.v1.Prefix
}
var file_npool_basetypes_v1_prefix_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_v1_prefix_proto_init() }
func file_npool_basetypes_v1_prefix_proto_init() {
	if File_npool_basetypes_v1_prefix_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_v1_prefix_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_v1_prefix_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_v1_prefix_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_v1_prefix_proto_enumTypes,
	}.Build()
	File_npool_basetypes_v1_prefix_proto = out.File
	file_npool_basetypes_v1_prefix_proto_rawDesc = nil
	file_npool_basetypes_v1_prefix_proto_goTypes = nil
	file_npool_basetypes_v1_prefix_proto_depIdxs = nil
}
