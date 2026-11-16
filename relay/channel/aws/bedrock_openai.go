package aws

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/relay/channel"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	relayconstant "github.com/QuantumNous/new-api/relay/constant"
	"github.com/QuantumNous/new-api/relaykit/types"
	"github.com/QuantumNous/new-api/setting/model_setting"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/gin-gonic/gin"
)

func newBedrockOpenAIRequest(c *gin.Context, info *relaycommon.RelayInfo, requestBody io.Reader) (*http.Request, error) {
	var path string
	switch {
	case info.RelayMode == relayconstant.RelayModeChatCompletions && info.RelayFormat == types.RelayFormatOpenAI:
		path = "/openai/v1/chat/completions"
	case info.RelayMode == relayconstant.RelayModeResponses &&
		(info.RelayFormat == types.RelayFormatOpenAIResponses || info.RelayFormat == types.RelayFormatOpenAI):
		path = "/openai/v1/responses"
	default:
		return nil, fmt.Errorf("Bedrock OpenAI mode supports Chat Completions and Responses")
	}

	parts := strings.Split(info.ApiKey, "|")
	if len(parts) < 2 || len(parts) > 4 {
		return nil, fmt.Errorf("invalid AWS credentials")
	}
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
		if parts[i] == "" {
			return nil, fmt.Errorf("invalid AWS credentials")
		}
	}
	region := parts[len(parts)-1]
	if strings.ContainsFunc(region, func(ch rune) bool {
		return (ch < 'a' || ch > 'z') && (ch < '0' || ch > '9') && ch != '-'
	}) {
		return nil, fmt.Errorf("invalid AWS region")
	}

	ctx := c.Request.Context()
	// Resolve the official regional endpoint. Channel base URLs and Host
	// overrides cannot redirect AWS credentials to a different host.
	endpoint, err := bedrockruntime.NewDefaultEndpointResolverV2().ResolveEndpoint(ctx, bedrockruntime.EndpointParameters{Region: aws.String(region)})
	if err != nil {
		return nil, fmt.Errorf("resolve Bedrock Runtime endpoint: %w", err)
	}
	endpoint.URI.Path = path

	payload, err := io.ReadAll(requestBody)
	if err != nil {
		return nil, fmt.Errorf("read Bedrock OpenAI request: %w", err)
	}
	var bodyFields map[string]json.RawMessage
	if err := common.Unmarshal(payload, &bodyFields); err != nil || bodyFields == nil {
		return nil, fmt.Errorf("Bedrock OpenAI request must be a JSON object")
	}
	if info.IsModelMapped && !info.IsChannelTest && (info.ChannelSetting.PassThroughBodyEnabled || model_setting.GetGlobalSettings().PassThroughRequestEnabled) {
		bodyFields["model"], err = common.Marshal(info.UpstreamModelName)
		if err != nil {
			return nil, fmt.Errorf("set Bedrock OpenAI model: %w", err)
		}
	}
	if info.RelayMode == relayconstant.RelayModeChatCompletions {
		if legacy, exists := bodyFields["max_tokens"]; exists {
			if _, configured := bodyFields["max_completion_tokens"]; !configured {
				bodyFields["max_completion_tokens"] = legacy
			}
			delete(bodyFields, "max_tokens")
		}
		if info.IsStream {
			bodyFields["stream_options"] = json.RawMessage(`{"include_usage":true}`)
		}
	}
	body, err := common.Marshal(bodyFields)
	if err != nil {
		return nil, fmt.Errorf("encode Bedrock OpenAI request: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.URI.String(), bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create Bedrock OpenAI request: %w", err)
	}
	channel.SetupApiRequestHeader(info, c, &req.Header)
	overrides, err := channel.ResolveHeaderOverride(info, c)
	if err != nil {
		return nil, err
	}
	for key, value := range overrides {
		if strings.EqualFold(key, "Host") || strings.EqualFold(key, "Authorization") || strings.HasPrefix(strings.ToLower(key), "x-amz-") {
			continue
		}
		req.Header.Set(key, value)
	}
	for _, connection := range req.Header.Values("Connection") {
		for name := range strings.SplitSeq(connection, ",") {
			req.Header.Del(strings.TrimSpace(name))
		}
	}
	for _, name := range []string{
		"Connection", "Proxy-Connection", "Keep-Alive", "Proxy-Authenticate",
		"Proxy-Authorization", "Te", "Trailer", "Transfer-Encoding", "Upgrade", "X-Api-Key",
	} {
		req.Header.Del(name)
	}
	req.Header.Set("Content-Type", "application/json")
	if len(parts) == 2 {
		req.Header.Set("Authorization", "Bearer "+parts[0])
		return req, nil
	}
	credentials := aws.Credentials{AccessKeyID: parts[0], SecretAccessKey: parts[1]}
	if len(parts) == 4 {
		credentials.SessionToken = parts[2]
	}
	hash := sha256.Sum256(body)
	if err := v4.NewSigner().SignHTTP(ctx, credentials, req, hex.EncodeToString(hash[:]), "bedrock", region, time.Now()); err != nil {
		return nil, fmt.Errorf("sign Bedrock OpenAI request: %w", err)
	}
	return req, nil
}
