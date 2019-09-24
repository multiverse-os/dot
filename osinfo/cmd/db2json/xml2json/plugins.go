package xml2json

import (
	"strings"
)

type (
	plugin interface {
		AddToDecoder(*Decoder) *Decoder
	}
	customTypeConverter struct {
		parseTypes []JSType
	}
	attrPrefixer    string
	contentPrefixer string
	excluder        []string
	nodesFormatter  struct {
		list []nodeFormatter
	}
	nodeFormatter struct {
		path   string
		plugin nodePlugin
	}
	nodePlugin interface {
		AddTo(*Node)
	}
	arrayFormatter struct{}
)

func WithTypeConverter(ts ...JSType) *customTypeConverter {
	return &customTypeConverter{parseTypes: ts}
}

func (tc *customTypeConverter) parseAsString(t JSType) bool {
	if t == String {
		return true
	}
	for i := 0; i < len(tc.parseTypes); i++ {
		if tc.parseTypes[i] == t {
			return false
		}
	}
	return true
}

func (tc *customTypeConverter) AddToDecoder(d *Decoder) *Decoder {
	return d
}

func (tc *customTypeConverter) Convert(s string) string {
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		s = s[1 : len(s)-1]
	}
	jsType := Str2JSType(s)
	if tc.parseAsString(jsType) {
		s = `"` + s + `"`
	}
	return s
}

func WithAttrPrefix(prefix string) *attrPrefixer {
	ap := attrPrefixer(prefix)
	return &ap
}

func (a *attrPrefixer) AddToDecoder(d *Decoder) *Decoder {
	d.attributePrefix = string((*a))
	return d
}

func WithContentPrefix(prefix string) *contentPrefixer {
	c := contentPrefixer(prefix)
	return &c
}

func (c *contentPrefixer) AddToDecoder(d *Decoder) *Decoder {
	d.contentPrefix = string((*c))
	return d
}

func ExcludeAttributes(attrs []string) *excluder {
	ex := excluder(attrs)
	return &ex
}

func (ex *excluder) AddToDecoder(d *Decoder) *Decoder {
	d.ExcludeAttributes([]string((*ex)))
	return d
}

// WithNodes formats specific nodes
func WithNodes(n ...nodeFormatter) *nodesFormatter {
	return &nodesFormatter{list: n}
}

func (nf *nodesFormatter) AddToDecoder(d *Decoder) *Decoder {
	d.AddFormatters(nf.list)
	return d
}

func NodePlugin(path string, plugin nodePlugin) nodeFormatter {
	return nodeFormatter{path: path, plugin: plugin}
}

func (nf *nodeFormatter) Format(node *Node) {
	child := node.GetChild(nf.path)
	if child != nil {
		nf.plugin.AddTo(child)
	}
}

func ToArray() *arrayFormatter {
	return &arrayFormatter{}
}

func (af *arrayFormatter) AddTo(n *Node) {
	n.ChildrenAlwaysAsArray = true
}
