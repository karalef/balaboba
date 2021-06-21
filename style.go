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

func (s Style) String() string {
	switch s {
	case NoStyle:
		return "Без стиля"
	case ConspiracyTheories:
		return "Теории заговора"
	case TVReports:
		return "ТВ-репортажи"
	case Tosts:
		return "Тосты"
	case GuyQuotes:
		return "Пацанские цитаты"
	case AdvertisingSlogans:
		return "Рекламные слоганы"
	case ShortStories:
		return "Короткие истории"
	case InstaCaptions:
		return "Подписи в instagram"
	case InBriefWikipedia:
		return "Короче, Википедия"
	case MovieSynopses:
		return "Синопсисы фильмов"
	case Horoscope:
		return "Гороскоп"
	case FolkWisdom:
		return "Народные мудрости"
	default:
		return ""
	}
}
