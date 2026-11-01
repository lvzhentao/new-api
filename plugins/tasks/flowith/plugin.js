// 模型能力来自对应上游项目 2026-10-31 的模型目录。
const PROVIDER = "flowith";
const CATALOG = {
  "dola-seedream-5-0-flash-260915": {"kind": "image", "durations": [], "min": 0, "max": 0, "resolutions": []},
  "dola-seedream-5-0-pro-260628": {"kind": "image", "durations": [], "min": 0, "max": 0, "resolutions": []},
  "gemini-2.5-flash-image": {"kind": "image", "durations": [], "min": 0, "max": 0, "resolutions": []},
  "gemini-3-pro-image": {"kind": "image", "durations": [], "min": 0, "max": 0, "resolutions": []},
  "gemini-3.1-flash-image": {"kind": "image", "durations": [], "min": 0, "max": 0, "resolutions": []},
  "gemini-3.1-flash-lite-image": {"kind": "image", "durations": [], "min": 0, "max": 0, "resolutions": []},
  "gemini-nano-banana-2.1": {"kind": "image", "durations": [], "min": 0, "max": 0, "resolutions": []},
  "gpt-image-2": {"kind": "image", "durations": [], "min": 0, "max": 0, "resolutions": []},
  "gpt-image-2.5-flare": {"kind": "image", "durations": [], "min": 0, "max": 0, "resolutions": []},
  "gpt-image-2.5-sunburst": {"kind": "image", "durations": [], "min": 0, "max": 0, "resolutions": []},
  "grok-imagine-image-2.0": {"kind": "image", "durations": [], "min": 0, "max": 0, "resolutions": []},
  "z-image-turbo": {"kind": "image", "durations": [], "min": 0, "max": 0, "resolutions": []},
  "kling-video/v2.6/pro/image-to-video": {"kind": "video", "durations": [5, 10], "min": 0, "max": 0, "resolutions": ["720P", "1080P", "2K", "4K"], "aspects": ["16:9", "9:16", "1:1"], "frames": true, "audio": true, "characters": false},
  "kling-video/v2.6/pro/text-to-video": {"kind": "video", "durations": [5, 10], "min": 0, "max": 0, "resolutions": ["720P", "1080P", "2K", "4K"], "aspects": ["16:9", "9:16", "1:1"], "frames": false, "audio": true, "characters": false},
  "kling-video/v3/4k/image-to-video": {"kind": "video", "durations": [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], "min": 0, "max": 0, "resolutions": [], "aspects": [], "frames": true, "audio": true, "characters": false},
  "kling-video/v3/4k/text-to-video": {"kind": "video", "durations": [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], "min": 0, "max": 0, "resolutions": [], "aspects": ["16:9", "9:16", "1:1"], "frames": false, "audio": true, "characters": false},
  "kling-video/v3/pro/image-to-video": {"kind": "video", "durations": [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], "min": 0, "max": 0, "resolutions": [], "aspects": [], "frames": true, "audio": true, "characters": false},
  "kling-video/v3/pro/text-to-video": {"kind": "video", "durations": [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], "min": 0, "max": 0, "resolutions": [], "aspects": ["16:9", "9:16", "1:1"], "frames": false, "audio": true, "characters": false},
  "kling-video/v3/standard/image-to-video": {"kind": "video", "durations": [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], "min": 0, "max": 0, "resolutions": [], "aspects": [], "frames": true, "audio": true, "characters": false},
  "kling-video/v3/standard/text-to-video": {"kind": "video", "durations": [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], "min": 0, "max": 0, "resolutions": [], "aspects": ["16:9", "9:16", "1:1"], "frames": false, "audio": true, "characters": false},
  "kling-video/v3/turbo/image-to-video": {"kind": "video", "durations": [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], "min": 0, "max": 0, "resolutions": ["720p", "1080p"], "aspects": [], "frames": false, "audio": false, "characters": false},
  "kling-video/v3/turbo/text-to-video": {"kind": "video", "durations": [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], "min": 0, "max": 0, "resolutions": ["720p", "1080p"], "aspects": ["16:9", "9:16", "1:1"], "frames": false, "audio": false, "characters": false},
  "kling-video/v4": {"kind": "video", "durations": [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30], "min": 0, "max": 0, "resolutions": ["720p", "1080p"], "aspects": ["auto", "16:9", "9:16", "1:1", "21:9"], "frames": true, "audio": false, "characters": false, "images": 10, "videos": 5},
  "minimax-h3-image-to-video": {"kind": "video", "durations": [4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], "min": 0, "max": 0, "resolutions": ["768P", "2K"], "aspects": [], "frames": true, "audio": false, "characters": false},
  "minimax-h3-text-to-video": {"kind": "video", "durations": [4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], "min": 0, "max": 0, "resolutions": ["768P", "2K"], "aspects": ["21:9", "16:9", "4:3", "1:1", "3:4", "9:16"], "frames": false, "audio": false, "characters": false},
  "seedance-2.0": {"kind": "video", "durations": [4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], "min": 0, "max": 0, "resolutions": ["480p", "720p", "1080p"], "aspects": ["16:9", "9:16", "1:1", "4:3", "3:4", "21:9"], "frames": true, "audio": true, "characters": true},
  "seedance-2.0-fast": {"kind": "video", "durations": [4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], "min": 0, "max": 0, "resolutions": ["480p", "720p"], "aspects": ["16:9", "9:16", "1:1", "4:3", "3:4", "21:9"], "frames": true, "audio": true, "characters": true},
  "seedance-2.5": {"kind": "video", "durations": [-1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30], "min": 0, "max": 0, "resolutions": ["480p", "720p", "1080p"], "aspects": ["adaptive", "16:9", "9:16", "1:1", "4:3", "3:4", "21:9"], "frames": true, "audio": true, "characters": true, "images": 30, "videos": 10},
  "veo-3.1-fast-generate": {"kind": "video", "durations": [4, 6, 8], "min": 0, "max": 0, "resolutions": [], "aspects": ["16:9", "9:16"], "frames": false, "audio": true, "characters": false},
  "veo-3.1-fast-image-to-video": {"kind": "video", "durations": [4, 6, 8], "min": 0, "max": 0, "resolutions": [], "aspects": ["16:9", "9:16"], "frames": false, "audio": true, "characters": false},
  "veo-3.1-generate": {"kind": "video", "durations": [4, 6, 8], "min": 0, "max": 0, "resolutions": [], "aspects": ["16:9", "9:16"], "frames": false, "audio": true, "characters": false},
  "veo-3.1-image-to-video": {"kind": "video", "durations": [4, 6, 8], "min": 0, "max": 0, "resolutions": [], "aspects": ["16:9", "9:16"], "frames": true, "audio": true, "characters": false},
  "wan/v2.6/image-to-video": {"kind": "video", "durations": [5, 10, 15], "min": 0, "max": 0, "resolutions": [], "aspects": ["16:9", "9:16", "1:1", "4:3", "3:4"], "frames": false, "audio": false, "characters": false},
  "wan/v2.6/text-to-video": {"kind": "video", "durations": [5, 10, 15], "min": 0, "max": 0, "resolutions": [], "aspects": ["16:9", "9:16", "1:1", "4:3", "3:4"], "frames": false, "audio": false, "characters": false},
  "wan/v3/image-to-video": {"kind": "video", "durations": [5, 10, 15, 20, 30], "min": 0, "max": 0, "resolutions": ["480p", "720p", "1080p"], "aspects": ["16:9", "9:16", "1:1", "4:3", "3:4"], "frames": true, "audio": true, "characters": false},
  "wan/v3/text-to-video": {"kind": "video", "durations": [5, 10, 15, 20, 30], "min": 0, "max": 0, "resolutions": ["480p", "720p", "1080p"], "aspects": ["16:9", "9:16", "1:1", "4:3", "3:4"], "frames": false, "audio": true, "characters": false},
};
const IMAGE_MODELS = Object.keys(CATALOG).filter(model => CATALOG[model].kind === "image").map(model => PROVIDER + "/" + model);
const VIDEO_MODELS = Object.keys(CATALOG).filter(model => CATALOG[model].kind === "video").map(model => PROVIDER + "/" + model);
export const meta = {
  apiVersion: 1, key: PROVIDER, name: PROVIDER === "buzzy" ? "Buzzy Media" : "Flowith Media",
  version: "1.3.0", author: {name: "klao"}, fetchMode: "per_task",
  description: {en: "Synchronous image and video generation", zh: "同步图片与视频生成"},
  models: IMAGE_MODELS.concat(VIDEO_MODELS),
  protocols: [{name: "openai_image", models: IMAGE_MODELS}, {name: "openai_video", models: VIDEO_MODELS}],
  usageSchema: {},
  usageProfiles: [
    {models: IMAGE_MODELS, schema: {image_count: {type: "number", unit: "count", description: {en: "Image generation unit price", zh: "图片生成单价"}}, resolution: {enum: ["model_default", "1k", "2k", "3k", "4k"], description: {en: "Output image resolution", zh: "输出图片分辨率"}}}},
    {models: VIDEO_MODELS, schema: {
      seconds: {type: "number", unit: "second", description: {en: "Video generation unit price", zh: "视频生成单价"}},
      resolution: {enum: ["360p", "480p", "720p", "768p", "1080p", "2k", "4k", "model_default"], description: {en: "Output video resolution", zh: "输出视频分辨率"}},
    }},
  ],
  routes: [
    {method: "POST", path: "/" + PROVIDER + "/v1/videos/generations", type: "submit", models: VIDEO_MODELS, decode: "video", render: "result"},
    {method: "GET", path: "/" + PROVIDER + "/v1/videos/:task_id", type: "query", render: "result"},
  ],
};

// GPT 的 4K 方图在 Flowith 对应 2880x2880，按目录档位而非最长边猜价。
const GPT_IMAGE_SIZES = {
  "1:1": ["1024x1024", "2048x2048", "2880x2880"],
  "3:2": ["1024x688", "2048x1360", "3456x2304"],
  "2:3": ["688x1024", "1360x2048", "2304x3456"],
  "4:3": ["1024x768", "2048x1536", "3264x2448"],
  "3:4": ["768x1024", "1536x2048", "2448x3264"],
  "16:9": ["1280x720", "2048x1152", "3840x2160"],
  "9:16": ["720x1280", "1152x2048", "2160x3840"],
  "2:1": ["2048x1024", "2688x1344", "3840x1920"],
  "1:2": ["1024x2048", "1344x2688", "1920x3840"],
  "21:9": ["1344x576", "2048x896", "3840x1648"],
  "9:21": ["576x1344", "896x2048", "1648x3840"],
};

// 同一归一化同时供上游参数与计费使用；缺省显式发送 1K，冲突拒绝。
function normalizeImageResolution(out, model) {
  const gpt = ["gpt-image-2", "gpt-image-2.5-flare", "gpt-image-2.5-sunburst"].includes(model);
  const banana = PROVIDER === "buzzy" ? ["Nano Banana 2", "Nano Banana 2.1", "Nano Banana Pro"].includes(model)
    : ["gemini-3.1-flash-image", "gemini-nano-banana-2.1", "gemini-3-pro-image"].includes(model);
  const grok = model === (PROVIDER === "buzzy" ? "grok-imagine-image" : "grok-imagine-image-2.0");
  const seedream = PROVIDER === "buzzy" && ["doubao-seedream-4-0-250828", "doubao-seedream-4-5-251128", "doubao-seedream-5-0-260128"].includes(model);
  if (!gpt && !banana && !grok && !seedream) return "model_default";
  const resolutions = seedream ? CATALOG[model].resolutions.map(value => value.toLowerCase()) : grok ? ["1k", "2k"] : ["1k", "2k", "4k"];
  let resolution, ratio;
  for (const field of ["resolution", "image_size", "size", "aspect_ratio"]) {
    if (out[field] === undefined) continue;
    if (typeof out[field] !== "string" || !out[field].trim()) throw new Error(field + " must be a non-empty string");
    const value = out[field].trim().toLowerCase();
    let nextResolution, nextRatio;
    if (field !== "aspect_ratio" && resolutions.includes(value)) nextResolution = value;
    else if (field !== "aspect_ratio" && value === "512" && PROVIDER === "flowith" && model === "gemini-3.1-flash-image") nextResolution = "512";
    else if (field === "size" || field === "aspect_ratio") {
      if (/^[1-9][0-9]*(?:\.[0-9]+)?:[1-9][0-9]*(?:\.[0-9]+)?$/.test(value) || (field === "aspect_ratio" && value === "auto" && grok && PROVIDER === "buzzy")) nextRatio = value;
      else {
        for (const [aspect, sizes] of Object.entries(GPT_IMAGE_SIZES)) {
          const index = sizes.indexOf(value);
          if (index >= 0) { nextRatio = aspect; nextResolution = ["1k", "2k", "4k"][index]; break; }
        }
        // 香蕉接受标准方图像素输入，不把 GPT 的特殊 4K 方图混为它的原生尺寸。
        if (banana && ["1024x1024", "2048x2048", "4096x4096"].includes(value)) {
          nextRatio = "1:1"; nextResolution = {"1024x1024":"1k", "2048x2048":"2k", "4096x4096":"4k"}[value];
        } else if (banana && nextResolution) throw new Error("Use resolution and aspect_ratio for banana models");
        if (!nextRatio) throw new Error("Unsupported image size; use 1k, 2k, 4k or a supported pixel size");
      }
    } else throw new Error("Unsupported image resolution; use 1k, 2k or 4k");
    if (resolution && nextResolution && resolution !== nextResolution) throw new Error("Image resolution parameters must agree");
    if (ratio && nextRatio && ratio !== nextRatio) throw new Error("Image aspect ratio parameters must agree");
    resolution = nextResolution || resolution;
    ratio = nextRatio || ratio;
  }
  resolution = resolution || resolutions[0];
  if (!resolutions.includes(resolution) && !(resolution === "512" && PROVIDER === "flowith" && model === "gemini-3.1-flash-image")) throw new Error("Image resolution is not supported by this model");
  ratio = ratio || "1:1";
  const ratios = grok ? ["1:1","3:4","4:3","9:16","16:9","2:3","3:2","9:19.5","19.5:9","9:20","20:9","1:2","2:1"].concat(PROVIDER === "buzzy" ? ["auto"] : [])
    : seedream ? ["1:1","4:3","3:4","16:9","9:16","3:2","2:3","21:9"]
    : gpt ? Object.keys(GPT_IMAGE_SIZES).filter(r => PROVIDER === "flowith" || !["2:1", "1:2"].includes(r))
    : PROVIDER === "buzzy" ? ["1:1","2:3","3:2","3:4","4:3","4:5","5:4","9:16","16:9","21:9"]
    : model === "gemini-3-pro-image" ? ["1:1","3:2","2:3","3:4","4:3","4:5","5:4","9:16","16:9","21:9"]
    : ["1:1","1:4","4:1","1:8","8:1","2:3","3:2","3:4","4:3","4:5","5:4","9:16","16:9","21:9"];
  if (!ratios.includes(ratio)) throw new Error("Unsupported image aspect ratio");
  delete out.size;
  delete out.resolution;
  delete out.image_size;
  out.aspect_ratio = ratio;
  if (PROVIDER === "buzzy") out.resolution = banana || seedream ? resolution.toUpperCase() : resolution;
  else if (gpt) out.aspect_ratio = GPT_IMAGE_SIZES[ratio][["1k", "2k", "4k"].indexOf(resolution)];
  else out.image_size = resolution;
  return resolution === "512" ? "1k" : resolution;
}

function normalize(body, action, model, finalModel) {
  if (!body || typeof body !== "object" || Array.isArray(body)) throw new Error("JSON object required");
  if (body.metadata != null && (typeof body.metadata !== "object" || Array.isArray(body.metadata))) throw new Error("metadata must be an object");
  const out = Object.assign({}, body.metadata || {}, body);
  delete out.metadata;
  const upstreamModel = model && model.startsWith(PROVIDER + "/") ? model.slice(PROVIDER.length + 1) : model;
  out.model = finalModel ? upstreamModel : model;
  if (typeof out.prompt !== "string" || !out.prompt.trim()) throw new Error("prompt is required");
  if (out.stream === true) throw new Error("Media streaming is not supported by this upstream");
  const maxCount = action === "image" ? 4 : 1;
  for (const field of ["n", "count"]) {
    if (out[field] !== undefined && (typeof out[field] !== "number" || !Number.isInteger(out[field]) || out[field] < 1 || out[field] > maxCount))
      throw new Error(field + " must be an integer between 1 and " + maxCount);
  }
  if (out.n !== undefined && out.count !== undefined && out.n !== out.count) throw new Error("n and count must agree");
  out.n = out.n === undefined ? (out.count === undefined ? 1 : out.count) : out.n;
  delete out.count;
  if (PROVIDER === "flowith") {
    for (const pair of [["image_urls", "reference_images"], ["video_urls", "reference_videos"]]) {
      if (out[pair[0]] !== undefined) {
        if (!Array.isArray(out[pair[0]]) || (out[pair[1]] !== undefined && !Array.isArray(out[pair[1]]))) throw new Error(pair[0] + " must be an array");
        out[pair[1]] = (out[pair[1]] || []).concat(out[pair[0]]);
        delete out[pair[0]];
      }
    }
    if (action === "image" && finalModel && !["gpt-image-2", "gpt-image-2.5-flare", "gpt-image-2.5-sunburst", "gemini-3.1-flash-image", "gemini-nano-banana-2.1", "gemini-3-pro-image", "grok-imagine-image-2.0"].includes(upstreamModel) && out.resolution !== undefined) {
      if (out.image_size !== undefined && out.image_size !== out.resolution) throw new Error("resolution and image_size must agree");
      out.image_size = out.resolution;
      delete out.resolution;
    }
  }
  if (action === "image" && finalModel) normalizeImageResolution(out, upstreamModel);
  if (action === "video") {
    const capability = finalModel ? CATALOG[upstreamModel] : undefined;
    for (const field of ["duration", "seconds"]) {
      if (out[field] === undefined) continue;
      const value = out[field];
      if ((typeof value !== "number" && (typeof value !== "string" || !/^[0-9]+$/.test(value))) || !Number.isInteger(Number(value)) || Number(value) < 1 || Number(value) > 3600)
        throw new Error(field + " must be a positive integer no greater than 3600; automatic duration is unsupported");
      out[field] = Number(value);
    }
    if (out.duration !== undefined && out.seconds !== undefined && out.duration !== out.seconds) throw new Error("duration and seconds must agree");
    let duration = out.duration === undefined ? out.seconds : out.duration;
    if (capability && capability.kind !== "video") throw new Error("A video model is required");
    if (capability) {
      // 在预扣和发送前执行同一校验，避免参数被上游忽略后仍按用户所选规格收费。
      for (const field of ["enhancement", "enhance", "upscale", "upscale_resolution", "middle_images", "last_image", "elements", "voice_ids", "video", "reference_audio", "reference_audios"]) {
        if (out[field] !== undefined) throw new Error(field + " is not supported by this video interface");
      }
      if (PROVIDER === "flowith" && out.audio_urls !== undefined) throw new Error("audio_urls is not supported by this video interface");
      for (const field of ["audio", "loop", "use_character_assets", "keep_original_audio", "character_orientation"]) {
        if (out[field] !== undefined && typeof out[field] !== "boolean") throw new Error(field + " must be a boolean");
      }
      if (out.audio === true && !capability.audio) throw new Error("audio control is not supported by this model");
      let aspect = out.aspect_ratio;
      if (aspect !== undefined && typeof aspect !== "string") throw new Error("aspect_ratio must be a string");
      if (out.size !== undefined) {
        if (typeof out.size !== "string") throw new Error("size must be an aspect ratio or pixel dimensions");
        let sizeAspect = out.size;
        const pixels = /^([1-9][0-9]*)x([1-9][0-9]*)$/i.exec(out.size);
        if (pixels) {
          const ratio = Number(pixels[1]) / Number(pixels[2]);
          sizeAspect = capability.aspects.find(value => {
            const parts = value.split(":");
            return parts.length === 2 && Math.abs(Number(parts[0]) / Number(parts[1]) - ratio) < 0.0001;
          });
          if (!sizeAspect) throw new Error("size does not match a supported aspect_ratio; use aspect_ratio explicitly");
        }
        if (aspect !== undefined && aspect !== sizeAspect) throw new Error("size and aspect_ratio must agree");
        aspect = sizeAspect;
        delete out.size;
      }
      if (aspect !== undefined && capability.aspects.length && !capability.aspects.includes(aspect)) throw new Error("aspect_ratio is not supported by this model");
      if (aspect !== undefined) out.aspect_ratio = aspect;
      if (out.image !== undefined) {
        const images = Array.isArray(out.image) ? out.image : [out.image];
        if (images.length !== 1 || typeof images[0] !== "string" || !/^https?:\/\//i.test(images[0])) throw new Error("image must be a single HTTP(S) URL; use reference_images for multiple images");
        if (out.image_url !== undefined && out.image_url !== images[0]) throw new Error("Frame aliases must agree");
        out.image_url = images[0];
        delete out.image;
      }
      for (const fields of [["start_frame", "first_frame", "image_url"], ["end_frame", "last_frame"]]) {
        let value;
        for (const field of fields) {
          if (out[field] === undefined) continue;
          if (typeof out[field] !== "string" || !/^https?:\/\//i.test(out[field])) throw new Error(field + " must be an HTTP(S) URL");
          if (value !== undefined && value !== out[field]) throw new Error("Frame aliases must agree");
          value = out[field];
          delete out[field];
        }
        if (value !== undefined) out[fields[0]] = value;
      }
      if (out.end_frame !== undefined && !capability.frames) throw new Error("end_frame is not supported by this model");
      if (out.end_frame && !out.start_frame) throw new Error("end_frame requires start_frame");
      for (const fields of [["image_urls", "reference_images"], ["video_urls", "reference_videos"], ["audio_urls"]]) {
        const kind = fields[0].split("_")[0] + "s";
        let count = kind === "images" ? Number(!!out.start_frame) + Number(!!out.end_frame) : 0;
        for (const field of fields) {
          if (out[field] === undefined) continue;
          if (!Array.isArray(out[field]) || out[field].some(value => typeof value !== "string" || !/^https?:\/\//i.test(value))) throw new Error(field + " must be an array of HTTP(S) URLs");
          count += out[field].length;
        }
        if (kind === "images" && out.image !== undefined) {
          const images = Array.isArray(out.image) ? out.image : [out.image];
          if (images.some(value => typeof value !== "string" || !/^https?:\/\//i.test(value))) throw new Error("image must contain HTTP(S) URLs");
          if (PROVIDER === "flowith" && (images.length !== 1 || out.start_frame)) throw new Error("Use reference_images for multiple images; image is a single start frame alias");
          count += images.length;
        }
        if (capability[kind] !== undefined && count > capability[kind]) throw new Error(kind + " exceeds this model's reference limit");
      }
      const allowed = capability.durations;
      if (duration === undefined) duration = allowed.find(value => value > 0) || capability.min || 5;
      if ((allowed.length && !allowed.includes(duration)) || (capability.min && duration < capability.min) || (capability.max && duration > capability.max))
        throw new Error("duration is outside the model's supported range");
    }
    if (capability) {
      const choices = capability.resolutions;
      if (out.resolution === undefined && choices.length) out.resolution = choices[0];
      if (out.resolution !== undefined) {
        if (typeof out.resolution !== "string") throw new Error("resolution must be a string");
        const match = choices.find(value => value.toLowerCase() === out.resolution.toLowerCase());
        if (!match) throw new Error("resolution is not supported by this model; omit it for the model default");
        out.resolution = match;
      }
      // 动作跟随的时长由参考视频决定，接口无成片时长；不能按虚构秒数收费。
      if (PROVIDER === "buzzy" && upstreamModel === "kling-motion-control") throw new Error("Motion-control billing requires verified reference video duration");
    }
    if (duration !== undefined) out.duration = duration;
    delete out.seconds;
    delete out.n;
    if (PROVIDER === "buzzy") out.count = 1;
  }
  return out;
}

function decode(ctx, action) {
  if (!ctx.body || ctx.body.kind !== "json") throw new Error("JSON body required; use reference image URLs");
  return {kind: "submit", model: ctx.model || ctx.body.value.model, action,
    requestBody: normalize(ctx.body.value, action, ctx.model || ctx.body.value.model)};
}

export function buildSubmitRequest(ctx) {
  const model = ctx.upstreamModel || ctx.model;
  const body = normalize(ctx.requestBody, ctx.action, model, true);
  if (!CATALOG[body.model] || CATALOG[body.model].kind !== ctx.action) throw new Error("Unsupported media model for this operation");
  return {url: ctx.baseUrl.replace(/\/+$/, "") + (ctx.action === "image" ? "/v1/images/generations" : "/v1/videos/generations"),
    method: "POST", headers: {Authorization: "Bearer " + ctx.apiKey, "Content-Type": "application/json"}, body};
}

function payloads(body) {
  const data = body && body.data;
  return (Array.isArray(data) ? data : data && typeof data === "object" ? [data] : []).filter(entry =>
    entry && ((typeof entry.url === "string" && entry.url.trim()) || (typeof entry.b64_json === "string" && entry.b64_json.trim())));
}

export function parseSubmitResponse(ctx, resp) {
  const body = resp.body || {};
  const entries = payloads(body);
  const url = (entries.find(entry => typeof entry.url === "string" && entry.url.trim()) || {}).url || "";
  const success = !body.error && entries.length > 0 && (ctx.action !== "video" || !!url);
  // 这是同步结果，不虚构可轮询的上游任务。
  return {taskId: ctx.publicTaskId || utils.uuid(), taskData: body,
    immediate: success ? {status: "SUCCESS", progress: "100%", url} : {status: "FAILURE", reason: "Upstream returned no completed media output"}};
}

export function extractUsage(ctx) {
  const body = normalize(ctx.requestBody, ctx.action, ctx.upstreamModel || ctx.model, true);
  if (ctx.action === "image") {
    const resolution = normalizeImageResolution(body, body.model);
    // 旧按张路径只接收数量倍率；枚举只用于表达式计费。
    return ctx.usagePurpose === "billing_ratios" ? {image_count: body.n} : {image_count: body.n, resolution};
  }
  if (!CATALOG[body.model]) throw new Error("Unsupported video model");
  // 时长和分辨率与最终发送的请求一致；上游 params 只是请求回显，不作为实际用量覆盖。
  const resolution = body.resolution ? body.resolution.toLowerCase() : "model_default";
  if (ctx.usagePurpose === "billing_ratios") throw new Error("Video requires per-second pricing with an explicit resolution expression");
  return {seconds: body.duration, resolution};
}

export function extractUsageOnComplete(task, result, body) {
  const entries = payloads(body);
  if (task.action !== "image") return {};
  const urls = entries.filter(entry => typeof entry.url === "string" && entry.url.trim()).length;
  const base64 = entries.filter(entry => typeof entry.b64_json === "string" && entry.b64_json.trim()).length;
  const count = Math.max(urls, base64);
  return count > 0 && count <= 4 ? {image_count: count} : {};
}

export function extractUsageOnSubmit(ctx, body) {
  return ctx.action === "image" ? extractUsageOnComplete(ctx, {}, body) : {};
}

export function buildQueryRequest() { throw new Error("This upstream has no asynchronous task query API"); }
export function parseTaskResult() { return {status: "UNKNOWN"}; }
export function listArtifacts(task) {
  return task.action === "video" && task.status === "SUCCESS" ? [{key: "video", type: "video", mimeType: "video/mp4"}] : [];
}
export function buildContentRequest(ctx) {
  const entry = payloads(ctx.data).find(item => typeof item.url === "string" && item.url.trim());
  if (!entry || !/^https?:\/\//.test(entry.url)) throw new Error("artifact_not_found");
  return {url: entry.url, method: ctx.clientRequest && ctx.clientRequest.method === "HEAD" ? "HEAD" : "GET", credentialless: true};
}
export const protocols = {
  openai_image: {decodeRequest: ctx => decode(ctx, "image"), render: (ctx, task) => Object.assign({}, task.data, {model: ctx.model, data: payloads(task.data)})},
  openai_video: {decodeRequest: ctx => decode(ctx, "video"), render: (ctx, task) => task.data || {}},
};
export const native = {video: ctx => decode(ctx, "video"), result: (ctx, task) => Object.assign({}, task.data, {id: task.task_id, status: task.status})};
