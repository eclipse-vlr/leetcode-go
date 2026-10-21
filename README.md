# Алгоритмы на Go

Здесь я собираю решения задач с LeetCode и готовлюсь к техническим собеседованиям на Go Backend Developer.

Стараюсь не просто получить правильный ответ: сначала разбираю идею, затем пишу код и тесты, проверяю крайние случаи и оцениваю сложность решения.

Это учебный репозиторий. Основные backend-проекты — [LinkForge](https://github.com/eclipse-vlr/linkforge) и [PC Club Admin](https://github.com/eclipse-vlr/pc-club-admin).

## Решённые задачи

| Тема | Задача | Сложность |
|---|---|---|
| Arrays & Hashing | [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | Easy |
| Arrays & Hashing | [Two Sum](https://leetcode.com/problems/two-sum/) | Easy |
| Arrays & Hashing | [Valid Anagram](https://leetcode.com/problems/valid-anagram/) | Easy |
| Arrays & Hashing | [Group Anagrams](https://leetcode.com/problems/group-anagrams/) | Medium |

**В работе:** [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/). Не отмечаю задачу решённой, пока код и тесты не опубликованы в репозитории.

## Как проверяю решения

Из корня репозитория:

```powershell
go test -count=1 -v ./...
```

Для каждой задачи по возможности добавляю табличные тесты. Они помогают проверять не только обычный пример, но и пустые входные данные, повторяющиеся элементы и другие граничные случаи.

## Структура

Решения по массивам и словарям лежат в [`arrays-hashing/`](arrays-hashing/). Код и тесты находятся рядом, чтобы удобно было запускать и повторять задачи.

## Что изучаю

Хеш-таблицы, подсчёт частот, работу со срезами и строками, два указателя и скользящее окно. Постепенно добавляю новые темы по мере практики.
