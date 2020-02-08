package publice

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"lzh/models"
	"net/http"
	"net/url"
	"strconv"
)

var transfers = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
var KEY = "taizhoushibogewangluokejibergnet"

//付款订单
type WithdrawOrder struct {
	XMLName        xml.Name `xml:"xml"`
	MchAppid       string   `xml:"mch_appid"`
	Mchid          string   `xml:"mchid"`
	DeviceInfo     string   `xml:"device_info"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	PartnerTradeNo string   `xml:"partner_trade_no"`
	Openid         string   `xml:"openid"`
	CheckName      string   `xml:"check_name"`
	Amount         int      `xml:"amount"`
	Desc           string   `xml:"desc"`
	SpbillCreateIp string   `xml:"spbill_create_ip"`
}

//付款订单结果
type WithdrawResult struct {
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	ResultCode     string `xml:"result_code"`
	ErrCodeDes     string `xml:"err_code_des"`
	PaymentNo      string `xml:"payment_no"`
	PartnerTradeNo string `xml:"partner_trade_no"`
}

//付款，成功返回自定义订单号，微信订单号，true，失败返回错误信息，false
func WithdrawMoney(appid, mchid, openid, amount, partnerTradeNo, desc, clientIp string) (string, string, bool) {
	order := WithdrawOrder{}
	order.MchAppid = appid
	order.Mchid = mchid
	order.Openid = openid
	order.Amount, _ = strconv.Atoi(amount)
	order.Desc = desc
	order.PartnerTradeNo = partnerTradeNo
	order.DeviceInfo = "WEB"
	order.CheckName = "NO_CHECK" //NO_CHECK：不校验真实姓名 FORCE_CHECK：强校验真实姓名
	order.SpbillCreateIp = clientIp
	order.NonceStr = BuildRandNumber(6) + BuildRandNumber(6) + BuildRandNumber(6) + BuildRandNumber(6) + BuildRandNumber(6)
	order.Sign = md5WithdrawOrder(order)
	xmlBody, _ := xml.MarshalIndent(order, " ", " ")
	resp, err := SecurePost(transfers, xmlBody)
	if err != nil {
		return err.Error(), "", false
	}
	defer resp.Body.Close()
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	var res WithdrawResult
	xmlerr := xml.Unmarshal(bodyByte, &res)
	if xmlerr != nil {
		return xmlerr.Error(), "", false
	}
	if res.ReturnCode == "SUCCESS" && res.ResultCode == "SUCCESS" {
		return res.PartnerTradeNo, res.PaymentNo, true
	}
	return res.ReturnMsg, res.ErrCodeDes, false
}

//md5签名
func md5WithdrawOrder(order WithdrawOrder) string {
	o := url.Values{}
	o.Add("mch_appid", order.MchAppid)
	o.Add("mchid", order.Mchid)
	o.Add("device_info", order.DeviceInfo)
	o.Add("partner_trade_no", order.PartnerTradeNo)
	o.Add("check_name", order.CheckName)
	o.Add("amount", strconv.Itoa(order.Amount))
	o.Add("spbill_create_ip", order.SpbillCreateIp)
	o.Add("desc", order.Desc)
	o.Add("nonce_str", order.NonceStr)
	o.Add("openid", order.Openid)
	r, _ := url.QueryUnescape(o.Encode())
	return models.ToMd5(r + "&key=" + KEY)
}

// 上面需要用到的ca证书操作

//ca证书的位置，需要绝对路径
var (
	wechatCertPath = `./cert/cert.pem`
	wechatKeyPath  = `./cert/key.pem`
	wechatCAPath   = `./cert/rootca.pem`
)
var _tlsConfig *tls.Config

//采用单例模式初始化ca
//func getTLSConfig() (*tls.Config, error) {
//	if _tlsConfig != nil {
//		return _tlsConfig, nil
//	}
//	// load cert
//	cert, err := tls.LoadX509KeyPair(wechatCertPath, wechatKeyPath)
//	if err != nil {
//		return nil, err
//	}
//	// load root ca
//	caData, err := ioutil.ReadFile(wechatCAPath)
//	if err != nil {
//		return nil, err
//	}
//	pool := x509.NewCertPool()
//	pool.AppendCertsFromPEM(caData)
//	_tlsConfig = &tls.Config{
//		Certificates: []tls.Certificate{cert},
//		RootCAs:      pool,
//	}
//	return _tlsConfig, nil
//}

// 取消rootca.pem
func getTLSConfig() (*tls.Config, error) {
	if _tlsConfig != nil {
		return _tlsConfig, nil
	}
	// load cert
	cert, err := tls.LoadX509KeyPair(wechatCertPath, wechatKeyPath)
	if err != nil {
		return nil, err
	}
	_tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	return _tlsConfig, nil
}

//携带ca证书的安全请求
func SecurePost(url string, xmlContent []byte) (*http.Response, error) {
	tlsConfig, err := getTLSConfig()
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: tr}
	return client.Post(
		url,
		"application/xml",
		bytes.NewBuffer(xmlContent))
}
