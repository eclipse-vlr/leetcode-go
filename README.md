# LeetCode на Go

Мои решения алгоритмических задач для подготовки к собеседованиям Go Backend Developer.

## Как работаю

1. Сначала понимаю условие и разбираю примеры.
2. Предлагаю свой подход до написания кода.
3. Пишу решение на Go и тесты в `*_test.go`.
4. Проверяю граничные случаи и оцениваю Time / Space Complexity.
5. После решения коротко описываю идею и при необходимости пробую решить задачу заново.

## Arrays & Hash Map — первый этап

| № | LeetCode | Уровень | Статус |
|---|---|---|---|
| 217 | [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | Easy | Не начата |
| 1 | [Two Sum](https://leetcode.com/problems/two-sum/) | Easy | Не начата |
| 242 | [Valid Anagram](https://leetcode.com/problems/valid-anagram/) | Easy | Не начата |
| 49 | [Group Anagrams](https://leetcode.com/problems/group-anagrams/) | Medium | Не начата |
| 347 | [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) | Medium | Не начата |

## Структура

```text
leetcode-go/
├── arrays-hashing/
│   └── README.md
├── .gitignore
├── go.mod
└── README.md
```

Каждой задаче посвящены файлы `название.go` и `название_test.go` в папке соответствующей темы.
Готовые решения заранее не добавлены: цель — научиться решать самому.

## Тестирование

После добавления первых Go-файлов:

```sh
go test ./...
```
