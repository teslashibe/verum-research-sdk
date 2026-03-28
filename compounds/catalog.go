package compounds

import "strings"

type Catalog struct {
	compounds []Compound
	byName    map[string]*Compound
	byAlias   map[string]*Compound
}

func Default() *Catalog {
	c := &Catalog{
		compounds: seedCompounds,
		byName:    make(map[string]*Compound),
		byAlias:   make(map[string]*Compound),
	}
	c.buildIndex()
	return c
}

func (c *Catalog) buildIndex() {
	c.byName = make(map[string]*Compound, len(c.compounds))
	c.byAlias = make(map[string]*Compound, len(c.compounds)*3)
	for i := range c.compounds {
		compound := &c.compounds[i]
		c.byName[normalize(compound.Name)] = compound
		for _, alias := range compound.Aliases {
			c.byAlias[normalize(alias)] = compound
		}
	}
}

func (c *Catalog) All() []Compound {
	out := make([]Compound, len(c.compounds))
	copy(out, c.compounds)
	return out
}

func (c *Catalog) Len() int {
	return len(c.compounds)
}

func (c *Catalog) FindByName(name string) (*Compound, bool) {
	key := normalize(name)
	if compound, ok := c.byName[key]; ok {
		return compound, true
	}
	if compound, ok := c.byAlias[key]; ok {
		return compound, true
	}
	return nil, false
}

func (c *Catalog) FindByID(id string) (*Compound, bool) {
	for i := range c.compounds {
		if c.compounds[i].ID == id {
			return &c.compounds[i], true
		}
	}
	return nil, false
}

func (c *Catalog) SearchByCategory(cat Category) []Compound {
	var results []Compound
	for _, compound := range c.compounds {
		if compound.Category == cat {
			results = append(results, compound)
		}
	}
	return results
}

func (c *Catalog) SearchByKeyword(keyword string) []Compound {
	key := normalize(keyword)
	var results []Compound
	for _, compound := range c.compounds {
		if containsNormalized(compound.Name, key) ||
			containsNormalized(compound.Mechanism, key) ||
			containsNormalized(compound.Description, key) ||
			anyContainsNormalized(compound.Aliases, key) {
			results = append(results, compound)
		}
	}
	return results
}

func (c *Catalog) Add(compound Compound) {
	c.compounds = append(c.compounds, compound)
	c.buildIndex()
}

func normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func containsNormalized(haystack, needle string) bool {
	return strings.Contains(normalize(haystack), needle)
}

func anyContainsNormalized(haystack []string, needle string) bool {
	for _, s := range haystack {
		if containsNormalized(s, needle) {
			return true
		}
	}
	return false
}
