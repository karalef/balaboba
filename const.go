package balaboba

// Warn1Rus is a first warning in russian.
const Warn1Rus = `Нейросеть не знает, что говорит, и может сказать всякое — если что, не обижайтесь.
Распространяя получившиеся тексты, помните об ответственности. (18+)`

// Warn1Eng is a first warning in english.
const Warn1Eng = `The neural network doesn’t really know what it’s saying, so it can say absolutely anything.
Don’t get offended if it says something that hurts your feelings.
When sharing the texts, make sure they’re not offensive or violate the law.`

// Warn2Rus is a second warning in russian.
const Warn2Rus = `Генератор может выдавать очень странные тексты.
Пожалуйста, будьте разумны, распространяя их.
Подумайте, не будет ли текст обидным для кого-то и не станет ли его публикация нарушением закона.`

// Warn2Eng is a second warning in english.
const Warn2Eng = `Balaboba can generate some very strange texts.
Please be reasonable when sharing them online.
Consider whether publishing the text may offend anyone or violate the law.`

// AboutRus Balaboba in russian.
const AboutRus = `Балабоба демонстрирует, как с помощью нейросетей семейства YaLM можно продолжать тексты на любую тему, сохраняя связность и заданный стиль.
Здесь эти нейросети используются для развлечения, но разрабатывались они для серьёзных задач — об этом можно почитать тут (https://yandex.ru/lab/yalm-howto).

У Балабобы нет своего мнения или знания.
Он умеет только подражать — писать тексты так, чтобы они были максимально похожи на реальные тексты из интернета.`

// AboutEng Balaboba in english.
const AboutEng = `Balaboba demonstrates how neural networks of the YaLM family can generate a coherent text on any topic, while maintaining a preset style.
In Balaboba’s case, YaLM is used for entertainment, but the models were originally developed for serious tasks. Learn more here (https://yandex.com/lab/yalm-howto-en).

Balaboba has no opinion or knowledge of its own. It can only imitate texts found on the internet.`

// BadQueryRus is a bad query response in russian.
const BadQueryRus = `Балабоба не принимает запросы на острые темы, например, про политику или религию.
Люди могут слишком серьёзно отнестись к сгенерированным текстам.

Вероятность того, что запрос задаёт одну из острых тем, определяет нейросеть, обученная на оценках случайных людей.
Но она может перестараться или, наоборот, что-то пропустить.`

// BadQueryEng is a bad query response in english.
const BadQueryEng = `Balaboba doesn’t accept queries about sensitive topics, like politics or religion.
People may take the generated text too seriously.

The probability that a query contains a triggering or sensitive topic is determined by a neural network trained on random people’s estimates.
It can overdo it or miss something.`
