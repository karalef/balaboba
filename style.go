package balaboba

// Style of generating text.
type Style uint8

// all styles
const (
	NoStyle            Style = iota // Без стиля
	ConspiracyTheories              // Теории заговора
	TVReports                       // ТВ-репортажи
	Tosts                           // Тосты
	GuyQuotes                       // Пацанские цитаты
	AdvertisingSlogans              // Рекламные слоганы
	ShortStories                    // Короткие истории
	InstaCaptions                   // Подписи в instagram
	InBriefWikipedia                // Короче, Википедия
	MovieSynopses                   // Синопсисы фильмов
	Horoscope                       // Гороскоп
	FolkWisdom                      // Народные мудрости
)

// Valid ...
func (s Style) Invalid() bool {
	return int(s) > len(strStyles)-1
}

func (s *Style) String() string {
	if s.Invalid() {
		*s = NoStyle
	}
	return strStyles[*s][0]
}

// Description of this style.
func (s *Style) Description() string {
	if s.Invalid() {
		*s = NoStyle
	}
	return strStyles[*s][1]
}

var strStyles = [...][2]string{
	NoStyle:            {"Без стиля", "Напишите что-нибудь и получите продолжение от Балабобы Без стиля"},
	ConspiracyTheories: {"Теории заговора", "Напишите, про что должна быть теория Теории заговора"},
	TVReports:          {"ТВ-репортажи", "Введите фразу, с которой мог бы начинаться сюжет ТВ-репортажи"},
	Tosts:              {"Тосты", "Напишите, с чего должен начинаться тост. Например, какое-нибудь обращение Тосты"},
	GuyQuotes:          {"Пацанские цитаты", "Введите утверждение, которому хотите придать весомость Пацанские цитаты"},
	AdvertisingSlogans: {"Рекламные слоганы", "Напишите название бренда, товара или имя бабушки Рекламные слоганы"},
	ShortStories:       {"Короткие истории", "Начните писать историю, а Балабоба продолжит — иногда страшно, но чаще смешно Короткие истории"},
	InstaCaptions:      {"Подписи в instagram", "Напишите, с чего должна начинаться подпись Подписи в instagram"},
	InBriefWikipedia:   {"Короче, Википедия", "Введите понятие, которое хотите определить Короче, Википедия"},
	MovieSynopses:      {"Синопсисы фильмов", "Напишите название фильма (существующего или нет), а Балабоба расскажет вам, о чем он Синопсисы фильмов"},
	Horoscope:          {"Гороскоп", "Введите знак зодиака и получите персональный гороскоп Гороскоп"},
	FolkWisdom:         {"Народные мудрости", "Напишите что-нибудь и получите народную мудрость Народные мудрости"},
}
