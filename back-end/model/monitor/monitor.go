package monitor

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"data-manager/config"
	"data-manager/exterror"
	"data-manager/model"
	webContext "data-manager/web/context"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"net/http"
	"time"

	pb "data-manager/model/monitor/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	dbNode = "node"
)

type monitorModel struct{}

var DefaultMonitorModel = &monitorModel{}

func (m *monitorModel) New(ctx *webContext.Context, nodeinfo model.Node) (string, error) {

	var opts []grpc.DialOption
	monitorAddr := nodeinfo.MonitorIp + ":" + nodeinfo.MonitorPort
	b, _ := ioutil.ReadFile(config.Config.HttpConf.GetCAPath())
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		logrus.Debug("credentials: failed to append certificates")
	}
	conf := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
		ServerName:         "platone",
	}
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(conf)))
	opts = append(opts, grpc.WithPerRPCCredentials(new(customeCredential)))

	conn, err := grpc.Dial(monitorAddr, opts...)
	if err != nil {
		logrus.Debug("did not connect: %v", err)
		return "", err
	}
	defer conn.Close()

	c := pb.NewNodeClient(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	logrus.Debug(nodeinfo)
	body := &pb.NodeNewRequest{
		Groupid:      nodeinfo.Groupid,
		Chainid:      nodeinfo.Chainid,
		Ip:           nodeinfo.NodeIp,
		Port:         strconv.FormatUint(uint64(nodeinfo.NodePort), 10),
		Rpcport:      nodeinfo.Rpcport,
		Wsport:       nodeinfo.Wsport,
		Dashport:     nodeinfo.Dashport,
		Password:     nodeinfo.Password,
		CreatorEnode: nodeinfo.CreatorEnode,
		Bootnodes:    nodeinfo.Bootnodes,
	}
	r, err := c.New(ctx1, body)
	if err != nil {
		logrus.Debug("could not greet: %v", err)
		return "", err

	}
	log.Printf("Message: %s", r.GetMessage())
	log.Printf("Result: %s", r.GetResult())
	log.Printf("Pubkey: %s", r.GetPubkey())

	filter := bson.M{}
	filter["nodepublickey"] = r.GetPubkey()

	collection := ctx.DBCtx.Collection(dbNode)
	var node model.Node
	node = nodeinfo
	node.ID = primitive.NewObjectID().Hex()

	node.NodePublicKey = r.GetPubkey()

	var res model.Node
	ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err = collection.FindOne(ctxBd, filter).Decode(&res)
	if nil != err {

		if _, err := collection.InsertOne(ctxBd, node); err != nil {
			return "", err
		}
	} else {
		err = exterror.ErrNodeHasRegistered
		return "", err
	}
	registerNodeInfo := &model.NodeOperateRequest{}
	registerNodeInfo.Name = nodeinfo.NodeName
	registerNodeInfo.Status = int32(nodeinfo.Status)
	registerNodeInfo.PublicKey = node.NodePublicKey
	registerNodeInfo.P2pPort = nodeinfo.NodePort

	if nodeinfo.NodeIp == "" {
		registerNodeInfo.ExternalIP = config.Config.HttpConf.GetIp()
		registerNodeInfo.InternalIP = config.Config.HttpConf.GetIp()
	} else {
		registerNodeInfo.ExternalIP = nodeinfo.NodeIp
		registerNodeInfo.InternalIP = nodeinfo.NodeIp
	}

	model.DefaultPlatonecliModel.AddNode(ctx, *registerNodeInfo)
	return node.NodePublicKey, nil
}

func (m *monitorModel) StartNode(ctx *webContext.Context, nodeinfo model.NodeOp) error {
	filter := bson.M{}
	filter["nodepublickey"] = nodeinfo.PublicKey

	collection := ctx.DBCtx.Collection(dbNode)

	var res model.Node
	ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err := collection.FindOne(ctxBd, filter).Decode(&res)

	MonitorAddr := res.MonitorIp + ":" + res.MonitorPort
	var opts []grpc.DialOption

	b, _ := ioutil.ReadFile(config.Config.HttpConf.GetCAPath())
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		logrus.Debug("credentials: failed to append certificates")
		return err

	}
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
		ServerName:         "platone",
	}
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	opts = append(opts, grpc.WithPerRPCCredentials(new(customeCredential)))

	conn, err := grpc.Dial(MonitorAddr, opts...)
	if err != nil {
		logrus.Debug("did not connect: %v", err)
		return err
	}
	defer conn.Close()

	c := pb.NewNodeClient(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Start(ctx1, &pb.NodeStartRequest{Groupid: nodeinfo.Groupid})
	if err != nil {
		logrus.Debug("could not greet: %v", err)
		return err
	}
	log.Printf("Greeting: %s", r.GetMessage())
	return nil
}

func (m *monitorModel) StopNode(ctx *webContext.Context, nodeinfo model.NodeOp) error {
	filter := bson.M{}
	filter["nodepublickey"] = nodeinfo.PublicKey

	collection := ctx.DBCtx.Collection(dbNode)

	var res model.Node
	ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err := collection.FindOne(ctxBd, filter).Decode(&res)

	MonitorAddr := res.MonitorIp + ":" + res.MonitorPort
	var opts []grpc.DialOption

	b, _ := ioutil.ReadFile(config.Config.HttpConf.GetCAPath())
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		logrus.Debug("credentials: failed to append certificates")
		return err

	}
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
		ServerName:         "platone",
	}
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	//use custome auth
	opts = append(opts, grpc.WithPerRPCCredentials(new(customeCredential)))

	conn, err := grpc.Dial(MonitorAddr, opts...)
	if err != nil {
		logrus.Debug("did not connect: %v", err)
		return err

	}
	defer conn.Close()

	c := pb.NewNodeClient(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Stop(ctx1, &pb.NodeStopRequest{Groupid: nodeinfo.Groupid})
	if err != nil {
		logrus.Debug("could not greet: %v", err)
		return err

	}
	log.Printf("Greeting: %s", r.GetMessage())
	return nil
}
func (m *monitorModel) RestartNode(ctx *webContext.Context, nodeInfo model.NodeOp) error {
	filter := bson.M{}
	filter["nodepublickey"] = nodeInfo.PublicKey

	collection := ctx.DBCtx.Collection(dbNode)

	var res model.Node
	ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err := collection.FindOne(ctxBd, filter).Decode(&res)

	MonitorAddr := res.MonitorIp + ":" + res.MonitorPort
	var opts []grpc.DialOption

	b, _ := ioutil.ReadFile(config.Config.HttpConf.GetCAPath())
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		logrus.Debug("credentials: failed to append certificates")
		return err
	}
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
		ServerName:         "platone",
	}
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	//use custome auth
	opts = append(opts, grpc.WithPerRPCCredentials(new(customeCredential)))

	conn, err := grpc.Dial(MonitorAddr, opts...)
	if err != nil {
		logrus.Debug("did not connect: %v", err)
		return err
	}
	defer conn.Close()

	c := pb.NewNodeClient(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ReStart(ctx1, &pb.NodeRestartRequest{Groupid: nodeInfo.Groupid})
	if err != nil {
		logrus.Debug("could not greet: %v", err)
		return err
	}
	log.Printf("Greeting: %s", r.GetMessage())
	return nil
}
