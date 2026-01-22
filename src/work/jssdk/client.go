package jssdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/jssdk"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/jssdk/response"
)

type Client struct {
	*jssdk.Client
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	jssdkClient, err := jssdk.NewClient(app)
	if err != nil {
		return nil, err
	}
	client := &Client{
		jssdkClient,
	}

	config := app.GetConfig()
	baseURI := config.GetString("http.base_uri", "/")

	client.TicketEndpoint = baseURI + "/cgi-bin/get_jsapi_ticket"

	client.OverrideGetAppID()

	return client, nil
}

func (comp *Client) OverrideGetAppID() {
	comp.GetAppID = func() string {
		config := comp.BaseClient.App.GetConfig()
		return config.GetString("corp_id", "")
	}
}

func (comp *Client) GetAgentConfigArray(
	request *http.Request,
	agentID int,
	url string,
	nonce string,
	timestamp int64,
) (*object.HashMap, error) {

	// url 为空时使用默认
	if url == "" {
		url = comp.GetUrl(request)
	}

	// nonce 为空时生成随机字符串
	if nonce == "" {
		nonce = object.QuickRandom(10)
	}

	// timestamp 为空时使用当前时间
	if timestamp == 0 {
		timestamp = time.Now().Unix()
	}

	// 获取 agent ticket
	ticketInfo, err := comp.GetAgentTicket(request.Context(), agentID, false, "agent_config")
	if err != nil {
		return nil, err
	}
	ticket := ticketInfo["ticket"].(string)

	// 生成签名
	signature := comp.GetTicketSignature(ticket, nonce, timestamp, url)

	return &object.HashMap{
		"corpid":    comp.GetAppID(),
		"agentid":   agentID,
		"nonceStr":  nonce,
		"timestamp": timestamp,
		"url":       url,
		"signature": signature,
	}, nil
}

func (comp *Client) GetTicket(ctx context.Context) (*response2.ResponseGetTicket, error) {
	result := &response2.ResponseGetTicket{}

	params := &object.StringMap{
		"type": "agent_config",
	}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/ticket/get", params, nil, result)

	return result, err
}

func (comp *Client) GetAgentTicket(
	ctx context.Context,
	agentID int,
	refresh bool,
	ticketType string,
) (object.HashMap, error) {

	if ticketType == "" {
		ticketType = "agent_config"
	}

	cacheKey := fmt.Sprintf(
		"powerwechat.work.jssdk.ticket.%d.%s.%s",
		agentID,
		ticketType,
		comp.GetAppID(),
	)

	if !refresh && comp.Cache.Has(cacheKey) {
		value, err := comp.Cache.Get(cacheKey, nil)
		if err == nil {
			if data, ok := value.(map[string]interface{}); ok {
				hashMap := object.HashMap(data)
				return hashMap, nil
			}
		}
	}

	resp := object.HashMap{}

	_, err := comp.BaseClient.RequestRaw(
		ctx,
		"cgi-bin/ticket/get",
		"GET",
		&object.HashMap{
			"query": &object.StringMap{
				"type": ticketType,
			},
		},
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}

	if resp.Get("errcode").(float64) != 0 {
		return nil, errors.New(resp.Get("errmsg").(string))
	}

	expiresIn := int(resp["expires_in"].(float64))
	ttl := time.Duration(expiresIn-500) * time.Second

	_ = comp.Cache.Set(cacheKey, resp, ttl)

	if !comp.Cache.Has(cacheKey) {
		return nil, errors.New("failed to cache jssdk ticket")
	}

	return resp, nil
}
