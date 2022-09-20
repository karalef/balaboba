package balaboba

import "context"

// Response contains generated text.
type Response struct {
	// Contains the query + generated continuation.
	//
	// If BadQuery is true it contains the
	// bad query text in the Client language.
	Text     string
	BadQuery bool

	raw  response
	lang Lang
}

type response struct {
	responseBase
	Query     string `json:"query"`
	Text      string `json:"text"`
	BadQuery  uint8  `json:"bad_query"`
	IsCached  uint8  `json:"is_cached"`
	Intro     int    `json:"intro"`
	Signature string `json:"signature"`
}

// Generate generates text with passed parameters.
func (c *Client) Generate(query string, style Style, filter ...bool) (*Response, error) {
	return c.GenerateContext(context.Background(), query, style, filter...)
}

// GenerateContext generates text with passed parameters.
// It uses the context for the request.
func (c *Client) GenerateContext(ctx context.Context, query string, style Style, filter ...bool) (*Response, error) {
	f := 0
	if len(filter) > 0 && filter[0] {
		f = 1
	}

	resp := Response{lang: c.lang}
	err := c.doContext(ctx, "text3", map[string]interface{}{
		"query": query, "intro": style.Value(c.lang), "filter": f,
	}, &resp.raw)
	if err != nil {
		return nil, err
	}

	resp.Text = resp.raw.Query + resp.raw.Text
	resp.BadQuery = resp.raw.BadQuery != 0
	if resp.BadQuery {
		if c.lang == Rus {
			resp.Text = BadQueryRus
		} else {
			resp.Text = BadQueryEng
		}
	}

	return &resp, nil
}

// Style of generating text.
type Style uint8

// all styles
const (
	Standart Style = iota
	UserManual
	Recipes
	ShortStories
	WikipediaSipmlified
	MovieSynopses
	FolkWisdom
)

func (s Style) intro(lang Lang) Intro {
	styles := &stylesRus
	if lang == Eng {
		styles = &stylesEng
	}
	if int(s) > len(styles) {
		return styles[0]
	}
	return styles[s]
}

// Description returns style description.
func (s Style) Description(lang Lang) (string, string) {
	i := s.intro(lang)
	return i.String, i.Description
}

// Value returns style code.
func (s Style) Value(lang Lang) int {
	return s.intro(lang).Style
}

var stylesRus = [...]Intro{
	Standart:            {0, "Без стиля", "Напишите что-нибудь и получите продолжение от Балабобы"},
	UserManual:          {24, "Инструкции по применению", "Перечислите несколько предметов, а Балабоба придумает, как их использовать"},
	Recipes:             {25, "Рецепты", "Перечислите съедобные ингредиенты, а Балабоба придумает рецепт с ними"},
	ShortStories:        {6, "Короткие истории", "Начните писать историю, а Балабобы продолжит — иногда страшно, но чаще смешно"},
	WikipediaSipmlified: {8, "Короче, Википедия", "Напишите какое-нибудь слово, а Балабоба даст этому определение"},
	MovieSynopses:       {9, "Синопсисы фильмов", "Напишите название фильма (существующего или нет), а Балабоба расскажет вам, о чем он"},
	FolkWisdom:          {11, "Народные мудрости", "Напишите что-нибудь и получите народную мудрость"},
}

var stylesEng = [...]Intro{
	Standart:            {32, "Standard", "Write something and Balaboba will continue"},
	UserManual:          {26, "User manual", "Give Balaboba a list of items and it'll come up with a way to use them"},
	Recipes:             {27, "Recipes", "Give Balaboba a list of edible ingredients and it'll come up with a recipe"},
	ShortStories:        {28, "Short stories", "Start writing a story and Balaboba will continue. Scary stories, funny stories – you name it."},
	WikipediaSipmlified: {29, "Wikipedia simplified", "Write a word and Balaboba will generate a definition for it"},
	MovieSynopses:       {30, "Movie synopses", "Write the name of a movie (real or made up) and Balaboba will tell you what it's about"},
	FolkWisdom:          {31, "Folk wisdom", "Write something and get a piece of folk wisdom"},
}
