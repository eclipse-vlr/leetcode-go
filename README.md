# Алгоритмы на Go

Здесь я собираю решения задач с LeetCode и готовлюсь к техническим собеседованиям на Go Backend Developer.

Сначала разбираю идею, затем пишу код и тесты, проверяю граничные случаи и оцениваю сложность алгоритма.

Это учебный репозиторий. Основные backend-проекты — [LinkForge](https://github.com/eclipse-vlr/linkforge) и [PC Club Admin](https://github.com/eclipse-vlr/pc-club-admin).

## Решённые задачи

| № | Задача | Тема | Сложность | Статус |
|---|---|---|---|---|
| 217 | [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | Arrays & Hashing | Easy | ✅ Решена |
| 1 | [Two Sum](https://leetcode.com/problems/two-sum/) | Arrays & Hashing | Easy | ✅ Решена |
| 242 | [Valid Anagram](https://leetcode.com/problems/valid-anagram/) | Arrays & Hashing | Easy | ✅ Решена |
| 49 | [Group Anagrams](https://leetcode.com/problems/group-anagrams/) | Arrays & Hashing | Medium | ✅ Решена |
| 347 | [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) | Arrays & Hashing | Medium | ✅ Решена |

**Прогресс базового блока Arrays & Hashing: 5 из 5.**

Для Top K Frequent Elements опубликовано решение через подсчёт частот и сортировку; вариант Bucket Sort за O(n) остаётся темой для повторения и оптимизации.

## План подготовки к собеседованиям в Яндекс

В репозитории есть отдельный список из **62 задач**, распределённых по темам: массивы, хеш-таблицы, два указателя, скользящее окно, стек, связные списки, деревья, графы и другие подходы.

➡️ **[Открыть список задач и статусы](YANDEX_ALGO_TASKS.md)**

Из этого списка на данный момент опубликованы решения **Two Sum**, **Valid Anagram** и **Group Anagrams**. Статусы в исходном PDF не отражают мой личный прогресс. Задачи Contains Duplicate и Top K Frequent Elements решены отдельно от подборки.

**Следующая тема — Two Pointers.** Для старта: [Move Zeroes (283)](https://leetcode.com/problems/move-zeroes/).

## Как проверяю решения

Из корня репозитория:

```powershell
go test -count=1 -v ./...
```

Для задач пишу табличные тесты и проверяю обычные и граничные случаи. Перед коммитом проверяю `git status`, затем отправляю изменения на GitHub. Сообщения коммитов пишу на английском.

## Структура

Решения по массивам и словарям лежат в [`arrays-hashing/`](arrays-hashing/). Код и тесты находятся рядом.

## Что изучаю

Hash Map, подсчёт частот, срезы и строки, Two Pointers, Sliding Window, Stack, Linked List, Binary Search, Trees, BFS/DFS и другие алгоритмические подходы.
