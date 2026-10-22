# Алгоритмы на Go

Решаю задачи с LeetCode и готовлюсь к техническим собеседованиям на Go Backend Developer. Для каждой задачи разбираю идею, реализацию на Go, тесты и сложность алгоритма.

## Решённые задачи

| № | Задача | Тема | Сложность |
|---|---|---|---|
| 217 | [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | Arrays & Hashing | Easy |
| 1 | [Two Sum](https://leetcode.com/problems/two-sum/) | Arrays & Hashing | Easy |
| 242 | [Valid Anagram](https://leetcode.com/problems/valid-anagram/) | Arrays & Hashing | Easy |
| 49 | [Group Anagrams](https://leetcode.com/problems/group-anagrams/) | Arrays & Hashing | Medium |
| 347 | [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) | Arrays & Hashing | Medium |
| 283 | [Move Zeroes](https://leetcode.com/problems/move-zeroes/) | Two Pointers | Easy |
| 125 | [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/) | Two Pointers | Easy |
| 167 | [Two Sum II](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) | Two Pointers | Medium |
| 977 | [Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/) | Two Pointers | Easy |

## Запуск тестов

```powershell
go test -count=1 -v ./...
```

Если Windows блокирует `vet.exe`, временно можно запустить тесты с `-vet=off`.
