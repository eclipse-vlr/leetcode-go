# LeetCode на Go

Решения алгоритмических задач на Go для подготовки к собеседованиям Go Backend Developer.

## Как работаю

1. Разбираю условие и примеры.
2. Придумываю алгоритм и оцениваю сложность.
3. Реализую решение на Go.
4. Пишу табличные тесты и проверяю граничные случаи.
5. Запускаю `go test -count=1 ./...` и отправляю решение на GitHub.

## Arrays & Hash Map

| № | Задача | Сложность | Статус |
|---|---|---|---|
| 217 | [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | Easy | ✅ Решена |
| 1 | [Two Sum](https://leetcode.com/problems/two-sum/) | Easy | ✅ Решена |
| 242 | [Valid Anagram](https://leetcode.com/problems/valid-anagram/) | Easy | ✅ Решена |
| 49 | [Group Anagrams](https://leetcode.com/problems/group-anagrams/) | Medium | ✅ Решена |
| 347 | [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) | Medium | ⏳ Следующая |

## Структура проекта

```text
leetcode-go/
├── arrays-hashing/
│   ├── README.md
│   ├── contains_duplicate.go
│   ├── contains_duplicate_test.go
│   ├── two_sum.go
│   ├── two_sum_test.go
│   ├── valid_anagram.go
│   ├── valid_anagram_test.go
│   ├── group_anagrams.go
│   └── group_anagrams_test.go
├── .gitignore
├── go.mod
└── README.md
```

## Запуск тестов

Из корня репозитория:

```sh
go test -count=1 -v ./...
```

## Изученные подходы

- Hash Set: поиск повторений за один проход
- Hash Map: хранение значений и индексов
- Frequency Counting: подсчёт частоты символов
- Grouping by Key: группировка анаграмм по отсортированному ключу

В задачах сохраняются исходные решения и тесты, чтобы можно было вернуться к ним и повторить алгоритмы.
