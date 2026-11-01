package model

import (
	"slices"
	"sync"
	"sync/atomic"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/pkg/jsplugin"
)

// TaskAliasTarget is one mapping-derived alias after cross-channel aggregation.
// Declared is empty when the same plugin resolves the alias to more than one
// declared tail (display-only; pin is disabled).
type TaskAliasTarget struct {
	Alias     string
	Declared  string
	PluginKey string
}

type taskAliasView struct {
	generation uint64
	expiresAt  time.Time
	byFold     map[string]TaskAliasTarget
	targets    map[string][]TaskAliasTarget
}

const taskAliasViewTTL = 60 * time.Second

var (
	taskAliasViewPtr   atomic.Pointer[taskAliasView]
	taskAliasRebuildMu sync.Mutex
)

// ResolveTaskModelAlias returns the mapping-derived alias target for name
// (exact or ASCII-folded). g is the caller's routing generation: a Number
// mismatch or TTL expiry rebuilds the view against that generation.
func ResolveTaskModelAlias(g *jsplugin.RoutingGeneration, name string) (TaskAliasTarget, bool) {
	if g == nil || name == "" {
		return TaskAliasTarget{}, false
	}
	view := loadFreshTaskAliasView(g)
	if view == nil {
		return TaskAliasTarget{}, false
	}
	target, ok := view.byFold[jsplugin.ASCIIFold(name)]
	return target, ok
}

// ResolveTaskModelAliasTargets preserves the provider-specific targets of a
// shared alias. Channel selection decides which provider may execute it.
func ResolveTaskModelAliasTargets(g *jsplugin.RoutingGeneration, name string) []TaskAliasTarget {
	if g == nil || name == "" {
		return nil
	}
	view := loadFreshTaskAliasView(g)
	if view == nil {
		return nil
	}
	return slices.Clone(view.targets[jsplugin.ASCIIFold(name)])
}

// ResolveTaskPluginModel matches a client model name against a plugin's
// declared models as the host protocols do: a case variant is served under
// the declared spelling, a channel alias under the alias. declared is the
// model of models the name stands for. A nil g matches exact names only.
func ResolveTaskPluginModel(g *jsplugin.RoutingGeneration, name string, models []string) (served, declared string, ok bool) {
	if slices.Contains(models, name) {
		return name, name, true
	}
	if canonical, found := g.CanonicalModel(name); found && slices.Contains(models, canonical) {
		return canonical, canonical, true
	}
	for _, target := range ResolveTaskModelAliasTargets(g, name) {
		if slices.Contains(models, target.Declared) {
			return target.Alias, target.Declared, true
		}
	}
	return "", "", false
}

func loadFreshTaskAliasView(g *jsplugin.RoutingGeneration) *taskAliasView {
	view := taskAliasViewPtr.Load()
	if taskAliasViewFresh(view, g.Number) {
		return view
	}
	taskAliasRebuildMu.Lock()
	defer taskAliasRebuildMu.Unlock()
	view = taskAliasViewPtr.Load()
	if taskAliasViewFresh(view, g.Number) {
		return view
	}
	rebuilt := buildTaskAliasView(g)
	taskAliasViewPtr.Store(rebuilt)
	return rebuilt
}

func taskAliasViewFresh(view *taskAliasView, generation uint64) bool {
	return view != nil && view.generation == generation && time.Now().Before(view.expiresAt)
}

func rebuildTaskAliasView() {
	taskAliasRebuildMu.Lock()
	defer taskAliasRebuildMu.Unlock()
	taskAliasViewPtr.Store(buildTaskAliasView(jsplugin.DefaultRegistry.Generation()))
}

type taskAliasDraft struct {
	spellings []string
	byPlugin  map[string]map[string]struct{}
}

func buildTaskAliasView(generation *jsplugin.RoutingGeneration) *taskAliasView {
	genNum := uint64(0)
	if generation != nil {
		genNum = generation.Number
	}
	view := &taskAliasView{
		generation: genNum,
		expiresAt:  time.Now().Add(taskAliasViewTTL),
		byFold:     make(map[string]TaskAliasTarget),
		targets:    make(map[string][]TaskAliasTarget),
	}
	if DB == nil {
		return view
	}

	var channels []Channel
	err := DB.Select("id", "type", "models", "model_mapping").
		Where("status = ?", common.ChannelStatusEnabled).
		Find(&channels).Error
	if err != nil {
		common.SysError(common.LogText("rebuild task alias view: %s", err.Error()))
		return view
	}

	drafts := make(map[string]*taskAliasDraft)
	for i := range channels {
		channel := &channels[i]
		mappingJSON := channel.GetModelMapping()
		if mappingJSON == "" || mappingJSON == "{}" {
			continue
		}
		modelMap := make(map[string]string)
		if err := common.UnmarshalJsonStr(mappingJSON, &modelMap); err != nil {
			common.SysError(common.LogText("task alias view: channel %d model_mapping: %s", channel.Id, err.Error()))
			continue
		}
		inModels := make(map[string]struct{})
		for _, modelName := range channel.GetModels() {
			inModels[modelName] = struct{}{}
		}
		for alias, mapped := range modelMap {
			if mapped == "" {
				continue
			}
			if _, exposed := inModels[alias]; !exposed {
				continue
			}
			tail, cyclic := followChannelModelMapping(modelMap, alias)
			if cyclic {
				common.SysError(common.LogText("task alias mapping cycle dropped: channel=%d key=%q", channel.Id, alias))
				continue
			}
			declared, ok := generation.CanonicalModel(tail)
			if !ok {
				continue
			}
			plugins := generation.PluginsByModel(declared)
			if len(plugins) == 0 {
				continue
			}
			fold := jsplugin.ASCIIFold(alias)
			draft := drafts[fold]
			if draft == nil {
				draft = &taskAliasDraft{byPlugin: make(map[string]map[string]struct{})}
				drafts[fold] = draft
			}
			draft.spellings = append(draft.spellings, alias)
			for _, plugin := range plugins {
				declareds := draft.byPlugin[plugin.Meta.Key]
				if declareds == nil {
					declareds = make(map[string]struct{})
					draft.byPlugin[plugin.Meta.Key] = declareds
				}
				declareds[declared] = struct{}{}
			}
		}
	}

	for _, draft := range drafts {
		alias := draft.spellings[0]
		for _, spelling := range draft.spellings[1:] {
			if spelling < alias {
				alias = spelling
			}
		}
		keys := make([]string, 0, len(draft.byPlugin))
		for key := range draft.byPlugin {
			keys = append(keys, key)
		}
		slices.Sort(keys)
		for _, key := range keys {
			declareds := draft.byPlugin[key]
			if len(declareds) != 1 {
				continue
			}
			for declared := range declareds {
				view.targets[jsplugin.ASCIIFold(alias)] = append(view.targets[jsplugin.ASCIIFold(alias)], TaskAliasTarget{Alias: alias, Declared: declared, PluginKey: key})
			}
		}
		if len(draft.byPlugin) != 1 {
			continue
		}

		var pluginKey, declared string
		for key, declareds := range draft.byPlugin {
			pluginKey = key
			if len(declareds) == 1 {
				for name := range declareds {
					declared = name
				}
			}
		}
		view.byFold[jsplugin.ASCIIFold(alias)] = TaskAliasTarget{
			Alias:     alias,
			Declared:  declared,
			PluginKey: pluginKey,
		}
	}
	return view
}

// followChannelModelMapping walks one channel's mapping the same way
// ModelMappedHelper does: visited-set cycle detection, self-map stops at
// the current hop, a non-self cycle is reported to the caller.
func followChannelModelMapping(modelMap map[string]string, start string) (string, bool) {
	current := start
	visited := map[string]bool{current: true}
	for {
		mapped, exists := modelMap[current]
		if !exists || mapped == "" {
			return current, false
		}
		if visited[mapped] {
			if mapped == current {
				return current, false
			}
			return "", true
		}
		visited[mapped] = true
		current = mapped
	}
}

// TaskModelTarget returns the declared plugin model selected by this channel's
// mapping. An unrecognized vendor tail leaves provider selection to the binding.
func (channel *Channel) TaskModelTarget(g *jsplugin.RoutingGeneration, name string) string {
	mapping := make(map[string]string)
	if g == nil || common.UnmarshalJsonStr(channel.GetModelMapping(), &mapping) != nil {
		return ""
	}
	tail, cyclic := followChannelModelMapping(mapping, name)
	if cyclic {
		return ""
	}
	declared, _ := g.CanonicalModel(tail)
	return declared
}
