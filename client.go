package main

import (
	"context"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/junaozun/rpcx-benchmark/pb"
	"github.com/junaozun/rpcx-benchmark/util"
	eclient "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/share"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var (
	helloClient pb.IHelloTestEtcdV3Client
	gvgClient   pb.IGvgTestEtcdV3Client
)

func initClient(serializeType protocol.SerializeType) {
	xclient, err := NewEtcdV3XClient("127.0.0.1:2379", "HelloTest", serializeType)
	if err != nil {
		return
	}
	xclient2, err := NewEtcdV3XClient("127.0.0.1:2379", "GvgTest", serializeType)
	if err != nil {
		return
	}
	helloClient = pb.NewHelloTestEtcdV3Client(xclient)
	gvgClient = pb.NewGvgTestEtcdV3Client(xclient2)
}

type Counter struct {
	Succ        int64 // 成功量
	Fail        int64 // 失败量
	Total       int64 // 总量
	Concurrency int64 // 并发量
	Cost        int64 // 总耗时 ms
}

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:8898", http.DefaultServeMux)
	}()
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	allSerializeType := make([]protocol.SerializeType, 0)
	//allSerializeType = append(allSerializeType, protocol.SerializeType(8)) // gogoprotobuf的服务端代码生成
	//allSerializeType = append(allSerializeType, protocol.SerializeType(7)) // gogoprotobuf,选择这个需要执行一下gogoproto.sh
	allSerializeType = append(allSerializeType, protocol.SerializeType(6)) // jsonCodec
	allSerializeType = append(allSerializeType, protocol.MsgPack)
	allSerializeType = append(allSerializeType, protocol.ProtoBuffer) // protobuf ,选择这个需要执行一下proto.sh
	allSerializeType = append(allSerializeType, protocol.JSON)
	for _, serializeType := range allSerializeType {
		f.SetCellValue("Sheet2", GetExecTableNum(serializeType)+"1", GetSValue(serializeType))
		var total int64
		realSerializeType := serializeType
		if serializeType == protocol.SerializeType(7) {
			realSerializeType = protocol.ProtoBuffer
		}
		if serializeType == protocol.SerializeType(8) {
			realSerializeType = protocol.ProtoBuffer
		}
		initClient(realSerializeType)
		for i := 1; i <= 10; i++ {
			var tps int64
			if serializeType == protocol.SerializeType(8) {
				tps = requestHello(GetSValue(serializeType), 10000000, 1000)
			} else {
				tps = requestGetGVGMap(GetSValue(serializeType), 10000000, 1000)
			}
			total += tps
			f.SetCellValue("Sheet2", "A"+strconv.Itoa(i+1), i)
			f.SetCellValue("Sheet2", GetExecTableNum(serializeType)+strconv.Itoa(i+1), tps)
		}
		f.SetCellValue("Sheet2", "A12", "avg")
		f.SetCellValue("Sheet2", GetExecTableNum(serializeType)+"12", total/10)
	}
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func GetExecTableNum(serializeType protocol.SerializeType) string {
	switch serializeType {
	case protocol.SerializeType(8):
		return "G"
	case protocol.SerializeType(7):
		return "F"
	case protocol.SerializeType(6):
		return "E"
	case protocol.SerializeType(3):
		return "D"
	case protocol.SerializeType(2):
		return "C"
	case protocol.SerializeType(1):
		return "B"
	}
	return ""
}

func GetSValue(serializeType protocol.SerializeType) string {
	switch serializeType {
	case protocol.SerializeType(8):
		return "gogoProtobuf_GenCode"
	case protocol.SerializeType(7):
		return "gogoProtobuf"
	case protocol.SerializeType(6):
		return "jsoniterCodec"
	case protocol.SerializeType(3):
		return "MsgPack"
	case protocol.SerializeType(2):
		return "ProtoBuf"
	case protocol.SerializeType(1):
		return "JSON"
	}
	return ""
}

// concurrency:500
// totalRequests:5000000
// 意味着启500个协程，每个协程发5000000/500=10000个请求
func requestHello(serializeType string, totalRequests int64, concurrency int64) (tps int64) {
	args := &pb.HelloRequest{
		Params: map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
		Url:        "www.shining3d.com/test/test/test",
		UserToken:  "9raksfjaskfj844r34tulwktjskgjs345823459325345ttjdsgfdfkkf340523548235472398",
		UserID:     "jkdjfasdf8wrfhhh3w43849rqhfuf934r8u23490582305423-045923-4-f3f3j4f394urt304tu",
		AppID:      "309481920435234j23kth2k5th3245t234905234523gdrggggg545t43634",
		AppKey:     "326344fd23j4r9234rjd923hr4j923hr4234rh2dr4",
		Ip:         "128.0.0.1",
		ClientID:   "430523ur3f93h20tgh34f00444r3h4rd33",
		CsrfToken:  "30r4234tuj293tjg234u9u2j32895t43895gt9h0dweui9hgsiorhtguwe0h5t3w4gh5hsdorghsoghsdkpghsejgh0w540ghudsgowg",
		TraceID:    "9343204jfoasjfg0935ufj43o5jgsjlgkdrgjslrj999999sfgsgj3p45u8t3295g2jspjgapjgflpgjeirgjhhi2og54950tgjdrsopgjspogjewpgi4jpw5u49gusdogjaospgoiergiuiwe",
		Body:       []byte(`dskafjlkfjwiwrfsakf439406356823048230456830945683045t28340ufr0ufowjehofjo32t4h0324u5023t50234582039852304ti2304tu230tjiu02ut023ut023u50235u234`),
		Header:     []byte(`dskafjlkfjwiwrfsakf439406356823048230456830945683045t28340ufr0ufowjehofjo32t4h0324u5023t50234582039852304ti2304tu230tjiu02ut023ut023u50235u234`),
		UserAgent:  []byte(`dskafjlkfjwiwrfsakf439406356823048230456830945683045t28340ufr0ufowjehofjo32t4h0324u5023t50234582039852304ti2304tu230tjiu02ut023ut023u50235u234`),
		Referer:    []byte(`dskafjlkfjwiwrfsakf439406356823048230456830945683045t28340ufr0ufowjehofjo32t4h0324u5023t50234582039852304ti2304tu230tjiu02ut023ut023u50235u234`),
		Queries:    []byte(`dskafjlkfjwiwrfsakf439406356823048230456830945683045t28340ufr0ufowjehofjo32t4h0324u5023t50234582039852304ti2304tu230tjiu02ut023ut023u50235u234`),
		EncryptAES: true,
	}
	counter := &Counter{
		Total:       totalRequests,
		Concurrency: concurrency,
	}
	perClientReqs := totalRequests / concurrency
	var wg sync.WaitGroup
	wg.Add(int(concurrency))
	startTime := time.Now().UnixNano()
	for i := int64(0); i < concurrency; i++ {

		go func(i int64) {
			for j := int64(0); j < perClientReqs; j++ {

				rsp, err := helloClient.SayHello(context.Background(), args)

				if err == nil && rsp.Result == "suxuefeng" {
					atomic.AddInt64(&counter.Succ, 1)
				} else {
					log.Info("rsp fail : %v", err)
					atomic.AddInt64(&counter.Fail, 1)
				}
			}

			wg.Done()
		}(i)
	}
	wg.Wait()
	counter.Cost = (time.Now().UnixNano() - startTime) / 1000000
	fmt.Printf(serializeType+" took %d ms for %d requests \n", counter.Cost, counter.Total)
	fmt.Printf("sent     requests      : %d\n", counter.Total)
	fmt.Printf("received requests      : %d\n", atomic.LoadInt64(&counter.Succ)+atomic.LoadInt64(&counter.Fail))
	fmt.Printf("received requests succ : %d\n", atomic.LoadInt64(&counter.Succ))
	fmt.Printf("received requests fail : %d\n", atomic.LoadInt64(&counter.Fail))
	tps = totalRequests * 1000 / counter.Cost
	fmt.Printf("throughput  (TPS)      : %d\n", tps)
	return
}

func requestGetGVGMap(serializeType string, totalRequests int64, concurrency int64) (tps int64) {
	args := &pb.HelloRequest{
		Params: map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
		Url:        "www.shining3d.com/test/test/test",
		UserToken:  "9raksfjaskfj844r34tulwktjskgjs345823459325345ttjdsgfdfkkf340523548235472398",
		UserID:     "jkdjfasdf8wrfhhh3w43849rqhfuf934r8u23490582305423-045923-4-f3f3j4f394urt304tu",
		AppID:      "309481920435234j23kth2k5th3245t234905234523gdrggggg545t43634",
		AppKey:     "326344fd23j4r9234rjd923hr4j923hr4234rh2dr4",
		Ip:         "128.0.0.1",
		ClientID:   "430523ur3f93h20tgh34f00444r3h4rd33",
		CsrfToken:  "30r4234tuj293tjg234u9u2j32895t43895gt9h0dweui9hgsiorhtguwe0h5t3w4gh5hsdorghsoghsdkpghsejgh0w540ghudsgowg",
		TraceID:    "9343204jfoasjfg0935ufj43o5jgsjlgkdrgjslrj999999sfgsgj3p45u8t3295g2jspjgapjgflpgjeirgjhhi2og54950tgjdrsopgjspogjewpgi4jpw5u49gusdogjaospgoiergiuiwe",
		Body:       []byte(`dskafjlkfjwiwrfsakf439406356823048230456830945683045t28340ufr0ufowjehofjo32t4h0324u5023t50234582039852304ti2304tu230tjiu02ut023ut023u50235u234`),
		Header:     []byte(`dskafjlkfjwiwrfsakf439406356823048230456830945683045t28340ufr0ufowjehofjo32t4h0324u5023t50234582039852304ti2304tu230tjiu02ut023ut023u50235u234`),
		UserAgent:  []byte(`dskafjlkfjwiwrfsakf439406356823048230456830945683045t28340ufr0ufowjehofjo32t4h0324u5023t50234582039852304ti2304tu230tjiu02ut023ut023u50235u234`),
		Referer:    []byte(`dskafjlkfjwiwrfsakf439406356823048230456830945683045t28340ufr0ufowjehofjo32t4h0324u5023t50234582039852304ti2304tu230tjiu02ut023ut023u50235u234`),
		Queries:    []byte(`dskafjlkfjwiwrfsakf439406356823048230456830945683045t28340ufr0ufowjehofjo32t4h0324u5023t50234582039852304ti2304tu230tjiu02ut023ut023u50235u234`),
		EncryptAES: true,
	}
	counter := &Counter{
		Total:       totalRequests,
		Concurrency: concurrency,
	}
	perClientReqs := totalRequests / concurrency
	var wg sync.WaitGroup
	wg.Add(int(concurrency))
	startTime := time.Now().UnixNano()
	for i := int64(0); i < concurrency; i++ {

		go func(i int64) {
			for j := int64(0); j < perClientReqs; j++ {

				rsp, err := gvgClient.GetMapInfo(context.Background(), args)

				if err == nil && rsp.Result == "suxuefeng" {
					atomic.AddInt64(&counter.Succ, 1)
				} else {
					log.Info("rsp fail : %v", err)
					atomic.AddInt64(&counter.Fail, 1)
				}
			}

			wg.Done()
		}(i)
	}
	wg.Wait()
	counter.Cost = (time.Now().UnixNano() - startTime) / 1000000
	fmt.Printf(serializeType+" took %d ms for %d requests \n", counter.Cost, counter.Total)
	fmt.Printf("sent     requests      : %d\n", counter.Total)
	fmt.Printf("received requests      : %d\n", atomic.LoadInt64(&counter.Succ)+atomic.LoadInt64(&counter.Fail))
	fmt.Printf("received requests succ : %d\n", atomic.LoadInt64(&counter.Succ))
	fmt.Printf("received requests fail : %d\n", atomic.LoadInt64(&counter.Fail))
	tps = totalRequests * 1000 / counter.Cost
	fmt.Printf("throughput  (TPS)      : %d\n", tps)
	return
}

func NewEtcdV3XClient(etcdAddr string, serverPath string, serializeType protocol.SerializeType) (client.XClient, error) {
	d, err := eclient.NewEtcdV3Discovery("/rpcx_test", serverPath, []string{etcdAddr}, true, nil)
	if err != nil {
		return nil, err
	}
	share.Codecs[protocol.SerializeType(6)] = &util.JsoniterCodec{}
	var DefaultOption = client.Option{
		Retries:             3,
		RPCPath:             share.DefaultRPCPath,
		ConnectTimeout:      time.Second,
		SerializeType:       serializeType,
		CompressType:        protocol.None,
		BackupLatency:       10 * time.Millisecond,
		MaxWaitForHeartbeat: 30 * time.Second,
		TCPKeepAlivePeriod:  time.Minute,
		BidirectionalBlock:  false,
	}
	xclient := client.NewXClient(serverPath, client.Failtry, client.RandomSelect, d, DefaultOption)
	return xclient, nil
}
