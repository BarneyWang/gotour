package main

import (
	"fmt"
	"strings"
)

// var c, python, java bool

// type Vertex struct {
// 	X int
// 	Y int
// }

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	// fmt.Println(v1, p, v2, v3)
	// fmt.Println("My favorite number is", rand.Intn(10))
	// fmt.Println(add(1, 2))
	// a, b := swap("hello", "world")
	// fmt.Println(split(100))
	// fmt.Println(a, b)
	// var i int
	// fmt.Println(i, c, python, java)

	// var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f)
	// var m uint
	// fmt.Println(x, y, z, m)

	// v := 1 // 修改这里！
	// fmt.Printf("v is of type %T\n", v)

	// fmt.Println("start")
	// for i := 1; i <= 100; i++ {
	// 	defer fmt.Println(i)
	// }
	var a int = 1
	var b *int = &a
	fmt.Println("a =", a)
	fmt.Println("b=", b)
	// fmt.Println("done")

	// i, j := 42, 2701

	// p := &i         // point to i
	// fmt.Println(*p) // read i through the pointer
	// *p = 21         // set i through the pointer
	// fmt.Println(i)  // see the new value of i

	// p = &j // point to j
	// fmt.Println(*p)
	// *p = *p / 37   // divide j through the pointer
	// fmt.Println(j) // see the new value of j

	// v := Vertex{1, 2}
	// p := &v
	// p.X = 1e9
	// v.X = 1e9
	// fmt.Println(v)

	// defer func() {
	// 	fmt.Println("c")
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("d")
	// }()
	// panic("ex")
	// x := "222"
	// fmt.Printf(testFor(x))
	// fmt.Printf(len(splitString(",")))
}

// func f() {
// 	fmt.Println("a")
// 	panic("ex")
// 	fmt.Println("b")
// 	fmt.Println("f")
// }

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {

	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 6 / 2
	y = sum - 1
	return
}
func testFor(x string) string {
	return x + "1"
}
func splitString(x string) []string {
	return strings.Split("ak-code-phabricator,acc-acid-et2,acc-acs-et2,acc-asm-et2,acc-fs-et2,acc-fs-ivr-et2,acc-kaproxy-et2,acc-metaq,acc-rec-et2,acc-sipcap-et2,acc-statserver-et2,accs-dm,accs-sal,address-server,addresstokenizer,adp,air,aladdin,alc,ali360,alicarecore,alicom-recommend,alipay_notify,allcbu-notify,amdc,amp-message,anywhere,aqcenter,arcticmaster,arcticsmartisan,argus,artisan,authcenter,baoxianmisc,baoxianpricing,baoxianrecom,baoxianrisk,bbq,bicester,birdsnestcenter,buc-acl-hsf,buc-acl-hsf-dp,buc-authcenter,buc-sso,buc-uc-hsf,buy2,buycenter,cachechecker,capacitymanager,carcenter,carts,cartsapi,cartscache,cdc,cdc-dump,chanel,channelcenter,chongzhibiz,chongzhidata,chongzhidispatch,chongzhimisc,chongzhimobile,cnalgo,cncbf,cndps,cnortools-center,cnortools-task,coincore,collina,collinadecryptor,collinafacade,communityuis,connected,consign,consigncenter,corona-server,credit.center,crmcenter,crmroute,ctr,cuntaonetwork,customercenter.interface,cybertron,d4s,daa-traffic-hub,daa-traffic-op,delivery,detail,detailskip,diablo,diamond,dna,doraemon,dosa,durex,e56,easygocore,ecardcpgw,ecardmisccenter,falcon9r,fivemin,flow,fundgateway,fundplatform,geb,gebsupport,gemini,glaucus,globaladdressservice,gns,gsearch,guide,harbor,havanaauth,havanamlogin,headline,hlservice,iforest,industry-tagcenter,industry.center,inventoryhotel,inventoryplatform,iopen-core,itemcenter,itemtools,jiajucenter,jingweiconsole,ju,juathena,jucol,jucube,judetail,juhsf,juseller,langyan,lexus,liaoyuan,lightspeed,loc-vdump,login,logisticscenter,logisticsmarket-center,logisticsmarket-task,luna,macao,macross,malldetail,malldetailskip,mallitemcenter,marketing-center,maybach,mclaren,member-business-center,membercenter,memberplatform,metaq,metaq-ops,metaqtransfer,minerva-base,mix2notify,mix3notify,mix_notify,motormaster,mtask,mtee,mteemodel,mtop,mtopplatform,myju,namesrv,notify,offer_cbu,omni-industry-center,omni-sourcing,onepiece,onlinecs,openair,openbase,orderplatform,palisade,pamirsqogir,pamirssupport,pbservice,pesystem3,phenix-center,pointcenter,polaris,price-service,promotioncenter,promotiontag,protectcenter,qccenter,quicksilver,quotacenter,ratecenter,refundface,refundface2,refundplatform,refundplatform2,relationplatform,relationrecommend,renqun,resourcecenter,resourcecenterops,rmc,robot,robot-portal,s-account,saipan,saleenginesync,scskuconsole,seattle,security-tesla,securitymatrix,shellcenter,shopcenter,sic,sirius,site_service_center,sitesupport,smartisan,stationmobile,stationplatform,subaru,sync,taokeeper,tbusergw,tddl,tmallbuy,tmallequity,tmallincubator,tmallreport,top,topapi,toptim,track-service,trade_notify,trade_sub2_notify,trade_sub_notify,tradecenter,tradedata.center,trademanager,tradeplatform,trippe,trippeunit,triprule,tsplatform,ucc,uicfinal,uiclogin,uicplus,ultima,umid,ump,upp,usa,ussc,vipserver,volvo,wallet,warehouse-resource-center,warehouserouteservice,wdn,widgetpt,wormholesource,wsearch,wsm,wsnssync,wtt-channel-server,wtt-irs,wtt-trade-service,wtt-trade-web,wukong-amorite,wukong-fastproxy,wukong-file,wukong-im-processor,wukong-oauth,wukong-proxy,wukong-push-amorite,wukong-register,wukong-shaka,wukong-syncserver,xflush-cloud,xspace-account,xspace-bss,xspace-payment,xspace-scheduler,yosemite,hadesP0App", ",")
}
