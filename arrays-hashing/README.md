# Массивы и словари в Go

Здесь разбираю задачи, в которых часто помогают `map`, подсчёт частот и группировка данных.

## Задачи

| Задача | Основная идея |
|---|---|
| [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | Запоминать уже встреченные элементы в `map` |
| [Two Sum](https://leetcode.com/problems/two-sum/) | Хранить значение и индекс, искать дополнение до суммы |
| [Valid Anagram](https://leetcode.com/problems/valid-anagram/) | Сравнивать количество вхождений символов |
| [Group Anagrams](https://leetcode.com/problems/group-anagrams/) | Объединять слова по общему ключу |

Задачу [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) добавлю в список решённых после публикации кода и тестов.

## Проверка

Из корня репозитория запускаю:

```powershell
go test -count=1 -v ./...
```

Рядом с решениями лежат тесты. По возможности проверяю пустые данные, повторы и нестандартные случаи.
