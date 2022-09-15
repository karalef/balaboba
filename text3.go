package balaboba

import "context"

// Response represents text generating response.
type Response struct {
	Query    string `json:"query"`
	Text     string `json:"text"`
	BadQuery int    `json:"bad_query"`
	Error    int    `json:"error"`
}

// FullText returns full text.
func (r Response) FullText() string {
	return r.Query + r.Text
}

// Generate generates with passed params.
func (c *Client) Generate(ctx context.Context, query string, style Style, filter ...int) (*Response, error) {
	f := 1
	if len(filter) > 0 {
		f = filter[0]
	}
	var resp Response
	return &resp, c.do(ctx, "text3", map[string]interface{}{
		"query": query, "intro": style.Value(c.lang), "filter": f,
	}, &resp)
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
func (s Style) Value(lang Lang) uint8 {
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
