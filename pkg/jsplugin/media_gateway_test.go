package jsplugin

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/QuantumNous/new-api/pkg/billingexpr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMediaGatewayPlugins(t *testing.T) {
	for _, provider := range []string{"buzzy", "flowith"} {
		t.Run(provider, func(t *testing.T) {
			source, err := os.ReadFile("../../plugins/tasks/" + provider + "/plugin.js")
			require.NoError(t, err)
			plugin, err := NewRegistry().Register(string(source), Options{Key: provider, Version: "test"})
			require.NoError(t, err)
			ctx := context.Background()
			videoModel := "grok-imagine"
			if provider == "flowith" {
				videoModel = "veo-3.1-generate"
			}
			decoded, err := plugin.Engine.CallPath(ctx, "protocols", []string{"openai_video", "decodeRequest"}, map[string]any{
				"model": "client-alias", "body": map[string]any{"kind": "json", "value": map[string]any{"prompt": "test"}},
			})
			require.NoError(t, err)
			requestBody := decoded.(map[string]any)["requestBody"].(map[string]any)
			assert.NotContains(t, requestBody, "duration", "decoder must not select defaults before channel mapping")
			videoContext := map[string]any{"model": "client-alias", "upstreamModel": provider + "/" + videoModel, "action": "video", "requestBody": requestBody}
			usage, err := plugin.Engine.Call(ctx, "extractUsage", videoContext)
			require.NoError(t, err)
			facts := usage.(map[string]any)
			videoContext["baseUrl"] = "https://example.com"
			videoContext["apiKey"] = "test-key"
			outbound, err := plugin.Engine.Call(ctx, "buildSubmitRequest", videoContext)
			require.NoError(t, err)
			videoBody := outbound.(map[string]any)["body"].(map[string]any)
			assert.EqualValues(t, facts["seconds"], videoBody["duration"])
			assert.Equal(t, videoModel, videoBody["model"])
			videoContext["usagePurpose"] = "billing_ratios"
			_, err = plugin.Engine.Call(ctx, "extractUsage", videoContext)
			require.ErrorContains(t, err, "per-second pricing")
			delete(videoContext, "usagePurpose")
			if provider == "buzzy" {
				assert.EqualValues(t, 6, facts["seconds"])
				assert.Equal(t, "480p", facts["resolution"])
				assert.Equal(t, "480p", videoBody["resolution"])
			} else {
				assert.EqualValues(t, 4, facts["seconds"])
				assert.Equal(t, "model_default", facts["resolution"])
			}
			for _, duration := range []any{-1, 0, 3601, false, "NaN", 1.5} {
				videoContext["requestBody"] = map[string]any{"prompt": "test", "metadata": map[string]any{"duration": duration}}
				_, err := plugin.Engine.Call(ctx, "extractUsage", videoContext)
				require.Error(t, err, "invalid metadata duration: %v", duration)
			}
			for _, count := range []any{0, -1, 5, 1.5, "2", true} {
				_, err := plugin.Engine.CallPath(ctx, "protocols", []string{"openai_image", "decodeRequest"}, map[string]any{
					"model": "gpt-image-2", "body": map[string]any{"kind": "json", "value": map[string]any{"model": "gpt-image-2", "prompt": "test", "n": count}},
				})
				require.Error(t, err, "invalid image count: %v", count)
			}
			result, err := plugin.Engine.Call(ctx, "buildSubmitRequest", map[string]any{
				"model": "alias", "upstreamModel": "gpt-image-2", "baseUrl": "https://example.com", "apiKey": "test-key",
				"action": "image", "requestBody": map[string]any{"model": "alias", "prompt": "test", "n": 2, "seed": 0, "audio": false},
			})
			require.NoError(t, err)
			request := result.(map[string]any)
			assert.Equal(t, "https://example.com/v1/images/generations", request["url"])
			body := request["body"].(map[string]any)
			assert.Equal(t, "gpt-image-2", body["model"])
			assert.EqualValues(t, 0, body["seed"])
			assert.Equal(t, false, body["audio"])
			for _, data := range []any{
				map[string]any{"url": "https://example.com/image.png"},
				[]any{map[string]any{"url": "https://example.com/image.png"}, map[string]any{"b64_json": "aGVsbG8="}, map[string]any{"revised_prompt": "test"}},
			} {
				usage, err := plugin.Engine.Call(ctx, "extractUsageOnComplete", map[string]any{"action": "image"}, map[string]any{"status": "SUCCESS"}, map[string]any{"data": data})
				require.NoError(t, err)
				assert.EqualValues(t, 1, usage.(map[string]any)["image_count"])
			}
			result, err = plugin.Engine.Call(ctx, "parseSubmitResponse", map[string]any{"action": "image", "publicTaskId": "task-test"}, map[string]any{"body": map[string]any{"data": []any{map[string]any{"revised_prompt": "empty"}}}})
			require.NoError(t, err)
			assert.Equal(t, "FAILURE", result.(map[string]any)["immediate"].(map[string]any)["status"])
		})
	}
}

func TestMediaGatewayImageResolution(t *testing.T) {
	for _, provider := range []string{"buzzy", "flowith"} {
		t.Run(provider, func(t *testing.T) {
			source, err := os.ReadFile("../../plugins/tasks/" + provider + "/plugin.js")
			require.NoError(t, err)
			plugin, err := NewRegistry().Register(string(source), Options{Key: provider})
			require.NoError(t, err)
			models := []string{"gpt-image-2", "gpt-image-2.5-flare", "gpt-image-2.5-sunburst", "Nano Banana 2", "Nano Banana 2.1", "Nano Banana Pro"}
			if provider == "flowith" {
				models = append(models[:3], "gemini-3.1-flash-image", "gemini-nano-banana-2.1", "gemini-3-pro-image")
			}
			for _, model := range models {
				for _, resolution := range []string{"1k", "2k", "4k"} {
					input := map[string]any{"model": "alias", "upstreamModel": provider + "/" + model, "action": "image", "baseUrl": "https://example.com", "apiKey": "test", "requestBody": map[string]any{"prompt": "test", "n": 2, "metadata": map[string]any{"resolution": resolution}}}
					usage, err := plugin.Engine.Call(context.Background(), "extractUsage", input)
					require.NoError(t, err)
					assert.Equal(t, resolution, usage.(map[string]any)["resolution"], model)
					assert.EqualValues(t, 2, usage.(map[string]any)["image_count"])
					rates := []float64{0.02, 0.04, 0.08}
					if strings.HasPrefix(model, "gpt-image") {
						rates = []float64{0.01, 0.05, 0.1}
					} else if strings.Contains(model, "2.1") {
						rates = []float64{0.015, 0.03, 0.06}
					}
					expression := fmt.Sprintf(`u("resolution") == "1k" ? tier("1k", u("image_count") * %g) : u("resolution") == "2k" ? tier("2k", u("image_count") * %g) : tier("4k", u("image_count") * %g)`, rates[0], rates[1], rates[2])
					facts := usage.(map[string]any)
					result, err := billingexpr.ComputeTieredQuotaWithRequest(&billingexpr.BillingSnapshot{ExprString: expression, ExprHash: billingexpr.ExprHashString(expression), GroupRatio: 1, QuotaPerUnit: 500000, ExprVersion: 1, TaskUsageBilling: true}, billingexpr.TokenParams{}, billingexpr.RequestInput{Usage: facts})
					require.NoError(t, err)
					unitPrice := map[string]float64{"1k": rates[0], "2k": rates[1], "4k": rates[2]}[resolution]
					assert.Equal(t, int(unitPrice*2*500000), result.ActualQuotaAfterGroup)
					// 完成回包只覆盖实出张数，已冻结的分辨率继续参与结算。
					complete, err := plugin.Engine.Call(context.Background(), "extractUsageOnComplete", map[string]any{"action": "image"}, map[string]any{}, map[string]any{"data": map[string]any{"url": "https://example.com/image.png"}})
					require.NoError(t, err)
					for k, v := range complete.(map[string]any) {
						facts[k] = v
					}
					result, err = billingexpr.ComputeTieredQuotaWithRequest(&billingexpr.BillingSnapshot{ExprString: expression, ExprHash: billingexpr.ExprHashString(expression), GroupRatio: 2, QuotaPerUnit: 500000, ExprVersion: 1, TaskUsageBilling: true}, billingexpr.TokenParams{}, billingexpr.RequestInput{Usage: facts})
					require.NoError(t, err)
					assert.Equal(t, int(unitPrice*2*500000), result.ActualQuotaAfterGroup)

					request, err := plugin.Engine.Call(context.Background(), "buildSubmitRequest", input)
					require.NoError(t, err)
					body := request.(map[string]any)["body"].(map[string]any)
					if provider == "flowith" && model[:3] == "gpt" {
						assert.Equal(t, map[string]string{"1k": "1024x1024", "2k": "2048x2048", "4k": "2880x2880"}[resolution], body["aspect_ratio"])
					} else if provider == "flowith" {
						assert.Equal(t, resolution, body["image_size"])
					} else {
						assert.Equal(t, resolution, strings.ToLower(body["resolution"].(string)))
					}
				}
			}
			for _, tc := range []struct {
				body map[string]any
				tier string
			}{
				{map[string]any{}, "1k"},
				{map[string]any{"size": "3840x2160"}, "4k"},
				{map[string]any{"image_size": "2K", "resolution": "2k"}, "2k"},
				{map[string]any{"size": "2048x2048", "resolution": "4k"}, ""},
				{map[string]any{"size": "auto"}, ""},
				{map[string]any{"resolution": "8k"}, ""},
				{map[string]any{"resolution": 0}, ""},
				{map[string]any{"aspect_ratio": "3840x2160", "resolution": "1k"}, ""},
				{map[string]any{"size": "1024x1024", "aspect_ratio": "16:9"}, ""},
			} {
				tc.body["prompt"] = "test"
				usage, err := plugin.Engine.Call(context.Background(), "extractUsage", map[string]any{"model": "gpt-image-2", "upstreamModel": provider + "/gpt-image-2", "action": "image", "requestBody": tc.body})
				if tc.tier == "" {
					require.Error(t, err)
					continue
				}
				require.NoError(t, err)
				assert.Equal(t, tc.tier, usage.(map[string]any)["resolution"])
			}
		})
	}
}

func TestMediaGatewayAdditionalImageTiers(t *testing.T) {
	for _, provider := range []string{"buzzy", "flowith"} {
		source, err := os.ReadFile("../../plugins/tasks/" + provider + "/plugin.js")
		require.NoError(t, err)
		plugin, err := NewRegistry().Register(string(source), Options{Key: provider})
		require.NoError(t, err)
		cases := []struct{ model, resolution, want string }{
			{"grok-imagine-image", "1k", "1k"}, {"grok-imagine-image", "2K", "2k"}, {"grok-imagine-image", "4k", ""},
			{"doubao-seedream-4-0-250828", "4k", "4k"}, {"doubao-seedream-4-0-250828", "3k", ""},
			{"doubao-seedream-4-5-251128", "", "2k"}, {"doubao-seedream-4-5-251128", "1k", ""},
			{"doubao-seedream-5-0-260128", "", "2k"}, {"doubao-seedream-5-0-260128", "3k", "3k"}, {"doubao-seedream-5-0-260128", "4k", "4k"},
		}
		if provider == "flowith" {
			cases = cases[:3]
			for i := range cases {
				cases[i].model = "grok-imagine-image-2.0"
			}
		}
		for _, tc := range cases {
			body := map[string]any{"prompt": "test", "n": 2}
			if tc.resolution != "" {
				body["metadata"] = map[string]any{"resolution": tc.resolution}
			}
			ctx := map[string]any{"model": tc.model, "upstreamModel": provider + "/" + tc.model, "action": "image", "baseUrl": "https://example.com", "apiKey": "test", "requestBody": body}
			facts, err := plugin.Engine.Call(context.Background(), "extractUsage", ctx)
			if tc.want == "" {
				require.Error(t, err, tc.model)
				continue
			}
			require.NoError(t, err)
			assert.Equal(t, tc.want, facts.(map[string]any)["resolution"])
			request, err := plugin.Engine.Call(context.Background(), "buildSubmitRequest", ctx)
			require.NoError(t, err)
			out := request.(map[string]any)["body"].(map[string]any)
			field := "resolution"
			if provider == "flowith" {
				field = "image_size"
			}
			assert.Equal(t, tc.want, strings.ToLower(out[field].(string)))
		}
	}
}

func TestMediaGatewayVideoParameters(t *testing.T) {
	for _, provider := range []string{"buzzy", "flowith"} {
		source, err := os.ReadFile("../../plugins/tasks/" + provider + "/plugin.js")
		require.NoError(t, err)
		plugin, err := NewRegistry().Register(string(source), Options{Key: provider})
		require.NoError(t, err)
		model := "Kling 4.0"
		if provider == "flowith" {
			model = "kling-video/v4"
		}
		for _, tc := range []struct {
			name      string
			body      map[string]any
			wantError string
		}{
			{"aspect", map[string]any{"aspect_ratio": "7:3"}, "aspect_ratio"},
			{"size conflict", map[string]any{"size": "9:16", "aspect_ratio": "16:9"}, "agree"},
			{"frame conflict", map[string]any{"start_frame": "https://example.com/a.png", "first_frame": "https://example.com/b.png"}, "agree"},
			{"unsupported enhancement", map[string]any{"enhancement": "4k"}, "not supported"},
			{"unsupported middle frames", map[string]any{"middle_images": []string{"https://example.com/a.png"}}, "not supported"},
			{"reference shape", map[string]any{"reference_images": "https://example.com/a.png"}, "array"},
			{"audio type", map[string]any{"audio": "false"}, "audio"},
		} {
			t.Run(provider+"/"+tc.name, func(t *testing.T) {
				tc.body["prompt"] = "test"
				_, err := plugin.Engine.Call(context.Background(), "extractUsage", map[string]any{"model": model, "upstreamModel": provider + "/" + model, "action": "video", "requestBody": map[string]any{"metadata": tc.body}})
				require.ErrorContains(t, err, tc.wantError)
			})
		}
		frameModel := "MiniMax-H3"
		frameBody := map[string]any{"prompt": "test", "start_frame": "https://example.com/start.png", "end_frame": "https://example.com/end.png", "reference_images": []string{"https://example.com/ref.png"}}
		if provider == "flowith" {
			frameModel = "seedance-2.5"
			frameBody = map[string]any{"prompt": "test", "image": "https://example.com/start.png", "end_frame": "https://example.com/end.png"}
		}
		frameRequest, frameErr := plugin.Engine.Call(context.Background(), "buildSubmitRequest", map[string]any{"model": frameModel, "upstreamModel": provider + "/" + frameModel, "action": "video", "baseUrl": "https://example.com", "apiKey": "test", "requestBody": frameBody})
		if provider == "buzzy" {
			require.ErrorContains(t, frameErr, "additional reference images")
		} else {
			require.NoError(t, frameErr)
			frameOut := frameRequest.(map[string]any)["body"].(map[string]any)
			assert.Equal(t, "https://example.com/start.png", frameOut["start_frame"])
			assert.Equal(t, "https://example.com/end.png", frameOut["end_frame"])
			assert.NotContains(t, frameOut, "image")
		}
		ctx := map[string]any{"model": model, "upstreamModel": provider + "/" + model, "action": "video", "baseUrl": "https://example.com", "apiKey": "test", "requestBody": map[string]any{"prompt": "test", "duration": 10, "resolution": "1080P", "size": "1920x1080"}}
		request, err := plugin.Engine.Call(context.Background(), "buildSubmitRequest", ctx)
		require.NoError(t, err)
		body := request.(map[string]any)["body"].(map[string]any)
		assert.Equal(t, "16:9", body["aspect_ratio"])
		assert.NotContains(t, body, "size")
		usage, err := plugin.Engine.Call(context.Background(), "extractUsage", ctx)
		require.NoError(t, err)
		facts := usage.(map[string]any)
		assert.EqualValues(t, 10, facts["seconds"])
		assert.Equal(t, "1080p", facts["resolution"])
		expression := `u("resolution") == "720p" ? tier("720p", u("seconds") * 0.157) : tier("1080p", u("seconds") * 0.235)`
		result, err := billingexpr.ComputeTieredQuotaWithRequest(&billingexpr.BillingSnapshot{ExprString: expression, ExprHash: billingexpr.ExprHashString(expression), GroupRatio: 1, QuotaPerUnit: 500000, ExprVersion: 1, TaskUsageBilling: true}, billingexpr.TokenParams{}, billingexpr.RequestInput{Usage: facts})
		require.NoError(t, err)
		assert.Equal(t, 1175000, result.ActualQuotaAfterGroup)
		complete, err := plugin.Engine.Call(context.Background(), "extractUsageOnComplete", map[string]any{"action": "video"}, map[string]any{"status": "SUCCESS"}, map[string]any{"data": []any{map[string]any{"url": "https://example.com/a.mp4"}}, "params": map[string]any{"duration": 1}})
		require.NoError(t, err)
		assert.Empty(t, complete, "request echoes are not measured duration and must not change the frozen charge")
	}
}
